package crud_generics

import (
	"github.com/gofiber/fiber/v2"
)

type CRUDRouteController[T any] struct {
	controllerRepo CRUDControllerRepo[T]
}

func NewCRUDRouteController[T any](crudController CRUDControllerRepo[T]) CRUDRouteController[T] {
	return CRUDRouteController[T]{crudController}
}

func (repo *CRUDRouteController[T]) Route(urlPath string, cr fiber.Router) {
	router := cr.Group(urlPath)

	router.Get("/", repo.controllerRepo.FindAll)
	router.Post("/", repo.controllerRepo.Create)
	router.Get("/:"+repo.controllerRepo.singleName+"Id", repo.controllerRepo.FindById)
	router.Patch("/:"+repo.controllerRepo.singleName+"Id", repo.controllerRepo.Update)
	router.Delete("/:"+repo.controllerRepo.singleName+"Id", repo.controllerRepo.Delete)
}
