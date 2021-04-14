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

func (categoryService *CategoryService) Insert(category *database.Category) (error, *database.Category) {
	return categoryService.repo.Insert(category)
}
