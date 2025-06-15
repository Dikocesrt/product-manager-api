package service

import "product-manager-api/internal/product/domain"

type ProductService interface {
	CreateProduct(product *domain.CreateProductRequest) (domain.ProductResponse, error)
	GetProductByID(id uint) (domain.ProductResponse, error)
	GetAllProducts() ([]domain.ProductResponse, error)
}