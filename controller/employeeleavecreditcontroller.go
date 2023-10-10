package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeLeaveCreditController interface {
	CountEmployeeLeaveCreditAll(c *gin.Context)
	FindEmployeeLeaveCredits(c *gin.Context)
	FindEmployeeLeaveCreditsOffset(c *gin.Context)
	CountSearchEmployeeLeaveCreditAll(c *gin.Context)
	FindSearchEmployeeLeaveCreditsOffset(c *gin.Context)
	CountEmployeeLeaveCreditEmpID(c *gin.Context)
	FindEmployeeLeaveCreditsOffsetByYear(c *gin.Context)
	CountSearchEmployeeLeaveCreditByEmp(c *gin.Context)
	FindSearchEmployeeLeaveCreditsOffsetByEmp(c *gin.Context)
	CountSearchEmployeeLeaveCreditByDepId(c *gin.Context)
	FindSearchEmployeeLeaveCreditsOffsetByDept(c *gin.Context)
	CountEmployeeLeaveCreditByDepId(c *gin.Context)
	FindEmployeeLeaveCreditsOffsetByDept(c *gin.Context)
	CountEmployeeLeaveCreditByDivAll(c *gin.Context)
	FindEmployeeLeaveCreditsOffsetByDivAll(c *gin.Context)
	CountSearchEmployeeLeaveCreditByDivIdAll(c *gin.Context)
	FindSearchEmployeeLeaveCreditsOffsetByDivIdAll(c *gin.Context)
	FindEmployeeLeaveCreditsExcEmp(c *gin.Context)
	FindEmployeeLeaveCreditsAll(c *gin.Context)
	CountEmployeeLeaveCreditByDepAll(c *gin.Context)
	FindEmployeeLeaveCreditsOffsetByDepAll(c *gin.Context)
	CountSearchEmployeeLeaveCreditByDepIdAll(c *gin.Context)
	FindSearchEmployeeLeaveCreditsOffsetByDepIdAll(c *gin.Context)
	// SearchEmployeeLeaveCredit(c *gin.Context)
	// CountSearchEmployeeLeaveCredit(c *gin.Context)
	FindEmployeeLeaveCreditById(c *gin.Context)
	// CountEmployeeLeaveCreditName(c *gin.Context)
	FindEmployeeLeaveCreditByEmpId(c *gin.Context)
	FindAllEmployeeLeaveCreditByEmpId(c *gin.Context)
	FindExcEmployeeLeaveCredit(c *gin.Context)
	InsertEmployeeLeaveCredit(c *gin.Context)
	UpdateEmployeeLeaveCredit(c *gin.Context)
	DeleteEmployeeLeaveCredit(c *gin.Context)
}

type employeeLeaveCreditController struct {
	employeeLeaveCreditService service.EmployeeLeaveCreditService
	jwtService                 service.JWTService
}

func NewEmployeeLeaveCreditController(employeeLeaveCreditServ service.EmployeeLeaveCreditService, jwtServ service.JWTService) EmployeeLeaveCreditController {
	return &employeeLeaveCreditController{
		employeeLeaveCreditService: employeeLeaveCreditServ,
		jwtService:                 jwtServ,
	}
}

