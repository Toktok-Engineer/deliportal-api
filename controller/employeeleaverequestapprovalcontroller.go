package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeLeaveRequestApprovalController interface {
	CountEmployeeLeaveRequestApprovalAll(c *gin.Context)
	FindEmployeeLeaveRequestApprovals(c *gin.Context)
	FindEmployeeLeaveRequestApprovalsOffset(c *gin.Context)
	SearchEmployeeLeaveRequestApproval(c *gin.Context)
	CountSearchEmployeeLeaveRequestApproval(c *gin.Context)
	FindEmployeeLeaveRequestApprovalById(c *gin.Context)
	// FindExcEmployeeLeaveRequestApproval(c *gin.Context)
	FindEmployeeLeaveRequestApprovalApproved(c *gin.Context)
	FindEmployeeLeaveRequestApprovalOpenApproved(c *gin.Context)
	FindEmployeeLeaveRequestApprovalByERId(c *gin.Context)
	// CountEmployeeLeaveRequestApprovalName(c *gin.Context)
	InsertEmployeeLeaveRequestApproval(c *gin.Context)
	UpdateEmployeeLeaveRequestApproval(c *gin.Context)
	DeleteEmployeeLeaveRequestApproval(c *gin.Context)
}

type employeeLeaveRequestApprovalController struct {
	employeeLeaveRequestApprovalService service.EmployeeLeaveRequestApprovalService
	jwtService                          service.JWTService
}

func NewEmployeeLeaveRequestApprovalController(employeeLeaveRequestApprovalServ service.EmployeeLeaveRequestApprovalService, jwtServ service.JWTService) EmployeeLeaveRequestApprovalController {
	return &employeeLeaveRequestApprovalController{
		employeeLeaveRequestApprovalService: employeeLeaveRequestApprovalServ,
		jwtService:                          jwtServ,
	}
}

func (b *employeeLeaveRequestApprovalController) CountEmployeeLeaveRequestApprovalAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	elrId, err := strconv.ParseInt(c.Param("elrId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employee leave request ID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err = b.employeeLeaveRequestApprovalService.CountEmployeeLeaveRequestApprovalAll(int(elrId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestApprovalController) FindEmployeeLeaveRequestApprovals(c *gin.Context) {
	var (
		employeeLeaveRequestApprovals []model.SelectEmployeeLeaveRequestApprovalParameter
		response                      helper.Response
	)
	employeeLeaveRequestApprovals, err := b.employeeLeaveRequestApprovalService.FindEmployeeLeaveRequestApprovals()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", employeeLeaveRequestApprovals)
		c.JSON(http.StatusOK, response)
	}
}

