package task

import (
	"github.com/gofiber/fiber/v2"
)

type TaskRouteController[T any] struct {
	taskController TaskController[T]
}

func NewTaskRouteController[T any](taskController TaskController[T]) TaskRouteController[T] {
	return TaskRouteController[T]{taskController}
}

func (repo *TaskRouteController[T]) Route(tr fiber.Router) {
	router := tr.Group("/tasks")

	router.Get("/", repo.taskController.FindAll)
	router.Get("/:taskId", repo.taskController.FindById)
	router.Post("/", repo.taskController.Create)
	router.Patch("/:taskId", repo.taskController.Update)
	router.Delete("/:taskId", repo.taskController.Delete)

	router.Patch("/setQuestions/:taskId", repo.taskController.SetQuestions)
}
