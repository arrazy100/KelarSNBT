package question

import (
	"context"
	"main/crud_generics"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection              *mongo.Collection
	controller              crud_generics.CRUDControllerRepo[QuestionDB]
	questionController      QuestionController[QuestionDB]
	questionRouteController QuestionRouteController[QuestionDB]
)

func Init(client *mongo.Client, ctx context.Context) {
	collection = client.Database(os.Getenv("MONGO_DBNAME")).Collection("daftar_soal")
	controller = crud_generics.NewCRUDControllerRepo[QuestionDB](
		collection,
		ctx,
		"questions",
		"question",
		func() interface{} {
			return &CreateQuestion{}
		},
		func() interface{} {
			return &UpdateQuestion{}
		},
	)

	questionController = NewQuestionController(controller)
	questionRouteController = NewQuestionRouteController(questionController)
}

func RouteInit(router fiber.Router) {
	questionRouteController.QuestionRoute(router)
	questionRouteController.AnswerRoute(router)
}
