package service

import (
	"context"
	"errors"
	// "errors"
	"go-foodease-be/dto"
	"go-foodease-be/helpers"

	"go-foodease-be/models"
	"go-foodease-be/repository"
)

type (
	StoreService interface {
		Register(ctx context.Context, req dto.StoreRegisterRequest) (dto.StoreResponse, error)
		GetStoreById(ctx context.Context, customerId string) (dto.StoreResponse, error)
		// GetCustomerByEmail(ctx context.Context, email string) (dto.CustomerResponse, error)
		VerifyLogin(ctx context.Context, req dto.StoreLoginRequest) (dto.StoreLoginResponse, error)
		DeleteAccount(ctx context.Context, id string) error
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

func (s *storeService) GetStoreById(ctx context.Context, storeId string) (dto.StoreResponse, error) {
	store, err := s.storeRepo.GetStoreById(ctx, nil, storeId)
	if err != nil {
		return dto.StoreResponse{}, err
	}

	return dto.StoreResponse{
		ID: store.ID,
		Email: store.Email,
		StoreName:store.StoreName,
		AddressId:store.AddressID,
	}, nil
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

func (s *storeService) DeleteAccount(ctx context.Context, id string) error {

	_, v_err := s.storeRepo.GetStoreById(ctx, nil, id)
	if v_err != nil {
		return v_err
	}

	err := s.storeRepo.DeleteAccount(ctx, nil, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *storeService) Register(ctx context.Context, req dto.StoreRegisterRequest) (dto.StoreResponse, error) {
	_, flag, _ := s.storeRepo.CheckEmail(ctx, nil, req.Email)
	if flag{
		return dto.StoreResponse{}, errors.New("email already exist")
	}

	newStore := models.Store{
		Email: req.Email,
		StoreName: req.StoreName,
		StorePassword: req.StorePassword,
		AddressID: nil,
	}

	storeReg, err := s.storeRepo.RegisterStore(ctx, nil, newStore)
	if err != nil {
		return dto.StoreResponse{}, err
	}

	return dto.StoreResponse{
		ID: storeReg.ID,
		Email: storeReg.Email,
		StoreName: storeReg.StoreName,
		AddressId: storeReg.AddressID,
	}, nil
}
