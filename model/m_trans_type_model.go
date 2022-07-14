package model

type Trans_Type struct {
	Id                string    `gorm:"primaryKey" json:"transTypeId"`
	Trans_Description string    `gorm:"size:100" json:"transDesc"`
	BaseModel         BaseModel `gorm:"embedded"`
}

func (tt Trans_Type) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "m_trans_type"
}