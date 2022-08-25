package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroupLicenseTypeController interface {
	CountGroupLicenseTypeAll(c *gin.Context)
	FindGroupLicenseTypes(c *gin.Context)
	FindGroupLicenseTypesOffset(c *gin.Context)
	SearchGroupLicenseType(c *gin.Context)
	CountSearchGroupLicenseType(c *gin.Context)
	FindGroupLicenseTypeById(c *gin.Context)
	FindExcGroupLicenseType(c *gin.Context)
	InsertGroupLicenseType(c *gin.Context)
	UpdateGroupLicenseType(c *gin.Context)
	DeleteGroupLicenseType(c *gin.Context)
}

type groupLicenseTypeController struct {
	groupLicenseTypeService service.GroupLicenseTypeService
	jwtService              service.JWTService
}

func NewGroupLicenseTypeController(groupLicenseTypeServ service.GroupLicenseTypeService, jwtServ service.JWTService) GroupLicenseTypeController {
	return &groupLicenseTypeController{
		groupLicenseTypeService: groupLicenseTypeServ,
		jwtService:              jwtServ,
	}
}

func (b *groupLicenseTypeController) CountGroupLicenseTypeAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.groupLicenseTypeService.CountGroupLicenseTypeAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *groupLicenseTypeController) FindGroupLicenseTypes(c *gin.Context) {
	var (
		businessunits []model.GroupLicenseType
		response      helper.Response
	)
	businessunits, err := b.groupLicenseTypeService.FindGroupLicenseTypes()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", businessunits)
		c.JSON(http.StatusOK, response)
	}
}

func (b *groupLicenseTypeController) FindGroupLicenseTypesOffset(c *gin.Context) {
	var (
		businessunits []model.GroupLicenseType
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
					businessunits, err = b.groupLicenseTypeService.FindGroupLicenseTypesOffset(int(limit), int(offset), order, dir)
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

func (b *groupLicenseTypeController) SearchGroupLicenseType(c *gin.Context) {
	var (
		businessunits []model.GroupLicenseType
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
						businessunits, err = b.groupLicenseTypeService.SearchGroupLicenseType(int(limit), int(offset), order, dir, search)
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

func (b *groupLicenseTypeController) CountSearchGroupLicenseType(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.groupLicenseTypeService.CountSearchGroupLicenseType(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *groupLicenseTypeController) FindGroupLicenseTypeById(c *gin.Context) {
	var (
		groupLicenseType model.GroupLicenseType
		response         helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		groupLicenseType, err = b.groupLicenseTypeService.FindGroupLicenseTypeById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", groupLicenseType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *groupLicenseTypeController) FindExcGroupLicenseType(c *gin.Context) {
	var (
		groupLicenseTypes []model.GroupLicenseType
		response          helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		groupLicenseTypes, err = b.groupLicenseTypeService.FindExcGroupLicenseType(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", groupLicenseTypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *groupLicenseTypeController) InsertGroupLicenseType(c *gin.Context) {
	var (
		groupLicenseType                model.GroupLicenseType
		response                        helper.Response
		CreateGroupLicenseTypeParameter model.CreateGroupLicenseTypeParameter
	)
	err := c.ShouldBindJSON(&CreateGroupLicenseTypeParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		groupLicenseType, err = b.groupLicenseTypeService.InsertGroupLicenseType(CreateGroupLicenseTypeParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register group license type", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", groupLicenseType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *groupLicenseTypeController) UpdateGroupLicenseType(c *gin.Context) {
	var (
		newData  model.GroupLicenseType
		oldData  model.GroupLicenseType
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
			oldData, err = b.groupLicenseTypeService.FindGroupLicenseTypeById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.GroupLicenseType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				groupLicenseType, err := b.groupLicenseTypeService.UpdateGroupLicenseType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update group license type", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", groupLicenseType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *groupLicenseTypeController) DeleteGroupLicenseType(c *gin.Context) {
	var (
		newData  model.GroupLicenseType
		oldData  model.GroupLicenseType
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
			oldData, err = b.groupLicenseTypeService.FindGroupLicenseTypeById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.GroupLicenseType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				groupLicenseType, err := b.groupLicenseTypeService.DeleteGroupLicenseType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete group license type", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", groupLicenseType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
