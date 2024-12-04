package service

import (
	"context"
	"errors"
	"go-foodease-be/repository"

	"github.com/google/uuid"
)

type (
	CategoryService interface {
		GetCategoryIdBySlug(ctx context.Context, slug string) (uuid.UUID, error)
	}

	categoryService struct {
		categoryRepo repository.CategoryRepository
	}
)

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) GetCategoryIdBySlug(ctx context.Context, slug string) (uuid.UUID, error) {
	if slug == "" {
		return uuid.Nil, errors.New("missing slug")
	}

	categoryId, err := s.categoryRepo.GetCategoryIdBySlug(ctx, nil, slug)
	if err != nil {
		return uuid.Nil, errors.New("failed to get category id")
	}

	return categoryId, nil

}