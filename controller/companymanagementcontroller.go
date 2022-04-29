package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyManagementController interface {
	FindCompanyManagements(c *gin.Context)
	FindCompanyManagementById(c *gin.Context)
	FindExcCompanyManagement(c *gin.Context)
	FindCompanyManagementByCompanyId(c *gin.Context)
	InsertCompanyManagement(c *gin.Context)
	UpdateCompanyManagement(c *gin.Context)
	DeleteCompanyManagement(c *gin.Context)
}

type companyManagementController struct {
	companyManagementService service.CompanyManagementService
	jwtService               service.JWTService
}

func NewCompanyManagementController(companyManagementServ service.CompanyManagementService, jwtServ service.JWTService) CompanyManagementController {
	return &companyManagementController{
		companyManagementService: companyManagementServ,
		jwtService:               jwtServ,
	}
}

func (b *companyManagementController) FindCompanyManagements(c *gin.Context) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
		response           helper.Response
	)
	companyManagements, err := b.companyManagementService.FindCompanyManagements()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", companyManagements)
		c.JSON(http.StatusOK, response)
	}
}

func (b *companyManagementController) FindCompanyManagementById(c *gin.Context) {
	var (
		companyManagement model.SelectCompanyManagementParameter
		response          helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagement, err = b.companyManagementService.FindCompanyManagementById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagement)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementController) FindExcCompanyManagement(c *gin.Context) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
		response           helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagements, err = b.companyManagementService.FindExcCompanyManagement(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagements)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementController) FindCompanyManagementByCompanyId(c *gin.Context) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
		response           helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param company id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyManagements, err = b.companyManagementService.FindCompanyManagementByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagements)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementController) InsertCompanyManagement(c *gin.Context) {
	var (
		companyManagement                model.CompanyManagement
		response                         helper.Response
		CreateCompanyManagementParameter model.CreateCompanyManagementParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyManagementParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyManagement, err = b.companyManagementService.InsertCompanyManagement(CreateCompanyManagementParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyManagement", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyManagement)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyManagementController) UpdateCompanyManagement(c *gin.Context) {
	var (
		newData  model.CompanyManagement
		oldData  model.SelectCompanyManagementParameter
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
			oldData, err = b.companyManagementService.FindCompanyManagementById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyManagementParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyManagement, err := b.companyManagementService.UpdateCompanyManagement(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyManagement", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyManagement)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyManagementController) DeleteCompanyManagement(c *gin.Context) {
	var (
		newData  model.CompanyManagement
		oldData  model.SelectCompanyManagementParameter
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
			oldData, err = b.companyManagementService.FindCompanyManagementById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyManagementParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyManagement, err := b.companyManagementService.DeleteCompanyManagement(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyManagement", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyManagement)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
