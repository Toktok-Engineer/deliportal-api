package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyAkteController interface {
	CountCompanyAkteAll(c *gin.Context)
	FindCompanyAktes(c *gin.Context)
	FindCompanyAktesOffset(c *gin.Context)
	SearchCompanyAkte(c *gin.Context)
	CountSearchCompanyAkte(c *gin.Context)
	FindCompanyAkteById(c *gin.Context)
	FindCompanyAkteByYear(c *gin.Context)
	FindExcCompanyAkteByYear(c *gin.Context)
	FindExcCompanyAkte(c *gin.Context)
	InsertCompanyAkte(c *gin.Context)
	UpdateCompanyAkte(c *gin.Context)
	DeleteCompanyAkte(c *gin.Context)
}

type companyAkteController struct {
	companyAkteService service.CompanyAkteService
	jwtService         service.JWTService
}

func NewCompanyAkteController(companyAkteServ service.CompanyAkteService, jwtServ service.JWTService) CompanyAkteController {
	return &companyAkteController{
		companyAkteService: companyAkteServ,
		jwtService:         jwtServ,
	}
}

func (b *companyAkteController) CountCompanyAkteAll(c *gin.Context) {
	var (
		response helper.Response
	)

	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.companyAkteService.CountCompanyAkteAll(int(companyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyAkteController) FindCompanyAktes(c *gin.Context) {
	var (
		companyAktes []model.SelectCompanyAkteParameter
		response     helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyAktes, err = b.companyAkteService.FindCompanyAktes(int(companyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyAktes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyAkteController) FindCompanyAktesOffset(c *gin.Context) {
	var (
		companyAktes []model.SelectCompanyAkteParameter
		response     helper.Response
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
						companyAktes, err = b.companyAkteService.FindCompanyAktesOffset(int(limit), int(offset), order, dir, int(companyId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", companyAktes)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *companyAkteController) SearchCompanyAkte(c *gin.Context) {
	var (
		companyAktes []model.SelectCompanyAkteParameter
		response     helper.Response
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
							companyAktes, err = b.companyAkteService.SearchCompanyAkte(int(limit), int(offset), order, dir, search, int(companyId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", companyAktes)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *companyAkteController) CountSearchCompanyAkte(c *gin.Context) {
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
			count, err := b.companyAkteService.CountSearchCompanyAkte(search, int(companyId))
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

func (b *companyAkteController) FindCompanyAkteById(c *gin.Context) {
	var (
		companyAkte model.SelectCompanyAkteParameter
		response    helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyAkte, err = b.companyAkteService.FindCompanyAkteById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyAkte)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyAkteController) FindCompanyAkteByYear(c *gin.Context) {
	var (
		companyAktes []model.SelectCompanyAkteParameter
		response     helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		year, err := strconv.ParseUint(c.Param("year"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			companyAktes, err = b.companyAkteService.FindCompanyAkteByYear(uint(companyId), uint(year))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", companyAktes)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *companyAkteController) FindExcCompanyAkteByYear(c *gin.Context) {
	var (
		companyAktes []model.SelectCompanyAkteParameter
		response     helper.Response
	)
	companyId, err := strconv.ParseInt(c.Param("companyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		year, err := strconv.ParseUint(c.Param("year"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param year was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			companyAkteId, err := strconv.ParseUint(c.Param("companyAkteId"), 0, 0)
			if err != nil {
				response = helper.BuildErrorResponse("No param companyAkteId was found", err.Error(), helper.EmptyObj{})
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
			} else {
				companyAktes, err = b.companyAkteService.FindExcCompanyAkteByYear(uint(companyId), uint(year), uint(companyAkteId))
				if err != nil {
					response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
					c.JSON(http.StatusNotFound, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyAktes)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyAkteController) FindExcCompanyAkte(c *gin.Context) {
	var (
		companyAktes []model.SelectCompanyAkteParameter
		response     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyAktes, err = b.companyAkteService.FindExcCompanyAkte(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyAktes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyAkteController) InsertCompanyAkte(c *gin.Context) {
	var (
		companyAkte                model.CompanyAkte
		response                   helper.Response
		CreateCompanyAkteParameter model.CreateCompanyAkteParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyAkteParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyAkte, err = b.companyAkteService.InsertCompanyAkte(CreateCompanyAkteParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyAkte", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyAkte)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyAkteController) UpdateCompanyAkte(c *gin.Context) {
	var (
		newData  model.CompanyAkte
		oldData  model.SelectCompanyAkteParameter
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
			oldData, err = b.companyAkteService.FindCompanyAkteById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyAkteParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyAkte, err := b.companyAkteService.UpdateCompanyAkte(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyAkte", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyAkte)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyAkteController) DeleteCompanyAkte(c *gin.Context) {
	var (
		newData  model.CompanyAkte
		oldData  model.SelectCompanyAkteParameter
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
			oldData, err = b.companyAkteService.FindCompanyAkteById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyAkteParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyAkte, err := b.companyAkteService.DeleteCompanyAkte(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyAkte", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyAkte)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
