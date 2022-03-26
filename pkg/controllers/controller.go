package controllers

import (
	"contact-service/pkg/models"
	"contact-service/pkg/services"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
)

// GetRoutes func show all routes.
// @Description  Retrieve all routes in this service.
// @Summary      Retrieve all routes
// @Tags         Routes Public
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Route
// @Router       / [get]
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

// GetContactById func retrieves a contact by given id.
// @Description  Retrieve a contact's data.
// @Summary      Retrieve a contact
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Contact Id"
// @Success      200  {object}  models.Contact
// @Failure      400  {object}  models.DefaultError
// @Failure      404  {object}  models.DefaultError
// @Router       /contacts/:id [get]
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

// GetContactsByQuery func retrieve contacts by given query parameters.
// @Description  Search contacts.
// @Summary      Search contacts.
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Param        organization_id  query     string  true  "Organization id"
// @Success      200              {object}  []models.Contact
// @Failure      400              {object}  models.DefaultError
// @Failure      500              {object}  models.DefaultError
// @Router       /contacts/search [get]
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

// CreateContact func Creates a contact.
// @Description  Creates a contact.
// @Summary      Creates a contact.
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.Contact
// @Failure      400  {object}  models.DefaultError
// @Failure      422  {object}  models.DefaultError
// @Failure      500  {object}  models.DefaultError
// @Router       /contacts [post]
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

// UpdateContact func Update a contact.
// @Description  Update a contact.
// @Summary      Update a contact.
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Contact Id"
// @Success      201  {object}  models.Contact
// @Failure      400  {object}  models.DefaultError
// @Failure      404  {object}  models.DefaultError
// @Failure      422  {object}  models.DefaultError
// @Router       /contacts/:id [put]
func UpdateContact(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(models.CreateContact_Request)

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(services.NewError(fiber.StatusBadRequest, "Requires parameter id not found"))
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

// DeleteContact func Delete a contact.
// @Description  Delete a contact.
// @Summary      Delete a contact.
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Contact Id"
// @Success      201  {object}  models.Contact
// @Failure      400  {object}  models.DefaultError
// @Failure      404  {object}  models.DefaultError
// @Failure      500  {object}  models.DefaultError
// @Router       /contacts/:id [delete]
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
