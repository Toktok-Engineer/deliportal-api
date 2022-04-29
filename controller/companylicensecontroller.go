package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyLicenseController interface {
	FindCompanyLicenses(c *gin.Context)
	FindCompanyLicenseById(c *gin.Context)
	FindExcCompanyLicense(c *gin.Context)
	FindCompanyLicenseByCompanyId(c *gin.Context)
	InsertCompanyLicense(c *gin.Context)
	UpdateCompanyLicense(c *gin.Context)
	UpdateCompanyLicenseStatus(c *gin.Context)
	UpdateCompanyLicenseDeactive(c *gin.Context)
	UpdateCompanyLicenseApprovedRenewalStatus(c *gin.Context)
	DeleteCompanyLicense(c *gin.Context)
	UpdateCompanyRemark(c *gin.Context)
}

type companyLicenseController struct {
	companyLicenseService service.CompanyLicenseService
	jwtService            service.JWTService
}

func NewCompanyLicenseController(companyLicenseServ service.CompanyLicenseService, jwtServ service.JWTService) CompanyLicenseController {
	return &companyLicenseController{
		companyLicenseService: companyLicenseServ,
		jwtService:            jwtServ,
	}
}

func (b *companyLicenseController) FindCompanyLicenses(c *gin.Context) {
	var (
		companyLicenses []model.SelectCompanyLicenseParameter
		response        helper.Response
	)
	companyLicenses, err := b.companyLicenseService.FindCompanyLicenses()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", companyLicenses)
		c.JSON(http.StatusOK, response)
	}
}

func (b *companyLicenseController) FindCompanyLicenseById(c *gin.Context) {
	var (
		companyLicense model.SelectCompanyLicenseParameter
		response       helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyLicense, err = b.companyLicenseService.FindCompanyLicenseById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyLicense)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyLicenseController) FindExcCompanyLicense(c *gin.Context) {
	var (
		companyLicenses []model.SelectCompanyLicenseParameter
		response        helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyLicenses, err = b.companyLicenseService.FindExcCompanyLicense(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyLicenses)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyLicenseController) FindCompanyLicenseByCompanyId(c *gin.Context) {
	var (
		companyLicenses []model.SelectCompanyLicenseParameter
		response        helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param company id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyLicenses, err = b.companyLicenseService.FindCompanyLicenseByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyLicenses)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyLicenseController) InsertCompanyLicense(c *gin.Context) {
	var (
		companyLicense                model.CompanyLicense
		response                      helper.Response
		CreateCompanyLicenseParameter model.CreateCompanyLicenseParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyLicenseParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyLicense, err = b.companyLicenseService.InsertCompanyLicense(CreateCompanyLicenseParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register companyLicense", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyLicense)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyLicenseController) UpdateCompanyLicense(c *gin.Context) {
	var (
		newData  model.CompanyLicense
		oldData  model.SelectCompanyLicenseParameter
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
			oldData, err = b.companyLicenseService.FindCompanyLicenseById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyLicenseParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyLicense, err := b.companyLicenseService.UpdateCompanyLicense(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyLicense", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyLicense)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyLicenseController) UpdateCompanyLicenseStatus(c *gin.Context) {
	var (
		newData  model.CompanyLicense
		oldData  model.SelectCompanyLicenseParameter
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
			oldData, err = b.companyLicenseService.FindCompanyLicenseById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyLicenseParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyLicense, err := b.companyLicenseService.UpdateCompanyLicenseStatus(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyLicense", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyLicense)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyLicenseController) UpdateCompanyLicenseDeactive(c *gin.Context) {
	var (
		newData  model.CompanyLicense
		oldData  model.SelectCompanyLicenseParameter
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
			oldData, err = b.companyLicenseService.FindCompanyLicenseById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyLicenseParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyLicense, err := b.companyLicenseService.UpdateCompanyLicenseDeactive(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyLicense", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyLicense)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyLicenseController) UpdateCompanyLicenseApprovedRenewalStatus(c *gin.Context) {
	var (
		newData  model.CompanyLicense
		oldData  model.SelectCompanyLicenseParameter
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
			oldData, err = b.companyLicenseService.FindCompanyLicenseById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyLicenseParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyLicense, err := b.companyLicenseService.UpdateCompanyLicenseApprovedRenewalStatus(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyLicense", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyLicense)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyLicenseController) DeleteCompanyLicense(c *gin.Context) {
	var (
		newData  model.CompanyLicense
		oldData  model.SelectCompanyLicenseParameter
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
			oldData, err = b.companyLicenseService.FindCompanyLicenseById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyLicenseParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyLicense, err := b.companyLicenseService.DeleteCompanyLicense(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete companyLicense", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyLicense)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyLicenseController) UpdateCompanyRemark(c *gin.Context) {
	var (
		newData  model.CompanyLicense
		oldData  model.SelectCompanyLicenseParameter
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
			oldData, err = b.companyLicenseService.FindCompanyLicenseById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectCompanyLicenseParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyLicense, err := b.companyLicenseService.UpdateCompanyRemark(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update companyLicense", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyLicense)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
