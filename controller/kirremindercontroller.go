package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KIRReminderController interface {
	CountKIRReminderAll(c *gin.Context)
	FindKIRReminders(c *gin.Context)
	FindKIRRemindersOffset(c *gin.Context)
	SearchKIRReminder(c *gin.Context)
	CountSearchKIRReminder(c *gin.Context)
	FindKIRReminderById(c *gin.Context)
	FindExcKIRReminder(c *gin.Context)
	InsertKIRReminder(c *gin.Context)
	UpdateKIRReminder(c *gin.Context)
	DeleteKIRReminder(c *gin.Context)
}

type kirReminderController struct {
	kirReminderService service.KIRReminderService
	jwtService         service.JWTService
}

func NewKIRReminderController(kirReminderServ service.KIRReminderService, jwtServ service.JWTService) KIRReminderController {
	return &kirReminderController{
		kirReminderService: kirReminderServ,
		jwtService:         jwtServ,
	}
}

func (b *kirReminderController) CountKIRReminderAll(c *gin.Context) {
	var (
		response helper.Response
	)
	VehicleID, err := strconv.ParseInt(c.Param("vehicle_id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param vehicle_id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.kirReminderService.CountKIRReminderAll(int(VehicleID))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *kirReminderController) FindKIRReminders(c *gin.Context) {
	var (
		kirreminders []model.SelectKIRReminderParameter
		response     helper.Response
	)
	kirreminders, err := b.kirReminderService.FindKIRReminders()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", kirreminders)
		c.JSON(http.StatusOK, response)
	}
}

func (b *kirReminderController) FindKIRRemindersOffset(c *gin.Context) {
	var (
		kirreminders []model.SelectKIRReminderParameter
		response     helper.Response
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
						kirreminders, err = b.kirReminderService.FindKIRRemindersOffset(int(VehicleID), int(limit), int(offset), order, dir)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", kirreminders)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *kirReminderController) SearchKIRReminder(c *gin.Context) {
	var (
		kirreminders []model.SelectKIRReminderParameter
		response     helper.Response
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
							kirreminders, err = b.kirReminderService.SearchKIRReminder(int(VehicleID), int(limit), int(offset), order, dir, search)
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", kirreminders)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *kirReminderController) CountSearchKIRReminder(c *gin.Context) {
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
			count, err := b.kirReminderService.CountSearchKIRReminder(int(VehicleID), search)
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

func (b *kirReminderController) FindKIRReminderById(c *gin.Context) {
	var (
		kirReminder model.SelectKIRReminderParameter
		response    helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		kirReminder, err = b.kirReminderService.FindKIRReminderById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", kirReminder)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *kirReminderController) FindExcKIRReminder(c *gin.Context) {
	var (
		kirReminders []model.SelectKIRReminderParameter
		response     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		kirReminders, err = b.kirReminderService.FindExcKIRReminder(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", kirReminders)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *kirReminderController) InsertKIRReminder(c *gin.Context) {
	var (
		kirReminder                model.KIRReminder
		response                   helper.Response
		CreateKIRReminderParameter model.CreateKIRReminderParameter
	)
	err := c.ShouldBindJSON(&CreateKIRReminderParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		kirReminder, err = b.kirReminderService.InsertKIRReminder(CreateKIRReminderParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register business unit", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", kirReminder)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *kirReminderController) UpdateKIRReminder(c *gin.Context) {
	var (
		newData  model.KIRReminder
		oldData  model.SelectKIRReminderParameter
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
			oldData, err = b.kirReminderService.FindKIRReminderById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.SelectKIRReminderParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				kirReminder, err := b.kirReminderService.UpdateKIRReminder(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update business unit", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", kirReminder)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *kirReminderController) DeleteKIRReminder(c *gin.Context) {
	var (
		newData  model.KIRReminder
		oldData  model.SelectKIRReminderParameter
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
			oldData, err = b.kirReminderService.FindKIRReminderById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.SelectKIRReminderParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				kirReminder, err := b.kirReminderService.DeleteKIRReminder(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete business unit", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", kirReminder)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
