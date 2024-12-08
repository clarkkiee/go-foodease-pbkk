package service

import (
	"context"
	"go-foodease-be/dto"
	"go-foodease-be/models"
	"go-foodease-be/repository"

	"github.com/google/uuid"
)

type (
	ProductService interface {
		CreateProduct(ctx context.Context, req dto.CreateProduct, storeID string) (dto.ProductResponse, error)
		UpdateProduct(ctx context.Context, productID string, req dto.UpdateProductRequest, storeID string) (uuid.UUID, error)
		GetMinimumProduct(ctx context.Context, productID string) (dto.GetMinimumProductResult, error)
		DeleteProduct(ctx context.Context, productID string, storeID string) error  
		GetProductById(ctx context.Context, productId string) (dto.ProductResponse, error)
		GetProductByStoreId(ctx context.Context, storeId string) ([]dto.ProductResponse, error)
		GetNearestProduct(ctx context.Context, customerCoord dto.CoordinatesResponse, limit string, offset string, distance string) ([]dto.ProductResponse, error)
	}

	productService struct {
		productRepo repository.ProductRepository
		jwtService  JWTService
	}
)

func NewProductService(productRepo repository.ProductRepository, jwtService JWTService) ProductService {
	return &productService{
		productRepo: productRepo,
		jwtService:  jwtService,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req dto.CreateProduct, storeID string) (dto.ProductResponse, error) {
	product := models.Product{
		ProductName:   req.ProductName,
		Description:   req.Description,
		PriceBefore:   req.PriceBefore,
		PriceAfter:    req.PriceAfter,
		ProductionTime: req.ProductionTime,
		ExpiredTime:   req.ExpiredTime,
		Stock:         req.Stock,
		CategoryID:    req.CategoryID,
		StoreId:       uuid.MustParse(storeID),
	}

	createdProduct, err := s.productRepo.CreateProduct(ctx, product, storeID)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		ID:            createdProduct.ID.String(),
		ProductName:   createdProduct.ProductName,
		Description:   createdProduct.Description,
		PriceBefore:   createdProduct.PriceBefore,
		PriceAfter:    createdProduct.PriceAfter,
		ProductionTime: createdProduct.ProductionTime,
		ExpiredTime:   createdProduct.ExpiredTime,
		Stock:         createdProduct.Stock,
		CategoryID:    createdProduct.CategoryID.String(),
		ImageID:       createdProduct.ImageID.String(),
		CreatedAt:     createdProduct.CreatedAt,
		UpdatedAt:     createdProduct.UpdatedAt,
	}, nil
}

func (s *productService) UpdateProduct(ctx context.Context, productID string, req dto.UpdateProductRequest, storeID string) (uuid.UUID, error) {

	parsedImageId, _ := uuid.Parse(req.ImageID)

	updatedProduct := models.Product{
		ProductName:   req.ProductName,
		Description:   req.Description,
		PriceBefore:   req.PriceBefore,
		PriceAfter:    req.PriceAfter,
		ProductionTime: req.ProductionTime,
		ExpiredTime:   req.ExpiredTime,
		CategoryID:    uuid.MustParse(req.CategorySlug),
		ImageID:       parsedImageId,
	}

	_, err := s.productRepo.UpdateProduct(ctx, productID, updatedProduct, storeID)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(productID), nil
}

func (s *productService) GetMinimumProduct(ctx context.Context, productID string) (dto.GetMinimumProductResult, error) {
	result, err := s.productRepo.GetMinimumProduct(ctx, nil, productID)
	if err != nil {
		return dto.GetMinimumProductResult{}, err
	}

	return result, nil
}

func (s *productService) DeleteProduct(ctx context.Context, productID string, storeID string) error {
    err := s.productRepo.DeleteProduct(ctx, productID, storeID)
    if err != nil {
        return err
    }
    return nil
}

func (s *productService) GetProductById(ctx context.Context, productId string) (dto.ProductResponse, error){
	product, err := s.productRepo.GetProductById(ctx, nil, productId)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		ID: product.ID.String(),
		ProductName: product.ProductName,
		Description: product.Description,
		PriceBefore: product.PriceBefore,
		PriceAfter: product.PriceAfter,
		ProductionTime: product.ProductionTime,
		ExpiredTime: product.ExpiredTime,
		Stock: product.Stock,
		CategoryID: product.CategoryID.String(),
		ImageID: product.ImageID.String(),
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}, nil
}

func (s *productService) GetProductByStoreId(ctx context.Context, storeId string) ([]dto.ProductResponse, error) {
	product, err := s.productRepo.GetProductByStoreId(ctx, nil, storeId)
	if err != nil {
		return []dto.ProductResponse{}, err
	}

	var res []dto.ProductResponse

	for _, item := range product {
		temp := dto.ProductResponse {
			ID: item.ID.String(),
			ProductName: item.ProductName,
			Description: item.Description,
			PriceBefore: item.PriceBefore,
			PriceAfter: item.PriceAfter,
			ProductionTime: item.ProductionTime,
			ExpiredTime: item.ExpiredTime,
			Stock: item.Stock,
			CategoryID: item.CategoryID.String(),
			ImageID: item.ImageID.String(),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}

		res = append(res, temp)
	}

	return res, nil
}

func (s *productService) GetNearestProduct(ctx context.Context, customerCoord dto.CoordinatesResponse, limit string, offset string, distance string) ([]dto.ProductResponse, error) {
	product, err := s.productRepo.GetNearestProduct(ctx, nil, customerCoord, limit, offset, distance)
	if err != nil {
		return []dto.ProductResponse{}, err
	}

	var res []dto.ProductResponse

	for _, item := range product {
		temp := dto.ProductResponse {
			ID: item.ID.String(),
			ProductName: item.ProductName,
			Description: item.Description,
			PriceBefore: item.PriceBefore,
			PriceAfter: item.PriceAfter,
			ProductionTime: item.ProductionTime,
			ExpiredTime: item.ExpiredTime,
			Stock: item.Stock,
			CategoryID: item.CategoryID.String(),
			ImageID: item.ImageID.String(),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}

		res = append(res, temp)
	}

	return res, nil
}