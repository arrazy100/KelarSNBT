package question_controllers

import (
	question_models "main/question/models"
	question_services "main/question/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	questionService question_services.QuestionService
}

func NewQuestionController(questionService question_services.QuestionService) QuestionController {
	return QuestionController{questionService}
}

// PostQuestion godoc
// @Summary create a new question
// @Schemes
// @Description Create a new question to choose for Task
// @Tags question
// @Produce json
// @Param question body question_models.CreateQuestion true "Question JSON"
// @Success 200 {object} question_models.QuestionDB
// @Router /questions [post]
func (qc *QuestionController) Create(ctx *gin.Context) {
	var question *question_models.CreateQuestion

	if err := ctx.ShouldBindJSON(&question); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())

		return
	}

	newQuestion, err := qc.questionService.Create(question)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newQuestion})
}

// GetQuestions godoc
// @Summary get questions
// @Schemes
// @Description Get all available questions
// @Tags question
// @Accept json
// @Produce json
// @Success 200 {array} question_models.QuestionDB
// @Router /questions [get]
func (qc *QuestionController) FindAll(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	questions, err := qc.questionService.FindAll(intPage, intLimit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(questions), "data": questions})
}

// GetQuestionById godoc
// @Summary get question by id
// @Schemes
// @Description Get question by id
// @Tags question
// @Accept json
// @Produce json
// @Param questionId path string true "Write question id"
// @Success 200 {object} question_models.QuestionDB
// @Router /questions/{questionId} [get]
func (qc *QuestionController) FindById(ctx *gin.Context) {
	questionId := ctx.Param("questionId")

	question, err := qc.questionService.FindById(questionId)

	if err != nil {
		if strings.Contains(err.Error(), "No document") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})

			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": question})
}

// DeleteQuestion godoc
// @Summary delete a question
// @Schemes
// @Description Delete a question by id
// @Tags question
// @Param questionId path string true "Write question id"
// @Success 204
// @Router /questions/{questionId} [delete]
func (qc *QuestionController) Delete(ctx *gin.Context) {
	questionId := ctx.Param("questionId")

	err := qc.questionService.Delete(questionId)

	if err != nil {
		if strings.Contains(err.Error(), "exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})

			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// PatchQuestion godoc
// @Summary Edit a question
// @Schemes
// @Description Edit a question by id
// @Tags question
// @Produce json
// @Param questionId path string true "Write question id"
// @Param question body question_models.UpdateQuestion true "Question JSON"
// @Success 200 {object} question_models.QuestionDB
// @Router /questions/{questionId} [patch]
func (qc *QuestionController) Update(ctx *gin.Context) {
	questionId := ctx.Param("questionId")

	var question *question_models.UpdateQuestion
	if err := ctx.ShouldBindJSON(&question); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	updatedQuestion, err := qc.questionService.Update(questionId, question)
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

// PostAnswer godoc
// @Summary add answer to a question
// @Schemes
// @Description Add an answer to a question
// @Tags answer
// @Produce json
// @Param questionId path string true "write question id"
// @Param answer body question_models.CreateAnswer true "Answer JSON"
// @Success 200 {object} question_models.AnswerDB
// @Router /answers/{questionId} [post]
func (qc *QuestionController) CreateAnswer(ctx *gin.Context) {
	questionId := ctx.Param("questionId")

	var answer *question_models.CreateAnswer
	if err := ctx.ShouldBindJSON(&answer); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	updatedQuestion, err := qc.questionService.CreateAnswer(questionId, answer)
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
// @Param question body question_models.UpdateAnswer true "Answer JSON"
// @Success 200 {object} question_models.QuestionDB
// @Router /answers/{questionId} [patch]
func (qc *QuestionController) UpdateAnswer(ctx *gin.Context) {
	questionId := ctx.Param("questionId")

	var answer *question_models.UpdateAnswer
	if err := ctx.ShouldBindJSON(&answer); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	updatedQuestion, err := qc.questionService.UpdateAnswer(questionId, answer)
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
func (qc *QuestionController) DeleteAnswer(ctx *gin.Context) {
	answerId := ctx.Param("answerId")

	err := qc.questionService.DeleteAnswer(answerId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
