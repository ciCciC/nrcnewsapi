// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "williamhall2894@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/ciCciC/nrcnewsapi"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "get": {
                "description": "Get category list",
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves category list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Category"
                            }
                        }
                    }
                }
            }
        },
		"/categories/nl": {
            "get": {
                "description": "Get category list in Dutch",
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves category list in Dutch",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Category"
                            }
                        }
                    }
                }
            }
        },
		"/category/all/:name": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ArticleItem"
                            }
                        }
                    }
                }
            }
        },
        "/category/astronomy": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ArticleItem"
                            }
                        }
                    }
                }
            }
        },
        "/category/astronomy/article": {
            "post": {
                "description": "Get a article with article item payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves article",
                "parameters": [
                    {
                        "description": "Get article",
                        "name": "articleitem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ArticleItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    }
                }
            }
        },
        "/category/games": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ArticleItem"
                            }
                        }
                    }
                }
            }
        },
        "/category/games/article": {
            "post": {
                "description": "Get a article with article item payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves article",
                "parameters": [
                    {
                        "description": "Get article",
                        "name": "articleitem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ArticleItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    }
                }
            }
        },
        "/category/news": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ArticleItem"
                            }
                        }
                    }
                }
            }
        },
        "/category/news/article": {
            "post": {
                "description": "Get a article with article item payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves article",
                "parameters": [
                    {
                        "description": "Get article",
                        "name": "articleitem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ArticleItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    }
                }
            }
        },
        "/category/physics": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ArticleItem"
                            }
                        }
                    }
                }
            }
        },
        "/category/physics/article": {
            "post": {
                "description": "Get a article with article item payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves article",
                "parameters": [
                    {
                        "description": "Get article",
                        "name": "articleitem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ArticleItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    }
                }
            }
        },
        "/category/technology": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves all articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ArticleItem"
                            }
                        }
                    }
                }
            }
        },
        "/category/technology/article": {
            "post": {
                "description": "Get a article with article item payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves article",
                "parameters": [
                    {
                        "description": "Get article",
                        "name": "articleitem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ArticleItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Article": {
            "type": "object",
            "properties": {
                "imageLink": {
                    "type": "string"
                },
                "pageLink": {
                    "type": "string"
                },
                "sectionList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Section"
                    }
                },
                "teaser": {
                    "type": "string",
                    "example": "BioTelemetry maakt systemen om hartpatiënten te diagnosticeren en op afstand in de gaten te houden. Philips zag de vraag naar zorg op afstand toenemen vanwege de pandemie."
                },
                "title": {
                    "type": "string",
                    "example": "Philips koopt Amerikaanse BioTelemetry voor 2,3 miljard euro"
                },
                "topic": {
                    "type": "string",
                    "example": "E-health"
                }
            }
        },
        "model.ArticleItem": {
            "type": "object",
            "properties": {
                "imageLink": {
                    "type": "string"
                },
                "pageLink": {
                    "type": "string"
                },
                "teaser": {
                    "type": "string",
                    "example": "BioTelemetry maakt systemen om hartpatiënten te diagnosticeren en op afstand in de gaten te houden. Philips zag de vraag naar zorg op afstand toenemen vanwege de pandemie."
                },
                "title": {
                    "type": "string",
                    "example": "Philips koopt Amerikaanse BioTelemetry voor 2,3 miljard euro"
                },
                "topic": {
                    "type": "string",
                    "example": "E-health"
                }
            }
        },
        "model.Category": {
            "type": "object",
            "properties": {
                "display": {
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                }
            }
        },
        "model.ContentBody": {
            "type": "object",
            "properties": {
                "cType": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                }
            }
        },
        "model.Section": {
            "type": "object",
            "properties": {
                "contentBody": {
                    "description": "Contents []string",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ContentBody"
                    }
                },
                "title": {
                    "type": "string"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
	swag.Register(swag.Name, &s{})
}
