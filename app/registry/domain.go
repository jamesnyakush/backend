package registry

import (
	"github.com/nyumbapoa/backend/app"
	"github.com/nyumbapoa/backend/app/auth"
	"github.com/nyumbapoa/backend/app/bill"
	"github.com/nyumbapoa/backend/app/building"
	"github.com/nyumbapoa/backend/app/county"
	"github.com/nyumbapoa/backend/app/feedback"
	"github.com/nyumbapoa/backend/app/house"
	"github.com/nyumbapoa/backend/app/permission"
	"github.com/nyumbapoa/backend/app/promocode"
	"github.com/nyumbapoa/backend/app/role"
	"github.com/nyumbapoa/backend/app/storage"
)

type Domain struct {
	Auth       auth.Interactor
	Bill       bill.Interactor
	Building   building.Interactor
	County     county.Interactor
	Feedback   feedback.Interactor
	House      house.Interactor
	Permission permission.Interactor
	Promocode  promocode.Interactor
	Role       role.Interactor
}

//channels *Channels
func NewDomain(config app.Config, database *storage.Database) *Domain {

	authRepo := auth.NewRepository(database)
	billRepo := bill.NewRepository(database)
	buildingRepo := building.NewRepository(database)
	//countyRepo := county
	//feedbackRepo := feedback.NewRepository(database)
	//houseRepo := house.NewRepository(database)
	//permissionRepo := permission.NewRepository(database)
	//promocodeRepo := promocode.NewRepository(database)
	//roleRepo := role.NewRepository(database)

	return &Domain{
		Auth:       auth.NewInteractor(config, authRepo),
		Bill:       bill.NewInteractor(config, billRepo),
		Building:   building.NewInteractor(config, buildingRepo),
		County:     county.NewInteractor(),
		Feedback:   feedback.NewInteractor(),
		House:      house.NewInteractor(),
		Permission: permission.NewInteractor(),
		Promocode:  promocode.NewInteractor(),
		Role:       role.NewInteractor(),
	}
}
