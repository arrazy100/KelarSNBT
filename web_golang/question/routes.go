package question

import (
	"github.com/gofiber/fiber/v2"
)

type QuestionRouteController[T any] struct {
	questionController QuestionController[T]
}

func NewQuestionRouteController[T any](questionController QuestionController[T]) QuestionRouteController[T] {
	return QuestionRouteController[T]{questionController}
}

func (repo *QuestionRouteController[T]) QuestionRoute(rg fiber.Router) {
	router := rg.Group("/questions")

	router.Get("/", repo.questionController.FindAll)
	router.Get("/:questionId", repo.questionController.FindById)
	router.Post("/", repo.questionController.Create)
	router.Patch("/:questionId", repo.questionController.Update)
	router.Delete("/:questionId", repo.questionController.Delete)
}

func (repo *QuestionRouteController[T]) AnswerRoute(rg fiber.Router) {
	router := rg.Group("/answers/")

	router.Post("/:questionId", repo.questionController.CreateAnswer)
	router.Patch("/:questionId", repo.questionController.UpdateAnswer)
	router.Delete("/:answerId", repo.questionController.DeleteAnswer)
}
