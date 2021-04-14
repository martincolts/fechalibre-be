package database

import "time"

type PlayerRepo struct {
	database *Database
}

func NewPlayerRepo(db *Database) *PlayerRepo {
	return &PlayerRepo{database: db}
}

func (pr *PlayerRepo) Insert(player *Player) (*Player, error) {
	p := *player
	p.Created = time.Now().Unix()
	result := pr.database.db.Create(&p)
	if result.Error == nil {
		return &p, nil
	} else {
		return nil, result.Error
	}
}

func (pr *PlayerRepo) GetAllPlayers() (error, *[]Player) {
	var players []Player
	if result := pr.database.db.Find(&players); result.Error == nil {
		return nil, &players
	} else {
		return result.Error, nil
	}
}

func (pr *PlayerRepo) GetPlayerById(id int) (error, *Player) {
	var players Player
	if result := pr.database.db.Find(&players, id); result.Error == nil {
		return nil, &players
	} else {
		return result.Error, nil
	}
}

func (pr *PlayerRepo) GetPlayerByUsername(username string) (*Player, error) {
	var player Player
	if result := pr.database.db.Where("username = ?", username).First(&player); result.Error == nil {
		return &player, nil
	} else {
		return nil, result.Error
	}
}
