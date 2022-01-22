package routes

import (
	"contact-service/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/", controllers.GetRoutes).Name("Root")
	route.Get("/contacts/search", controllers.GetContactsByQuery).Name("SearchConstacts")
	route.Get("/contacts/:id", controllers.GetContactById).Name("GetContact")
	route.Post("/contacts", controllers.CreateContact).Name("CreateContact")
	route.Put("/contacts/:id", controllers.UpdateContact).Name("UpdateContact")
	route.Delete("/contacts/:id", controllers.DeleteContact).Name("DeleteContact")
}
