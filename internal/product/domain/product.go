package domain

import "time"

type ProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Price int `json:"price" validate:"required,gt=0"`
}

type ProductResponse struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price int `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}