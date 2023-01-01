package crud_generics

import (
	"context"
	"main/logs"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

func (repo *CRUDControllerRepo[T]) FindAll(ctx *gin.Context) {
	logs.Info("Requesting all " + repo.name)

	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return
	}

	datas, err := FindAllService[T](repo.collection, repo.ctx, intPage, intLimit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return
	}
	logs.Debug("Getting all " + repo.name + " from db")

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(datas), "data": datas})
}

func (repo *CRUDControllerRepo[T]) Create(ctx *gin.Context) {
	logs.Info("User requesting to create a " + repo.singleName)

	data := repo.createModelConstructor()

	logs.Debug("Validating " + repo.singleName + " from JSON")
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		logs.Warning(err.Error())

		return
	}

	newData, err := CreateService[T](repo.collection, repo.ctx, data)
	logs.DebugObject("Inserting new "+repo.singleName+" to database.", data)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return
	}
	logs.DebugObject("New "+repo.singleName+" saved.", data)

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newData})
}

func (repo *CRUDControllerRepo[T]) FindById(ctx *gin.Context) {
	logs.Info("Requsting a " + repo.singleName + "  by id")
	objectId := ctx.Param(repo.singleName + "Id")

	data, err := FindByIdService[T](repo.collection, repo.ctx, objectId)
	logs.Debug("Getting a " + repo.singleName + " from db by id")

	if err != nil {
		if strings.Contains(err.Error(), "No document") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			logs.Error(err.Error())

			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
	logs.Debug("Success finding a " + repo.singleName)
}

func (repo *CRUDControllerRepo[T]) Update(ctx *gin.Context) {
	logs.Info("User requesting to create a " + repo.singleName)

	objectId := ctx.Param(repo.singleName + "Id")

	data := repo.updateModelConstructor()

	logs.Debug("Validating " + repo.singleName + " from JSON")
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		logs.Warning(err.Error())

		return
	}

	updatedData, err := UpdateService[T](repo.collection, repo.ctx, objectId, data)
	logs.DebugObject("Updating a "+repo.singleName+" from database.", data)

	if err != nil {
		if strings.Contains(err.Error(), "not exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			logs.Error(err.Error())

			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return
	}
	logs.DebugObject("A "+repo.singleName+" is updated.", data)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedData})
}

func (repo *CRUDControllerRepo[T]) Delete(ctx *gin.Context) {
	logs.Info("Requsting to delete a " + repo.singleName + " by id")
	objectId := ctx.Param(repo.singleName + "Id")

	err := DeleteService[T](repo.collection, repo.ctx, objectId)
	logs.Debug("Deleting a " + repo.singleName + " by id")

	if err != nil {
		if strings.Contains(err.Error(), "exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			logs.Error(err.Error())

			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return
	}

	ctx.JSON(http.StatusNoContent, nil)
	logs.Debug("Success deleting a " + repo.singleName)
}

func (repo *CRUDControllerRepo[T]) GetCollection() *mongo.Collection {
	return repo.collection
}

func (repo *CRUDControllerRepo[T]) GetContext() context.Context {
	return repo.ctx
}
