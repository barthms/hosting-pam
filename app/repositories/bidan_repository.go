package repositories

import (
	"monitoring-service/app/models"
	"time"

	"gorm.io/gorm"
)

type BidanListItem struct {
	ID          int32     `gorm:"column:id" json:"id"`
	PendudukID  int32     `gorm:"column:penduduk_id" json:"penduduk_id"`
	NamaLengkap string    `gorm:"column:nama_lengkap" json:"nama_lengkap"`
	NIK         string    `gorm:"column:nik" json:"nik"`
	Kecamatan   string    `gorm:"column:kecamatan" json:"kecamatan"`
	Desa        string    `gorm:"column:desa" json:"desa"`
	NoSTR       string    `gorm:"column:no_str" json:"no_str"`
	NoSIPB      string    `gorm:"column:no_sipb" json:"no_sipb"`
	Status      string    `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type BidanRepository struct {
	db *gorm.DB
}

func NewBidanRepository(db *gorm.DB) *BidanRepository {
	return &BidanRepository{db: db}
}

func (r *BidanRepository) Create(data *models.Bidan) error {
	return r.db.Create(data).Error
}

func (r *BidanRepository) FindByID(id int32) (*models.Bidan, error) {
	var data models.Bidan
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&data).Error
	return &data, err
}

func (r *BidanRepository) FindByPendudukID(pendudukID int32) (*models.Bidan, error) {
	var data models.Bidan
	err := r.db.Where("penduduk_id = ? AND deleted_at IS NULL", pendudukID).First(&data).Error
	return &data, err
}

func (r *BidanRepository) FindAnyByPendudukID(pendudukID int32) (*models.Bidan, error) {
	var data models.Bidan
	err := r.db.Unscoped().Where("penduduk_id = ?", pendudukID).First(&data).Error
	return &data, err
}

func (r *BidanRepository) Update(data *models.Bidan) error {
	return r.db.Save(data).Error
}

func (r *BidanRepository) SetStatus(id int32, status string) error {
	now := time.Now()
	return r.db.Model(&models.Bidan{}).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": now,
		}).Error
}

func (r *BidanRepository) List(desa string) ([]BidanListItem, error) {
	var rows []BidanListItem

	q := r.db.Table("bidan b").
		Select("b.id, b.penduduk_id, p.nama_lengkap, p.nik, p.kecamatan, p.desa, b.no_str, b.no_sipb, b.status, b.created_at, b.updated_at").
		Joins("JOIN penduduk p ON p.id = b.penduduk_id").
		Where("b.deleted_at IS NULL AND p.deleted_at IS NULL").
		Order("b.id DESC")

	if desa != "" {
		q = q.Where("COALESCE(p.desa, '') = ?", desa)
	}

	err := q.Scan(&rows).Error
	return rows, err
}