func (b *employeeLeaveCreditController) CountEmployeeLeaveCreditAll(c *gin.Context) {
	var (
		response helper.Response
	)

	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.employeeLeaveCreditService.CountEmployeeLeaveCreditAll(int(year))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveCreditController) FindEmployeeLeaveCredits(c *gin.Context) {
	var (
		employeeLeaveCredits []model.EmployeeLeaveCredit
		response             helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param empId was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			employeeLeaveCredits, err = b.employeeLeaveCreditService.FindEmployeeLeaveCredits(int(year), int(empId))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *employeeLeaveCreditController) FindEmployeeLeaveCreditsOffset(c *gin.Context) {
	var (
		employeeLeaveCredits []model.SelectEmployeeLeaveCredit
		response             helper.Response
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
					year, err := strconv.ParseInt(c.Param("year"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						employeeLeaveCredits, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditsOffset(int(limit), int(offset), order, dir, int(year))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveCreditController) CountSearchEmployeeLeaveCreditAll(c *gin.Context) {
	var (
		response helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		search := c.Param("search")
		if search == "" {
			response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.employeeLeaveCreditService.CountSearchEmployeeLeaveCreditAll(int(year), search)
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

func (b *employeeLeaveCreditController) FindSearchEmployeeLeaveCreditsOffset(c *gin.Context) {
	var (
		employeeLeaveCredits []model.SelectEmployeeLeaveCredit
		response             helper.Response
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
						year, err := strconv.ParseInt(c.Param("year"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							employeeLeaveCredits, err = b.employeeLeaveCreditService.FindSearchEmployeeLeaveCreditsOffset(int(limit), int(offset), order, dir, int(year), search)
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveCreditController) CountEmployeeLeaveCreditEmpID(c *gin.Context) {
	var (
		response helper.Response
	)

	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param employee id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.employeeLeaveCreditService.CountEmployeeLeaveCreditEmpID(int(year), int(empId))
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

func (b *employeeLeaveCreditController) FindEmployeeLeaveCreditsOffsetByYear(c *gin.Context) {
	var (
		employeeLeaveCredits []model.SelectEmployeeLeaveCredit
		response             helper.Response
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
					year, err := strconv.ParseInt(c.Param("year"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param employee id was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							employeeLeaveCredits, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditsOffsetByYear(int(limit), int(offset), order, dir, int(year), int(empId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveCreditController) CountSearchEmployeeLeaveCreditByEmp(c *gin.Context) {
	var (
		response helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		search := c.Param("search")
		if search == "" {
			response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
			if err != nil {
				response = helper.BuildErrorResponse("No param employee id was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				count, err := b.employeeLeaveCreditService.CountSearchEmployeeLeaveCreditByEmp(int(year), int(empId), search)
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

func (b *employeeLeaveCreditController) FindSearchEmployeeLeaveCreditsOffsetByEmp(c *gin.Context) {
	var (
		employeeLeaveCredits []model.SelectEmployeeLeaveCredit
		response             helper.Response
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
						year, err := strconv.ParseInt(c.Param("year"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							empId, err := strconv.ParseInt(c.Param("empId"), 0, 0)
							if err != nil {
								response = helper.BuildErrorResponse("No param employee id was found", err.Error(), helper.EmptyObj{})
								c.AbortWithStatusJSON(http.StatusBadRequest, response)
							} else {
								employeeLeaveCredits, err = b.employeeLeaveCreditService.FindSearchEmployeeLeaveCreditsOffsetByEmp(int(limit), int(offset), order, dir, int(year), int(empId), search)
								if err != nil {
									response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
									c.JSON(http.StatusNotFound, response)
								} else {
									response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
									c.JSON(http.StatusOK, response)
								}
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveCreditController) CountSearchEmployeeLeaveCreditByDepId(c *gin.Context) {
	var (
		response helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		search := c.Param("search")
		if search == "" {
			response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			deptId, err := strconv.ParseInt(c.Param("deptId"), 0, 0)
			if err != nil {
				response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {

				positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
				if err != nil {
					response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						count, err := b.employeeLeaveCreditService.CountSearchEmployeeLeaveCreditByDepId(int(year), int(deptId), int(positionId), search, int(groupId))
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
	}
}

func (b *employeeLeaveCreditController) FindSearchEmployeeLeaveCreditsOffsetByDept(c *gin.Context) {
	var (
		employeeLeaveCredits []model.SelectEmployeeLeaveCredit
		response             helper.Response
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
						year, err := strconv.ParseInt(c.Param("year"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							deptId, err := strconv.ParseInt(c.Param("deptId"), 0, 0)
							if err != nil {
								response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
								c.AbortWithStatusJSON(http.StatusBadRequest, response)
							} else {
								positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
								if err != nil {
									response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
									c.AbortWithStatusJSON(http.StatusBadRequest, response)
								} else {

									groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
									if err != nil {
										response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
										c.AbortWithStatusJSON(http.StatusBadRequest, response)
									} else {
										employeeid, err := strconv.ParseInt(c.Param("employeeid"), 0, 0)
										if err != nil {
											response = helper.BuildErrorResponse("No param group employee id was found", err.Error(), helper.EmptyObj{})
											c.AbortWithStatusJSON(http.StatusBadRequest, response)
										} else {
											sectionid, err := strconv.ParseInt(c.Param("sectionid"), 0, 0)
											if err != nil {
												response = helper.BuildErrorResponse("No param group section id was found", err.Error(), helper.EmptyObj{})
												c.AbortWithStatusJSON(http.StatusBadRequest, response)
											} else {
												employeeLeaveCredits, err = b.employeeLeaveCreditService.FindSearchEmployeeLeaveCreditsOffsetByDept(int(limit), int(offset), order, dir, int(year), int(deptId), int(positionId), search, int(groupId), int(employeeid), int(sectionid))
												if err != nil {
													response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
													c.JSON(http.StatusNotFound, response)
												} else {
													response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
													c.JSON(http.StatusOK, response)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveCreditController) CountEmployeeLeaveCreditByDepId(c *gin.Context) {
	var (
		response helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		deptId, err := strconv.ParseInt(c.Param("deptId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
			if err != nil {
				response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
				if err != nil {
					response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					count, err := b.employeeLeaveCreditService.CountEmployeeLeaveCreditByDepId(int(year), int(deptId), int(positionId), int(groupId))
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
}

func (b *employeeLeaveCreditController) FindEmployeeLeaveCreditsOffsetByDept(c *gin.Context) {
	var (
		employeeLeaveCredits []model.SelectEmployeeLeaveCredit
		response             helper.Response
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
					year, err := strconv.ParseInt(c.Param("year"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						deptId, err := strconv.ParseInt(c.Param("deptId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
							if err != nil {
								response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
								c.AbortWithStatusJSON(http.StatusBadRequest, response)
							} else {
								groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
								if err != nil {
									response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
									c.AbortWithStatusJSON(http.StatusBadRequest, response)
								} else {
									employeeid, err := strconv.ParseInt(c.Param("employeeid"), 0, 0)
									if err != nil {
										response = helper.BuildErrorResponse("No param group employee id was found", err.Error(), helper.EmptyObj{})
										c.AbortWithStatusJSON(http.StatusBadRequest, response)
									} else {
										sectionid, err := strconv.ParseInt(c.Param("sectionid"), 0, 0)
										if err != nil {
											response = helper.BuildErrorResponse("No param group section id was found", err.Error(), helper.EmptyObj{})
											c.AbortWithStatusJSON(http.StatusBadRequest, response)
										} else {
											employeeLeaveCredits, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditsOffsetByDept(int(limit), int(offset), order, dir, int(year), int(deptId), int(positionId), int(groupId), int(employeeid), int(sectionid))
											if err != nil {
												response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
												c.JSON(http.StatusNotFound, response)
											} else {
												response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
												c.JSON(http.StatusOK, response)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveCreditController) CountEmployeeLeaveCreditByDivAll(c *gin.Context) {
	var (
		response helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		divId, err := strconv.ParseInt(c.Param("divId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param division id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
			if err != nil {
				response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
				if err != nil {
					response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					count, err := b.employeeLeaveCreditService.CountEmployeeLeaveCreditByDivAll(int(year), int(divId), int(positionId), int(groupId))
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
}

func (b *employeeLeaveCreditController) FindEmployeeLeaveCreditsOffsetByDivAll(c *gin.Context) {
	var (
		employeeLeaveCredits []model.SelectEmployeeLeaveCredit
		response             helper.Response
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
					year, err := strconv.ParseInt(c.Param("year"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						divId, err := strconv.ParseInt(c.Param("divId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param division id was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
							if err != nil {
								response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
								c.AbortWithStatusJSON(http.StatusBadRequest, response)
							} else {
								groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
								if err != nil {
									response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
									c.AbortWithStatusJSON(http.StatusBadRequest, response)
								} else {
									employeeLeaveCredits, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditsOffsetByDivAll(int(limit), int(offset), order, dir, int(year), int(divId), int(positionId), int(groupId))
									if err != nil {
										response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
										c.JSON(http.StatusNotFound, response)
									} else {
										response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
										c.JSON(http.StatusOK, response)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveCreditController) CountSearchEmployeeLeaveCreditByDivIdAll(c *gin.Context) {
	var (
		response helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		search := c.Param("search")
		if search == "" {
			response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			divId, err := strconv.ParseInt(c.Param("divId"), 0, 0)
			if err != nil {
				response = helper.BuildErrorResponse("No param division id was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {

				positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
				if err != nil {
					response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						count, err := b.employeeLeaveCreditService.CountSearchEmployeeLeaveCreditByDivIdAll(int(year), int(divId), int(positionId), search, int(groupId))
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
	}
}

func (b *employeeLeaveCreditController) FindSearchEmployeeLeaveCreditsOffsetByDivIdAll(c *gin.Context) {
	var (
		employeeLeaveCredits []model.SelectEmployeeLeaveCredit
		response             helper.Response
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
						year, err := strconv.ParseInt(c.Param("year"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							divId, err := strconv.ParseInt(c.Param("divId"), 0, 0)
							if err != nil {
								response = helper.BuildErrorResponse("No param division id was found", err.Error(), helper.EmptyObj{})
								c.AbortWithStatusJSON(http.StatusBadRequest, response)
							} else {
								positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
								if err != nil {
									response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
									c.AbortWithStatusJSON(http.StatusBadRequest, response)
								} else {

									groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
									if err != nil {
										response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
										c.AbortWithStatusJSON(http.StatusBadRequest, response)
									} else {
										employeeLeaveCredits, err = b.employeeLeaveCreditService.FindSearchEmployeeLeaveCreditsOffsetByDivIdAll(int(limit), int(offset), order, dir, int(year), int(divId), int(positionId), search, int(groupId))
										if err != nil {
											response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
											c.JSON(http.StatusNotFound, response)
										} else {
											response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
											c.JSON(http.StatusOK, response)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveCreditController) FindEmployeeLeaveCreditsExcEmp(c *gin.Context) {
	var (
		employeeLeaveCredits []model.EmployeeLeaveCredit
		response             helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		empId := c.Param("empId")
		if empId == "" {
			response = helper.BuildErrorResponse("No param empId was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			employeeLeaveCredits, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditsExcEmp(int(year), empId)
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *employeeLeaveCreditController) FindEmployeeLeaveCreditsAll(c *gin.Context) {
	var (
		employeeLeaveCredits []model.EmployeeLeaveCredit
		response             helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveCredits, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditsAll(int(year))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveCreditController) CountEmployeeLeaveCreditByDepAll(c *gin.Context) {
	var (
		response helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		depId, err := strconv.ParseInt(c.Param("depId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
			if err != nil {
				response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
				if err != nil {
					response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					count, err := b.employeeLeaveCreditService.CountEmployeeLeaveCreditByDepAll(int(year), int(depId), int(positionId), int(groupId))
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
}

func (b *employeeLeaveCreditController) FindEmployeeLeaveCreditsOffsetByDepAll(c *gin.Context) {
	var (
		employeeLeaveCredits []model.SelectEmployeeLeaveCredit
		response             helper.Response
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
					year, err := strconv.ParseInt(c.Param("year"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						depId, err := strconv.ParseInt(c.Param("depId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
							if err != nil {
								response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
								c.AbortWithStatusJSON(http.StatusBadRequest, response)
							} else {
								groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
								if err != nil {
									response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
									c.AbortWithStatusJSON(http.StatusBadRequest, response)
								} else {
									employeeLeaveCredits, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditsOffsetByDepAll(int(limit), int(offset), order, dir, int(year), int(depId), int(positionId), int(groupId))
									if err != nil {
										response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
										c.JSON(http.StatusNotFound, response)
									} else {
										response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
										c.JSON(http.StatusOK, response)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (b *employeeLeaveCreditController) CountSearchEmployeeLeaveCreditByDepIdAll(c *gin.Context) {
	var (
		response helper.Response
	)
	year, err := strconv.ParseInt(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		search := c.Param("search")
		if search == "" {
			response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			depId, err := strconv.ParseInt(c.Param("depId"), 0, 0)
			if err != nil {
				response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {

				positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
				if err != nil {
					response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						count, err := b.employeeLeaveCreditService.CountSearchEmployeeLeaveCreditByDepIdAll(int(year), int(depId), int(positionId), search, int(groupId))
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
	}
}

func (b *employeeLeaveCreditController) FindSearchEmployeeLeaveCreditsOffsetByDepIdAll(c *gin.Context) {
	var (
		employeeLeaveCredits []model.SelectEmployeeLeaveCredit
		response             helper.Response
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
						year, err := strconv.ParseInt(c.Param("year"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							depId, err := strconv.ParseInt(c.Param("depId"), 0, 0)
							if err != nil {
								response = helper.BuildErrorResponse("No param department id was found", err.Error(), helper.EmptyObj{})
								c.AbortWithStatusJSON(http.StatusBadRequest, response)
							} else {
								positionId, err := strconv.ParseInt(c.Param("positionId"), 0, 0)
								if err != nil {
									response = helper.BuildErrorResponse("No param position id was found", err.Error(), helper.EmptyObj{})
									c.AbortWithStatusJSON(http.StatusBadRequest, response)
								} else {

									groupId, err := strconv.ParseInt(c.Param("groupId"), 0, 0)
									if err != nil {
										response = helper.BuildErrorResponse("No param group company id was found", err.Error(), helper.EmptyObj{})
										c.AbortWithStatusJSON(http.StatusBadRequest, response)
									} else {
										employeeLeaveCredits, err = b.employeeLeaveCreditService.FindSearchEmployeeLeaveCreditsOffsetByDepIdAll(int(limit), int(offset), order, dir, int(year), int(depId), int(positionId), search, int(groupId))
										if err != nil {
											response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
											c.JSON(http.StatusNotFound, response)
										} else {
											response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
											c.JSON(http.StatusOK, response)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// func (b *employeeLeaveCreditController) SearchEmployeeLeaveCredit(c *gin.Context) {
// 	var (
// 		employeeLeaveCredits []model.EmployeeLeaveCredit
// 		response             helper.Response
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
// 						employeeLeaveCredits, err = b.employeeLeaveCreditService.SearchEmployeeLeaveCredit(int(limit), int(offset), order, dir, search)
// 						if err != nil {
// 							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
// 							c.JSON(http.StatusNotFound, response)
// 						} else {
// 							response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
// 							c.JSON(http.StatusOK, response)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func (b *employeeLeaveCreditController) CountSearchEmployeeLeaveCredit(c *gin.Context) {
// 	var (
// 		response helper.Response
// 	)
// 	search := c.Param("search")
// 	if search == "" {
// 		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 	} else {
// 		count, err := b.employeeLeaveCreditService.CountSearchEmployeeLeaveCredit(search)
// 		if err != nil {
// 			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
// 			c.JSON(http.StatusNotFound, response)
// 		} else {
// 			response = helper.BuildResponse(true, "OK", count)
// 			c.JSON(http.StatusOK, response)
// 		}
// 	}
// }

func (b *employeeLeaveCreditController) FindEmployeeLeaveCreditById(c *gin.Context) {
	var (
		employeeLeaveCredit model.EmployeeLeaveCredit
		response            helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveCredit, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveCredit)
			c.JSON(http.StatusOK, response)
		}
	}
}

// func (b *employeeLeaveCreditController) CountEmployeeLeaveCreditName(c *gin.Context) {
// 	var (
// 		response helper.Response
// 	)
// 	search := c.Param("search")
// 	if search == "" {
// 		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
// 		c.AbortWithStatusJSON(http.StatusBadRequest, response)
// 	} else {
// 		count, err := b.employeeLeaveCreditService.CountEmployeeLeaveCreditName(search)
// 		if err != nil {
// 			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
// 			c.JSON(http.StatusNotFound, response)
// 		} else {
// 			response = helper.BuildResponse(true, "OK", count)
// 			c.JSON(http.StatusOK, response)
// 		}
// 	}
// }

func (b *employeeLeaveCreditController) FindEmployeeLeaveCreditByEmpId(c *gin.Context) {
	var (
		employeeLeaveCredit model.EmployeeLeaveCredit
		response            helper.Response
	)
	empid, err := strconv.ParseUint(c.Param("empid"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param empid was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveCredit, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditByEmpId(uint(empid))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveCredit)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveCreditController) FindAllEmployeeLeaveCreditByEmpId(c *gin.Context) {
	var (
		employeeLeaveCredit []model.EmployeeLeaveCredit
		response            helper.Response
	)
	empid, err := strconv.ParseUint(c.Param("empid"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param empid was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveCredit, err = b.employeeLeaveCreditService.FindAllEmployeeLeaveCreditByEmpId(uint(empid))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveCredit)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveCreditController) FindExcEmployeeLeaveCredit(c *gin.Context) {
	var (
		employeeLeaveCredits []model.EmployeeLeaveCredit
		response             helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveCredits, err = b.employeeLeaveCreditService.FindExcEmployeeLeaveCredit(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveCredits)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveCreditController) InsertEmployeeLeaveCredit(c *gin.Context) {
	var (
		employeeLeaveCredit                model.EmployeeLeaveCredit
		response                           helper.Response
		CreateEmployeeLeaveCreditParameter model.CreateEmployeeLeaveCreditParameter
	)
	err := c.ShouldBindJSON(&CreateEmployeeLeaveCreditParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		employeeLeaveCredit, err = b.employeeLeaveCreditService.InsertEmployeeLeaveCredit(CreateEmployeeLeaveCreditParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register employeeLeaveCredit", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", employeeLeaveCredit)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *employeeLeaveCreditController) UpdateEmployeeLeaveCredit(c *gin.Context) {
	var (
		newData  model.EmployeeLeaveCredit
		oldData  model.EmployeeLeaveCredit
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
			oldData, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.EmployeeLeaveCredit{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employeeLeaveCredit, err := b.employeeLeaveCreditService.UpdateEmployeeLeaveCredit(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update employeeLeaveCredit", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employeeLeaveCredit)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *employeeLeaveCreditController) DeleteEmployeeLeaveCredit(c *gin.Context) {
	var (
		newData  model.EmployeeLeaveCredit
		oldData  model.EmployeeLeaveCredit
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
			oldData, err = b.employeeLeaveCreditService.FindEmployeeLeaveCreditById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.EmployeeLeaveCredit{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				employeeLeaveCredit, err := b.employeeLeaveCreditService.DeleteEmployeeLeaveCredit(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete employeeLeaveCredit", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", employeeLeaveCredit)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
