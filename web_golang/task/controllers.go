package task

import (
	"errors"
	"main/common"
	"main/crud_generics"
	"main/logs"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type TaskController[T any] struct {
	genericController crud_generics.CRUDControllerRepo[T]
}

func NewTaskController[T any](genericController crud_generics.CRUDControllerRepo[T]) TaskController[T] {
	return TaskController[T]{genericController}
}

// GetTasks godoc
// @Summary get tasks
// @Schemes
// @Description Get all available tasks
// @Tags task
// @Accept json
// @Produce json
// @Param page query int false "write page number"
// @Param limit query int false "write limit number"
// @Success 200 {array} TaskDB
// @Router /tasks [get]
func (repo *TaskController[T]) FindAll(ctx *fiber.Ctx) error {
	repo.genericController.FindAll(ctx)

	return nil
}

// GetTaskById godoc
// @Summary get task by id
// @Schemes
// @Description Get task by id
// @Tags task
// @Accept json
// @Produce json
// @Param taskId path string true "Write task id"
// @Success 200 {object} TaskDB
// @Router /tasks/{taskId} [get]
func (repo *TaskController[T]) FindById(ctx *fiber.Ctx) error {
	repo.genericController.FindById(ctx)

	return nil
}

// PostTask godoc
// @Summary create a new task
// @Schemes
// @Description Create a new task for events
// @Tags task
// @Produce json
// @Param task body CreateTask true "Task JSON"
// @Success 200 {object} TaskDB
// @Router /tasks [post]
func (repo *TaskController[T]) Create(ctx *fiber.Ctx) error {
	repo.genericController.Create(ctx)

	return nil
}

// PatchTask godoc
// @Summary Edit a task
// @Schemes
// @Description Edit a task by id
// @Tags task
// @Produce json
// @Param taskId path string true "Write task id"
// @Param task body UpdateTask true "Task JSON"
// @Success 200 {object} TaskDB
// @Router /tasks/{taskId} [patch]
func (repo *TaskController[T]) Update(ctx *fiber.Ctx) error {
	repo.genericController.Update(ctx)

	return nil
}

// DeleteTask godoc
// @Summary delete a task
// @Schemes
// @Description Delete a task by id
// @Tags task
// @Param taskId path string true "Write task id"
// @Success 204
// @Router /tasks/{taskId} [delete]
func (repo *TaskController[T]) Delete(ctx *fiber.Ctx) error {
	repo.genericController.Delete(ctx)

	return nil
}

// SetQuestionsToTask godoc
// @Summary set questions for a task
// @Schemes
// @Description Set questions for a task by task id
// @Tags task
// @Produce json
// @Param taskId path string true "write task id"
// @Param task body SetQuestion true "Task JSON"
// @Success 200 {object} SetQuestion
// @Router /tasks/setQuestions/{taskId} [patch]
func (repo *TaskController[T]) SetQuestions(ctx *fiber.Ctx) error {
	logs.Info("User requesting to set questions")

	taskId := ctx.Params("taskId")

	var questions *SetQuestion

	logs.Debug("Parsing questions from JSON")
	if err := ctx.BodyParser(&questions); err != nil {
		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

		return err
	}

	logs.Debug("Validating questions from JSON")
	errs := common.Validate(questions)
	if errs != nil {
		ctx.Status(http.StatusBadRequest).JSON(errs)

		encoded, _ := ctx.App().Config().JSONEncoder(errs)
		logs.Error("Failed to validate questions from JSON. " + string(encoded))

		return errors.New(string(encoded))
	}

	updatedTask, err := SetQuestionsService(repo.genericController.GetCollection(), repo.genericController.GetContext(), taskId, questions)
	logs.DebugObject("Saving all question to database.", questions)

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
	logs.DebugObject("All questions are saved.", questions)

	ctx.SendStatus(http.StatusOK)
	ctx.JSON(fiber.Map{"status": "success", "data": updatedTask})

	return nil
}
