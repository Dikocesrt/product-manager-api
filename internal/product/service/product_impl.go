package service

import (
	"product-manager-api/internal/product/domain"
	"product-manager-api/internal/product/entity"
	"product-manager-api/internal/product/repository"
	"product-manager-api/pkg"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		productRepository: productRepository,
	}
}

func (s *ProductServiceImpl) CreateProduct(request *domain.ProductRequest) (domain.ProductResponse, error) {
	// Validate the request
	if err := pkg.Validate(request); err != nil {
		return domain.ProductResponse{}, err
	}

	// Create the product entity
	product := &entity.Product{
		Name:        request.Name,
		Price:       request.Price,
	}

	// Save the product to the repository
	newProduct, err := s.productRepository.CreateProduct(product)
	if err != nil {
		return domain.ProductResponse{}, err
	}

	return domain.ProductResponse{ID: newProduct.ID, Name: product.Name, Price: product.Price, CreatedAt: newProduct.CreatedAt, UpdatedAt: newProduct.UpdatedAt}, nil
}

func (s *ProductServiceImpl) GetProductByID(id uint) (domain.ProductResponse, error) {
	product, err := s.productRepository.GetProductByID(id)
	if err != nil {
		return domain.ProductResponse{}, err
	}
	
	return domain.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}, nil
}

func (s *ProductServiceImpl) GetAllProducts() ([]domain.ProductResponse, error) {
	products, err := s.productRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	
	var productResponses []domain.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, domain.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}
	
	return productResponses, nil
}

func (s *ProductServiceImpl) UpdateProduct(id uint, request *domain.ProductRequest) (domain.ProductResponse, error) {
	// Validate the request
	if err := pkg.Validate(request); err != nil {
		return domain.ProductResponse{}, err
	}
	
	// Create product entity for update
	product := &entity.Product{
		Name:  request.Name,
		Price: request.Price,
	}
	
	// Update via repository
	updatedProduct, err := s.productRepository.UpdateProduct(id, product)
	if err != nil {
		return domain.ProductResponse{}, err
	}
	
	// Return response
	return domain.ProductResponse{
		ID:        updatedProduct.ID,
		Name:      updatedProduct.Name,
		Price:     updatedProduct.Price,
		CreatedAt: updatedProduct.CreatedAt,
		UpdatedAt: updatedProduct.UpdatedAt,
	}, nil
}

func (s *ProductServiceImpl) DeleteProduct(id uint) error {
	return s.productRepository.DeleteProduct(id)
}