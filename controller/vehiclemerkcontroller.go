package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehicleMerkController interface {
	CountVehicleMerkAll(c *gin.Context)
	FindVehicleMerks(c *gin.Context)
	FindVehicleMerksOffset(c *gin.Context)
	SearchVehicleMerk(c *gin.Context)
	CountSearchVehicleMerk(c *gin.Context)
	FindVehicleMerkById(c *gin.Context)
	FindExcVehicleMerk(c *gin.Context)
	InsertVehicleMerk(c *gin.Context)
	UpdateVehicleMerk(c *gin.Context)
	DeleteVehicleMerk(c *gin.Context)
}

type vehicleMerkController struct {
	vehicleMerkService service.VehicleMerkService
	jwtService         service.JWTService
}

func NewVehicleMerkController(vehicleMerkServ service.VehicleMerkService, jwtServ service.JWTService) VehicleMerkController {
	return &vehicleMerkController{
		vehicleMerkService: vehicleMerkServ,
		jwtService:         jwtServ,
	}
}

func (b *vehicleMerkController) CountVehicleMerkAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.vehicleMerkService.CountVehicleMerkAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *vehicleMerkController) FindVehicleMerks(c *gin.Context) {
	var (
		vehiclemerks []model.VehicleMerk
		response     helper.Response
	)
	vehiclemerks, err := b.vehicleMerkService.FindVehicleMerks()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", vehiclemerks)
		c.JSON(http.StatusOK, response)
	}
}

func (b *vehicleMerkController) FindVehicleMerksOffset(c *gin.Context) {
	var (
		vehiclemerks []model.VehicleMerk
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
					vehiclemerks, err = b.vehicleMerkService.FindVehicleMerksOffset(int(limit), int(offset), order, dir)
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

func (b *vehicleMerkController) SearchVehicleMerk(c *gin.Context) {
	var (
		vehiclemerks []model.VehicleMerk
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
						vehiclemerks, err = b.vehicleMerkService.SearchVehicleMerk(int(limit), int(offset), order, dir, search)
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

func (b *vehicleMerkController) CountSearchVehicleMerk(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.vehicleMerkService.CountSearchVehicleMerk(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleMerkController) FindVehicleMerkById(c *gin.Context) {
	var (
		vehicleMerk model.VehicleMerk
		response    helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		vehicleMerk, err = b.vehicleMerkService.FindVehicleMerkById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleMerk)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleMerkController) FindExcVehicleMerk(c *gin.Context) {
	var (
		vehicleMerks []model.VehicleMerk
		response     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		vehicleMerks, err = b.vehicleMerkService.FindExcVehicleMerk(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleMerks)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleMerkController) InsertVehicleMerk(c *gin.Context) {
	var (
		vehicleMerk                model.VehicleMerk
		response                   helper.Response
		CreateVehicleMerkParameter model.CreateVehicleMerkParameter
	)
	err := c.ShouldBindJSON(&CreateVehicleMerkParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		vehicleMerk, err = b.vehicleMerkService.InsertVehicleMerk(CreateVehicleMerkParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register vehicle merk", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleMerk)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleMerkController) UpdateVehicleMerk(c *gin.Context) {
	var (
		newData  model.VehicleMerk
		oldData  model.VehicleMerk
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
			oldData, err = b.vehicleMerkService.FindVehicleMerkById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.VehicleMerk{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				vehicleMerk, err := b.vehicleMerkService.UpdateVehicleMerk(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update vehicle merk", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", vehicleMerk)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *vehicleMerkController) DeleteVehicleMerk(c *gin.Context) {
	var (
		newData  model.VehicleMerk
		oldData  model.VehicleMerk
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
			oldData, err = b.vehicleMerkService.FindVehicleMerkById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.VehicleMerk{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				vehicleMerk, err := b.vehicleMerkService.DeleteVehicleMerk(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete vehicle merk", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", vehicleMerk)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
