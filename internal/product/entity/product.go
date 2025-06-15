package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(255);not null"`
	Price       float64 `gorm:"type:integer;not null"`
}

func (Product) TableName() string {
	return "products"
}

func ToProductModel(product Product) Product {
	return Product{
		Model: product.Model,
		Name:  product.Name,
		Price: product.Price,
	}
}