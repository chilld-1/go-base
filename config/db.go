package config

import (
	"log"
	"time"
	"xiangmu/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {

	dsn := AppConfig.Database.Dsn
	log.Printf("Database DSN: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database ,got error: %v", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns) //空闲连接数量
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns) // open max
	sqlDB.SetConnMaxLifetime(time.Hour)                    //max time
	if err != nil {
		log.Fatalf("Failed to config database ,got error: %v", err)
	}
	global.Db = db
}
