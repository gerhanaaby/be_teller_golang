package main

import (
	"log"
	"teller/db"
	"teller/models"
	"teller/routes"

	"github.com/wpcodevo/golang-gorm-postgres/initializers"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func init() {
	db.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	
	routes.Routes()	
}
