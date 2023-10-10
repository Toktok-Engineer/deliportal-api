package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LeaveTypesController interface {
	CountLeaveTypesAll(c *gin.Context)
	FindLeaveTypess(c *gin.Context)
	FindLeaveTypessOffset(c *gin.Context)
	SearchLeaveTypes(c *gin.Context)
	CountSearchLeaveTypes(c *gin.Context)
	FindLeaveTypesById(c *gin.Context)
	CountLeaveTypesName(c *gin.Context)
	FindExcLeaveTypes(c *gin.Context)
	InsertLeaveTypes(c *gin.Context)
	UpdateLeaveTypes(c *gin.Context)
	DeleteLeaveTypes(c *gin.Context)
}

type leaveTypesController struct {
	leaveTypesService service.LeaveTypesService
	jwtService        service.JWTService
}

func NewLeaveTypesController(leaveTypesServ service.LeaveTypesService, jwtServ service.JWTService) LeaveTypesController {
	return &leaveTypesController{
		leaveTypesService: leaveTypesServ,
		jwtService:        jwtServ,
	}
}

func (b *leaveTypesController) CountLeaveTypesAll(c *gin.Context) {
	var (
		response helper.Response
	)

	count, err := b.leaveTypesService.CountLeaveTypesAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *leaveTypesController) FindLeaveTypess(c *gin.Context) {
	var (
		leaveTypess []model.LeaveTypes
		response    helper.Response
	)
	leaveTypess, err := b.leaveTypesService.FindLeaveTypess()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", leaveTypess)
		c.JSON(http.StatusOK, response)
	}
}

func (b *leaveTypesController) FindLeaveTypessOffset(c *gin.Context) {
	var (
		leaveTypess []model.LeaveTypes
		response    helper.Response
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
					leaveTypess, err = b.leaveTypesService.FindLeaveTypessOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", leaveTypess)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *leaveTypesController) SearchLeaveTypes(c *gin.Context) {
	var (
		leaveTypess []model.LeaveTypes
		response    helper.Response
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
						leaveTypess, err = b.leaveTypesService.SearchLeaveTypes(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", leaveTypess)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *leaveTypesController) CountSearchLeaveTypes(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.leaveTypesService.CountSearchLeaveTypes(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *leaveTypesController) FindLeaveTypesById(c *gin.Context) {
	var (
		leaveTypes model.LeaveTypes
		response   helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		leaveTypes, err = b.leaveTypesService.FindLeaveTypesById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", leaveTypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *leaveTypesController) CountLeaveTypesName(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.leaveTypesService.CountLeaveTypesName(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *leaveTypesController) FindExcLeaveTypes(c *gin.Context) {
	var (
		leaveTypess []model.LeaveTypes
		response    helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		leaveTypess, err = b.leaveTypesService.FindExcLeaveTypes(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", leaveTypess)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *leaveTypesController) InsertLeaveTypes(c *gin.Context) {
	var (
		leaveTypes                model.LeaveTypes
		response                  helper.Response
		CreateLeaveTypesParameter model.CreateLeaveTypesParameter
	)
	err := c.ShouldBindJSON(&CreateLeaveTypesParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		leaveTypes, err = b.leaveTypesService.InsertLeaveTypes(CreateLeaveTypesParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register leaveTypes", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", leaveTypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *leaveTypesController) UpdateLeaveTypes(c *gin.Context) {
	var (
		newData  model.LeaveTypes
		oldData  model.LeaveTypes
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
			oldData, err = b.leaveTypesService.FindLeaveTypesById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.LeaveTypes{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				leaveTypes, err := b.leaveTypesService.UpdateLeaveTypes(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update leaveTypes", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", leaveTypes)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *leaveTypesController) DeleteLeaveTypes(c *gin.Context) {
	var (
		newData  model.LeaveTypes
		oldData  model.LeaveTypes
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
			oldData, err = b.leaveTypesService.FindLeaveTypesById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.LeaveTypes{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				leaveTypes, err := b.leaveTypesService.DeleteLeaveTypes(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete leaveTypes", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", leaveTypes)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
