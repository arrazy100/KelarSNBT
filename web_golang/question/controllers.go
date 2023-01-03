package question

import (
	"errors"
	"main/common"
	"main/crud_generics"
	"main/logs"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
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
// @Security ApiKeyAuth
func (repo *QuestionController[T]) FindAll(ctx *fiber.Ctx) error {
	repo.genericController.FindAll(ctx)

	return nil
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
func (repo *QuestionController[T]) FindById(ctx *fiber.Ctx) error {
	repo.genericController.FindById(ctx)

	return nil
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
func (repo *QuestionController[T]) Create(ctx *fiber.Ctx) error {
	repo.genericController.Create(ctx)

	return nil
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
func (repo *QuestionController[T]) Update(ctx *fiber.Ctx) error {
	repo.genericController.Update(ctx)

	return nil
}

// DeleteQuestion godoc
// @Summary delete a question
// @Schemes
// @Description Delete a question by id
// @Tags question
// @Param questionId path string true "Write question id"
// @Success 204
// @Router /questions/{questionId} [delete]
func (repo *QuestionController[T]) Delete(ctx *fiber.Ctx) error {
	repo.genericController.Delete(ctx)

	return nil
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
func (repo *QuestionController[T]) CreateAnswer(ctx *fiber.Ctx) error {
	logs.Info("User requesting to create an answer")

	questionId := ctx.Params("questionId")

	logs.Debug("Parsing answer from JSON")
	var answer *CreateAnswer
	if err := ctx.BodyParser(&answer); err != nil {
		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

		return err
	}

	logs.Debug("Validating answer from JSON")
	errs := common.Validate(answer)
	if errs != nil {
		ctx.Status(http.StatusBadRequest).JSON(errs)

		encoded, _ := ctx.App().Config().JSONEncoder(errs)
		logs.Error("Failed to validate answer from JSON. " + string(encoded))

		return errors.New(string(encoded))
	}

	updatedQuestion, err := CreateAnswerService(repo.genericController.GetCollection(), repo.genericController.GetContext(), questionId, answer)
	logs.DebugObject("Saving answer to database.", answer)

	if err != nil {
		if strings.Contains(err.Error(), "not exists") {
			ctx.SendStatus(http.StatusNotFound)
			ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

			return err
		}

		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

		return err
	}
	logs.DebugObject("Answer is saved.", answer)

	ctx.SendStatus(http.StatusOK)
	ctx.JSON(fiber.Map{"status": "success", "data": updatedQuestion})

	return nil
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
func (repo *QuestionController[T]) UpdateAnswer(ctx *fiber.Ctx) error {
	logs.Info("User requesting to update an answer")

	questionId := ctx.Params("questionId")

	logs.Debug("Parsing answer from JSON")
	var answer *UpdateAnswer
	if err := ctx.BodyParser(&answer); err != nil {
		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

		return err
	}

	logs.Debug("Validating answer from JSON")
	errs := common.Validate(answer)
	if errs != nil {
		ctx.Status(http.StatusBadRequest).JSON(errs)

		encoded, _ := ctx.App().Config().JSONEncoder(errs)
		logs.Error("Failed to validate answer from JSON. " + string(encoded))

		return errors.New(string(encoded))
	}

	updatedQuestion, err := UpdateAnswerService(repo.genericController.GetCollection(), repo.genericController.GetContext(), questionId, answer)
	logs.DebugObject("Updating answer to database.", answer)

	if err != nil {
		if strings.Contains(err.Error(), "not exists") {
			ctx.SendStatus(http.StatusNotFound)
			ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

			return err
		}

		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

		return err
	}
	logs.DebugObject("Answer is updated.", answer)

	ctx.SendStatus(http.StatusOK)
	ctx.JSON(fiber.Map{"status": "success", "data": updatedQuestion})

	return nil
}

// DeleteAnswer godoc
// @Summary delete an answer
// @Schemes
// @Description Delete an answer by id
// @Tags answer
// @Param answerId path string true "Write answer id"
// @Success 204
// @Router /answers/{answerId} [delete]
func (repo *QuestionController[T]) DeleteAnswer(ctx *fiber.Ctx) error {
	logs.Info("User requesting to delete an answer")

	answerId := ctx.Params("answerId")

	err := DeleteAnswerService(repo.genericController.GetCollection(), repo.genericController.GetContext(), answerId)
	logs.Debug("Deleting an answer by id")

	if err != nil {
		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

		return err
	}
	logs.Debug("Success deleting an answer")

	ctx.SendStatus(http.StatusNoContent)

	return nil
}
