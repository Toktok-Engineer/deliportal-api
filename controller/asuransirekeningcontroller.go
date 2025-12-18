package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AsuransiRekeningController interface {
	CountAsuransiRekeningAll(c *gin.Context)
	FindAsuransiRekenings(c *gin.Context)
	FindAsuransiRekeningsOffset(c *gin.Context)
	SearchAsuransiRekening(c *gin.Context)
	CountSearchAsuransiRekening(c *gin.Context)
	FindAsuransiRekeningById(c *gin.Context)
	FindExcAsuransiRekening(c *gin.Context)
	InsertAsuransiRekening(c *gin.Context)
	UpdateAsuransiRekening(c *gin.Context)
	DeleteAsuransiRekening(c *gin.Context)
}

type asuransiRekeningController struct {
	asuransiRekeningService service.AsuransiRekeningService
	jwtService              service.JWTService
}

func NewAsuransiRekeningController(asuransiRekeningServ service.AsuransiRekeningService, jwtServ service.JWTService) AsuransiRekeningController {
	return &asuransiRekeningController{
		asuransiRekeningService: asuransiRekeningServ,
		jwtService:              jwtServ,
	}
}

func (b *asuransiRekeningController) CountAsuransiRekeningAll(c *gin.Context) {
	var (
		response helper.Response
	)

	AsuransiID, err := strconv.ParseInt(c.Param("asuransi_id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param Asuransi ID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.asuransiRekeningService.CountAsuransiRekeningAll(int(AsuransiID))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiRekeningController) FindAsuransiRekenings(c *gin.Context) {
	var (
		vehiclemerks []model.AsuransiRekening
		response     helper.Response
	)
	vehiclemerks, err := b.asuransiRekeningService.FindAsuransiRekenings()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", vehiclemerks)
		c.JSON(http.StatusOK, response)
	}
}

func (b *asuransiRekeningController) FindAsuransiRekeningsOffset(c *gin.Context) {
	var (
		vehiclemerks []model.AsuransiRekening
		response     helper.Response
	)
	AsuransiID, err := strconv.ParseInt(c.Param("asuransi_id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param Asuransi ID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
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
						vehiclemerks, err = b.asuransiRekeningService.FindAsuransiRekeningsOffset(int(AsuransiID), int(limit), int(offset), order, dir)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", vehiclemerks)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *asuransiRekeningController) SearchAsuransiRekening(c *gin.Context) {
	var (
		vehiclemerks []model.AsuransiRekening
		response     helper.Response
	)

	AsuransiID, err := strconv.ParseInt(c.Param("asuransi_id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param Asuransi ID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
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
							vehiclemerks, err = b.asuransiRekeningService.SearchAsuransiRekening(int(AsuransiID), int(limit), int(offset), order, dir, search)
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", vehiclemerks)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *asuransiRekeningController) CountSearchAsuransiRekening(c *gin.Context) {
	var (
		response helper.Response
	)
	AsuransiID, err := strconv.ParseInt(c.Param("asuransi_id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param Asuransi ID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		search := c.Param("search")
		if search == "" {
			response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.asuransiRekeningService.CountSearchAsuransiRekening(int(AsuransiID), search)
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

func (b *asuransiRekeningController) FindAsuransiRekeningById(c *gin.Context) {
	var (
		asuransiRekening model.AsuransiRekening
		response         helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		asuransiRekening, err = b.asuransiRekeningService.FindAsuransiRekeningById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", asuransiRekening)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiRekeningController) FindExcAsuransiRekening(c *gin.Context) {
	var (
		asuransiRekenings []model.AsuransiRekening
		response          helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		asuransiRekenings, err = b.asuransiRekeningService.FindExcAsuransiRekening(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", asuransiRekenings)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiRekeningController) InsertAsuransiRekening(c *gin.Context) {
	var (
		asuransiRekening                model.AsuransiRekening
		response                        helper.Response
		CreateAsuransiRekeningParameter model.CreateAsuransiRekeningParameter
	)
	err := c.ShouldBindJSON(&CreateAsuransiRekeningParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		asuransiRekening, err = b.asuransiRekeningService.InsertAsuransiRekening(CreateAsuransiRekeningParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register asuransiRekening", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", asuransiRekening)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiRekeningController) UpdateAsuransiRekening(c *gin.Context) {
	var (
		newData  model.AsuransiRekening
		oldData  model.AsuransiRekening
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
			oldData, err = b.asuransiRekeningService.FindAsuransiRekeningById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.AsuransiRekening{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				asuransiRekening, err := b.asuransiRekeningService.UpdateAsuransiRekening(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update asuransiRekening", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", asuransiRekening)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *asuransiRekeningController) DeleteAsuransiRekening(c *gin.Context) {
	var (
		newData  model.AsuransiRekening
		oldData  model.AsuransiRekening
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
			oldData, err = b.asuransiRekeningService.FindAsuransiRekeningById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.AsuransiRekening{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				asuransiRekening, err := b.asuransiRekeningService.DeleteAsuransiRekening(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete asuransiRekening", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", asuransiRekening)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
