package models

import "time"

type CatatanPelayananTrimester2 struct {
	IDCatatanT2                     int32      `gorm:"primaryKey" json:"id_catatan_t2"`
	KehamilanID                     int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan                       *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`
	TanggalPeriksaStampParaf        *time.Time `gorm:"type:date" json:"tanggal_periksa_stamp_paraf"`
	KeluhanPemeriksaanTindakanSaran string     `json:"keluhan_pemeriksaan_tindakan_saran"`
	TanggalKembali                  *time.Time `gorm:"type:date" json:"tanggal_kembali"`
}

func (CatatanPelayananTrimester2) TableName() string {
	return "catatan_pelayanan_trimester_2"
}
