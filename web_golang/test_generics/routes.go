package test_generics

import (
	"main/crud_generics"

	"github.com/gin-gonic/gin"
)

type TestRouteController[T any] struct {
	testController TestController[T]
	genericRoute   crud_generics.CRUDRouteController[T]
}

func NewTestRouteController[T any](testController TestController[T], genericRoute crud_generics.CRUDRouteController[T]) TestRouteController[T] {
	return TestRouteController[T]{testController, genericRoute}
}

func (repo *TestRouteController[T]) Route(router *gin.RouterGroup) {
	repo.genericRoute.Route("/tests", router)

	router.PATCH("/tests/setQuestions/:testId", repo.testController.SetData)
}
