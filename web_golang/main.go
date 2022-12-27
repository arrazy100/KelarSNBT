package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	http.HandleFunc("/soal", requestHandler)
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNSTRING")))

	if err != nil {
		fmt.Println(err.Error())
	}

	collection := client.Database(os.Getenv("MONGO_DBNAME")).Collection("daftar_soal")

	data := map[string]interface{}{}

	err = json.NewDecoder(req.Body).Decode(&data)

	if err != nil {
		fmt.Println(err.Error())
	}

	switch req.Method {
	case "GET":
		response, err = getRecords(collection, ctx)
	case "POST":
		response, err = createRecord(collection, ctx, data)
	}

	if err != nil {
		response = map[string]interface{}{"error": err.Error}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}

}
