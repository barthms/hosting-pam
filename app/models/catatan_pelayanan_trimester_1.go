package models

import "time"

type CatatanPelayananTrimester1 struct {
	IDCatatan                       int32      `gorm:"primaryKey" json:"id_catatan"`
	KehamilanID                     int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan                       *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID;onDelete:CASCADE" json:"kehamilan,omitempty"`
	TanggalPeriksaStampParaf        *time.Time `gorm:"type:date" json:"tanggal_periksa_stamp_paraf"`
	KeluhanPemeriksaanTindakanSaran string     `json:"keluhan_pemeriksaan_tindakan_saran"`
	TanggalKembali                  *time.Time `gorm:"type:date" json:"tanggal_kembali"`
}

func (CatatanPelayananTrimester1) TableName() string {
	return "catatan_pelayanan_trimester_1"
}
