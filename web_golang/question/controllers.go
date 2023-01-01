package question

import (
	"main/crud_generics"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type QuestionController[T any] struct {
	genericController crud_generics.CRUDControllerRepo[T]
}

func NewQuestionController[T any](genericController crud_generics.CRUDControllerRepo[T]) QuestionController[T] {
	return QuestionController[T]{genericController}
}

// GetQuestions godoc
// @Summary get questions
// @Schemes
// @Description Get all available questions
// @Tags question
// @Accept json
// @Produce json
// @Param page query int false "write page number"
// @Param limit query int false "write limit number"
// @Success 200 {array} QuestionDB
// @Router /questions [get]
func (repo *QuestionController[T]) FindAll(ctx *gin.Context) {
	repo.genericController.FindAll(ctx)
}

// GetQuestionById godoc
// @Summary get question by id
// @Schemes
// @Description Get question by id
// @Tags question
// @Accept json
// @Produce json
// @Param questionId path string true "Write question id"
// @Success 200 {object} QuestionDB
// @Router /questions/{questionId} [get]
func (repo *QuestionController[T]) FindById(ctx *gin.Context) {
	repo.genericController.FindById(ctx)
}

// PostQuestion godoc
// @Summary create a new question
// @Schemes
// @Description Create a new question to choose for Task
// @Tags question
// @Produce json
// @Param question body CreateQuestion true "Question JSON"
// @Success 200 {object} QuestionDB
// @Router /questions [post]
func (repo *QuestionController[T]) Create(ctx *gin.Context) {
	repo.genericController.Create(ctx)
}

// PatchQuestion godoc
// @Summary Edit a question
// @Schemes
// @Description Edit a question by id
// @Tags question
// @Produce json
// @Param questionId path string true "Write question id"
// @Param question body UpdateQuestion true "Question JSON"
// @Success 200 {object} QuestionDB
// @Router /questions/{questionId} [patch]
func (repo *QuestionController[T]) Update(ctx *gin.Context) {
	repo.genericController.Update(ctx)
}

// DeleteQuestion godoc
// @Summary delete a question
// @Schemes
// @Description Delete a question by id
// @Tags question
// @Param questionId path string true "Write question id"
// @Success 204
// @Router /questions/{questionId} [delete]
func (repo *QuestionController[T]) Delete(ctx *gin.Context) {
	repo.genericController.Delete(ctx)
}

// PostAnswer godoc
// @Summary add answer to a question
// @Schemes
// @Description Add an answer to a question
// @Tags answer
// @Produce json
// @Param questionId path string true "write question id"
// @Param answer body CreateAnswer true "Answer JSON"
// @Success 200 {object} AnswerDB
// @Router /answers/{questionId} [post]
func (repo *QuestionController[T]) CreateAnswer(ctx *gin.Context) {
	questionId := ctx.Param("questionId")

	var answer *CreateAnswer
	if err := ctx.ShouldBindJSON(&answer); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	updatedQuestion, err := CreateAnswerService(repo.genericController.GetCollection(), repo.genericController.GetContext(), questionId, answer)

	if err != nil {
		if strings.Contains(err.Error(), "not exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})

			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedQuestion})
}

// PatchAnswer godoc
// @Summary Edit an answer
// @Schemes
// @Description Edit an answer by question id
// @Tags answer
// @Produce json
// @Param questionId path string true "Write question id"
// @Param question body UpdateAnswer true "Answer JSON"
// @Success 200 {object} QuestionDB
// @Router /answers/{questionId} [patch]
func (repo *QuestionController[T]) UpdateAnswer(ctx *gin.Context) {
	questionId := ctx.Param("questionId")

	var answer *UpdateAnswer
	if err := ctx.ShouldBindJSON(&answer); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	updatedQuestion, err := UpdateAnswerService(repo.genericController.GetCollection(), repo.genericController.GetContext(), questionId, answer)

	if err != nil {
		if strings.Contains(err.Error(), "not exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})

			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedQuestion})
}

// DeleteAnswer godoc
// @Summary delete an answer
// @Schemes
// @Description Delete an answer by id
// @Tags answer
// @Param answerId path string true "Write answer id"
// @Success 204
// @Router /answers/{answerId} [delete]
func (repo *QuestionController[T]) DeleteAnswer(ctx *gin.Context) {
	answerId := ctx.Param("answerId")

	err := DeleteAnswerService(repo.genericController.GetCollection(), repo.genericController.GetContext(), answerId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
