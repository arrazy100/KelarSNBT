package question_routes

import (
	question_controllers "main/question/controllers"

	"github.com/gin-gonic/gin"
)

type QuestionRouteController struct {
	questionController question_controllers.QuestionController
}

func NewQuestionRouteController(questionController question_controllers.QuestionController) QuestionRouteController {
	return QuestionRouteController{questionController}
}

func (qrc *QuestionRouteController) QuestionRoute(qr *gin.RouterGroup) {
	router := qr.Group("/questions")

	router.GET("/", qrc.questionController.FindAll)
	router.GET("/:questionId", qrc.questionController.FindById)
	router.POST("/", qrc.questionController.Create)
	router.PATCH("/:questionId", qrc.questionController.Update)
	router.DELETE("/:questionId", qrc.questionController.Delete)

	router.POST("/createAnswer/:questionId", qrc.questionController.CreateAnswer)
	router.PATCH("/updateAnswer/:questionId", qrc.questionController.UpdateAnswer)
	router.DELETE("/deleteAnswer/:answerId", qrc.questionController.DeleteAnswer)
}
