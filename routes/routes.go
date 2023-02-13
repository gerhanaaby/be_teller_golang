package routes

import (
	"golang/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	r.POST("/postcustomer", controllers.PostCustomer)
	r.POST("/postitem", controllers.PostItem)
	r.POST("/postorder", controllers.PostOrder)

	r.GET("/getorder/:orderNo", controllers.GetOrder)

	r.PUT("/updateorder/:orderNo", controllers.UpdateOrder)

	r.Run(":5000")
}