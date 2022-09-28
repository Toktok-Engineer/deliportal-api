package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type SectionService interface {
	CountSectionAll() (count int64, err error)
	FindSections() (sectionOutput []model.SelectSectionParameter, err error)
	FindSectionsOffset(limit int, offset int, order string, dir string) (sectionOutput []model.SelectSectionParameter, err error)
	SearchSection(limit int, offset int, order string, dir string, search string) (sectionOutput []model.SelectSectionParameter, err error)
	CountSearchSection(search string) (count int64, err error)
	FindSectionById(id uint) (sectionOutput model.SelectSectionParameter, err error)
	FindExcSection(depId uint, id uint) (sectionOutput []model.SelectSectionParameter, err error)
	FindSectionByDepId(depId uint) (sectionOutput []model.SelectSectionParameter, err error)
	FindSectionByDivisionID(divId uint) (sectionOutput []model.SelectSectionParameter, err error)
	CountSectionName(search string) (count int64, err error)
	InsertSection(section model.CreateSectionParameter) (sectionOutput model.Section, err error)
	UpdateSection(section model.Section, id uint) (sectionOutput model.Section, err error)
	DeleteSection(section model.Section, id uint) (sectionOutput model.Section, err error)
}

type sectionService struct {
	sectionRepository repository.SectionRepository
}

func NewSectionService(sectionRep repository.SectionRepository) SectionService {
	return &sectionService{
		sectionRepository: sectionRep,
	}
}

func (service *sectionService) CountSectionAll() (count int64, err error) {
	res, err := service.sectionRepository.CountSectionAll()
	return res, err
}

func (service *sectionService) FindSections() (sectionOutput []model.SelectSectionParameter, err error) {
	res, err := service.sectionRepository.FindSections()
	return res, err
}

func (service *sectionService) FindSectionsOffset(limit int, offset int, order string, dir string) (sectionOutput []model.SelectSectionParameter, err error) {
	res, err := service.sectionRepository.FindSectionsOffset(limit, offset, order, dir)
	return res, err
}

func (service *sectionService) SearchSection(limit int, offset int, order string, dir string, search string) (sectionOutput []model.SelectSectionParameter, err error) {
	res, err := service.sectionRepository.SearchSection(limit, offset, order, dir, search)
	return res, err
}

func (service *sectionService) CountSearchSection(search string) (count int64, err error) {
	res, err := service.sectionRepository.CountSearchSection(search)
	return res, err
}

func (service *sectionService) FindSectionById(id uint) (sectionOutput model.SelectSectionParameter, err error) {
	res, err := service.sectionRepository.FindSectionById(id)
	return res, err
}

func (service *sectionService) FindExcSection(depId uint, id uint) (sectionOutput []model.SelectSectionParameter, err error) {
	res, err := service.sectionRepository.FindExcSection(depId, id)
	return res, err
}

func (service *sectionService) FindSectionByDepId(depId uint) (sectionOutput []model.SelectSectionParameter, err error) {
	res, err := service.sectionRepository.FindSectionByDepId(depId)
	return res, err
}

func (service *sectionService) FindSectionByDivisionID(divId uint) (sectionOutput []model.SelectSectionParameter, err error) {
	res, err := service.sectionRepository.FindSectionByDivisionID(divId)
	return res, err
}

func (service *sectionService) CountSectionName(search string) (count int64, err error) {
	res, err := service.sectionRepository.CountSectionName(search)
	return res, err
}

func (service *sectionService) InsertSection(section model.CreateSectionParameter) (sectionOutput model.Section, err error) {
	newSection := model.Section{}
	err1 := smapping.FillStruct(&newSection, smapping.MapFields(&section))
	if err != nil {
		return newSection, err1
	}
	res, err := service.sectionRepository.InsertSection(newSection)
	return res, err
}

func (service *sectionService) UpdateSection(section model.Section, id uint) (sectionOutput model.Section, err error) {
	newSection := model.Section{}
	err1 := smapping.FillStruct(&newSection, smapping.MapFields(&section))
	if err != nil {
		return newSection, err1
	}
	res, err := service.sectionRepository.UpdateSection(newSection, id)
	return res, err
}

func (service *sectionService) DeleteSection(section model.Section, id uint) (sectionOutput model.Section, err error) {
	newSection := model.Section{}
	err1 := smapping.FillStruct(&newSection, smapping.MapFields(&section))
	if err != nil {
		return newSection, err1
	}
	res, err := service.sectionRepository.UpdateSection(newSection, id)
	return res, err
}
