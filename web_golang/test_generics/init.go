package test_generics

import (
	"context"
	"main/crud_generics"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection          *mongo.Collection
	controller          crud_generics.CRUDControllerRepo[TestDB]
	routeController     crud_generics.CRUDRouteController[TestDB]
	testController      TestController[TestDB]
	testRouteController TestRouteController[TestDB]
)

func Init(client *mongo.Client, ctx context.Context) {
	collection = client.Database(os.Getenv("MONGO_DBNAME")).Collection("daftar_tes")
	controller = crud_generics.NewCRUDControllerRepo[TestDB](
		collection,
		ctx,
		"tests",
		"test",
		func() interface{} {
			return &CreateTest{}
		},
		func() interface{} {
			return &UpdateTest{}
		},
	)
	routeController = crud_generics.NewCRUDRouteController(controller)

	testController = NewTestController(controller)
	testRouteController = NewTestRouteController(testController, routeController)
}

func RouteInit(router fiber.Router) {
	testRouteController.Route(router)
}
