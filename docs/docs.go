// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/cloneNamespace": {
            "post": {
                "description": "Clone a namespace and its objects to a new namespace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Clone a namespace",
                "parameters": [
                    {
                        "description": "Namespace clone request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.NSClonerRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces": {
            "get": {
                "description": "Get all namespaces in the cluster",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all namespaces",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/configmaps/display": {
            "get": {
                "description": "Display all config maps in the specified namespace",
                "produces": [
                    "application/json"
                ],
                "summary": "Display config maps for a specific namespace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace name",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/deployments": {
            "get": {
                "description": "Get all deployments in the specified namespace",
                "produces": [
                    "application/json"
                ],
                "summary": "Get deployments for a specific namespace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace name",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/deployments/display": {
            "get": {
                "description": "Display all deployments in the specified namespace",
                "produces": [
                    "application/json"
                ],
                "summary": "Display deployments for a specific namespace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace name",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/secrets/display": {
            "get": {
                "description": "Display all secrets in the specified namespace",
                "produces": [
                    "application/json"
                ],
                "summary": "Display secrets for a specific namespace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace name",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/updateConfigMap": {
            "post": {
                "description": "Update a config map in a specific namespace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a config map",
                "parameters": [
                    {
                        "description": "ConfigMap Update Request Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ConfigMapPatchRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/updateDeploymentImage": {
            "post": {
                "description": "Update the image of a deployment in a specific namespace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update deployment image",
                "parameters": [
                    {
                        "description": "Deployment Image Set Request Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.DeploymentPatchRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/updateSecret": {
            "post": {
                "description": "Update a secret in a specific namespace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a secret",
                "parameters": [
                    {
                        "description": "Secret Update Request Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SecretPatchRequestBody"
                        }
                    },
                    {
                        "description": "Secret patch request body",
                        "name": "secretPatchRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SecretPatchRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ConfigMapPatchRequestBody": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                }
            }
        },
        "controllers.DeploymentPatchRequestBody": {
            "type": "object",
            "properties": {
                "container": {
                    "type": "string"
                },
                "deployment": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                }
            }
        },
        "controllers.NSClonerRequestBody": {
            "type": "object",
            "properties": {
                "sourceNamespace": {
                    "type": "string"
                },
                "targetNamespace": {
                    "type": "string"
                }
            }
        },
        "controllers.SecretPatchRequestBody": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Kubernetes Namespace Cloner API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
