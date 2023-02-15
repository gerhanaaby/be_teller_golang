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

func ConnectDB()  error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
	os.Getenv(`PSQL_HOST`),
			5432,
			os.Getenv(`PSQL_USER`),
			os.Getenv(`PSQL_PASS`),
			os.Getenv(`PSQL_DBNAME`))
	fmt.Println(`PSQL---------------->`+psqlInfo)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.User{}, &models.Skn{})
	return nil
}
func GetDB() *gorm.DB {
	return db
}
