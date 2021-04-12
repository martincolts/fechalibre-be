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
		services.NewCategoryService,
		database.NewDatabase,
		injector.NewEvent,
	)
	return &injector.Event{}
}
