package database

type CategoryRepo struct {
	database *Database
}

func NewCategoryRepo(db *Database) *CategoryRepo {
	return &CategoryRepo{database: db}
}

func (cr *CategoryRepo) Insert(category *Category) (*Category, error) {
	result := cr.database.db.Create(category)
	if result.Error == nil {
		return category, nil
	} else {
		return nil, result.Error
	}
}
