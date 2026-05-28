package models

type TipeLembar string

const (
	TipeLembarA    TipeLembar = "A"
	TipeLembarB    TipeLembar = "B"
	TipeLembarUmum TipeLembar = "umum"
)

type StatusGizi string

const (
	StatusGiziBuruk    StatusGizi = "gizi_buruk"
	StatusGiziKurang   StatusGizi = "gizi_kurang"
	StatusGiziBaik     StatusGizi = "gizi_baik(normal)"
	StatusGiziLebih    StatusGizi = "gizi_lebih"
	StatusGiziObesitas StatusGizi = "obesitas"
)

// Role name constants
const (
	RoleNameIbu = "ibu"
	// RoleNameKaderPosyandu = "kader_posyandu"
	RoleNamePetugas = "petugas_kesehatan"
	RoleNameAdmin   = "admin"
)