func (b *employeeLeaveRequestApprovalController) FindEmployeeLeaveRequestApprovalsOffset(c *gin.Context) {
	var (
		employeeLeaveRequestApprovals []model.SelectEmployeeLeaveRequestApprovalParameter
		response                      helper.Response
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
					elrId, err := strconv.ParseInt(c.Param("elrId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param employee leave request ID was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						employeeLeaveRequestApprovals, err = b.employeeLeaveRequestApprovalService.FindEmployeeLeaveRequestApprovalsOffset(int(limit), int(offset), order, dir, int(elrId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", employeeLeaveRequestApprovals)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveRequestApprovalController) SearchEmployeeLeaveRequestApproval(c *gin.Context) {
	var (
		employeeLeaveRequestApprovals []model.SelectEmployeeLeaveRequestApprovalParameter
		response                      helper.Response
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
						elrId, err := strconv.ParseInt(c.Param("elrId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param employee leave request ID was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							employeeLeaveRequestApprovals, err = b.employeeLeaveRequestApprovalService.SearchEmployeeLeaveRequestApproval(int(limit), int(offset), order, dir, search, int(elrId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", employeeLeaveRequestApprovals)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveRequestApprovalController) CountSearchEmployeeLeaveRequestApproval(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		elrId, err := strconv.ParseInt(c.Param("elrId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param employee leave request ID was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.employeeLeaveRequestApprovalService.CountSearchEmployeeLeaveRequestApproval(search, int(elrId))
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

func (b *employeeLeaveRequestApprovalController) FindEmployeeLeaveRequestApprovalById(c *gin.Context) {
	var (
		employeeLeaveRequestApproval model.SelectEmployeeLeaveRequestApprovalParameter
		response                     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestApproval, err = b.employeeLeaveRequestApprovalService.FindEmployeeLeaveRequestApprovalById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestApproval)
			c.JSON(http.StatusOK, response)
		}
	}
}

// func (b *employeeLeaveRequestApprovalController) FindExcEmployeeLeaveRequestApproval(c *gin.Context) {
// 	var (
// 		employeeLeaveRequestApprovals []model.SelectEmployeeLeaveRequestApprovalParameter
// 		response                      helper.Response
// 	)
// 	divId, err := strconv.ParseUint(c.Param("divId"), 0, 0)
// 	if err != nil {
// 		response = helper.BuildErrorResponse("No param divId was found", err.Error(), helper.EmptyObj{})
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 	} else {
// 		id, err := strconv.ParseUint(c.Param("id"), 0, 0)
// 		if err != nil {
// 			response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
// 			c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		} else {
// 			employeeLeaveRequestApprovals, err = b.employeeLeaveRequestApprovalService.FindExcEmployeeLeaveRequestApproval(uint(divId), uint(id))
// 			if err != nil {
// 				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
// 				c.JSON(http.StatusNotFound, response)
// 			} else {
// 				response = helper.BuildResponse(true, "OK", employeeLeaveRequestApprovals)
// 				c.JSON(http.StatusOK, response)
// 			}
// 		}
// 	}
// }

func (b *employeeLeaveRequestApprovalController) FindEmployeeLeaveRequestApprovalApproved(c *gin.Context) {
	var (
		employeeLeaveRequestApprovals []model.SelectEmployeeLeaveRequestApprovalParameter
		response                      helper.Response
	)
	elrId, err := strconv.ParseUint(c.Param("elrId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employee leave request id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestApprovals, err = b.employeeLeaveRequestApprovalService.FindEmployeeLeaveRequestApprovalApproved(uint(elrId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestApprovals)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestApprovalController) FindEmployeeLeaveRequestApprovalOpenApproved(c *gin.Context) {
	var (
		employeeLeaveRequestApprovals []model.SelectEmployeeLeaveRequestApprovalParameter
		response                      helper.Response
	)
	elrId, err := strconv.ParseUint(c.Param("elrId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employee leave request id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestApprovals, err = b.employeeLeaveRequestApprovalService.FindEmployeeLeaveRequestApprovalOpenApproved(uint(elrId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestApprovals)
			c.JSON(http.StatusOK, response)
		}
	}
}
func (b *employeeLeaveRequestApprovalController) FindEmployeeLeaveRequestApprovalByERId(c *gin.Context) {
	var (
		employeeLeaveRequestApprovals model.SelectEmployeeLeaveRequestApprovalParameter
		response                      helper.Response
	)
	elrId, err := strconv.ParseUint(c.Param("elrId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employee leave request id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestApprovals, err = b.employeeLeaveRequestApprovalService.FindEmployeeLeaveRequestApprovalByERId(uint(elrId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestApprovals)
			c.JSON(http.StatusOK, response)
		}
	}
}

// func (b *employeeLeaveRequestApprovalController) CountEmployeeLeaveRequestApprovalName(c *gin.Context) {
// 	var (
// 		response helper.Response
// 	)
// 	search := c.Param("search")
// 	if search == "" {
// 		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 	} else {
// 		count, err := b.employeeLeaveRequestApprovalService.CountEmployeeLeaveRequestApprovalName(search)
// 		if err != nil {
// 			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
// 			c.JSON(http.StatusNotFound, response)
// 		} else {
// 			response = helper.BuildResponse(true, "OK", count)
// 			c.JSON(http.StatusOK, response)
// 		}
// 	}
// }

func (b *employeeLeaveRequestApprovalController) InsertEmployeeLeaveRequestApproval(c *gin.Context) {
	var (
		employeeLeaveRequestApproval                model.EmployeeLeaveRequestApproval
		response                                    helper.Response
		CreateEmployeeLeaveRequestApprovalParameter model.CreateEmployeeLeaveRequestApprovalParameter
	)
	err := c.ShouldBindJSON(&CreateEmployeeLeaveRequestApprovalParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveRequestApproval, err = b.employeeLeaveRequestApprovalService.InsertEmployeeLeaveRequestApproval(CreateEmployeeLeaveRequestApprovalParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register employeeLeaveRequestApproval", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveRequestApproval)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveRequestApprovalController) UpdateEmployeeLeaveRequestApproval(c *gin.Context) {
	var (
		newData  model.EmployeeLeaveRequestApproval
		oldData  model.SelectEmployeeLeaveRequestApprovalParameter
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
			oldData, err = b.employeeLeaveRequestApprovalService.FindEmployeeLeaveRequestApprovalById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectEmployeeLeaveRequestApprovalParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employeeLeaveRequestApproval, err := b.employeeLeaveRequestApprovalService.UpdateEmployeeLeaveRequestApproval(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update employeeLeaveRequestApproval", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employeeLeaveRequestApproval)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *employeeLeaveRequestApprovalController) DeleteEmployeeLeaveRequestApproval(c *gin.Context) {
	var (
		newData  model.EmployeeLeaveRequestApproval
		oldData  model.SelectEmployeeLeaveRequestApprovalParameter
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
			oldData, err = b.employeeLeaveRequestApprovalService.FindEmployeeLeaveRequestApprovalById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectEmployeeLeaveRequestApprovalParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employeeLeaveRequestApproval, err := b.employeeLeaveRequestApprovalService.DeleteEmployeeLeaveRequestApproval(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete employeeLeaveRequestApproval", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employeeLeaveRequestApproval)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
