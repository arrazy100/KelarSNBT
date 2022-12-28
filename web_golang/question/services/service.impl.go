package question_services

import (
	"context"
	"errors"
	question_models "main/question/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuestionServiceImpl struct {
	questionCollection *mongo.Collection
	ctx                context.Context
}

func NewQuestionService(questionCollection *mongo.Collection, ctx context.Context) QuestionService {
	return &QuestionServiceImpl{questionCollection, ctx}
}

func (q *QuestionServiceImpl) Create(question *question_models.CreateQuestion) (*question_models.QuestionDB, error) {
	for i, attr := range question.Answers {
		if attr.Id.IsZero() {
			question.Answers[i].Id = primitive.NewObjectID()
		}
	}

	res, err := q.questionCollection.InsertOne(q.ctx, question)

	if err != nil {
		return nil, err
	}

	var newQuestion *question_models.QuestionDB
	query := bson.M{"_id": res.InsertedID}

	if err = q.questionCollection.FindOne(q.ctx, query).Decode(&newQuestion); err != nil {
		return nil, err
	}

	return newQuestion, nil
}

func (q *QuestionServiceImpl) FindAll(page int, limit int) ([]*question_models.QuestionDB, error) {
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

	cursor, err := q.questionCollection.Find(q.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(q.ctx)

	var questions []*question_models.QuestionDB

	for cursor.Next(q.ctx) {
		question := &question_models.QuestionDB{}
		err := cursor.Decode(question)

		if err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(questions) == 0 {
		return []*question_models.QuestionDB{}, nil
	}

	return questions, nil
}

func (q *QuestionServiceImpl) FindById(id string) (*question_models.QuestionDB, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	var question *question_models.QuestionDB

	if err := q.questionCollection.FindOne(q.ctx, query).Decode(&question); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("No document with Id " + id)
		}

		return nil, err
	}

	return question, nil
}

func (q *QuestionServiceImpl) Delete(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := q.questionCollection.DeleteOne(q.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("No document with Id " + id + " exists")
	}

	return nil
}

func (q *QuestionServiceImpl) Update(id string, data *question_models.UpdateQuestion) (*question_models.QuestionDB, error) {
	for i, attr := range data.Answers {
		if attr.Id.IsZero() {
			data.Answers[i].Id = primitive.NewObjectID()
		}
	}

	var doc interface{} = &data

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := q.questionCollection.FindOneAndUpdate(q.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedQuestion *question_models.QuestionDB

	if err := res.Decode(&updatedQuestion); err != nil {
		return nil, errors.New("Question with Id " + id + " not exists")
	}

	return updatedQuestion, nil
}

func (q *QuestionServiceImpl) CreateAnswer(id string, data *question_models.CreateAnswer) (*question_models.QuestionDB, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	createData := question_models.AnswerDB{
		Id:            primitive.NewObjectID(),
		Content:       data.Content,
		CorrectAnswer: data.CorrectAnswer,
	}

	createFields := bson.M{
		"answers": createData,
	}

	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$push", Value: createFields}}

	res := q.questionCollection.FindOneAndUpdate(q.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1).SetUpsert(true))

	var updatedQuestion *question_models.QuestionDB

	if err := res.Decode(&updatedQuestion); err != nil {
		return nil, errors.New("Question with Id " + id + " not exists")
	}

	return updatedQuestion, nil
}

func (q *QuestionServiceImpl) UpdateAnswer(id string, data *question_models.UpdateAnswer) (*question_models.QuestionDB, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	updateFields := bson.M{
		"answers.$[e].content":        &data.Content,
		"answers.$[e].correct_answer": &data.CorrectAnswer,
	}

	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: updateFields}}
	arrayFiltersOpt := options.ArrayFilters{Filters: []interface{}{bson.D{{Key: "e._id", Value: &data.Id}}}}
	res := q.questionCollection.FindOneAndUpdate(q.ctx, query, update, options.FindOneAndUpdate().SetArrayFilters(arrayFiltersOpt).SetReturnDocument(1))

	var updatedQuestion *question_models.QuestionDB

	if err := res.Decode(&updatedQuestion); err != nil {
		return nil, errors.New("Question with Id " + id + " not exists")
	}

	return updatedQuestion, nil
}

func (q *QuestionServiceImpl) DeleteAnswer(id string) error {
	answerId, _ := primitive.ObjectIDFromHex(id)

	deleteFields := bson.M{
		"answers": bson.M{
			"_id": answerId,
		},
	}

	query := bson.M{}
	update := bson.M{"$pull": deleteFields}
	_, err := q.questionCollection.UpdateOne(q.ctx, query, update)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
