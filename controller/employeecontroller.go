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
	CountEmployeeAll(c *gin.Context)
	FindEmployees(c *gin.Context)
	FindEmployeesOffset(c *gin.Context)
	SearchEmployee(c *gin.Context)
	CountSearchEmployee(c *gin.Context)
	FindEmployeeById(c *gin.Context)
	FindEmployeeByNik(c *gin.Context)
	FindExcEmployee(c *gin.Context)
	FindEmployeeByDepartment(c *gin.Context)
	FindEmployeeByPosition(c *gin.Context)
	FindEmployeeByDivIdAndDepId(c *gin.Context)
	FindEmployeeByDate(c *gin.Context)
	FindEmployeeCuti(c *gin.Context)
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

func (b *employeeController) CountEmployeeAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	count, err := b.employeeService.CountEmployeeAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeController) FindEmployees(c *gin.Context) {
	var (
		employees []model.SelectEmployeeParameter
		response  helper.Response
	)
	employees, err := b.employeeService.FindEmployees()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", employees)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeController) FindEmployeesOffset(c *gin.Context) {
	var (
		employees []model.SelectEmployeeParameter
		response  helper.Response
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
					employees, err = b.employeeService.FindEmployeesOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", employees)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *employeeController) SearchEmployee(c *gin.Context) {
	var (
		employees []model.SelectEmployeeParameter
		response  helper.Response
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
				response = helper.BuildErrorResponse("No param order was found", "No data with given order", helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				dir := c.Param("dir")
				if dir == "" {
					response = helper.BuildErrorResponse("No param dir was found", "No data with given dir", helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					search := c.Param("search")
					if search == "" {
						response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						employees, err = b.employeeService.SearchEmployee(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", employees)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *employeeController) CountSearchEmployee(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.employeeService.CountSearchEmployee(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
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
	} else {
		employee, err = b.employeeService.FindEmployeeById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employee)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeController) FindEmployeeByNik(c *gin.Context) {
	var (
		response helper.Response
	)
	nik := c.Param("nik")
	if nik == "" {
		response = helper.BuildErrorResponse("No param nik found", "No data with given department name", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employee, err := b.employeeService.FindEmployeeByNik(nik)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employee)
			c.JSON(http.StatusOK, response)
		}
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
	} else {
		employees, err = b.employeeService.FindExcEmployee(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employees)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeController) FindEmployeeByDivIdAndDepId(c *gin.Context) {
	var (
		response helper.Response
	)
	DivId, err := strconv.ParseUint(c.Param("DivId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param division id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		DepId, err := strconv.ParseUint(c.Param("DepId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			employees, err := b.employeeService.FindEmployeeByDivIdAndDepId(uint(DivId), uint(DepId))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", employees)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *employeeController) FindEmployeeByDepartment(c *gin.Context) {
	var (
		response helper.Response
	)
	DepId, err := strconv.ParseUint(c.Param("department"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employees, err := b.employeeService.FindEmployeeByDepartment(int(DepId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employees)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeController) FindEmployeeByPosition(c *gin.Context) {
	var (
		response helper.Response
	)
	group, err := strconv.ParseUint(c.Param("group"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param group id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		division := c.Param("division")
		if division == "" {
			response = helper.BuildErrorResponse("No param division name was found", "No data with given division name", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			department := c.Param("department")
			if department == "" {
				response = helper.BuildErrorResponse("No param department name was found", "No data with given department name", helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				position := c.Param("position")
				if position == "" {
					response = helper.BuildErrorResponse("No param position name was found", "No data with given position name", helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					employees, err := b.employeeService.FindEmployeeByPosition(int(group), division, department, position)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", employees)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *employeeController) FindEmployeeCuti(c *gin.Context) {
	var (
		response helper.Response
	)
	group, err := strconv.ParseUint(c.Param("group"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param group id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		section, err := strconv.ParseUint(c.Param("section"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param section id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			division, err := strconv.ParseUint(c.Param("division"), 0, 0)
			if err != nil {
				response = helper.BuildErrorResponse("No param division id was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				department, err := strconv.ParseUint(c.Param("department"), 0, 0)
				if err != nil {
					response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					position, err := strconv.ParseUint(c.Param("position"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						employees, err := b.employeeService.FindEmployeeCuti(int(group), int(section), int(division), int(department), int(position))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", employees)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *employeeController) FindEmployeeByDate(c *gin.Context) {
	var (
		response helper.Response
	)
	date, err := strconv.ParseFloat(c.Param("date"), 64)
	if err != nil {
		response = helper.BuildErrorResponse("No param date was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employees, err := b.employeeService.FindEmployeeByDate(float64(date))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employees)
			c.JSON(http.StatusOK, response)
		}
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
			response = helper.BuildErrorResponse("Failed to register employee", err.Error(), helper.EmptyObj{})
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
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		err := c.ShouldBindJSON(&newData)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			oldData, err = b.employeeService.FindEmployeeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectEmployeeParameter{})) {
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
	}
}

func (b *employeeController) DeleteEmployee(c *gin.Context) {
	var (
		newData  model.Employee
		oldData  model.SelectEmployeeParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		err := c.ShouldBindJSON(&newData)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			oldData, err = b.employeeService.FindEmployeeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectEmployeeParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employee, err := b.employeeService.DeleteEmployee(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete employee", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employee)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
