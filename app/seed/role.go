package seed

import (
	"log"

	"gorm.io/gorm"

	"monitoring-service/app/models"
)

func SeederRole(db *gorm.DB) error {
	log.Println("🌱 Seed: roles...")

	defaultRoles := []string{
		"Admin",
		"Tenaga-kesehatan",
		"Dokter",
		"Kader",
		"Bidan",
		"Ibu",
	}

	for _, roleName := range defaultRoles {
		role := models.Role{}

		err := db.Where(&models.Role{Name: roleName}).
			FirstOrCreate(&role, models.Role{
				Name: roleName,
			}).Error

		if err != nil {
			return err
		}

		log.Printf("✅ Role: %s", roleName)
	}

	log.Println("✅ Seeder role selesai")
	return nil
}
