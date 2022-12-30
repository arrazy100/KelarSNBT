package main

import (
	"context"
	"fmt"
	"log"
	question_controllers "main/question/controllers"
	question_routes "main/question/routes"
	question_services "main/question/services"
	task_controllers "main/task/controllers"
	task_routes "main/task/routes"
	task_services "main/task/services"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	docs "main/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	server *gin.Engine

	questionCollection      *mongo.Collection
	questionService         question_services.QuestionService
	questionController      question_controllers.QuestionController
	questionRouteController question_routes.QuestionRouteController

	taskCollection      *mongo.Collection
	taskService         task_services.TaskService
	taskController      task_controllers.TaskController
	taskRouteController task_routes.TaskRouteController
)

// @title           Kelar SNBT
// @version         1.0
// @description     No Description
// @termsOfService  https://tos.kelarsnbt.com

// @contact.name   Muhammad Afdhal Arrazy
// @contact.url    https://github.com/arrazy100
// @contact.email  afdhalarrazy111@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3001
// @BasePath  /api/
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

	taskCollection = client.Database(os.Getenv("MONGO_DBNAME")).Collection("daftar_tugas")
	taskService = task_services.NewTaskService(taskCollection, ctx)
	taskController = task_controllers.NewTaskController(taskService)
	taskRouteController = task_routes.NewTaskRouteController(taskController)

	server = gin.Default()
}

func startGinServer() {
	router := server.Group("/api")

	questionRouteController.QuestionRoute(router)
	questionRouteController.AnswerRoute(router)
	taskRouteController.TaskRoute(router)

	docs.SwaggerInfo.BasePath = "/api/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Fatal(server.Run(":8080"))
}
