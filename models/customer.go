package models

import (
	"go-foodease-be/helpers"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Email          string     `json:"email"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Password       string     `gorm:"column:customer_password" json:"password"`
	ActiveAddressId *uuid.UUID `gorm:"type:uuid;default:null" json:"active_address_id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	c.Password, err = helpers.HashPassword(c.Password)
	if err != nil {
		return err
	}

	return nil
}