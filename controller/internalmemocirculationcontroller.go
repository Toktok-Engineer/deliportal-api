package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InternalMemoCirculationController interface {
	CountInternalMemoCirculationAll(c *gin.Context)
	FindInternalMemoCirculations(c *gin.Context)
	FindInternalMemoCirculationStat(c *gin.Context)
	FindInternalMemoCirculationListStat(c *gin.Context)
	FindInternalMemoCirculationSeq(c *gin.Context)
	FindInternalMemoCirculationsOffset(c *gin.Context)
	SearchInternalMemoCirculation(c *gin.Context)
	CountSearchInternalMemoCirculation(c *gin.Context)
	FindInternalMemoCirculationById(c *gin.Context)
	FindExcInternalMemoCirculation(c *gin.Context)
	FindInternalMemoCirculationByCompanyId(c *gin.Context)
	InsertInternalMemoCirculation(c *gin.Context)
	UpdateInternalMemoCirculation(c *gin.Context)
	DeleteInternalMemoCirculation(c *gin.Context)
	UpdateInternalMemoCirculationApprove(c *gin.Context)
	CountInternalMemoCirculationJoinIM(c *gin.Context)
	FindInternalMemoCirculationsJoinIMOffset(c *gin.Context)
	SearchInternalMemoCirculationJoinIM(c *gin.Context)
	CountSearchInternalMemoCirculationJoinIM(c *gin.Context)
}

type internalMemoCirculationController struct {
	internalMemoCirculationService service.InternalMemoCirculationService
	jwtService                     service.JWTService
}

func NewInternalMemoCirculationController(internalMemoCirculationServ service.InternalMemoCirculationService, jwtServ service.JWTService) InternalMemoCirculationController {
	return &internalMemoCirculationController{
		internalMemoCirculationService: internalMemoCirculationServ,
		jwtService:                     jwtServ,
	}
}

