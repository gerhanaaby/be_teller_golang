package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, Authorization, access-control-allow-origin, access-control-allow-headers"},
	}))

	
	golang := r.Group("golang") 
	{
		RunV1(golang)
		RunV2(golang)
	}
	r.Run(":5000")
}
