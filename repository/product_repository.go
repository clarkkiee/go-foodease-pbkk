package repository

import (
	"context"
	"go-foodease-be/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product models.Product, storeID string) (models.Product, error)
	UpdateProduct(ctx context.Context, productID string, updatedProduct models.Product, storeID string) (uuid.UUID, error)
	GetProductById(ctx context.Context, tx *gorm.DB, productId string) (models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
func (r *productRepository) CreateProduct(ctx context.Context, product models.Product, storeID string) (models.Product, error) {
	if err := r.db.WithContext(ctx).Create(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, productID string, updatedProduct models.Product, storeID string) (uuid.UUID, error) {
	var product models.Product
	
	if err := r.db.WithContext(ctx).Where("id = ? AND store_id = ?", productID, storeID).First(&product).Error; err != nil {
		return uuid.Nil, err 
	}

	if err := r.db.WithContext(ctx).Model(&product).Updates(map[string]interface{}{
		"product_name":     updatedProduct.ProductName,
		"description":      updatedProduct.Description,
		"price_before":     updatedProduct.PriceBefore,
		"price_after":      updatedProduct.PriceAfter,
		"production_time":  updatedProduct.ProductionTime,
		"expired_time":     updatedProduct.ExpiredTime,
		"stock":            updatedProduct.Stock,
		"category_id":      updatedProduct.CategoryID,
		"image_id":         updatedProduct.ImageID,
	}).Error; err != nil {
		return uuid.Nil, err
	}

	return product.ID, nil
}

func (r *productRepository) GetProductById(ctx context.Context, tx *gorm.DB, productId string) (models.Product, error) {
	if tx == nil {
		tx = r.db
	}

	var product models.Product
	
	if err := tx.WithContext(ctx).Where("id = ?", productId).Take(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}