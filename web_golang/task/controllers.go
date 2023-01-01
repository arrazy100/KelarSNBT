package task

import (
	crud_controllers "main/crud/controllers"
	crud_services "main/crud/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	crudService    crud_services.CRUDService
	CrudController crud_controllers.CRUDController
}

func NewTaskController(param crud_services.Param, name string, singleName string) TaskController {
	crudService := crud_services.NewCRUDService(param)
	crudController := crud_controllers.NewCRUDController(crudService, name, singleName)

	return TaskController{crudService, crudController}
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
func (c *TaskController) FindAll(ctx *gin.Context) {
	c.CrudController.FindAll(ctx)
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
func (c *TaskController) FindById(ctx *gin.Context) {
	c.CrudController.FindById(ctx)
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
func (c *TaskController) Create(ctx *gin.Context) {
	c.CrudController.Create(ctx)
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
func (c *TaskController) Update(ctx *gin.Context) {
	c.CrudController.Update(ctx)
}

// DeleteTask godoc
// @Summary delete a task
// @Schemes
// @Description Delete a task by id
// @Tags task
// @Param taskId path string true "Write task id"
// @Success 204
// @Router /tasks/{taskId} [delete]
func (c *TaskController) Delete(ctx *gin.Context) {
	c.CrudController.Delete(ctx)
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
func (c *TaskController) SetQuestions(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	var questions *SetQuestion
	if err := ctx.ShouldBindJSON(&questions); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	param := c.crudService.GetParameter()
	updatedTask, err := SetQuestionsService(param.Collection, param.Ctx, taskId, questions)

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
