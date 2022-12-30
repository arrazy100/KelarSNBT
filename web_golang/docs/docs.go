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
        "termsOfService": "https://tos.kelarsnbt.com",
        "contact": {
            "name": "Muhammad Afdhal Arrazy",
            "url": "https://github.com/arrazy100",
            "email": "afdhalarrazy111@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/answers/{answerId}": {
            "delete": {
                "description": "Delete an answer by id",
                "tags": [
                    "answer"
                ],
                "summary": "delete an answer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write answer id",
                        "name": "answerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/answers/{questionId}": {
            "post": {
                "description": "Add an answer to a question",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "answer"
                ],
                "summary": "add answer to a question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "write question id",
                        "name": "questionId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Answer JSON",
                        "name": "answer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/question_models.CreateAnswer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/question_models.AnswerDB"
                        }
                    }
                }
            },
            "patch": {
                "description": "Edit an answer by question id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "answer"
                ],
                "summary": "Edit an answer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write question id",
                        "name": "questionId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Answer JSON",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/question_models.UpdateAnswer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/question_models.QuestionDB"
                        }
                    }
                }
            }
        },
        "/questions": {
            "get": {
                "description": "Get all available questions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "question"
                ],
                "summary": "get questions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/question_models.QuestionDB"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new question to choose for Task",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "question"
                ],
                "summary": "create a new question",
                "parameters": [
                    {
                        "description": "Question JSON",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/question_models.CreateQuestion"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/question_models.QuestionDB"
                        }
                    }
                }
            }
        },
        "/questions/{questionId}": {
            "get": {
                "description": "Get question by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "question"
                ],
                "summary": "get question by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write question id",
                        "name": "questionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/question_models.QuestionDB"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a question by id",
                "tags": [
                    "question"
                ],
                "summary": "delete a question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write question id",
                        "name": "questionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Edit a question by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "question"
                ],
                "summary": "Edit a question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write question id",
                        "name": "questionId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Question JSON",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/question_models.UpdateQuestion"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/question_models.QuestionDB"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "question_models.AnswerDB": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "correct_answer": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "question_models.CreateAnswer": {
            "type": "object",
            "required": [
                "content"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "correct_answer": {
                    "type": "boolean"
                }
            }
        },
        "question_models.CreateQuestion": {
            "type": "object",
            "required": [
                "materi",
                "question"
            ],
            "properties": {
                "materi": {
                    "type": "integer"
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "question_models.QuestionDB": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/question_models.AnswerDB"
                    }
                },
                "id": {
                    "type": "string"
                },
                "materi": {
                    "type": "integer"
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "question_models.UpdateAnswer": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "correct_answer": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "question_models.UpdateQuestion": {
            "type": "object",
            "properties": {
                "materi": {
                    "type": "integer"
                },
                "question": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3001",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "Kelar SNBT",
	Description:      "No Description",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
