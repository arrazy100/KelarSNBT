package task_routes

import (
	task_controllers "main/task/controllers"

	"github.com/gin-gonic/gin"
)

type TaskRouteController struct {
	taskController task_controllers.TaskController
}

func NewTaskRouteController(taskController task_controllers.TaskController) TaskRouteController {
	return TaskRouteController{taskController}
}

func (trc *TaskRouteController) TaskRoute(tr *gin.RouterGroup) {
	router := tr.Group("/tasks")

	router.GET("/", trc.taskController.FindAll)
	router.GET("/:taskId", trc.taskController.FindById)
	router.POST("/", trc.taskController.Create)
	router.PATCH("/:taskId", trc.taskController.Update)
	router.DELETE("/:taskId", trc.taskController.Delete)

	router.PATCH("/setQuestions/:taskId", trc.taskController.SetQuestions)
}
