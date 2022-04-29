package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyLicenseRenewalTracingController interface {
	FindCompanyLicenseRenewalTracings(c *gin.Context)
	FindCompanyLicenseRenewalTracingById(c *gin.Context)
	FindExcCompanyLicenseRenewalTracing(c *gin.Context)
	InsertCompanyLicenseRenewalTracing(c *gin.Context)
	UpdateCompanyLicenseRenewalTracing(c *gin.Context)
	DeleteCompanyLicenseRenewalTracing(c *gin.Context)
}

type companyLicenseRenewalTracingController struct {
	companyLicenseRenewalTracingService service.CompanyLicenseRenewalTracingService
	jwtService                          service.JWTService
}

func NewCompanyLicenseRenewalTracingController(companyLicenseRenewalTracingServ service.CompanyLicenseRenewalTracingService, jwtServ service.JWTService) CompanyLicenseRenewalTracingController {
	return &companyLicenseRenewalTracingController{
		companyLicenseRenewalTracingService: companyLicenseRenewalTracingServ,
		jwtService:                          jwtServ,
	}
}

func (b *companyLicenseRenewalTracingController) FindCompanyLicenseRenewalTracings(c *gin.Context) {
	var (
		companyLicenseRenewalTracings []model.CompanyLicenseRenewalTracing
		response                      helper.Response
	)
	companyLicenseRenewalTracings, err := b.companyLicenseRenewalTracingService.FindCompanyLicenseRenewalTracings()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", companyLicenseRenewalTracings)
		c.JSON(http.StatusOK, response)
	}
}

func (b *companyLicenseRenewalTracingController) FindCompanyLicenseRenewalTracingById(c *gin.Context) {
	var (
		companyLicenseRenewalTracing model.CompanyLicenseRenewalTracing
		response                     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyLicenseRenewalTracing, err = b.companyLicenseRenewalTracingService.FindCompanyLicenseRenewalTracingById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyLicenseRenewalTracing)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyLicenseRenewalTracingController) FindExcCompanyLicenseRenewalTracing(c *gin.Context) {
	var (
		companyLicenseRenewalTracings []model.CompanyLicenseRenewalTracing
		response                      helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyLicenseRenewalTracings, err = b.companyLicenseRenewalTracingService.FindExcCompanyLicenseRenewalTracing(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyLicenseRenewalTracings)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyLicenseRenewalTracingController) InsertCompanyLicenseRenewalTracing(c *gin.Context) {
	var (
		companyLicenseRenewalTracing                model.CompanyLicenseRenewalTracing
		response                                    helper.Response
		CreateCompanyLicenseRenewalTracingParameter model.CreateCompanyLicenseRenewalTracingParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyLicenseRenewalTracingParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyLicenseRenewalTracing, err = b.companyLicenseRenewalTracingService.InsertCompanyLicenseRenewalTracing(CreateCompanyLicenseRenewalTracingParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyLicenseRenewalTracing", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyLicenseRenewalTracing)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyLicenseRenewalTracingController) UpdateCompanyLicenseRenewalTracing(c *gin.Context) {
	var (
		newData  model.CompanyLicenseRenewalTracing
		oldData  model.CompanyLicenseRenewalTracing
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
			oldData, err = b.companyLicenseRenewalTracingService.FindCompanyLicenseRenewalTracingById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyLicenseRenewalTracing{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyLicenseRenewalTracing, err := b.companyLicenseRenewalTracingService.UpdateCompanyLicenseRenewalTracing(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyLicenseRenewalTracing", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyLicenseRenewalTracing)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyLicenseRenewalTracingController) DeleteCompanyLicenseRenewalTracing(c *gin.Context) {
	var (
		newData  model.CompanyLicenseRenewalTracing
		oldData  model.CompanyLicenseRenewalTracing
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
			oldData, err = b.companyLicenseRenewalTracingService.FindCompanyLicenseRenewalTracingById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyLicenseRenewalTracing{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyLicenseRenewalTracing, err := b.companyLicenseRenewalTracingService.DeleteCompanyLicenseRenewalTracing(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyLicenseRenewalTracing", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyLicenseRenewalTracing)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
