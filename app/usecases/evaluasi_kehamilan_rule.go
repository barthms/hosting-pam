package usecases

import (
	"fmt"
	"monitoring-service/app/models"
	"strings"
)

func GeneratePenjelasanSingle(
	current *models.GrafikEvaluasiKehamilan,
	history []models.GrafikEvaluasiKehamilan,
) (string, string) {

	var janinStatus []string
	var ibuStatus []string
	var suplemenInfo []string
	score := 0

	// ======================
	// 1. ANALISIS JANIN
	// ======================
	
	// TFU vs USIA
	if current.UsiaGestasiMinggu != nil && current.TinggiFundusUteriCm != nil {
		usia := float64(*current.UsiaGestasiMinggu)
		tfu := *current.TinggiFundusUteriCm
		if usia >= 20 {
			selisih := tfu - usia
			if selisih < -2 {
				janinStatus = append(janinStatus, "TFU rendah (Waspada Pertumbuhan Janin Terhambat)")
				score += 2
			} else if selisih > 2 {
				janinStatus = append(janinStatus, "TFU tinggi (Perlu cek cairan ketuban/bayi besar)")
				score += 1
			} else {
				janinStatus = append(janinStatus, "Pertumbuhan TFU normal")
			}
		}
	}

	// Trend TFU (History)
	if len(history) > 0 {
		last := history[len(history)-1]
		if last.TinggiFundusUteriCm != nil && current.TinggiFundusUteriCm != nil {
			if *current.TinggiFundusUteriCm <= *last.TinggiFundusUteriCm {
				janinStatus = append(janinStatus, "⚠️ Grafik TFU mendatar/turun (Risiko gawat janin)")
				score += 3
			}
		}
	}

	// DJJ
	if current.DenyutJantungBayiXMenit != nil {
		djj := *current.DenyutJantungBayiXMenit
		if djj < 110 || djj > 160 {
			janinStatus = append(janinStatus, fmt.Sprintf("DJJ tidak normal: %d bpm (Gawat Janin)", djj))
			score += 4
		} else {
			janinStatus = append(janinStatus, "Detak jantung normal")
		}
	}

	// Gerakan Janin
	if current.GerakanBayi != nil {
		g := strings.ToLower(*current.GerakanBayi)
		if g == "kurang" || g == "-" || g == "tidak ada" {
			janinStatus = append(janinStatus, "Gerakan janin berkurang")
			score += 4
		} else {
			janinStatus = append(janinStatus, "Gerakan aktif")
		}
	}

	// ======================
	// 2. ANALISIS IBU
	// ======================
	
	// Tekanan Darah
	hipertensi := false
	if current.TekananDarahSistole != nil && current.TekananDarahDiastole != nil {
		s := *current.TekananDarahSistole
		d := *current.TekananDarahDiastole
		
		if s >= 140 || d >= 90 {
			ibuStatus = append(ibuStatus, fmt.Sprintf("Hipertensi (%d/%d)", s, d))
			score += 3
			hipertensi = true
		} else if s >= 130 || d >= 80 {
			ibuStatus = append(ibuStatus, fmt.Sprintf("Tekanan darah meningkat (%d/%d) - Perlu waspada", s, d))
			score += 1
		} else {
			ibuStatus = append(ibuStatus, "Tekanan darah normal")
		}
	}

	// Urin & Preeklampsia
	proteinPositif := false
	if current.UrinProtein != nil {
		p := strings.ToLower(*current.UrinProtein)
		if strings.Contains(p, "+") || strings.Contains(p, "positif") {
			ibuStatus = append(ibuStatus, "Protein urin positif (+)")
			score += 2
			proteinPositif = true
		}
	}
	if current.UrinReduksi != nil {
		r := strings.ToLower(*current.UrinReduksi)
		if strings.Contains(r, "+") || strings.Contains(r, "positif") {
			ibuStatus = append(ibuStatus, "Reduksi urin positif (Waspada Diabetes)")
			score += 1
		}
	}

	if hipertensi && proteinPositif {
		ibuStatus = append(ibuStatus, "‼️ Indikasi kuat Preeklampsia (Segera Rujuk)")
		score += 4
	}

	// Hemoglobin
	if current.Hemoglobin != nil {
		hb := *current.Hemoglobin
		if hb < 11 {
			ibuStatus = append(ibuStatus, fmt.Sprintf("Anemia (Hb rendah: %.1f)", hb))
			score += 2
		} else {
			ibuStatus = append(ibuStatus, fmt.Sprintf("Hb normal (%.1f)", hb))
		}
	}

	// ======================
	// 3. KEPATUHAN & RISK
	// ======================
	if current.TabletTambahDarah != nil && *current.TabletTambahDarah > 0 {
		suplemenInfo = append(suplemenInfo, fmt.Sprintf("Fe (%d butir)", *current.TabletTambahDarah))
	}
	if current.Kalsium != nil && strings.ToLower(*current.Kalsium) == "ya" {
		suplemenInfo = append(suplemenInfo, "Kalsium")
	}
	if current.Aspirin != nil && strings.ToLower(*current.Aspirin) == "ya" {
		suplemenInfo = append(suplemenInfo, "Aspirin")
	}

	// Penentuan Risk Level
	risk := "Rendah"
	if score >= 7 {
		risk = "Tinggi"
	} else if score >= 3 {
		risk = "Sedang"
	}

	// Gabungkan hasil menjadi narasi yang rapi
	var finalMessage []string
	if len(janinStatus) > 0 {
		finalMessage = append(finalMessage, "[JANIN]: "+strings.Join(janinStatus, ", "))
	}
	if len(ibuStatus) > 0 {
		finalMessage = append(finalMessage, "[IBU]: "+strings.Join(ibuStatus, ", "))
	}
	if len(suplemenInfo) > 0 {
		finalMessage = append(finalMessage, "[SUPLEMEN]: "+strings.Join(suplemenInfo, ", "))
	}

	if len(finalMessage) == 0 {
		return "Hasil evaluasi dalam batas normal.", "Rendah"
	}

	return strings.Join(finalMessage, " | ") + ".", risk
}