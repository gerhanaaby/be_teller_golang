package routes

import (
	"teller/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	r.POST("/postcustomer", controllers.PostCustomer)
	r.POST("/postitem", controllers.PostItem)
	r.POST("/postorder", controllers.PostOrder)
	r.POST("/postskn", controllers.PostSkn)
	r.POST("/auth/user/login", controllers.UserLoginController)
	r.POST("/teller/hostinq", controllers.HostInquiry)

	r.GET("/getorder/:orderNo", controllers.GetOrder)

	r.PUT("/updateorder/:orderNo", controllers.UpdateOrder)

	// r.POST("/user/auth/register")

	// r.POST("/user/auth/verify", controllers.UserLoginVerify)

	r.Run(":5000")
}
