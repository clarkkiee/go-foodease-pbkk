package repository

import (
	"context"
	"go-foodease-be/models"

	"gorm.io/gorm"
)

type (
	CustomerRepository interface {
		// RegisterCustomer(ctx context.Context, tx *gorm.DB, customer models.Customer) (models.Customer, error)
		GetCustomerById(ctx context.Context, tx *gorm.DB, customerId string) (models.Customer, error)
		GetCustomerByEmail(ctx context.Context, tx *gorm.DB, email string) (models.Customer, error)
	}

	customerRepository struct {
		db *gorm.DB
	}
)

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}


// func (r *customerRepository) RegisterCustomer(ctx *context.Context, tx *gorm.DB, customer models.Customer) (models.Customer, error){}

func (r *customerRepository) GetCustomerById(ctx context.Context, tx *gorm.DB, customerId string) (models.Customer, error) {
	if tx == nil {
		tx = r.db
	}

	var customer models.Customer
	
	if err := tx.WithContext(ctx).Where("id = ?", customerId).Take(&customer).Error; err != nil {
		return models.Customer{}, err
	}

	return customer, nil

}

func (r *customerRepository) GetCustomerByEmail(ctx context.Context, tx *gorm.DB, email string) (models.Customer, error) {
	if tx == nil {
		tx = r.db
	}

	var customer models.Customer

	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&customer).Error; err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}