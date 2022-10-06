package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Slug  string  `json:"slug" gorm:"type:varchar(250);index:idx_slug"`
	Name  string  `json:"name" gorm:"type:varchar(250)"`
	Sku   string  `json:"sku" gorm:"type:varchar(100);index:idx_sku"`
	Price float32 `json:"price" gorm:"type:float;index:idx_price"`
}
