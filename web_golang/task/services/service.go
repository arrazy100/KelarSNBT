package task_services

import task_models "main/task/models"

type TaskService interface {
	Create(*task_models.CreateTask) (*task_models.TaskDB, error)
	Update(string, *task_models.UpdateTask) (*task_models.TaskDB, error)
	FindById(string) (*task_models.TaskDB, error)
	FindAll(page int, limit int) ([]*task_models.TaskDB, error)
	Delete(string) error

	SetQuestions(string, *task_models.SetQuestion) (*task_models.TaskDB, error)
}
