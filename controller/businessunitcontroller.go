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

	var businessUnits []model.BusinessUnit = b.businessUnitService.FindBusinessUnits()
	res := helper.BuildResponse(true, "OK", businessUnits)

	c.JSON(http.StatusOK, res)
}

func (b *businessUnitController) FindBusinessUnitById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var businessUnit model.BusinessUnit = b.businessUnitService.FindBusinessUnitById(uint(id))
	if (businessUnit == model.BusinessUnit{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", businessUnit)
		c.JSON(http.StatusOK, res)
	}
}

func (b *businessUnitController) InsertBusinessUnit(c *gin.Context) {
	var input model.CreateBusinessUnitParameter
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		result := b.businessUnitService.InsertBusinessUnit(input)
		response := helper.BuildResponse(true, "OK", result)
		c.JSON(http.StatusOK, response)
	}
}

func (b *businessUnitController) UpdateBusinessUnit(c *gin.Context) {
	var newData model.BusinessUnit
	var oldData model.BusinessUnit

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	oldData = b.businessUnitService.FindBusinessUnitById(uint(id))
	if (oldData == model.BusinessUnit{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		if err := c.ShouldBindJSON(&newData); err != nil {
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
			// return
		} else {
			result := b.businessUnitService.UpdateBusinessUnit(newData)
			response := helper.BuildResponse(true, "OK", result)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *businessUnitController) DeleteBusinessUnit(c *gin.Context) {
	var businessUnit model.BusinessUnit

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userId, err := strconv.ParseUint(c.Param("userId"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param userId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	businessUnit = b.businessUnitService.FindBusinessUnitById(uint(id))
	if (businessUnit == model.BusinessUnit{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		result := b.businessUnitService.DeleteBusinessUnit(businessUnit, uint(userId))
		response := helper.BuildResponse(true, "OK", result)
		c.JSON(http.StatusOK, response)
	}
}
