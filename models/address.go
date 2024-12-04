package models

import (
	"go-foodease-be/types"
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
    Street string `json:"street"`
    Coordinates types.Coordinates `gorm:"type:geometry(Point,4326)" json:"coordinates"`
    CustomerID  uuid.UUID  `gorm:"type:uuid;" json:"customer_id,omitempty"`
	Customer    Customer   `gorm:"foreignKey:CustomerID" json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}