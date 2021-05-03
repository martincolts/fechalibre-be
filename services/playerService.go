package services

import "tincho.example/database"

type PlayerService struct {
	playerRepo *database.PlayerRepo
}

func NewPlayerService(pr *database.PlayerRepo) *PlayerService {
	return &PlayerService{playerRepo: pr}
}

func (pr *PlayerService) Insert(player *database.Player) (*database.Player, error) {
	return pr.playerRepo.Insert(player)
}

func (pr *PlayerService) InsertAdmin(player *database.Player) (*database.Player, error) {
	return pr.playerRepo.InsertAdmin(player)
}

func (pr *PlayerService) GetAllPlayers() (*[]database.Player, error) {
	return pr.playerRepo.GetAllPlayers()
}

func (pr *PlayerService) GetPlayerById(id int) (*database.Player, error) {
	return pr.playerRepo.GetPlayerById(id)
}

func (pr *PlayerService) GetPlayerByUsername(username string) (*database.Player, error) {
	return pr.playerRepo.GetPlayerByUsername(username)
}
