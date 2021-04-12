package services

import (
	"github.com/gin-gonic/gin"
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

func Insert() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
