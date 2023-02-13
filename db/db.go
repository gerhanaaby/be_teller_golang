package db

import (
	"fmt"
	"golang/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "katasandi12"
	dbname   = "golangdev"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Debug().AutoMigrate(&models.Customer{}, &models.Order{}, &models.Item{})
}
func GetDB() *gorm.DB {
	return db
}
