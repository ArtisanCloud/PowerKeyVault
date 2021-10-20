package controllers

import (
	. "github.com/ArtisanCloud/PowerKeyVault/app/services"
	"github.com/gin-gonic/gin"
)

type MerchantAPIController struct {
	ServiceMerchant *MerchantService
}


func NewMerchantAPIController() (ctl *MerchantAPIController) {

	return &MerchantAPIController{
		ServiceMerchant: NewMerchantService(nil),
	}
}


func APIMerchantGetList(context *gin.Context) {

}


func APIMerchantCreate(context *gin.Context) {

}

func APIMerchantUpdate(context *gin.Context) {

}

func APIMerchantDelete(context *gin.Context) {

}

