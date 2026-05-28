package main

// import (
// 	"fmt"
// 	"log"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Penduduk struct {
// 	ID           int32
// 	NamaLengkap  string
// 	JenisKelamin string
// }

// type Anak struct {
// 	ID         int32
// 	PendudukID int32
// 	Penduduk   *Penduduk `gorm:"foreignKey:PendudukID"`
// }

// func main() {
// 	dsn := "host=localhost user=postgres password=postgres dbname=kia_db port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var anak Anak
// 	if err := db.Preload("Penduduk").First(&anak, 21).Error; err != nil {
// 		fmt.Printf("Error finding anak: %v\n", err)
// 	} else {
// 		fmt.Printf("Anak ID: %d, Nama: %s, Gender: %s\n", anak.ID, anak.Penduduk.NamaLengkap, anak.Penduduk.JenisKelamin)
// 	}

// 	var count int64
// 	db.Table("master_standar_antropometri").Count(&count)
// 	fmt.Printf("Total Master Standar: %d\n", count)
	
// 	var samples []map[string]interface{}
// 	db.Table("master_standar_antropometri").Limit(5).Find(&samples)
// 	fmt.Printf("Sample standards: %v\n", samples)
// }
