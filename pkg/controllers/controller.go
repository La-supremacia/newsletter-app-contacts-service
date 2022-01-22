package controllers

import (
	"contact-service/pkg/models"
	"contact-service/pkg/services"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
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

func GetContactById(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(errors.New("required param id not found"))
	}

	baseModel := &models.Contact{}
	coll := mgm.Coll(baseModel)
	err := coll.FindByID(id, baseModel)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(baseModel)
}

func GetContactsByQuery(c *fiber.Ctx) error {
	organizationId := c.Query("organization_id")

	baseModel := &models.Contact{}
	coll := mgm.Coll(baseModel)
	result := []models.Contact{}

	query := bson.M{}
	if organizationId != "" {
		query = bson.M{"organization_id": bson.M{operator.Eq: organizationId}}
	}

	err := coll.SimpleFind(&result, query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(result)
}

func CreateContact(c *fiber.Ctx) error {
	body := new(models.CreateContact_Request)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	createdContact := services.NewContact(body)

	err := mgm.Coll(createdContact).Create(createdContact)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	fmt.Println("Successfully created a new Contact", createdContact.Name, createdContact.LastName)

	return c.Status(fiber.StatusOK).JSON(createdContact)
}

func UpdateContact(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(models.CreateContact_Request)

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(errors.New("required param id not found"))
	}
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	baseModel := &models.Contact{}
	coll := mgm.Coll(baseModel)
	err := coll.FindByID(id, baseModel)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	baseModel.Name = body.Name
	baseModel.LastName = body.LastName
	baseModel.Email = body.Email
	baseModel.Phone = body.Phone
	err = coll.Update(baseModel)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(baseModel)
}

func DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(errors.New("required param id not found"))
	}

	baseModel := &models.Contact{}
	coll := mgm.Coll(baseModel)
	_ = coll.FindByID(id, baseModel)
	err := coll.Delete(baseModel)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(fiber.Map{
		"sucess":  true,
		"message": "The organization was successfully deleted",
	})
}
