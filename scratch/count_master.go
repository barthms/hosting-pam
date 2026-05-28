package main

import (
	"fmt"
	"log"
	"monitoring-service/pkg/config"
	"monitoring-service/pkg/database"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed read config: %v", err)
	}

	cfg := config.NewConfig()

	db, err := database.GetConnection(cfg.Postgres().Read.ToArgs(database.Postgres, database.ReadConn, nil))
	if err != nil {
		log.Fatalf("failed connect db: %v", err)
	}

	var count int64
	if err := db.Table("master_standar_antropometri").Count(&count).Error; err != nil {
		log.Fatalf("count query failed: %v", err)
	}

	fmt.Printf("master_standar_antropometri count=%d\n", count)

	var sample []map[string]interface{}
	if err := db.Table("master_standar_antropometri").Limit(5).Order("parameter, jenis_kelamin, nilai_sumbu_x").Find(&sample).Error; err != nil {
		log.Fatalf("sample query failed: %v", err)
	}

	fmt.Printf("sample rows=%v\n", sample)
}
