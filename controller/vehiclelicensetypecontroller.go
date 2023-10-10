package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehicleLicenseTypeController interface {
	CountVehicleLicenseTypeAll(c *gin.Context)
	FindVehicleLicenseTypes(c *gin.Context)
	FindVehicleLicenseTypesOffset(c *gin.Context)
	SearchVehicleLicenseType(c *gin.Context)
	CountSearchVehicleLicenseType(c *gin.Context)
	FindVehicleLicenseTypeById(c *gin.Context)
	FindExcVehicleLicenseType(c *gin.Context)
	InsertVehicleLicenseType(c *gin.Context)
	UpdateVehicleLicenseType(c *gin.Context)
	DeleteVehicleLicenseType(c *gin.Context)
}

type vehicleLicenseTypeController struct {
	vehicleLicenseTypeService service.VehicleLicenseTypeService
	jwtService                service.JWTService
}

func NewVehicleLicenseTypeController(vehicleLicenseTypeServ service.VehicleLicenseTypeService, jwtServ service.JWTService) VehicleLicenseTypeController {
	return &vehicleLicenseTypeController{
		vehicleLicenseTypeService: vehicleLicenseTypeServ,
		jwtService:                jwtServ,
	}
}

func (b *vehicleLicenseTypeController) CountVehicleLicenseTypeAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.vehicleLicenseTypeService.CountVehicleLicenseTypeAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *vehicleLicenseTypeController) FindVehicleLicenseTypes(c *gin.Context) {
	var (
		vehiclelicensetypes []model.VehicleLicenseType
		response            helper.Response
	)
	vehiclelicensetypes, err := b.vehicleLicenseTypeService.FindVehicleLicenseTypes()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", vehiclelicensetypes)
		c.JSON(http.StatusOK, response)
	}
}

func (b *vehicleLicenseTypeController) FindVehicleLicenseTypesOffset(c *gin.Context) {
	var (
		vehiclelicensetypes []model.VehicleLicenseType
		response            helper.Response
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
					vehiclelicensetypes, err = b.vehicleLicenseTypeService.FindVehicleLicenseTypesOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", vehiclelicensetypes)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *vehicleLicenseTypeController) SearchVehicleLicenseType(c *gin.Context) {
	var (
		vehiclelicensetypes []model.VehicleLicenseType
		response            helper.Response
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
						vehiclelicensetypes, err = b.vehicleLicenseTypeService.SearchVehicleLicenseType(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", vehiclelicensetypes)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *vehicleLicenseTypeController) CountSearchVehicleLicenseType(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.vehicleLicenseTypeService.CountSearchVehicleLicenseType(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleLicenseTypeController) FindVehicleLicenseTypeById(c *gin.Context) {
	var (
		vehicleLicenseType model.VehicleLicenseType
		response           helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		vehicleLicenseType, err = b.vehicleLicenseTypeService.FindVehicleLicenseTypeById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleLicenseType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleLicenseTypeController) FindExcVehicleLicenseType(c *gin.Context) {
	var (
		vehicleLicenseTypes []model.VehicleLicenseType
		response            helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		vehicleLicenseTypes, err = b.vehicleLicenseTypeService.FindExcVehicleLicenseType(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleLicenseTypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleLicenseTypeController) InsertVehicleLicenseType(c *gin.Context) {
	var (
		vehicleLicenseType                model.VehicleLicenseType
		response                          helper.Response
		CreateVehicleLicenseTypeParameter model.CreateVehicleLicenseTypeParameter
	)
	err := c.ShouldBindJSON(&CreateVehicleLicenseTypeParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		vehicleLicenseType, err = b.vehicleLicenseTypeService.InsertVehicleLicenseType(CreateVehicleLicenseTypeParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register vehicle licensetype", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleLicenseType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleLicenseTypeController) UpdateVehicleLicenseType(c *gin.Context) {
	var (
		newData  model.VehicleLicenseType
		oldData  model.VehicleLicenseType
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
			oldData, err = b.vehicleLicenseTypeService.FindVehicleLicenseTypeById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.VehicleLicenseType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				vehicleLicenseType, err := b.vehicleLicenseTypeService.UpdateVehicleLicenseType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update vehicle licensetype", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", vehicleLicenseType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *vehicleLicenseTypeController) DeleteVehicleLicenseType(c *gin.Context) {
	var (
		newData  model.VehicleLicenseType
		oldData  model.VehicleLicenseType
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
			oldData, err = b.vehicleLicenseTypeService.FindVehicleLicenseTypeById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.VehicleLicenseType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				vehicleLicenseType, err := b.vehicleLicenseTypeService.DeleteVehicleLicenseType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete vehicle licensetype", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", vehicleLicenseType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
