// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/sharings": {
            "post": {
                "description": "Inserts a new sharing",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "insert-sharing",
                "parameters": [
                    {
                        "description": "Shared content",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateSharing_Params"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Sharing"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/sharings/{hash}": {
            "get": {
                "description": "Get sharing by hash",
                "produces": [
                    "application/json"
                ],
                "operationId": "get-sharing-by-hash",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Sharing hash",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Sharing"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 409
                },
                "message": {
                    "type": "string",
                    "example": "Some Error"
                }
            }
        },
        "dto.CreateSharing_Params": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": "aaaabbbbcccc"
                }
            }
        },
        "dto.Sharing": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": "aaaabbbbcccc"
                },
                "hash": {
                    "type": "string",
                    "example": "11c85195ae99540ac07f80e2905e6e39aaefc4ac94cd380f366e79ba83560566"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Visualtez Storage API",
	Description:      "API documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}