package main

// import (
//     "log"
//     "monitoring-service/app/seeders"
//     "monitoring-service/pkg/config"
//     "monitoring-service/pkg/database"
//     "github.com/spf13/viper"
// )

// func main() {
//     viper.SetConfigFile(".env")
//     if err := viper.ReadInConfig(); err != nil {
//         log.Fatalf("failed read config: %v", err)
//     }

//     cfg := config.NewConfig()

//     db, err := database.GetConnection(cfg.Postgres().Read.ToArgs(database.Postgres, database.ReadConn, nil))
//     if err != nil {
//         log.Fatalf("failed connect db: %v", err)
//     }

//     log.Println("Seeding master standar BBU...")
//     if err := seeders.NewMasterStandarBBUSeeder(db).Seed(); err != nil {
//         log.Fatalf("BBU seeding failed: %v", err)
//     }

//     log.Println("Seeding master standar TBU...")
//     if err := seeders.NewMasterStandarTBUSeeder(db).Seed(); err != nil {
//         log.Fatalf("TBU seeding failed: %v", err)
//     }

//     log.Println("Seeding master standar BBTB...")
//     if err := seeders.NewMasterStandarBBTBSeeder(db).Seed(); err != nil {
//         log.Fatalf("BBTB seeding failed: %v", err)
//     }

//     log.Println("Seeding master standar IMTU...")
//     if err := seeders.NewMasterStandarIMTUSeeder(db).Seed(); err != nil {
//         log.Fatalf("IMTU seeding failed: %v", err)
//     }

//     log.Println("Seeding master standar LKU...")
//     if err := seeders.NewMasterStandarLKUSeeder(db).Seed(); err != nil {
//         log.Fatalf("LKU seeding failed: %v", err)
//     }

//     log.Println("All master standar seeded successfully")
// }
