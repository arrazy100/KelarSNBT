package task_models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskDB struct {
	Id        primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	StartDate time.Time            `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   time.Time            `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Questions []primitive.ObjectID `json:"questions,omitempty" bson:"questions,omitempty"`
}

type CreateTask struct {
	StartDate time.Time            `json:"start_date" bson:"start_date" binding:"required"`
	EndDate   time.Time            `json:"end_date" bson:"end_date" binding:"required"`
	Questions []primitive.ObjectID `json:"questions" bson:"questions" binding:"required"`
}

type UpdateTask struct {
	StartDate time.Time            `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   time.Time            `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Questions []primitive.ObjectID `json:"questions,omitempty" bson:"questions,omitempty"`
}
