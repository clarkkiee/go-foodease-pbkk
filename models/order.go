package models

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusInCartUnselected Status = "in-cart-unselected"
	StatusInCartSelected   Status = "in-cart-selected"
	StatusWaiting          Status = "waiting"
	StatusProcessed        Status = "processed"
	StatusReady            Status = "ready"
	StatusDone             Status = "done"
	StatusCancelled        Status = "cancelled"
	StatusRejected         Status = "rejected"
)

type Order struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Status Status `gorm:"type:enum('in-cart-unselected', 'in-cart-selected', 'waiting', 'processed', 'ready', 'done', 'cancelled', 'rejected')" json:"status"`
	CustomerID uuid.UUID `gorm:"type:uuid" json:"customer_id"`
	Customer Customer `gorm:"foreignKey:CustomerID" json:"-"`
	StoreID uuid.UUID `gorm:"type:uuid" json:"store_id"`
	Store Store `gorm:"foreignKey:StoreID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}