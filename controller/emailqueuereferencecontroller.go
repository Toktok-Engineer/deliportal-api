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

type EmailQueueReferenceController interface {
	FindEmailQueueReferences(c *gin.Context)
	FindEmailQueueReferenceById(c *gin.Context)
	FindExcEmailQueueReference(c *gin.Context)
	InsertEmailQueueReference(c *gin.Context)
	UpdateEmailQueueReference(c *gin.Context)
	DeleteEmailQueueReference(c *gin.Context)
}

type emailQueueReferenceController struct {
	emailQueueReferenceService service.EmailQueueReferenceService
	jwtService                 service.JWTService
}

func NewEmailQueueReferenceController(emailQueueReferenceServ service.EmailQueueReferenceService, jwtServ service.JWTService) EmailQueueReferenceController {
	return &emailQueueReferenceController{
		emailQueueReferenceService: emailQueueReferenceServ,
		jwtService:                 jwtServ,
	}
}

func (b *emailQueueReferenceController) FindEmailQueueReferences(c *gin.Context) {
	var (
		emailQueueReferences []model.SelectEmailQueueReferenceParameter
		response             helper.Response
	)
	emailQueueReferences, err := b.emailQueueReferenceService.FindEmailQueueReferences()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", emailQueueReferences)
		c.JSON(http.StatusOK, response)
	}
}

func (b *emailQueueReferenceController) FindEmailQueueReferenceById(c *gin.Context) {
	var (
		emailQueueReference model.SelectEmailQueueReferenceParameter
		response            helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		emailQueueReference, err = b.emailQueueReferenceService.FindEmailQueueReferenceById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", emailQueueReference)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *emailQueueReferenceController) FindExcEmailQueueReference(c *gin.Context) {
	var (
		emailQueueReferences []model.SelectEmailQueueReferenceParameter
		response             helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		emailQueueReferences, err = b.emailQueueReferenceService.FindExcEmailQueueReference(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", emailQueueReferences)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *emailQueueReferenceController) InsertEmailQueueReference(c *gin.Context) {
	var (
		emailQueueReference                model.EmailQueueReference
		response                           helper.Response
		CreateEmailQueueReferenceParameter model.CreateEmailQueueReferenceParameter
	)
	err := c.ShouldBindJSON(&CreateEmailQueueReferenceParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		emailQueueReference, err = b.emailQueueReferenceService.InsertEmailQueueReference(CreateEmailQueueReferenceParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register emailQueueReference", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", emailQueueReference)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *emailQueueReferenceController) UpdateEmailQueueReference(c *gin.Context) {
	var (
		newData  model.EmailQueueReference
		oldData  model.SelectEmailQueueReferenceParameter
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
			oldData, err = b.emailQueueReferenceService.FindEmailQueueReferenceById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectEmailQueueReferenceParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				emailQueueReference, err := b.emailQueueReferenceService.UpdateEmailQueueReference(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update emailQueueReference", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", emailQueueReference)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *emailQueueReferenceController) DeleteEmailQueueReference(c *gin.Context) {
	var (
		newData  model.EmailQueueReference
		oldData  model.SelectEmailQueueReferenceParameter
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
			oldData, err = b.emailQueueReferenceService.FindEmailQueueReferenceById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectEmailQueueReferenceParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				emailQueueReference, err := b.emailQueueReferenceService.DeleteEmailQueueReference(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete emailQueueReference", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", emailQueueReference)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
