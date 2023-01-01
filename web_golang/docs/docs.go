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
                            "$ref": "#/definitions/question.CreateAnswer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/question.AnswerDB"
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
                            "$ref": "#/definitions/question.UpdateAnswer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/question.QuestionDB"
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
                "parameters": [
                    {
                        "type": "integer",
                        "description": "write page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "write limit number",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/question.QuestionDB"
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
                            "$ref": "#/definitions/question.CreateQuestion"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/question.QuestionDB"
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
                            "$ref": "#/definitions/question.QuestionDB"
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
                            "$ref": "#/definitions/question.UpdateQuestion"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/question.QuestionDB"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "Get all available tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "get tasks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "write page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "write limit number",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/task_models.TaskDB"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new task for events",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "create a new task",
                "parameters": [
                    {
                        "description": "Task JSON",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task_models.CreateTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task_models.TaskDB"
                        }
                    }
                }
            }
        },
        "/tasks/setQuestions/{taskId}": {
            "patch": {
                "description": "Set questions for a task by task id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "set questions for a task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "write task id",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task JSON",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task_models.SetQuestion"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task_models.SetQuestion"
                        }
                    }
                }
            }
        },
        "/tasks/{taskId}": {
            "get": {
                "description": "Get task by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "get task by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write task id",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task_models.TaskDB"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a task by id",
                "tags": [
                    "task"
                ],
                "summary": "delete a task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write task id",
                        "name": "taskId",
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
                "description": "Edit a task by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Edit a task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write task id",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task JSON",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task_models.UpdateTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task_models.TaskDB"
                        }
                    }
                }
            }
        },
        "/tests": {
            "get": {
                "description": "Get all available tests",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "get tests",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "write page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "write limit number",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/test.TestDB"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new test",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "create a new test",
                "parameters": [
                    {
                        "description": "Task JSON",
                        "name": "test",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/test.CreateTest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/test.TestDB"
                        }
                    }
                }
            }
        },
        "/tests/{testId}": {
            "get": {
                "description": "Get test by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "get test by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write test id",
                        "name": "testId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/test.TestDB"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a test by id",
                "tags": [
                    "test"
                ],
                "summary": "delete a test",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write test id",
                        "name": "testId",
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
                "description": "Edit a test by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "Edit a test",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Write test id",
                        "name": "testId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Test JSON",
                        "name": "test",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/test.UpdateTest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/test.TestDB"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "question.AnswerDB": {
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
        "question.CreateAnswer": {
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
        "question.CreateQuestion": {
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
        "question.QuestionDB": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/question.AnswerDB"
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
        "question.UpdateAnswer": {
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
        "question.UpdateQuestion": {
            "type": "object",
            "properties": {
                "materi": {
                    "type": "integer"
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "task.CreateTask": {
            "type": "object",
            "required": [
                "end_date",
                "name",
                "start_date"
            ],
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "task.SetQuestion": {
            "type": "object",
            "required": [
                "questions"
            ],
            "properties": {
                "questions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "task.TaskDB": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "task.UpdateTask": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "task_models.CreateTask": {
            "type": "object",
            "required": [
                "end_date",
                "name",
                "start_date"
            ],
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "task_models.SetQuestion": {
            "type": "object",
            "required": [
                "questions"
            ],
            "properties": {
                "questions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "task_models.TaskDB": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "task_models.UpdateTask": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "test.CreateTest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "test.TestDB": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "test.UpdateTest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
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
