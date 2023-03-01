package db

import (
	"fmt"
	"teller/inits"
	"teller/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() error {

	models.ApiMap = make(map[string]models.Apis)

	if db, err = gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
			inits.Cfg.DBHost,
			inits.Cfg.DBPort,
			inits.Cfg.DBUserName,
			inits.Cfg.DBUserPassword,
			inits.Cfg.DBName)), &gorm.Config{}); err != nil {
		return err
	}

	db.AutoMigrate(
		&models.User{}, 
		&models.Skn{}, 
		&models.Nasabah{}, 
		&models.Apis{}, 
		&models.InquiryTransfer{}, 
		&models.InternalTransfer{},
		&models.GetDetail{}, 
		&models.Advice{})
    
	if err = db.Table(`apis`).Scan(&models.ApiList).Error; err != nil  {
		return err
	}
	fmt.Println(models.ApiList[0].Name)
	fmt.Println(models.ApiList[1].Name)

	for _, api := range models.ApiList{
		models.ApiMap[api.Name] = api
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
