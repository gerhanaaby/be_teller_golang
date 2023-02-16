package db

import (
	"fmt"
	"teller/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "P@SS1234"
	dbname   = "teller"
)

<<<<<<< Updated upstream
var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
=======
func ConnectDB()  error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		os.Getenv(`PSQL_HOST`),
		os.Getenv(`PSQL_PORT`),
		os.Getenv(`PSQL_USER`),
		os.Getenv(`PSQL_PASS`),
		os.Getenv(`PSQL_DBNAME`))
>>>>>>> Stashed changes
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
<<<<<<< Updated upstream
	db.Debug().AutoMigrate(&models.Customer{}, &models.Order{}, &models.Item{})
=======
	db.AutoMigrate(&models.User{}, &models.Skn{}, &models.Apis{})

	if err = db.Debug().Raw("SELECT * FROM apis").Scan(&models.ApiList).Error; err != nil  {
		return err
	}

	return nil
>>>>>>> Stashed changes
}
func GetDB() *gorm.DB {
	return db
}
