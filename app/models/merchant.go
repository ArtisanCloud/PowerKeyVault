package models

type Merchant struct {
	Model

	Name     string `gorm:"column:name" json:"name"`
	MchID    string `gorm:"column:mch_id" json:"mchID"`
	Secret   string `gorm:"column:secret" json:"secret"`
	SecretV3 string `gorm:"column:secret_v3" json:"secretV3"`
	CertPath string `gorm:"column:cert_path" json:"certPath"`
	KeyPath  string `gorm:"column:key_path" json:"keyPath"`
	SerialNo string `gorm:"column:serial_no" json:"serialNO"`
}
