package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getRecords(collection *mongo.Collection, ctx context.Context) (map[string]interface{}, error) {
	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var daftar_soal []bson.M

	for cur.Next(ctx) {
		var soal bson.M

		if err = cur.Decode(&soal); err != nil {
			return nil, err
		}

		daftar_soal = append(daftar_soal, soal)
	}

	res := map[string]interface{}{}
	res = map[string]interface{}{
		"data": daftar_soal,
	}

	return res, nil
}

func createRecord(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	req, err := collection.InsertOne(ctx, data)

	if err != nil {

		return nil, err

	}

	insertedId := req.InsertedID

	res := map[string]interface{}{
		"data": map[string]interface{}{
			"insertedId": insertedId,
		},
	}

	return res, nil
}
