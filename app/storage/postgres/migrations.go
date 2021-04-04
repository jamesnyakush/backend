package postgres

import (
	"github.com/nyumbapoa/backend/app/storage"
)

func Migrate(database *storage.Database) {

	/*
		err := database.DB.AutoMigrate(
				model.User{},
				model.Role{},
				model.Permission{},
				model.Bill{},
				model.Building{},
				model.County{},
				model.Feedback{},
				model.House{},
				model.HouseImage{},
				model.HouseType{},
				model.Notice{},
				model.PaymentMethod{},
				model.PromoCode{},
				model.Transaction{},
				model.UtilityProvider{},
				model.UtilityType{},
			)

			if err != nil {
				log.Println(err)
			}
	*/
}
