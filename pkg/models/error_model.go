package models

type DefaultError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
