package main

import (
	mid "contact-service/pkg/middleware"
	"contact-service/pkg/routes"
	"contact-service/pkg/utils"
	"contact-service/platform/database"
	"os"

	_ "contact-service/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title           Contacts microservice for newsletter-app
// @version         1.0
// @description     This service manage the contacts of the app.

// @contact.name   API Support
// @contact.email  lasupremaciadelpuntoycoma@gmail.com

// @host      https://newsletter-app-contact-service.herokuapp.com/
// @BasePath  /api/v1

// @securitydefinitions.apikey
// @in header
// @name Authorization
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	utils.GoDotEnvVariable("PORT")
	database.Init()
	mid.FiberMiddleware(app)
	routes.PublicRoutes(app)
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
