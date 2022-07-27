// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/products": {
            "get": {
                "description": "Get products csv",
                "produces": [
                    "application/csv"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get Products CSV",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "Create products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create Product",
                "parameters": [
                    {
                        "description": "create product",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.CreateProductInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/core.Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/appError.Element"
                            }
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "put": {
                "description": "Update products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Update Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update product",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/core.UpdateProductInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/appError.Element"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/restHandler.response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Delete Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/restHandler.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "appError.Element": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "core.CreateProductInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "core.Product": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "core.UpdateProductInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "restHandler.response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "Product app REST-API",
	Description:      "Application for adding/getting products and download CSV-file",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}