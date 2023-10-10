package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyShareholderHistoryDetailController interface {
	CountCompanyShareholderHistoryDetailAll(c *gin.Context)
	FindCompanyShareholderHistoryDetails(c *gin.Context)
	FindCompanyShareholderHistoryDetailsOffset(c *gin.Context)
	SearchCompanyShareholderHistoryDetail(c *gin.Context)
	CountSearchCompanyShareholderHistoryDetail(c *gin.Context)
	FindCompanyShareholderHistoryDetailByCompanyId(c *gin.Context)
	FindCompanyShareholderHistoryDetailById(c *gin.Context)
	FindExcCompanyShareholderHistoryDetail(c *gin.Context)
	InsertCompanyShareholderHistoryDetail(c *gin.Context)
	UpdateCompanyShareholderHistoryDetail(c *gin.Context)
	DeleteCompanyShareholderHistoryDetail(c *gin.Context)
}

type companyShareholderHistoryDetailController struct {
	companyShareholderHistoryDetailService service.CompanyShareholderHistoryDetailService
	jwtService                             service.JWTService
}

func NewCompanyShareholderHistoryDetailController(companyShareholderHistoryDetailServ service.CompanyShareholderHistoryDetailService, jwtServ service.JWTService) CompanyShareholderHistoryDetailController {
	return &companyShareholderHistoryDetailController{
		companyShareholderHistoryDetailService: companyShareholderHistoryDetailServ,
		jwtService:                             jwtServ,
	}
}

func (b *companyShareholderHistoryDetailController) CountCompanyShareholderHistoryDetailAll(c *gin.Context) {
	var (
		response helper.Response
	)
	companyShareholderHistoryDetailHistoryId, err := strconv.ParseInt(c.Param("companyShareholderHistoryDetailHistoryId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyShareholderHistoryDetailHistoryId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.companyShareholderHistoryDetailService.CountCompanyShareholderHistoryDetailAll(int(companyShareholderHistoryDetailHistoryId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryDetailController) FindCompanyShareholderHistoryDetails(c *gin.Context) {
	var (
		companyShareholderHistoryDetails []model.CompanyShareholderHistoryDetail
		response                         helper.Response
	)
	companyShareholderHistoryDetailHistoryId, err := strconv.ParseInt(c.Param("companyShareholderHistoryDetailHistoryId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyShareholderHistoryDetailHistoryId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistoryDetails, err = b.companyShareholderHistoryDetailService.FindCompanyShareholderHistoryDetails(int(companyShareholderHistoryDetailHistoryId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholderHistoryDetails)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryDetailController) FindCompanyShareholderHistoryDetailsOffset(c *gin.Context) {
	var (
		companyShareholderHistoryDetails []model.SelectCompanyShareholderHistoryDetail
		response                         helper.Response
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
					companyShareholderHistoryDetailHistoryId, err := strconv.ParseInt(c.Param("companyShareholderHistoryDetailHistoryId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param companyShareholderHistoryDetailHistoryId was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						companyShareholderHistoryDetails, err = b.companyShareholderHistoryDetailService.FindCompanyShareholderHistoryDetailsOffset(int(limit), int(offset), order, dir, int(companyShareholderHistoryDetailHistoryId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", companyShareholderHistoryDetails)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *companyShareholderHistoryDetailController) SearchCompanyShareholderHistoryDetail(c *gin.Context) {
	var (
		companyShareholderHistoryDetails []model.SelectCompanyShareholderHistoryDetail
		response                         helper.Response
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
						companyShareholderHistoryDetailHistoryId, err := strconv.ParseInt(c.Param("companyShareholderHistoryDetailHistoryId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param companyShareholderHistoryDetailHistoryId was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							companyShareholderHistoryDetails, err = b.companyShareholderHistoryDetailService.SearchCompanyShareholderHistoryDetail(int(limit), int(offset), order, dir, search, int(companyShareholderHistoryDetailHistoryId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", companyShareholderHistoryDetails)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *companyShareholderHistoryDetailController) CountSearchCompanyShareholderHistoryDetail(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistoryDetailHistoryId, err := strconv.ParseInt(c.Param("companyShareholderHistoryDetailHistoryId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param companyShareholderHistoryDetailHistoryId was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.companyShareholderHistoryDetailService.CountSearchCompanyShareholderHistoryDetail(search, int(companyShareholderHistoryDetailHistoryId))
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

func (b *companyShareholderHistoryDetailController) FindCompanyShareholderHistoryDetailByCompanyId(c *gin.Context) {
	var (
		companyShareholderHistoryDetails []model.CompanyShareholderHistoryDetail
		response                         helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistoryDetails, err = b.companyShareholderHistoryDetailService.FindCompanyShareholderHistoryDetailByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholderHistoryDetails)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryDetailController) FindCompanyShareholderHistoryDetailById(c *gin.Context) {
	var (
		companyShareholderHistoryDetail model.CompanyShareholderHistoryDetail
		response                        helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistoryDetail, err = b.companyShareholderHistoryDetailService.FindCompanyShareholderHistoryDetailById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholderHistoryDetail)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryDetailController) FindExcCompanyShareholderHistoryDetail(c *gin.Context) {
	var (
		companyShareholderHistoryDetails []model.CompanyShareholderHistoryDetail
		response                         helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistoryDetails, err = b.companyShareholderHistoryDetailService.FindExcCompanyShareholderHistoryDetail(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholderHistoryDetails)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryDetailController) InsertCompanyShareholderHistoryDetail(c *gin.Context) {
	var (
		companyShareholderHistoryDetail                model.CompanyShareholderHistoryDetail
		response                                       helper.Response
		CreateCompanyShareholderHistoryDetailParameter model.CreateCompanyShareholderHistoryDetailParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyShareholderHistoryDetailParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyShareholderHistoryDetail, err = b.companyShareholderHistoryDetailService.InsertCompanyShareholderHistoryDetail(CreateCompanyShareholderHistoryDetailParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyShareholderHistoryDetail", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholderHistoryDetail)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderHistoryDetailController) UpdateCompanyShareholderHistoryDetail(c *gin.Context) {
	var (
		newData  model.CompanyShareholderHistoryDetail
		oldData  model.CompanyShareholderHistoryDetail
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
			oldData, err = b.companyShareholderHistoryDetailService.FindCompanyShareholderHistoryDetailById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyShareholderHistoryDetail{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyShareholderHistoryDetail, err := b.companyShareholderHistoryDetailService.UpdateCompanyShareholderHistoryDetail(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyShareholderHistoryDetail", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyShareholderHistoryDetail)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyShareholderHistoryDetailController) DeleteCompanyShareholderHistoryDetail(c *gin.Context) {
	var (
		newData  model.CompanyShareholderHistoryDetail
		oldData  model.CompanyShareholderHistoryDetail
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
			oldData, err = b.companyShareholderHistoryDetailService.FindCompanyShareholderHistoryDetailById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyShareholderHistoryDetail{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyShareholderHistoryDetail, err := b.companyShareholderHistoryDetailService.DeleteCompanyShareholderHistoryDetail(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyShareholderHistoryDetail", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyShareholderHistoryDetail)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
