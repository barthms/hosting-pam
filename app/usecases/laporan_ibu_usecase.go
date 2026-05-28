// app/usecases/laporan_ibu_usecase.go
package usecases

import (
	"fmt"
	"time"

	"monitoring-service/app/models"
	"monitoring-service/app/repositories"

	"github.com/xuri/excelize/v2"
)

type LaporanIbuUsecase interface {
	ExportExcelLaporanIbu() (string, error)
	GetLaporanIbu() ([]models.LaporanIbu, error)
}

type laporanIbuUsecase struct {
	repo repositories.LaporanIbuRepository
}

func NewLaporanIbuUsecase(
	repo repositories.LaporanIbuRepository,
) LaporanIbuUsecase {
	return &laporanIbuUsecase{repo}
}

func (u *laporanIbuUsecase) GetLaporanIbu() ([]models.LaporanIbu, error) {
	return u.repo.GetLaporanIbu()
}

func (u *laporanIbuUsecase) ExportExcelLaporanIbu() (string, error) {

	data, err := u.repo.GetLaporanIbu()
	if err != nil {
		return "", err
	}

	f := excelize.NewFile()

	sheet := "Laporan Ibu"

	f.SetSheetName("Sheet1", sheet)

	headers := []string{
		"NIK",
		"Nama Ibu",
		"Nama Suami",
		"Tanggal Lahir",
		"HPHT",
		"HPL",
		"Usia Kehamilan",
		"Trimester",
		"Gravida",
		"Paritas",
		"Abortus",
		"BB Awal",
		"Tinggi Badan",
		"IMT",
		"LILA",
		"Tekanan Darah",
		"Sistole",
		"Diastole",
		"Tinggi Fundus",
		"Hb",
		"Golongan Darah",
		"Status Imunisasi",
		"Tripel Eliminasi",
		"Kunjungan ANC",
		"Tindakan",
		"Kecamatan",
		"Desa",
	}

	for i, h := range headers {
		cell := fmt.Sprintf("%c1", 65+i)
		f.SetCellValue(sheet, cell, h)
	}

	for i, d := range data {

		row := i + 2

		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), d.NIK)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), d.NamaIbu)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), d.NamaSuami)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), d.TanggalLahir.Format("2006-01-02"))
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), d.HPHT.Format("2006-01-02"))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), d.HPL.Format("2006-01-02"))
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), d.UsiaKehamilan)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), d.Trimester)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), d.Gravida)
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), d.Paritas)
		f.SetCellValue(sheet, fmt.Sprintf("K%d", row), d.Abortus)
		f.SetCellValue(sheet, fmt.Sprintf("L%d", row), d.BBAwal)
		f.SetCellValue(sheet, fmt.Sprintf("M%d", row), d.TinggiBadan)
		f.SetCellValue(sheet, fmt.Sprintf("N%d", row), d.IMT)
		f.SetCellValue(sheet, fmt.Sprintf("O%d", row), d.LILA)
		f.SetCellValue(sheet, fmt.Sprintf("P%d", row), d.TekananDarah)
		f.SetCellValue(sheet, fmt.Sprintf("Q%d", row), d.Sistole)
		f.SetCellValue(sheet, fmt.Sprintf("R%d", row), d.Diastole)
		f.SetCellValue(sheet, fmt.Sprintf("S%d", row), d.TinggiFundus)
		f.SetCellValue(sheet, fmt.Sprintf("T%d", row), d.Hb)
		f.SetCellValue(sheet, fmt.Sprintf("U%d", row), d.GolonganDarah)
		f.SetCellValue(sheet, fmt.Sprintf("V%d", row), d.StatusImunisasi)
		f.SetCellValue(sheet, fmt.Sprintf("W%d", row), d.TripelEliminasi)
		f.SetCellValue(sheet, fmt.Sprintf("X%d", row), d.KunjunganANC)
		f.SetCellValue(sheet, fmt.Sprintf("Y%d", row), d.Tindakan)
		f.SetCellValue(sheet, fmt.Sprintf("Z%d", row), d.Kecamatan)
		f.SetCellValue(sheet, fmt.Sprintf("AA%d", row), d.Desa)
	}

	filename := fmt.Sprintf(
		"laporan_ibu_%s.xlsx",
		time.Now().Format("20060102_150405"),
	)

	err = f.SaveAs(filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}