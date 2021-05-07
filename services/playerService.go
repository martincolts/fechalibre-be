package services

import (
	"errors"

	"tincho.example/database"
	"tincho.example/dtos"
)

type PlayerService struct {
	playerRepo *database.PlayerRepo
}

func NewPlayerService(pr *database.PlayerRepo) *PlayerService {
	return &PlayerService{playerRepo: pr}
}

func (ps *PlayerService) Insert(player *database.Player) (*database.Player, error) {
	return ps.playerRepo.Insert(player)
}

func (ps *PlayerService) InsertAdmin(player *database.Player) (*database.Player, error) {
	return ps.playerRepo.InsertAdmin(player)
}

func (ps *PlayerService) GetAllPlayers() (*[]database.Player, error) {
	return ps.playerRepo.GetAllPlayers()
}

func (ps *PlayerService) GetPlayerById(id int64) (*database.Player, error) {
	return ps.playerRepo.GetPlayerById(id)
}

func (ps *PlayerService) GetPlayerByUsername(username string) (*database.Player, error) {
	return ps.playerRepo.GetPlayerByUsername(username)
}

func (ps *PlayerService) UpdatePassword(dto dtos.UpdatePasswordPayload, username string) (*database.Player, error) {

	player, error := ps.GetPlayerByUsername(username)
	if error != nil {
		return nil, error
	}
	if player.Password != dto.OldPassword {
		return nil, errors.New("incorrect current password")
	}
	return ps.playerRepo.UpdatePassword(dto.NewPassword, username)
}
