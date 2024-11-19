package dto

import "github.com/google/uuid"

type (
	CustomerResponse struct {
		ID uuid.UUID `json:"id"`
		Email string `json:"email"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		ActiveAddressId *uuid.UUID `json:"active_address_id,omitempty"`
	}

	CustomerLoginRequest struct {	
		Email string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`	
	}

	CustomerLoginResponse struct {
		Token string `json:"token"`
		ID string `json:"id"`
	}

	CustomerRegisterRequest struct {
		Email string `json:"email" form:"email" binding:"required"`
		FirstName string `json:"first_name" form:"first_name" binding:"required"`
		LastName string `json:"last_name" form:"last_name" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
		ActiveAddressId *uuid.UUID `json:"active_address_id"`
	}
)