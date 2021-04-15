package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nyumbapoa/backend/app"
	"github.com/nyumbapoa/backend/app/registry"
	"github.com/nyumbapoa/backend/app/routing/error_handlers"
)

func Router(domain *registry.Domain, configs app.Config) *fiber.App {

	srv := fiber.New(
		fiber.Config{ErrorHandler: error_handlers.ErrorHandler},
	)

	apiGroup := srv.Group("/api")
	apiGroup.Use(logger.New())

	apiRouteGroup(apiGroup, domain, configs)

	return srv
}

//
func apiRouteGroup(api fiber.Router, domain *registry.Domain, config app.Config) {
	/*
		api.Post("/login/:user_type", user_handlers.Authenticate(domain, configs))
		api.Post("/user/:user_type", user_handlers.Register(domain))
	*/

	// Auth (login,register,forgot-password, verify, Fetch All) endpoint
	auth := api.Group("/auth")
	auth.Post("/login")
	auth.Post("/register")
	auth.Post("/forgot-password")
	auth.Put("/change-password")

	// Building (Add,Verify, Update, Delete, Fetch All, Verify)
	building := api.Group("/buildings")
	building.Post("/")
	building.Get("/")
	building.Put("/:id")
	building.Put("/:id/verify")
	building.Delete("/:id")

	// House (Add,Verify, Update, Delete ,Fetch All, Verify)
	houses := api.Group("/houses")
	houses.Post("/")
	houses.Get("/")
	houses.Put("/:id")
	houses.Put("/:id/verify")
	houses.Delete("/:id")

	//

}
