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
		GetProductById(ctx context.Context, productId string) (dto.ProductResponse, error)
	}

	productService struct {
		productRepo repository.ProductRepository
		jwtService JWTService
	}
)

func NewProductService(productRepo repository.ProductRepository, jwtService JWTService) ProductService {
	return &productService{
		productRepo: productRepo,
		jwtService: jwtService,
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
		CategoryID: uuid.MustParse(req.CategorySlug),
		ImageID: parsedImageId,
    }

	_, err := s.productRepo.UpdateProduct(ctx, productID, updatedProduct, storeID)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(productID), nil

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