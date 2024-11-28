package service

import (
	"context"
	// "errors"
	"go-foodease-be/dto"
	"go-foodease-be/helpers"
	// "go-foodease-be/models"
	"go-foodease-be/repository"
)

type (
	StoreService interface {
		// Register(ctx context.Context, req dto.CustomerRegisterRequest) (dto.CustomerResponse, error)
		// GetCustomerById(ctx context.Context, customerId string) (dto.CustomerResponse, error)
		// GetCustomerByEmail(ctx context.Context, email string) (dto.CustomerResponse, error)
		VerifyLogin(ctx context.Context, req dto.StoreLoginRequest) (dto.StoreLoginResponse, error)
		// DeleteAccount(ctx context.Context, id string) error
	}

	storeService struct {
		storeRepo repository.StoreRepository
		jwtService JWTService
	}
)

func NewStoreService(storeRepo repository.StoreRepository, jwtService JWTService) StoreService {
	return &storeService{
		storeRepo: storeRepo,
		jwtService: jwtService,
	}
}

func (s *storeService) VerifyLogin(ctx context.Context, req dto.StoreLoginRequest) (dto.StoreLoginResponse, error) {
	store, flag, err := s.storeRepo.CheckEmail(ctx, nil, req.Email)
	if err!= nil || !flag {
		return dto.StoreLoginResponse{}, err
	}
	
	checkPassword, err := helpers.ValidatePassword(req.StorePassword, store.StorePassword)
	if err != nil || !checkPassword {
		return dto.StoreLoginResponse{}, err
	}


	token := s.jwtService.GenerateToken(store.ID.String())

	return dto.StoreLoginResponse{
		Token: token,
		ID: store.ID.String(),
	}, nil
}
