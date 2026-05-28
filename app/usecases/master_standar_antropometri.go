package usecases

// import (
// 	"math"
// 	"monitoring-service/app/models"
// 	"monitoring-service/pkg/customerror"
// 	"strings"
// )

// func (m *Main) GetMasterStandar(parameter, jenisKelamin string) ([]models.MasterStandarResponse, error) {
// 	rows, err := m.repository.GetMasterStandarByFilter(parameter, jenisKelamin)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := make([]models.MasterStandarResponse, 0, len(rows))
// 	for _, row := range rows {
// 		item := models.MasterStandarResponse{
// 			ID:           row.ID,
// 			Parameter:    row.Parameter,
// 			JenisKelamin: string(row.JenisKelamin),
// 			NilaiSumbuX:  row.NilaiSumbuX,
// 			SD3Neg:       row.SD3Neg,
// 			SD2Neg:       row.SD2Neg,
// 			SD1Neg:       row.SD1Neg,
// 			Median:       row.Median,
// 			SD1Pos:       row.SD1Pos,
// 			SD2Pos:       row.SD2Pos,
// 			SD3Pos:       row.SD3Pos,
// 		}

// 		if row.Parameter == ParamLKU {
// 			item.UsiaBulan = int(math.Round(row.NilaiSumbuX))
// 			item.UsiaYM = formatBulanToYM(row.NilaiSumbuX)
// 		}

// 		result = append(result, item)
// 	}

// 	return result, nil
// }

// func (m *Main) CreateMasterStandar(req *models.CreateMasterStandarRequest) error {
// 	if strings.TrimSpace(req.Parameter) == "" {
// 		return customerror.NewBadRequestError("parameter wajib diisi")
// 	}

// 	nilaiSumbuX := req.NilaiSumbuX
// 	if strings.TrimSpace(req.Parameter) == ParamLKU && strings.TrimSpace(req.UsiaReferensi) != "" {
// 		parsed, err := parseUsiaReferensiToBulan(req.UsiaReferensi)
// 		if err != nil {
// 			return err
// 		}
// 		nilaiSumbuX = parsed
// 	}

// 	model := &models.MasterStandarAntropometri{
// 		Parameter:    req.Parameter,
// 		JenisKelamin: models.GenderType(sanitizeGender(req.JenisKelamin)),
// 		NilaiSumbuX:  nilaiSumbuX,
// 		SD3Neg:       req.SD3Neg,
// 		SD2Neg:       req.SD2Neg,
// 		SD1Neg:       req.SD1Neg,
// 		Median:       req.Median,
// 		SD1Pos:       req.SD1Pos,
// 		SD2Pos:       req.SD2Pos,
// 		SD3Pos:       req.SD3Pos,
// 	}

// 	return m.repository.CreateMasterStandar(model)
// }
