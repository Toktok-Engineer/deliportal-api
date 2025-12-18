package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AsuransiController interface {
	CountAsuransiAll(c *gin.Context)
	FindAsuransis(c *gin.Context)
	FindAsuransisOffset(c *gin.Context)
	SearchAsuransi(c *gin.Context)
	CountSearchAsuransi(c *gin.Context)
	FindAsuransiById(c *gin.Context)
	FindExcAsuransi(c *gin.Context)
	InsertAsuransi(c *gin.Context)
	UpdateAsuransi(c *gin.Context)
	DeleteAsuransi(c *gin.Context)
}

type asuransiController struct {
	asuransiService service.AsuransiService
	jwtService         service.JWTService
}

func NewAsuransiController(asuransiServ service.AsuransiService, jwtServ service.JWTService) AsuransiController {
	return &asuransiController{
		asuransiService: asuransiServ,
		jwtService:         jwtServ,
	}
}

func (b *asuransiController) CountAsuransiAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.asuransiService.CountAsuransiAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *asuransiController) FindAsuransis(c *gin.Context) {
	var (
		vehiclemerks []model.Asuransi
		response     helper.Response
	)
	vehiclemerks, err := b.asuransiService.FindAsuransis()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", vehiclemerks)
		c.JSON(http.StatusOK, response)
	}
}

func (b *asuransiController) FindAsuransisOffset(c *gin.Context) {
	var (
		vehiclemerks []model.Asuransi
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
					vehiclemerks, err = b.asuransiService.FindAsuransisOffset(int(limit), int(offset), order, dir)
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

func (b *asuransiController) SearchAsuransi(c *gin.Context) {
	var (
		vehiclemerks []model.Asuransi
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
						vehiclemerks, err = b.asuransiService.SearchAsuransi(int(limit), int(offset), order, dir, search)
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

func (b *asuransiController) CountSearchAsuransi(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.asuransiService.CountSearchAsuransi(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiController) FindAsuransiById(c *gin.Context) {
	var (
		asuransi model.Asuransi
		response    helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		asuransi, err = b.asuransiService.FindAsuransiById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", asuransi)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiController) FindExcAsuransi(c *gin.Context) {
	var (
		asuransis []model.Asuransi
		response     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		asuransis, err = b.asuransiService.FindExcAsuransi(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", asuransis)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiController) InsertAsuransi(c *gin.Context) {
	var (
		asuransi                model.Asuransi
		response                   helper.Response
		CreateAsuransiParameter model.CreateAsuransiParameter
	)
	err := c.ShouldBindJSON(&CreateAsuransiParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		asuransi, err = b.asuransiService.InsertAsuransi(CreateAsuransiParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register asuransi", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", asuransi)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiController) UpdateAsuransi(c *gin.Context) {
	var (
		newData  model.Asuransi
		oldData  model.Asuransi
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
			oldData, err = b.asuransiService.FindAsuransiById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.Asuransi{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				asuransi, err := b.asuransiService.UpdateAsuransi(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update asuransi", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", asuransi)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *asuransiController) DeleteAsuransi(c *gin.Context) {
	var (
		newData  model.Asuransi
		oldData  model.Asuransi
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
			oldData, err = b.asuransiService.FindAsuransiById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.Asuransi{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				asuransi, err := b.asuransiService.DeleteAsuransi(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete asuransi", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", asuransi)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
