package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BusinessUnitController interface {
	FindBusinessUnits(c *gin.Context)
	FindBusinessUnitById(c *gin.Context)
	FindExcBusinessUnit(c *gin.Context)
	InsertBusinessUnit(c *gin.Context)
	UpdateBusinessUnit(c *gin.Context)
	DeleteBusinessUnit(c *gin.Context)
}

type businessUnitController struct {
	businessUnitService service.BusinessUnitService
	jwtService          service.JWTService
}

func NewBusinessUnitController(businessUnitServ service.BusinessUnitService, jwtServ service.JWTService) BusinessUnitController {
	return &businessUnitController{
		businessUnitService: businessUnitServ,
		jwtService:          jwtServ,
	}
}

func (b *businessUnitController) FindBusinessUnits(c *gin.Context) {
	var (
		businessUnits []model.BusinessUnit
		response      helper.Response
	)
	businessUnits, err := b.businessUnitService.FindBusinessUnits()

	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		response = helper.BuildResponse(true, "OK", businessUnits)
		c.JSON(http.StatusOK, response)
	}
}

func (b *businessUnitController) FindBusinessUnitById(c *gin.Context) {
	var (
		businessUnit model.BusinessUnit
		response     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		businessUnit, err = b.businessUnitService.FindBusinessUnitById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", businessUnit)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *businessUnitController) FindExcBusinessUnit(c *gin.Context) {
	var (
		businessUnits []model.BusinessUnit
		response      helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		businessUnits, err = b.businessUnitService.FindExcBusinessUnit(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", businessUnits)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *businessUnitController) InsertBusinessUnit(c *gin.Context) {
	var (
		businessUnit                model.BusinessUnit
		response                    helper.Response
		CreateBusinessUnitParameter model.CreateBusinessUnitParameter
	)
	err := c.ShouldBindJSON(&CreateBusinessUnitParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		businessUnit, err = b.businessUnitService.InsertBusinessUnit(CreateBusinessUnitParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register business unit", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", businessUnit)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *businessUnitController) UpdateBusinessUnit(c *gin.Context) {
	var (
		newData  model.BusinessUnit
		oldData  model.BusinessUnit
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		err = c.ShouldBindJSON(&newData)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			oldData, err = b.businessUnitService.FindBusinessUnitById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.BusinessUnit{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				businessUnit, err := b.businessUnitService.UpdateBusinessUnit(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update business unit", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", businessUnit)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *businessUnitController) DeleteBusinessUnit(c *gin.Context) {
	var (
		newData  model.BusinessUnit
		oldData  model.BusinessUnit
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		err = c.ShouldBindJSON(&newData)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			oldData, err = b.businessUnitService.FindBusinessUnitById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.BusinessUnit{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				businessUnit, err := b.businessUnitService.DeleteBusinessUnit(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete business unit", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", businessUnit)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
