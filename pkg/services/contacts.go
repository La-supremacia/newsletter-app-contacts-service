package services

import "contact-service/pkg/models"

func NewContact(request *models.CreateContact_Request) *models.Contact {
	return &models.Contact{
		Name:           request.Name,
		LastName:       request.LastName,
		Email:          request.Email,
		Phone:          request.Phone,
		OrganizationId: request.OrganizationId,
	}
}
