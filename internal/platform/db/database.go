package db

import (
	"fmt"

	"github.com/Lakasha-hub/micro-yoga-mp.git/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg config.Config) {
	dsn := cfg.MySQLDSN

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connected")
	// TODO: Create models
	DB.AutoMigrate()
}
