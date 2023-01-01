package task

import (
	"main/crud_generics"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
func (repo *TaskController[T]) FindAll(ctx *gin.Context) {
	repo.genericController.FindAll(ctx)
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
func (repo *TaskController[T]) FindById(ctx *gin.Context) {
	repo.genericController.FindById(ctx)
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
func (repo *TaskController[T]) Create(ctx *gin.Context) {
	repo.genericController.Create(ctx)
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
func (repo *TaskController[T]) Update(ctx *gin.Context) {
	repo.genericController.Update(ctx)
}

// DeleteTask godoc
// @Summary delete a task
// @Schemes
// @Description Delete a task by id
// @Tags task
// @Param taskId path string true "Write task id"
// @Success 204
// @Router /tasks/{taskId} [delete]
func (repo *TaskController[T]) Delete(ctx *gin.Context) {
	repo.genericController.Delete(ctx)
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
func (repo *TaskController[T]) SetQuestions(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	var questions *SetQuestion
	if err := ctx.ShouldBindJSON(&questions); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	updatedTask, err := SetQuestionsService(repo.genericController.GetCollection(), repo.genericController.GetContext(), taskId, questions)

	if err != nil {
		if strings.Contains(err.Error(), "not exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})

			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedTask})
}
