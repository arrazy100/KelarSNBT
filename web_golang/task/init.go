package task

import (
	"context"
	crud_services "main/crud/services"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	taskCollection      *mongo.Collection
	taskParam           crud_services.Param
	taskController      TaskController
	taskRouteController TaskRouteController
)

func Init(client *mongo.Client, ctx context.Context) {
	taskCollection = client.Database(os.Getenv("MONGO_DBNAME")).Collection("daftar_tugas")
	taskParam = crud_services.Param{
		Collection:  taskCollection,
		Ctx:         ctx,
		ResultModel: TaskDB{},
		CreateModel: CreateTask{},
		UpdateModel: UpdateTask{},
	}
	taskController = NewTaskController(taskParam, "tasks", "task")
	taskRouteController = NewTaskRouteController(taskController)
}

func RouteInit(router *gin.RouterGroup) {
	taskRouteController.TaskRoute(router)
}
