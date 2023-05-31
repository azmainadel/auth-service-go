package db

import (
	"authentication/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresDb() *gorm.DB {
	DSN_STRING := "test"
	db, err := gorm.Open(postgres.Open(DSN_STRING), &gorm.Config{})

	if err != nil {
		panic("DB connection failed!")
	}

	db.AutoMigrate(&models.User{})
	return db
}
