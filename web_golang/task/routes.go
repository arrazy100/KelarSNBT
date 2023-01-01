package task

import "github.com/gin-gonic/gin"

type TaskRouteController[T any] struct {
	taskController TaskController[T]
}

func NewTaskRouteController[T any](taskController TaskController[T]) TaskRouteController[T] {
	return TaskRouteController[T]{taskController}
}

func (repo *TaskRouteController[T]) Route(tr *gin.RouterGroup) {
	router := tr.Group("/tasks")

	router.GET("/", repo.taskController.FindAll)
	router.GET("/:taskId", repo.taskController.FindById)
	router.POST("/", repo.taskController.Create)
	router.PATCH("/:taskId", repo.taskController.Update)
	router.DELETE("/:taskId", repo.taskController.Delete)

	router.PATCH("/setQuestions/:taskId", repo.taskController.SetQuestions)
}
