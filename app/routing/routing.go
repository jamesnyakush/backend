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

		// create group at /api/admin
		admin := api.Group("/admin", middleware.AuthByBearerToken(configs.Secret))
		admin.Post("/assign-float", user_handlers.AssignFloat(domain.Admin))
		admin.Post("/update-charge", user_handlers.UpdateCharge(domain.Tariff))
		admin.Get("/get-tariff", user_handlers.GetTariff(domain.Tariff))
		admin.Put("/super-agent-status", user_handlers.UpdateSuperAgentStatus(domain.Agent))
	*/
}
