package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyManagementHistoryController interface {
	CountCompanyManagementHistoryAll(c *gin.Context)
	FindCompanyManagementHistorys(c *gin.Context)
	FindCompanyManagementHistorysOffset(c *gin.Context)
	SearchCompanyManagementHistory(c *gin.Context)
	CountSearchCompanyManagementHistory(c *gin.Context)
	FindCompanyManagementHistoryById(c *gin.Context)
	FindExcCompanyManagementHistory(c *gin.Context)
	FindCompanyManagementHistoryByCompanyId(c *gin.Context)
	InsertCompanyManagementHistory(c *gin.Context)
	UpdateCompanyManagementHistory(c *gin.Context)
	DeleteCompanyManagementHistory(c *gin.Context)
}

type companyManagementHistoryController struct {
	companyManagementHistoryService service.CompanyManagementHistoryService
	jwtService                      service.JWTService
}

func NewCompanyManagementHistoryController(companyManagementHistoryServ service.CompanyManagementHistoryService, jwtServ service.JWTService) CompanyManagementHistoryController {
	return &companyManagementHistoryController{
		companyManagementHistoryService: companyManagementHistoryServ,
		jwtService:                      jwtServ,
	}
}

func (b *companyManagementHistoryController) CountCompanyManagementHistoryAll(c *gin.Context) {
	var (
		response helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.companyManagementHistoryService.CountCompanyManagementHistoryAll(int(companyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryController) FindCompanyManagementHistorys(c *gin.Context) {
	var (
		companyManagementHistorys []model.CompanyManagementHistory
		response                  helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistorys, err = b.companyManagementHistoryService.FindCompanyManagementHistorys(int(companyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementHistorys)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryController) FindCompanyManagementHistorysOffset(c *gin.Context) {
	var (
		companyManagementHistorys []model.SelectCompanyManagementHistoryParameter
		response                  helper.Response
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
						companyManagementHistorys, err = b.companyManagementHistoryService.FindCompanyManagementHistorysOffset(int(limit), int(offset), order, dir, int(companyId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", companyManagementHistorys)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *companyManagementHistoryController) SearchCompanyManagementHistory(c *gin.Context) {
	var (
		companyManagementHistorys []model.SelectCompanyManagementHistoryParameter
		response                  helper.Response
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
							companyManagementHistorys, err = b.companyManagementHistoryService.SearchCompanyManagementHistory(int(limit), int(offset), order, dir, search, int(companyId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", companyManagementHistorys)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *companyManagementHistoryController) CountSearchCompanyManagementHistory(c *gin.Context) {
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
			count, err := b.companyManagementHistoryService.CountSearchCompanyManagementHistory(search, int(companyId))
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

func (b *companyManagementHistoryController) FindCompanyManagementHistoryById(c *gin.Context) {
	var (
		companyManagementHistory model.SelectCompanyManagementHistoryParameter
		response                 helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistory, err = b.companyManagementHistoryService.FindCompanyManagementHistoryById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementHistory)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryController) FindExcCompanyManagementHistory(c *gin.Context) {
	var (
		companyManagementHistorys []model.SelectCompanyManagementHistoryParameter
		response                  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistorys, err = b.companyManagementHistoryService.FindExcCompanyManagementHistory(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementHistorys)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryController) FindCompanyManagementHistoryByCompanyId(c *gin.Context) {
	var (
		companyManagementHistorys []model.SelectCompanyManagementHistoryParameter
		response                  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param company id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistorys, err = b.companyManagementHistoryService.FindCompanyManagementHistoryByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementHistorys)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryController) InsertCompanyManagementHistory(c *gin.Context) {
	var (
		companyManagementHistory                model.CompanyManagementHistory
		response                                helper.Response
		CreateCompanyManagementHistoryParameter model.CreateCompanyManagementHistoryParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyManagementHistoryParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistory, err = b.companyManagementHistoryService.InsertCompanyManagementHistory(CreateCompanyManagementHistoryParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyManagementHistory", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementHistory)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryController) UpdateCompanyManagementHistory(c *gin.Context) {
	var (
		newData  model.CompanyManagementHistory
		oldData  model.SelectCompanyManagementHistoryParameter
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
			oldData, err = b.companyManagementHistoryService.FindCompanyManagementHistoryById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyManagementHistoryParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyManagementHistory, err := b.companyManagementHistoryService.UpdateCompanyManagementHistory(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyManagementHistory", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyManagementHistory)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyManagementHistoryController) DeleteCompanyManagementHistory(c *gin.Context) {
	var (
		newData  model.CompanyManagementHistory
		oldData  model.SelectCompanyManagementHistoryParameter
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
			oldData, err = b.companyManagementHistoryService.FindCompanyManagementHistoryById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyManagementHistoryParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyManagementHistory, err := b.companyManagementHistoryService.DeleteCompanyManagementHistory(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyManagementHistory", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyManagementHistory)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
