package injector

import (
	"tincho.example/database"
	"tincho.example/services"
)

type Event struct {
	categoryService *services.CategoryService
	playerService   *services.PlayerService
	database        *database.Database
}

func NewEvent(
	cs *services.CategoryService,
	ps *services.PlayerService,
	database *database.Database,
) *Event {
	return &Event{
		categoryService: cs,
		playerService:   ps,
		database:        database,
	}
}

func (e Event) Start() {

}

func (e Event) GetCategoryService() *services.CategoryService {
	return e.categoryService
}

func (e Event) GetPlayerService() *services.PlayerService {
	return e.playerService
}

func (e Event) GetDatabase() *database.Database {
	return e.database
}

func InitializeEvent() *Event {
	databaseDatabase := database.NewDatabase()
	categoryRepo := database.NewCategoryRepo(databaseDatabase)
	categoryService := services.NewCategoryService(categoryRepo)
	playerRepo := database.NewPlayerRepo(databaseDatabase)
	playerService := services.NewPlayerService(playerRepo)
	event := NewEvent(categoryService, playerService, databaseDatabase)
	return event
}
