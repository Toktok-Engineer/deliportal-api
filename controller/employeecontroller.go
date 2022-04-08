package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

type EmployeeController interface {
	FindEmployees(c *gin.Context)
	FindEmployeeById(c *gin.Context)
	FindEmployeeByNik(c *gin.Context)
	FindExcEmployee(c *gin.Context)
	InsertEmployee(c *gin.Context)
	UpdateEmployee(c *gin.Context)
	DeleteEmployee(c *gin.Context)
}

type employeeController struct {
	employeeService service.EmployeeService
	jwtService      service.JWTService
}

func NewEmployeeController(employeeServ service.EmployeeService, jwtServ service.JWTService) EmployeeController {
	return &employeeController{
		employeeService: employeeServ,
		jwtService:      jwtServ,
	}
}

func (b *employeeController) FindEmployees(c *gin.Context) {
	var (
		employees []model.SelectEmployeeParameter
		response  helper.Response
	)
	employees, err := b.employeeService.FindEmployees()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), nil)
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", employees)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeController) FindEmployeeById(c *gin.Context) {
	var (
		employee model.SelectEmployeeParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	employee, err = b.employeeService.FindEmployeeById(uint(id))
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", employee)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeController) FindEmployeeByNik(c *gin.Context) {
	var (
		employee model.SelectEmployeeParameter
		response helper.Response
	)
	nik, err := strconv.ParseUint(c.Param("nik"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	employee, err = b.employeeService.FindEmployeeByNik(uint(nik))
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", employee)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeController) FindExcEmployee(c *gin.Context) {
	var (
		employees []model.SelectEmployeeParameter
		response  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	employees, err = b.employeeService.FindExcEmployee(uint(id))
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", employees)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeController) InsertEmployee(c *gin.Context) {
	var (
		employee                model.Employee
		response                helper.Response
		CreateEmployeeParameter model.CreateEmployeeParameter
	)
	err := c.ShouldBindJSON(&CreateEmployeeParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		employee, err = b.employeeService.InsertEmployee(CreateEmployeeParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register employee", err.Error(), nil)
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", employee)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeController) UpdateEmployee(c *gin.Context) {
	var (
		newData  model.Employee
		oldData  model.SelectEmployeeParameter
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

	oldData, err = b.employeeService.FindEmployeeById(uint(id))
	if (cmp.Equal(oldData, model.SelectEmployeeParameter{})) {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		employee, err := b.employeeService.UpdateEmployee(newData, uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Failed to update employee", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", employee)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeController) DeleteEmployee(c *gin.Context) {
	var (
		newData  model.Employee
		oldData  model.SelectEmployeeParameter
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

	oldData, err = b.employeeService.FindEmployeeById(uint(id))
	if (cmp.Equal(oldData, model.SelectEmployeeParameter{})) {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, oldData)
	} else {
		employee, err := b.employeeService.DeleteEmployee(newData, uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Failed to update employee", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", employee)
			c.JSON(http.StatusOK, response)
		}
	}
}
