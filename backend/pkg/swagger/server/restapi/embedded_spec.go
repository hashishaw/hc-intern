// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HTTP server in Go with Swagger endpoints definition.",
    "title": "go-rest-api",
    "version": "0.1.0"
  },
  "paths": {
    "/api/interview/{id}": {
      "get": {
        "description": "Returns whether the id is valid.",
        "parameters": [
          {
            "type": "string",
            "description": "The interview session id.",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Returns interview found."
          },
          "404": {
            "description": "Returns interview not found."
          }
        }
      }
    },
    "/api/schedule/{candidateId}": {
      "post": {
        "description": "Returns an interview id.",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "The id of the candidate.",
            "name": "candidateId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the interview id for that candidate.",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "checkHealth",
        "responses": {
          "200": {
            "description": "OK message.",
            "schema": {
              "type": "string",
              "enum": [
                "OK"
              ]
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HTTP server in Go with Swagger endpoints definition.",
    "title": "go-rest-api",
    "version": "0.1.0"
  },
  "paths": {
    "/api/interview/{id}": {
      "get": {
        "description": "Returns whether the id is valid.",
        "parameters": [
          {
            "type": "string",
            "description": "The interview session id.",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Returns interview found."
          },
          "404": {
            "description": "Returns interview not found."
          }
        }
      }
    },
    "/api/schedule/{candidateId}": {
      "post": {
        "description": "Returns an interview id.",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "The id of the candidate.",
            "name": "candidateId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Returns the interview id for that candidate.",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "checkHealth",
        "responses": {
          "200": {
            "description": "OK message.",
            "schema": {
              "type": "string",
              "enum": [
                "OK"
              ]
            }
          }
        }
      }
    }
  }
}`))
}