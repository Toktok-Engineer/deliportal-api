package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InternalMemoTracingController interface {
	CountInternalMemoTracingAll(c *gin.Context)
	FindInternalMemoTracings(c *gin.Context)
	FindInternalMemoTracingById(c *gin.Context)
	FindExcInternalMemoTracing(c *gin.Context)
	InsertInternalMemoTracing(c *gin.Context)
	UpdateInternalMemoTracing(c *gin.Context)
	DeleteInternalMemoTracing(c *gin.Context)
}

type internalMemoTracingController struct {
	internalMemoTracingService service.InternalMemoTracingService
	jwtService                 service.JWTService
}

func NewInternalMemoTracingController(internalMemoTracingServ service.InternalMemoTracingService, jwtServ service.JWTService) InternalMemoTracingController {
	return &internalMemoTracingController{
		internalMemoTracingService: internalMemoTracingServ,
		jwtService:                 jwtServ,
	}
}

func (b *internalMemoTracingController) CountInternalMemoTracingAll(c *gin.Context) {
	var (
		response helper.Response
	)
	internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.internalMemoTracingService.CountInternalMemoTracingAll(int(internalMemoId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoTracingController) FindInternalMemoTracings(c *gin.Context) {
	var (
		internalmemotracings []model.InternalMemoTracing
		response             helper.Response
	)
	internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalmemotracings, err = b.internalMemoTracingService.FindInternalMemoTracings(int(internalMemoId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalmemotracings)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoTracingController) FindInternalMemoTracingById(c *gin.Context) {
	var (
		internalMemoTracing model.InternalMemoTracing
		response            helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoTracing, err = b.internalMemoTracingService.FindInternalMemoTracingById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoTracing)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoTracingController) FindExcInternalMemoTracing(c *gin.Context) {
	var (
		internalMemoTracings []model.InternalMemoTracing
		response             helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoTracings, err = b.internalMemoTracingService.FindExcInternalMemoTracing(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoTracings)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoTracingController) InsertInternalMemoTracing(c *gin.Context) {
	var (
		internalMemoTracing                model.InternalMemoTracing
		response                           helper.Response
		CreateInternalMemoTracingParameter model.CreateInternalMemoTracingParameter
	)
	err := c.ShouldBindJSON(&CreateInternalMemoTracingParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		internalMemoTracing, err = b.internalMemoTracingService.InsertInternalMemoTracing(CreateInternalMemoTracingParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register internal memo tracing", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoTracing)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoTracingController) UpdateInternalMemoTracing(c *gin.Context) {
	var (
		newData  model.InternalMemoTracing
		oldData  model.InternalMemoTracing
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
			oldData, err = b.internalMemoTracingService.FindInternalMemoTracingById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.InternalMemoTracing{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemoTracing, err := b.internalMemoTracingService.UpdateInternalMemoTracing(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update internal memo tracing", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemoTracing)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *internalMemoTracingController) DeleteInternalMemoTracing(c *gin.Context) {
	var (
		newData  model.InternalMemoTracing
		oldData  model.InternalMemoTracing
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
			oldData, err = b.internalMemoTracingService.FindInternalMemoTracingById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.InternalMemoTracing{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemoTracing, err := b.internalMemoTracingService.DeleteInternalMemoTracing(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete internal memo tracing", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemoTracing)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
