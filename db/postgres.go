package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresDb() *gorm.DB {
	DSN_STRING := "test"
	db, err := gorm.Open(postgres.Open(DSN_STRING), &gorm.Config{})

	if err != nil {
		panic("DB connection failed!")
	}

	return db
}
