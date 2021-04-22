package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nyumbapoa/backend/app"
	"github.com/nyumbapoa/backend/app/registry"
	"github.com/nyumbapoa/backend/app/routing/auth_handlers"
	"github.com/nyumbapoa/backend/app/routing/building_handlers"
	"github.com/nyumbapoa/backend/app/routing/caretaker_handlers"
	"github.com/nyumbapoa/backend/app/routing/error_handlers"
	"github.com/nyumbapoa/backend/app/routing/house_handlers"
	"github.com/nyumbapoa/backend/app/routing/land_handlers"
)

func Router(domain *registry.Domain, configs app.Config) *fiber.App {
	srv := fiber.New(fiber.Config{ErrorHandler: error_handlers.ErrorHandler})

	apiGroup := srv.Group("/api")
	apiGroup.Use(logger.New())

	apiRouteGroup(apiGroup, domain, configs)

	return srv
}

//
func apiRouteGroup(api fiber.Router, domain *registry.Domain, config app.Config) {

	// Auth (login,register,forgot-password, verify, Fetch All) endpoint
	api.Post("/login", auth_handlers.LoginUser(domain.Auth, config))
	api.Post("/register", auth_handlers.RegisterUser(domain.Auth))
	api.Post("/forgot-password", auth_handlers.ForgotPassword(domain.Auth))
	api.Put("/change-password", auth_handlers.ChangePassword(domain.Auth))

	// Building (Add,Verify, Update, Delete, Fetch All, Verify)
	building := api.Group("/buildings")
	building.Post("/", building_handlers.AddBuilding(domain.Building))
	building.Get("/", building_handlers.FetchAllBuildings(domain.Building))
	building.Get("/:user_id", building_handlers.FetchUserBuildings(domain.Building))
	building.Get("/:id/houses", building_handlers.FetchBuildingHouses(domain.Building))
	building.Put("/:id", building_handlers.UpdateBuilding(domain.Building))
	building.Put("/:id/verify", building_handlers.VerifyBuilding(domain.Building))
	building.Delete("/:id", building_handlers.DeleteBuilding(domain.Building))

	// House (Add,Verify, Update, Delete ,Fetch All, Verify)
	houses := api.Group("/houses")
	houses.Post("/", house_handlers.AddHouse(domain.House))
	//houses.Get("/")
	houses.Put("/:id", house_handlers.UpdateHouse(domain.House))
	houses.Put("/:id/verify", house_handlers.VerifyHouse(domain.House))
	houses.Delete("/:id", house_handlers.DeleteHouse(domain.House))

	// Land (Add,Verify, Update, Delete ,Fetch All, Verify)
	land := api.Group("/lands")
	land.Post("/", land_handlers.AddLand())
	//land.Get("/")
	land.Put("/:id/verify", land_handlers.VerifyLand())
	land.Put("/:id", land_handlers.UpdateLand())
	land.Delete("/:id", land_handlers.DeleteLand())

	// Caretaker (Add,Verify, Update, Delete ,Fetch All, Verify)
	caretaker := api.Group("/caretakers")
	caretaker.Post("/", caretaker_handlers.CreateCaretaker())
	caretaker.Put("/:id/assign/:id", caretaker_handlers.AssignCaretaker())

}
