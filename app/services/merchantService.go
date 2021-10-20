package service

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/models"
	"github.com/gin-gonic/gin"
)

type MerchantService struct {
	Merchant *models.Merchant
}

func NewMerchantService(c *gin.Context) *MerchantService {
	return &MerchantService{}
}