package dto

import "time"

type CreateProductRequest struct {
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description"`
	PriceBefore   float64    `json:"price_before"`
	PriceAfter    float64    `json:"price_after"`
	ProductionTime time.Time `json:"production_time"`
	ExpiredTime   time.Time `json:"expired_time"`
	Stock         uint64    `json:"stock"`
	CategoryID    string    `json:"category_id"`
	ImageID       *string   `json:"image_id"`
}

type ProductResponse struct {
	ID            string    `json:"id"`
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description"`
	PriceBefore   float64    `json:"price_before"`
	PriceAfter    float64    `json:"price_after"`
	ProductionTime time.Time `json:"production_time"`
	ExpiredTime   time.Time `json:"expired_time"`
	Stock         uint64    `json:"stock"`
	CategoryID    string    `json:"category_id"`
	ImageID       *string   `json:"image_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type UpdateProductRequest struct {
    ProductName   string    `json:"product_name"`
    Description   string    `json:"description"`
    PriceBefore   float64    `json:"price_before"`
    PriceAfter    float64    `json:"price_after"`
    Stock         uint64    `json:"stock"`
    CategoryID    string    `json:"category_id"`
    ImageID       *string   `json:"image_id"`
}
