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
		// AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, Authorization, access-control-allow-origin, access-control-allow-headers"},
	}))

	v1 := r.Group("v1")
	{

		UserRoutes := v1.Group("user") 
		{
			// UserRoutes.POST("/auth/login", controllers.UserLoginController)
			UserRoutes.POST("/create", controllers.CreateUser)
	
			UserAuthRoutes := UserRoutes.Group("auth")
			{
				UserAuthRoutes.POST("/login", controllers.UserLoginController)
			}
	
			UserTransRoutes := UserRoutes.Group("transac")
			{
				UserTransRoutes.POST("/postskn", controllers.SKN)                           ///user/transac/postskn
				UserTransRoutes.POST("/postinquirytransfer", controllers.PostInquiryTransfer)   // inquiry transfer
				UserTransRoutes.POST("/postinternaltransfer", controllers.PostInternalTransfer) // internal transfer
				UserTransRoutes.POST("/postgetdetail", controllers.PostGetDetail)               // get detail
				UserTransRoutes.POST("/postgetdetail2", controllers.PostGetDetail2)               // get detail
				UserTransRoutes.POST("/postadvice", controllers.PostAdvice)                     // Advice
			}
	
			UserNsbRoutes := UserRoutes.Group("nasabah") 
			{
				UserNsbRoutes.GET("/:cif/get", controllers.GetNasabahByCIF)
			}
	
			UserUtilityRoutes := UserRoutes.Group("utils")
			{
				UserUtilityRoutes.GET("/getb64/:cif", controllers.GetBase64ByCif)
			}
		}
	}

	v2 := r.Group("v2")
	{
		UserRoutes := v2.Group("user")
		{
	
			// UserRoutes.POST("/auth/login", controllers.UserLoginController)
			UserRoutes.POST("/create", controllers.CreateUser)
	
			UserAuthRoutes := UserRoutes.Group("auth")
			{
				UserAuthRoutes.POST("/login", controllers.UserLoginController)
			}
	
			UserTransRoutes := UserRoutes.Group("transac")
			{
				UserTransRoutes.POST("/postskn", controllers.SKN)                           ///user/transac/postskn
				UserTransRoutes.POST("/postinquirytransfer", controllers.PostInquiryTransfer)   // inquiry transfer
				UserTransRoutes.POST("/postinternaltransfer", controllers.PostInternalTransfer) // internal transfer
				UserTransRoutes.POST("/postgetdetail", controllers.PostGetDetail)               // get detail
				UserTransRoutes.POST("/postgetdetail2", controllers.PostGetDetail2)               // get detail
				UserTransRoutes.POST("/postadvice", controllers.PostAdvice)                     // Advice
			}
	
			UserNsbRoutes := UserRoutes.Group("nasabah") 
			{
				UserNsbRoutes.GET("/:cif/get", controllers.GetNasabahByCIF)
			}
	
			UserUtilityRoutes := UserRoutes.Group("utils")
			{
				UserUtilityRoutes.GET("/getb64/:cif", controllers.GetBase64ByCif)
			}
		}
	
	}


	r.Run(":5000")
}
