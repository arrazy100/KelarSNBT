package question

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateAnswerService(collection *mongo.Collection, ctx context.Context, id string, data *CreateAnswer) (*QuestionDB, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	createData := AnswerDB{
		Id:            primitive.NewObjectID(),
		Content:       data.Content,
		CorrectAnswer: data.CorrectAnswer,
	}

	createFields := bson.M{
		"answers": createData,
	}

	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$push", Value: createFields}}

	res := collection.FindOneAndUpdate(ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1).SetUpsert(true))

	var updatedQuestion *QuestionDB

	if err := res.Decode(&updatedQuestion); err != nil {
		return nil, errors.New("Question with Id " + id + " not exists")
	}

	return updatedQuestion, nil
}

func UpdateAnswerService(collection *mongo.Collection, ctx context.Context, id string, data *UpdateAnswer) (*QuestionDB, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	updateFields := bson.M{
		"answers.$[e].content":        &data.Content,
		"answers.$[e].correct_answer": &data.CorrectAnswer,
	}

	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: updateFields}}
	arrayFiltersOpt := options.ArrayFilters{Filters: []interface{}{bson.D{{Key: "e._id", Value: &data.Id}}}}
	res := collection.FindOneAndUpdate(ctx, query, update, options.FindOneAndUpdate().SetArrayFilters(arrayFiltersOpt).SetReturnDocument(1))

	var updatedQuestion *QuestionDB

	if err := res.Decode(&updatedQuestion); err != nil {
		return nil, errors.New("Question with Id " + id + " not exists")
	}

	return updatedQuestion, nil
}

func DeleteAnswerService(collection *mongo.Collection, ctx context.Context, id string) error {
	answerId, _ := primitive.ObjectIDFromHex(id)

	deleteFields := bson.M{
		"answers": bson.M{
			"_id": answerId,
		},
	}

	query := bson.M{}
	update := bson.M{"$pull": deleteFields}
	_, err := collection.UpdateOne(ctx, query, update)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
