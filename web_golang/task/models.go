package task

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskDB struct {
	Id        primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string               `json:"name,omitempty" bson:"name,omitempty"`
	StartDate time.Time            `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   time.Time            `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Questions []primitive.ObjectID `json:"questions,omitempty" bson:"questions,omitempty"`
}

type CreateTask struct {
	Name      string    `json:"name" bson:"name" validate:"required"`
	StartDate time.Time `json:"start_date" bson:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" bson:"end_date" validate:"required"`
}

type UpdateTask struct {
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	StartDate time.Time `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty" bson:"end_date,omitempty"`
}

type SetQuestion struct {
	Questions []primitive.ObjectID `json:"questions" bson:"questions" binding:"required"`
}
