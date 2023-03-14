package routes

import (
	"teller/controllers"

	"github.com/gin-gonic/gin"
)

func RunV2(r *gin.RouterGroup) {
	v1 := r.Group("v2")
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
				UserTransRoutes.POST("/postskn", controllers.SKN)                               ///user/transac/postskn
				UserTransRoutes.POST("/postinquirytransfer", controllers.PostInquiryTransfer)   // inquiry transfer
				UserTransRoutes.POST("/postinternaltransfer", controllers.PostInternalTransfer) // internal transfer
				UserTransRoutes.POST("/postgetdetail", controllers.PostGetDetail)               // get detail
				UserTransRoutes.POST("/postgetdetail2", controllers.PostGetDetail2)             // get detail
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
}