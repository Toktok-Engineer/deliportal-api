package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LocationController interface {
	FindLocations(c *gin.Context)
	FindLocationById(c *gin.Context)
	FindExcLocation(c *gin.Context)
	InsertLocation(c *gin.Context)
	UpdateLocation(c *gin.Context)
	DeleteLocation(c *gin.Context)
}

type locationController struct {
	locationService service.LocationService
	jwtService      service.JWTService
}

func NewLocationController(locationServ service.LocationService, jwtServ service.JWTService) LocationController {
	return &locationController{
		locationService: locationServ,
		jwtService:      jwtServ,
	}
}

func (b *locationController) FindLocations(c *gin.Context) {
	var (
		locations []model.Location
		response  helper.Response
	)
	locations, err := b.locationService.FindLocations()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", locations)
		c.JSON(http.StatusOK, response)
	}
}

func (b *locationController) FindLocationById(c *gin.Context) {
	var (
		location model.Location
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		location, err = b.locationService.FindLocationById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", location)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *locationController) FindExcLocation(c *gin.Context) {
	var (
		locations []model.Location
		response  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		locations, err = b.locationService.FindExcLocation(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", locations)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *locationController) InsertLocation(c *gin.Context) {
	var (
		location                model.Location
		response                helper.Response
		CreateLocationParameter model.CreateLocationParameter
	)
	err := c.ShouldBindJSON(&CreateLocationParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		location, err = b.locationService.InsertLocation(CreateLocationParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register location", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", location)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *locationController) UpdateLocation(c *gin.Context) {
	var (
		newData  model.Location
		oldData  model.Location
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
			oldData, err = b.locationService.FindLocationById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.Location{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				location, err := b.locationService.UpdateLocation(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update location", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", location)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *locationController) DeleteLocation(c *gin.Context) {
	var (
		newData  model.Location
		oldData  model.Location
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
			oldData, err = b.locationService.FindLocationById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.Location{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				location, err := b.locationService.DeleteLocation(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete location", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", location)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
