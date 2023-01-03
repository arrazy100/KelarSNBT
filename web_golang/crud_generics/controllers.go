package crud_generics

import (
	"context"
	"errors"
	"main/logs"
	"net/http"
	"strconv"
	"strings"

	"main/common"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type CRUDControllerRepo[T any] struct {
	collection             *mongo.Collection
	ctx                    context.Context
	name                   string
	singleName             string
	createModelConstructor func() interface{}
	updateModelConstructor func() interface{}
}

func NewCRUDControllerRepo[T any](collection *mongo.Collection, ctx context.Context, name string, singleName string, createModelConstructor func() interface{}, updateModelConstructor func() interface{}) CRUDControllerRepo[T] {
	return CRUDControllerRepo[T]{
		collection:             collection,
		ctx:                    ctx,
		name:                   name,
		singleName:             singleName,
		createModelConstructor: createModelConstructor,
		updateModelConstructor: updateModelConstructor,
	}
}

func (repo *CRUDControllerRepo[T]) FindAll(ctx *fiber.Ctx) error {
	logs.Info("Requesting all " + repo.name)

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return err
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return err
	}

	datas, err := FindAllService[T](repo.collection, repo.ctx, intPage, intLimit)
	if err != nil {
		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return err
	}
	logs.Debug("Getting all " + repo.name + " from db")

	ctx.SendStatus(http.StatusOK)
	ctx.JSON(fiber.Map{"status": "success", "results": len(datas), "data": datas})

	return nil
}

func (repo *CRUDControllerRepo[T]) Create(ctx *fiber.Ctx) error {
	logs.Info("User requesting to create a " + repo.singleName)

	data := repo.createModelConstructor()

	logs.Debug("Parsing " + repo.singleName + " from JSON")
	if err := ctx.BodyParser(&data); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		logs.Warning(err.Error())

		return err
	}

	logs.Debug("Validating " + repo.singleName + " from JSON")
	errs := common.Validate(data)
	if errs != nil {
		ctx.Status(http.StatusBadRequest).JSON(errs)

		encoded, _ := ctx.App().Config().JSONEncoder(errs)
		logs.Error("Failed to validate " + repo.singleName + " from JSON. " + string(encoded))

		return errors.New(string(encoded))
	}

	newData, err := CreateService[T](repo.collection, repo.ctx, data)
	logs.DebugObject("Inserting new "+repo.singleName+" to database.", data)

	if err != nil {
		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})

		return err
	}
	logs.DebugObject("New "+repo.singleName+" saved.", data)

	ctx.SendStatus(http.StatusCreated)
	ctx.JSON(fiber.Map{"status": "success", "data": newData})

	return nil
}

func (repo *CRUDControllerRepo[T]) FindById(ctx *fiber.Ctx) error {
	logs.Info("Requsting a " + repo.singleName + "  by id")
	objectId := ctx.Params(repo.singleName + "Id")

	data, err := FindByIdService[T](repo.collection, repo.ctx, objectId)
	logs.Debug("Getting a " + repo.singleName + " from db by id")

	if err != nil {
		if strings.Contains(err.Error(), "No document") {
			ctx.SendStatus(http.StatusNotFound)
			ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
			logs.Error(err.Error())

			return err
		}

		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return err
	}

	ctx.SendStatus(http.StatusOK)
	ctx.JSON(fiber.Map{"status": "success", "data": data})
	logs.Debug("Success finding a " + repo.singleName)

	return nil
}

func (repo *CRUDControllerRepo[T]) Update(ctx *fiber.Ctx) error {
	logs.Info("User requesting to update a " + repo.singleName)

	objectId := ctx.Params(repo.singleName + "Id")

	data := repo.updateModelConstructor()

	logs.Debug("Parsing " + repo.singleName + " from JSON")
	if err := ctx.BodyParser(&data); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		logs.Warning(err.Error())

		return err
	}

	logs.Debug("Validating " + repo.singleName + " from JSON")
	errs := common.Validate(data)
	if errs != nil {
		ctx.Status(http.StatusBadRequest).JSON(errs)

		encoded, _ := ctx.App().Config().JSONEncoder(errs)
		logs.Error("Failed to validate " + repo.singleName + " from JSON. " + string(encoded))

		return errors.New(string(encoded))
	}

	updatedData, err := UpdateService[T](repo.collection, repo.ctx, objectId, data)
	logs.DebugObject("Updating a "+repo.singleName+" from database.", data)

	if err != nil {
		if strings.Contains(err.Error(), "not exists") {
			ctx.SendStatus(http.StatusNotFound)
			ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
			logs.Error(err.Error())

			return err
		}

		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return err
	}
	logs.DebugObject("A "+repo.singleName+" is updated.", data)

	ctx.SendStatus(http.StatusOK)
	ctx.JSON(fiber.Map{"status": "success", "data": updatedData})

	return nil
}

func (repo *CRUDControllerRepo[T]) Delete(ctx *fiber.Ctx) error {
	logs.Info("Requsting to delete a " + repo.singleName + " by id")
	objectId := ctx.Params(repo.singleName + "Id")

	err := DeleteService[T](repo.collection, repo.ctx, objectId)
	logs.Debug("Deleting a " + repo.singleName + " by id")

	if err != nil {
		if strings.Contains(err.Error(), "exists") {
			ctx.SendStatus(http.StatusNotFound)
			ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
			logs.Error(err.Error())

			return err
		}

		ctx.SendStatus(http.StatusBadGateway)
		ctx.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return err
	}

	ctx.SendStatus(http.StatusNoContent)
	logs.Debug("Success deleting a " + repo.singleName)

	return nil
}

func (repo *CRUDControllerRepo[T]) GetCollection() *mongo.Collection {
	return repo.collection
}

func (repo *CRUDControllerRepo[T]) GetContext() context.Context {
	return repo.ctx
}

func (repo *CRUDControllerRepo[T]) GetSingleName() string {
	return repo.singleName
}
