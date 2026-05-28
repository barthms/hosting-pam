package usecases

import (
	"errors"
	"fmt"
	"monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"strconv"
	"strings"
	"time"
)

var ErrWarnaTinjaForbidden = errors.New("akses warna tinja ditolak")

type WarnaTinjaUsecase interface {
	GetByAnakIDForIbu(anakID string, userID uint) ([]models.WarnaTinjaAnak, error)
	SaveForIbu(req models.WarnaTinjaSaveRequest, userID uint) (*models.WarnaTinjaAnak, error)
}

type warnaTinjaUsecase struct {
	repo repositories.WarnaTinjaRepository
}

func NewWarnaTinjaUsecase(repo repositories.WarnaTinjaRepository) WarnaTinjaUsecase {
	return &warnaTinjaUsecase{repo: repo}
}

func parseAnakID(anakID string) (uint, error) {
	anakID = strings.TrimSpace(anakID)
	if anakID == "" {
		return 0, errors.New("anak_id wajib diisi")
	}
	val, err := strconv.ParseUint(anakID, 10, 32)
	if err != nil || val <= 0 {
		return 0, fmt.Errorf("anak_id harus berupa angka lebih dari 0")
	}
	return uint(val), nil
}

func (u *warnaTinjaUsecase) GetByAnakIDForIbu(anakID string, userID uint) ([]models.WarnaTinjaAnak, error) {
	id, err := parseAnakID(anakID)
	if err != nil {
		return nil, err
	}

	ok, err := u.repo.IsAnakMilikIbu(userID, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrWarnaTinjaForbidden
	}

	return u.repo.FindByAnakID(id)
}

func (u *warnaTinjaUsecase) SaveForIbu(req models.WarnaTinjaSaveRequest, userID uint) (*models.WarnaTinjaAnak, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ok, err := u.repo.IsAnakMilikIbu(userID, req.AnakID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrWarnaTinjaForbidden
	}

	tanggal, err := time.Parse("2006-01-02", req.TanggalCatat)
	if err != nil {
		return nil, errors.New("format tanggal_catat harus YYYY-MM-DD")
	}

	data := &models.WarnaTinjaAnak{
		AnakID:       req.AnakID,
		PeriodeKey:   req.PeriodeKey,
		PeriodeLabel: models.WarnaTinjaLabelForPeriode(req.PeriodeKey),
		TanggalCatat: &tanggal,
		NomorWarna:   req.NomorWarna,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := u.repo.Upsert(data); err != nil {
		return nil, err
	}

	return data, nil
}
