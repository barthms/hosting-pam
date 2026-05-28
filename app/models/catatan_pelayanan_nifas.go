// models/catatan_pelayanan_nifas.go
package models

import "time"

type CatatanPelayananNifas struct {
    IDCatatanNifas                  int32      `gorm:"primaryKey;autoIncrement" json:"id_catatan_nifas"`
    KehamilanID                     int32      `json:"kehamilan_id"`
    KunjunganKe                     string     `json:"kunjungan_ke"`
    TanggalPeriksaStampParaf        *time.Time `json:"tanggal_periksa_stamp_paraf"`
    KeluhanPemeriksaanTindakanSaran string     `json:"keluhan_pemeriksaan_tindakan_saran"`
    TanggalKembali                  *time.Time `json:"tanggal_kembali"`
}

func (CatatanPelayananNifas) TableName() string {
    return "catatan_pelayanan_nifas"
}