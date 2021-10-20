package controllers

import (
	. "github.com/ArtisanCloud/PowerKeyVault/app/services"
	"github.com/gin-gonic/gin"
)

type ConfigAPIController struct {
	*APIController
	ServiceConfig *ConfigService
}


func NewConfigAPIController() (ctl *ConfigAPIController) {
	return &ConfigAPIController{
		ServiceConfig: NewConfigService(nil),
	}
}

func APIConfigGetList(context *gin.Context) {

}


func APIConfigBind(context *gin.Context) {

}

func APIConfigUpdate(context *gin.Context) {

}

func APIConfigUnBind(context *gin.Context) {

}