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

type InternalMemoController interface {
	CountInternalMemoAll(c *gin.Context)
	FindInternalMemos(c *gin.Context)
	FindInternalMemosOffset(c *gin.Context)
	SearchInternalMemo(c *gin.Context)
	CountSearchInternalMemo(c *gin.Context)
	CountInternalMemoByDept(c *gin.Context)
	FindInternalMemoById(c *gin.Context)
	FindExcInternalMemo(c *gin.Context)
	InsertInternalMemo(c *gin.Context)
	UpdateInternalMemo(c *gin.Context)
	DeleteInternalMemo(c *gin.Context)
	UpdateInternalMemoApprove(c *gin.Context)
	UploadIMDocument(c *gin.Context)
}

type internalMemoController struct {
	internalMemoService service.InternalMemoService
	jwtService          service.JWTService
}

func NewInternalMemoController(internalMemoServ service.InternalMemoService, jwtServ service.JWTService) InternalMemoController {
	return &internalMemoController{
		internalMemoService: internalMemoServ,
		jwtService:          jwtServ,
	}
}

func (b *internalMemoController) CountInternalMemoAll(c *gin.Context) {
	var (
		response helper.Response
	)

	employeeID, err := strconv.ParseInt(c.Param("employeeID"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param employeeID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.internalMemoService.CountInternalMemoAll(int(employeeID))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoController) FindInternalMemos(c *gin.Context) {
	var (
		internalMemos []model.SelectInternalMemoParameter
		response      helper.Response
	)
	internalMemos, err := b.internalMemoService.FindInternalMemos()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", internalMemos)
		c.JSON(http.StatusOK, response)
	}
}

func (b *internalMemoController) FindInternalMemosOffset(c *gin.Context) {
	var (
		internalMemos []model.SelectInternalMemoParameter
		response      helper.Response
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
					employeeID, err := strconv.ParseInt(c.Param("employeeID"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param employeeID was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						internalMemos, err = b.internalMemoService.FindInternalMemosOffset(int(limit), int(offset), order, dir, int(employeeID))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", internalMemos)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *internalMemoController) SearchInternalMemo(c *gin.Context) {
	var (
		internalMemos []model.SelectInternalMemoParameter
		response      helper.Response
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
						employeeID, err := strconv.ParseInt(c.Param("employeeID"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param employeeID was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							internalMemos, err = b.internalMemoService.SearchInternalMemo(int(limit), int(offset), order, dir, search, int(employeeID))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", internalMemos)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *internalMemoController) CountSearchInternalMemo(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		employeeID, err := strconv.ParseInt(c.Param("employeeID"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param employeeID was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.internalMemoService.CountSearchInternalMemo(search, int(employeeID))
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

func (b *internalMemoController) CountInternalMemoByDept(c *gin.Context) {
	var (
		response helper.Response
	)
	deptName := c.Param("deptName")
	if deptName == "" {
		response = helper.BuildErrorResponse("No param deptName was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.internalMemoService.CountInternalMemoByDept(deptName)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoController) FindInternalMemoById(c *gin.Context) {
	var (
		internalMemo model.SelectInternalMemoParameter
		response     helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemo, err = b.internalMemoService.FindInternalMemoById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemo)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoController) FindExcInternalMemo(c *gin.Context) {
	var (
		internalMemos []model.SelectInternalMemoParameter
		response      helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		internalMemos, err = b.internalMemoService.FindExcInternalMemo(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemos)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoController) InsertInternalMemo(c *gin.Context) {
	var (
		internalMemo                model.InternalMemo
		response                    helper.Response
		CreateInternalMemoParameter model.CreateInternalMemoParameter
	)
	err := c.ShouldBindJSON(&CreateInternalMemoParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		internalMemo, err = b.internalMemoService.InsertInternalMemo(CreateInternalMemoParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register internalMemo", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", internalMemo)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *internalMemoController) UpdateInternalMemo(c *gin.Context) {
	var (
		newData  model.InternalMemo
		oldData  model.SelectInternalMemoParameter
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
			oldData, err = b.internalMemoService.FindInternalMemoById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectInternalMemoParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemo, err := b.internalMemoService.UpdateInternalMemo(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update internalMemo", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemo)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *internalMemoController) DeleteInternalMemo(c *gin.Context) {
	var (
		newData  model.InternalMemo
		oldData  model.SelectInternalMemoParameter
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
			oldData, err = b.internalMemoService.FindInternalMemoById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectInternalMemoParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemo, err := b.internalMemoService.DeleteInternalMemo(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete internalMemo", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemo)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *internalMemoController) UpdateInternalMemoApprove(c *gin.Context) {
	var (
		newData  model.InternalMemo
		oldData  model.SelectInternalMemoParameter
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
			oldData, err = b.internalMemoService.FindInternalMemoById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectInternalMemoParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemo, err := b.internalMemoService.UpdateInternalMemoApprove(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update internalMemo", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemo)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *internalMemoController) UploadIMDocument(c *gin.Context) {
	var (
		newData  model.InternalMemo
		oldData  model.SelectInternalMemoParameter
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
			oldData, err = b.internalMemoService.FindInternalMemoById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectInternalMemoParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				internalMemo, err := b.internalMemoService.UploadIMDocument(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update internalMemo", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", internalMemo)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
