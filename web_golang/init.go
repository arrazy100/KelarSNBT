package main

import (
	"context"
	"fmt"
	"main/question"
	"main/task"
	"main/test_generics"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
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

	server = gin.Default()
}
