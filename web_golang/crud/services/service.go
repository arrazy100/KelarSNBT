package crud_services

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
)

type Param struct {
	Collection  *mongo.Collection
	Ctx         context.Context
	ResultModel interface{}
	CreateModel interface{}
	UpdateModel interface{}
}

type CRUDService interface {
	Create(interface{}) (interface{}, error)
	Update(string, interface{}) (interface{}, error)
	FindAll(page int, limit int) (interface{}, error)
	FindById(string) (interface{}, error)
	Delete(string) error

	GetResultModel() reflect.Type
	GetCreateModel() reflect.Type
	GetUpdateModel() reflect.Type
	GetParameter() Param
}
