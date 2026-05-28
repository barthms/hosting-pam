package usecases

import (
	// "fmt"
	"monitoring-service/app/models"
)

func (m *Main) GetJumlahKunjunganByStatus() (
	[]models.StatusKunjunganCountResponse,
	error,
) {

	rows, err :=
		m.repository.
			GetJumlahKunjunganByStatus()

	if err != nil {
		return nil, err
	}

	response :=
		[]models.StatusKunjunganCountResponse{}

	for _, row := range rows {

		response =
			append(
				response,
				models.StatusKunjunganCountResponse{
					StatusID:        row.StatusID,
					StatusKunjungan: row.StatusKunjungan,
					JumlahKunjungan: row.JumlahKunjungan,
				},
			)
	}

	return response, nil
}