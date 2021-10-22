package models

const PLATFORM_WECHAT = "WeChat"
const PLATFORM_ALIPAY = "AliPay"

type Platform struct {
	MyModel

	Name string `gorm:"column:name" json:"name"`
}
