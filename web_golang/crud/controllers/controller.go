package crud_controllers

import (
	crud_services "main/crud/services"
	"main/logs"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CRUDController struct {
	service    crud_services.CRUDService
	name       string
	singleName string
}

func NewCRUDController(service crud_services.CRUDService, name string, singleName string) CRUDController {
	return CRUDController{service, name, singleName}
}

func (c *CRUDController) Create(ctx *gin.Context) {
	logs.Info("User requesting to create a " + c.singleName)

	modelType := c.service.GetCreateModel()
	data := reflect.New(modelType).Interface()

	logs.Debug("Validating " + c.singleName + " from JSON")
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		logs.Warning(err.Error())

		return
	}

	newData, err := c.service.Create(data)
	logs.DebugObject("Inserting new "+c.singleName+" to database.", data)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return
	}
	logs.DebugObject("New "+c.singleName+" saved.", data)

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newData})
}

func (c *CRUDController) Update(ctx *gin.Context) {
	logs.Info("User requesting to create a " + c.singleName)

	objectId := ctx.Param(c.GetIdName())

	modelType := c.service.GetUpdateModel()
	data := reflect.New(modelType).Interface()

	logs.Debug("Validating " + c.singleName + " from JSON")
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		logs.Warning(err.Error())

		return
	}

	updatedData, err := c.service.Update(objectId, data)
	logs.DebugObject("Updating a "+c.singleName+" from database.", data)

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
	logs.DebugObject("A "+c.singleName+" is updated.", data)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedData})

}

func (c *CRUDController) FindAll(ctx *gin.Context) {
	logs.Info("Requesting all " + c.name)

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

	tasks, err := c.service.FindAll(intPage, intLimit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		logs.Error(err.Error())

		return
	}
	logs.Debug("Getting all " + c.name + " from db")

	len_tasks := reflect.ValueOf(tasks).Len()

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len_tasks, "data": tasks})
}

func (c *CRUDController) FindById(ctx *gin.Context) {
	logs.Info("Requsting a " + c.singleName + "  by id")
	objectId := ctx.Param(c.GetIdName())

	data, err := c.service.FindById(objectId)
	logs.Debug("Getting a " + c.singleName + " from db by id")

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
	logs.Debug("Success finding a " + c.singleName)
}

func (c *CRUDController) Delete(ctx *gin.Context) {
	logs.Info("Requsting to delete a " + c.singleName + " by id")
	objectId := ctx.Param(c.GetIdName())

	err := c.service.Delete(objectId)
	logs.Debug("Deleting a " + c.singleName + " by id")

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
	logs.Debug("Success deleting a " + c.singleName)
}

func (c *CRUDController) GetName() string {
	return c.name
}

func (c *CRUDController) GetSingleName() string {
	return c.singleName
}

func (c *CRUDController) GetIdName() string {
	return c.singleName + "Id"
}
