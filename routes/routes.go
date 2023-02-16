package routes

import (
	"teller/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
        // AllowOrigins: []string{"* or Write Alowed URL"},
        AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
    }))

    UserRoutes := r.Group("user")
    {

		// UserAuthRoutes := ""
		UserRoutes.POST("/auth/login", controllers.UserLoginController)

		UserTransRoutes := UserRoutes.Group("transac")
		{
		UserTransRoutes.POST("/postskn", controllers.PostSkn) ///user/transac/postskn

		}
	}

	r.Run(":5000")
}
