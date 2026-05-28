package models

import "time"

type CatatanPelayananTrimester3 struct {
	IDCatatanT3                     int32      `gorm:"primaryKey" json:"id_catatan_t3"`
	KehamilanID                     int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan                       *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID;onDelete:CASCADE" json:"kehamilan,omitempty"`
	TanggalPeriksaStampParaf        *time.Time `gorm:"type:date" json:"tanggal_periksa_stamp_paraf"`
	KeluhanPemeriksaanTindakanSaran string     `json:"keluhan_pemeriksaan_tindakan_saran"`
	TanggalKembali                  *time.Time `gorm:"type:date" json:"tanggal_kembali"`
}

func (CatatanPelayananTrimester3) TableName() string {
	return "catatan_pelayanan_trimester_3"
}
