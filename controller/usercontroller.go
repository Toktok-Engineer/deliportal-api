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
	FindUsers(c *gin.Context)
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
	UName := c.Param("UName")
	if UName == "" {
		response = helper.BuildErrorResponse("No param UName was found", "No data with given UName", helper.EmptyObj{})
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
