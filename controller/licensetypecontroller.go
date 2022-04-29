package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LicenseTypeController interface {
	FindLicenseTypes(c *gin.Context)
	FindLicenseTypeById(c *gin.Context)
	FindExcLicenseType(c *gin.Context)
	InsertLicenseType(c *gin.Context)
	UpdateLicenseType(c *gin.Context)
	DeleteLicenseType(c *gin.Context)
}

type licenseTypeController struct {
	licenseTypeService service.LicenseTypeService
	jwtService         service.JWTService
}

func NewLicenseTypeController(licenseTypeServ service.LicenseTypeService, jwtServ service.JWTService) LicenseTypeController {
	return &licenseTypeController{
		licenseTypeService: licenseTypeServ,
		jwtService:         jwtServ,
	}
}

func (b *licenseTypeController) FindLicenseTypes(c *gin.Context) {
	var (
		licenseTypes []model.LicenseType
		response     helper.Response
	)
	licenseTypes, err := b.licenseTypeService.FindLicenseTypes()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", licenseTypes)
		c.JSON(http.StatusOK, response)
	}
}

func (b *licenseTypeController) FindLicenseTypeById(c *gin.Context) {
	var (
		licenseType model.LicenseType
		response    helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		licenseType, err = b.licenseTypeService.FindLicenseTypeById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", licenseType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *licenseTypeController) FindExcLicenseType(c *gin.Context) {
	var (
		licenseTypes []model.LicenseType
		response     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		licenseTypes, err = b.licenseTypeService.FindExcLicenseType(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", licenseTypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *licenseTypeController) InsertLicenseType(c *gin.Context) {
	var (
		licenseType                model.LicenseType
		response                   helper.Response
		CreateLicenseTypeParameter model.CreateLicenseTypeParameter
	)
	err := c.ShouldBindJSON(&CreateLicenseTypeParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		licenseType, err = b.licenseTypeService.InsertLicenseType(CreateLicenseTypeParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register licenseType", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", licenseType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *licenseTypeController) UpdateLicenseType(c *gin.Context) {
	var (
		newData  model.LicenseType
		oldData  model.LicenseType
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
			oldData, err = b.licenseTypeService.FindLicenseTypeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.LicenseType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				licenseType, err := b.licenseTypeService.UpdateLicenseType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update licenseType", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", licenseType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *licenseTypeController) DeleteLicenseType(c *gin.Context) {
	var (
		newData  model.LicenseType
		oldData  model.LicenseType
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
			oldData, err = b.licenseTypeService.FindLicenseTypeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.LicenseType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				licenseType, err := b.licenseTypeService.DeleteLicenseType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete licenseType", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", licenseType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
