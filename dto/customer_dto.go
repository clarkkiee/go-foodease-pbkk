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
)