package test_generics

import (
	"main/crud_generics"

	"github.com/gofiber/fiber/v2"
)

type TestRouteController[T any] struct {
	testController TestController[T]
	genericRoute   crud_generics.CRUDRouteController[T]
}

func NewTestRouteController[T any](testController TestController[T], genericRoute crud_generics.CRUDRouteController[T]) TestRouteController[T] {
	return TestRouteController[T]{testController, genericRoute}
}

func (repo *TestRouteController[T]) Route(router fiber.Router) {
	repo.genericRoute.Route("/tests", router)

	router.Patch("/tests/setQuestions/:testId", repo.testController.SetData)
}
