package database

type PlayerRole string

const (
	ADMIN PlayerRole = "ADMIN"
	USER  PlayerRole = "USER"
)

type Player struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id        int64 `db:"id"`
	Created   int64
	Birthdate int64
	Name      string     `db:",size:50"`   // Column size set to 50
	Lastname  string     `db:",size:50"`   // Column size set to 50
	DNI       string     `db:",size:1024"` // Set both column name and size
	Password  string     `db:",size:1024"`
	Role      PlayerRole `db:",size:8"`
	Username  string     `db:",size:50"`
}
