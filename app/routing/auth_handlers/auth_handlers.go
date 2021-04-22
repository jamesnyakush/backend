package auth_handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nyumbapoa/backend/app"
	"github.com/nyumbapoa/backend/app/auth"
)

func LoginUser(i auth.Interactor, config app.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx, "<a>Welcome To Login</a>")
		return nil
	}
}

func RegisterUser(i auth.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx, "Welcome to Register")
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
