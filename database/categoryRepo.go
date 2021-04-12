package database

type CategoryRepo struct {
	database *Database
}

func NewCategoryRepo(db *Database) *CategoryRepo {
	return &CategoryRepo{database: db}
}

func (cr *CategoryRepo) Insert(category *Category) (error, *Category) {
	result := cr.database.db.Create(category)
	if result.Error == nil {
		return nil, category
	} else {
		return result.Error, nil
	}
}
