package model

type Menu_Price struct {
	Id string `gorm:"primaryKey"`

	MenuId string

	Menus Menu `gorm:"foreignKey:MenuId" json:"menu"`

	Price     int       `json:"menuPrice"`
	BaseModel BaseModel `gorm:"embedded"`
}

// MenuId []Menu `json:"menu"`

func (mp Menu_Price) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "m_menu_price"
}