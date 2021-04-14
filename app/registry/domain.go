package registry

import (
	"github.com/nyumbapoa/backend/app/auth"
	"github.com/nyumbapoa/backend/app/bill"
	"github.com/nyumbapoa/backend/app/building"
	"github.com/nyumbapoa/backend/app/county"
	"github.com/nyumbapoa/backend/app/feedback"
)

type Domain struct {
	Auth     auth.Interactor
	Bill     bill.Interactor
	Building building.Interactor
	County   county.Interactor
	Feedback feedback.Interactor
}
