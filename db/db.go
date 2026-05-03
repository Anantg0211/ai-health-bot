package db

import (
	"ai-powered-health-bot/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conf := config.GetConfig()

	dsn := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v?charset=utf8mb4,utf8&parseTime=True`, conf.GetString("database.user"), conf.GetString("database.password"), conf.GetString("database.host"), conf.GetString("database.port"), conf.GetString("database.name"))
	if dsn == "" {
		return
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	DB = db
}