package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/nyumbapoa/backend/api/v1/auth"
	"github.com/nyumbapoa/backend/api/v1/bill"
	"github.com/nyumbapoa/backend/api/v1/building"
	"github.com/nyumbapoa/backend/api/v1/county"
	"github.com/nyumbapoa/backend/api/v1/feedback"
	"github.com/nyumbapoa/backend/api/v1/house"
	"github.com/nyumbapoa/backend/api/v1/notice"
	"github.com/nyumbapoa/backend/api/v1/payment"
	"github.com/nyumbapoa/backend/api/v1/permission"
	"github.com/nyumbapoa/backend/api/v1/promocode"
	"github.com/nyumbapoa/backend/api/v1/role"
	"github.com/nyumbapoa/backend/api/v1/utility"
	"github.com/nyumbapoa/backend/internal/config"
	"github.com/nyumbapoa/backend/internal/store"
)

func New(enableCORS bool, cfg *config.Config, store store.Store) (*chi.Mux, error) {

	authResource := auth.NewResource(store, cfg)
	billResource := bill.NewResource(store, cfg)
	buildingResource := building.NewResource(store, cfg)
	countyResource := county.NewResource(store, cfg)
	feedbackResource := feedback.NewResource(store, cfg)
	houseResource := house.NewResource(store, cfg)
	noticeResource := notice.NewResource(store, cfg)
	paymentResource := payment.NewResource(store, cfg)
	permissionResource := permission.NewResource(store, cfg)
	promocodeResource := promocode.NewResource(store, cfg)
	roleResource := role.NewResource(store, cfg)
	utilityResource := utility.NewResource(store, cfg)

	r := chi.NewRouter()

	r.Mount("/v1/auth", authResource.Router())
	r.Mount("/v1/bill", billResource.Router())
	r.Mount("/v1/building", buildingResource.Router())
	r.Mount("/v1/county", countyResource.Router())
	r.Mount("/v1/feedback", feedbackResource.Router())
	r.Mount("/v1/house", houseResource.Router())
	r.Mount("/v1/notice", noticeResource.Router())
	r.Mount("/v1/payment", paymentResource.Router())
	r.Mount("/v1/permission", permissionResource.Router())
	r.Mount("/v1/promocode", promocodeResource.Router())
	r.Mount("/v1/role", roleResource.Router())
	r.Mount("/v1/utility", utilityResource.Router())

	return r, nil
}

func corsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Accept-Encoding", "Content-Type", "Content-Length", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   true,
		MaxAge:             86400,
		OptionsPassthrough: false,
	})
}
