package service

import (
	"context"
	"errors"
	"go-foodease-be/dto"
	"go-foodease-be/helpers"
	"go-foodease-be/models"
	"go-foodease-be/repository"
)

type (
	CustomerService interface {
		Register(ctx context.Context, req dto.CustomerRegisterRequest) (dto.CustomerResponse, error)
		GetCustomerById(ctx context.Context, customerId string) (dto.CustomerResponse, error)
		GetCustomerByEmail(ctx context.Context, email string) (dto.CustomerResponse, error)
		VerifyLogin(ctx context.Context, req dto.CustomerLoginRequest) (dto.CustomerLoginResponse, error)
		DeleteAccount(ctx context.Context, id string) error
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

func (s *customerService) Register(ctx context.Context, req dto.CustomerRegisterRequest) (dto.CustomerResponse, error) {
	_, flag, _ := s.customerRepo.CheckEmail(ctx, nil, req.Email)
	if flag {
		return dto.CustomerResponse{}, errors.New("email already exist")
	}

	newCustomer := models.Customer{
		Email: req.Email,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Password: req.Password,
		ActiveAddressId: nil,
	}

	custReg, err := s.customerRepo.RegisterCustomer(ctx, nil, newCustomer)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	return dto.CustomerResponse{
		ID: custReg.ID,
		Email: custReg.Email,
		FirstName: custReg.FirstName,
		LastName: custReg.LastName,
		ActiveAddressId: custReg.ActiveAddressId,
	}, nil
}

func (s *customerService) DeleteAccount(ctx context.Context, id string) error {
	err := s.customerRepo.DeleteAccount(ctx, nil, id)
	if err != nil {
		return err
	}

	return nil
}