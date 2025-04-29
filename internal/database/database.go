package database

import (
	"fmt"
	"log"
	"time"

	"core.yuyuid.id/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(conf config.Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s Timezone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.SSlMode,
		conf.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), config.GormConfig())
	if err != nil {
		log.Fatal("Failed to connet to the database: ", err)
		panic("Failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	return db
}
