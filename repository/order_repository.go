package repository

import (
	"context"
	"go-foodease-be/models"

	"gorm.io/gorm"
)

type (
	OrderRepository interface {
		AddtoCart(ctx context.Context, tx *gorm.DB) (string, error)
		GetOrderInCart(ctx context.Context, tx *gorm.DB, storeId string, customerId string) (string, error)
	}

	orderRepository struct {
		db *gorm.DB
	}
)

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) AddtoCart(ctx context.Context, tx *gorm.DB) (string, error){
	return "", nil
}

func (r *orderRepository) GetOrderInCart(ctx context.Context, tx *gorm.DB, storeId string, customerId string) (string, error) {
	if tx == nil {
		tx = r.db
	}
	
	var order models.Order
	if err := tx.WithContext(ctx).Select("id").Where("customer_id = ? AND store_id = ? AND status IN ('in-cart-selected', 'in-cart-unselected')", customerId, storeId).First(&order).Error; err != nil {
		return "", err
	}

	return order.ID.String(), nil
}