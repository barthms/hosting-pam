package repositories

import (
	"testing"
	"time"

	"monitoring-service/app/models"

	"gorm.io/gorm"
)

// Mock Database untuk testing
func createMockDB(t *testing.T) *gorm.DB {
	// Ini adalah placeholder - dalam praktik nyata, gunakan testcontainers atau SQLite in-memory
	db, err := gorm.Open(gorm.Dialector(nil), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	return db
}

// Test CreatePerawatan
func TestCreatePerawatan(t *testing.T) {
	db := createMockDB(t)
	repo := NewPerawatanRepository(db)

	perawatan := &models.Perawatan{
		AnakID:            1,
		KategoriCapaianID: 1,
		Jawaban:           boolPtr(true),
		TanggalPeriksa:    timePtr(time.Now()),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err := repo.CreatePerawatan(perawatan)
	if err != nil {
		t.Errorf("CreatePerawatan() failed: %v", err)
	}

	if perawatan.ID == 0 {
		t.Error("CreatePerawatan() did not set ID")
	}
}

// Test GetPerawatanByID - not found
func TestGetPerawatanByID_NotFound(t *testing.T) {
	db := createMockDB(t)
	repo := NewPerawatanRepository(db)

	result, err := repo.GetPerawatanByID(999)
	if err == nil {
		t.Error("Expected error for non-existent ID")
	}
	if result != nil {
		t.Error("Expected nil result for non-existent ID")
	}
}

// Test GetKategoriCapaianByRentangUsia
func TestGetKategoriCapaianByRentangUsia(t *testing.T) {
	db := createMockDB(t)
	repo := NewPerawatanRepository(db)

	_, err := repo.GetKategoriCapaianByRentangUsia("0-12 Bulan")
	// Tidak ada error jika database kosong
	if err != nil && err.Error() != "gagal mengambil kategori capaian" {
		t.Errorf("Unexpected error: %v", err)
	}
}

// Helper functions
func boolPtr(b bool) *bool {
	return &b
}

func timePtr(t time.Time) *time.Time {
	return &t
}
