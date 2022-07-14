package model

type Bill_Detail struct {
	Id            string `gorm:"primaryKey" json:"billDetailId"`
	Bill_Id       string
	Bill          Bill `gorm:"foreignKey:Bill_Id" json:"billId"`
	Menu_Price_Id string
	Menu_Price    Menu_Price `gorm:"foreignKey:	Menu_Price_Id" json:"menuPriceId"`
	Qty           int        `json:"billQty"`
	BaseModel     BaseModel  `gorm:"embedded"`
}

func (tbd Bill_Detail) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "t_bill_detail"
}