package auth_handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nyumbapoa/backend/app"
	"github.com/nyumbapoa/backend/app/auth"
	"github.com/nyumbapoa/backend/app/routing/responses"
	"net/http"
)

func LoginUser(i auth.Interactor, config app.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx, "<a>Welcome To Login</a>")
		return nil
	}
}

func RegisterUser(i auth.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params auth.RegisterParams
		_ = ctx.BodyParser(&params)

		err := params.Validate()

		if err != nil {
			return err
		}

		// register admin
		adm, err := i.Register(params)
		if err != nil {
			return err
		}

		// we use a presenter to reformat the response of admin.
		_ = ctx.Status(http.StatusOK).JSON(responses.RegistrationResponse(adm.ID))

		return nil
	}
}

func ForgotPassword(i auth.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx,"Welcome to Forgot pass")
		return nil
	}
}

func ChangePassword(i auth.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx,"Welcome to Change Password")
		return nil
	}
}
