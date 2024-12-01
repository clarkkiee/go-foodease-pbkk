package repository

import (
	"context"
	"go-foodease-be/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	StoreRepository interface {
		RegisterStore(ctx context.Context, tx *gorm.DB, store models.Store) (models.Store, error)
		GetStoreById(ctx context.Context, tx *gorm.DB, customerId string) (models.Store, error)
		// GetCustomerByEmail(ctx context.Context, tx *gorm.DB, email string) (models.Customer, error)
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (models.Store, bool, error)
		DeleteAccount(ctx context.Context, tx *gorm.DB, id string) error
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

func (r *storeRepository) RegisterStore(ctx context.Context, tx *gorm.DB, store models.Store) (models.Store, error) {
	if tx == nil {
		tx = r.db
	}
	
	if err := tx.WithContext(ctx).Create(&store).Error; err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (r *storeRepository) GetStoreById(ctx context.Context, tx *gorm.DB, storeId string) (models.Store, error) {
	if tx == nil {
		tx = r.db
	}

	var store models.Store
	
	if err := tx.WithContext(ctx).Where("id = ?", storeId).Take(&store).Error; err != nil {
		return models.Store{}, err
	}

	return store, nil

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

func (r *storeRepository) DeleteAccount(ctx context.Context, tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&models.Store{}, uuid.MustParse(id)).Error; err != nil {
		return err
	}

	return nil
}