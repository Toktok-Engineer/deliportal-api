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

type FormController interface {
	CountFormAll(c *gin.Context)
	FindForms(c *gin.Context)
	FindFormsOffset(c *gin.Context)
	SearchForm(c *gin.Context)
	CountSearchForm(c *gin.Context)
	FindFormJoinRole(c *gin.Context)
	FindFormByRole(c *gin.Context)
	FindFormByType(c *gin.Context)
	FindExcFormByType(c *gin.Context)
	FindFormById(c *gin.Context)
	FindFormByFormTypeId(c *gin.Context)
	FindExcForm(c *gin.Context)
	FindFormHead(c *gin.Context)
	FindFormHeadDetail(c *gin.Context)
	FindExcFormHead(c *gin.Context)
	FindExcFormOnly(c *gin.Context)
	InsertForm(c *gin.Context)
	UpdateForm(c *gin.Context)
	DeleteForm(c *gin.Context)
}

type formController struct {
	formService service.FormService
	jwtService  service.JWTService
}

func NewFormController(formServ service.FormService, jwtServ service.JWTService) FormController {
	return &formController{
		formService: formServ,
		jwtService:  jwtServ,
	}
}

func (b *formController) CountFormAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	count, err := b.formService.CountFormAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *formController) FindForms(c *gin.Context) {
	var (
		forms    []model.SelectFormParameter
		response helper.Response
	)
	forms, err := b.formService.FindForms()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", forms)
		c.JSON(http.StatusOK, response)
	}
}

func (b *formController) FindFormsOffset(c *gin.Context) {
	var (
		forms    []model.SelectFormParameter
		response helper.Response
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
					forms, err = b.formService.FindFormsOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", forms)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *formController) SearchForm(c *gin.Context) {
	var (
		forms    []model.SelectFormParameter
		response helper.Response
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
						forms, err = b.formService.SearchForm(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", forms)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *formController) CountSearchForm(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.formService.CountSearchForm(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formController) FindFormJoinRole(c *gin.Context) {
	var (
		forms    []model.SelectFormCRUDParameter
		response helper.Response
	)
	uId, err := strconv.ParseUint(c.Param("uId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param uId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		fpId, err := strconv.ParseUint(c.Param("fpId"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param fpId was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			forms, err = b.formService.FindFormJoinRole(uint(uId), uint(fpId))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", forms)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *formController) FindFormByRole(c *gin.Context) {
	var (
		forms    []model.SelectFormCRUDParameter
		response helper.Response
	)
	uId, err := strconv.ParseUint(c.Param("uId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param uId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		forms, err = b.formService.FindFormByRole(uint(uId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", forms)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formController) FindFormByType(c *gin.Context) {
	var (
		forms    []model.SelectFormParameter
		response helper.Response
	)
	tyId, err := strconv.ParseUint(c.Param("tyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param tyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		forms, err = b.formService.FindFormByType(uint(tyId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", forms)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formController) FindExcFormByType(c *gin.Context) {
	var (
		forms    []model.SelectFormParameter
		response helper.Response
	)
	tyId, err := strconv.ParseUint(c.Param("tyId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param tyId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		id, err := strconv.ParseUint(c.Param("id"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			forms, err = b.formService.FindExcFormByType(uint(tyId), uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", forms)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *formController) FindFormById(c *gin.Context) {
	var (
		form     model.SelectFormParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		form, err = b.formService.FindFormById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", form)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formController) FindFormByFormTypeId(c *gin.Context) {
	var (
		forms    []model.SelectFormParameter
		response helper.Response
	)
	ftId, err := strconv.ParseUint(c.Param("ftId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param ftId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		forms, err = b.formService.FindFormByFormTypeId(uint(ftId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", forms)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formController) FindExcForm(c *gin.Context) {
	var (
		forms    []model.SelectFormParameter
		response helper.Response
	)
	ftId, err := strconv.ParseUint(c.Param("ftId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param ftId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		id, err := strconv.ParseUint(c.Param("id"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			forms, err = b.formService.FindExcForm(uint(ftId), uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", forms)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *formController) FindFormHead(c *gin.Context) {
	var (
		forms    []model.SelectFormParameter
		response helper.Response
	)
	ftId, err := strconv.ParseUint(c.Param("ftId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param ftId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		forms, err = b.formService.FindFormHead(uint(ftId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", forms)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formController) FindFormHeadDetail(c *gin.Context) {
	var (
		form     model.SelectFormParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		form, err = b.formService.FindFormHeadDetail(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", form)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formController) FindExcFormHead(c *gin.Context) {
	var (
		forms    []model.SelectFormParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		forms, err = b.formService.FindExcFormHead(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", forms)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formController) FindExcFormOnly(c *gin.Context) {
	var (
		forms    []model.SelectFormParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		forms, err = b.formService.FindExcFormOnly(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", forms)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formController) InsertForm(c *gin.Context) {
	var (
		form                model.Form
		response            helper.Response
		CreateFormParameter model.CreateFormParameter
	)
	err := c.ShouldBindJSON(&CreateFormParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		form, err = b.formService.InsertForm(CreateFormParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register form", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", form)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formController) UpdateForm(c *gin.Context) {
	var (
		newData  model.Form
		oldData  model.SelectFormParameter
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
			oldData, err = b.formService.FindFormById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectEmployeeParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				form, err := b.formService.UpdateForm(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update form", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", form)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
func (b *formController) DeleteForm(c *gin.Context) {
	var (
		newData  model.Form
		oldData  model.SelectFormParameter
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
			oldData, err = b.formService.FindFormById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectEmployeeParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				form, err := b.formService.DeleteForm(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete form", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", form)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
