package postgres

import (
	"github.com/nyumbapoa/backend/app"
	"github.com/nyumbapoa/backend/app/storage"
)

// NewDatabase creates a new Database object
func NewDatabase(config app.Config) (*storage.Database, error) {

	var err error

	// var db *storage.Database
	db := new(storage.Database)

	/*	var conn *gorm.DB
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s",
			config.DB.User, config.DB.Password, config.DB.DBName, config.DB.Host, config.DB.Port,
		)
		conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	*/

	if err != nil {
		return nil, err
	}
	//db.DB = conn

	return db, nil
}
