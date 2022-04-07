package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type SectionService interface {
	FindSections() (sectionOutput []model.SelectSectionParameter, err error)
	FindSectionById(id uint) (sectionOutput model.SelectSectionParameter, err error)
	FindExcSection(depId uint, id uint) (sectionOutput []model.SelectSectionParameter, err error)
	FindSectionByDepId(depId uint) (sectionOutput []model.SelectSectionParameter, err error)
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

func (service *sectionService) FindSections() (sectionOutput []model.SelectSectionParameter, err error) {
	res, err := service.sectionRepository.FindSections()
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

func (service *sectionService) InsertSection(section model.CreateSectionParameter) (sectionOutput model.Section, err error) {
	newSection := model.Section{}
	err1 := smapping.FillStruct(&newSection, smapping.MapFields(&section))
	if err != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.sectionRepository.InsertSection(newSection)
	return res, err
}

func (service *sectionService) UpdateSection(section model.Section, id uint) (sectionOutput model.Section, err error) {
	newSection := model.Section{}
	err1 := smapping.FillStruct(&newSection, smapping.MapFields(&section))
	if err != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.sectionRepository.UpdateSection(newSection, id)
	return res, err
}

func (service *sectionService) DeleteSection(section model.Section, id uint) (sectionOutput model.Section, err error) {
	newSection := model.Section{}
	err1 := smapping.FillStruct(&newSection, smapping.MapFields(&section))
	if err != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.sectionRepository.UpdateSection(newSection, id)
	return res, err
}
