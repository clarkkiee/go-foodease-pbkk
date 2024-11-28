package service

import (
	"context"
	"go-foodease-be/dto"
	"go-foodease-be/models"
	"go-foodease-be/repository"
	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req dto.CreateProductRequest, storeID string) (dto.ProductResponse, error)
}

type productService struct {
	productRepo repository.ProductRepository
	jwtService JWTService

}


func NewProductService(productRepo repository.ProductRepository, jwtService JWTService) ProductService {
	return &productService{
		productRepo: productRepo,
		jwtService: jwtService,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req dto.CreateProductRequest, storeID string) (dto.ProductResponse, error) {
	product := models.Product{
		ProductName:   req.ProductName,
		Description:   req.Description,
		PriceBefore:   req.PriceBefore,
		PriceAfter:    req.PriceAfter,
		ProductionTime: req.ProductionTime,
		ExpiredTime:   req.ExpiredTime,
		Stock:         req.Stock,
		CategoryID:    uuid.MustParse(req.CategoryID),
		StoreId:       uuid.MustParse(storeID),
	}

	if req.ImageID != nil {
		imageID := uuid.MustParse(*req.ImageID)
		product.ImageID = &imageID
	}

	createdProduct, err := s.productRepo.CreateProduct(ctx, product, storeID)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	// Handle ImageID as a string or nil
	var imageIDStr *string
	if createdProduct.ImageID != nil {
		str := createdProduct.ImageID.String()
		imageIDStr = &str
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
		ImageID:       imageIDStr,  // Correct handling of *string
		CreatedAt:     createdProduct.CreatedAt,
		UpdatedAt:     createdProduct.UpdatedAt,
	}, nil
}

