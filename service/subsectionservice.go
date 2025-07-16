package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type SubSectionService interface {
	CountSubSectionAll() (count int64, err error)
	FindSubSections() (subSectionOutput []model.SelectSubSectionParameter, err error)
	FindSubSectionsOffset(limit int, offset int, order string, dir string) (subSectionOutput []model.SelectSubSectionParameter, err error)
	SearchSubSection(limit int, offset int, order string, dir string, search string) (subSectionOutput []model.SelectSubSectionParameter, err error)
	CountSearchSubSection(search string) (count int64, err error)
	FindSubSectionById(id uint) (subSectionOutput model.SelectSubSectionParameter, err error)
	FindExcSubSection(sectionId uint, id uint) (subSectionOutput []model.SelectSubSectionParameter, err error)
	FindSubSectionBySecId(sectionId uint) (subSectionOutput []model.SelectSubSectionParameter, err error)
	FindSubSectionByDepId(depId uint) (subSectionOutput []model.SelectSubSectionParameter, err error)
	FindSubSectionByDivisionID(divId uint) (subSectionOutput []model.SelectSubSectionParameter, err error)
	CountSubSectionName(search string) (count int64, err error)
	InsertSubSection(subSection model.CreateSubSectionParameter) (subSectionOutput model.SubSection, err error)
	UpdateSubSection(subSection model.SubSection, id uint) (subSectionOutput model.SubSection, err error)
	DeleteSubSection(subSection model.SubSection, id uint) (subSectionOutput model.SubSection, err error)
}

type subSectionService struct {
	subSectionRepository repository.SubSectionRepository
}

func NewSubSectionService(subSectionRep repository.SubSectionRepository) SubSectionService {
	return &subSectionService{
		subSectionRepository: subSectionRep,
	}
}

func (service *subSectionService) CountSubSectionAll() (count int64, err error) {
	res, err := service.subSectionRepository.CountSubSectionAll()
	return res, err
}

func (service *subSectionService) FindSubSections() (subSectionOutput []model.SelectSubSectionParameter, err error) {
	res, err := service.subSectionRepository.FindSubSections()
	return res, err
}

func (service *subSectionService) FindSubSectionsOffset(limit int, offset int, order string, dir string) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	res, err := service.subSectionRepository.FindSubSectionsOffset(limit, offset, order, dir)
	return res, err
}

func (service *subSectionService) SearchSubSection(limit int, offset int, order string, dir string, search string) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	res, err := service.subSectionRepository.SearchSubSection(limit, offset, order, dir, search)
	return res, err
}

func (service *subSectionService) CountSearchSubSection(search string) (count int64, err error) {
	res, err := service.subSectionRepository.CountSearchSubSection(search)
	return res, err
}

func (service *subSectionService) FindSubSectionById(id uint) (subSectionOutput model.SelectSubSectionParameter, err error) {
	res, err := service.subSectionRepository.FindSubSectionById(id)
	return res, err
}

func (service *subSectionService) FindExcSubSection(sectionId uint, id uint) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	res, err := service.subSectionRepository.FindExcSubSection(sectionId, id)
	return res, err
}

func (service *subSectionService) FindSubSectionBySecId(sectionId uint) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	res, err := service.subSectionRepository.FindSubSectionBySecId(sectionId)
	return res, err
}

func (service *subSectionService) FindSubSectionByDepId(depId uint) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	res, err := service.subSectionRepository.FindSubSectionByDepId(depId)
	return res, err
}

func (service *subSectionService) FindSubSectionByDivisionID(divId uint) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	res, err := service.subSectionRepository.FindSubSectionByDivisionID(divId)
	return res, err
}

func (service *subSectionService) CountSubSectionName(search string) (count int64, err error) {
	res, err := service.subSectionRepository.CountSubSectionName(search)
	return res, err
}

func (service *subSectionService) InsertSubSection(subSection model.CreateSubSectionParameter) (subSectionOutput model.SubSection, err error) {
	newSubSection := model.SubSection{}
	err1 := smapping.FillStruct(&newSubSection, smapping.MapFields(&subSection))
	if err != nil {
		return newSubSection, err1
	}
	res, err := service.subSectionRepository.InsertSubSection(newSubSection)
	return res, err
}

func (service *subSectionService) UpdateSubSection(subSection model.SubSection, id uint) (subSectionOutput model.SubSection, err error) {
	newSubSection := model.SubSection{}
	err1 := smapping.FillStruct(&newSubSection, smapping.MapFields(&subSection))
	if err != nil {
		return newSubSection, err1
	}
	res, err := service.subSectionRepository.UpdateSubSection(newSubSection, id)
	return res, err
}

func (service *subSectionService) DeleteSubSection(subSection model.SubSection, id uint) (subSectionOutput model.SubSection, err error) {
	newSubSection := model.SubSection{}
	err1 := smapping.FillStruct(&newSubSection, smapping.MapFields(&subSection))
	if err != nil {
		return newSubSection, err1
	}
	res, err := service.subSectionRepository.UpdateSubSection(newSubSection, id)
	return res, err
}
