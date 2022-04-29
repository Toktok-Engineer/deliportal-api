package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyShareholderController interface {
	FindCompanyShareholders(c *gin.Context)
	FindCompanyShareholderByCompanyId(c *gin.Context)
	FindCompanyShareholderById(c *gin.Context)
	FindExcCompanyShareholder(c *gin.Context)
	InsertCompanyShareholder(c *gin.Context)
	UpdateCompanyShareholder(c *gin.Context)
	DeleteCompanyShareholder(c *gin.Context)
}

type companyShareholderController struct {
	companyShareholderService service.CompanyShareholderService
	jwtService                service.JWTService
}

func NewCompanyShareholderController(companyShareholderServ service.CompanyShareholderService, jwtServ service.JWTService) CompanyShareholderController {
	return &companyShareholderController{
		companyShareholderService: companyShareholderServ,
		jwtService:                jwtServ,
	}
}

func (b *companyShareholderController) FindCompanyShareholders(c *gin.Context) {
	var (
		companyShareholders []model.CompanyShareholder
		response            helper.Response
	)
	companyShareholders, err := b.companyShareholderService.FindCompanyShareholders()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", companyShareholders)
		c.JSON(http.StatusOK, response)
	}
}

func (b *companyShareholderController) FindCompanyShareholderByCompanyId(c *gin.Context) {
	var (
		companyShareholders []model.CompanyShareholder
		response            helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholders, err = b.companyShareholderService.FindCompanyShareholderByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholders)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderController) FindCompanyShareholderById(c *gin.Context) {
	var (
		companyShareholder model.CompanyShareholder
		response           helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholder, err = b.companyShareholderService.FindCompanyShareholderById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholder)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderController) FindExcCompanyShareholder(c *gin.Context) {
	var (
		companyShareholders []model.CompanyShareholder
		response            helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyShareholders, err = b.companyShareholderService.FindExcCompanyShareholder(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholders)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderController) InsertCompanyShareholder(c *gin.Context) {
	var (
		companyShareholder                model.CompanyShareholder
		response                          helper.Response
		CreateCompanyShareholderParameter model.CreateCompanyShareholderParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyShareholderParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyShareholder, err = b.companyShareholderService.InsertCompanyShareholder(CreateCompanyShareholderParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyShareholder", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyShareholder)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyShareholderController) UpdateCompanyShareholder(c *gin.Context) {
	var (
		newData  model.CompanyShareholder
		oldData  model.CompanyShareholder
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
			oldData, err = b.companyShareholderService.FindCompanyShareholderById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyShareholder{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyShareholder, err := b.companyShareholderService.UpdateCompanyShareholder(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyShareholder", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyShareholder)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyShareholderController) DeleteCompanyShareholder(c *gin.Context) {
	var (
		newData  model.CompanyShareholder
		oldData  model.CompanyShareholder
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
			oldData, err = b.companyShareholderService.FindCompanyShareholderById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.CompanyShareholder{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyShareholder, err := b.companyShareholderService.DeleteCompanyShareholder(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyShareholder", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyShareholder)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
