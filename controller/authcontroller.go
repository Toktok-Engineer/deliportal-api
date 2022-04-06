package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	RenewToken(c *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (a *authController) Login(c *gin.Context) {
	var loginParameter model.LoginParameter
	var loginResponse model.LoginResponse
	err := c.ShouldBind(&loginParameter)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := a.authService.VerifyCredential(loginParameter.Username, loginParameter.Password)
	if user, ok := authResult.(model.User); ok {
		if user.ID != 0 {
			now := time.Now()
			iat := now.Unix()
			exp := now.Add(time.Hour * 24).Unix()
			expRefreshToken := now.AddDate(0, 0, 7).Unix()

			generatedToken := a.jwtService.GenerateToken(uint64(user.ID), user.Username, iat, exp)
			generatedRefreshToken := a.jwtService.GenerateRefreshToken(uint64(user.ID), user.Username, iat, expRefreshToken)

			loginResponse.User = user
			loginResponse.AccessToken = generatedToken
			loginResponse.RefreshToken = generatedRefreshToken

			response := helper.BuildResponse(true, "OK", loginResponse)
			c.JSON(http.StatusOK, response)
		} else {
			response := helper.BuildResponse(false, "User not found", nil)
			c.JSON(http.StatusOK, response)
		}

		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (a *authController) Register(c *gin.Context) {
	var createUserParameter model.CreateUserParameter
	var loginResponse model.LoginResponse
	var response helper.Response

	errDTO := c.ShouldBind(&createUserParameter)
	if errDTO != nil {
		response = helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !a.authService.IsDuplicateUsername(createUserParameter.Username) {
		response = helper.BuildErrorResponse("Failed to process request", "Duplicate username", helper.EmptyObj{})
		c.JSON(http.StatusConflict, response)
	} else {
		createdUser, err := a.authService.CreateUser(createUserParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register user", err.Error(), nil)
			c.JSON(http.StatusBadRequest, response)
		} else {
			now := time.Now()
			iat := now.Unix()
			exp := now.Add(time.Hour * 24).Unix()
			expRefreshToken := now.AddDate(0, 0, 7).Unix()

			accessToken := a.jwtService.GenerateToken(uint64(createdUser.ID), createdUser.Username, iat, exp)
			refreshToken := a.jwtService.GenerateRefreshToken(uint64(createdUser.ID), createdUser.Username, iat, expRefreshToken)

			loginResponse.User = createdUser
			loginResponse.AccessToken = accessToken
			loginResponse.RefreshToken = refreshToken

			response = helper.BuildResponse(true, "OK", loginResponse)
			c.JSON(http.StatusCreated, response)
		}
	}
}

func (a *authController) RenewToken(c *gin.Context) {
	var renewToken model.RenewToken
	err := c.ShouldBind(&renewToken)

	var renewTokenResult model.RenewTokenResult

	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if a.authService.IsUserRegistered(renewToken.ID, renewToken.Username) {
		response := helper.BuildErrorResponse("Failed to process request", "User not registered", helper.EmptyObj{})
		c.JSON(http.StatusConflict, response)
		return
	} else {

		now := time.Now()
		iat := now.Unix()
		exp := now.Add(time.Hour * 24).Unix()
		expRefreshToken := now.AddDate(0, 0, 7).Unix()

		generatedToken := a.jwtService.GenerateToken(renewToken.ID, renewToken.Username, iat, exp)
		generatedRefreshToken := a.jwtService.GenerateRefreshToken(renewToken.ID, renewToken.Username, iat, expRefreshToken)

		renewTokenResult.AccessToken = generatedToken
		renewTokenResult.RefreshToken = generatedRefreshToken

		response := helper.BuildResponse(true, "OK", renewTokenResult)
		c.JSON(http.StatusOK, response)
		return
	}
}
