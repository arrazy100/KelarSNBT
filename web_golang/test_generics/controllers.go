package test_generics

import (
	"main/crud_generics"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type TestController[T any] struct {
	genericController crud_generics.CRUDControllerRepo[T]
}

func NewTestController[T any](genericController crud_generics.CRUDControllerRepo[T]) TestController[T] {
	return TestController[T]{genericController}
}

func (repo *TestController[T]) SetData(ctx *fiber.Ctx) error {
	testId := ctx.Params("testId")

	var data *SetData
	if err := ctx.BodyParser(&data); err != nil {
		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

		return err
	}

	updateddData, err := SetDataService[T](repo.genericController.GetCollection(), repo.genericController.GetContext(), testId, data)

	if err != nil {
		if strings.Contains(err.Error(), "not exists") {
			ctx.SendStatus(http.StatusNotFound)
			ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

			return err
		}

		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

		return err
	}

	ctx.SendStatus(http.StatusOK)
	ctx.JSON(fiber.Map{"status": "success", "data": updateddData})

	return nil
}
