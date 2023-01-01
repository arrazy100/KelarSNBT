package crud_generics

import (
	"github.com/gin-gonic/gin"
)

type CRUDRouteController[T any] struct {
	controllerRepo CRUDControllerRepo[T]
}

func NewCRUDRouteController[T any](crudController CRUDControllerRepo[T]) CRUDRouteController[T] {
	return CRUDRouteController[T]{crudController}
}

func (repo *CRUDRouteController[T]) Route(urlPath string, cr *gin.RouterGroup) {
	router := cr.Group(urlPath)

	router.GET("/", repo.controllerRepo.FindAll)
	router.POST("/", repo.controllerRepo.Create)
	router.GET("/:"+repo.controllerRepo.singleName+"Id", repo.controllerRepo.FindById)
	router.PATCH("/:"+repo.controllerRepo.singleName+"Id", repo.controllerRepo.Update)
	router.DELETE("/:"+repo.controllerRepo.singleName+"Id", repo.controllerRepo.Delete)
}
