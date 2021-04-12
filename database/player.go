package database

type Player struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id        int64 `db:"id"`
	Created   int64
	Birthdate int64
	Name      string `db:",size:50"`   // Column size set to 50
	Lastname  string `db:",size:50"`   // Column size set to 50
	DNI       string `db:",size:1024"` // Set both column name and size
	Password  string `db:",size:1024"`
}
