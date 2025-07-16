package controller

import (
	"deliportal-api/helper"
	"deliportal-api/model"
	"deliportal-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubSectionController interface {
	CountSubSectionAll(c *gin.Context)
	FindSubSections(c *gin.Context)
	FindSubSectionsOffset(c *gin.Context)
	SearchSubSection(c *gin.Context)
	CountSearchSubSection(c *gin.Context)
	FindSubSectionById(c *gin.Context)
	FindExcSubSection(c *gin.Context)
	FindSubSectionBySecId(c *gin.Context)
	FindSubSectionByDepId(c *gin.Context)
	FindSubSectionByDivisionID(c *gin.Context)
	CountSubSectionName(c *gin.Context)
	InsertSubSection(c *gin.Context)
	UpdateSubSection(c *gin.Context)
	DeleteSubSection(c *gin.Context)
}

type subSubSectionController struct {
	subSubSectionService service.SubSectionService
	jwtService           service.JWTService
}

func NewSubSectionController(subSubSectionServ service.SubSectionService, jwtServ service.JWTService) SubSectionController {
	return &subSubSectionController{
		subSubSectionService: subSubSectionServ,
		jwtService:           jwtServ,
	}
}

func (b *subSubSectionController) CountSubSectionAll(c *gin.Context) {
	var (
		count    int64
		response helper.Response
	)

	count, err := b.subSubSectionService.CountSubSectionAll()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", count)
		c.JSON(http.StatusOK, response)
	}
}

func (b *subSubSectionController) FindSubSections(c *gin.Context) {
	var (
		subSubSections []model.SelectSubSectionParameter
		response       helper.Response
	)
	subSubSections, err := b.subSubSectionService.FindSubSections()
	if err != nil {
		response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusNotFound, response)
	} else {
		response = helper.BuildResponse(true, "OK", subSubSections)
		c.JSON(http.StatusOK, response)
	}
}

func (b *subSubSectionController) FindSubSectionsOffset(c *gin.Context) {
	var (
		subSubSections []model.SelectSubSectionParameter
		response       helper.Response
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
					subSubSections, err = b.subSubSectionService.FindSubSectionsOffset(int(limit), int(offset), order, dir)
					if err != nil {
						response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
						c.JSON(http.StatusNotFound, response)
					} else {
						response = helper.BuildResponse(true, "OK", subSubSections)
						c.JSON(http.StatusOK, response)
					}
				}
			}
		}
	}
}

func (b *subSubSectionController) SearchSubSection(c *gin.Context) {
	var (
		subSubSections []model.SelectSubSectionParameter
		response       helper.Response
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
						subSubSections, err = b.subSubSectionService.SearchSubSection(int(limit), int(offset), order, dir, search)
						if err != nil {
							response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
							c.JSON(http.StatusNotFound, response)
						} else {
							response = helper.BuildResponse(true, "OK", subSubSections)
							c.JSON(http.StatusOK, response)
						}
					}
				}
			}
		}
	}
}

func (b *subSubSectionController) CountSearchSubSection(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.subSubSectionService.CountSearchSubSection(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *subSubSectionController) FindSubSectionById(c *gin.Context) {
	var (
		subSubSection model.SelectSubSectionParameter
		response      helper.Response
	)
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		subSubSection, err = b.subSubSectionService.FindSubSectionById(uint(id))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", subSubSection)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *subSubSectionController) FindExcSubSection(c *gin.Context) {
	var (
		subSubSections []model.SelectSubSectionParameter
		response       helper.Response
	)
	sectionId, err := strconv.ParseUint(c.Param("sectionId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param sectionId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		id, err := strconv.ParseUint(c.Param("id"), 0, 0)
		if err != nil {
			response = helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			subSubSections, err = b.subSubSectionService.FindExcSubSection(uint(sectionId), uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				response = helper.BuildResponse(true, "OK", subSubSections)
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

func (b *subSubSectionController) FindSubSectionBySecId(c *gin.Context) {
	var (
		subSubSections []model.SelectSubSectionParameter
		response       helper.Response
	)
	sectionId, err := strconv.ParseUint(c.Param("sectionId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param sectionId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		subSubSections, err = b.subSubSectionService.FindSubSectionBySecId(uint(sectionId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", subSubSections)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *subSubSectionController) FindSubSectionByDepId(c *gin.Context) {
	var (
		subSubSections []model.SelectSubSectionParameter
		response       helper.Response
	)
	depId, err := strconv.ParseUint(c.Param("depId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param depId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		subSubSections, err = b.subSubSectionService.FindSubSectionByDepId(uint(depId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", subSubSections)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *subSubSectionController) FindSubSectionByDivisionID(c *gin.Context) {
	var (
		subSubSections []model.SelectSubSectionParameter
		response       helper.Response
	)
	divId, err := strconv.ParseUint(c.Param("divId"), 0, 0)
	if err != nil {
		response = helper.BuildErrorResponse("No param divId was found", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		subSubSections, err = b.subSubSectionService.FindSubSectionByDivisionID(uint(divId))
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", subSubSections)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *subSubSectionController) CountSubSectionName(c *gin.Context) {
	var (
		response helper.Response
	)
	search := c.Param("search")
	if search == "" {
		response = helper.BuildErrorResponse("No param search was found", "No data with given search", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		count, err := b.subSubSectionService.CountSubSectionName(search)
		if err != nil {
			response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusNotFound, response)
		} else {
			response = helper.BuildResponse(true, "OK", count)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *subSubSectionController) InsertSubSection(c *gin.Context) {
	var (
		subSubSection             model.SubSection
		response                  helper.Response
		CreateSubSectionParameter model.CreateSubSectionParameter
	)
	err := c.ShouldBindJSON(&CreateSubSectionParameter)
	if err != nil {
		response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
	} else {
		subSubSection, err = b.subSubSectionService.InsertSubSection(CreateSubSectionParameter)
		if err != nil {
			response = helper.BuildErrorResponse("Failed to register subSubSection", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, response)
		} else {
			response = helper.BuildResponse(true, "OK", subSubSection)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (b *subSubSectionController) UpdateSubSection(c *gin.Context) {
	var (
		newData  model.SubSection
		oldData  model.SelectSubSectionParameter
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
			oldData, err = b.subSubSectionService.FindSubSectionById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectSubSectionParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				subSubSection, err := b.subSubSectionService.UpdateSubSection(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to update subSubSection", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", subSubSection)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}

func (b *subSubSectionController) DeleteSubSection(c *gin.Context) {
	var (
		newData  model.SubSection
		oldData  model.SelectSubSectionParameter
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
			oldData, err = b.subSubSectionService.FindSubSectionById(uint(id))
			if err != nil {
				response = helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else if (oldData == model.SelectSubSectionParameter{}) {
				response = helper.BuildErrorResponse("Data not found", err.Error(), helper.EmptyObj{})
				c.JSON(http.StatusNotFound, response)
			} else {
				subSubSection, err := b.subSubSectionService.DeleteSubSection(newData, uint(id))
				if err != nil {
					response = helper.BuildErrorResponse("Failed to delete subSubSection", err.Error(), helper.EmptyObj{})
					c.AbortWithStatusJSON(http.StatusBadRequest, response)
				} else {
					response = helper.BuildResponse(true, "OK", subSubSection)
					c.JSON(http.StatusOK, response)
				}
			}
		}
	}
}
