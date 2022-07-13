package model

type Menu_Price struct {
	Id      string `gorm:"primaryKey" json:"id"`
	Menu_Id int
	Menu    Menu `gorm:"foreignKey:Menu_Id" json:"menu"`

	Price     int       `json:"menuPrice"`
	BaseModel BaseModel `gorm:"embedded"`
}

func (mp Menu_Price) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "m_menu_price"
}