package repository

import (
	"context"
	"go-foodease-be/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product models.Product, storeID string) (models.Product, error)
	UpdateProduct(ctx context.Context, productID string, updatedProduct models.Product, storeID string) (models.Product, error) // <-- Tambahkan ini
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

func (r *productRepository) UpdateProduct(ctx context.Context, productID string, updatedProduct models.Product, storeID string) (models.Product, error) {
	var product models.Product
	
	
	if err := r.db.WithContext(ctx).Where("id = ? AND store_id = ?", productID, storeID).First(&product).Error; err != nil {
		return models.Product{}, err 
	}

	// Perbarui data produk menggunakan Updates() untuk menghindari perubahan satu per satu
	if err := r.db.WithContext(ctx).Model(&product).Updates(map[string]interface{}{
		"product_name":     updatedProduct.ProductName,
		"description":      updatedProduct.Description,
		"price_before":     updatedProduct.PriceBefore,
		"price_after":      updatedProduct.PriceAfter,
		//"production_time":  updatedProduct.ProductionTime,
		//"expired_time":     updatedProduct.ExpiredTime,
		"stock":            updatedProduct.Stock,
		"category_id":      updatedProduct.CategoryID,
		"image_id":         updatedProduct.ImageID,
	}).Error; err != nil {
		return models.Product{}, err
	}

	// Kembalikan produk yang telah diperbarui
	return product, nil
}
