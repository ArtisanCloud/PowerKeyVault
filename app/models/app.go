package models

import "github.com/ArtisanCloud/PowerLibs/object"

type App struct {
	*MyModel

	Name   string `gorm:"column:name" json:"name"`
	AppID  string `gorm:"column:app_id" json:"appID"`
	Secret string `gorm:"column:secret" json:"secret"`
}

func NewApp(mapObject *object.Collection) *App {

	return &App{
		MyModel: NewMyModel(),
	}
}
