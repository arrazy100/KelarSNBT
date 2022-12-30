package achievement_models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AchievementDB struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	TaskId    primitive.ObjectID `json:"task_id,omitempty" bson:"task_id,omitempty"`
	Grade     int                `json:"grade,omitempty" bson:"grade,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateAchievement struct {
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id" binding:"required"`
	TaskId    primitive.ObjectID `json:"task_id" bson:"task_id" binding:"required"`
	Grade     int                `json:"grade" bson:"grade" binding:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateAchievement struct {
	UserId    primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	TaskId    primitive.ObjectID `json:"task_id,omitempty" bson:"task_id,omitempty"`
	Grade     int                `json:"grade,omitempty" bson:"grade,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
