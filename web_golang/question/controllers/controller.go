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

func (qc *QuestionController) DeleteAnswer(ctx *gin.Context) {
	answerId := ctx.Param("answerId")

	err := qc.questionService.DeleteAnswer(answerId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
