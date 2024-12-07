package repository

import (
	"context"
	"go-foodease-be/dto"
	"go-foodease-be/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	OrderRepository interface {
		AddtoCart(ctx context.Context, tx *gorm.DB) (string, error)
		GetOrderInCart(ctx context.Context, tx *gorm.DB, storeId string, customerId string) (string, error)
		CreateNewOrder(ctx context.Context, tx *gorm.DB, storeId string, customerId string) (string, error)
		GetOrderProduct(ctx context.Context, tx *gorm.DB, customerId string, orderId string, productId string) (dto.GetOrderProductResult, error)
		CreateOrderProduct(ctx context.Context, tx *gorm.DB, orderId string, productId string) (string, error)
		IncreaseOrderProductQuantity(ctx context.Context, tx *gorm.DB, customerId string, orderId string, productId string) (string, error)
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
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}

	return order.ID.String(), nil
}

func (r *orderRepository) CreateNewOrder(ctx context.Context, tx *gorm.DB, storeId string, customerId string) (string, error) {
	if tx == nil {
		tx = r.db
	}

	newOrder := models.Order{
		CustomerID: uuid.MustParse(customerId),
		StoreID: uuid.MustParse(storeId),
	}

	if err := tx.WithContext(ctx).Select("customer_id", "store_id").Create(&newOrder).Error; err != nil {
		return "", err
	}

	return newOrder.ID.String(), nil

}

func (r *orderRepository) GetOrderProduct(ctx context.Context, tx *gorm.DB, customerId string, orderId string, productId string) (dto.GetOrderProductResult, error) {
	if tx == nil {
		tx = r.db
	}


	var res dto.GetOrderProductResult

	if err := tx.WithContext(ctx).Model(&models.OrderProductDetail{}).Select("orderproductdetails.id", "orderproductdetails.quantity", "orders.customer_id").InnerJoins("orderproductdetails").Where("order_id = ? AND product_id = ? AND orders.customer_id = ?", orderId, productId, customerId).Scan(&res).Error; err != nil {
		return dto.GetOrderProductResult{}, err
	}

	return res, nil
}

func (r *orderRepository) CreateOrderProduct(ctx context.Context, tx *gorm.DB, orderId string, productId string) (string, error) {
	if tx == nil {
		tx = r.db
	}

	newOrderProduct := models.OrderProductDetail{
		OrderID: uuid.MustParse(orderId),
		ProductID: uuid.MustParse(productId),
		Quantity: 1,
	}

	if err := tx.WithContext(ctx).Model(&models.OrderProductDetail{}).Create(newOrderProduct).Error; err != nil {
		return "", err
	}

	return newOrderProduct.ID.String(), nil
}

func (r *orderRepository) IncreaseOrderProductQuantity(ctx context.Context, tx *gorm.DB, customerId string, orderId string, productId string) (string, error) {
	if tx == nil {
		tx = r.db
	}

	subQuery := tx.WithContext(ctx).Model(&models.Order{}).Select("id").Where("id = ? AND customer_id = ?", orderId, customerId)

	var updatedRow models.OrderProductDetail
	if err := tx.WithContext(ctx).Model(updatedRow).Clauses(clause.Returning{}).Where("product_id = ? AND order_id = (?)", productId, subQuery).Update("quantity", gorm.Expr("quantity + 1")).Error; err != nil {
		return "", err
	}

	return updatedRow.ID.String(), nil

}