package routes

import (
	"teller/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
    }))

    apiRoutes := r.Group("api/item")
    {
		apiRoutes.POST("/postskn", controllers.PostSkn)
		apiRoutes.POST("/user/auth/login", controllers.UserLoginController)
	}

	r.Run(":5000")
}
