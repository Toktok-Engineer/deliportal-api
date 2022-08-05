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
	CountEmailQueueReferenceAll(c *gin.Context)
	FindEmailQueueReferences(c *gin.Context)
	FindEmailQueueReferencesOffset(c *gin.Context)
	SearchEmailQueueReference(c *gin.Context)
	CountSearchEmailQueueReference(c *gin.Context)
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

func (b *emailQueueReferenceController) CountEmailQueueReferenceAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	count, err := b.emailQueueReferenceService.CountEmailQueueReferenceAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
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

func (b *emailQueueReferenceController) FindEmailQueueReferencesOffset(c *gin.Context) {
	var (
		emailQueueReferences []model.SelectEmailQueueReferenceParameter
		response             helper.Response
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
					emailQueueReferences, err = b.emailQueueReferenceService.FindEmailQueueReferencesOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", emailQueueReferences)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *emailQueueReferenceController) SearchEmailQueueReference(c *gin.Context) {
	var (
		emailQueueReferences []model.SelectEmailQueueReferenceParameter
		response             helper.Response
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
						emailQueueReferences, err = b.emailQueueReferenceService.SearchEmailQueueReference(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", emailQueueReferences)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *emailQueueReferenceController) CountSearchEmailQueueReference(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.emailQueueReferenceService.CountSearchEmailQueueReference(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
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
