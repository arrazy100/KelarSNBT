package task

import (
	"context"
	"main/crud_generics"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection          *mongo.Collection
	controller          crud_generics.CRUDControllerRepo[TaskDB]
	taskController      TaskController[TaskDB]
	taskRouteController TaskRouteController[TaskDB]
)

func Init(client *mongo.Client, ctx context.Context) {
	collection = client.Database(os.Getenv("MONGO_DBNAME")).Collection("daftar_tugas")
	controller = crud_generics.NewCRUDControllerRepo[TaskDB](
		collection,
		ctx,
		"tasks",
		"task",
		func() interface{} {
			return &CreateTask{}
		},
		func() interface{} {
			return &UpdateTask{}
		},
	)

	taskController = NewTaskController(controller)
	taskRouteController = NewTaskRouteController(taskController)
}

func RouteInit(router *gin.RouterGroup) {
	taskRouteController.Route(router)
}
