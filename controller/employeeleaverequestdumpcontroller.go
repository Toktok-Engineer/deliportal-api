package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeLeaveRequestDumpController interface {
	CountEmployeeLeaveRequestDumpAll(c *gin.Context)
	FindEmployeeLeaveRequestDumps(c *gin.Context)
	FindEmployeeLeaveRequestDumpsOffset(c *gin.Context)
	// SearchEmployeeLeaveRequestDump(c *gin.Context)
	// CountSearchEmployeeLeaveRequestDump(c *gin.Context)
	FindEmployeeLeaveRequestDumpById(c *gin.Context)
	FindExcEmployeeLeaveRequestDump(c *gin.Context)
	InsertEmployeeLeaveRequestDump(c *gin.Context)
	UpdateEmployeeLeaveRequestDump(c *gin.Context)
	DeleteEmployeeLeaveRequestDump(c *gin.Context)
}

type employeeLeaveRequestDumpController struct {
	employeeLeaveRequestDumpService service.EmployeeLeaveRequestDumpService
	jwtService                      service.JWTService
}

func NewEmployeeLeaveRequestDumpController(employeeLeaveRequestDumpServ service.EmployeeLeaveRequestDumpService, jwtServ service.JWTService) EmployeeLeaveRequestDumpController {
	return &employeeLeaveRequestDumpController{
		employeeLeaveRequestDumpService: employeeLeaveRequestDumpServ,
		jwtService:                      jwtServ,
	}
}

func (b *employeeLeaveRequestDumpController) CountEmployeeLeaveRequestDumpAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.employeeLeaveRequestDumpService.CountEmployeeLeaveRequestDumpAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeLeaveRequestDumpController) FindEmployeeLeaveRequestDumps(c *gin.Context) {
	var (
		businessunits []model.EmployeeLeaveRequestDump
		response      helper.Response
	)
	businessunits, err := b.employeeLeaveRequestDumpService.FindEmployeeLeaveRequestDumps()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", businessunits)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeLeaveRequestDumpController) FindEmployeeLeaveRequestDumpsOffset(c *gin.Context) {
	var (
		businessunits []model.EmployeeLeaveRequestDump
		response      helper.Response
	)

	limit, err := strconv.ParseInt(c.Param("limit"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param limit was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		offset, err := strconv.ParseInt(c.Param("offset"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param offset was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			order := c.Param("order")
			if order == "" {
				response = helper.BuildErrorResponse("No param order was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				dir := c.Param("dir")
				if dir == "" {
					response = helper.BuildErrorResponse("No param dir was found", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					businessunits, err = b.employeeLeaveRequestDumpService.FindEmployeeLeaveRequestDumpsOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", businessunits)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

// func (b *employeeLeaveRequestDumpController) SearchEmployeeLeaveRequestDump(c *gin.Context) {
// 	var (
// 		businessunits []model.EmployeeLeaveRequestDump
// 		response      helper.Response
// 	)

// 	limit, err := strconv.ParseInt(c.Param("limit"), 0, 0)
// 	if err != nil {
// 		response = helper.BuildErrorResponse("No param limit was found", err.Error(), helper.EmptyObj{})
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 	} else {
// 		offset, err := strconv.ParseInt(c.Param("offset"), 0, 0)
// 		if err != nil {
// 			response = helper.BuildErrorResponse("No param offset was found", err.Error(), helper.EmptyObj{})
// 			c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		} else {
// 			order := c.Param("order")
// 			if order == "" {
// 				response = helper.BuildErrorResponse("No param order was found", "No data with given order", helper.EmptyObj{})
// 				c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 			} else {
// 				dir := c.Param("dir")
// 				if dir == "" {
// 					response = helper.BuildErrorResponse("No param dir was found", "No data with given dir", helper.EmptyObj{})
// 					c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 				} else {
// 					search := c.Param("search")
// 					if search == "" {
// 						response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
// 						c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 					} else {
// 						businessunits, err = b.employeeLeaveRequestDumpService.SearchEmployeeLeaveRequestDump(int(limit), int(offset), order, dir, search)
// 						if err != nil {
// 							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
// 							c.JSON(http.StatusNotFound, response)
// 						} else {
// 							response = helper.BuildResponse(true, "OK", businessunits)
// 							c.JSON(http.StatusOK, response)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func (b *employeeLeaveRequestDumpController) CountSearchEmployeeLeaveRequestDump(c *gin.Context) {
// 	var (
// 		response helper.Response
// 	)
// 	search := c.Param("search")
// 	if search == "" {
// 		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 	} else {
// 		count, err := b.employeeLeaveRequestDumpService.CountSearchEmployeeLeaveRequestDump(search)
// 		if err != nil {
// 			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
// 			c.JSON(http.StatusNotFound, response)
// 		} else {
// 			response = helper.BuildResponse(true, "OK", count)
// 			c.JSON(http.StatusOK, response)
// 		}
// 	}
// }

func (b *employeeLeaveRequestDumpController) FindEmployeeLeaveRequestDumpById(c *gin.Context) {
	var (
		employeeLeaveRequestDump model.EmployeeLeaveRequestDump
		response                 helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestDump, err = b.employeeLeaveRequestDumpService.FindEmployeeLeaveRequestDumpById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestDump)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestDumpController) FindExcEmployeeLeaveRequestDump(c *gin.Context) {
	var (
		employeeLeaveRequestDumps []model.EmployeeLeaveRequestDump
		response                  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestDumps, err = b.employeeLeaveRequestDumpService.FindExcEmployeeLeaveRequestDump(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestDumps)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestDumpController) InsertEmployeeLeaveRequestDump(c *gin.Context) {
	var (
		employeeLeaveRequestDump                model.EmployeeLeaveRequestDump
		response                                helper.Response
		CreateEmployeeLeaveRequestDumpParameter model.CreateEmployeeLeaveRequestDumpParameter
	)
	err := c.ShouldBindJSON(&CreateEmployeeLeaveRequestDumpParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestDump, err = b.employeeLeaveRequestDumpService.InsertEmployeeLeaveRequestDump(CreateEmployeeLeaveRequestDumpParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register employee leave request dump", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestDump)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestDumpController) UpdateEmployeeLeaveRequestDump(c *gin.Context) {
	var (
		newData  model.EmployeeLeaveRequestDump
		oldData  model.EmployeeLeaveRequestDump
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
			oldData, err = b.employeeLeaveRequestDumpService.FindEmployeeLeaveRequestDumpById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.EmployeeLeaveRequestDump{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employeeLeaveRequestDump, err := b.employeeLeaveRequestDumpService.UpdateEmployeeLeaveRequestDump(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update employee leave request dump", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employeeLeaveRequestDump)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *employeeLeaveRequestDumpController) DeleteEmployeeLeaveRequestDump(c *gin.Context) {
	var (
		newData  model.EmployeeLeaveRequestDump
		oldData  model.EmployeeLeaveRequestDump
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
			oldData, err = b.employeeLeaveRequestDumpService.FindEmployeeLeaveRequestDumpById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.EmployeeLeaveRequestDump{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employeeLeaveRequestDump, err := b.employeeLeaveRequestDumpService.DeleteEmployeeLeaveRequestDump(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete employee leave request dump", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employeeLeaveRequestDump)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
