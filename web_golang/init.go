package main

import (
	"context"
	"fmt"
	"log"
	"main/question"
	"main/task"
	"main/test_generics"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server     *fiber.App
	outputFile *os.File
)

func init() {
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNSTRING")))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	question.Init(client, ctx)
	task.Init(client, ctx)
	test_generics.Init(client, ctx)

	server = fiber.New(fiber.Config{
		Prefork:       true,
		StrictRouting: true,
		CaseSensitive: true,
	})

	SetLog()
}

func SetLog() {
	var err error

	outputFile, err = os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(outputFile)

	server.Use(logger.New(logger.Config{
		Output: outputFile,
	}))
}
