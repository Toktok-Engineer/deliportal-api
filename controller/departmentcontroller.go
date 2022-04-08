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
	var (
		departments []model.SelectDepartmentParameter
		response    helper.Response
	)
	departments, err := b.departmentService.FindDepartments()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", departments)
		c.JSON(http.StatusOK, response)
	}
}

func (b *departmentController) FindDepartmentById(c *gin.Context) {
	var (
		department model.SelectDepartmentParameter
		response   helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	department, err = b.departmentService.FindDepartmentById(uint(id))
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", department)
		c.JSON(http.StatusOK, response)
	}
}

func (b *departmentController) FindExcDepartment(c *gin.Context) {
	var (
		departments []model.SelectDepartmentParameter
		response    helper.Response
	)
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

	departments, err = b.departmentService.FindExcDepartment(uint(divId), uint(id))
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", departments)
		c.JSON(http.StatusOK, response)
	}
}

func (b *departmentController) FindDepartmentByDivId(c *gin.Context) {
	var (
		departments []model.SelectDepartmentParameter
		response    helper.Response
	)
	divId, err := strconv.ParseUint(c.Param("divId"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	departments, err = b.departmentService.FindDepartmentByDivId(uint(divId))
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", departments)
		c.JSON(http.StatusOK, response)
	}
}

func (b *departmentController) InsertDepartment(c *gin.Context) {
	var (
		department                model.Department
		response                  helper.Response
		CreateDepartmentParameter model.CreateDepartmentParameter
	)
	err := c.ShouldBindJSON(&CreateDepartmentParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		department, err = b.departmentService.InsertDepartment(CreateDepartmentParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register department", err.Error(), nil)
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", department)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *departmentController) UpdateDepartment(c *gin.Context) {
	var (
		newData  model.Department
		oldData  model.SelectDepartmentParameter
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
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	oldData, err = b.departmentService.FindDepartmentById(uint(id))
	if (oldData == model.SelectDepartmentParameter{}) {
		res := helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		department, err := b.departmentService.UpdateDepartment(newData, uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Failed to update department", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", department)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *departmentController) DeleteDepartment(c *gin.Context) {
	var (
		newData  model.Department
		oldData  model.SelectDepartmentParameter
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
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	oldData, err = b.departmentService.FindDepartmentById(uint(id))
	if (oldData == model.SelectDepartmentParameter{}) {
		res := helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
	} else {
		department, err := b.departmentService.DeleteDepartment(newData, uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Failed to update department", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", department)
			c.JSON(http.StatusOK, response)
		}
	}
}
