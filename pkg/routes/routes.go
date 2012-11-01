package routes

import (
	"contact-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/", controllers.GetRoutes).Name("Root")
	route.Post("/", controllers.CreateContact).Name("CreateContact")
}
