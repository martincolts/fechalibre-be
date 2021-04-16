package main

import (
	"github.com/google/wire"
	"tincho.example/conf"
	"tincho.example/database"
	"tincho.example/injector"
	"tincho.example/services"
)

func InitializeEvent(filePath string) *injector.Event {
	wire.Build(
		database.NewCategoryRepo,
		database.NewPlayerRepo,
		database.NewDatabase,
		services.NewCategoryService,
		services.NewPlayerService,
		injector.NewEvent,
		conf.NewConf,
	)
	return &injector.Event{}
}
