package model

type Table struct {
	Id                string    `gorm:"primaryKey" json:"tableId"`
	Table_Description string    `gorm:"size:50" json:"tableDesc"`
	Is_Available      int       `gorm:"default:0" json:"tableIsAvailable"`
	BaseModel         BaseModel `gorm:"embedded"`
}

func (t Table) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "m_table"
}