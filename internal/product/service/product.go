package service

import "product-manager-api/internal/product/domain"

type ProductService interface {
	CreateProduct(product *domain.ProductRequest) (domain.ProductResponse, error)
	GetProductByID(id uint) (domain.ProductResponse, error)
	GetAllProducts() ([]domain.ProductResponse, error)
	UpdateProduct(id uint, product *domain.ProductRequest) (domain.ProductResponse, error)
	DeleteProduct(id uint) error
}