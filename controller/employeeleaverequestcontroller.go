package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeLeaveRequestController interface {
	CountEmployeeLeaveRequestAll(c *gin.Context)
	CountEmployeeLeaveRequestApprovalAllDt(c *gin.Context)
	FindEmployeeLeaveRequests(c *gin.Context)
	FindEmployeeLeaveRequestsOffset(c *gin.Context)
	SearchEmployeeLeaveRequest(c *gin.Context)
	CountSearchEmployeeLeaveRequest(c *gin.Context)
	FindEmployeeLeaveRequestById(c *gin.Context)
	CountEmployeeLeaveRequestAllPIC(c *gin.Context)
	FindEmployeeLeaveRequestsPICOffset(c *gin.Context)
	CountSearchEmployeeLeaveRequestPIC(c *gin.Context)
	SearchEmployeeLeaveRequestPIC(c *gin.Context)
	FindEmployeeLeaveRequestPICById(c *gin.Context)
	CountEmployeeLeaveRequestDaftarCuti(c *gin.Context)
	FindEmployeeLeaveRequestsDaftarCutiOffset(c *gin.Context)
	SearchEmployeeLeaveRequestDaftarCuti(c *gin.Context)
	CountSearchEmployeeLeaveRequestDaftarCuti(c *gin.Context)
	FindEmployeeLeaveRequestByEmpId(c *gin.Context)
	FindEmployeeLeaveRequestByEmpIdDate(c *gin.Context)
	FindEmployeeLeaveRequestByEmpIdExcDate(c *gin.Context)
	FindEmployeeLeaveRequestDraftOrAFA(c *gin.Context)
	FindEmployeeLeaveRequestPeriodDate(c *gin.Context)
	CountEmployeeLeaveRequestPersonaliaAllPIC(c *gin.Context)
	FindEmployeeLeaveRequestsPersonaliaPICOffset(c *gin.Context)
	CountSearchEmployeeLeaveRequestPersonaliaPIC(c *gin.Context)
	SearchEmployeeLeaveRequestPersonaliaPIC(c *gin.Context)
	FindAllEmployeeLeaveRequestByEmpId(c *gin.Context)
	FindEmployeeLeaveRequestFromDate(c *gin.Context)
	FindEmployeeLeaveRequestToDate(c *gin.Context)
	FindEmployeeLeaveRequestDate(c *gin.Context)
	FindAllEmployeeLeaveRequestByEmpIdDate(c *gin.Context)
	// FindExcEmployeeLeaveRequest(c *gin.Context)
	// FindEmployeeLeaveRequestByNWTId(c *gin.Context)
	// CountEmployeeLeaveRequestName(c *gin.Context)
	InsertEmployeeLeaveRequest(c *gin.Context)
	UpdateEmployeeLeaveRequest(c *gin.Context)
	DeleteEmployeeLeaveRequest(c *gin.Context)
}

type employeeLeaveRequestController struct {
	employeeLeaveRequestService service.EmployeeLeaveRequestService
	jwtService                  service.JWTService
}

func NewEmployeeLeaveRequestController(employeeLeaveRequestServ service.EmployeeLeaveRequestService, jwtServ service.JWTService) EmployeeLeaveRequestController {
	return &employeeLeaveRequestController{
		employeeLeaveRequestService: employeeLeaveRequestServ,
		jwtService:                  jwtServ,
	}
}

