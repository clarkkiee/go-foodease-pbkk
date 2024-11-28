package models

import (
	"go-foodease-be/helpers"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Email       string    `gorm:"uniqueIndex:idx_store_email" json:"email"`
	StoreName   string    `json:"store_name"`
	Description string    `json:"description"`
	StorePassword string  `gorm:"column:store_password" json:"store_password"`
	AddressID   *uuid.UUID `gorm:"type:uuid;default:null" json:"address_id"`
	FreeTime    string    `json:"free_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *Store) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	c.StorePassword, err = helpers.HashPassword(c.StorePassword)
	if err != nil {
		return err
	}

	return nil
}