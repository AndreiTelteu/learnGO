package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Slug        string `json:"slug" gorm:"type:varchar(250);index:idx_slug"`
	Name        string `json:"name" gorm:"type:varchar(250)"`
	Description string `json:"description" gorm:"type:varchar(250)"`
	Image       string `json:"image" gorm:"type:varchar(500)"`

	Products []*Product `gorm:"many2many:product_categories;"`
}
