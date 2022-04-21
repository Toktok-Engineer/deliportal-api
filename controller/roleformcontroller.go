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
	FindRoleForms(c *gin.Context)
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
