package main

import (
	mid "auth-service/pkg/middleware"
	"auth-service/pkg/routes"
	"auth-service/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	utils.GoDotEnvVariable("PORT")
	//database.Init()
	mid.FiberMiddleware(app)
	routes.PublicRoutes(app)
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
