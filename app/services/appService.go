package service

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/models"
	"github.com/ArtisanCloud/PowerKeyVault/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type AppService struct {
	App *models.App
}

func NewAppService(c *gin.Context) *AppService {
	return &AppService{}
}

func (service *AppService) Upsert(uniqueName string, apps []*models.App) error {

	if len(apps) <= 0 {
		return nil
	}
	result := database.DBConnection.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: uniqueName}},
		DoUpdates: clause.AssignmentColumns(models.GetModelFields(models.App{})),
	}).Create(&apps)

	return result.Error
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
