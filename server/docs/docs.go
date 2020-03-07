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
        "/transcode/{id}": {
            "get": {
                "description": "Get transcode status by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transcode"
                ],
                "summary": "Get transcode status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Transcoder"
                        }
                    },
                    "400": {
                        "description": "Failure to parse the id",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Failure to find the id",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
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
                            "$ref": "#/definitions/server.UploadAudioResp"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
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
                            "$ref": "#/definitions/server.UploadImageResp"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Transcoder": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "completed_at": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "list": {
                    "type": "string"
                },
                "original": {
                    "type": "string"
                },
                "percentage": {
                    "type": "integer"
                },
                "upload_id": {
                    "type": "string"
                }
            }
        },
        "server.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "server.UploadAudioResp": {
            "type": "object",
            "properties": {
                "duration": {
                    "type": "number"
                },
                "file_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "transcoder_id": {
                    "type": "string"
                }
            }
        },
        "server.UploadImageResp": {
            "type": "object",
            "properties": {
                "filename": {
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
	Host:        "localhost:8081",
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
