package models

import "github.com/ArtisanCloud/PowerLibs/object"

// merchant_to_app_config 数据表结构
type MerchantToAppConfig struct {
	*MyRelationship

	Merchant *Merchant `gorm:"foreignKey:MerchantUUID;references:UUID"`
	App      *App      `gorm:"foreignKey:AppUUID;references:UUID"`

	//common fields
	Status       int8              `gorm:"column:status" json:"status"`
	ConfigType   int8              `gorm:"column:config_type" json:"ConfigType"`
	MerchantUUID object.NullString `gorm:"column:merchant_uuid;not null;index:index_merchant_uuid" json:"merchantUUID"`
	AppUUID      object.NullString `gorm:"column:app_uuid;not null;index:index_app_uuid" json:"appUUID"`
}

const TABLE_NAME_Merchant_To_App_Config = "merchant_to_app_config"
