package main

import (
	"log"
	"main/auth"
	"main/question"
	"main/task"
	"main/test_generics"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "main/docs"
)

var (
	router fiber.Router
)

func startFiberServer() {
	router = server.Group("/api")

	question.RouteInit(router)
	task.RouteInit(router)
	test_generics.RouteInit(router)

	SetAuthMiddleware()

	server.Get("/api/swagger/*", swagger.HandlerDefault)

	log.Fatal(server.Listen(":8080"))
}

func SetAuthMiddleware() {
	server.Use(auth.AuthAdminMiddleware)
}
