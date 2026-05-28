package usecases

// import (
// 	"strconv"
// 	"strings"

// 	"monitoring-service/app/models"
// 	"monitoring-service/pkg/customerror"
// )

// func mapAnakToResponse(data models.Anak) models.AnakResponse {
// 	res := models.AnakResponse{
// 		ID:              data.ID,
// 		IbuID:           data.IbuID,
// 		KependudukanID:  data.KependudukanID,
// 		NoKartuKeluarga: data.NoKartuKeluarga,
// 		NamaAnak:        data.NamaAnak,
// 		JenisKelamin:    data.JenisKelamin,
// 		TanggalLahir:    data.TanggalLahir,
// 		BeratLahir:      data.BeratLahir,
// 		TinggiLahir:     data.TinggiLahir,
// 	}

// 	if strings.TrimSpace(res.NamaAnak) == "" && data.Kependudukan != nil {
// 		res.NamaAnak = data.Kependudukan.Nama
// 	}
// 	if strings.TrimSpace(res.JenisKelamin) == "" && data.Kependudukan != nil {
// 		res.JenisKelamin = data.Kependudukan.JenisKelamin
// 	}
// 	if strings.TrimSpace(res.TanggalLahir) == "" && data.Kependudukan != nil {
// 		res.TanggalLahir = data.Kependudukan.TanggalLahir
// 	}
// 	if res.NoKartuKeluarga == 0 && data.Kependudukan != nil && data.Kependudukan.NoKartuKeluarga != nil {
// 		res.NoKartuKeluarga = data.Kependudukan.NoKartuKeluarga.NoKartuKeluarga
// 	}

// 	return res
// }

// func mapAnakListToResponse(data []models.Anak) []models.AnakResponse {
// 	res := make([]models.AnakResponse, 0, len(data))
// 	for _, val := range data {
// 		res = append(res, mapAnakToResponse(val))
// 	}
// 	return res
// }

// func (m *Main) GetAllAnak() ([]models.AnakResponse, error) {
// 	data, err := m.repository.GetAllAnak()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return mapAnakListToResponse(data), nil
// }

// func (m *Main) GetAnakById(anakID string) (*models.AnakResponse, error) {
// 	id, err := strconv.ParseUint(strings.TrimSpace(anakID), 10, 64)
// 	if err != nil || id == 0 {
// 		return nil, customerror.NewBadRequestError("anak_id tidak valid")
// 	}

// 	data, err := m.repository.GetAnakByID(uint(id))
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := mapAnakToResponse(*data)
// 	return &res, nil
// }

// func (m *Main) GetAnak(namaAnak, namaIbu, noKK string) ([]models.AnakResponse, *models.Pagination, error) {
// 	namaAnak = strings.TrimSpace(namaAnak)
// 	namaIbu = strings.TrimSpace(namaIbu)
// 	noKK = strings.TrimSpace(noKK)

// 	if namaAnak == "" && namaIbu == "" && noKK == "" {
// 		return nil, nil, customerror.NewBadRequestError("minimal satu parameter pencarian wajib diisi: nama, nama_ibu, atau no_kk")
// 	}

// 	data, err := m.repository.SearchAnak(namaAnak, namaIbu, noKK)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	res := mapAnakListToResponse(data)
// 	meta := &models.Pagination{
// 		Total: len(res),
// 	}

// 	return res, meta, nil
// }
