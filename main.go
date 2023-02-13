package main

import (
	"golang/db"
	"golang/routes"
)

func init() {
	db.ConnectDB()
}

func main() {

	routes.Routes()

}
