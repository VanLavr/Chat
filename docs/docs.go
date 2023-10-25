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
        "/chatroom": {
            "put": {
                "description": "Update an existing chatroom",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chatroom"
                ],
                "summary": "Update chatroom",
                "operationId": "update-chatroom",
                "parameters": [
                    {
                        "description": "Chatroom information",
                        "name": "chatroom",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Chatroom"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new chatroom",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chatroom"
                ],
                "summary": "Create chatroom",
                "operationId": "create-chatroom",
                "parameters": [
                    {
                        "description": "Chatroom information",
                        "name": "chatroom",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Chatroom"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an existing chatroom",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chatroom"
                ],
                "summary": "Delete chatroom",
                "operationId": "delete-chatroom",
                "parameters": [
                    {
                        "description": "Chatroom deletion information",
                        "name": "DeleteChatDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.DeleteChatDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/chatroom/{id}": {
            "get": {
                "description": "Get a chatroom by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chatroom"
                ],
                "summary": "Get chatroom by ID",
                "operationId": "get-chatroom-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Chatroom ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/chatrooms/{limit}": {
            "get": {
                "description": "Get a list of chatrooms",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chatroom"
                ],
                "summary": "Get chatrooms",
                "operationId": "get-chatrooms",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Number of chatrooms to retrieve",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/message": {
            "put": {
                "description": "Update a message with new content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Update a message",
                "operationId": "update-message",
                "parameters": [
                    {
                        "description": "Message object",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/message/{id}": {
            "delete": {
                "description": "Delete a message with a specified ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Delete a message",
                "operationId": "delete-message",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Message ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/messages/{chat}/{limit}": {
            "get": {
                "description": "Retrieve messages for a specific chat with a specified limit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Get chat messages",
                "operationId": "get-chat-messages",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Chat ID",
                        "name": "chat",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit of messages to retrieve",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/messages/{limit}": {
            "get": {
                "description": "Retrieve messages with a specified limit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Get messages",
                "operationId": "get-messages",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit of messages to retrieve",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/messages/{user}/{limit}": {
            "get": {
                "description": "Retrieve messages for a specific user with a specified limit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Get user messages",
                "operationId": "get-user-messages",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit of messages to retrieve",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/user": {
            "put": {
                "description": "Update an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a user",
                "operationId": "update-user",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete a user",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/user/enterChatroom": {
            "post": {
                "description": "Enter a chatroom",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chatroom"
                ],
                "summary": "Enter chatroom",
                "operationId": "enter-chatroom",
                "parameters": [
                    {
                        "description": "User chat information",
                        "name": "EnterChatDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.EnterChatDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/user/jwt": {
            "get": {
                "description": "Generate a JWT token for the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Generate JWT token",
                "operationId": "get-jwt",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Retrieve a user with a specified ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get a user",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/user/{uid}/chatroom/{cid}": {
            "post": {
                "description": "Join a chatroom with the specified user ID and chatroom ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Join a chatroom",
                "operationId": "join-chatroom",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Chatroom ID",
                        "name": "cid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/user/{uid}/leaveRoom/{chatroom_id}": {
            "get": {
                "description": "Leave a chatroom",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chatroom"
                ],
                "summary": "Leave chatroom",
                "operationId": "leave-chatroom",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Chatroom ID",
                        "name": "chatroom_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/users/{limit}": {
            "get": {
                "description": "Retrieve a list of users with a specified limit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get users",
                "operationId": "get-users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit of users to retrieve",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "delivery.DeleteChatDTO": {
            "type": "object",
            "properties": {
                "cid": {
                    "type": "integer"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "delivery.EnterChatDTO": {
            "type": "object",
            "properties": {
                "cid": {
                    "type": "integer"
                },
                "roomPassword": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "models.Chatroom": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "chat_id": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "sended": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "content": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object"
        }
    },
    "securityDefinitions": {
        "JWTAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "chat",
	Description:      "pidorpidorasi",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
