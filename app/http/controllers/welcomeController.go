package controllers

import (
	. "github.com/ArtisanCloud/PowerKeyVault/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WelcomeController struct {
	ServiceWelcome *WelcomeService
}

func NewWelcomeController() (ctl *WelcomeController) {

	return &WelcomeController{
		ServiceWelcome: NewWelcomeService(),
	}
}

func WebGetHome(context *gin.Context) {
	ctl := NewWelcomeController()

	r := ctl.ServiceWelcome.GetWelcome()

	context.JSON(http.StatusOK, r)
}
