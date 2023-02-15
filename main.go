package main

import (
	"teller/db"
	"teller/routes"

	"github.com/joho/godotenv"
)

func init() {

	//   /home/golang/app/teller/env
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	if err := db.ConnectDB(); err != nil{
		panic(err)
	}
}


func main() {
	// Init()
	routes.Routes()	
}
