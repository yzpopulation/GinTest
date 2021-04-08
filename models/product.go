package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
func (b *Product) TableName() string {
	return "products"
}