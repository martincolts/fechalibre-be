package injector

import (
	"tincho.example/database"
	"tincho.example/services"
)

type Event struct {
	categoryService *services.CategoryService
}

func NewEvent(cs *services.CategoryService) *Event {
	return &Event{categoryService: cs}
}

func (e Event) Start() {

}

func (e Event) GetCategoryService() *services.CategoryService {
	return e.categoryService
}

func InitializeEvent() *Event {
	databaseDatabase := database.NewDatabase()
	categoryRepo := database.NewCategoryRepo(databaseDatabase)
	categoryService := services.NewCategoryService(categoryRepo)
	event := NewEvent(categoryService)
	return event
}
