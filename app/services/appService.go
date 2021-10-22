package service

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/models"
	"github.com/ArtisanCloud/PowerKeyVault/database"
	"github.com/ArtisanCloud/PowerLibs/fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type AppService struct {
	App *models.App
}

func NewAppService(c *gin.Context) *AppService {
	return &AppService{}
}

func (service *AppService) Create() (*models.App, error) {
	db := database.DBConnection
	fmt.Dump(service.App)
	dbRes := db.Create(service.App)
	if dbRes.RowsAffected > 0 {
		return service.App, nil
	}
	return nil, dbRes.Error
}

func (service *AppService) Index(page, pageSize int) ([]models.App, error) {

	var apps []models.App
	db := database.DBConnection
	dbRes := db.Model(&models.User{}).Offset(page).Limit(pageSize).Find(&apps)
	if dbRes.Error != nil {
		return apps, nil
	}
	return nil, dbRes.Error
}

func (service *AppService) Delete(id int) (bool, error) {
	db := database.DBConnection
	dbRes := db.Delete(&models.App{}, id)
	if dbRes.Error != nil {
		return false, dbRes.Error
	}
	return true, nil
}

func (service *AppService) Update(id int) (bool, error) {
	db := database.DBConnection
	dbRes := db.Model(&models.App{}).Where("id =?", id).Updates(map[string]interface{}{
		"name":       service.App.Name,
		"app_id":     service.App.AppID,
		"secret":     service.App.Secret,
		"updated_at": time.Now(),
	})

	if dbRes.RowsAffected > 0 {
		return true, nil
	}

	return false, dbRes.Error
}
