package model

import "gorm.io/datatypes"

type Bill struct {
	Id string `gorm:"primaryKey" json:"billId"`

	Trans_Date datatypes.Time

	// Customer_Id []*Discount `gorm:"many2many:customer_discount"` //many to many
	Table_Id      string
	Table         Table `gorm:"foreignKey:Table_Id"`
	Trans_Type_Id string
	Trans_Type    Trans_Type `gorm:"foreignKey:Trans_Type_Id"`
	BaseModel     BaseModel  `gorm:"embedded"`
}

func (tr Bill) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "t_bill"
}