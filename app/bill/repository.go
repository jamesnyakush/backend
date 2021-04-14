package bill

import "github.com/nyumbapoa/backend/app/storage"

type Repository interface {

}

// NewRepository creates and returns a new instance of admin repository
func NewRepository(database *storage.Database) Repository {
	return &repository{
		db: database,
	}
}

type repository struct {
	db *storage.Database
}
