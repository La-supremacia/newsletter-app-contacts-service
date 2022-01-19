package controllers

import (
	"contact-service/pkg/models"
	"contact-service/pkg/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)

func GetRoutes(c *fiber.Ctx) error {
	rootRoute := c.App().GetRoute("Root")
	createRoute := c.App().GetRoute("CreateContact")

	var ruts = [2]models.Route{
		services.NewRoute(rootRoute.Path, rootRoute.Method, rootRoute.Name, ""),
		services.NewRoute(createRoute.Path, createRoute.Method, createRoute.Name, "name:string,last_name:string,email:string,organization_id:string"),
	}
	return c.Status(fiber.StatusOK).JSON(ruts)
}

func CreateContact(c *fiber.Ctx) error {
	body := new(models.CreateContact_Request)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	createdContact := services.NewContact(body)

	err := mgm.Coll(createdContact).Create(createdContact)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Successfully created a new Contact", createdContact.Name, createdContact.LastName)

	return c.Status(fiber.StatusOK).JSON(createdContact)
}
