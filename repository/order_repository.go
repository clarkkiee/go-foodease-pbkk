package repository

import (
	"context"
	"errors"
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
		GetOrderById(ctx context.Context, tx *gorm.DB, orderId string) (dto.OrderDetails, error) 
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
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

	if err := tx.WithContext(ctx).Model(&models.OrderProductDetail{}).Select("order_product_details.id", "order_product_details.quantity", "orders.customer_id").Joins("INNER JOIN orders ON orders.id = order_product_details.order_id").Where("order_id = ? AND product_id = ? AND orders.customer_id = ?", orderId, productId, customerId).Scan(&res).Error; err != nil {
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
		Selected: true,
	}

	if err := tx.WithContext(ctx).Model(&models.OrderProductDetail{}).Create(&newOrderProduct).Error; err != nil {
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
	if err := tx.WithContext(ctx).Model(&updatedRow).Clauses(clause.Returning{}).Where("product_id = ? AND order_id = (?)", productId, subQuery).Update("quantity", gorm.Expr("quantity + 1")).Error; err != nil {
		return "", err
	}

	return updatedRow.ID.String(), nil

}

func (r *orderRepository) GetOrderById(ctx context.Context, tx *gorm.DB, orderId string) (dto.OrderDetails, error) {
	if tx == nil {
		tx = r.db
	}
	
	var order dto.OrderDetails

	if err := tx.WithContext(ctx).Raw(`
	SELECT 
		o.id,
		o.status,
		s.id AS "store_id",
		s.store_name,
		c.id AS "customer_id",
		c.email AS "customer_email",
		CONCAT(c.first_name, c.last_name) AS "customer_name",
		a.street AS "address_street",
		ST_AsText(coordinates) AS coordinates,
		p.id AS "product_id",
		od.id AS "order_product_id",
		od.selected AS "order_product_selected",
		od.quantity AS "order_product_quantity",
		p.product_name,
		p.price_before,
		p.price_after,
		p.stock,
		o.created_at,
		o.updated_at
	FROM orders o
	INNER JOIN order_product_details od ON o.id = od.order_id
	INNER JOIN stores s ON s.id = o.store_id
	INNER JOIN products p ON p.id = od.product_id
	INNER JOIN customers c ON c.id = o.customer_id
	LEFT JOIN addresses a ON c.active_address_id = a.id
	WHERE o.id = ?
	`, orderId).Scan(&order).Error; err != nil {
		return dto.OrderDetails{}, err	
	}

	return order, nil
}