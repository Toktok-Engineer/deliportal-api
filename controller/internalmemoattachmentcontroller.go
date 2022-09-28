package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InternalMemoAttachmentController interface {
	CountInternalMemoAttachmentAll(c *gin.Context)
	FindInternalMemoAttachments(c *gin.Context)
	FindInternalMemoAttachmentsOffset(c *gin.Context)
	SearchInternalMemoAttachment(c *gin.Context)
	CountSearchInternalMemoAttachment(c *gin.Context)
	FindInternalMemoAttachmentById(c *gin.Context)
	FindExcInternalMemoAttachment(c *gin.Context)
	FindInternalMemoAttachmentByCompanyId(c *gin.Context)
	InsertInternalMemoAttachment(c *gin.Context)
	UpdateInternalMemoAttachment(c *gin.Context)
	DeleteInternalMemoAttachment(c *gin.Context)
}

type internalMemoAttachmentController struct {
	internalMemoAttachmentService service.InternalMemoAttachmentService
	jwtService                    service.JWTService
}

func NewInternalMemoAttachmentController(internalMemoAttachmentServ service.InternalMemoAttachmentService, jwtServ service.JWTService) InternalMemoAttachmentController {
	return &internalMemoAttachmentController{
		internalMemoAttachmentService: internalMemoAttachmentServ,
		jwtService:                    jwtServ,
	}
}

func (b *internalMemoAttachmentController) CountInternalMemoAttachmentAll(c *gin.Context) {
	var (
		response helper.Response
	)
	internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.internalMemoAttachmentService.CountInternalMemoAttachmentAll(int(internalMemoId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoAttachmentController) FindInternalMemoAttachments(c *gin.Context) {
	var (
		internalMemoAttachments []model.InternalMemoAttachment
		response                helper.Response
	)
	internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoAttachments, err = b.internalMemoAttachmentService.FindInternalMemoAttachments(int(internalMemoId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoAttachments)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoAttachmentController) FindInternalMemoAttachmentsOffset(c *gin.Context) {
	var (
		internalMemoAttachments []model.SelectInternalMemoAttachmentParameter
		response                helper.Response
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
					internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						internalMemoAttachments, err = b.internalMemoAttachmentService.FindInternalMemoAttachmentsOffset(int(limit), int(offset), order, dir, int(internalMemoId))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", internalMemoAttachments)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *internalMemoAttachmentController) SearchInternalMemoAttachment(c *gin.Context) {
	var (
		internalMemoAttachments []model.SelectInternalMemoAttachmentParameter
		response                helper.Response
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
						internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							internalMemoAttachments, err = b.internalMemoAttachmentService.SearchInternalMemoAttachment(int(limit), int(offset), order, dir, search, int(internalMemoId))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", internalMemoAttachments)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *internalMemoAttachmentController) CountSearchInternalMemoAttachment(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoId, err := strconv.ParseInt(c.Param("internalMemoId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param internalMemoId was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.internalMemoAttachmentService.CountSearchInternalMemoAttachment(search, int(internalMemoId))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", count)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *internalMemoAttachmentController) FindInternalMemoAttachmentById(c *gin.Context) {
	var (
		internalMemoAttachment model.SelectInternalMemoAttachmentParameter
		response               helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoAttachment, err = b.internalMemoAttachmentService.FindInternalMemoAttachmentById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoAttachment)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoAttachmentController) FindExcInternalMemoAttachment(c *gin.Context) {
	var (
		internalMemoAttachments []model.SelectInternalMemoAttachmentParameter
		response                helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoAttachments, err = b.internalMemoAttachmentService.FindExcInternalMemoAttachment(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoAttachments)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoAttachmentController) FindInternalMemoAttachmentByCompanyId(c *gin.Context) {
	var (
		internalMemoAttachments []model.SelectInternalMemoAttachmentParameter
		response                helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param company id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemoAttachments, err = b.internalMemoAttachmentService.FindInternalMemoAttachmentByCompanyId(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoAttachments)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoAttachmentController) InsertInternalMemoAttachment(c *gin.Context) {
	var (
		internalMemoAttachment                model.InternalMemoAttachment
		response                              helper.Response
		CreateInternalMemoAttachmentParameter model.CreateInternalMemoAttachmentParameter
	)
	err := c.ShouldBindJSON(&CreateInternalMemoAttachmentParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		internalMemoAttachment, err = b.internalMemoAttachmentService.InsertInternalMemoAttachment(CreateInternalMemoAttachmentParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register internalMemoAttachment", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemoAttachment)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoAttachmentController) UpdateInternalMemoAttachment(c *gin.Context) {
	var (
		newData  model.InternalMemoAttachment
		oldData  model.SelectInternalMemoAttachmentParameter
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
			oldData, err = b.internalMemoAttachmentService.FindInternalMemoAttachmentById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectInternalMemoAttachmentParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemoAttachment, err := b.internalMemoAttachmentService.UpdateInternalMemoAttachment(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update internalMemoAttachment", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemoAttachment)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *internalMemoAttachmentController) DeleteInternalMemoAttachment(c *gin.Context) {
	var (
		newData  model.InternalMemoAttachment
		oldData  model.SelectInternalMemoAttachmentParameter
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
			oldData, err = b.internalMemoAttachmentService.FindInternalMemoAttachmentById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectInternalMemoAttachmentParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemoAttachment, err := b.internalMemoAttachmentService.DeleteInternalMemoAttachment(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete internalMemoAttachment", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemoAttachment)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
