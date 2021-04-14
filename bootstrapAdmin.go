package main

import (
	"time"

	"tincho.example/database"
	"tincho.example/injector"
)

var adminPlayer database.Player = database.Player{
	Created:   time.Now().Unix(),
	Birthdate: 0,
	Name:      "admin",
	Lastname:  "admin",
	DNI:       "dniAdmin",
	Password:  "adminPassword",
	Role:      database.ADMIN,
	Username:  "admin",
}

func addAdminUser(event *injector.Event) {
	playerService := event.GetPlayerService()
	if player, _ := playerService.GetPlayerByUsername(adminPlayer.Username); player == nil {
		playerService.Insert(&adminPlayer)
	}
}
