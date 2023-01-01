package crud_generics

import (
	"context"
	"errors"
	"main/logs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateService[T any](collection *mongo.Collection, ctx context.Context, data interface{}) (*T, error) {
	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	var newData *T

	query := bson.M{"_id": res.InsertedID}

	if err = collection.FindOne(ctx, query).Decode(&newData); err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	return newData, nil
}

func FindAllService[T any](collection *mongo.Collection, ctx context.Context, page int, limit int) ([]T, error) {
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

	cursor, err := collection.Find(ctx, query, &opt)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	defer cursor.Close(ctx)

	var datas []T

	for cursor.Next(ctx) {
		var data T
		err := cursor.Decode(&data)

		if err != nil {
			return nil, err
		}

		datas = append(datas, data)
	}

	if err := cursor.Err(); err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	if len(datas) == 0 {
		logs.Warning("List is empty, returning empty array")
		return []T{}, nil
	}

	return datas, nil
}

func FindByIdService[T any](collection *mongo.Collection, ctx context.Context, id string) (*T, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	var data *T

	if err := collection.FindOne(ctx, query).Decode(&data); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("No document with Id " + id)
		}

		return nil, err
	}

	return data, nil
}

func UpdateService[T any](collection *mongo.Collection, ctx context.Context, id string, data interface{}) (*T, error) {
	var doc interface{} = &data

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	options := options.FindOneAndUpdate().SetReturnDocument(1)
	res := collection.FindOneAndUpdate(ctx, query, update, options)

	var updatedData *T

	if err := res.Decode(&updatedData); err != nil {
		return nil, errors.New("Data with Id " + id + " not exists")
	}

	return updatedData, nil
}

func DeleteService[T any](collection *mongo.Collection, ctx context.Context, id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := collection.DeleteOne(ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("No document with Id " + id + " exists")
	}

	return nil
}
