package auth

import "github.com/nyumbapoa/backend/app/storage"

type Repository interface {
}

func NewRepository(database *storage.Database) Repository {
	return &repository{
		db: database,
	}
}

type repository struct {
	db *storage.Database
}
