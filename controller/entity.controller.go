package controller

import (
	"net/http"
	"payment/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntityController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewEntityController(servicesManager services.ServiceManager) *EntityController {
	return &EntityController{
		serviceManager: &servicesManager,
	}
}

// get list
func (handler *EntityController) GetListEntity(e *gin.Context) {
	entities, err := handler.serviceManager.EntityService.FindAllEntity(e)
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
	}

	e.JSON(http.StatusOK, entities)
}

func (handler *EntityController) GetEntityTypeByID(e *gin.Context) {
	id, _ := strconv.Atoi(e.Param("id"))
	entities, err := handler.serviceManager.EntityService.FindEntityById(e, int32(id))
	if err != nil {
		e.JSON(http.StatusInternalServerError, err)
		return
	}
	e.JSON(http.StatusOK, entities)
}
