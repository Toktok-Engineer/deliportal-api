package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DepartmentController interface {
	FindDepartments(c *gin.Context)
	FindDepartmentById(c *gin.Context)
	FindExcDepartment(c *gin.Context)
	FindDepartmentByDivId(c *gin.Context)
	InsertDepartment(c *gin.Context)
	UpdateDepartment(c *gin.Context)
	DeleteDepartment(c *gin.Context)
}

type departmentController struct {
	departmentService service.DepartmentService
	jwtService        service.JWTService
}

func NewDepartmentController(departmentServ service.DepartmentService, jwtServ service.JWTService) DepartmentController {
	return &departmentController{
		departmentService: departmentServ,
		jwtService:        jwtServ,
	}
}

func (b *departmentController) FindDepartments(c *gin.Context) {

	var departments []model.SelectDepartmentParameter = b.departmentService.FindDepartments()
	res := helper.BuildResponse(true, "OK", departments)

	c.JSON(http.StatusOK, res)
}

func (b *departmentController) FindDepartmentById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var department model.SelectDepartmentParameter = b.departmentService.FindDepartmentById(uint(id))
	if (department == model.SelectDepartmentParameter{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", department)
		c.JSON(http.StatusOK, res)
	}
}

func (b *departmentController) FindExcDepartment(c *gin.Context) {
	divId, err := strconv.ParseUint(c.Param("divId"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param divId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var departments []model.SelectDepartmentParameter = b.departmentService.FindExcDepartment(uint(divId), uint(id))
	res := helper.BuildResponse(true, "OK", departments)
	c.JSON(http.StatusOK, res)
}

func (b *departmentController) FindDepartmentByDivId(c *gin.Context) {
	divId, err := strconv.ParseUint(c.Param("divId"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var departments []model.SelectDepartmentParameter = b.departmentService.FindDepartmentByDivId(uint(divId))
	res := helper.BuildResponse(true, "OK", departments)
	c.JSON(http.StatusOK, res)
}

func (b *departmentController) InsertDepartment(c *gin.Context) {
	var input model.CreateDepartmentParameter
	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		result := b.departmentService.InsertDepartment(input)
		response := helper.BuildResponse(true, "OK", result)
		c.JSON(http.StatusOK, response)
	}
}

func (b *departmentController) UpdateDepartment(c *gin.Context) {
	var newData model.Department
	var oldData model.SelectDepartmentParameter

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	oldData = b.departmentService.FindDepartmentById(uint(id))
	if (oldData == model.SelectDepartmentParameter{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		if err := c.ShouldBindJSON(&newData); err != nil {
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			result := b.departmentService.UpdateDepartment(newData, uint(id))
			response := helper.BuildResponse(true, "OK", result)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *departmentController) DeleteDepartment(c *gin.Context) {
	var newData model.Department
	var oldData model.SelectDepartmentParameter

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	oldData = b.departmentService.FindDepartmentById(uint(id))
	if (oldData == model.SelectDepartmentParameter{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		if err := c.ShouldBindJSON(&newData); err != nil {
			response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			result := b.departmentService.DeleteDepartment(newData, uint(id))
			response := helper.BuildResponse(true, "OK", result)
			c.JSON(http.StatusOK, response)
		}
	}
}
