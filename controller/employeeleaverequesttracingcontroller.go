package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeLeaveRequestTracingController interface {
	CountEmployeeLeaveRequestTracingAll(c *gin.Context)
	FindEmployeeLeaveRequestTracings(c *gin.Context)
	FindEmployeeLeaveRequestTracingById(c *gin.Context)
	FindExcEmployeeLeaveRequestTracing(c *gin.Context)
	InsertEmployeeLeaveRequestTracing(c *gin.Context)
	UpdateEmployeeLeaveRequestTracing(c *gin.Context)
	DeleteEmployeeLeaveRequestTracing(c *gin.Context)
}

type employeeLeaveRequestTracingController struct {
	employeeLeaveRequestTracingService service.EmployeeLeaveRequestTracingService
	jwtService                         service.JWTService
}

func NewEmployeeLeaveRequestTracingController(employeeLeaveRequestTracingServ service.EmployeeLeaveRequestTracingService, jwtServ service.JWTService) EmployeeLeaveRequestTracingController {
	return &employeeLeaveRequestTracingController{
		employeeLeaveRequestTracingService: employeeLeaveRequestTracingServ,
		jwtService:                         jwtServ,
	}
}

func (b *employeeLeaveRequestTracingController) CountEmployeeLeaveRequestTracingAll(c *gin.Context) {
	var (
		response helper.Response
	)
	employeeLeaveRequestId, err := strconv.ParseInt(c.Param("employeeLeaveRequestId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employeeLeaveRequestId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.employeeLeaveRequestTracingService.CountEmployeeLeaveRequestTracingAll(int(employeeLeaveRequestId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestTracingController) FindEmployeeLeaveRequestTracings(c *gin.Context) {
	var (
		internalmemotracings []model.EmployeeLeaveRequestTracing
		response             helper.Response
	)
	employeeLeaveRequestId, err := strconv.ParseInt(c.Param("employeeLeaveRequestId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employeeLeaveRequestId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalmemotracings, err = b.employeeLeaveRequestTracingService.FindEmployeeLeaveRequestTracings(int(employeeLeaveRequestId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalmemotracings)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestTracingController) FindEmployeeLeaveRequestTracingById(c *gin.Context) {
	var (
		employeeLeaveRequestTracing model.EmployeeLeaveRequestTracing
		response                    helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestTracing, err = b.employeeLeaveRequestTracingService.FindEmployeeLeaveRequestTracingById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestTracing)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestTracingController) FindExcEmployeeLeaveRequestTracing(c *gin.Context) {
	var (
		employeeLeaveRequestTracings []model.EmployeeLeaveRequestTracing
		response                     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestTracings, err = b.employeeLeaveRequestTracingService.FindExcEmployeeLeaveRequestTracing(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestTracings)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestTracingController) InsertEmployeeLeaveRequestTracing(c *gin.Context) {
	var (
		employeeLeaveRequestTracing                model.EmployeeLeaveRequestTracing
		response                                   helper.Response
		CreateEmployeeLeaveRequestTracingParameter model.CreateEmployeeLeaveRequestTracingParameter
	)
	err := c.ShouldBindJSON(&CreateEmployeeLeaveRequestTracingParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestTracing, err = b.employeeLeaveRequestTracingService.InsertEmployeeLeaveRequestTracing(CreateEmployeeLeaveRequestTracingParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register internal memo tracing", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestTracing)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestTracingController) UpdateEmployeeLeaveRequestTracing(c *gin.Context) {
	var (
		newData  model.EmployeeLeaveRequestTracing
		oldData  model.EmployeeLeaveRequestTracing
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
			oldData, err = b.employeeLeaveRequestTracingService.FindEmployeeLeaveRequestTracingById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.EmployeeLeaveRequestTracing{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employeeLeaveRequestTracing, err := b.employeeLeaveRequestTracingService.UpdateEmployeeLeaveRequestTracing(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update internal memo tracing", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employeeLeaveRequestTracing)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *employeeLeaveRequestTracingController) DeleteEmployeeLeaveRequestTracing(c *gin.Context) {
	var (
		newData  model.EmployeeLeaveRequestTracing
		oldData  model.EmployeeLeaveRequestTracing
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
			oldData, err = b.employeeLeaveRequestTracingService.FindEmployeeLeaveRequestTracingById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.EmployeeLeaveRequestTracing{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employeeLeaveRequestTracing, err := b.employeeLeaveRequestTracingService.DeleteEmployeeLeaveRequestTracing(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete internal memo tracing", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employeeLeaveRequestTracing)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
