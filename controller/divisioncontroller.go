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

	var divisions []model.Division = b.divisionService.FindDivisions()
	res := helper.BuildResponse(true, "OK", divisions)

	c.JSON(http.StatusOK, res)
}

func (b *divisionController) FindDivisionById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var division model.Division = b.divisionService.FindDivisionById(uint(id))
	if (division == model.Division{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", division)
		c.JSON(http.StatusOK, res)
	}
}

func (b *divisionController) FindExcDivision(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var divisions []model.Division = b.divisionService.FindExcDivision(uint(id))
	res := helper.BuildResponse(true, "OK", divisions)
	c.JSON(http.StatusOK, res)
}

func (b *divisionController) InsertDivision(c *gin.Context) {
	var input model.CreateDivisionParameter
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		result := b.divisionService.InsertDivision(input)
		response := helper.BuildResponse(true, "OK", result)
		c.JSON(http.StatusOK, response)
	}
}

func (b *divisionController) UpdateDivision(c *gin.Context) {
	var newData model.Division
	var oldData model.Division

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	oldData = b.divisionService.FindDivisionById(uint(id))
	if (oldData == model.Division{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		if err := c.ShouldBindJSON(&newData); err != nil {
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			result := b.divisionService.UpdateDivision(newData, uint(id))
			response := helper.BuildResponse(true, "OK", result)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *divisionController) DeleteDivision(c *gin.Context) {

	var newData model.Division
	var oldData model.Division

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	oldData = b.divisionService.FindDivisionById(uint(id))
	if (oldData == model.Division{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		if err := c.ShouldBindJSON(&newData); err != nil {
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			result := b.divisionService.DeleteDivision(newData, uint(id))
			response := helper.BuildResponse(true, "OK", result)
			c.JSON(http.StatusOK, response)
		}
	}
}
