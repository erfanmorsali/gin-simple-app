package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectToDb() (*gorm.DB, error) {
	dsn := "host=localhost user=erfan password=123456 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return db, nil
}
