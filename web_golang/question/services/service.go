package question_services

import question_models "main/question/models"

type QuestionService interface {
	Create(*question_models.CreateQuestion) (*question_models.QuestionDB, error)
	Update(string, *question_models.UpdateQuestion) (*question_models.QuestionDB, error)
	FindById(string) (*question_models.QuestionDB, error)
	FindAll(page int, limit int) ([]*question_models.QuestionDB, error)
	Delete(string) error

	CreateAnswer(string, *question_models.CreateAnswer) (*question_models.QuestionDB, error)
	UpdateAnswer(string, *question_models.UpdateAnswer) (*question_models.QuestionDB, error)
	DeleteAnswer(string) error
}
