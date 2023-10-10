package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehicleTypeController interface {
	CountVehicleTypeAll(c *gin.Context)
	FindVehicleTypes(c *gin.Context)
	FindVehicleTypesOffset(c *gin.Context)
	SearchVehicleType(c *gin.Context)
	CountSearchVehicleType(c *gin.Context)
	FindVehicleTypeById(c *gin.Context)
	FindExcVehicleType(c *gin.Context)
	InsertVehicleType(c *gin.Context)
	UpdateVehicleType(c *gin.Context)
	DeleteVehicleType(c *gin.Context)
}

type vehicleTypeController struct {
	vehicleTypeService service.VehicleTypeService
	jwtService         service.JWTService
}

func NewVehicleTypeController(vehicleTypeServ service.VehicleTypeService, jwtServ service.JWTService) VehicleTypeController {
	return &vehicleTypeController{
		vehicleTypeService: vehicleTypeServ,
		jwtService:         jwtServ,
	}
}

func (b *vehicleTypeController) CountVehicleTypeAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.vehicleTypeService.CountVehicleTypeAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *vehicleTypeController) FindVehicleTypes(c *gin.Context) {
	var (
		vehicletypes []model.VehicleType
		response     helper.Response
	)
	vehicletypes, err := b.vehicleTypeService.FindVehicleTypes()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", vehicletypes)
		c.JSON(http.StatusOK, response)
	}
}

func (b *vehicleTypeController) FindVehicleTypesOffset(c *gin.Context) {
	var (
		vehicletypes []model.VehicleType
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
					vehicletypes, err = b.vehicleTypeService.FindVehicleTypesOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", vehicletypes)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *vehicleTypeController) SearchVehicleType(c *gin.Context) {
	var (
		vehicletypes []model.VehicleType
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
						vehicletypes, err = b.vehicleTypeService.SearchVehicleType(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", vehicletypes)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *vehicleTypeController) CountSearchVehicleType(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.vehicleTypeService.CountSearchVehicleType(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleTypeController) FindVehicleTypeById(c *gin.Context) {
	var (
		vehicleType model.VehicleType
		response    helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		vehicleType, err = b.vehicleTypeService.FindVehicleTypeById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleTypeController) FindExcVehicleType(c *gin.Context) {
	var (
		vehicleTypes []model.VehicleType
		response     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		vehicleTypes, err = b.vehicleTypeService.FindExcVehicleType(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleTypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleTypeController) InsertVehicleType(c *gin.Context) {
	var (
		vehicleType                model.VehicleType
		response                   helper.Response
		CreateVehicleTypeParameter model.CreateVehicleTypeParameter
	)
	err := c.ShouldBindJSON(&CreateVehicleTypeParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		vehicleType, err = b.vehicleTypeService.InsertVehicleType(CreateVehicleTypeParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register vehicle type", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleTypeController) UpdateVehicleType(c *gin.Context) {
	var (
		newData  model.VehicleType
		oldData  model.VehicleType
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
			oldData, err = b.vehicleTypeService.FindVehicleTypeById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.VehicleType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				vehicleType, err := b.vehicleTypeService.UpdateVehicleType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update vehicle type", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", vehicleType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *vehicleTypeController) DeleteVehicleType(c *gin.Context) {
	var (
		newData  model.VehicleType
		oldData  model.VehicleType
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
			oldData, err = b.vehicleTypeService.FindVehicleTypeById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.VehicleType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				vehicleType, err := b.vehicleTypeService.DeleteVehicleType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete vehicle type", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", vehicleType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
