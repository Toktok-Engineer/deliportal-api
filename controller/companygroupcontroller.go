package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyGroupController interface {
	CountCompanyGroupAll(c *gin.Context)
	FindCompanyGroups(c *gin.Context)
	FindCompanyGroupsOffset(c *gin.Context)
	SearchCompanyGroup(c *gin.Context)
	CountSearchCompanyGroup(c *gin.Context)
	FindCompanyGroupById(c *gin.Context)
	FindExcCompanyGroup(c *gin.Context)
	InsertCompanyGroup(c *gin.Context)
	UpdateCompanyGroup(c *gin.Context)
	DeleteCompanyGroup(c *gin.Context)
}

type companyGroupController struct {
	companyGroupService service.CompanyGroupService
	jwtService          service.JWTService
}

func NewCompanyGroupController(companyGroupServ service.CompanyGroupService, jwtServ service.JWTService) CompanyGroupController {
	return &companyGroupController{
		companyGroupService: companyGroupServ,
		jwtService:          jwtServ,
	}
}

func (b *companyGroupController) CountCompanyGroupAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.companyGroupService.CountCompanyGroupAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *companyGroupController) FindCompanyGroups(c *gin.Context) {
	var (
		businessunits []model.CompanyGroup
		response      helper.Response
	)
	businessunits, err := b.companyGroupService.FindCompanyGroups()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", businessunits)
		c.JSON(http.StatusOK, response)
	}
}

func (b *companyGroupController) FindCompanyGroupsOffset(c *gin.Context) {
	var (
		businessunits []model.CompanyGroup
		response      helper.Response
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
					businessunits, err = b.companyGroupService.FindCompanyGroupsOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", businessunits)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *companyGroupController) SearchCompanyGroup(c *gin.Context) {
	var (
		businessunits []model.CompanyGroup
		response      helper.Response
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
						businessunits, err = b.companyGroupService.SearchCompanyGroup(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", businessunits)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *companyGroupController) CountSearchCompanyGroup(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.companyGroupService.CountSearchCompanyGroup(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyGroupController) FindCompanyGroupById(c *gin.Context) {
	var (
		companyGroup model.CompanyGroup
		response     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyGroup, err = b.companyGroupService.FindCompanyGroupById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyGroup)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyGroupController) FindExcCompanyGroup(c *gin.Context) {
	var (
		companyGroups []model.CompanyGroup
		response      helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		companyGroups, err = b.companyGroupService.FindExcCompanyGroup(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyGroups)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyGroupController) InsertCompanyGroup(c *gin.Context) {
	var (
		companyGroup                model.CompanyGroup
		response                    helper.Response
		CreateCompanyGroupParameter model.CreateCompanyGroupParameter
	)
	err := c.ShouldBindJSON(&CreateCompanyGroupParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		companyGroup, err = b.companyGroupService.InsertCompanyGroup(CreateCompanyGroupParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register company group", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", companyGroup)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *companyGroupController) UpdateCompanyGroup(c *gin.Context) {
	var (
		newData  model.CompanyGroup
		oldData  model.CompanyGroup
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
			oldData, err = b.companyGroupService.FindCompanyGroupById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.CompanyGroup{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyGroup, err := b.companyGroupService.UpdateCompanyGroup(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update company group", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyGroup)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *companyGroupController) DeleteCompanyGroup(c *gin.Context) {
	var (
		newData  model.CompanyGroup
		oldData  model.CompanyGroup
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
			oldData, err = b.companyGroupService.FindCompanyGroupById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.CompanyGroup{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				companyGroup, err := b.companyGroupService.DeleteCompanyGroup(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete company group", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", companyGroup)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
