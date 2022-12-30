package task_services

import (
	"context"
	"errors"
	task_models "main/task/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskServiceImpl struct {
	taskCollection *mongo.Collection
	ctx            context.Context
}

func NewTaskService(taskCollection *mongo.Collection, ctx context.Context) TaskService {
	return &TaskServiceImpl{taskCollection, ctx}
}

func (t *TaskServiceImpl) Create(task *task_models.CreateTask) (*task_models.TaskDB, error) {
	res, err := t.taskCollection.InsertOne(t.ctx, task)

	if err != nil {
		return nil, err
	}

	var newTask *task_models.TaskDB
	query := bson.M{"_id": res.InsertedID}

	if err = t.taskCollection.FindOne(t.ctx, query).Decode(&newTask); err != nil {
		return nil, err
	}

	return newTask, nil
}

func (t *TaskServiceImpl) Update(id string, data *task_models.UpdateTask) (*task_models.TaskDB, error) {
	var doc interface{} = &data

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	options := options.FindOneAndUpdate().SetReturnDocument(1)
	res := t.taskCollection.FindOneAndUpdate(t.ctx, query, update, options)

	var updatedTask *task_models.TaskDB

	if err := res.Decode(&updatedTask); err != nil {
		return nil, errors.New("Task with Id " + id + " not exists")
	}

	return updatedTask, nil
}

func (t *TaskServiceImpl) FindById(id string) (*task_models.TaskDB, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	var task *task_models.TaskDB

	if err := t.taskCollection.FindOne(t.ctx, query).Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("No document with Id " + id)
		}

		return nil, err
	}

	return task, nil
}

func (t *TaskServiceImpl) FindAll(page int, limit int) ([]*task_models.TaskDB, error) {
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

	cursor, err := t.taskCollection.Find(t.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(t.ctx)

	var tasks []*task_models.TaskDB

	for cursor.Next(t.ctx) {
		task := &task_models.TaskDB{}
		err := cursor.Decode(task)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return []*task_models.TaskDB{}, nil
	}

	return tasks, nil
}

func (t *TaskServiceImpl) Delete(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := t.taskCollection.DeleteOne(t.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("No document with Id " + id + " exists")
	}

	return nil
}

func (t *TaskServiceImpl) SetQuestions(id string, data *task_models.SetQuestion) (*task_models.TaskDB, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	updateFields := bson.M{
		"questions": bson.M{"$each": &data.Questions},
	}
	removeFields := bson.M{
		"questions": []primitive.ObjectID{},
	}

	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$push", Value: updateFields}}
	removeAll := bson.D{{Key: "$set", Value: removeFields}}
	options := options.FindOneAndUpdate().SetReturnDocument(1).SetUpsert(true)

	_, err := t.taskCollection.UpdateOne(t.ctx, query, removeAll)
	if err != nil {
		return nil, err
	}

	res := t.taskCollection.FindOneAndUpdate(t.ctx, query, update, options)

	var updatedTask *task_models.TaskDB

	if err := res.Decode(&updatedTask); err != nil {
		return nil, errors.New("Task with Id " + id + " not exists")
	}

	return updatedTask, nil
}
