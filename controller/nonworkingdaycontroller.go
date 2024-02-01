package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NonWorkingDayController interface {
	CountNonWorkingDayAll(c *gin.Context)
	FindNonWorkingDays(c *gin.Context)
	FindNonWorkingDaysCuti(c *gin.Context)
	FindNonWorkingDaysAllDate(c *gin.Context)
	FindNonWorkingDaysOffset(c *gin.Context)
	SearchNonWorkingDay(c *gin.Context)
	CountNonWorkingDayExPersonalia(c *gin.Context)
	FindNonWorkingDaysExPersonaliaOffset(c *gin.Context)
	SearchNonWorkingDayExPersonalia(c *gin.Context)
	CountSearchNonWorkingDayExPersonalia(c *gin.Context)
	CountSearchNonWorkingDay(c *gin.Context)
	FindNonWorkingDayById(c *gin.Context)
	FindExcNonWorkingDay(c *gin.Context)
	FindNonWorkingDayByNWTId(c *gin.Context)
	CountNonWorkingDayName(c *gin.Context)
	FindNonWorkingDaybyDate(c *gin.Context)
	InsertNonWorkingDay(c *gin.Context)
	UpdateNonWorkingDay(c *gin.Context)
	DeleteNonWorkingDay(c *gin.Context)
}

type nonWorkingDayController struct {
	nonWorkingDayService service.NonWorkingDayService
	jwtService           service.JWTService
}

func NewNonWorkingDayController(nonWorkingDayServ service.NonWorkingDayService, jwtServ service.JWTService) NonWorkingDayController {
	return &nonWorkingDayController{
		nonWorkingDayService: nonWorkingDayServ,
		jwtService:           jwtServ,
	}
}

func (b *nonWorkingDayController) CountNonWorkingDayAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	count, err := b.nonWorkingDayService.CountNonWorkingDayAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *nonWorkingDayController) FindNonWorkingDaysCuti(c *gin.Context) {
	var (
		nonWorkingDays []model.SelectNonWorkingDayParameter
		response       helper.Response
	)
	nonWorkingDays, err := b.nonWorkingDayService.FindNonWorkingDaysCuti()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", nonWorkingDays)
		c.JSON(http.StatusOK, response)
	}
}

func (b *nonWorkingDayController) FindNonWorkingDaysAllDate(c *gin.Context) {
	var (
		nonWorkingDays []model.SelectNonWorkingDayParameter
		response       helper.Response
	)
	nonWorkingDays, err := b.nonWorkingDayService.FindNonWorkingDaysAllDate()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", nonWorkingDays)
		c.JSON(http.StatusOK, response)
	}
}

func (b *nonWorkingDayController) FindNonWorkingDays(c *gin.Context) {
	var (
		nonWorkingDays []model.SelectNonWorkingDayParameter
		response       helper.Response
	)
	nonWorkingDays, err := b.nonWorkingDayService.FindNonWorkingDays()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", nonWorkingDays)
		c.JSON(http.StatusOK, response)
	}
}

