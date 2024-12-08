package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Slug string `json:"slug"`
	CategoryName string `json:"category_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}