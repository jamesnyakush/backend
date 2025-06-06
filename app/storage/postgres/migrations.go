package postgres

import (
	"github.com/nyumbapoa/backend/app/model"
	"github.com/nyumbapoa/backend/app/storage"
	"log"
)

func Migrate(database *storage.Database) {

	err := database.DB.AutoMigrate(
		model.User{},
		model.Role{},
		model.Permission{},
		model.Bill{},
		model.Building{},
		model.BuildingImage{},
		model.County{},
		model.Feedback{},
		model.House{},
		model.HouseImage{},
		model.Notice{},
		model.PromoCode{},
		model.Transaction{},
		model.UtilityProvider{},
		model.UtilityType{},
		model.Land{},
	)

	if err != nil {
		log.Println(err)
	}
}