func (b *nonWorkingDayController) FindNonWorkingDaysOffset(c *gin.Context) {
	var (
		nonWorkingDays []model.SelectNonWorkingDayParameter
		response       helper.Response
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
					nonWorkingDays, err = b.nonWorkingDayService.FindNonWorkingDaysOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", nonWorkingDays)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *nonWorkingDayController) SearchNonWorkingDay(c *gin.Context) {
	var (
		nonWorkingDays []model.SelectNonWorkingDayParameter
		response       helper.Response
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
						nonWorkingDays, err = b.nonWorkingDayService.SearchNonWorkingDay(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", nonWorkingDays)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *nonWorkingDayController) CountSearchNonWorkingDay(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.nonWorkingDayService.CountSearchNonWorkingDay(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonWorkingDayController) FindNonWorkingDayById(c *gin.Context) {
	var (
		nonWorkingDay model.SelectNonWorkingDayParameter
		response      helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		nonWorkingDay, err = b.nonWorkingDayService.FindNonWorkingDayById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", nonWorkingDay)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonWorkingDayController) CountNonWorkingDayExPersonalia(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	count, err := b.nonWorkingDayService.CountNonWorkingDayExPersonalia()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}
func (b *nonWorkingDayController) FindNonWorkingDaysExPersonaliaOffset(c *gin.Context) {
	var (
		nonWorkingDays []model.SelectNonWorkingDayParameter
		response       helper.Response
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
					nonWorkingDays, err = b.nonWorkingDayService.FindNonWorkingDaysExPersonaliaOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", nonWorkingDays)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *nonWorkingDayController) SearchNonWorkingDayExPersonalia(c *gin.Context) {
	var (
		nonWorkingDays []model.SelectNonWorkingDayParameter
		response       helper.Response
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
						nonWorkingDays, err = b.nonWorkingDayService.SearchNonWorkingDayExPersonalia(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", nonWorkingDays)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *nonWorkingDayController) CountSearchNonWorkingDayExPersonalia(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.nonWorkingDayService.CountSearchNonWorkingDayExPersonalia(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonWorkingDayController) FindExcNonWorkingDay(c *gin.Context) {
	var (
		nonWorkingDays []model.SelectNonWorkingDayParameter
		response       helper.Response
	)
	nwtId, err := strconv.ParseUint(c.Param("nwtId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param non working type id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		id, err := strconv.ParseUint(c.Param("id"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			nonWorkingDays, err = b.nonWorkingDayService.FindExcNonWorkingDay(uint(nwtId), uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", nonWorkingDays)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *nonWorkingDayController) FindNonWorkingDayByNWTId(c *gin.Context) {
	var (
		nonWorkingDays []model.SelectNonWorkingDayParameter
		response       helper.Response
	)
	nwtId, err := strconv.ParseUint(c.Param("nwtId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param non working type id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		nonWorkingDays, err = b.nonWorkingDayService.FindNonWorkingDayByNWTId(uint(nwtId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", nonWorkingDays)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonWorkingDayController) CountNonWorkingDayName(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.nonWorkingDayService.CountNonWorkingDayName(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonWorkingDayController) FindNonWorkingDaybyDate(c *gin.Context) {
	var (
		nonWorkingDays []model.SelectNonWorkingDayParameter
		response       helper.Response
	)
	date, err := strconv.ParseFloat(c.Param("date"), 64)
	if err != nil {
		response = helper.BuildErrorResponse("No param date was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		nonWorkingDays, err = b.nonWorkingDayService.FindNonWorkingDaybyDate(float64(date))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", nonWorkingDays)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonWorkingDayController) InsertNonWorkingDay(c *gin.Context) {
	var (
		nonWorkingDay                model.NonWorkingDay
		response                     helper.Response
		CreateNonWorkingDayParameter model.CreateNonWorkingDayParameter
	)
	err := c.ShouldBindJSON(&CreateNonWorkingDayParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		nonWorkingDay, err = b.nonWorkingDayService.InsertNonWorkingDay(CreateNonWorkingDayParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register nonWorkingDay", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", nonWorkingDay)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *nonWorkingDayController) UpdateNonWorkingDay(c *gin.Context) {
	var (
		newData  model.NonWorkingDay
		oldData  model.SelectNonWorkingDayParameter
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
			oldData, err = b.nonWorkingDayService.FindNonWorkingDayById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectNonWorkingDayParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				nonWorkingDay, err := b.nonWorkingDayService.UpdateNonWorkingDay(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update nonWorkingDay", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", nonWorkingDay)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *nonWorkingDayController) DeleteNonWorkingDay(c *gin.Context) {
	var (
		newData  model.NonWorkingDay
		oldData  model.SelectNonWorkingDayParameter
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
			oldData, err = b.nonWorkingDayService.FindNonWorkingDayById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectNonWorkingDayParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				nonWorkingDay, err := b.nonWorkingDayService.DeleteNonWorkingDay(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete nonWorkingDay", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", nonWorkingDay)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
