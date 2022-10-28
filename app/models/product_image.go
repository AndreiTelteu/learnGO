package models

type ProductImage struct {
	ID        uint    `gorm:"primarykey"`
	File      string  `json:"file" gorm:"type:varchar(250)"`
	Order     int32   `json:"order" gorm:"index:idx_order"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignkey:id"`
}
