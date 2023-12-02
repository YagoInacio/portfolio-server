// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Yago Faran",
            "email": "yagofaran@gmail.com"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/images/icons/{name}": {
            "get": {
                "description": "This endpoint returns the icon file by its name",
                "produces": [
                    "image/png",
                    "image/jpeg",
                    "image/svg+xml",
                    "image/gif"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "Get icon file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/images/{name}": {
            "get": {
                "description": "This endpoint returns the image file by its name",
                "produces": [
                    "image/png",
                    "image/jpeg",
                    "image/svg+xml",
                    "image/gif"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "Get image file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/technologies": {
            "get": {
                "description": "This endpoint lists all technologies that are marked as display: true",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Technologies"
                ],
                "summary": "List Technologies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/usecases.ListTechnologiesOutput"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "This endpoint creates a new technology",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Technologies"
                ],
                "summary": "Creates a new technology",
                "parameters": [
                    {
                        "description": "creation attributes",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecases.TechnologyInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecases.TechnologyOutput"
                        }
                    }
                }
            }
        },
        "/technologies/{id}": {
            "patch": {
                "description": "This endpoint alters a technology's display parameter",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Technologies"
                ],
                "summary": "Alters technology display",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Technology Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "true",
                            "false"
                        ],
                        "type": "string",
                        "description": "display value to be modified to",
                        "name": "display",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecases.AlterTechnologyDisplayOutput"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "usecases.AlterTechnologyDisplayOutput": {
            "type": "object",
            "properties": {
                "display": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "src": {
                    "type": "string"
                }
            }
        },
        "usecases.ListTechnologiesOutput": {
            "type": "object",
            "properties": {
                "display": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "src": {
                    "type": "string"
                }
            }
        },
        "usecases.TechnologyInput": {
            "type": "object",
            "properties": {
                "display": {
                    "description": "whether it will be displayed on the skills page",
                    "type": "boolean",
                    "example": true
                },
                "name": {
                    "description": "tech name",
                    "type": "string",
                    "example": "Go"
                },
                "src": {
                    "description": "tech icon file name",
                    "type": "string",
                    "example": "go.png"
                }
            }
        },
        "usecases.TechnologyOutput": {
            "type": "object",
            "properties": {
                "display": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "src": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Portfolio server",
	Description:      "This API is designed to improve the management of project and experience information for your portfolio website, ensuring a seamless rendering and effortless updating process. Additionally, it facilitates interaction with Firebase Storage, allowing you to efficiently load and display icons and images, enhancing the overall visual experience of your portfolio.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
