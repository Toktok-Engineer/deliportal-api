package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImDefaultCirculationController interface {
	CountImDefaultCirculationAll(c *gin.Context)
	FindImDefaultCirculations(c *gin.Context)
	FindImDefaultCirculationsOffset(c *gin.Context)
	SearchImDefaultCirculation(c *gin.Context)
	CountSearchImDefaultCirculation(c *gin.Context)
	FindImDefaultCirculationById(c *gin.Context)
	FindExcImDefaultCirculation(c *gin.Context)
	InsertImDefaultCirculation(c *gin.Context)
	UpdateImDefaultCirculation(c *gin.Context)
	DeleteImDefaultCirculation(c *gin.Context)
}

type imDefaultCirculationController struct {
	imDefaultCirculationService service.ImDefaultCirculationService
	jwtService                  service.JWTService
}

func NewImDefaultCirculationController(imDefaultCirculationServ service.ImDefaultCirculationService, jwtServ service.JWTService) ImDefaultCirculationController {
	return &imDefaultCirculationController{
		imDefaultCirculationService: imDefaultCirculationServ,
		jwtService:                  jwtServ,
	}
}

func (b *imDefaultCirculationController) CountImDefaultCirculationAll(c *gin.Context) {
	var (
		response helper.Response
	)

	companyGroupID, err := strconv.ParseInt(c.Param("companyGroupID"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyGroupID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.imDefaultCirculationService.CountImDefaultCirculationAll(int(companyGroupID))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *imDefaultCirculationController) FindImDefaultCirculations(c *gin.Context) {
	var (
		imdefaultcirculations []model.ImDefaultCirculation
		response              helper.Response
	)
	companyGroupID, err := strconv.ParseInt(c.Param("companyGroupID"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param companyGroupID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		imdefaultcirculations, err = b.imDefaultCirculationService.FindImDefaultCirculations(int(companyGroupID))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", imdefaultcirculations)
			c.JSON(http.StatusOK, response)

		}
	}
}

func (b *imDefaultCirculationController) FindImDefaultCirculationsOffset(c *gin.Context) {
	var (
		imdefaultcirculations []model.SelectImDefaultCirculationParameter
		response              helper.Response
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
						imdefaultcirculations, err = b.imDefaultCirculationService.FindImDefaultCirculationsOffset(int(limit), int(offset), order, dir, int(companyGroupID))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", imdefaultcirculations)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *imDefaultCirculationController) SearchImDefaultCirculation(c *gin.Context) {
	var (
		imdefaultcirculations []model.SelectImDefaultCirculationParameter
		response              helper.Response
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
							imdefaultcirculations, err = b.imDefaultCirculationService.SearchImDefaultCirculation(int(limit), int(offset), order, dir, search, int(companyGroupID))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", imdefaultcirculations)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *imDefaultCirculationController) CountSearchImDefaultCirculation(c *gin.Context) {
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
			count, err := b.imDefaultCirculationService.CountSearchImDefaultCirculation(search, int(companyGroupID))
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

func (b *imDefaultCirculationController) FindImDefaultCirculationById(c *gin.Context) {
	var (
		imDefaultCirculation model.ImDefaultCirculation
		response             helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		imDefaultCirculation, err = b.imDefaultCirculationService.FindImDefaultCirculationById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", imDefaultCirculation)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *imDefaultCirculationController) FindExcImDefaultCirculation(c *gin.Context) {
	var (
		imDefaultCirculations []model.ImDefaultCirculation
		response              helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		imDefaultCirculations, err = b.imDefaultCirculationService.FindExcImDefaultCirculation(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", imDefaultCirculations)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *imDefaultCirculationController) InsertImDefaultCirculation(c *gin.Context) {
	var (
		imDefaultCirculation                model.ImDefaultCirculation
		response                            helper.Response
		CreateImDefaultCirculationParameter model.CreateImDefaultCirculationParameter
	)
	err := c.ShouldBindJSON(&CreateImDefaultCirculationParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		imDefaultCirculation, err = b.imDefaultCirculationService.InsertImDefaultCirculation(CreateImDefaultCirculationParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register im default circulation", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", imDefaultCirculation)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *imDefaultCirculationController) UpdateImDefaultCirculation(c *gin.Context) {
	var (
		newData  model.ImDefaultCirculation
		oldData  model.ImDefaultCirculation
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		err = c.ShouldBindJSON(&newData)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			oldData, err = b.imDefaultCirculationService.FindImDefaultCirculationById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.ImDefaultCirculation{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				imDefaultCirculation, err := b.imDefaultCirculationService.UpdateImDefaultCirculation(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update im default circulation", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", imDefaultCirculation)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *imDefaultCirculationController) DeleteImDefaultCirculation(c *gin.Context) {
	var (
		newData  model.ImDefaultCirculation
		oldData  model.ImDefaultCirculation
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		err = c.ShouldBindJSON(&newData)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			oldData, err = b.imDefaultCirculationService.FindImDefaultCirculationById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.ImDefaultCirculation{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				imDefaultCirculation, err := b.imDefaultCirculationService.DeleteImDefaultCirculation(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete im default circulation", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", imDefaultCirculation)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
