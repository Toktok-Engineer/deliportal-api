package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyShareholderHistoryController interface {
	CountCompanyShareholderHistoryAll(c *gin.Context)
	FindCompanyShareholderHistorys(c *gin.Context)
	FindCompanyShareholderHistorysOffset(c *gin.Context)
	SearchCompanyShareholderHistory(c *gin.Context)
	CountSearchCompanyShareholderHistory(c *gin.Context)
	FindCompanyShareholderHistoryById(c *gin.Context)
	FindExcCompanyShareholderHistory(c *gin.Context)
	FindCompanyShareholderHistoryByCompanyId(c *gin.Context)
	InsertCompanyShareholderHistory(c *gin.Context)
	UpdateCompanyShareholderHistory(c *gin.Context)
	DeleteCompanyShareholderHistory(c *gin.Context)
}

type companyShareholderHistoryController struct {
	companyShareholderHistoryService service.CompanyShareholderHistoryService
	jwtService                       service.JWTService
}

func NewCompanyShareholderHistoryController(companyShareholderHistoryServ service.CompanyShareholderHistoryService, jwtServ service.JWTService) CompanyShareholderHistoryController {
	return &companyShareholderHistoryController{
		companyShareholderHistoryService: companyShareholderHistoryServ,
		jwtService:                       jwtServ,
	}
}

func (b *companyShareholderHistoryController) CountCompanyShareholderHistoryAll(c *gin.Context) {
	var (
		response helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.companyShareholderHistoryService.CountCompanyShareholderHistoryAll(int(companyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryController) FindCompanyShareholderHistorys(c *gin.Context) {
	var (
		companyShareholderHistorys []model.CompanyShareholderHistory
		response                   helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistorys, err = b.companyShareholderHistoryService.FindCompanyShareholderHistorys(int(companyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholderHistorys)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryController) FindCompanyShareholderHistorysOffset(c *gin.Context) {
	var (
		companyShareholderHistorys []model.SelectCompanyShareholderHistoryParameter
		response                   helper.Response
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
						companyShareholderHistorys, err = b.companyShareholderHistoryService.FindCompanyShareholderHistorysOffset(int(limit), int(offset), order, dir, int(companyId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", companyShareholderHistorys)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *companyShareholderHistoryController) SearchCompanyShareholderHistory(c *gin.Context) {
	var (
		companyShareholderHistorys []model.SelectCompanyShareholderHistoryParameter
		response                   helper.Response
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
							companyShareholderHistorys, err = b.companyShareholderHistoryService.SearchCompanyShareholderHistory(int(limit), int(offset), order, dir, search, int(companyId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", companyShareholderHistorys)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *companyShareholderHistoryController) CountSearchCompanyShareholderHistory(c *gin.Context) {
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
			count, err := b.companyShareholderHistoryService.CountSearchCompanyShareholderHistory(search, int(companyId))
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

func (b *companyShareholderHistoryController) FindCompanyShareholderHistoryById(c *gin.Context) {
	var (
		companyShareholderHistory model.SelectCompanyShareholderHistoryParameter
		response                  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistory, err = b.companyShareholderHistoryService.FindCompanyShareholderHistoryById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholderHistory)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryController) FindExcCompanyShareholderHistory(c *gin.Context) {
	var (
		companyShareholderHistorys []model.SelectCompanyShareholderHistoryParameter
		response                   helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistorys, err = b.companyShareholderHistoryService.FindExcCompanyShareholderHistory(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholderHistorys)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryController) FindCompanyShareholderHistoryByCompanyId(c *gin.Context) {
	var (
		companyShareholderHistorys []model.SelectCompanyShareholderHistoryParameter
		response                   helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param company id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistorys, err = b.companyShareholderHistoryService.FindCompanyShareholderHistoryByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholderHistorys)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryController) InsertCompanyShareholderHistory(c *gin.Context) {
	var (
		companyShareholderHistory                model.CompanyShareholderHistory
		response                                 helper.Response
		CreateCompanyShareholderHistoryParameter model.CreateCompanyShareholderHistoryParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyShareholderHistoryParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistory, err = b.companyShareholderHistoryService.InsertCompanyShareholderHistory(CreateCompanyShareholderHistoryParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyShareholderHistory", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholderHistory)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryController) UpdateCompanyShareholderHistory(c *gin.Context) {
	var (
		newData  model.CompanyShareholderHistory
		oldData  model.SelectCompanyShareholderHistoryParameter
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
			oldData, err = b.companyShareholderHistoryService.FindCompanyShareholderHistoryById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyShareholderHistoryParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyShareholderHistory, err := b.companyShareholderHistoryService.UpdateCompanyShareholderHistory(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyShareholderHistory", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyShareholderHistory)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyShareholderHistoryController) DeleteCompanyShareholderHistory(c *gin.Context) {
	var (
		newData  model.CompanyShareholderHistory
		oldData  model.SelectCompanyShareholderHistoryParameter
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
			oldData, err = b.companyShareholderHistoryService.FindCompanyShareholderHistoryById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyShareholderHistoryParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyShareholderHistory, err := b.companyShareholderHistoryService.DeleteCompanyShareholderHistory(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyShareholderHistory", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyShareholderHistory)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
