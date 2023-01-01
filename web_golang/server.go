package main

import (
	"log"
	docs "main/docs"
	"main/question"
	"main/task"
	"main/test_generics"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func startGinServer() {
	router := server.Group("/api")

	question.RouteInit(router)
	task.RouteInit(router)
	test_generics.RouteInit(router)

	docs.SwaggerInfo.BasePath = "/api/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Fatal(server.Run(":8080"))
}
