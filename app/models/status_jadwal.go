package models

type StatusJadwal struct {
	ID         uint   `gorm:"column:id;primaryKey" json:"id"`
	NamaStatus string `gorm:"column:nama_status;type:text" json:"nama_status"`
}

func (StatusJadwal) TableName() string {
	return "status_jadwal"
}
