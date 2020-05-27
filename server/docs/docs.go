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
        "contact": {
            "name": "BitSong",
            "email": "hello@bitsong.io"
        },
        "license": {
            "name": "CC0",
            "url": "https://creativecommons.org/share-your-work/public-domain/cc0/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/upload/audio": {
            "post": {
                "description": "Upload, transcode and publish to ipfs an audio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "upload"
                ],
                "summary": "Upload and transcode audio file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Audio file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.UploadCidResp"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorJson"
                        }
                    }
                }
            }
        },
        "/upload/image": {
            "post": {
                "description": "Upload, create and publish to ipfs an image",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "upload"
                ],
                "summary": "Upload and create image file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.UploadCidResp"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorJson"
                        }
                    }
                }
            }
        },
        "/upload/manifest": {
            "post": {
                "description": "Upload, create and publish to ipfs a manifest data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "upload"
                ],
                "summary": "Upload and create raw data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Manifest",
                        "name": "manifest",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.UploadCidResp"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorJson"
                        }
                    }
                }
            }
        },
        "/upload/{cid}/status": {
            "get": {
                "description": "Get upload status by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "upload"
                ],
                "summary": "Get upload status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CID",
                        "name": "cid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.UploadStatusResp"
                        }
                    },
                    "400": {
                        "description": "Failure to parse the id",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorJson"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorJson"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "server.ErrorJson": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "$ref": "#/definitions/server.ErrorJsonBody"
                }
            }
        },
        "server.ErrorJsonBody": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "server.UploadCidResp": {
            "type": "object",
            "properties": {
                "cid": {
                    "type": "string"
                },
                "filename": {
                    "type": "string"
                }
            }
        },
        "server.UploadStatusResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "percentage": {
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
	Version:     "0.1",
	Host:        "localhost:1347",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "bitsongms API Docs",
	Description: "Swagger documentation for the BitSong Media Server service API.",
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
