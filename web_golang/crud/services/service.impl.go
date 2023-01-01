package crud_services

import (
	"errors"
	"main/logs"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CRUDServiceImpl struct {
	Parameter Param
}

func NewCRUDService(param Param) CRUDService {
	return &CRUDServiceImpl{param}
}

func (c *CRUDServiceImpl) Create(data interface{}) (interface{}, error) {
	res, err := c.Parameter.Collection.InsertOne(c.Parameter.Ctx, data)

	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	newData := reflect.New(c.GetResultModel()).Interface()

	query := bson.M{"_id": res.InsertedID}

	if err = c.Parameter.Collection.FindOne(c.Parameter.Ctx, query).Decode(newData); err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	return newData, nil
}

func (c *CRUDServiceImpl) Update(id string, data interface{}) (interface{}, error) {
	var doc interface{} = &data

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	options := options.FindOneAndUpdate().SetReturnDocument(1)
	res := c.Parameter.Collection.FindOneAndUpdate(c.Parameter.Ctx, query, update, options)

	updatedData := reflect.New(c.GetResultModel()).Interface()

	if err := res.Decode(updatedData); err != nil {
		return nil, errors.New("Data with Id " + id + " not exists")
	}

	return updatedData, nil
}

func (c *CRUDServiceImpl) FindAll(page int, limit int) (interface{}, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))

	query := bson.M{}

	cursor, err := c.Parameter.Collection.Find(c.Parameter.Ctx, query, &opt)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	defer cursor.Close(c.Parameter.Ctx)

	modelType := c.GetResultModel()
	datas := reflect.MakeSlice(reflect.SliceOf(modelType), 0, 0)

	for cursor.Next(c.Parameter.Ctx) {
		data := reflect.New(c.GetResultModel()).Interface()
		err := cursor.Decode(data)

		if err != nil {
			return nil, err
		}

		datas = reflect.Append(datas, reflect.ValueOf(data).Elem())
	}

	if err := cursor.Err(); err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	if datas.Len() == 0 {
		logs.Warning("List is empty, returning empty array")
	}

	return datas.Interface(), nil
}

func (c *CRUDServiceImpl) FindById(id string) (interface{}, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	data := reflect.New(c.GetResultModel()).Interface()

	if err := c.Parameter.Collection.FindOne(c.Parameter.Ctx, query).Decode(data); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("No document with Id " + id)
		}

		return nil, err
	}

	return data, nil
}

func (c *CRUDServiceImpl) Delete(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := c.Parameter.Collection.DeleteOne(c.Parameter.Ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("No document with Id " + id + " exists")
	}

	return nil
}

func (c *CRUDServiceImpl) GetResultModel() reflect.Type {
	return reflect.TypeOf(c.Parameter.ResultModel)
}

func (c *CRUDServiceImpl) GetCreateModel() reflect.Type {
	return reflect.TypeOf(c.Parameter.CreateModel)
}

func (c *CRUDServiceImpl) GetUpdateModel() reflect.Type {
	return reflect.TypeOf(c.Parameter.UpdateModel)
}

func (c *CRUDServiceImpl) GetParameter() Param {
	return c.Parameter
}
