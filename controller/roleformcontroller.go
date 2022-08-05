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

type RoleFormController interface {
	CountRoleFormAll(c *gin.Context)
	FindRoleForms(c *gin.Context)
	FindRoleFormsOffset(c *gin.Context)
	SearchRoleForm(c *gin.Context)
	CountSearchRoleForm(c *gin.Context)
	FindRoleFormById(c *gin.Context)
	FindRoleFormByFormId(c *gin.Context)
	FindExcRoleForm(c *gin.Context)
	FindExcRoleFormOnly(c *gin.Context)
	InsertRoleForm(c *gin.Context)
	UpdateRoleForm(c *gin.Context)
	DeleteRoleForm(c *gin.Context)
}

type roleFormController struct {
	roleFormService service.RoleFormService
	jwtService      service.JWTService
}

func NewRoleFormController(roleFormServ service.RoleFormService, jwtServ service.JWTService) RoleFormController {
	return &roleFormController{
		roleFormService: roleFormServ,
		jwtService:      jwtServ,
	}
}

func (b *roleFormController) CountRoleFormAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	count, err := b.roleFormService.CountRoleFormAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *roleFormController) FindRoleForms(c *gin.Context) {
	var (
		roleForms []model.SelectRoleFormParameter
		response  helper.Response
	)
	roleForms, err := b.roleFormService.FindRoleForms()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", roleForms)
		c.JSON(http.StatusOK, response)
	}
}

func (b *roleFormController) FindRoleFormsOffset(c *gin.Context) {
	var (
		roleForms []model.SelectRoleFormParameter
		response  helper.Response
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
					roleForms, err = b.roleFormService.FindRoleFormsOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", roleForms)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *roleFormController) SearchRoleForm(c *gin.Context) {
	var (
		roleForms []model.SelectRoleFormParameter
		response  helper.Response
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
						roleForms, err = b.roleFormService.SearchRoleForm(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", roleForms)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *roleFormController) CountSearchRoleForm(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.roleFormService.CountSearchRoleForm(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *roleFormController) FindRoleFormById(c *gin.Context) {
	var (
		roleForm model.SelectRoleFormParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		roleForm, err = b.roleFormService.FindRoleFormById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", roleForm)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *roleFormController) FindRoleFormByFormId(c *gin.Context) {
	var (
		response helper.Response
	)
	fid, err := strconv.ParseUint(c.Param("fid"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param form id was found", "No data with given UName", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		rid, err := strconv.ParseUint(c.Param("rid"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param role id was found", "No data with given UName", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			roleForm, err := b.roleFormService.FindRoleFormByFormId(uint(fid), uint(rid))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", roleForm)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *roleFormController) FindExcRoleForm(c *gin.Context) {
	var (
		roleForms []model.SelectRoleFormParameter
		response  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		rid, err := strconv.ParseUint(c.Param("rid"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param role id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			roleForms, err = b.roleFormService.FindExcRoleForm(uint(id), uint(rid))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", roleForms)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *roleFormController) FindExcRoleFormOnly(c *gin.Context) {
	var (
		roleForms []model.SelectRoleFormParameter
		response  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		roleForms, err = b.roleFormService.FindExcRoleFormOnly(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", roleForms)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *roleFormController) InsertRoleForm(c *gin.Context) {
	var (
		roleForm                model.RoleForm
		response                helper.Response
		CreateRoleFormParameter model.CreateRoleFormParameter
	)
	err := c.ShouldBindJSON(&CreateRoleFormParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		roleForm, err = b.roleFormService.InsertRoleForm(CreateRoleFormParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register roleForm", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", roleForm)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *roleFormController) UpdateRoleForm(c *gin.Context) {
	var (
		newData  model.RoleForm
		oldData  model.SelectRoleFormParameter
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
			oldData, err = b.roleFormService.FindRoleFormById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectRoleFormParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				roleForm, err := b.roleFormService.UpdateRoleForm(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update roleForm", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", roleForm)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
func (b *roleFormController) DeleteRoleForm(c *gin.Context) {
	var (
		newData  model.RoleForm
		oldData  model.SelectRoleFormParameter
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
			oldData, err = b.roleFormService.FindRoleFormById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectRoleFormParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				roleForm, err := b.roleFormService.DeleteRoleForm(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete roleForm", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", roleForm)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
