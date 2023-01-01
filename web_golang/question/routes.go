package question

import (
	"github.com/gin-gonic/gin"
)

type QuestionRouteController[T any] struct {
	questionController QuestionController[T]
}

func NewQuestionRouteController[T any](questionController QuestionController[T]) QuestionRouteController[T] {
	return QuestionRouteController[T]{questionController}
}

func (repo *QuestionRouteController[T]) QuestionRoute(rg *gin.RouterGroup) {
	router := rg.Group("/questions")

	router.GET("/", repo.questionController.FindAll)
	router.GET("/:questionId", repo.questionController.FindById)
	router.POST("/", repo.questionController.Create)
	router.PATCH("/:questionId", repo.questionController.Update)
	router.DELETE("/:questionId", repo.questionController.Delete)
}

func (repo *QuestionRouteController[T]) AnswerRoute(rg *gin.RouterGroup) {
	router := rg.Group("/answers/")

	router.POST("/:questionId", repo.questionController.CreateAnswer)
	router.PATCH("/:questionId", repo.questionController.UpdateAnswer)
	router.DELETE("/:answerId", repo.questionController.DeleteAnswer)
}
