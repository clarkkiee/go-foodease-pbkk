package repository

import (
	"context"
	"go-foodease-be/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	CategoryRepository interface {
		GetCategoryIdBySlug(ctx context.Context, tx *gorm.DB, slug string) (uuid.UUID, error)
	}

	categoryRepository struct {
		db *gorm.DB
	}
)
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetCategoryIdBySlug(ctx context.Context, tx *gorm.DB, slug string) (uuid.UUID, error) {
	if tx == nil {
		tx = r.db
	}

	var category models.Category
	if err := tx.WithContext(ctx).Where("slug = ?", slug).First(&category).Error; err != nil {
		return uuid.Nil, err
	}

	return category.ID, nil

}