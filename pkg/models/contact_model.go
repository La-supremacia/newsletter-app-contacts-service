package models

import "github.com/kamva/mgm/v3"

type Contact struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `db:"name" bson:"name" json:"name" validate:"required,lte=255"`
	LastName         string `db:"last_name" bson:"last_name" json:"last_name,omitempty" validate:"lte=255"`
	Email            string `db:"email" bson:"email" json:"email" validate:"required,email,lte=255"`
	Phone            string `db:"phone" bson:"phone" json:"phone" validate:"required,lte=255"`
	OrganizationId   string `db:"organization_id" bson:"organization_id" json:"organization_id" validate:"required"`
}

type CreateContact_Request struct {
	Name           string `json:"name" validate:"required,lte=255"`
	LastName       string `json:"last_name,omitempty" validate:"lte=255"`
	Email          string `json:"email" validate:"required,email,lte=255"`
	Phone          string `json:"phone" validate:"required,lte=255"`
	OrganizationId string `json:"organization_id" validate:"required"`
}
