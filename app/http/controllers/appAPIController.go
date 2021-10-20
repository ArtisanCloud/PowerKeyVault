package controllers

import (
	. "github.com/ArtisanCloud/PowerKeyVault/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppAPIController struct {
	*APIController
	ServiceApp *AppService
}

func NewAppAPIController(c *gin.Context) (ctl *AppAPIController) {
	return &AppAPIController{
		ServiceApp: NewAppService(c),
	}
}

func APIAPPGetList(context *gin.Context) {

}

func APIAPPCreate(context *gin.Context) {
	ctl := NewAppAPIController(context)

	r := ctl.ServiceApp.Create()

	context.JSON(http.StatusOK, r)

}

func APIAPPUpdate(context *gin.Context) {

}

func APIAPPDelete(context *gin.Context) {

}
