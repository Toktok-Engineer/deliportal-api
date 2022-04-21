package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SectionController interface {
	FindSections(c *gin.Context)
	FindSectionById(c *gin.Context)
	FindExcSection(c *gin.Context)
	FindSectionByDepId(c *gin.Context)
	InsertSection(c *gin.Context)
	UpdateSection(c *gin.Context)
	DeleteSection(c *gin.Context)
}

type sectionController struct {
	sectionService service.SectionService
	jwtService     service.JWTService
}

func NewSectionController(sectionServ service.SectionService, jwtServ service.JWTService) SectionController {
	return &sectionController{
		sectionService: sectionServ,
		jwtService:     jwtServ,
	}
}

func (b *sectionController) FindSections(c *gin.Context) {
	var (
		sections []model.SelectSectionParameter
		response helper.Response
	)
	sections, err := b.sectionService.FindSections()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", sections)
		c.JSON(http.StatusOK, response)
	}
}

func (b *sectionController) FindSectionById(c *gin.Context) {
	var (
		section  model.SelectSectionParameter
		response helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		section, err = b.sectionService.FindSectionById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", section)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *sectionController) FindExcSection(c *gin.Context) {
	var (
		sections []model.SelectSectionParameter
		response helper.Response
	)
	depId, err := strconv.ParseUint(c.Param("depId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param depId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		id, err := strconv.ParseUint(c.Param("id"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			sections, err = b.sectionService.FindExcSection(uint(depId), uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", sections)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *sectionController) FindSectionByDepId(c *gin.Context) {
	var (
		sections []model.SelectSectionParameter
		response helper.Response
	)
	depId, err := strconv.ParseUint(c.Param("depId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param depId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		sections, err = b.sectionService.FindSectionByDepId(uint(depId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", sections)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *sectionController) InsertSection(c *gin.Context) {
	var (
		section                model.Section
		response               helper.Response
		CreateSectionParameter model.CreateSectionParameter
	)
	err := c.ShouldBindJSON(&CreateSectionParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		section, err = b.sectionService.InsertSection(CreateSectionParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register section", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", section)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *sectionController) UpdateSection(c *gin.Context) {
	var (
		newData  model.Section
		oldData  model.SelectSectionParameter
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
			oldData, err = b.sectionService.FindSectionById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectSectionParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				section, err := b.sectionService.UpdateSection(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update section", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", section)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *sectionController) DeleteSection(c *gin.Context) {
	var (
		newData  model.Section
		oldData  model.SelectSectionParameter
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
			oldData, err = b.sectionService.FindSectionById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectSectionParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				section, err := b.sectionService.DeleteSection(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete section", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", section)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
