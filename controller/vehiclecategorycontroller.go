package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehicleCategoryController interface {
	CountVehicleCategoryAll(c *gin.Context)
	FindVehicleCategorys(c *gin.Context)
	FindVehicleCategorysOffset(c *gin.Context)
	SearchVehicleCategory(c *gin.Context)
	CountSearchVehicleCategory(c *gin.Context)
	FindVehicleCategoryById(c *gin.Context)
	FindExcVehicleCategory(c *gin.Context)
	InsertVehicleCategory(c *gin.Context)
	UpdateVehicleCategory(c *gin.Context)
	DeleteVehicleCategory(c *gin.Context)
}

type vehicleCategoryController struct {
	vehicleCategoryService service.VehicleCategoryService
	jwtService             service.JWTService
}

func NewVehicleCategoryController(vehicleCategoryServ service.VehicleCategoryService, jwtServ service.JWTService) VehicleCategoryController {
	return &vehicleCategoryController{
		vehicleCategoryService: vehicleCategoryServ,
		jwtService:             jwtServ,
	}
}

func (b *vehicleCategoryController) CountVehicleCategoryAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.vehicleCategoryService.CountVehicleCategoryAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *vehicleCategoryController) FindVehicleCategorys(c *gin.Context) {
	var (
		vehiclecategorys []model.VehicleCategory
		response         helper.Response
	)
	vehiclecategorys, err := b.vehicleCategoryService.FindVehicleCategorys()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", vehiclecategorys)
		c.JSON(http.StatusOK, response)
	}
}

func (b *vehicleCategoryController) FindVehicleCategorysOffset(c *gin.Context) {
	var (
		vehiclecategorys []model.VehicleCategory
		response         helper.Response
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
					vehiclecategorys, err = b.vehicleCategoryService.FindVehicleCategorysOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", vehiclecategorys)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *vehicleCategoryController) SearchVehicleCategory(c *gin.Context) {
	var (
		vehiclecategorys []model.VehicleCategory
		response         helper.Response
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
						vehiclecategorys, err = b.vehicleCategoryService.SearchVehicleCategory(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", vehiclecategorys)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *vehicleCategoryController) CountSearchVehicleCategory(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.vehicleCategoryService.CountSearchVehicleCategory(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleCategoryController) FindVehicleCategoryById(c *gin.Context) {
	var (
		vehicleCategory model.VehicleCategory
		response        helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		vehicleCategory, err = b.vehicleCategoryService.FindVehicleCategoryById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleCategory)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleCategoryController) FindExcVehicleCategory(c *gin.Context) {
	var (
		vehicleCategorys []model.VehicleCategory
		response         helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		vehicleCategorys, err = b.vehicleCategoryService.FindExcVehicleCategory(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleCategorys)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleCategoryController) InsertVehicleCategory(c *gin.Context) {
	var (
		vehicleCategory                model.VehicleCategory
		response                       helper.Response
		CreateVehicleCategoryParameter model.CreateVehicleCategoryParameter
	)
	err := c.ShouldBindJSON(&CreateVehicleCategoryParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		vehicleCategory, err = b.vehicleCategoryService.InsertVehicleCategory(CreateVehicleCategoryParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register vehicle category", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", vehicleCategory)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *vehicleCategoryController) UpdateVehicleCategory(c *gin.Context) {
	var (
		newData  model.VehicleCategory
		oldData  model.VehicleCategory
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
			oldData, err = b.vehicleCategoryService.FindVehicleCategoryById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.VehicleCategory{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				vehicleCategory, err := b.vehicleCategoryService.UpdateVehicleCategory(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update vehicle category", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", vehicleCategory)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *vehicleCategoryController) DeleteVehicleCategory(c *gin.Context) {
	var (
		newData  model.VehicleCategory
		oldData  model.VehicleCategory
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
			oldData, err = b.vehicleCategoryService.FindVehicleCategoryById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.VehicleCategory{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				vehicleCategory, err := b.vehicleCategoryService.DeleteVehicleCategory(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete vehicle category", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", vehicleCategory)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
