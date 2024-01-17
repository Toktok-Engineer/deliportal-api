package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyManagementController interface {
	CountCompanyManagementAll(c *gin.Context)
	FindCompanyManagements(c *gin.Context)
	FindCompanyManagementsOffset(c *gin.Context)
	SearchCompanyManagement(c *gin.Context)
	CountSearchCompanyManagement(c *gin.Context)
	FindCompanyManagementById(c *gin.Context)
	FindExcCompanyManagement(c *gin.Context)
	FindCompanyManagementByCompanyId(c *gin.Context)
	InsertCompanyManagement(c *gin.Context)
	UpdateCompanyManagement(c *gin.Context)
	DeleteCompanyManagement(c *gin.Context)
	ReportCompanyManagement(c *gin.Context)
}

type companyManagementController struct {
	companyManagementService service.CompanyManagementService
	jwtService               service.JWTService
}

func NewCompanyManagementController(companyManagementServ service.CompanyManagementService, jwtServ service.JWTService) CompanyManagementController {
	return &companyManagementController{
		companyManagementService: companyManagementServ,
		jwtService:               jwtServ,
	}
}

func (b *companyManagementController) CountCompanyManagementAll(c *gin.Context) {
	var (
		response helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.companyManagementService.CountCompanyManagementAll(int(companyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementController) FindCompanyManagements(c *gin.Context) {
	var (
		companyManagements []model.CompanyManagement
		response           helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagements, err = b.companyManagementService.FindCompanyManagements(int(companyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagements)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementController) FindCompanyManagementsOffset(c *gin.Context) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
		response           helper.Response
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
					companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						companyManagements, err = b.companyManagementService.FindCompanyManagementsOffset(int(limit), int(offset), order, dir, int(companyId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", companyManagements)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *companyManagementController) SearchCompanyManagement(c *gin.Context) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
		response           helper.Response
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
						companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							companyManagements, err = b.companyManagementService.SearchCompanyManagement(int(limit), int(offset), order, dir, search, int(companyId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", companyManagements)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *companyManagementController) CountSearchCompanyManagement(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.companyManagementService.CountSearchCompanyManagement(search, int(companyId))
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

func (b *companyManagementController) FindCompanyManagementById(c *gin.Context) {
	var (
		companyManagement model.SelectCompanyManagementParameter
		response          helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagement, err = b.companyManagementService.FindCompanyManagementById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagement)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementController) FindExcCompanyManagement(c *gin.Context) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
		response           helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagements, err = b.companyManagementService.FindExcCompanyManagement(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagements)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementController) FindCompanyManagementByCompanyId(c *gin.Context) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
		response           helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param company id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagements, err = b.companyManagementService.FindCompanyManagementByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagements)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementController) InsertCompanyManagement(c *gin.Context) {
	var (
		companyManagement                model.CompanyManagement
		response                         helper.Response
		CreateCompanyManagementParameter model.CreateCompanyManagementParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyManagementParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyManagement, err = b.companyManagementService.InsertCompanyManagement(CreateCompanyManagementParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyManagement", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagement)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementController) UpdateCompanyManagement(c *gin.Context) {
	var (
		newData  model.CompanyManagement
		oldData  model.SelectCompanyManagementParameter
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
			oldData, err = b.companyManagementService.FindCompanyManagementById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyManagementParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyManagement, err := b.companyManagementService.UpdateCompanyManagement(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyManagement", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyManagement)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyManagementController) DeleteCompanyManagement(c *gin.Context) {
	var (
		newData  model.CompanyManagement
		oldData  model.SelectCompanyManagementParameter
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
			oldData, err = b.companyManagementService.FindCompanyManagementById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyManagementParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyManagement, err := b.companyManagementService.DeleteCompanyManagement(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyManagement", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyManagement)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyManagementController) ReportCompanyManagement(c *gin.Context) {
	var (
		companyManagement []model.SelectCompanyManagementReport
		response          helper.Response
	)
	year, err := strconv.ParseUint(c.Param("year"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		groupId, err := strconv.ParseUint(c.Param("groupId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			companyManagement, err = b.companyManagementService.ReportCompanyManagement(int(year), int(groupId))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", companyManagement)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}
