package service

import (
	"context"
	"fmt"
	"go-foodease-be/dto"
	"go-foodease-be/repository"
)

type (
	OrderService interface {
		AddToCart(ctx context.Context, customerId string, productId string) ()
	}

	orderService struct {
		orderRepo repository.OrderRepository
		productRepo repository.ProductRepository
	}
)

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		productRepo: productRepo,
	}
}

func (s *orderService) AddToCart(ctx context.Context, customerId string, productId string) () {
	
	//cek produk ada atau tidak, cek juga stock nya
	product, err := s.productRepo.GetMinimumProduct(ctx, nil, productId)
	if err != nil {
		return
	}

	//cek apakah sudah ada order di cart (by customer id dan store id)
	order, err := s.orderRepo.GetOrderInCart(ctx, nil, product.StoreID, customerId)
	if err != nil {
		return
	}

	//jika tidak ada record, buat record order baru 
	//jika sudah ada, ambil order id nya saja
	orderID := ""
	if orderID == "" || len(order) == 0 {
		//buat record order 
		newOrderId, _ := s.orderRepo.CreateNewOrder(ctx, nil, product.StoreID, customerId)
		orderID = newOrderId
	} else {
		//simpan order id nya
		orderID = order
	}

	//cek data di order apakah sudah ada produk yang dimaksud di order cart
	orderProduct, err := s.orderRepo.GetOrderProduct(ctx, nil, customerId, orderID, product.ID)
	if err != nil {
		return 
	}

	var orderProductID string
	if orderProduct == (dto.GetOrderProductResult{}) {
		//jika belum ada, buat record order product baru
		newOrderProductID, err := s.orderRepo.CreateOrderProduct(ctx, nil, orderID, product.ID)
		if err != nil {
			return
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
			return
		} else {
			// increase order product qty
			s.orderRepo.IncreaseOrderProductQuantity(ctx, nil, customerId, orderID, product.ID)
		}
		fmt.Print(orderProductID)
	}

	//terakhir -> dapatkan data order yang baru saja dibuat by ID
}