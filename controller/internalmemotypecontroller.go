package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InternalMemoTypeController interface {
	CountInternalMemoTypeAll(c *gin.Context)
	FindInternalMemoTypes(c *gin.Context)
	FindInternalMemoTypesOffset(c *gin.Context)
	SearchInternalMemoType(c *gin.Context)
	CountSearchInternalMemoType(c *gin.Context)
	FindInternalMemoTypeById(c *gin.Context)
	FindExcInternalMemoType(c *gin.Context)
	InsertInternalMemoType(c *gin.Context)
	UpdateInternalMemoType(c *gin.Context)
	DeleteInternalMemoType(c *gin.Context)
}

type internalMemoTypeController struct {
	internalMemoTypeService service.InternalMemoTypeService
	jwtService              service.JWTService
}

func NewInternalMemoTypeController(internalMemoTypeServ service.InternalMemoTypeService, jwtServ service.JWTService) InternalMemoTypeController {
	return &internalMemoTypeController{
		internalMemoTypeService: internalMemoTypeServ,
		jwtService:              jwtServ,
	}
}

func (b *internalMemoTypeController) CountInternalMemoTypeAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.internalMemoTypeService.CountInternalMemoTypeAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *internalMemoTypeController) FindInternalMemoTypes(c *gin.Context) {
	var (
		internalmemotypes []model.InternalMemoType
		response          helper.Response
	)
	internalmemotypes, err := b.internalMemoTypeService.FindInternalMemoTypes()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", internalmemotypes)
		c.JSON(http.StatusOK, response)
	}
}

func (b *internalMemoTypeController) FindInternalMemoTypesOffset(c *gin.Context) {
	var (
		internalmemotypes []model.InternalMemoType
		response          helper.Response
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
					internalmemotypes, err = b.internalMemoTypeService.FindInternalMemoTypesOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", internalmemotypes)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *internalMemoTypeController) SearchInternalMemoType(c *gin.Context) {
	var (
		internalmemotypes []model.InternalMemoType
		response          helper.Response
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
						internalmemotypes, err = b.internalMemoTypeService.SearchInternalMemoType(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", internalmemotypes)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *internalMemoTypeController) CountSearchInternalMemoType(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.internalMemoTypeService.CountSearchInternalMemoType(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoTypeController) FindInternalMemoTypeById(c *gin.Context) {
	var (
		internalMemoType model.InternalMemoType
		response         helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoType, err = b.internalMemoTypeService.FindInternalMemoTypeById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoTypeController) FindExcInternalMemoType(c *gin.Context) {
	var (
		internalMemoTypes []model.InternalMemoType
		response          helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoTypes, err = b.internalMemoTypeService.FindExcInternalMemoType(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoTypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoTypeController) InsertInternalMemoType(c *gin.Context) {
	var (
		internalMemoType                model.InternalMemoType
		response                        helper.Response
		CreateInternalMemoTypeParameter model.CreateInternalMemoTypeParameter
	)
	err := c.ShouldBindJSON(&CreateInternalMemoTypeParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		internalMemoType, err = b.internalMemoTypeService.InsertInternalMemoType(CreateInternalMemoTypeParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register internal memo type", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoTypeController) UpdateInternalMemoType(c *gin.Context) {
	var (
		newData  model.InternalMemoType
		oldData  model.InternalMemoType
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
			oldData, err = b.internalMemoTypeService.FindInternalMemoTypeById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.InternalMemoType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemoType, err := b.internalMemoTypeService.UpdateInternalMemoType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update internal memo type", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemoType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *internalMemoTypeController) DeleteInternalMemoType(c *gin.Context) {
	var (
		newData  model.InternalMemoType
		oldData  model.InternalMemoType
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
			oldData, err = b.internalMemoTypeService.FindInternalMemoTypeById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.InternalMemoType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemoType, err := b.internalMemoTypeService.DeleteInternalMemoType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete internal memo type", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemoType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
