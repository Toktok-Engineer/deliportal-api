package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyGroupCompanyController interface {
	CountCompanyGroupCompanyAll(c *gin.Context)
	FindCompanyGroupCompanys(c *gin.Context)
	FindCompanyGroupCompanysOffset(c *gin.Context)
	SearchCompanyGroupCompany(c *gin.Context)
	CountSearchCompanyGroupCompany(c *gin.Context)
	FindCompanyGroupCompanyById(c *gin.Context)
	InsertCompanyGroupCompany(c *gin.Context)
	UpdateCompanyGroupCompany(c *gin.Context)
	DeleteCompanyGroupCompany(c *gin.Context)
}

type companyGroupCompanyController struct {
	companyGroupCompanyService service.CompanyGroupCompanyService
	jwtService                 service.JWTService
}

func NewCompanyGroupCompanyController(companyGroupCompanyServ service.CompanyGroupCompanyService, jwtServ service.JWTService) CompanyGroupCompanyController {
	return &companyGroupCompanyController{
		companyGroupCompanyService: companyGroupCompanyServ,
		jwtService:                 jwtServ,
	}
}

func (b *companyGroupCompanyController) CountCompanyGroupCompanyAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	companyGroupID, err := strconv.ParseInt(c.Param("companyGroupID"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyGroupID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err = b.companyGroupCompanyService.CountCompanyGroupCompanyAll(int(companyGroupID))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyGroupCompanyController) FindCompanyGroupCompanys(c *gin.Context) {
	var (
		companyGroupCompanys []model.SelectCompanyGroupCompanyParameter
		response             helper.Response
	)
	companyGroupCompanys, err := b.companyGroupCompanyService.FindCompanyGroupCompanys()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", companyGroupCompanys)
		c.JSON(http.StatusOK, response)
	}
}

func (b *companyGroupCompanyController) FindCompanyGroupCompanysOffset(c *gin.Context) {
	var (
		companyGroupCompanys []model.SelectCompanyGroupCompanyParameter
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
					companyGroupID, err := strconv.ParseInt(c.Param("companyGroupID"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param companyGroupID was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						companyGroupCompanys, err = b.companyGroupCompanyService.FindCompanyGroupCompanysOffset(int(limit), int(offset), order, dir, int(companyGroupID))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", companyGroupCompanys)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *companyGroupCompanyController) SearchCompanyGroupCompany(c *gin.Context) {
	var (
		companyGroupCompanys []model.SelectCompanyGroupCompanyParameter
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
						companyGroupID, err := strconv.ParseInt(c.Param("companyGroupID"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param companyGroupID was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							companyGroupCompanys, err = b.companyGroupCompanyService.SearchCompanyGroupCompany(int(limit), int(offset), order, dir, search, int(companyGroupID))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", companyGroupCompanys)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *companyGroupCompanyController) CountSearchCompanyGroupCompany(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyGroupID, err := strconv.ParseInt(c.Param("companyGroupID"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param companyGroupID was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.companyGroupCompanyService.CountSearchCompanyGroupCompany(search, int(companyGroupID))
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

func (b *companyGroupCompanyController) FindCompanyGroupCompanyById(c *gin.Context) {
	var (
		companyGroupCompany model.SelectCompanyGroupCompanyParameter
		response            helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyGroupCompany, err = b.companyGroupCompanyService.FindCompanyGroupCompanyById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyGroupCompany)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyGroupCompanyController) InsertCompanyGroupCompany(c *gin.Context) {
	var (
		companyGroupCompany                model.CompanyGroupCompany
		response                           helper.Response
		CreateCompanyGroupCompanyParameter model.CreateCompanyGroupCompanyParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyGroupCompanyParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyGroupCompany, err = b.companyGroupCompanyService.InsertCompanyGroupCompany(CreateCompanyGroupCompanyParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyGroupCompany", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyGroupCompany)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyGroupCompanyController) UpdateCompanyGroupCompany(c *gin.Context) {
	var (
		newData  model.CompanyGroupCompany
		oldData  model.SelectCompanyGroupCompanyParameter
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
			oldData, err = b.companyGroupCompanyService.FindCompanyGroupCompanyById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyGroupCompanyParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyGroupCompany, err := b.companyGroupCompanyService.UpdateCompanyGroupCompany(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyGroupCompany", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyGroupCompany)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyGroupCompanyController) DeleteCompanyGroupCompany(c *gin.Context) {
	var (
		newData  model.CompanyGroupCompany
		oldData  model.SelectCompanyGroupCompanyParameter
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
			oldData, err = b.companyGroupCompanyService.FindCompanyGroupCompanyById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyGroupCompanyParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyGroupCompany, err := b.companyGroupCompanyService.DeleteCompanyGroupCompany(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyGroupCompany", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyGroupCompany)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
