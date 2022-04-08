package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PositionController interface {
	FindPositions(c *gin.Context)
	FindPositionById(c *gin.Context)
	FindExcPosition(c *gin.Context)
	InsertPosition(c *gin.Context)
	UpdatePosition(c *gin.Context)
	DeletePosition(c *gin.Context)
}

type positionController struct {
	positionService service.PositionService
	jwtService      service.JWTService
}

func NewPositionController(positionServ service.PositionService, jwtServ service.JWTService) PositionController {
	return &positionController{
		positionService: positionServ,
		jwtService:      jwtServ,
	}
}

func (b *positionController) FindPositions(c *gin.Context) {
	var (
		positions []model.Position
		response  helper.Response
	)
	positions, err := b.positionService.FindPositions()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", positions)
		c.JSON(http.StatusOK, response)
	}
}

func (b *positionController) FindPositionById(c *gin.Context) {
	var (
		position model.Position
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	position, err = b.positionService.FindPositionById(uint(id))
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", position)
		c.JSON(http.StatusOK, response)
	}
}

func (b *positionController) FindExcPosition(c *gin.Context) {
	var (
		positions []model.Position
		response  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	positions, err = b.positionService.FindExcPosition(uint(id))
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", positions)
		c.JSON(http.StatusOK, response)
	}
}

func (b *positionController) InsertPosition(c *gin.Context) {
	var (
		position                model.Position
		response                helper.Response
		CreatePositionParameter model.CreatePositionParameter
	)
	err := c.ShouldBindJSON(&CreatePositionParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		position, err = b.positionService.InsertPosition(CreatePositionParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register position", err.Error(), nil)
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", position)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *positionController) UpdatePosition(c *gin.Context) {
	var (
		newData  model.Position
		oldData  model.Position
		response helper.Response
	)

	err := c.ShouldBindJSON(&newData)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	oldData, err = b.positionService.FindPositionById(uint(id))
	if (oldData == model.Position{}) {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		position, err := b.positionService.UpdatePosition(newData, uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Failed to update position", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", position)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *positionController) DeletePosition(c *gin.Context) {
	var (
		newData  model.Position
		oldData  model.Position
		response helper.Response
	)

	err := c.ShouldBindJSON(&newData)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	oldData, err = b.positionService.FindPositionById(uint(id))
	if (oldData == model.Position{}) {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		position, err := b.positionService.DeletePosition(newData, uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Failed to delete position", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", position)
			c.JSON(http.StatusOK, response)
		}
	}
}
