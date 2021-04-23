package building_handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nyumbapoa/backend/app/building"
)

func AddBuilding(i building.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx, "Building Added Successfully")
		return nil
	}
}

func VerifyBuilding(i building.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx, "Building Verified Successfully")
		return nil
	}
}

func FetchAllBuildings(i building.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx, "Buildings")
		return nil
	}
}

func FetchUserBuildings(i building.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx, "Buildings")
		return nil
	}
}
func FetchBuildingHouses(i building.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx, "Buildings")
		return nil
	}
}

func UpdateBuilding(i building.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx, "Building updated Successfully")
		return nil
	}
}

func DeleteBuilding(i building.Interactor) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Fprintf(ctx, "Building Deleted Successfully")
		return nil
	}
}