func (b *internalMemoCirculationController) CountInternalMemoCirculationAll(c *gin.Context) {
	var (
		response helper.Response
	)
	internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.internalMemoCirculationService.CountInternalMemoCirculationAll(int(internalMemoId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoCirculationController) FindInternalMemoCirculationSeq(c *gin.Context) {
	var (
		internalMemoCirculation model.SelectInternalMemoCirculationParameter
		response                helper.Response
	)
	internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeID, err := strconv.ParseInt(c.Param("employeeID"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param employeeID was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			internalMemoCirculation, err = b.internalMemoCirculationService.FindInternalMemoCirculationSeq(int(internalMemoId), int(employeeID))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", internalMemoCirculation)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *internalMemoCirculationController) FindInternalMemoCirculationStat(c *gin.Context) {
	var (
		internalMemoCirculation model.SelectInternalMemoCirculationParameter
		response                helper.Response
	)
	internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		cirStat, err := strconv.ParseInt(c.Param("cirStat"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param cirStat was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			internalMemoCirculation, err = b.internalMemoCirculationService.FindInternalMemoCirculationStat(int(internalMemoId), int(cirStat))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", internalMemoCirculation)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *internalMemoCirculationController) FindInternalMemoCirculationListStat(c *gin.Context) {
	var (
		internalMemoCirculations []model.SelectInternalMemoCirculationParameter
		response                 helper.Response
	)
	internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		cirStat, err := strconv.ParseInt(c.Param("cirStat"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param cirStat was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			internalMemoCirculations, err = b.internalMemoCirculationService.FindInternalMemoCirculationListStat(int(internalMemoId), int(cirStat))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", internalMemoCirculations)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *internalMemoCirculationController) FindInternalMemoCirculations(c *gin.Context) {
	var (
		internalMemoCirculations []model.InternalMemoCirculation
		response                 helper.Response
	)
	internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoCirculations, err = b.internalMemoCirculationService.FindInternalMemoCirculations(int(internalMemoId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoCirculations)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoCirculationController) FindInternalMemoCirculationsOffset(c *gin.Context) {
	var (
		internalMemoCirculations []model.SelectInternalMemoCirculationParameter
		response                 helper.Response
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
					internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						internalMemoCirculations, err = b.internalMemoCirculationService.FindInternalMemoCirculationsOffset(int(limit), int(offset), order, dir, int(internalMemoId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", internalMemoCirculations)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *internalMemoCirculationController) SearchInternalMemoCirculation(c *gin.Context) {
	var (
		internalMemoCirculations []model.SelectInternalMemoCirculationParameter
		response                 helper.Response
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
						internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							internalMemoCirculations, err = b.internalMemoCirculationService.SearchInternalMemoCirculation(int(limit), int(offset), order, dir, search, int(internalMemoId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", internalMemoCirculations)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *internalMemoCirculationController) CountSearchInternalMemoCirculation(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.internalMemoCirculationService.CountSearchInternalMemoCirculation(search, int(internalMemoId))
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

func (b *internalMemoCirculationController) FindInternalMemoCirculationById(c *gin.Context) {
	var (
		internalMemoCirculation model.SelectInternalMemoCirculationParameter
		response                helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoCirculation, err = b.internalMemoCirculationService.FindInternalMemoCirculationById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoCirculation)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoCirculationController) FindExcInternalMemoCirculation(c *gin.Context) {
	var (
		internalMemoCirculations []model.SelectInternalMemoCirculationParameter
		response                 helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoCirculations, err = b.internalMemoCirculationService.FindExcInternalMemoCirculation(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoCirculations)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoCirculationController) FindInternalMemoCirculationByCompanyId(c *gin.Context) {
	var (
		internalMemoCirculations []model.SelectInternalMemoCirculationParameter
		response                 helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param company id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoCirculations, err = b.internalMemoCirculationService.FindInternalMemoCirculationByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoCirculations)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoCirculationController) InsertInternalMemoCirculation(c *gin.Context) {
	var (
		internalMemoCirculation                model.InternalMemoCirculation
		response                               helper.Response
		CreateInternalMemoCirculationParameter model.CreateInternalMemoCirculationParameter
	)
	err := c.ShouldBindJSON(&CreateInternalMemoCirculationParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		internalMemoCirculation, err = b.internalMemoCirculationService.InsertInternalMemoCirculation(CreateInternalMemoCirculationParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register internalMemoCirculation", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoCirculation)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoCirculationController) UpdateInternalMemoCirculation(c *gin.Context) {
	var (
		newData  model.InternalMemoCirculation
		oldData  model.SelectInternalMemoCirculationParameter
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
			oldData, err = b.internalMemoCirculationService.FindInternalMemoCirculationById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectInternalMemoCirculationParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemoCirculation, err := b.internalMemoCirculationService.UpdateInternalMemoCirculation(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update internalMemoCirculation", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemoCirculation)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *internalMemoCirculationController) DeleteInternalMemoCirculation(c *gin.Context) {
	var (
		newData  model.InternalMemoCirculation
		oldData  model.SelectInternalMemoCirculationParameter
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
			oldData, err = b.internalMemoCirculationService.FindInternalMemoCirculationById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectInternalMemoCirculationParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemoCirculation, err := b.internalMemoCirculationService.DeleteInternalMemoCirculation(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete internalMemoCirculation", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemoCirculation)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *internalMemoCirculationController) UpdateInternalMemoCirculationApprove(c *gin.Context) {
	var (
		newData  model.InternalMemoCirculation
		oldData  model.SelectInternalMemoCirculationParameter
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
			oldData, err = b.internalMemoCirculationService.FindInternalMemoCirculationById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectInternalMemoCirculationParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemoCirculation, err := b.internalMemoCirculationService.UpdateInternalMemoCirculationApprove(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update internalMemoCirculation", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemoCirculation)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *internalMemoCirculationController) CountInternalMemoCirculationJoinIM(c *gin.Context) {
	var (
		response helper.Response
	)
	employeeID, err := strconv.ParseInt(c.Param("employeeID"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employeeID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.internalMemoCirculationService.CountInternalMemoCirculationJoinIM(int(employeeID))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}
func (b *internalMemoCirculationController) FindInternalMemoCirculationsJoinIMOffset(c *gin.Context) {
	var (
		internalMemoCirculations []model.SelectInternalMemoCirculationJoinIMParameter
		response                 helper.Response
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
					employeeID, err := strconv.ParseInt(c.Param("employeeID"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param employeeID was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						internalMemoCirculations, err = b.internalMemoCirculationService.FindInternalMemoCirculationsJoinIMOffset(int(limit), int(offset), order, dir, int(employeeID))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", internalMemoCirculations)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *internalMemoCirculationController) SearchInternalMemoCirculationJoinIM(c *gin.Context) {
	var (
		internalMemoCirculations []model.SelectInternalMemoCirculationJoinIMParameter
		response                 helper.Response
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
						employeeID, err := strconv.ParseInt(c.Param("employeeID"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param employeeID was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							internalMemoCirculations, err = b.internalMemoCirculationService.SearchInternalMemoCirculationJoinIM(int(limit), int(offset), order, dir, search, int(employeeID))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", internalMemoCirculations)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *internalMemoCirculationController) CountSearchInternalMemoCirculationJoinIM(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeID, err := strconv.ParseInt(c.Param("employeeID"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param employeeID was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.internalMemoCirculationService.CountSearchInternalMemoCirculationJoinIM(search, int(employeeID))
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
