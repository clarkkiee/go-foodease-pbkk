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
	OrderProductSelected bool      `json:"order_product_selected"`
	OrderProductQuantity uint64       `json:"order_product_quantity"`
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

func convertToGetOrderSchema(order OrderDetails){
	
}