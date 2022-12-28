package main

import (
	"context"
	"fmt"
	"log"
	question_controllers "main/question/controllers"
	question_routes "main/question/routes"
	question_services "main/question/services"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine

	questionCollection      *mongo.Collection
	questionService         question_services.QuestionService
	questionController      question_controllers.QuestionController
	questionRouteController question_routes.QuestionRouteController
)

func main() {
	startGinServer()
}

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

	questionCollection = client.Database(os.Getenv("MONGO_DBNAME")).Collection("daftar_soal")
	questionService = question_services.NewQuestionService(questionCollection, ctx)
	questionController = question_controllers.NewQuestionController(questionService)
	questionRouteController = question_routes.NewQuestionRouteController(questionController)

	server = gin.Default()
}

func startGinServer() {
	router := server.Group("/api")

	questionRouteController.QuestionRoute(router)
	log.Fatal(server.Run(":8080"))
}
