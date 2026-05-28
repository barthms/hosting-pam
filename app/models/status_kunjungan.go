package models

type StatusKunjungan struct {
	ID         uint   `gorm:"column:id;primaryKey" json:"id"`
	NamaStatus string `gorm:"column:nama_status;type:text" json:"nama_status"`
}

func (StatusKunjungan) TableName() string {
	return "status_kunjungan"
}

type StatusKunjunganCountResponse struct {
	StatusID        uint   `json:"status_id"`
	StatusKunjungan string `json:"status_kunjungan"`
	JumlahKunjungan int64  `json:"jumlah_kunjungan"`
}