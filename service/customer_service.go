package service

import (
	"context"
	"fmt"
	"go-foodease-be/dto"
	"go-foodease-be/helpers"
	"go-foodease-be/repository"
)

type (
	CustomerService interface {
		GetCustomerById(ctx context.Context, customerId string) (dto.CustomerResponse, error)
		GetCustomerByEmail(ctx context.Context, email string) (dto.CustomerResponse, error)
		VerifyLogin(ctx context.Context, req dto.CustomerLoginRequest) (dto.CustomerLoginResponse, error)
	}

	customerService struct {
		customerRepo repository.CustomerRepository
		jwtService JWTService
	}
)

func NewCustomerService(customerRepo repository.CustomerRepository, jwtService JWTService) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
		jwtService: jwtService,
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

func (s *customerService) VerifyLogin(ctx context.Context, req dto.CustomerLoginRequest) (dto.CustomerLoginResponse, error) {
	cust, flag, err := s.customerRepo.CheckEmail(ctx, nil, req.Email)
	if err!= nil || !flag {
		return dto.CustomerLoginResponse{}, err
	}
	
	fmt.Printf("in DB: %v\n", cust)
	checkPassword, err := helpers.ValidatePassword(req.Password, cust.Password)
	if err != nil || !checkPassword {
		return dto.CustomerLoginResponse{}, err
	}


	token := s.jwtService.GenerateToken(cust.ID.String())

	return dto.CustomerLoginResponse{
		Token: token,
		ID: cust.ID.String(),
	}, nil
}
