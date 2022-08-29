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
	CountLicenseTypeAll(c *gin.Context)
	FindLicenseTypes(c *gin.Context)
	FindLicenseTypesOffset(c *gin.Context)
	SearchLicenseType(c *gin.Context)
	CountSearchLicenseType(c *gin.Context)
	FindLicenseTypeById(c *gin.Context)
	FindExcLicenseType(c *gin.Context)
	FindExcCompleteLicenseType(c *gin.Context)
	FindLicenseTypeByGroupLT(c *gin.Context)
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

func (b *licenseTypeController) CountLicenseTypeAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.licenseTypeService.CountLicenseTypeAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *licenseTypeController) FindLicenseTypes(c *gin.Context) {
	var (
		licenseTypes []model.SelectLicenseTypeParameter
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

func (b *licenseTypeController) FindLicenseTypesOffset(c *gin.Context) {
	var (
		licenseTypes []model.SelectLicenseTypeParameter
		response     helper.Response
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
					licenseTypes, err = b.licenseTypeService.FindLicenseTypesOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", licenseTypes)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *licenseTypeController) SearchLicenseType(c *gin.Context) {
	var (
		licenseTypes []model.SelectLicenseTypeParameter
		response     helper.Response
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
						licenseTypes, err = b.licenseTypeService.SearchLicenseType(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", licenseTypes)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *licenseTypeController) CountSearchLicenseType(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.licenseTypeService.CountSearchLicenseType(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *licenseTypeController) FindLicenseTypeById(c *gin.Context) {
	var (
		licenseType model.SelectLicenseTypeParameter
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
		licenseTypes []model.SelectLicenseTypeParameter
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

func (b *licenseTypeController) FindExcCompleteLicenseType(c *gin.Context) {
	var (
		licenseTypes []model.SelectLicenseTypeParameter
		response     helper.Response
	)
	groupLT, err := strconv.ParseUint(c.Param("groupLT"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param groupLT was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		id, err := strconv.ParseUint(c.Param("id"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			licenseTypes, err = b.licenseTypeService.FindExcCompleteLicenseType(uint(groupLT), uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", licenseTypes)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *licenseTypeController) FindLicenseTypeByGroupLT(c *gin.Context) {
	var (
		licenseTypes []model.SelectLicenseTypeParameter
		response     helper.Response
	)
	groupLT, err := strconv.ParseUint(c.Param("groupLT"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param division id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		licenseTypes, err = b.licenseTypeService.FindLicenseTypeByGroupLT(uint(groupLT))
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
		oldData  model.SelectLicenseTypeParameter
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
			} else if (oldData == model.SelectLicenseTypeParameter{}) {
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
		oldData  model.SelectLicenseTypeParameter
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
			} else if (oldData == model.SelectLicenseTypeParameter{}) {
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
