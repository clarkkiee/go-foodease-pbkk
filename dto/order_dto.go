package dto

import (
	
	"go-foodease-be/types"
	"time"
)

type GetOrderProductResult struct {
	CustomerID string `json:"customer_id"`
	ID         string `json:"id"`
	Quantity   uint64 `json:"quantity"`
}

type AddToCartSchema struct {
	ProductId string `json:"product_id"`
}

type OrderDetails struct {
	ID                   string       `json:"id"`
	Status               string    `json:"status"`
	StoreID              string       `json:"store_id"`
	StoreName            string    `json:"store_name"`
	CustomerID           string       `json:"customer_id"`
	CustomerEmail        string    `json:"customer_email"`
	CustomerName         string    `json:"customer_name"`
	AddressStreet        string    `json:"address_street"`
	Coordinates          types.Coordinates    `json:"coordinates"`
	ProductID            string       `json:"product_id"`
	OrderProductID       string       `json:"order_product_id"`
	OrderProductSelected bool      `json:"selected"`
	OrderProductQuantity uint64       `json:"quantity"`
	ProductName          string    `json:"product_name"`
	PriceBefore          float64   `json:"price_before"`
	PriceAfter           float64   `json:"price_after"`
	Stock                uint64       `json:"stock"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type GetOrderSchema struct {
	ID string `json:"id"`
	Status string `json:"status"`
	Store StoreResponse `json:"store"`
	Customer CustomerResponse `json:"customer"`
	
}

type Coordinates struct {
	Longitude float64 `json:"Longitude"`
	Latitude  float64 `json:"Latitude"`
}

type Product struct {
	ID              string  `json:"product_id"`
	Name            string  `json:"product_name"`
	Selected        bool    `json:"order_product_selected"`
	Quantity        uint64     `json:"order_product_quantity"`
	PriceBefore     float64 `json:"price_before"`
	PriceAfter      float64 `json:"price_after"`
	Stock           uint64     `json:"stock"`
	ImageURL        *string `json:"image_url"`
}

type Address struct {
	Street   string  `json:"address_street"`
	Longitude float64 `json:"Longitude"`
	Latitude  float64 `json:"Latitude"`
}

type Customer struct {
	ID           string  `json:"customer_id"`
	Email        string  `json:"customer_email"`
	DisplayName  string  `json:"customer_name"`
	Address      Address `json:"address"`
}

type Store struct {
	ID          string `json:"store_id"`
	DisplayName string `json:"store_name"`
}

type Order struct {
	ID          string   `json:"id"`
	Status      string   `json:"status"`
	Store       Store    `json:"store"`
	Customer    Customer `json:"user"`
	Products    []Product `json:"products"`
	TotalPrice  float64  `json:"total_price"`
}

func ConvertToGetOrderSchema(orderDetails []OrderDetails) Order {
	var totalPrice float64

	var products []Product

	if orderDetails[0].Status == "in-cart-unselected" {
		totalPrice = 0
	} else {
		for _, order := range orderDetails {
			if order.OrderProductSelected {
				totalPrice += order.PriceAfter * float64(order.OrderProductQuantity)
			}

			products = append(products, Product{
				ID:          order.ProductID,
				Name:        order.ProductName,
				Selected:    order.OrderProductSelected,
				Quantity:    order.OrderProductQuantity,
				PriceBefore: order.PriceBefore,
				PriceAfter:  order.PriceAfter,
				Stock:       order.Stock,
			})
		}
	}

	orderSchema := Order{
		ID:     orderDetails[0].ID,
		Status: orderDetails[0].Status,
		Store: Store{
			ID:          orderDetails[0].StoreID,
			DisplayName: orderDetails[0].StoreName,
		},
		Customer: Customer{
			ID:          orderDetails[0].CustomerID,
			Email:       orderDetails[0].CustomerEmail,
			DisplayName: orderDetails[0].CustomerName,
			Address: Address{
				Street:   orderDetails[0].AddressStreet,
				Longitude: orderDetails[0].Coordinates.Longitude,
				Latitude:  orderDetails[0].Coordinates.Latitude,
			},
		},
		Products:   products,
		TotalPrice: totalPrice,
	}

	return orderSchema
}

type ProductDetailsInCart struct {
	ID string `json:"id"`
	Selected bool `json:"selected"`
	Quantity uint64 `json:"quantity"`
	ProductName string `json:"product_name"`
	PriceBefore float64 `json:"price_before"`
	PriceAfter float64 `json:"price_after"`
	Stock uint64 `json:"stock"`
	ImageUrl string `json:"image_url"`
}


type OrderCart struct {
	ID string `json:"id"`
	Status string `json:"status"`
	Store Store `json:"store"`
	Products []ProductDetailsInCart `json:"products"`
	TotalPrice float64 `json:"total_price"`
}

type GetUserCartResults struct {
	Orders     []OrderCart `json:"orders"`
	TotalPrice float64 `json:"total_price"`
}