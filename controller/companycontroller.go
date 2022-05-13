package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

type CompanyController interface {
	FindCompanys(c *gin.Context)
	FindCompanyApprove(c *gin.Context)
	FindCompanyById(c *gin.Context)
	FindExcCompany(c *gin.Context)
	InsertCompany(c *gin.Context)
	UpdateCompany(c *gin.Context)
	UpdateCompanyApprove(c *gin.Context)
	UpdateCompanyDeactive(c *gin.Context)
	DeleteCompany(c *gin.Context)
}

type companyController struct {
	companyService service.CompanyService
	jwtService     service.JWTService
}

func NewCompanyController(companyServ service.CompanyService, jwtServ service.JWTService) CompanyController {
	return &companyController{
		companyService: companyServ,
		jwtService:     jwtServ,
	}
}

func (b *companyController) FindCompanys(c *gin.Context) {
	var (
		companys []model.SelectCompanyParameter
		response helper.Response
	)
	companys, err := b.companyService.FindCompanys()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", companys)
		c.JSON(http.StatusOK, response)
	}
}

func (b *companyController) FindCompanyApprove(c *gin.Context) {
	var (
		companys []model.SelectCompanyParameter
		response helper.Response
	)
	companys, err := b.companyService.FindCompanyApprove()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", companys)
		c.JSON(http.StatusOK, response)
	}
}

func (b *companyController) FindCompanyById(c *gin.Context) {
	var (
		company  model.SelectCompanyParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		company, err = b.companyService.FindCompanyById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", company)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyController) FindExcCompany(c *gin.Context) {
	var (
		companys []model.SelectCompanyParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companys, err = b.companyService.FindExcCompany(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companys)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyController) InsertCompany(c *gin.Context) {
	var (
		company                model.Company
		response               helper.Response
		CreateCompanyParameter model.CreateCompanyParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		company, err = b.companyService.InsertCompany(CreateCompanyParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register company", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", company)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyController) UpdateCompany(c *gin.Context) {
	var (
		newData  model.Company
		oldData  model.SelectCompanyParameter
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
			oldData, err = b.companyService.FindCompanyById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectCompanyParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				company, err := b.companyService.UpdateCompany(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update company", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", company)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyController) UpdateCompanyApprove(c *gin.Context) {
	var (
		newData  model.Company
		oldData  model.SelectCompanyParameter
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
			oldData, err = b.companyService.FindCompanyById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectCompanyParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				company, err := b.companyService.UpdateCompanyApprove(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update company", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", company)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyController) UpdateCompanyDeactive(c *gin.Context) {
	var (
		newData  model.Company
		oldData  model.SelectCompanyParameter
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
			oldData, err = b.companyService.FindCompanyById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectCompanyParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				company, err := b.companyService.UpdateCompanyDeactive(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update company", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", company)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyController) DeleteCompany(c *gin.Context) {
	var (
		newData  model.Company
		oldData  model.SelectCompanyParameter
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
			oldData, err = b.companyService.FindCompanyById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectCompanyParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				company, err := b.companyService.DeleteCompany(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete company", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", company)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
