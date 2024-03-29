package controllers

import (
	"contact-service/pkg/models"
	"contact-service/pkg/services"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRoutes(c *fiber.Ctx) error {
	rootRoute := c.App().GetRoute("Root")
	getByIdRoute := c.App().GetRoute("GetContact")
	searchRoute := c.App().GetRoute("SearchContacts")
	createRoute := c.App().GetRoute("CreateContact")
	updateRoute := c.App().GetRoute("UpdateContact")
	deleteRoute := c.App().GetRoute("DeleteContact")

	var ruts = [6]models.Route{
		services.NewRoute(rootRoute.Path, rootRoute.Method, rootRoute.Name, ""),
		services.NewRoute(getByIdRoute.Path, getByIdRoute.Method, getByIdRoute.Name, ""),
		services.NewRoute(searchRoute.Path, searchRoute.Method, searchRoute.Name, "organization_id:string"),
		services.NewRoute(createRoute.Path, createRoute.Method, createRoute.Name, "name:string,last_name:string,email:string,phone:string,organization_id:string"),
		services.NewRoute(updateRoute.Path, updateRoute.Method, updateRoute.Name, "name:string,last_name:string,email:string,phone:string"),
		services.NewRoute(deleteRoute.Path, deleteRoute.Method, deleteRoute.Name, ""),
	}
	return c.Status(fiber.StatusOK).JSON(ruts)
}

func GetContactById(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "required param id not found")
	}

	baseModel := &models.Contact{}
	coll := mgm.Coll(baseModel)
	err := coll.FindByID(id, baseModel)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	return c.JSON(baseModel)
}

func GetContactsByQuery(c *fiber.Ctx) error {
	organizationId := c.Query("organization_id")

	if organizationId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "required query param organization_id not found")
	}

	baseModel := &models.Contact{}
	coll := mgm.Coll(baseModel)
	result := []models.Contact{}

	query := bson.M{"organization_id": bson.M{operator.Eq: organizationId}}
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
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(createdContact)
}

func UpdateContact(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(models.CreateContact_Request)

	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "required param id not found")
	}
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	baseModel := &models.Contact{}
	coll := mgm.Coll(baseModel)
	err := coll.FindByID(id, baseModel)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	baseModel.Name = body.Name
	baseModel.LastName = body.LastName
	baseModel.Email = body.Email
	baseModel.Phone = body.Phone
	err = coll.Update(baseModel)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(baseModel)
}

func DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "required param id not found")
	}

	baseModel := &models.Contact{}
	coll := mgm.Coll(baseModel)
	err := coll.FindByID(id, baseModel)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	err = coll.Delete(baseModel)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(fiber.Map{
		"sucess":  true,
		"message": "The contact was successfully deleted",
	})
}
