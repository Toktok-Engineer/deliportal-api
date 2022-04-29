package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmailQueueTypeController interface {
	FindEmailQueueTypes(c *gin.Context)
	FindEmailQueueTypeById(c *gin.Context)
	FindExcEmailQueueType(c *gin.Context)
	InsertEmailQueueType(c *gin.Context)
	UpdateEmailQueueType(c *gin.Context)
	DeleteEmailQueueType(c *gin.Context)
}

type emailQueueTypeController struct {
	emailQueueTypeService service.EmailQueueTypeService
	jwtService            service.JWTService
}

func NewEmailQueueTypeController(emailQueueTypeServ service.EmailQueueTypeService, jwtServ service.JWTService) EmailQueueTypeController {
	return &emailQueueTypeController{
		emailQueueTypeService: emailQueueTypeServ,
		jwtService:            jwtServ,
	}
}

func (b *emailQueueTypeController) FindEmailQueueTypes(c *gin.Context) {
	var (
		emailQueueTypes []model.EmailQueueType
		response        helper.Response
	)
	emailQueueTypes, err := b.emailQueueTypeService.FindEmailQueueTypes()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", emailQueueTypes)
		c.JSON(http.StatusOK, response)
	}
}

func (b *emailQueueTypeController) FindEmailQueueTypeById(c *gin.Context) {
	var (
		emailQueueType model.EmailQueueType
		response       helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		emailQueueType, err = b.emailQueueTypeService.FindEmailQueueTypeById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", emailQueueType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *emailQueueTypeController) FindExcEmailQueueType(c *gin.Context) {
	var (
		emailQueueTypes []model.EmailQueueType
		response        helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		emailQueueTypes, err = b.emailQueueTypeService.FindExcEmailQueueType(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", emailQueueTypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *emailQueueTypeController) InsertEmailQueueType(c *gin.Context) {
	var (
		emailQueueType                model.EmailQueueType
		response                      helper.Response
		CreateEmailQueueTypeParameter model.CreateEmailQueueTypeParameter
	)
	err := c.ShouldBindJSON(&CreateEmailQueueTypeParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		emailQueueType, err = b.emailQueueTypeService.InsertEmailQueueType(CreateEmailQueueTypeParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register emailQueueType", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", emailQueueType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *emailQueueTypeController) UpdateEmailQueueType(c *gin.Context) {
	var (
		newData  model.EmailQueueType
		oldData  model.EmailQueueType
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
			oldData, err = b.emailQueueTypeService.FindEmailQueueTypeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.EmailQueueType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				emailQueueType, err := b.emailQueueTypeService.UpdateEmailQueueType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update emailQueueType", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", emailQueueType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *emailQueueTypeController) DeleteEmailQueueType(c *gin.Context) {
	var (
		newData  model.EmailQueueType
		oldData  model.EmailQueueType
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
			oldData, err = b.emailQueueTypeService.FindEmailQueueTypeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.EmailQueueType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				emailQueueType, err := b.emailQueueTypeService.DeleteEmailQueueType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete emailQueueType", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", emailQueueType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
