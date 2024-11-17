package service

import (
	"context"
	"go-foodease-be/dto"
	"go-foodease-be/repository"
)

type (
	CustomerService interface {
		GetCustomerById(ctx context.Context, customerId string) (dto.CustomerResponse, error)
		GetCustomerByEmail(ctx context.Context, email string) (dto.CustomerResponse, error)
	}

	customerService struct {
		customerRepo repository.CustomerRepository
	}
)

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

func (s *customerService) GetCustomerById(ctx context.Context, customerId string) (dto.CustomerResponse, error) {
	customer, err := s.customerRepo.GetCustomerById(ctx, nil, customerId)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	return dto.CustomerResponse{
		ID: customer.ID,
		Email: customer.Email,
		FirstName: customer.FirstName,
		LastName: customer.LastName,
		ActiveAddressId: customer.ActiveAddressId,
	}, nil
}

func (s *customerService) GetCustomerByEmail(ctx context.Context, email string) (dto.CustomerResponse, error) {
	customer, err := s.customerRepo.GetCustomerByEmail(ctx, nil, email)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	return dto.CustomerResponse{
		ID: customer.ID,
		Email: customer.Email,
		FirstName: customer.FirstName,
		LastName: customer.LastName,
		ActiveAddressId: customer.ActiveAddressId,
	}, nil
}
