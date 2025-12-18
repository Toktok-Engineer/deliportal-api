package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AsuransiReminderController interface {
	CountAsuransiReminderAll(c *gin.Context)
	FindAsuransiReminders(c *gin.Context)
	FindAsuransiRemindersOffset(c *gin.Context)
	SearchAsuransiReminder(c *gin.Context)
	CountSearchAsuransiReminder(c *gin.Context)
	FindAsuransiReminderById(c *gin.Context)
	FindExcAsuransiReminder(c *gin.Context)
	InsertAsuransiReminder(c *gin.Context)
	UpdateAsuransiReminder(c *gin.Context)
	DeleteAsuransiReminder(c *gin.Context)
}

type asuransiReminderController struct {
	asuransiReminderService service.AsuransiReminderService
	jwtService              service.JWTService
}

func NewAsuransiReminderController(asuransiReminderServ service.AsuransiReminderService, jwtServ service.JWTService) AsuransiReminderController {
	return &asuransiReminderController{
		asuransiReminderService: asuransiReminderServ,
		jwtService:              jwtServ,
	}
}

func (b *asuransiReminderController) CountAsuransiReminderAll(c *gin.Context) {
	var (
		response helper.Response
	)
	VehicleID, err := strconv.ParseInt(c.Param("vehicle_id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param vehicle_id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.asuransiReminderService.CountAsuransiReminderAll(int(VehicleID))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiReminderController) FindAsuransiReminders(c *gin.Context) {
	var (
		asuransireminders []model.SelectAsuransiReminderParameter
		response          helper.Response
	)
	asuransireminders, err := b.asuransiReminderService.FindAsuransiReminders()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", asuransireminders)
		c.JSON(http.StatusOK, response)
	}
}

func (b *asuransiReminderController) FindAsuransiRemindersOffset(c *gin.Context) {
	var (
		asuransireminders []model.SelectAsuransiReminderParameter
		response          helper.Response
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
						asuransireminders, err = b.asuransiReminderService.FindAsuransiRemindersOffset(int(VehicleID), int(limit), int(offset), order, dir)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", asuransireminders)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *asuransiReminderController) SearchAsuransiReminder(c *gin.Context) {
	var (
		asuransireminders []model.SelectAsuransiReminderParameter
		response          helper.Response
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
							asuransireminders, err = b.asuransiReminderService.SearchAsuransiReminder(int(VehicleID), int(limit), int(offset), order, dir, search)
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", asuransireminders)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *asuransiReminderController) CountSearchAsuransiReminder(c *gin.Context) {
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
			count, err := b.asuransiReminderService.CountSearchAsuransiReminder(int(VehicleID), search)
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

func (b *asuransiReminderController) FindAsuransiReminderById(c *gin.Context) {
	var (
		asuransiReminder model.SelectAsuransiReminderParameter
		response         helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		asuransiReminder, err = b.asuransiReminderService.FindAsuransiReminderById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", asuransiReminder)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiReminderController) FindExcAsuransiReminder(c *gin.Context) {
	var (
		asuransiReminders []model.SelectAsuransiReminderParameter
		response          helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		asuransiReminders, err = b.asuransiReminderService.FindExcAsuransiReminder(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", asuransiReminders)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiReminderController) InsertAsuransiReminder(c *gin.Context) {
	var (
		asuransiReminder                model.AsuransiReminder
		response                        helper.Response
		CreateAsuransiReminderParameter model.CreateAsuransiReminderParameter
	)
	err := c.ShouldBindJSON(&CreateAsuransiReminderParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		asuransiReminder, err = b.asuransiReminderService.InsertAsuransiReminder(CreateAsuransiReminderParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register business unit", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", asuransiReminder)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *asuransiReminderController) UpdateAsuransiReminder(c *gin.Context) {
	var (
		newData  model.AsuransiReminder
		oldData  model.SelectAsuransiReminderParameter
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
			oldData, err = b.asuransiReminderService.FindAsuransiReminderById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.SelectAsuransiReminderParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				asuransiReminder, err := b.asuransiReminderService.UpdateAsuransiReminder(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update business unit", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", asuransiReminder)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *asuransiReminderController) DeleteAsuransiReminder(c *gin.Context) {
	var (
		newData  model.AsuransiReminder
		oldData  model.SelectAsuransiReminderParameter
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
			oldData, err = b.asuransiReminderService.FindAsuransiReminderById(uint(id))
			if err != nil {
				res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, res)
			} else if (oldData == model.SelectAsuransiReminderParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				asuransiReminder, err := b.asuransiReminderService.DeleteAsuransiReminder(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete business unit", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", asuransiReminder)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
