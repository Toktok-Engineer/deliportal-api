package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NonWorkingTypeController interface {
	CountNonWorkingTypeAll(c *gin.Context)
	FindNonWorkingTypes(c *gin.Context)
	FindNonWorkingTypesOffset(c *gin.Context)
	SearchNonWorkingType(c *gin.Context)
	CountSearchNonWorkingType(c *gin.Context)
	FindNonWorkingTypeById(c *gin.Context)
	CountNonWorkingTypeName(c *gin.Context)
	FindExcNonWorkingType(c *gin.Context)
	InsertNonWorkingType(c *gin.Context)
	UpdateNonWorkingType(c *gin.Context)
	DeleteNonWorkingType(c *gin.Context)
}

type nonworkingtypeController struct {
	nonworkingtypeService service.NonWorkingTypeService
	jwtService            service.JWTService
}

func NewNonWorkingTypeController(nonworkingtypeServ service.NonWorkingTypeService, jwtServ service.JWTService) NonWorkingTypeController {
	return &nonworkingtypeController{
		nonworkingtypeService: nonworkingtypeServ,
		jwtService:            jwtServ,
	}
}

func (b *nonworkingtypeController) CountNonWorkingTypeAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.nonworkingtypeService.CountNonWorkingTypeAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *nonworkingtypeController) FindNonWorkingTypes(c *gin.Context) {
	var (
		nonworkingtypes []model.NonWorkingType
		response        helper.Response
	)
	nonworkingtypes, err := b.nonworkingtypeService.FindNonWorkingTypes()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", nonworkingtypes)
		c.JSON(http.StatusOK, response)
	}
}

func (b *nonworkingtypeController) FindNonWorkingTypesOffset(c *gin.Context) {
	var (
		nonworkingtypes []model.NonWorkingType
		response        helper.Response
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
					nonworkingtypes, err = b.nonworkingtypeService.FindNonWorkingTypesOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", nonworkingtypes)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *nonworkingtypeController) SearchNonWorkingType(c *gin.Context) {
	var (
		nonworkingtypes []model.NonWorkingType
		response        helper.Response
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
						nonworkingtypes, err = b.nonworkingtypeService.SearchNonWorkingType(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", nonworkingtypes)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *nonworkingtypeController) CountSearchNonWorkingType(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.nonworkingtypeService.CountSearchNonWorkingType(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonworkingtypeController) FindNonWorkingTypeById(c *gin.Context) {
	var (
		nonworkingtype model.NonWorkingType
		response       helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		nonworkingtype, err = b.nonworkingtypeService.FindNonWorkingTypeById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", nonworkingtype)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonworkingtypeController) CountNonWorkingTypeName(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.nonworkingtypeService.CountNonWorkingTypeName(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonworkingtypeController) FindExcNonWorkingType(c *gin.Context) {
	var (
		nonworkingtypes []model.NonWorkingType
		response        helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		nonworkingtypes, err = b.nonworkingtypeService.FindExcNonWorkingType(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", nonworkingtypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonworkingtypeController) InsertNonWorkingType(c *gin.Context) {
	var (
		nonworkingtype                model.NonWorkingType
		response                      helper.Response
		CreateNonWorkingTypeParameter model.CreateNonWorkingTypeParameter
	)
	err := c.ShouldBindJSON(&CreateNonWorkingTypeParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		nonworkingtype, err = b.nonworkingtypeService.InsertNonWorkingType(CreateNonWorkingTypeParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register nonworkingtype", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", nonworkingtype)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonworkingtypeController) UpdateNonWorkingType(c *gin.Context) {
	var (
		newData  model.NonWorkingType
		oldData  model.NonWorkingType
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
			oldData, err = b.nonworkingtypeService.FindNonWorkingTypeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.NonWorkingType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				nonworkingtype, err := b.nonworkingtypeService.UpdateNonWorkingType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update nonworkingtype", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", nonworkingtype)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *nonworkingtypeController) DeleteNonWorkingType(c *gin.Context) {
	var (
		newData  model.NonWorkingType
		oldData  model.NonWorkingType
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
			oldData, err = b.nonworkingtypeService.FindNonWorkingTypeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.NonWorkingType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				nonworkingtype, err := b.nonworkingtypeService.DeleteNonWorkingType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete nonworkingtype", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", nonworkingtype)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
