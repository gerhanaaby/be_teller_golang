package db

import (
	"fmt"
	"os"
	"teller/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	os.Getenv(`PSQL_HOST`),
			os.Getenv(`PSQL_PORT`),
			os.Getenv(`PSQL_USER`),
			os.Getenv(`PSQL_PASS`),
			os.Getenv(`PSQL_DBNAME`))
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Skn{})
}
func GetDB() *gorm.DB {
	return db
}
