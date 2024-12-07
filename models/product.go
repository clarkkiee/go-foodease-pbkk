package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	PriceBefore float64 `json:"price_before"`
	PriceAfter float64 `json:"price_after"`
	ProductionTime time.Time `gorm:"type:time" json:"production_time"`
	ExpiredTime time.Time `gorm:"type:time" json:"expired_time"`
	Stock uint64 `json:"stock"`
	StoreId uuid.UUID `gorm:"type:uuid" json:"store_id"`
	CategoryID uuid.UUID `gorm:"type:uuid" json:"category_id"`
	ImageID *uuid.UUID `gorm:"type:uuid;" json:"image_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

