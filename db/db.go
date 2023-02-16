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
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	os.Getenv(`PSQL_HOST`),
	os.Getenv(`PSQL_PORT`),
			os.Getenv(`PSQL_USER`),
			os.Getenv(`PSQL_PASS`),
			os.Getenv(`PSQL_DBNAME`))
	fmt.Println(`PSQL---------------->`+psqlInfo)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.User{}, &models.Skn{}, &models.Apis{})

	if err = db.Debug().Raw("SELECT * FROM apis").Scan(&models.ApiList).Error; err != nil  {
		return err
	}
	return nil
}
func GetDB() *gorm.DB {
	return db
}
