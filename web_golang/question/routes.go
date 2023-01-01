package question

import (
	"github.com/gin-gonic/gin"
)

type QuestionRouteController struct {
	questionController QuestionController
}

func NewQuestionRouteController(questionController QuestionController) QuestionRouteController {
	return QuestionRouteController{questionController}
}

func (rc *QuestionRouteController) QuestionRoute(rg *gin.RouterGroup) {
	router := rg.Group("/questions")

	router.GET("/", rc.questionController.FindAll)
	router.GET("/:questionId", rc.questionController.FindById)
	router.POST("/", rc.questionController.Create)
	router.PATCH("/:questionId", rc.questionController.Update)
	router.DELETE("/:questionId", rc.questionController.Delete)
}

func (rc *QuestionRouteController) AnswerRoute(rg *gin.RouterGroup) {
	router := rg.Group("/answers/")

	router.POST("/:questionId", rc.questionController.CreateAnswer)
	router.PATCH("/:questionId", rc.questionController.UpdateAnswer)
	router.DELETE("/:answerId", rc.questionController.DeleteAnswer)
}
