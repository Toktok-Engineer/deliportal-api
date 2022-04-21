package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FormTypeController interface {
	FindFormTypes(c *gin.Context)
	FindFormTypeById(c *gin.Context)
	FindExcFormType(c *gin.Context)
	InsertFormType(c *gin.Context)
	UpdateFormType(c *gin.Context)
	DeleteFormType(c *gin.Context)
}

type formTypeController struct {
	formTypeService service.FormTypeService
	jwtService      service.JWTService
}

func NewFormTypeController(formTypeServ service.FormTypeService, jwtServ service.JWTService) FormTypeController {
	return &formTypeController{
		formTypeService: formTypeServ,
		jwtService:      jwtServ,
	}
}

func (b *formTypeController) FindFormTypes(c *gin.Context) {
	var (
		formTypes []model.FormType
		response  helper.Response
	)
	formTypes, err := b.formTypeService.FindFormTypes()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", formTypes)
		c.JSON(http.StatusOK, response)
	}
}

func (b *formTypeController) FindFormTypeById(c *gin.Context) {
	var (
		formType model.FormType
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		formType, err = b.formTypeService.FindFormTypeById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", formType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formTypeController) FindExcFormType(c *gin.Context) {
	var (
		formTypes []model.FormType
		response  helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		formTypes, err = b.formTypeService.FindExcFormType(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", formTypes)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formTypeController) InsertFormType(c *gin.Context) {
	var (
		formType                model.FormType
		response                helper.Response
		CreateFormTypeParameter model.CreateFormTypeParameter
	)
	err := c.ShouldBindJSON(&CreateFormTypeParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		formType, err = b.formTypeService.InsertFormType(CreateFormTypeParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register formType", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", formType)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *formTypeController) UpdateFormType(c *gin.Context) {
	var (
		newData  model.FormType
		oldData  model.FormType
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
			oldData, err = b.formTypeService.FindFormTypeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.FormType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				formType, err := b.formTypeService.UpdateFormType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update formType", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", formType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *formTypeController) DeleteFormType(c *gin.Context) {
	var (
		newData  model.FormType
		oldData  model.FormType
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
			oldData, err = b.formTypeService.FindFormTypeById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.FormType{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				formType, err := b.formTypeService.DeleteFormType(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete formType", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", formType)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
