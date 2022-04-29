package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyManagementTypeController interface {
	FindCompanyManagementTypes(c *gin.Context)
	FindCompanyManagementTypeById(c *gin.Context)
	FindExcCompanyManagementType(c *gin.Context)
	InsertCompanyManagementType(c *gin.Context)
	UpdateCompanyManagementType(c *gin.Context)
	DeleteCompanyManagementType(c *gin.Context)
}

type companyManagementTypeController struct {
	companyManagementTypeService service.CompanyManagementTypeService
	jwtService                   service.JWTService
}

func NewCompanyManagementTypeController(companyManagementTypeServ service.CompanyManagementTypeService, jwtServ service.JWTService) CompanyManagementTypeController {
	return &companyManagementTypeController{
		companyManagementTypeService: companyManagementTypeServ,
		jwtService:                   jwtServ,
	}
}

func (b *companyManagementTypeController) FindCompanyManagementTypes(c *gin.Context) {
	var (
		companyManagementTypes []model.CompanyManagementType
		response               helper.Response
	)
	companyManagementTypes, err := b.companyManagementTypeService.FindCompanyManagementTypes()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", companyManagementTypes)
		c.JSON(http.StatusOK, response)
	}
}

func (b *companyManagementTypeController) FindCompanyManagementTypeById(c *gin.Context) {
	var (
		companyManagementType model.CompanyManagementType
		response              helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementType, err = b.companyManagementTypeService.FindCompanyManagementTypeById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementTypeController) FindExcCompanyManagementType(c *gin.Context) {
	var (
		companyManagementTypes []model.CompanyManagementType
		response               helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagementTypes, err = b.companyManagementTypeService.FindExcCompanyManagementType(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementTypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementTypeController) InsertCompanyManagementType(c *gin.Context) {
	var (
		companyManagementType                model.CompanyManagementType
		response                             helper.Response
		CreateCompanyManagementTypeParameter model.CreateCompanyManagementTypeParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyManagementTypeParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyManagementType, err = b.companyManagementTypeService.InsertCompanyManagementType(CreateCompanyManagementTypeParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyManagementType", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagementType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementTypeController) UpdateCompanyManagementType(c *gin.Context) {
	var (
		newData  model.CompanyManagementType
		oldData  model.CompanyManagementType
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
			oldData, err = b.companyManagementTypeService.FindCompanyManagementTypeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyManagementType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyManagementType, err := b.companyManagementTypeService.UpdateCompanyManagementType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyManagementType", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyManagementType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyManagementTypeController) DeleteCompanyManagementType(c *gin.Context) {
	var (
		newData  model.CompanyManagementType
		oldData  model.CompanyManagementType
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
			oldData, err = b.companyManagementTypeService.FindCompanyManagementTypeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyManagementType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyManagementType, err := b.companyManagementTypeService.DeleteCompanyManagementType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyManagementType", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyManagementType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
