package question

import (
	"context"
	crud_services "main/crud/services"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	questionCollection      *mongo.Collection
	questionParam           crud_services.Param
	questionController      QuestionController
	questionRouteController QuestionRouteController
)

func Init(client *mongo.Client, ctx context.Context) {
	questionCollection = client.Database(os.Getenv("MONGO_DBNAME")).Collection("daftar_soal")
	questionParam = crud_services.Param{
		Collection:  questionCollection,
		Ctx:         ctx,
		ResultModel: QuestionDB{},
		CreateModel: CreateQuestion{},
		UpdateModel: UpdateQuestion{},
	}
	questionController = NewQuestionController(questionParam, "questions", "question")
	questionRouteController = NewQuestionRouteController(questionController)
}

func RouteInit(router *gin.RouterGroup) {
	questionRouteController.QuestionRoute(router)
	questionRouteController.AnswerRoute(router)
}
