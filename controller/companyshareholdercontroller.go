package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyShareholderController interface {
	CountCompanyShareholderAll(c *gin.Context)
	FindCompanyShareholders(c *gin.Context)
	FindCompanyShareholdersOffset(c *gin.Context)
	SearchCompanyShareholder(c *gin.Context)
	CountSearchCompanyShareholder(c *gin.Context)
	FindCompanyShareholderByCompanyId(c *gin.Context)
	FindCompanyShareholderById(c *gin.Context)
	FindExcCompanyShareholder(c *gin.Context)
	InsertCompanyShareholder(c *gin.Context)
	UpdateCompanyShareholder(c *gin.Context)
	DeleteCompanyShareholder(c *gin.Context)
	ReportCompanyShareholder(c *gin.Context)
}

type companyShareholderController struct {
	companyShareholderService service.CompanyShareholderService
	jwtService                service.JWTService
}

func NewCompanyShareholderController(companyShareholderServ service.CompanyShareholderService, jwtServ service.JWTService) CompanyShareholderController {
	return &companyShareholderController{
		companyShareholderService: companyShareholderServ,
		jwtService:                jwtServ,
	}
}

func (b *companyShareholderController) CountCompanyShareholderAll(c *gin.Context) {
	var (
		response helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.companyShareholderService.CountCompanyShareholderAll(int(companyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderController) FindCompanyShareholders(c *gin.Context) {
	var (
		companyShareholders []model.CompanyShareholder
		response            helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholders, err = b.companyShareholderService.FindCompanyShareholders(int(companyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholders)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderController) FindCompanyShareholdersOffset(c *gin.Context) {
	var (
		companyShareholders []model.SelectCompanyShareholder
		response            helper.Response
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
						companyShareholders, err = b.companyShareholderService.FindCompanyShareholdersOffset(int(limit), int(offset), order, dir, int(companyId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", companyShareholders)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *companyShareholderController) SearchCompanyShareholder(c *gin.Context) {
	var (
		companyShareholders []model.SelectCompanyShareholder
		response            helper.Response
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
							companyShareholders, err = b.companyShareholderService.SearchCompanyShareholder(int(limit), int(offset), order, dir, search, int(companyId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", companyShareholders)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *companyShareholderController) CountSearchCompanyShareholder(c *gin.Context) {
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
			count, err := b.companyShareholderService.CountSearchCompanyShareholder(search, int(companyId))
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

func (b *companyShareholderController) FindCompanyShareholderByCompanyId(c *gin.Context) {
	var (
		companyShareholders []model.CompanyShareholder
		response            helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholders, err = b.companyShareholderService.FindCompanyShareholderByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholders)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderController) FindCompanyShareholderById(c *gin.Context) {
	var (
		companyShareholder model.CompanyShareholder
		response           helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholder, err = b.companyShareholderService.FindCompanyShareholderById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholder)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderController) FindExcCompanyShareholder(c *gin.Context) {
	var (
		companyShareholders []model.CompanyShareholder
		response            helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholders, err = b.companyShareholderService.FindExcCompanyShareholder(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholders)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderController) InsertCompanyShareholder(c *gin.Context) {
	var (
		companyShareholder                model.CompanyShareholder
		response                          helper.Response
		CreateCompanyShareholderParameter model.CreateCompanyShareholderParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyShareholderParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyShareholder, err = b.companyShareholderService.InsertCompanyShareholder(CreateCompanyShareholderParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyShareholder", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholder)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderController) UpdateCompanyShareholder(c *gin.Context) {
	var (
		newData  model.CompanyShareholder
		oldData  model.CompanyShareholder
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
			oldData, err = b.companyShareholderService.FindCompanyShareholderById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyShareholder{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyShareholder, err := b.companyShareholderService.UpdateCompanyShareholder(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyShareholder", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyShareholder)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyShareholderController) DeleteCompanyShareholder(c *gin.Context) {
	var (
		newData  model.CompanyShareholder
		oldData  model.CompanyShareholder
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
			oldData, err = b.companyShareholderService.FindCompanyShareholderById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyShareholder{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyShareholder, err := b.companyShareholderService.DeleteCompanyShareholder(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyShareholder", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyShareholder)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyShareholderController) ReportCompanyShareholder(c *gin.Context) {
	var (
		companyShareholder []model.SelectCompanyShareholderReport
		response           helper.Response
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
			companyShareholder, err = b.companyShareholderService.ReportCompanyShareholder(int(year), int(groupId))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", companyShareholder)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}
