package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmailQueueController interface {
	FindEmailQueues(c *gin.Context)
	FindEmailQueueById(c *gin.Context)
	FindExcEmailQueue(c *gin.Context)
	FindEmailQueueByStatus(c *gin.Context)
	InsertEmailQueue(c *gin.Context)
	UpdateEmailQueue(c *gin.Context)
	DeleteEmailQueue(c *gin.Context)
}

type emailQueueController struct {
	emailQueueService service.EmailQueueService
	jwtService        service.JWTService
}

func NewEmailQueueController(emailQueueServ service.EmailQueueService, jwtServ service.JWTService) EmailQueueController {
	return &emailQueueController{
		emailQueueService: emailQueueServ,
		jwtService:        jwtServ,
	}
}

func (b *emailQueueController) FindEmailQueues(c *gin.Context) {
	var (
		emailQueues []model.SelectEmailQueueParameter
		response    helper.Response
	)
	emailQueues, err := b.emailQueueService.FindEmailQueues()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", emailQueues)
		c.JSON(http.StatusOK, response)
	}
}

func (b *emailQueueController) FindEmailQueueById(c *gin.Context) {
	var (
		emailQueue model.SelectEmailQueueParameter
		response   helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		emailQueue, err = b.emailQueueService.FindEmailQueueById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", emailQueue)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *emailQueueController) FindExcEmailQueue(c *gin.Context) {
	var (
		emailQueues []model.SelectEmailQueueParameter
		response    helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		emailQueues, err = b.emailQueueService.FindExcEmailQueue(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", emailQueues)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *emailQueueController) FindEmailQueueByStatus(c *gin.Context) {
	var (
		emailQueues []model.SelectEmailQueueParameter
		response    helper.Response
	)
	status, err := strconv.ParseUint(c.Param("status"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param status was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		emailQueues, err = b.emailQueueService.FindEmailQueueByStatus(uint(status))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", emailQueues)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *emailQueueController) InsertEmailQueue(c *gin.Context) {
	var (
		emailQueue                model.EmailQueue
		response                  helper.Response
		CreateEmailQueueParameter model.CreateEmailQueueParameter
	)
	err := c.ShouldBindJSON(&CreateEmailQueueParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		emailQueue, err = b.emailQueueService.InsertEmailQueue(CreateEmailQueueParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register emailQueue", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", emailQueue)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *emailQueueController) UpdateEmailQueue(c *gin.Context) {
	var (
		newData  model.EmailQueue
		oldData  model.SelectEmailQueueParameter
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
			oldData, err = b.emailQueueService.FindEmailQueueById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectEmailQueueParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				emailQueue, err := b.emailQueueService.UpdateEmailQueue(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update emailQueue", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", emailQueue)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *emailQueueController) DeleteEmailQueue(c *gin.Context) {
	var (
		newData  model.EmailQueue
		oldData  model.SelectEmailQueueParameter
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
			oldData, err = b.emailQueueService.FindEmailQueueById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectEmailQueueParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				emailQueue, err := b.emailQueueService.DeleteEmailQueue(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete emailQueue", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", emailQueue)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
