package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DivisionController interface {
	FindDivisions(c *gin.Context)
	FindDivisionById(c *gin.Context)
	FindExcDivision(c *gin.Context)
	InsertDivision(c *gin.Context)
	UpdateDivision(c *gin.Context)
	DeleteDivision(c *gin.Context)
}

type divisionController struct {
	divisionService service.DivisionService
	jwtService      service.JWTService
}

func NewDivisionController(divisionServ service.DivisionService, jwtServ service.JWTService) DivisionController {
	return &divisionController{
		divisionService: divisionServ,
		jwtService:      jwtServ,
	}
}

func (b *divisionController) FindDivisions(c *gin.Context) {
	var (
		divisions []model.Division
		response  helper.Response
	)
	divisions, err := b.divisionService.FindDivisions()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", divisions)
		c.JSON(http.StatusOK, response)
	}
}

func (b *divisionController) FindDivisionById(c *gin.Context) {
	var (
		division model.Division
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	division, err = b.divisionService.FindDivisionById(uint(id))
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", division)
		c.JSON(http.StatusOK, response)
	}
}

func (b *divisionController) FindExcDivision(c *gin.Context) {
	var (
		divisions []model.Division
		response  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	divisions, err = b.divisionService.FindExcDivision(uint(id))
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", divisions)
		c.JSON(http.StatusOK, response)
	}
}

func (b *divisionController) InsertDivision(c *gin.Context) {
	var (
		division                model.Division
		response                helper.Response
		CreateDivisionParameter model.CreateDivisionParameter
	)
	err := c.ShouldBindJSON(&CreateDivisionParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		division, err = b.divisionService.InsertDivision(CreateDivisionParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register division", err.Error(), nil)
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", division)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *divisionController) UpdateDivision(c *gin.Context) {
	var (
		newData  model.Division
		oldData  model.Division
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

	oldData, err = b.divisionService.FindDivisionById(uint(id))
	if (oldData == model.Division{}) {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		division, err := b.divisionService.UpdateDivision(newData, uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Failed to update division", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", division)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *divisionController) DeleteDivision(c *gin.Context) {
	var (
		newData  model.Division
		oldData  model.Division
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

	oldData, err = b.divisionService.FindDivisionById(uint(id))
	if (oldData == model.Division{}) {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		division, err := b.divisionService.DeleteDivision(newData, uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Failed to delete division", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", division)
			c.JSON(http.StatusOK, response)
		}
	}
}
