package models

import (
	"time"

	"gorm.io/gorm"
)

// type DeteksiDiniPenyimpangan struct {
// 	ID                string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
// 	AnakID            string         `json:"anak_id" gorm:"type:varchar(36);not null;index"`
// 	Anak              *Anak          `json:"anak,omitempty" gorm:"foreignKey:AnakID"`
// 	TenagaKesehatanID string         `json:"tenaga_kesehatan_id" gorm:"type:varchar(36);not null;index"`
// 	TenagaKesehatan   *User          `json:"tenaga_kesehatan,omitempty" gorm:"foreignKey:TenagaKesehatanID"`
// 	Bulanke           int            `json:"bulan" gorm:"not null"`
// 	Tanggal           time.Time      `json:"tanggal" gorm:"not null"`
// 	BBperU            string         `json:"bb_u" gorm:"type:enum('SK','K','N','RBBL');not null"`
// 	BBpertb           string         `json:"bb_t" gorm:"type:enum('G B','GK','GN','RGL','Gl','O');not null"`
// 	TBperU            string         `json:"tb_u" gorm:"type:enum('SP','P','N','Ti');not null"`
// 	LkperU            string         `json:"lk_u" gorm:"type:enum('Mi','N','Ma');not null"`
// 	LilA              string         `json:"lila" gorm:"type:enum('GB','GK','N');not null"`
// 	KPSP              string         `json:"kpsp" gorm:"type:enum('Ds','Dm','Dp');not null"`
// 	TDD               string         `json:"tdd" gorm:"type:enum('N','R');not null"`
// 	TDL               string         `json:"tdl" gorm:"type:enum('N','R');not null"`
// 	KMPE              string         `json:"kmpe" gorm:"type:enum('N','R');not null"`
// 	M_CHAT_REVISED    string         `json:"m_chat_revised" gorm:"type:enum('N','R');not null"`
// 	ACTRS             string         `json:"actrs" gorm:"type:enum('N','R');not null"`
// 	HasilPKAT         string         `json:"hasil_pkat" gorm:"not null"`
// 	Tindakan          string         `json:"tindakan" gorm:"not null"`
// 	KunjunganUlang    time.Time      `json:"kunjungan_ulang" gorm:"not null"`
// 	CreatedAt         time.Time      `json:"created_at"`
// 	UpdatedAt         time.Time      `json:"updated_at"`
// 	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
// }
type DeteksiDiniPenyimpangan struct {
	ID                int32         `json:"id" gorm:"primaryKey;autoIncrement"`

	AnakID            int32         `json:"anak_id" gorm:"not null;index"`
	Anak              *Anak          `json:"anak,omitempty" gorm:"foreignKey:AnakID"`

	TenagaKesehatanID int32         `json:"tenaga_kesehatan_id" gorm:"not null;index"`
	TenagaKesehatan   *User          `json:"tenaga_kesehatan,omitempty" gorm:"foreignKey:TenagaKesehatanID"`

	BulanKe           int            `json:"bulan_ke" gorm:"not null"`
	Tanggal           time.Time      `json:"tanggal" gorm:"not null"`

	BBperU            string         `json:"bb_u" gorm:"type:varchar(10);not null"`
	BBperTB           string         `json:"bb_tb" gorm:"type:varchar(10);not null"`
	TBperU            string         `json:"tb_u" gorm:"type:varchar(10);not null"`
	LKperU            string         `json:"lk_u" gorm:"type:varchar(10);not null"`
	LILA              string         `json:"lila" gorm:"type:varchar(10);not null"`

	KPSP              string         `json:"kpsp" gorm:"type:varchar(10);not null"`
	TDD               string         `json:"tdd" gorm:"type:varchar(5);not null"`
	TDL               string         `json:"tdl" gorm:"type:varchar(5);not null"`
	KMPE              string         `json:"kmpe" gorm:"type:varchar(5);not null"`
	MCHATRevised      string         `json:"m_chat_revised" gorm:"type:varchar(5);not null"`
	ACTRS             string         `json:"actrs" gorm:"type:varchar(5);not null"`

	HasilPKAT         string         `json:"hasil_pkat" gorm:"type:text;not null"`
	Tindakan          string         `json:"tindakan" gorm:"type:text;not null"`

	KunjunganUlang    time.Time      `json:"kunjungan_ulang" gorm:"not null"`

	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func (DeteksiDiniPenyimpangan) TableName() string {
	return "deteksi_dini_penyimpangan"
}

// func (D *DeteksiDiniPenyimpangan) BeforeCreate(tx *gorm.DB) error {
// 	if D.ID == "" {
// 		D.ID = uuid.New().String()
// 	}
// 	return nil
// }
