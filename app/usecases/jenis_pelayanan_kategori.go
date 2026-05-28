package usecases

import "monitoring-service/app/repositories"

// =====================
// INTERFACE
// =====================
type JenisPelayananUsecase interface {
	GetJenisPelayanan(kategoriID, periodeID int32) ([]JenisPelayananResponse, error)
}

// =====================
// STRUCT
// =====================
type jenisPelayananUsecase struct {
	repo repositories.JenisPelayananRepository
}

// =====================
// CONSTRUCTOR
// =====================
func NewJenisPelayananUsecase(repo repositories.JenisPelayananRepository) JenisPelayananUsecase {
	return &jenisPelayananUsecase{repo}
}

// =====================
// RESPONSE
// =====================
type JenisPelayananResponse struct {
	JenisPelayananID int32  `json:"jenis_pelayanan_id"`
	Nama             string `json:"nama"`
	TipeInput        string `json:"tipe_input"`
	GroupName        string `json:"group_name"`
	Section          string `json:"section"`
}

// =====================
// METHOD
// =====================
func (u *jenisPelayananUsecase) GetJenisPelayanan(kategoriID, periodeID int32) ([]JenisPelayananResponse, error) {
	data, err := u.repo.GetByKategoriAndPeriode(kategoriID, periodeID)
	if err != nil {
		return nil, err
	}

	var result []JenisPelayananResponse

	for _, d := range data {
		result = append(result, JenisPelayananResponse{
			JenisPelayananID: d.ID,
			Nama:             d.Nama,
			TipeInput:        d.TipeInput,
			GroupName:        d.GroupName,
			Section:          d.Section,
		})
	}

	return result, nil
}