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

type UserRoleController interface {
	CountUserRoleAll(c *gin.Context)
	FindUserRoles(c *gin.Context)
	FindUserRolesOffset(c *gin.Context)
	SearchUserRole(c *gin.Context)
	CountSearchUserRole(c *gin.Context)
	FindUserRoleById(c *gin.Context)
	FindUserRoleByUserId(c *gin.Context)
	FindExcUserRole(c *gin.Context)
	FindExcUserRoleOnly(c *gin.Context)
	InsertUserRole(c *gin.Context)
	UpdateUserRole(c *gin.Context)
	DeleteUserRole(c *gin.Context)
}

type userRoleController struct {
	userRoleService service.UserRoleService
	jwtService      service.JWTService
}

func NewUserRoleController(userRoleServ service.UserRoleService, jwtServ service.JWTService) UserRoleController {
	return &userRoleController{
		userRoleService: userRoleServ,
		jwtService:      jwtServ,
	}
}

func (b *userRoleController) CountUserRoleAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	count, err := b.userRoleService.CountUserRoleAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *userRoleController) FindUserRoles(c *gin.Context) {
	var (
		userRoles []model.SelectUserRoleParameter
		response  helper.Response
	)
	userRoles, err := b.userRoleService.FindUserRoles()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", userRoles)
		c.JSON(http.StatusOK, response)
	}
}

func (b *userRoleController) FindUserRolesOffset(c *gin.Context) {
	var (
		userRoles []model.SelectUserRoleParameter
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
					userRoles, err = b.userRoleService.FindUserRolesOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", userRoles)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *userRoleController) SearchUserRole(c *gin.Context) {
	var (
		userRoles []model.SelectUserRoleParameter
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
						userRoles, err = b.userRoleService.SearchUserRole(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", userRoles)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *userRoleController) CountSearchUserRole(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.userRoleService.CountSearchUserRole(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userRoleController) FindUserRoleById(c *gin.Context) {
	var (
		userRole model.SelectUserRoleParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		userRole, err = b.userRoleService.FindUserRoleById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", userRole)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userRoleController) FindUserRoleByUserId(c *gin.Context) {
	var (
		response helper.Response
	)
	uid, err := strconv.ParseUint(c.Param("uid"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param uid was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		userRole, err := b.userRoleService.FindUserRoleByUserId(uint(uid))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", userRole)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userRoleController) FindExcUserRole(c *gin.Context) {
	var (
		userRoles []model.SelectUserRoleParameter
		response  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		uid, err := strconv.ParseUint(c.Param("uid"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param uid was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			userRoles, err = b.userRoleService.FindExcUserRole(uint(id), uint(uid))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", userRoles)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *userRoleController) FindExcUserRoleOnly(c *gin.Context) {
	var (
		userRoles []model.SelectUserRoleParameter
		response  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		userRoles, err = b.userRoleService.FindExcUserRoleOnly(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", userRoles)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userRoleController) InsertUserRole(c *gin.Context) {
	var (
		userRole                model.UserRole
		response                helper.Response
		CreateUserRoleParameter model.CreateUserRoleParameter
	)
	err := c.ShouldBindJSON(&CreateUserRoleParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		userRole, err = b.userRoleService.InsertUserRole(CreateUserRoleParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register userRole", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", userRole)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userRoleController) UpdateUserRole(c *gin.Context) {
	var (
		newData  model.UserRole
		oldData  model.SelectUserRoleParameter
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
			oldData, err = b.userRoleService.FindUserRoleById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectEmployeeParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				userRole, err := b.userRoleService.UpdateUserRole(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update userRole", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", userRole)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
func (b *userRoleController) DeleteUserRole(c *gin.Context) {
	var (
		newData  model.UserRole
		oldData  model.SelectUserRoleParameter
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
			oldData, err = b.userRoleService.FindUserRoleById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectEmployeeParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				userRole, err := b.userRoleService.DeleteUserRole(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete userRole", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", userRole)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
