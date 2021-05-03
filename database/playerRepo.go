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
	p.Password = "1234"
	p.Created = time.Now().Unix()
	result := pr.database.db.Create(&p)
	if result.Error == nil {
		return &p, nil
	} else {
		return nil, result.Error
	}
}

func (pr *PlayerRepo) InsertAdmin(player *Player) (*Player, error) {
	p := *player
	p.Created = time.Now().Unix()
	result := pr.database.db.Create(&p)
	if result.Error == nil {
		return &p, nil
	} else {
		return nil, result.Error
	}
}

func (pr *PlayerRepo) GetAllPlayers() (*[]Player, error) {
	var players []Player
	var playerToReturn []Player
	if result := pr.database.db.Find(&players); result.Error == nil {
		copyPlayers(players, &playerToReturn)
		return &playerToReturn, nil
	} else {
		return nil, result.Error
	}
}

func (pr *PlayerRepo) GetPlayerById(id int) (*Player, error) {
	var players Player
	if result := pr.database.db.Find(&players, id); result.Error == nil {
		players.Password = ""
		return &players, nil
	} else {
		return nil, result.Error
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

func copyPlayers(players []Player, toCopy *[]Player) {
	for _, player := range players {
		newPlayer := Player{
			Id:        player.Id,
			Created:   player.Created,
			Birthdate: player.Birthdate,
			Name:      player.Name,
			Lastname:  player.Lastname,
			DNI:       player.DNI,
			Password:  "",
			Role:      player.Role,
			Username:  player.Username,
		}
		*toCopy = append(*toCopy, newPlayer)
	}
}
