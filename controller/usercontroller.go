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

type UserController interface {
	CountUserAll(c *gin.Context)
	FindUsersAll(c *gin.Context)
	FindUsers(c *gin.Context)
	FindUsersOffset(c *gin.Context)
	SearchUser(c *gin.Context)
	CountSearchUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByUName(c *gin.Context)
	FindExcUser(c *gin.Context)
	InsertUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userServ service.UserService, jwtServ service.JWTService) UserController {
	return &userController{
		userService: userServ,
		jwtService:  jwtServ,
	}
}

func (b *userController) CountUserAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	count, err := b.userService.CountUserAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *userController) FindUsersAll(c *gin.Context) {
	var (
		user     model.SelectUserParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		user, err = b.userService.FindUsersAll(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", user)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userController) FindUsers(c *gin.Context) {
	var (
		users    []model.SelectUserParameter
		response helper.Response
	)
	users, err := b.userService.FindUsers()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", users)
		c.JSON(http.StatusOK, response)
	}
}

func (b *userController) FindUsersOffset(c *gin.Context) {
	var (
		users    []model.SelectUserParameter
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
					users, err = b.userService.FindUsersOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", users)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *userController) SearchUser(c *gin.Context) {
	var (
		users    []model.SelectUserParameter
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
						users, err = b.userService.SearchUser(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", users)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *userController) CountSearchUser(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.userService.CountSearchUser(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userController) FindUserById(c *gin.Context) {
	var (
		user     model.SelectUserParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		user, err = b.userService.FindUserById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", user)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userController) FindUserByUName(c *gin.Context) {
	var (
		response helper.Response
	)
	UName := c.Param("uName")
	if UName == "" {
		response = helper.BuildErrorResponse("No param uName was found", "No data with given UName", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		user, err := b.userService.FindUserByUName(UName)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", user)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userController) FindExcUser(c *gin.Context) {
	var (
		users    []model.SelectUserParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		users, err = b.userService.FindExcUser(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", users)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userController) InsertUser(c *gin.Context) {
	var (
		user                model.User
		response            helper.Response
		CreateUserParameter model.CreateUserParameter
	)
	err := c.ShouldBindJSON(&CreateUserParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		user, err = b.userService.InsertUser(CreateUserParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register user", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", user)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *userController) UpdateUser(c *gin.Context) {
	var (
		newData  model.User
		oldData  model.SelectUserParameter
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
			oldData, err = b.userService.FindUserById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectEmployeeParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				user, err := b.userService.UpdateUser(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update user", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", user)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
func (b *userController) DeleteUser(c *gin.Context) {
	var (
		newData  model.User
		oldData  model.SelectUserParameter
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
			oldData, err = b.userService.FindUserById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (cmp.Equal(oldData, model.SelectEmployeeParameter{})) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				user, err := b.userService.DeleteUser(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete user", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", user)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
