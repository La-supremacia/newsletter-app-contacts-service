package main

import (
	mid "contact-service/pkg/middleware"
	"contact-service/pkg/routes"
	"contact-service/pkg/utils"
	"contact-service/platform/database"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	utils.GoDotEnvVariable("PORT")
	database.Init()
	mid.FiberMiddleware(app)
	routes.PublicRoutes(app)
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
