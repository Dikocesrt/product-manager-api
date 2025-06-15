package repository

import "product-manager-api/internal/product/entity"

type ProductRepository interface {
	CreateProduct(product *entity.Product) (entity.Product, error)
	GetProductByID(id uint) (entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
}