package dto

import (
	"time"

	"github.com/google/uuid"
)

type Position struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Item struct {
	Position Position `json:"position"`
}

type GeocodeAPIResponse struct {
	Items []Item `json:"items"`
}

type (

	CreateNewAddressRequest struct {
		Street    string  `json:"street" binding:"required"`
		Village string `json:"village" binding:"required"`
		SubDistrict string `json:"sub_district" binding:"required"`
		City string `json:"city" binding:"required"`
		Province string `json:"province" binding:"required"`
	}

	CoordinatesResponse struct {
		Longitude float64 `json:"longitude"`
		Latitude float64 `json:"latitude"`
	}

	AddressResponse struct {
		ID uuid.UUID `json:"ID"`
		Street string `json:"street"`
		Longitude float64 `json:"longitude"`
		Latitude float64 `json:"latitude"`	
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)