package test_generics

import (
	"main/crud_generics"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type TestController[T any] struct {
	genericController crud_generics.CRUDControllerRepo[T]
}

func NewTestController[T any](genericController crud_generics.CRUDControllerRepo[T]) TestController[T] {
	return TestController[T]{genericController}
}

func (repo *TestController[T]) SetData(ctx *gin.Context) {
	testId := ctx.Param("testId")

	var data *SetData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	updateddData, err := SetDataService[T](repo.genericController.GetCollection(), repo.genericController.GetContext(), testId, data)

	if err != nil {
		if strings.Contains(err.Error(), "not exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})

			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updateddData})
}
