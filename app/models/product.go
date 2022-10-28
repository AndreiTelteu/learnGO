package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Slug         string  `json:"slug" gorm:"type:varchar(250);index:idx_slug"`
	Name         string  `json:"name" gorm:"type:varchar(250)"`
	Sku          string  `json:"sku" gorm:"type:varchar(100);index:idx_sku"`
	Price        float32 `json:"price" gorm:"type:float;index:idx_price"`
	RegularPrice float32 `json:"regular_price" gorm:"type:float;index:idx_regular_price"`
	SalePrice    float32 `json:"sale_price" gorm:"type:float;index:idx_sale_price"`
	Description  string  `json:"description" gorm:"type:text"`

	Images     []*ProductImage
	Categories []*Category `gorm:"many2many:product_categories;"`
}
