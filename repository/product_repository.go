package repository

import (
	"context"
	"go-foodease-be/dto"
	"go-foodease-be/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product models.Product, storeID string) (models.Product, error)
	UpdateProduct(ctx context.Context, productID string, updatedProduct models.Product, storeID string) (uuid.UUID, error) 
	DeleteProduct(ctx context.Context, productID string, storeID string) error
	GetMinimumProduct(ctx context.Context, tx *gorm.DB, productId string) (dto.GetMinimumProductResult, error)
	GetProductById(ctx context.Context, tx *gorm.DB, productId string) (models.Product, error)
	GetProductByStoreId(ctx context.Context, tx *gorm.DB, storeId string) ([]models.Product, error)
	GetNearestProduct(ctx context.Context, tx *gorm.DB, customerCoord dto.CoordinatesResponse, limit string, offset string, distance string) ([]models.Product, error)
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
	if err := tx.WithContext(ctx).Model(&models.Product{}).Where("id = ?", productId).Take(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *productRepository) GetMinimumProduct(ctx context.Context, tx *gorm.DB, productId string) (dto.GetMinimumProductResult, error) {
	if tx == nil {
		tx = r.db
	}

	var product models.Product
	if err := tx.WithContext(ctx).Model(&models.Product{}).Select("id, stock, store_id, category_id").Where("id = ?", productId).Take(&product).Error; err != nil {
		return dto.GetMinimumProductResult{}, err
	}

	res := dto.GetMinimumProductResult{
		ID: product.ID.String(),
		Stock: product.Stock,
		StoreID: product.StoreId.String(),
		CategoryID: product.CategoryID.String(),
	}

	return res, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, productID string, storeID string) error {
    var product models.Product
    if err := r.db.WithContext(ctx).Where("id = ? AND store_id = ?", productID, storeID).First(&product).Error; err != nil {
        return err
    }

    if err := r.db.WithContext(ctx).Delete(&product).Error; err != nil {
        return err
    }

    return nil
}

func (r *productRepository) GetProductByStoreId(ctx context.Context, tx *gorm.DB, storeId string) ([]models.Product, error) {

	if tx == nil {
		tx = r.db
	}

	var product []models.Product
	if err := tx.WithContext(ctx).Model(&models.Product{}).Where("store_id = ?", storeId).Find(&product).Scan(&product).Error; err != nil {
		return []models.Product{}, err
	}

	return product, nil
}

func (r *productRepository) GetNearestProduct(ctx context.Context, tx *gorm.DB, customerCoord dto.CoordinatesResponse, limit string, offset string, distance string) ([]models.Product, error){
	if tx == nil {
		tx = r.db
	}

	var product []models.Product
	err := tx.Raw(`SELECT
		p.id,
		p.product_name,
		p.description,
		p.price_before,
		p.price_after,
		p.production_time,
		p.expired_time,
		p.stock,
		p.image_id,
		s.store_name,
		a.street,
		ST_X(a.coordinates::geometry) as "address_longitude",
		ST_Y(a.coordinates::geometry) as "address_latitude",
		ST_DISTANCE(a.coordinates, :user_coordinates) as "address_distance",
		c.slug,
		c.category_name,
		p.updated_at,
		p.created_at
	FROM product p 
	INNER JOIN store s ON p.store_id = s.id
	INNER JOIN address a ON a.id = s.address_id
	INNER JOIN category c ON c.id = p.category_id
	WHERE 
		ST_DISTANCE(a.coordinates, :?) < :?
	ORDER BY
		ST_DISTANCE(a.coordinates, :?) ASC,
		p.id ASC,
		p.updated_at DESC
	LIMIT :? OFFSET :?;`, customerCoord, distance, customerCoord, limit, offset).Scan(&product).Error

	if err != nil {
		return []models.Product{}, err
	}

	return product, nil
}
