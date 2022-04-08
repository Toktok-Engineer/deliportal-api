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
	var response helper.Response

	businessUnits, err := b.businessUnitService.FindBusinessUnits()

	if err != nil {
		response = helper.BuildErrorResponse("Failed to register user", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		response = helper.BuildResponse(true, "OK", businessUnits)
		c.JSON(http.StatusOK, response)
	}
}

func (b *businessUnitController) FindBusinessUnitById(c *gin.Context) {
	var response helper.Response

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		businessUnit, err := b.businessUnitService.FindBusinessUnitById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", businessUnit)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *businessUnitController) InsertBusinessUnit(c *gin.Context) {
	var response helper.Response

	var input model.CreateBusinessUnitParameter
	if err := c.ShouldBindJSON(&input); err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		businessUnit, err := b.businessUnitService.InsertBusinessUnit(input)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", businessUnit)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *businessUnitController) UpdateBusinessUnit(c *gin.Context) {
	var response helper.Response
	var newData model.BusinessUnit
	var oldData model.BusinessUnit

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		oldData, err = b.businessUnitService.FindBusinessUnitById(uint(id))
		if err != nil {
			res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, res)
		} else if (oldData == model.BusinessUnit{}) {
			res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
			c.JSON(http.StatusNotFound, res)
		} else {
			if err = c.ShouldBindJSON(&newData); err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusBadRequest, response)
			} else {
				businessUnit, err := b.businessUnitService.UpdateBusinessUnit(newData)
				if err != nil {
					response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
					c.JSON(http.StatusNotFound, response)
				} else {
					response = helper.BuildResponse(true, "OK", businessUnit)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *businessUnitController) DeleteBusinessUnit(c *gin.Context) {
	var response helper.Response

	var businessUnit model.BusinessUnit

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		userId, err := strconv.ParseUint(c.Param("userId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param userId was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			businessUnit, err = b.businessUnitService.FindBusinessUnitById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (businessUnit == model.BusinessUnit{}) {
				response = helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				result, err := b.businessUnitService.DeleteBusinessUnit(businessUnit, uint(userId))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
					c.JSON(http.StatusNotFound, response)
				} else {
					response := helper.BuildResponse(true, "OK", result)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
