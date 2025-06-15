package handler

import (
	"net/http"
	"product-manager-api/internal/product/domain"
	"product-manager-api/internal/product/service"
	"product-manager-api/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var request domain.ProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewBaseErrorResponse("Invalid request format"))
		return
	}

	response, err := h.productService.CreateProduct(&request)
	if err != nil {
		if validationErrors, ok := err.(pkg.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, pkg.NewValidationErrorResponse("Validation failed", validationErrors))
			return
		}

		c.JSON(pkg.ConvertErrorCode(err), pkg.NewBaseErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, pkg.NewBaseSuccessResponse("Product created successfully", response))
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	// Parse the product ID from the URL
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewBaseErrorResponse("Invalid product ID"))
		return
	}

	// Get the product from the service
	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(pkg.ConvertErrorCode(err), pkg.NewBaseErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pkg.NewBaseSuccessResponse("Product retrieved successfully", product))
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	// Get all products from the service
	products, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(pkg.ConvertErrorCode(err), pkg.NewBaseErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pkg.NewBaseSuccessResponse("Products retrieved successfully", products))
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	// Parse the product ID from the URL
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewBaseErrorResponse("Invalid product ID"))
		return
	}

	// Parse request body
	var request domain.ProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewBaseErrorResponse("Invalid request format"))
		return
	}

	// Call service to update product
	product, err := h.productService.UpdateProduct(uint(id), &request)
	if err != nil {
		if validationErrors, ok := err.(pkg.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, pkg.NewValidationErrorResponse("Validation failed", validationErrors))
			return
		}

		c.JSON(pkg.ConvertErrorCode(err), pkg.NewBaseErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pkg.NewBaseSuccessResponse("Product updated successfully", product))
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	// Parse the product ID from the URL
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewBaseErrorResponse("Invalid product ID"))
		return
	}

	// Call service to delete product
	err = h.productService.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(pkg.ConvertErrorCode(err), pkg.NewBaseErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pkg.NewBaseSuccessResponse("Product deleted successfully", nil))
}