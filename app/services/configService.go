package service

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/models"
	"github.com/gin-gonic/gin"
)

type ConfigService struct {
	Config *models.MerchantToAppConfig
}

func NewConfigService(c *gin.Context) *ConfigService {
	return &ConfigService{}
}