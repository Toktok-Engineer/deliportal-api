package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyManagementHistoryDetailController interface {
	CountCompanyManagementHistoryDetailAll(c *gin.Context)
	FindCompanyManagementHistoryDetails(c *gin.Context)
	FindCompanyManagementHistoryDetailsOffset(c *gin.Context)
	SearchCompanyManagementHistoryDetail(c *gin.Context)
	CountSearchCompanyManagementHistoryDetail(c *gin.Context)
	FindCompanyManagementHistoryDetailByCompanyId(c *gin.Context)
	FindCompanyManagementHistoryDetailById(c *gin.Context)
	FindExcCompanyManagementHistoryDetail(c *gin.Context)
	InsertCompanyManagementHistoryDetail(c *gin.Context)
	UpdateCompanyManagementHistoryDetail(c *gin.Context)
	DeleteCompanyManagementHistoryDetail(c *gin.Context)
}

type companyManagementHistoryDetailController struct {
	companyManagementHistoryDetailService service.CompanyManagementHistoryDetailService
	jwtService                            service.JWTService
}

func NewCompanyManagementHistoryDetailController(companyManagementHistoryDetailServ service.CompanyManagementHistoryDetailService, jwtServ service.JWTService) CompanyManagementHistoryDetailController {
	return &companyManagementHistoryDetailController{
		companyManagementHistoryDetailService: companyManagementHistoryDetailServ,
		jwtService:                            jwtServ,
	}
}

func (b *companyManagementHistoryDetailController) CountCompanyManagementHistoryDetailAll(c *gin.Context) {
	var (
		response helper.Response
	)
	companyManagementHistoryDetailHistoryId, err := strconv.ParseInt(c.Param("companyManagementHistoryDetailHistoryId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyManagementHistoryDetailHistoryId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.companyManagementHistoryDetailService.CountCompanyManagementHistoryDetailAll(int(companyManagementHistoryDetailHistoryId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryDetailController) FindCompanyManagementHistoryDetails(c *gin.Context) {
	var (
		companyManagementHistoryDetails []model.CompanyManagementHistoryDetail
		response                        helper.Response
	)
	companyManagementHistoryDetailHistoryId, err := strconv.ParseInt(c.Param("companyManagementHistoryDetailHistoryId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyManagementHistoryDetailHistoryId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistoryDetails, err = b.companyManagementHistoryDetailService.FindCompanyManagementHistoryDetails(int(companyManagementHistoryDetailHistoryId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementHistoryDetails)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryDetailController) FindCompanyManagementHistoryDetailsOffset(c *gin.Context) {
	var (
		companyManagementHistoryDetails []model.SelectCompanyManagementHistoryDetail
		response                        helper.Response
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
					companyManagementHistoryDetailHistoryId, err := strconv.ParseInt(c.Param("companyManagementHistoryDetailHistoryId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param companyManagementHistoryDetailHistoryId was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						companyManagementHistoryDetails, err = b.companyManagementHistoryDetailService.FindCompanyManagementHistoryDetailsOffset(int(limit), int(offset), order, dir, int(companyManagementHistoryDetailHistoryId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", companyManagementHistoryDetails)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *companyManagementHistoryDetailController) SearchCompanyManagementHistoryDetail(c *gin.Context) {
	var (
		companyManagementHistoryDetails []model.SelectCompanyManagementHistoryDetail
		response                        helper.Response
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
						companyManagementHistoryDetailHistoryId, err := strconv.ParseInt(c.Param("companyManagementHistoryDetailHistoryId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param companyManagementHistoryDetailHistoryId was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							companyManagementHistoryDetails, err = b.companyManagementHistoryDetailService.SearchCompanyManagementHistoryDetail(int(limit), int(offset), order, dir, search, int(companyManagementHistoryDetailHistoryId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", companyManagementHistoryDetails)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *companyManagementHistoryDetailController) CountSearchCompanyManagementHistoryDetail(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistoryDetailHistoryId, err := strconv.ParseInt(c.Param("companyManagementHistoryDetailHistoryId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param companyManagementHistoryDetailHistoryId was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.companyManagementHistoryDetailService.CountSearchCompanyManagementHistoryDetail(search, int(companyManagementHistoryDetailHistoryId))
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

func (b *companyManagementHistoryDetailController) FindCompanyManagementHistoryDetailByCompanyId(c *gin.Context) {
	var (
		companyManagementHistoryDetails []model.CompanyManagementHistoryDetail
		response                        helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistoryDetails, err = b.companyManagementHistoryDetailService.FindCompanyManagementHistoryDetailByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementHistoryDetails)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryDetailController) FindCompanyManagementHistoryDetailById(c *gin.Context) {
	var (
		companyManagementHistoryDetail model.CompanyManagementHistoryDetail
		response                       helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistoryDetail, err = b.companyManagementHistoryDetailService.FindCompanyManagementHistoryDetailById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementHistoryDetail)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryDetailController) FindExcCompanyManagementHistoryDetail(c *gin.Context) {
	var (
		companyManagementHistoryDetails []model.CompanyManagementHistoryDetail
		response                        helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistoryDetails, err = b.companyManagementHistoryDetailService.FindExcCompanyManagementHistoryDetail(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementHistoryDetails)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryDetailController) InsertCompanyManagementHistoryDetail(c *gin.Context) {
	var (
		companyManagementHistoryDetail                model.CompanyManagementHistoryDetail
		response                                      helper.Response
		CreateCompanyManagementHistoryDetailParameter model.CreateCompanyManagementHistoryDetailParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyManagementHistoryDetailParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyManagementHistoryDetail, err = b.companyManagementHistoryDetailService.InsertCompanyManagementHistoryDetail(CreateCompanyManagementHistoryDetailParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyManagementHistoryDetail", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementHistoryDetail)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementHistoryDetailController) UpdateCompanyManagementHistoryDetail(c *gin.Context) {
	var (
		newData  model.CompanyManagementHistoryDetail
		oldData  model.CompanyManagementHistoryDetail
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
			oldData, err = b.companyManagementHistoryDetailService.FindCompanyManagementHistoryDetailById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyManagementHistoryDetail{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyManagementHistoryDetail, err := b.companyManagementHistoryDetailService.UpdateCompanyManagementHistoryDetail(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyManagementHistoryDetail", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyManagementHistoryDetail)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyManagementHistoryDetailController) DeleteCompanyManagementHistoryDetail(c *gin.Context) {
	var (
		newData  model.CompanyManagementHistoryDetail
		oldData  model.CompanyManagementHistoryDetail
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
			oldData, err = b.companyManagementHistoryDetailService.FindCompanyManagementHistoryDetailById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyManagementHistoryDetail{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyManagementHistoryDetail, err := b.companyManagementHistoryDetailService.DeleteCompanyManagementHistoryDetail(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyManagementHistoryDetail", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyManagementHistoryDetail)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
