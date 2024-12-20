package service

import (
	"context"
	"errors"
	"fmt"
	"go-foodease-be/dto"
	"go-foodease-be/repository"

	"gorm.io/gorm"
)

type (
	OrderService interface {
		AddToCart(ctx context.Context, customerId string, productId string) (dto.Order, error)
		GetCustomerCart(ctx context.Context, customerId string) (dto.GetUserCartResults, error)
	}

	orderService struct {
		orderRepo repository.OrderRepository
		productRepo repository.ProductRepository
		db *gorm.DB
	}
)

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository, db *gorm.DB) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		productRepo: productRepo,
		db: db,
	}
}

func (s *orderService) AddToCart(ctx context.Context, customerId string, productId string) (dto.Order, error) {
	
	//cek produk ada atau tidak, cek juga stock nya
	product, err := s.productRepo.GetMinimumProduct(ctx, nil, productId)
	if err != nil {
		return dto.Order{}, err
	}

	if product.Stock <= 0 {
		return dto.Order{}, errors.New("product is sold out")
	}

	tx := s.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	//cek apakah sudah ada order di cart (by customer id dan store id)
	order, err := s.orderRepo.GetOrderInCart(ctx, tx, product.StoreID, customerId)
	if err != nil {
		tx.Rollback()
		return dto.Order{}, err
	}

	//jika tidak ada record, buat record order baru 
	//jika sudah ada, ambil order id nya saja
	orderID := ""
	if order == "" || len(order) == 0 {
		//buat record order 
		newOrderId, _ := s.orderRepo.CreateNewOrder(ctx, tx, product.StoreID, customerId)
		orderID = newOrderId
	} else {
		//simpan order id nya
		orderID = order
	}

	//cek data di order apakah sudah ada produk yang dimaksud di order cart
	orderProduct, err := s.orderRepo.GetOrderProduct(ctx, tx, customerId, orderID, product.ID)
	if err != nil {
		tx.Rollback()
		return dto.Order{}, err
	}

	var orderProductID string
	if orderProduct == (dto.GetOrderProductResult{}) {
		//jika belum ada, buat record order product baru
		newOrderProductID, err := s.orderRepo.CreateOrderProduct(ctx, tx, orderID, product.ID)
		if err != nil {
			tx.Rollback()
			return dto.Order{}, err
		}
		orderProductID = newOrderProductID
		// orderProductID
	} else {
		//jika sudah ada, ambil order oroduct idnya, cek quantity dan stock
		orderProductID = orderProduct.ID
		//jika quantity tidak melebihi stock, increase quantity product
		//jika quantity > error, stok tidak cukup
		if orderProduct.Quantity >= product.Stock {
			// rollback
			tx.Rollback()
			return dto.Order{}, err
		} else {
			// increase order product qty
			_, err := s.orderRepo.IncreaseOrderProductQuantity(ctx, tx, customerId, orderID, product.ID)
			if err != nil {
				tx.Rollback()
				return dto.Order{}, err
			}
		}
		fmt.Print(orderProductID)
	}
	// dapatin data order yang baru saja dibuat by ID

	newOrderProduct, err := s.orderRepo.GetOrderById(ctx, tx, orderID)
	if err != nil {
		tx.Rollback()
		return dto.Order{}, err
	}

	tx.Commit()

	formattedRes := dto.ConvertToGetOrderSchema([]dto.OrderDetails{newOrderProduct})

	return formattedRes, nil
}

func (s *orderService) GetCustomerCart(ctx context.Context, customerId string) (dto.GetUserCartResults, error) {
	res, err := s.orderRepo.GetUserCartByCustomer(ctx, nil, customerId)
	if err != nil {
		return dto.GetUserCartResults{}, err
	}

	return res, nil
}