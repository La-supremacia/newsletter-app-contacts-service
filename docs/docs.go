// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "lasupremaciadelpuntoycoma@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Retrieve all routes in this service.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Routes Public"
                ],
                "summary": "Retrieve all routes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Route"
                            }
                        }
                    }
                }
            }
        },
        "/contacts": {
            "post": {
                "description": "Creates a contact.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contacts"
                ],
                "summary": "Creates a contact.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Contact info",
                        "name": "contact",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateContact_Request"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Contact"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    }
                }
            }
        },
        "/contacts/search": {
            "get": {
                "description": "Search contacts.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contacts"
                ],
                "summary": "Search contacts.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Organization id",
                        "name": "organization_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Contact"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    }
                }
            }
        },
        "/contacts/{id}": {
            "get": {
                "description": "Retrieve a contact's data.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contacts"
                ],
                "summary": "Retrieve a contact",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contact Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Contact"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a contact.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contacts"
                ],
                "summary": "Update a contact.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contact Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Contact info",
                        "name": "contact",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.CreateContact_Request"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Contact"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a contact.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contacts"
                ],
                "summary": "Delete a contact.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contact Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Contact"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.DefaultError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Contact": {
            "type": "object",
            "required": [
                "email",
                "name",
                "organization_id",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 255
                },
                "name": {
                    "type": "string",
                    "maxLength": 255
                },
                "organization_id": {
                    "type": "string"
                },
                "phone": {
                    "type": "string",
                    "maxLength": 255
                }
            }
        },
        "models.CreateContact_Request": {
            "type": "object",
            "required": [
                "email",
                "name",
                "organization_id",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 255
                },
                "name": {
                    "type": "string",
                    "maxLength": 255
                },
                "organization_id": {
                    "type": "string"
                },
                "phone": {
                    "type": "string",
                    "maxLength": 255
                }
            }
        },
        "models.DefaultError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.Route": {
            "type": "object",
            "properties": {
                "method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "params": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "newsletter-app-contact-service.herokuapp.com",
	BasePath:         "/api/v1",
	Schemes:          []string{"https"},
	Title:            "Contacts microservice for newsletter-app",
	Description:      "This service manage the contacts of the app.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
