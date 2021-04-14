package main

import (
	"github.com/google/wire"
	"tincho.example/database"
	"tincho.example/injector"
	"tincho.example/services"
)

func InitializeEvent() *injector.Event {
	wire.Build(
		database.NewCategoryRepo,
		database.NewPlayerRepo,
		database.NewDatabase,
		services.NewCategoryService,
		services.NewPlayerService,
		injector.NewEvent,
	)
	return &injector.Event{}
}