func (b *employeeLeaveRequestController) CountEmployeeLeaveRequestAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)
	empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param empId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err = b.employeeLeaveRequestService.CountEmployeeLeaveRequestAll(int(empId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestController) CountEmployeeLeaveRequestApprovalAllDt(c *gin.Context) {
	var (
		response helper.Response
	)
	year := c.Param("year")
	if year == "" {
		response = helper.BuildErrorResponse("No param year was found", "No data with given year", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.employeeLeaveRequestService.CountEmployeeLeaveRequestApprovalAllDt(year)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequests(c *gin.Context) {
	var (
		employeeLeaveRequests []model.SelectEmployeeLeaveRequestParameter
		response              helper.Response
	)
	employeeLeaveRequests, err := b.employeeLeaveRequestService.FindEmployeeLeaveRequests()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestsOffset(c *gin.Context) {
	var (
		employeeLeaveRequests []model.SelectEmployeeLeaveRequestParameter
		response              helper.Response
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
					empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param empId was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						employeeLeaveRequests, err = b.employeeLeaveRequestService.FindEmployeeLeaveRequestsOffset(int(limit), int(offset), order, dir, int(empId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) SearchEmployeeLeaveRequest(c *gin.Context) {
	var (
		employeeLeaveRequests []model.SelectEmployeeLeaveRequestParameter
		response              helper.Response
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
						empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param empId was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							employeeLeaveRequests, err = b.employeeLeaveRequestService.SearchEmployeeLeaveRequest(int(limit), int(offset), order, dir, search, int(empId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) CountSearchEmployeeLeaveRequest(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param empId was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.employeeLeaveRequestService.CountSearchEmployeeLeaveRequest(search, int(empId))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", count)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestById(c *gin.Context) {
	var (
		employeeLeaveRequest model.SelectEmployeeLeaveRequestParameter
		response             helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequest, err = b.employeeLeaveRequestService.FindEmployeeLeaveRequestById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequest)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestController) CountEmployeeLeaveRequestAllPIC(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)
	empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param empId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err = b.employeeLeaveRequestService.CountEmployeeLeaveRequestAllPIC(int(empId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}
func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestsPICOffset(c *gin.Context) {
	var (
		employeeLeaveRequests []model.SelectEmployeeLeaveRequestPICParameter
		response              helper.Response
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
					empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param empId was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						employeeLeaveRequests, err = b.employeeLeaveRequestService.FindEmployeeLeaveRequestsPICOffset(int(limit), int(offset), order, dir, int(empId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) SearchEmployeeLeaveRequestPIC(c *gin.Context) {
	var (
		employeeLeaveRequests []model.SelectEmployeeLeaveRequestPICParameter
		response              helper.Response
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
						empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param empId was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							employeeLeaveRequests, err = b.employeeLeaveRequestService.SearchEmployeeLeaveRequestPIC(int(limit), int(offset), order, dir, search, int(empId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) CountSearchEmployeeLeaveRequestPIC(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param empId was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.employeeLeaveRequestService.CountSearchEmployeeLeaveRequestPIC(search, int(empId))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", count)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestPICById(c *gin.Context) {
	var (
		employeeLeaveRequest model.SelectEmployeeLeaveRequestPICParameter
		response             helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequest, err = b.employeeLeaveRequestService.FindEmployeeLeaveRequestPICById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequest)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestController) CountEmployeeLeaveRequestDaftarCuti(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)
	count, err := b.employeeLeaveRequestService.CountEmployeeLeaveRequestDaftarCuti()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestsDaftarCutiOffset(c *gin.Context) {
	var (
		employeeLeaveRequests []model.SelectEmployeeLeaveRequestParameter
		response              helper.Response
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
					employeeLeaveRequests, err = b.employeeLeaveRequestService.FindEmployeeLeaveRequestsDaftarCutiOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) SearchEmployeeLeaveRequestDaftarCuti(c *gin.Context) {
	var (
		employeeLeaveRequests []model.SelectEmployeeLeaveRequestParameter
		response              helper.Response
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
						employeeLeaveRequests, err = b.employeeLeaveRequestService.SearchEmployeeLeaveRequestDaftarCuti(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) CountSearchEmployeeLeaveRequestDaftarCuti(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.employeeLeaveRequestService.CountSearchEmployeeLeaveRequestDaftarCuti(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestByEmpId(c *gin.Context) {
	var (
		response helper.Response
	)
	empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employee id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		year := c.Param("year")
		if year == "" {
			response = helper.BuildErrorResponse("No param year was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.employeeLeaveRequestService.FindEmployeeLeaveRequestByEmpId(uint(empId), year)
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", count)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestByEmpIdDate(c *gin.Context) {
	var (
		response helper.Response
	)
	empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employee id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		year := c.Param("year")
		if year == "" {
			response = helper.BuildErrorResponse("No param year was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.employeeLeaveRequestService.FindEmployeeLeaveRequestByEmpIdDate(uint(empId), year)
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", count)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestByEmpIdExcDate(c *gin.Context) {
	var (
		response helper.Response
	)
	elrId, err := strconv.ParseInt(c.Param("elrId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employee leave request id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param employee id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			year := c.Param("year")
			if year == "" {
				response = helper.BuildErrorResponse("No param year was found", "No data with given search", helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				count, err := b.employeeLeaveRequestService.FindEmployeeLeaveRequestByEmpIdExcDate(uint(elrId), uint(empId), year)
				if err != nil {
					response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
					c.JSON(http.StatusNotFound, response)
				} else {
					response = helper.BuildResponse(true, "OK", count)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestDraftOrAFA(c *gin.Context) {
	var (
		response helper.Response
	)
	empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employee id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.employeeLeaveRequestService.FindEmployeeLeaveRequestDraftOrAFA(uint(empId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestPeriodDate(c *gin.Context) {
	var (
		response helper.Response
	)
	date, err := strconv.ParseFloat(c.Param("date"), 64)
	if err != nil {
		response = helper.BuildErrorResponse("No param date was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.employeeLeaveRequestService.FindEmployeeLeaveRequestPeriodDate(float64(date))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestController) CountEmployeeLeaveRequestPersonaliaAllPIC(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)
	count, err := b.employeeLeaveRequestService.CountEmployeeLeaveRequestPersonaliaAllPIC()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}
func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestsPersonaliaPICOffset(c *gin.Context) {
	var (
		employeeLeaveRequests []model.SelectEmployeeLeaveRequestPICParameter
		response              helper.Response
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
					employeeLeaveRequests, err = b.employeeLeaveRequestService.FindEmployeeLeaveRequestsPersonaliaPICOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) SearchEmployeeLeaveRequestPersonaliaPIC(c *gin.Context) {
	var (
		employeeLeaveRequests []model.SelectEmployeeLeaveRequestPICParameter
		response              helper.Response
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
						employeeLeaveRequests, err = b.employeeLeaveRequestService.SearchEmployeeLeaveRequestPersonaliaPIC(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) CountSearchEmployeeLeaveRequestPersonaliaPIC(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.employeeLeaveRequestService.CountSearchEmployeeLeaveRequestPersonaliaPIC(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestController) FindAllEmployeeLeaveRequestByEmpId(c *gin.Context) {
	var (
		response helper.Response
	)
	empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employee id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		year := c.Param("year")
		if year == "" {
			response = helper.BuildErrorResponse("No param year was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			date, err := strconv.ParseFloat(c.Param("date"), 64)
			if err != nil {
				response = helper.BuildErrorResponse("No param date was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				count, err := b.employeeLeaveRequestService.FindAllEmployeeLeaveRequestByEmpId(uint(empId), year, date)
				if err != nil {
					response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
					c.JSON(http.StatusNotFound, response)
				} else {
					response = helper.BuildResponse(true, "OK", count)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestFromDate(c *gin.Context) {
	var (
		response helper.Response
	)
	date, err := strconv.ParseFloat(c.Param("date"), 64)
	if err != nil {
		response = helper.BuildErrorResponse("No param date was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.employeeLeaveRequestService.FindEmployeeLeaveRequestFromDate(float64(date))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestToDate(c *gin.Context) {
	var (
		response helper.Response
	)
	date, err := strconv.ParseFloat(c.Param("date"), 64)
	if err != nil {
		response = helper.BuildErrorResponse("No param date from was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		date2, err := strconv.ParseFloat(c.Param("date2"), 64)
		if err != nil {
			response = helper.BuildErrorResponse("No param date to was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.employeeLeaveRequestService.FindEmployeeLeaveRequestToDate(float64(date), float64(date2))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", count)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestDate(c *gin.Context) {
	var (
		response helper.Response
	)
	date, err := strconv.ParseFloat(c.Param("date"), 64)
	if err != nil {
		response = helper.BuildErrorResponse("No param date from was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		date2, err := strconv.ParseFloat(c.Param("date2"), 64)
		if err != nil {
			response = helper.BuildErrorResponse("No param date to was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.employeeLeaveRequestService.FindEmployeeLeaveRequestDate(float64(date), float64(date2))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", count)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *employeeLeaveRequestController) FindAllEmployeeLeaveRequestByEmpIdDate(c *gin.Context) {
	var (
		response helper.Response
	)
	empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employee id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		year := c.Param("year")
		if year == "" {
			response = helper.BuildErrorResponse("No param year was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			date, err := strconv.ParseFloat(c.Param("date"), 64)
			if err != nil {
				response = helper.BuildErrorResponse("No param date was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				count, err := b.employeeLeaveRequestService.FindAllEmployeeLeaveRequestByEmpIdDate(uint(empId), year, date)
				if err != nil {
					response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
					c.JSON(http.StatusNotFound, response)
				} else {
					response = helper.BuildResponse(true, "OK", count)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

// func (b *employeeLeaveRequestController) FindExcEmployeeLeaveRequest(c *gin.Context) {
// 	var (
// 		employeeLeaveRequests []model.SelectEmployeeLeaveRequestParameter
// 		response       helper.Response
// 	)
// 	nwtId, err := strconv.ParseUint(c.Param("nwtId"), 0, 0)
// 	if err != nil {
// 		response = helper.BuildErrorResponse("No param non working type id was found", err.Error(), helper.EmptyObj{})
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 	} else {
// 		id, err := strconv.ParseUint(c.Param("id"), 0, 0)
// 		if err != nil {
// 			response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
// 			c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		} else {
// 			employeeLeaveRequests, err = b.employeeLeaveRequestService.FindExcEmployeeLeaveRequest(uint(nwtId), uint(id))
// 			if err != nil {
// 				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
// 				c.JSON(http.StatusNotFound, response)
// 			} else {
// 				response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
// 				c.JSON(http.StatusOK, response)
// 			}
// 		}
// 	}
// }

// func (b *employeeLeaveRequestController) FindEmployeeLeaveRequestByNWTId(c *gin.Context) {
// 	var (
// 		employeeLeaveRequests []model.SelectEmployeeLeaveRequestParameter
// 		response       helper.Response
// 	)
// 	nwtId, err := strconv.ParseUint(c.Param("nwtId"), 0, 0)
// 	if err != nil {
// 		response = helper.BuildErrorResponse("No param non working type id was found", err.Error(), helper.EmptyObj{})
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 	} else {
// 		employeeLeaveRequests, err = b.employeeLeaveRequestService.FindEmployeeLeaveRequestByNWTId(uint(nwtId))
// 		if err != nil {
// 			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
// 			c.JSON(http.StatusNotFound, response)
// 		} else {
// 			response = helper.BuildResponse(true, "OK", employeeLeaveRequests)
// 			c.JSON(http.StatusOK, response)
// 		}
// 	}
// }

// func (b *employeeLeaveRequestController) CountEmployeeLeaveRequestName(c *gin.Context) {
// 	var (
// 		response helper.Response
// 	)
// 	search := c.Param("search")
// 	if search == "" {
// 		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 	} else {
// 		count, err := b.employeeLeaveRequestService.CountEmployeeLeaveRequestName(search)
// 		if err != nil {
// 			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
// 			c.JSON(http.StatusNotFound, response)
// 		} else {
// 			response = helper.BuildResponse(true, "OK", count)
// 			c.JSON(http.StatusOK, response)
// 		}
// 	}
// }

func (b *employeeLeaveRequestController) InsertEmployeeLeaveRequest(c *gin.Context) {
	var (
		employeeLeaveRequest                model.EmployeeLeaveRequest
		response                            helper.Response
		CreateEmployeeLeaveRequestParameter model.CreateEmployeeLeaveRequestParameter
	)
	err := c.ShouldBindJSON(&CreateEmployeeLeaveRequestParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequest, err = b.employeeLeaveRequestService.InsertEmployeeLeaveRequest(CreateEmployeeLeaveRequestParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register employeeLeaveRequest", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequest)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestController) UpdateEmployeeLeaveRequest(c *gin.Context) {
	var (
		newData  model.EmployeeLeaveRequest
		oldData  model.SelectEmployeeLeaveRequestParameter
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
			oldData, err = b.employeeLeaveRequestService.FindEmployeeLeaveRequestById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectEmployeeLeaveRequestParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employeeLeaveRequest, err := b.employeeLeaveRequestService.UpdateEmployeeLeaveRequest(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update employeeLeaveRequest", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employeeLeaveRequest)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *employeeLeaveRequestController) DeleteEmployeeLeaveRequest(c *gin.Context) {
	var (
		newData  model.EmployeeLeaveRequest
		oldData  model.SelectEmployeeLeaveRequestParameter
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
			oldData, err = b.employeeLeaveRequestService.FindEmployeeLeaveRequestById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectEmployeeLeaveRequestParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employeeLeaveRequest, err := b.employeeLeaveRequestService.DeleteEmployeeLeaveRequest(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete employeeLeaveRequest", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employeeLeaveRequest)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
