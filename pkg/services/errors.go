package services

import "contact-service/pkg/models"

func NewError(status int, message string) *models.DefaultError {
	return &models.DefaultError{
		Status:  status,
		Message: message,
	}
}
