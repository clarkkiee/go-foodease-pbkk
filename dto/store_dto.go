package dto

import (
	// "time"



	"github.com/google/uuid"
)

type (
	StoreResponse struct {
		ID uuid.UUID `json:"id"`
		Email string `json:"email"`
		StoreName string `json:"store_name"`
		AddressId *uuid.UUID `json:"address_id,omitempty"`
	}

	StoreLoginRequest struct {	
		Email string `json:"email" form:"email" binding:"required"`
		StorePassword string `json:"store_password" form:"store_password" binding:"required"`	
	}

	StoreLoginResponse struct {
		Token string `json:"token"`
		ID string `json:"id"`
	}

	StoreRegisterRequest struct {
		Email string `json:"email" form:"email" binding:"required"`
		StoreName string `json:"store_name" form:"store_name" binding:"required"`
		StorePassword string `json:"store_password" form:"store_password" binding:"required"`
		StorePasswordConfirm string `json:"store_password_confirm" form:"store_password_confirm" binding:"required"`
		FreeTime string `json:"free_time" form:"free_time" binding:"required"`
		Address CreateNewAddressRequest `json:"address" form:"address" binding:"required"`
	}
)