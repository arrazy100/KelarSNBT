package question

import (
	crud_controllers "main/crud/controllers"
	crud_services "main/crud/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	crudService    crud_services.CRUDService
	CrudController crud_controllers.CRUDController
}

func NewQuestionController(param crud_services.Param, name string, singleName string) QuestionController {
	crudService := crud_services.NewCRUDService(param)
	crudController := crud_controllers.NewCRUDController(crudService, name, singleName)

	return QuestionController{crudService, crudController}
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
func (c *QuestionController) FindAll(ctx *gin.Context) {
	c.CrudController.FindAll(ctx)
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
func (c *QuestionController) FindById(ctx *gin.Context) {
	c.CrudController.FindById(ctx)
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
func (c *QuestionController) Create(ctx *gin.Context) {
	c.CrudController.Create(ctx)
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
func (c *QuestionController) Update(ctx *gin.Context) {
	c.CrudController.Update(ctx)
}

// DeleteQuestion godoc
// @Summary delete a question
// @Schemes
// @Description Delete a question by id
// @Tags question
// @Param questionId path string true "Write question id"
// @Success 204
// @Router /questions/{questionId} [delete]
func (c *QuestionController) Delete(ctx *gin.Context) {
	c.CrudController.Delete(ctx)
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
func (qc *QuestionController) CreateAnswer(ctx *gin.Context) {
	questionId := ctx.Param("questionId")

	var answer *CreateAnswer
	if err := ctx.ShouldBindJSON(&answer); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	param := qc.crudService.GetParameter()
	updatedQuestion, err := CreateAnswerService(param.Collection, param.Ctx, questionId, answer)

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
func (qc *QuestionController) UpdateAnswer(ctx *gin.Context) {
	questionId := ctx.Param("questionId")

	var answer *UpdateAnswer
	if err := ctx.ShouldBindJSON(&answer); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	param := qc.crudService.GetParameter()
	updatedQuestion, err := UpdateAnswerService(param.Collection, param.Ctx, questionId, answer)

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

	param := qc.crudService.GetParameter()
	err := DeleteAnswerService(param.Collection, param.Ctx, answerId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
