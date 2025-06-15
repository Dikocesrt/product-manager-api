package repository

import (
	"context"
	"encoding/json"
	"product-manager-api/config"
	"product-manager-api/internal/product/entity"
	"product-manager-api/pkg"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl {
	db.AutoMigrate(&entity.Product{})
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (r *ProductRepositoryImpl) CreateProduct(product *entity.Product) (entity.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return entity.Product{}, err
	}
	ctx := context.Background()
	
	// Cache individual product
	cacheKey := "product-manager:product:id:" + strconv.FormatUint(uint64(product.ID), 10)
	if productJSON, err := json.Marshal(product); err == nil {
		config.Redis.Set(ctx, cacheKey, productJSON, 15*time.Minute)
	}
	
	// Invalidate the all products cache
	r.invalidateProductsCache()

	return *product, nil
}

func (r *ProductRepositoryImpl) GetProductByID(id uint) (entity.Product, error) {
	// First, try to get from Redis
	ctx := context.Background()
	cacheKey := "product-manager:product:id:" + strconv.FormatUint(uint64(id), 10)
	
	// Check if product exists in Redis cache
	productJSON, err := config.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		// Product found in cache
		var product entity.Product
		if err := json.Unmarshal([]byte(productJSON), &product); err == nil {
			return product, nil
		}
	}
	
	// If not in cache or error in unmarshaling, get from database
	var product entity.Product
	if err := r.db.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Product{}, pkg.ErrProductNotFound
		}
		return entity.Product{}, err
	}
	
	// Store in Redis cache for future requests (cache for 15 minutes)
	if productJSON, err := json.Marshal(product); err == nil {
		config.Redis.Set(ctx, cacheKey, productJSON, 15*time.Minute)
	}
	
	return product, nil
}

func (r *ProductRepositoryImpl) GetAllProducts() ([]entity.Product, error) {
	// First, try to get from Redis
	ctx := context.Background()
	cacheKey := "product-manager:products:all"
	
	// Check if products list exists in Redis cache
	productsJSON, err := config.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		// Products found in cache
		var products []entity.Product
		if err := json.Unmarshal([]byte(productsJSON), &products); err == nil {
			return products, nil
		}
	}
	
	// If not in cache or error in unmarshaling, get from database
	var products []entity.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	
	// Store in Redis cache for future requests (cache for 5 minutes)
	// We use shorter TTL for list of all products as this could change frequently
	if productsJSON, err := json.Marshal(products); err == nil {
		config.Redis.Set(ctx, cacheKey, productsJSON, 5*time.Minute)
	}
	
	return products, nil
}

// Helper function to invalidate the all products cache
func (r *ProductRepositoryImpl) invalidateProductsCache() {
	ctx := context.Background()
	config.Redis.Del(ctx, "product-manager:products:all")
}