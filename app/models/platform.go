package models


const PLATFORM_WECHAT = "WeChat"
const PLATFORM_ALIPAY = "AliPay"

type Platform struct {
	Model

	Name string `gorm:"column:name" json:"name"`
}
