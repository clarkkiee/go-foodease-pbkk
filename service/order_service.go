package service

import (
	"context"
	"fmt"
	"go-foodease-be/repository"
)

type (
	OrderService interface {
		AddToCart(ctx context.Context, storeId string, customerId string) ()
	}

	orderService struct {
		orderRepo repository.OrderRepository
	}
)

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (s *orderService) AddToCart(ctx context.Context, storeId string, customerId string) () {
	
	//cek produk ada atau tidak, cek juga stock nya

	//cek apakah sudah ada order di cart (by customer id dan store id)
	order, err := s.orderRepo.GetOrderInCart(ctx, nil, storeId, customerId)
	if err != nil {
		return
	}

	//jika tidak ada record, buat record order baru 
	//jika sudah ada, ambil order id nya saja
	orderID := ""
	if order == "" || len(order) == 0 {
		//buat record order 
		s.orderRepo
	} else {
		//simpan order id nya
		orderID = order
	}

	fmt.Println("DARI SERVICE NI BOUSSS ", orderID)

	//cek data di order apakah sudah ada produk yang dimaksud di order cart

	//jika belum ada, buat record order product baru

	//jika sudah ada, ambil order oroduct idnya, cek quantity dan stock

	//jika quantity tidak melebihi stock, increase quantity product

	//jika quantity > error, stok tidak cukup

	//terakhir -> dapatkan data order yang baru saja dibuat by ID
}