package model

type Inventory struct {
	Id          int    `gorm:"column:id;primaryKey;"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Picture     string `gorm:"column:picture"`
	Price       int    `gorm:"column:price"`
	Count       int    `gorm:"column:count"`
}
