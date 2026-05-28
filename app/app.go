package app

import (
	"fmt"
	"log"
	"monitoring-service/app/controllers"

	"time"

	// "monitoring-service/app/models"
	"monitoring-service/app/repositories"
	"monitoring-service/app/routes"

	// "monitoring-service/app/seed"

	"monitoring-service/app/usecases"
	"monitoring-service/pkg/config"
	"monitoring-service/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Main struct {
	cfg        *config.Config
	database   Database
	repo       *repositories.Main
	usecase    *usecases.Main
	controller *controllers.Main
	router     *echo.Echo
	cron       *cron.Cron
}

type Database struct {
	MySQL    *gorm.DB
	Postgres *gorm.DB
}

func New() *Main {
	return new(Main)
}
func (m *Main) startCronJob() {
	// Asumsikan usecase.KehamilanUsecase memiliki method UpdateAllActiveGestationalAge()
	kehamilanUC := m.usecase.Kehamilan // atau m.usecase.KehUsecase, sesuaikan dengan field di usecases.Main

	c := cron.New(cron.WithLocation(time.Local))
	// Jadwalkan setiap hari jam 01:00
	_, err := c.AddFunc("0 1 * * *", func() {
		log.Println("[CRON] Memulai update otomatis usia kehamilan...")
		if err := kehamilanUC.UpdateAllActiveGestationalAge(); err != nil {
			log.Printf("[CRON] Gagal update: %v", err)
		} else {
			log.Println("[CRON] Update usia kehamilan selesai.")
		}
	})
	if err != nil {
		log.Fatalf("[CRON] Gagal menjadwalkan job: %v", err)
	}

	// Jadwalkan pengingat imunisasi setiap hari jam 07:00
	_, err = c.AddFunc("0 7 * * *", func() {
		log.Println("[CRON] Memulai pengiriman pengingat imunisasi otomatis...")
		if err := m.usecase.SendScheduledImmunizationReminders(); err != nil {
			log.Printf("[CRON] Gagal mengirim pengingat imunisasi: %v", err)
		} else {
			log.Println("[CRON] Pengiriman pengingat imunisasi selesai.")
		}
	})
	if err != nil {
		log.Fatalf("[CRON] Gagal menjadwalkan job pengingat imunisasi: %v", err)
	}

	c.Start()
	m.cron = c
	log.Println("[CRON] Scheduler berjalan (usia kehamilan pukul 01:00, pengingat imunisasi pukul 07:00).")
}

func (m *Main) Init() (err error) {
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	m.cfg = config.NewConfig()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	m.database.Postgres, err = database.GetConnection(m.cfg.Postgres().Read.ToArgs(database.Postgres, database.ReadConn, nil))

	if err != nil {
		panic("❌ Gagal konek ke database: " + err.Error())
	}
	fmt.Println("✅ BERHASIL KONEK KE DATABASE")

	// Jalankan migrasi schema agar tabel baru selalu tersedia saat aplikasi start.
	// err = models.AutoMigrate(m.database.Postgres)
	// if err != nil {
	// 	return
	// }

	// Seeder
	// err = seed.RunAllSeed(m.database.Postgres)
	// if err != nil {
	// 	return
	// }
	// SEEDER setelah migrate
	// seeder kependudukan + anak
	// kependudukanSeeder := seeders.NewKependudukanSeeder(m.database.Postgres)
	// if err := kependudukanSeeder.Seed(); err != nil {
	// 	return err
	// }
	// kategoriTandaSakitSeeder := seeders.NewKategoriTandaSakitSeeder(m.database.Postgres)
	// if err := kategoriTandaSakitSeeder.Seed(); err != nil {
	// 	return err
	// }

	// // seeder master standar TBU
	// masterTBUSeeder := seeders.NewMasterStandarTBUSeeder(m.database.Postgres)
	// if err := masterTBUSeeder.Seed(); err != nil {
	// 	return err
	// }
	// masterBBTBSeeder := seeders.NewMasterStandarBBTBSeeder(m.database.Postgres)
	// if err := masterBBTBSeeder.Seed(); err != nil {
	// 	return err
	// }
	// masterBBUSeeder := seeders.NewMasterStandarBBUSeeder(m.database.Postgres)
	// if err := masterBBUSeeder.Seed(); err != nil {
	// 	return err
	// }
	// masterIMTUSeeder := seeders.NewMasterStandarIMTUSeeder(m.database.Postgres)
	// if err := masterIMTUSeeder.Seed(); err != nil {
	// 	return err
	// }
	// masterLKUSeeder := seeders.NewMasterStandarLKUSeeder(m.database.Postgres)
	// if err := masterLKUSeeder.Seed(); err != nil {
	// 	return err
	// }
	// kategoriCapaianSeeder := seeders.NewKategoriCapaianSeeder(m.database.Postgres)
	// if err := kategoriCapaianSeeder.Seed(); err != nil {
	// 	return err
	// }

	// SEEDER setelah migrate
	// seeder kependudukan + anak
	// kependudukanSeeder := seeders.NewKependudukanSeeder(m.database.Postgres)
	// if err := kependudukanSeeder.Seed(); err != nil {
	// 	return err
	// }

	// // seeder master standar TBU
	// masterTBUSeeder := seeders.NewMasterStandarTBUSeeder(m.database.Postgres)
	// if err := masterTBUSeeder.Seed(); err != nil {
	// 	return err
	// }
	// masterBBTBSeeder := seeders.NewMasterStandarBBTBSeeder(m.database.Postgres)
	// if err := masterBBTBSeeder.Seed(); err != nil {
	// 	return err
	// }
	// masterBBUSeeder := seeders.NewMasterStandarBBUSeeder(m.database.Postgres)
	// if err := masterBBUSeeder.Seed(); err != nil {
	// 	return err
	// }
	// masterIMTUSeeder := seeders.NewMasterStandarIMTUSeeder(m.database.Postgres)
	// if err := masterIMTUSeeder.Seed(); err != nil {
	// 	return err
	// }
	// masterLKUSeeder := seeders.NewMasterStandarLKUSeeder(m.database.Postgres)
	// if err := masterLKUSeeder.Seed(); err != nil {
	// 	return err
	// }
	// kategoriCapaianSeeder := seeders.NewKategoriCapaianSeeder(m.database.Postgres)
	// if err := kategoriCapaianSeeder.Seed(); err != nil {
	// 	return err
	// }
	// kategoriCapaianSeeder := seeders.NewKategoriCapaianSeeder(m.database.Postgres)
	// if err := kategoriCapaianSeeder.Seed(); err != nil {
	// 	return err
	// }

	m.repo = repositories.Init(repositories.Options{
		Config:   m.cfg,
		Postgres: m.database.Postgres,
	})
	m.usecase = usecases.Init(usecases.Options{
		Config:     m.cfg,
		Repository: m.repo,
	})
	m.controller = controllers.Init(controllers.Options{
		Config:   m.cfg,
		UseCases: m.usecase,
	})

	m.router = e
	go m.startCronJob()

	routes.ConfigureRouter(e, m.controller)
	return err
}

func (m *Main) Run() (err error) {
	defer m.close()

	m.router.Start(":" + m.cfg.ServicePort)
	return
}

func (m *Main) close() {
	if m.database.MySQL != nil {
		if db, err := m.database.MySQL.DB(); err == nil {
			db.Close()
		}
	}

	if m.database.Postgres != nil {
		if db, err := m.database.Postgres.DB(); err == nil {
			db.Close()
		}
	}
}
