package dto

type GetOrderProductResult struct {
	CustomerID string `json:"customer_id"`
	ID string `json:"id"`
	Quantity uint64 `json:"quantity"`
}

type AddToCartSchema struct {
	ProductId string `json:"product_id"`
}