package database

type PlayerRole string

const (
	ADMIN PlayerRole = "ADMIN"
	USER  PlayerRole = "USER"
)

type Player struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id        int64      `gorm:"id"`
	Created   int64      `gorm:"not null"`
	Birthdate int64      `gorm:"not null"`
	Name      string     `gorm:"not null;size:50"`   // Column size set to 50
	Lastname  string     `gorm:"not null;size:50"`   // Column size set to 50
	DNI       string     `gorm:"not null;size:1024"` // Set both column name and size
	Password  string     `gorm:"not null;size:1024" json:",omitempty"`
	Role      PlayerRole `gorm:"not null;size:8"`
	Username  string     `gorm:"not null;unique;size:50"`
}
