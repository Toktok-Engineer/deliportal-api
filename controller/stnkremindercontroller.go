package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type STNKReminderController interface {
	CountSTNKReminderAll(c *gin.Context)
	FindSTNKReminders(c *gin.Context)
	FindSTNKRemindersOffset(c *gin.Context)
	SearchSTNKReminder(c *gin.Context)
	CountSearchSTNKReminder(c *gin.Context)
	FindSTNKReminderById(c *gin.Context)
	FindExcSTNKReminder(c *gin.Context)
	InsertSTNKReminder(c *gin.Context)
	UpdateSTNKReminder(c *gin.Context)
	DeleteSTNKReminder(c *gin.Context)
}

type stnkReminderController struct {
	stnkReminderService service.STNKReminderService
	jwtService          service.JWTService
}

func NewSTNKReminderController(stnkReminderServ service.STNKReminderService, jwtServ service.JWTService) STNKReminderController {
	return &stnkReminderController{
		stnkReminderService: stnkReminderServ,
		jwtService:          jwtServ,
	}
}

func (b *stnkReminderController) CountSTNKReminderAll(c *gin.Context) {
	var (
		response helper.Response
	)
	VehicleID, err := strconv.ParseInt(c.Param("vehicle_id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param vehicle_id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.stnkReminderService.CountSTNKReminderAll(int(VehicleID))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *stnkReminderController) FindSTNKReminders(c *gin.Context) {
	var (
		stnkreminders []model.SelectSTNKReminderParameter
		response      helper.Response
	)
	stnkreminders, err := b.stnkReminderService.FindSTNKReminders()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", stnkreminders)
		c.JSON(http.StatusOK, response)
	}
}

func (b *stnkReminderController) FindSTNKRemindersOffset(c *gin.Context) {
	var (
		stnkreminders []model.SelectSTNKReminderParameter
		response      helper.Response
	)

	VehicleID, err := strconv.ParseInt(c.Param("vehicle_id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param vehicle_id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
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
						stnkreminders, err = b.stnkReminderService.FindSTNKRemindersOffset(int(VehicleID), int(limit), int(offset), order, dir)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", stnkreminders)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *stnkReminderController) SearchSTNKReminder(c *gin.Context) {
	var (
		stnkreminders []model.SelectSTNKReminderParameter
		response      helper.Response
	)

	VehicleID, err := strconv.ParseInt(c.Param("vehicle_id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param vehicle_id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
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
							stnkreminders, err = b.stnkReminderService.SearchSTNKReminder(int(VehicleID), int(limit), int(offset), order, dir, search)
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", stnkreminders)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *stnkReminderController) CountSearchSTNKReminder(c *gin.Context) {
	var (
		response helper.Response
	)

	VehicleID, err := strconv.ParseInt(c.Param("vehicle_id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param vehicle_id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		search := c.Param("search")
		if search == "" {
			response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.stnkReminderService.CountSearchSTNKReminder(int(VehicleID), search)
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

func (b *stnkReminderController) FindSTNKReminderById(c *gin.Context) {
	var (
		stnkReminder model.SelectSTNKReminderParameter
		response     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		stnkReminder, err = b.stnkReminderService.FindSTNKReminderById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", stnkReminder)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *stnkReminderController) FindExcSTNKReminder(c *gin.Context) {
	var (
		stnkReminders []model.SelectSTNKReminderParameter
		response      helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		stnkReminders, err = b.stnkReminderService.FindExcSTNKReminder(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", stnkReminders)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *stnkReminderController) InsertSTNKReminder(c *gin.Context) {
	var (
		stnkReminder                model.STNKReminder
		response                    helper.Response
		CreateSTNKReminderParameter model.CreateSTNKReminderParameter
	)
	err := c.ShouldBindJSON(&CreateSTNKReminderParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		stnkReminder, err = b.stnkReminderService.InsertSTNKReminder(CreateSTNKReminderParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register business unit", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", stnkReminder)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *stnkReminderController) UpdateSTNKReminder(c *gin.Context) {
	var (
		newData  model.STNKReminder
		oldData  model.SelectSTNKReminderParameter
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
			oldData, err = b.stnkReminderService.FindSTNKReminderById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.SelectSTNKReminderParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				stnkReminder, err := b.stnkReminderService.UpdateSTNKReminder(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update business unit", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", stnkReminder)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *stnkReminderController) DeleteSTNKReminder(c *gin.Context) {
	var (
		newData  model.STNKReminder
		oldData  model.SelectSTNKReminderParameter
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
			oldData, err = b.stnkReminderService.FindSTNKReminderById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.SelectSTNKReminderParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				stnkReminder, err := b.stnkReminderService.DeleteSTNKReminder(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete business unit", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", stnkReminder)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
