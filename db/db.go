package db

import (
	"fmt"
	"os"
	"teller/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error	
)

var PSQLDB *gorm.DB

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	os.Getenv(`PSQL_HOST`),
			os.Getenv(`PSQL_PORT`),
			os.Getenv(`PSQL_USER`),
			os.Getenv(`PSQL_PASS`),
			os.Getenv(`PSQL_DBNAME`))
	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	PSQLDB = DB
	DB.AutoMigrate(&models.User{}, &models.Skn{})
}
func GetDB() *gorm.DB {
	return DB
}
