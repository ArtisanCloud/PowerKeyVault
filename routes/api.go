package routes

import (
	. "github.com/ArtisanCloud/PowerKeyVault/app/http/controllers"
	. "github.com/ArtisanCloud/PowerKeyVault/app/http/middleware"
	"github.com/ArtisanCloud/PowerKeyVault/app/http/request"
	"github.com/gin-gonic/gin"
)


func InitializeAPIRoutes(router *gin.Engine) {

	apiRouter := router.Group("/api")
	{
		apiRouter.Use(Maintenance, AuthAPI, AuthWeb)
		{
			// Handle the index route
			apiRouter.GET("/", APIGetHome)

			apiRouter.GET("/merchant/get", request.ValidateRequestLogin, APIMerchantGetList)
			apiRouter.POST("/merchant/get", request.ValidateRequestLogin, APIMerchantCreate)
			apiRouter.PUT("/merchant/get", request.ValidateRequestLogin, APIMerchantUpdate)
			apiRouter.DELETE("/merchant/get", request.ValidateRequestLogin, APIMerchantDelete)

			apiRouter.GET("/app/get", request.ValidateRequestLogin, APIAPPGetList)
			apiRouter.POST("/app/get", request.ValidateRequestLogin, APIAPPCreate)
			apiRouter.PUT("/app/get", request.ValidateRequestLogin, APIAPPUpdate)
			apiRouter.DELETE("/app/get", request.ValidateRequestLogin, APIAPPDelete)

			apiRouter.GET("/config/list", request.ValidateRequestLogin, APIConfigGetList)
			apiRouter.POST("/config/bind", request.ValidateRequestLogin, APIConfigBind)
			apiRouter.PUT("/config/update", request.ValidateRequestLogin, APIConfigUpdate)
			apiRouter.DELETE("/config/unbind", request.ValidateRequestLogin, APIConfigUnBind)

			//apiRouter.POST("/make", ValidateRequestMakeWelcome, ctlWelcome.APIMakeWelcome)
			// apiRouter.PUT("/somePut", putting)
			// apiRouter.DELETE("/someDelete", deleting)
			// apiRouter.PATCH("/somePatch", patching)
			// apiRouter.HEAD("/someHead", head)
			// apiRouter.OPTIONS("/someOptions", options)

		}
	}
}
