package question

import "go.mongodb.org/mongo-driver/bson/primitive"

type AnswerDB struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content       string             `json:"content,omitempty" bson:"content,omitempty"`
	CorrectAnswer bool               `json:"correct_answer,omitempty" bson:"correct_answer,omitempty"`
}

type CreateAnswer struct {
	Content       string `json:"content" bson:"content" validate:"required"`
	CorrectAnswer bool   `json:"correct_answer,omitempty" bson:"correct_answer,omitempty"`
}

type UpdateAnswer struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content       string             `json:"content,omitempty" bson:"content,omitempty"`
	CorrectAnswer bool               `json:"correct_answer,omitempty" bson:"correct_answer,omitempty"`
}

type QuestionDB struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Materi   int                `json:"materi,omitempty" bson:"materi,omitempty"`
	Question string             `json:"question,omitempty" bson:"question,omitempty"`
	Answers  []AnswerDB         `json:"answers,omitempty" bson:"answers,omitempty"`
}

type CreateQuestion struct {
	Materi   int    `json:"materi" bson:"materi" validate:"required"`
	Question string `json:"question" bson:"question" validate:"required"`
}

type UpdateQuestion struct {
	Materi   int    `json:"materi,omitempty" bson:"materi,omitempty"`
	Question string `json:"question,omitempty" bson:"question,omitempty"`
}
