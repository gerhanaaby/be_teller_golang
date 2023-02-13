package main

import (
	"golang/controllers"
	"golang/db"

	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectDB()
}

func main() {

	r := gin.Default()

	r.POST("/postcustomer", controllers.PostCustomer)
	r.POST("/postitem", controllers.PostItem)
	r.POST("/postorder", controllers.PostOrder)

	r.GET("/getorder/:orderNo", controllers.GetOrder)

	r.PUT("/updateorder/:orderNo", controllers.UpdateOrder)

	r.Run(":5000")
}
