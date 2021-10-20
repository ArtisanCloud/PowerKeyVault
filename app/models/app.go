package models

type App struct {
	Model

	Name   string `gorm:"column:name" json:"name"`
	AppID  string `gorm:"column:app_id" json:"appID"`
	Secret string `gorm:"column:secret" json:"secret"`
}
