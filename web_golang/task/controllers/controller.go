package task_controllers

import (
	task_models "main/task/models"
	task_services "main/task/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService task_services.TaskService
}

func NewTaskController(taskService task_services.TaskService) TaskController {
	return TaskController{taskService}
}

// PostTask godoc
// @Summary create a new task
// @Schemes
// @Description Create a new task for events
// @Tags task
// @Produce json
// @Param task body task_models.CreateTask true "Task JSON"
// @Success 200 {object} task_models.TaskDB
// @Router /tasks [post]
func (tc *TaskController) Create(ctx *gin.Context) {
	var task *task_models.CreateTask

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())

		return
	}

	newTask, err := tc.taskService.Create(task)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newTask})
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
// @Success 200 {array} task_models.TaskDB
// @Router /tasks [get]
func (tc *TaskController) FindAll(ctx *gin.Context) {
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

	tasks, err := tc.taskService.FindAll(intPage, intLimit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(tasks), "data": tasks})
}

// GetTaskById godoc
// @Summary get task by id
// @Schemes
// @Description Get task by id
// @Tags task
// @Accept json
// @Produce json
// @Param taskId path string true "Write task id"
// @Success 200 {object} task_models.TaskDB
// @Router /tasks/{taskId} [get]
func (tc *TaskController) FindById(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	task, err := tc.taskService.FindById(taskId)

	if err != nil {
		if strings.Contains(err.Error(), "No document") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})

			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": task})
}

// DeleteTask godoc
// @Summary delete a task
// @Schemes
// @Description Delete a task by id
// @Tags task
// @Param taskId path string true "Write task id"
// @Success 204
// @Router /tasks/{taskId} [delete]
func (tc *TaskController) Delete(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	err := tc.taskService.Delete(taskId)

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

// PatchTask godoc
// @Summary Edit a task
// @Schemes
// @Description Edit a task by id
// @Tags task
// @Produce json
// @Param taskId path string true "Write task id"
// @Param task body task_models.UpdateTask true "Task JSON"
// @Success 200 {object} task_models.TaskDB
// @Router /tasks/{taskId} [patch]
func (tc *TaskController) Update(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	var task *task_models.UpdateTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	updatedTask, err := tc.taskService.Update(taskId, task)
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

// SetQuestionsToTask godoc
// @Summary set questions for a task
// @Schemes
// @Description Set questions for a task by task id
// @Tags task
// @Produce json
// @Param taskId path string true "write task id"
// @Param task body task_models.SetQuestion true "Task JSON"
// @Success 200 {object} task_models.SetQuestion
// @Router /tasks/setQuestions/{taskId} [patch]
func (tc *TaskController) SetQuestions(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	var questions *task_models.SetQuestion
	if err := ctx.ShouldBindJSON(&questions); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	updatedTask, err := tc.taskService.SetQuestions(taskId, questions)
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
