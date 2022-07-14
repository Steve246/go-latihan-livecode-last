package model

type Menu struct {
	Id        string    `gorm:"not null" json:"id"`
	Menu_Name string    `gorm:"size:50; not null" json:"menu_name"`
	BaseModel BaseModel `gorm:"embedded"`
}

func (m Menu) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "mst_menu"
}
