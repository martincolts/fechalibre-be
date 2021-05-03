package services

import (
	"tincho.example/database"
)

type CategoryService struct {
	repo *database.CategoryRepo
}

func NewCategoryService(r *database.CategoryRepo) *CategoryService {
	return &CategoryService{repo: r}
}

func (categoryService *CategoryService) Insert(category *database.Category) (*database.Category, error) {
	return categoryService.repo.Insert(category)
}
