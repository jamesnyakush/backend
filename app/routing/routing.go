package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nyumbapoa/backend/app"
)

/*func Router(domain *registry.Domain, config app.Config) *fiber.App {

	srv := fiber.New(
		fiber.Config{ErrorHandler: error_handlers.ErrorHandler},
	)

	apiGroup := srv.Group("/api")
	apiGroup.Use(logger.New())

	//apiRouteGroup(apiGroup, domain, config)

	return srv
}
*/
 //domain *registry.Domain,
func apiRouteGroup(api fiber.Router, config app.Config) {
	/*
		api.Post("/login/:user_type", user_handlers.Authenticate(domain, config))
		api.Post("/user/:user_type", user_handlers.Register(domain))

		// create group at /api/admin
		admin := api.Group("/admin", middleware.AuthByBearerToken(config.Secret))
		admin.Post("/assign-float", user_handlers.AssignFloat(domain.Admin))
		admin.Post("/update-charge", user_handlers.UpdateCharge(domain.Tariff))
		admin.Get("/get-tariff", user_handlers.GetTariff(domain.Tariff))
		admin.Put("/super-agent-status", user_handlers.UpdateSuperAgentStatus(domain.Agent))*/
}
