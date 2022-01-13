package controllers

import (
	"auth-service/pkg/models"
	"auth-service/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func GetRoutes(c *fiber.Ctx) error {
	rootRoute := c.App().GetRoute("Root")

	var ruts = [1]models.Route{
		services.NewRoute(rootRoute.Path, rootRoute.Method, rootRoute.Name, ""),
	}
	return c.Status(fiber.StatusOK).JSON(ruts)
}
