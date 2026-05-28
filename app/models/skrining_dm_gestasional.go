package models

type SkriningDMGestasional struct {
	IDSkriningDM int32      `gorm:"primaryKey" json:"id_skrining_dm"`
	KehamilanID  int32      `gorm:"not null;index" json:"kehamilan_id"`
	Kehamilan    *Kehamilan `gorm:"foreignKey:KehamilanID;references:ID" json:"kehamilan,omitempty"`

	GulaDarahPuasaHasil              string `json:"gula_darah_puasa_hasil"`
	GulaDarahPuasaRencana            string `json:"gula_darah_puasa_rencana_tindak_lanjut"`
	GulaDarah2JamPostPrandialHasil   string `json:"gula_darah_2_jam_post_prandial_hasil"`
	GulaDarah2JamPostPrandialRencana string `json:"gula_darah_2_jam_post_prandial_rencana_tindak_lanjut"`
}

func (SkriningDMGestasional) TableName() string {
	return "skrining_dm_gestasional"
}
