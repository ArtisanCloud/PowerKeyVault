package service

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/models"
	"github.com/gin-gonic/gin"
)

type AppService struct {
	App *models.App
}

func NewAppService(c *gin.Context) *AppService {
	return &AppService{}
}

func (service *AppService)Create() *models.App  {

	return nil
}