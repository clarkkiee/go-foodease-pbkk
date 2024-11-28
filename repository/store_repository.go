package repository

import (
	"context"
	"go-foodease-be/models"

	// "github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	StoreRepository interface {
		// RegisterCustomer(ctx context.Context, tx *gorm.DB, customer models.Customer) (models.Customer, error)
		// GetCustomerById(ctx context.Context, tx *gorm.DB, customerId string) (models.Customer, error)
		// GetCustomerByEmail(ctx context.Context, tx *gorm.DB, email string) (models.Customer, error)
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (models.Store, bool, error)
		// DeleteAccount(ctx context.Context, tx *gorm.DB, id string) error
	}

	storeRepository struct {
		db *gorm.DB
	}
)

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{
		db: db,
	}
}

func (r *storeRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (models.Store, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var store models.Store
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&store).Error; err != nil {
		return models.Store{}, false, err
	}

	return store, true, nil
}
