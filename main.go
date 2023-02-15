package main

import (
	"teller/db"
	"teller/routes"

	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(`/home/golang/app/teller/env`); err != nil {
		return
	}

	db.ConnectDB()
}


func main() {
	
	routes.Routes()	
}
