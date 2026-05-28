package repositories

import (
	"monitoring-service/app/models"
	"strings"
	"time"

	"gorm.io/gorm"
)

type KartuKeluargaRepository struct {
	db *gorm.DB
}

func NewKartuKeluargaRepository(db *gorm.DB) *KartuKeluargaRepository {
	return &KartuKeluargaRepository{db: db}
}

func (r *KartuKeluargaRepository) Create(kk *models.KartuKeluarga) error {
	return r.db.Create(kk).Error
}

func (r *KartuKeluargaRepository) FindByNoKK(noKK string) (*models.KartuKeluarga, error) {
	var kk models.KartuKeluarga
	err := r.db.Where("no_kk = ? AND deleted_at IS NULL", noKK).First(&kk).Error
	return &kk, err
}

func (r *KartuKeluargaRepository) FindByID(id int64) (*models.KartuKeluarga, error) {
	var kk models.KartuKeluarga
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&kk).Error
	return &kk, err
}

func (r *KartuKeluargaRepository) FindByNoKKExceptID(noKK string, exceptID int64) (*models.KartuKeluarga, error) {
	var kk models.KartuKeluarga
	err := r.db.Where("no_kk = ? AND id <> ? AND deleted_at IS NULL", noKK, exceptID).First(&kk).Error
	return &kk, err
}

func (r *KartuKeluargaRepository) Update(kk *models.KartuKeluarga) error {
	return r.db.Save(kk).Error
}

func (r *KartuKeluargaRepository) SoftDeleteByID(id int64) error {
	now := time.Now()
	return r.db.Model(&models.KartuKeluarga{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"deleted_at": now,
			"updated_at": now,
		}).Error
}

func (r *KartuKeluargaRepository) ListPaginated(search string, page int, limit int, sortBy string, sortDir string) ([]models.KartuKeluarga, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	allowedSortBy := map[string]string{
		"created_at": "created_at",
		"updated_at": "updated_at",
		"no_kk":      "no_kk",
	}
	column, ok := allowedSortBy[strings.ToLower(strings.TrimSpace(sortBy))]
	if !ok {
		column = "created_at"
	}

	direction := strings.ToUpper(strings.TrimSpace(sortDir))
	if direction != "ASC" {
		direction = "DESC"
	}

	query := r.db.Model(&models.KartuKeluarga{}).Where("deleted_at IS NULL")

	if trimmed := strings.TrimSpace(search); trimmed != "" {
		like := "%" + trimmed + "%"
		query = query.Where(`
			no_kk ILIKE ? OR EXISTS (
				SELECT 1
				FROM penduduk p
				WHERE p.kartu_keluarga_id = kartu_keluarga.id
				  AND p.deleted_at IS NULL
				  AND (p.nik ILIKE ? OR p.nama_lengkap ILIKE ?)
			)
		`, like, like, like)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var list []models.KartuKeluarga
	err := query.
		Order(column + " " + direction).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&list).Error

	return list, total, err
}
