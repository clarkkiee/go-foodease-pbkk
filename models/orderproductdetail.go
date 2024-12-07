package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderProductDetail struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
	Product Product `gorm:"foreignKey:ProductID" json:"-"`
	OrderID uuid.UUID `gorm:"type:uuid" json:"order_id"`
	Order Order `gorm:"foreignKey:OrderID" json:"-"`
	Quantity uint64 `json:"quantity"`
	Selected bool `json:"selected"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}