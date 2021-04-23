package auth

import (
	"github.com/gofrs/uuid"
	"github.com/jackc/pgconn"
	"github.com/nyumbapoa/backend/app/errors"
	"github.com/nyumbapoa/backend/app/model"
	"github.com/nyumbapoa/backend/app/storage"
)

type Repository interface {
	Add(model.User) (model.User, error)
	GetByID(uuid.UUID) (model.User, error)
	GetByEmail(string) (model.User, error)
	Update(model.User) error
	Delete(model.User) error
}

func NewRepository(database *storage.Database) Repository {
	return &repository{
		db: database,
	}
}

type repository struct {
	db *storage.Database
}

func (r repository) Add(user model.User) (model.User, error) {
	// add new admin to administrators table, if query return violation of unique key column,
	// we know that the admin with given record exists and return that admin instead
	result := r.db.Model(model.User{}).Create(&user)
	if err := result.Error; err != nil {
		// we check if the error is a postgres unique constraint violation
		if pgerr, ok := err.(*pgconn.PgError); ok && pgerr.Code == "23505" {
			return user, errors.Error{Code: errors.ECONFLICT, Message: errors.ErrUserExists}
		}
		return model.User{}, errors.Error{Err: result.Error, Code: errors.EINTERNAL}
	}

	return user, nil
}

func (r repository) GetByID(u uuid.UUID) (model.User, error) {
	panic("implement me")
}

func (r repository) GetByEmail(email string) (model.User, error) {
	panic("implement me")
}

func (r repository) Update(user model.User) error {
	panic("implement me")
}

func (r repository) Delete(user model.User) error {
	panic("implement me")
}
