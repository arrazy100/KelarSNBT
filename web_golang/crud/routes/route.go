package crud_routes

import (
	crud_controllers "main/crud/controllers"

	"github.com/gin-gonic/gin"
)

type CRUDRouteController struct {
	crudController crud_controllers.CRUDController
}

func NewCRUDRouteController(crudController crud_controllers.CRUDController) CRUDRouteController {
	return CRUDRouteController{crudController}
}

func (crc *CRUDRouteController) CRUDRoute(urlPath string, cr *gin.RouterGroup) {
	router := cr.Group(urlPath)

	router.GET("/", crc.crudController.FindAll)
	router.POST("/", crc.crudController.Create)
	router.PATCH("/:"+crc.crudController.GetIdName(), crc.crudController.Update)
	router.GET("/:"+crc.crudController.GetIdName(), crc.crudController.FindById)
	router.DELETE("/:"+crc.crudController.GetIdName(), crc.crudController.Delete)
}
