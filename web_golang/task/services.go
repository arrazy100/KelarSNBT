package task

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetQuestionsService(collection *mongo.Collection, ctx context.Context, id string, data *SetQuestion) (*TaskDB, error) {
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

	_, err := collection.UpdateOne(ctx, query, removeAll)
	if err != nil {
		return nil, err
	}

	res := collection.FindOneAndUpdate(ctx, query, update, options)

	var updatedTask *TaskDB

	if err := res.Decode(&updatedTask); err != nil {
		return nil, errors.New("Task with Id " + id + " not exists")
	}

	return updatedTask, nil
}
