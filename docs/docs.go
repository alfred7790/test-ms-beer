// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/v1/beers": {
            "get": {
                "description": "a page and a limit of results per page is optional.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beers"
                ],
                "summary": "returns a list of paginated beers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Request page, default 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "number of results per page, default 50",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.BeerList"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Used to create a new beer",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beers"
                ],
                "summary": "returns details about a new beer was created",
                "parameters": [
                    {
                        "description": "struct to create a new beer",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.BeerDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Beer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    }
                }
            }
        },
        "/v1/beers/{beerid}": {
            "get": {
                "description": "Used to find a beer",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beers"
                ],
                "summary": "returns a beer info searching by beerId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Beer ID",
                        "name": "beerid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Beer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Used to update the beer's info",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beers"
                ],
                "summary": "returns details about a beer was updated",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Beer ID",
                        "name": "beerid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "struct to update a beer",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.BeerDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Beer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    }
                }
            }
        },
        "/v1/beers/{beerid}/boxprice": {
            "get": {
                "description": "List the price of a box of beer",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beers"
                ],
                "summary": "returns a beer info searching by beerId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Beer ID",
                        "name": "beerid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Currency, default USD",
                        "name": "currency",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Quantity, default 6",
                        "name": "quantity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.BeerBox"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utilities.FailureResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Beer": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Bohemia"
                },
                "sku": {
                    "type": "string",
                    "example": "BHE001"
                },
                "unitPrice": {
                    "type": "number",
                    "example": 1.2
                }
            }
        },
        "entity.BeerBox": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Bohemia"
                },
                "prices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.BeerBoxPrice"
                    }
                },
                "quantity": {
                    "type": "integer",
                    "example": 10
                },
                "sku": {
                    "type": "string",
                    "example": "BHE001"
                },
                "unitPrice": {
                    "type": "number",
                    "example": 1.2
                }
            }
        },
        "entity.BeerBoxPrice": {
            "type": "object",
            "properties": {
                "currencyName": {
                    "type": "string",
                    "example": "USD"
                },
                "currencyPrice": {
                    "type": "number",
                    "example": 1.2
                },
                "total": {
                    "type": "number",
                    "example": 12
                }
            }
        },
        "entity.BeerDTO": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Bohemia"
                },
                "sku": {
                    "type": "string",
                    "example": "BHE001"
                },
                "unitPrice": {
                    "type": "number",
                    "example": 1.2
                }
            }
        },
        "entity.BeerList": {
            "type": "object",
            "properties": {
                "beers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Beer"
                    }
                },
                "total": {
                    "type": "integer",
                    "example": 100
                }
            }
        },
        "utilities.FailureResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string",
                    "example": "Error message for developers"
                },
                "message": {
                    "type": "string",
                    "example": "Error message for users"
                },
                "status": {
                    "type": "string",
                    "example": "failure"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Test API",
	Description: "test api",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}