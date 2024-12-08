package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateProductRequest struct {
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description"`
	PriceBefore   float64    `json:"price_before"`
	PriceAfter    float64    `json:"price_after"`
	ProductionTime string `json:"production_time"`
	ExpiredTime   string `json:"expired_time"`
	Stock         uint64    `json:"stock"`
	CategorySlug    string    `json:"category_slug"`
	ImageID       string   `json:"image_id,omitempty"`
}

type CreateProduct struct {
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description"`
	PriceBefore   float64    `json:"price_before"`
	PriceAfter    float64    `json:"price_after"`
	ProductionTime string `json:"production_time"`
	ExpiredTime   string `json:"expired_time"`
	Stock         uint64    `json:"stock"`
	CategoryID    uuid.UUID    `json:"category_id"`
	ImageID       string   `json:"image_id,omitempty"`
}

type ProductResponse struct {
	ID            string    `json:"id"`
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description"`
	PriceBefore   float64    `json:"price_before"`
	PriceAfter    float64    `json:"price_after"`
	ProductionTime string `json:"production_time"`
	ExpiredTime   string `json:"expired_time"`
	Stock         uint64    `json:"stock"`
	CategoryID    string    `json:"category_id"`
	ImageID       string   `json:"image_id,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UpdateProductRequest struct {
    ProductName   string    `json:"product_name"`
    Description   string    `json:"description"`
    PriceBefore   float64    `json:"price_before"`
    PriceAfter    float64    `json:"price_after"`
	ProductionTime string	`json:"production_time"`
	ExpiredTime 	string 	`json:"expired_time"`
	CategorySlug 	string `json:"category_slug"`
	ImageID       string   `json:"image_id,omitempty"`
}

type GetMinimumProductResult struct {
	ID string `json:"id"`
	Stock uint64 `json:"stock"`
	StoreID string `json:"store_id"`
	CategoryID string `json:"category_id"`
}

