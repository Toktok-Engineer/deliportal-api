package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserCompanyRestrictionController interface {
	CountUserCompanyRestrictionAll(c *gin.Context)
	FindUserCompanyRestrictions(c *gin.Context)
	FindUserCompanyRestrictionsOffset(c *gin.Context)
	SearchUserCompanyRestriction(c *gin.Context)
	CountSearchUserCompanyRestriction(c *gin.Context)
	FindUserCompanyRestrictionById(c *gin.Context)
	FindUserCompanyRestrictionByUserId(c *gin.Context)
	FindExcUserCompanyRestriction(c *gin.Context)
	InsertUserCompanyRestriction(c *gin.Context)
	UpdateUserCompanyRestriction(c *gin.Context)
	DeleteUserCompanyRestriction(c *gin.Context)
}

type usercompanyrestrictionController struct {
	usercompanyrestrictionService service.UserCompanyRestrictionService
	jwtService                    service.JWTService
}

func NewUserCompanyRestrictionController(usercompanyrestrictionServ service.UserCompanyRestrictionService, jwtServ service.JWTService) UserCompanyRestrictionController {
	return &usercompanyrestrictionController{
		usercompanyrestrictionService: usercompanyrestrictionServ,
		jwtService:                    jwtServ,
	}
}

func (b *usercompanyrestrictionController) CountUserCompanyRestrictionAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	usernameID, err := strconv.ParseInt(c.Param("usernameID"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param usernameID was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err = b.usercompanyrestrictionService.CountUserCompanyRestrictionAll(int(usernameID))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *usercompanyrestrictionController) FindUserCompanyRestrictions(c *gin.Context) {
	var (
		usercompanyrestrictions []model.SelectUserCompanyRestrictionParameter
		response                helper.Response
	)
	usercompanyrestrictions, err := b.usercompanyrestrictionService.FindUserCompanyRestrictions()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", usercompanyrestrictions)
		c.JSON(http.StatusOK, response)
	}
}

func (b *usercompanyrestrictionController) FindUserCompanyRestrictionsOffset(c *gin.Context) {
	var (
		usercompanyrestrictions []model.SelectUserCompanyRestrictionParameter
		response                helper.Response
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
					usernameID, err := strconv.ParseInt(c.Param("usernameID"), 0, 0)
					if err != nil {
						response = helper.BuildErrorResponse("No param usernameID was found", err.Error(), helper.EmptyObj{})
						c.AbortWithStatusJSON(http.StatusBadRequest, response)
					} else {
						usercompanyrestrictions, err = b.usercompanyrestrictionService.FindUserCompanyRestrictionsOffset(int(limit), int(offset), order, dir, int(usernameID))
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", usercompanyrestrictions)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *usercompanyrestrictionController) SearchUserCompanyRestriction(c *gin.Context) {
	var (
		usercompanyrestrictions []model.SelectUserCompanyRestrictionParameter
		response                helper.Response
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
						usernameID, err := strconv.ParseInt(c.Param("usernameID"), 0, 0)
						if err != nil {
							response = helper.BuildErrorResponse("No param usernameID was found", err.Error(), helper.EmptyObj{})
							c.AbortWithStatusJSON(http.StatusBadRequest, response)
						} else {
							usercompanyrestrictions, err = b.usercompanyrestrictionService.SearchUserCompanyRestriction(int(limit), int(offset), order, dir, search, int(usernameID))
							if err != nil {
								response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
								c.JSON(http.StatusNotFound, response)
							} else {
								response = helper.BuildResponse(true, "OK", usercompanyrestrictions)
								c.JSON(http.StatusOK, response)
							}
						}
					}
				}
			}
		}
	}
}

func (b *usercompanyrestrictionController) CountSearchUserCompanyRestriction(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		usernameID, err := strconv.ParseInt(c.Param("usernameID"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param usernameID was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			count, err := b.usercompanyrestrictionService.CountSearchUserCompanyRestriction(search, int(usernameID))
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

func (b *usercompanyrestrictionController) FindUserCompanyRestrictionById(c *gin.Context) {
	var (
		usercompanyrestriction model.SelectUserCompanyRestrictionParameter
		response               helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		usercompanyrestriction, err = b.usercompanyrestrictionService.FindUserCompanyRestrictionById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", usercompanyrestriction)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *usercompanyrestrictionController) FindUserCompanyRestrictionByUserId(c *gin.Context) {
	var (
		usercompanyrestrictions []model.SelectUserCompanyRestrictionParameter
		response                helper.Response
	)
	uid, err := strconv.ParseUint(c.Param("uid"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param uid was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		usercompanyrestrictions, err = b.usercompanyrestrictionService.FindUserCompanyRestrictionByUserId(uint(uid))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", usercompanyrestrictions)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *usercompanyrestrictionController) FindExcUserCompanyRestriction(c *gin.Context) {
	var (
		usercompanyrestrictions []model.SelectUserCompanyRestrictionParameter
		response                helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		usercompanyrestrictions, err = b.usercompanyrestrictionService.FindExcUserCompanyRestriction(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", usercompanyrestrictions)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *usercompanyrestrictionController) InsertUserCompanyRestriction(c *gin.Context) {
	var (
		usercompanyrestriction                model.UserCompanyRestriction
		response                              helper.Response
		CreateUserCompanyRestrictionParameter model.CreateUserCompanyRestrictionParameter
	)
	err := c.ShouldBindJSON(&CreateUserCompanyRestrictionParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		usercompanyrestriction, err = b.usercompanyrestrictionService.InsertUserCompanyRestriction(CreateUserCompanyRestrictionParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register usercompanyrestriction", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", usercompanyrestriction)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *usercompanyrestrictionController) UpdateUserCompanyRestriction(c *gin.Context) {
	var (
		newData  model.UserCompanyRestriction
		oldData  model.SelectUserCompanyRestrictionParameter
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
			oldData, err = b.usercompanyrestrictionService.FindUserCompanyRestrictionById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectUserCompanyRestrictionParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				usercompanyrestriction, err := b.usercompanyrestrictionService.UpdateUserCompanyRestriction(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update usercompanyrestriction", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", usercompanyrestriction)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *usercompanyrestrictionController) DeleteUserCompanyRestriction(c *gin.Context) {
	var (
		newData  model.UserCompanyRestriction
		oldData  model.SelectUserCompanyRestrictionParameter
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
			oldData, err = b.usercompanyrestrictionService.FindUserCompanyRestrictionById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectUserCompanyRestrictionParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				usercompanyrestriction, err := b.usercompanyrestrictionService.DeleteUserCompanyRestriction(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete usercompanyrestriction", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", usercompanyrestriction)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
