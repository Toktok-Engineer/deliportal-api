package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type ImDefaultCirculationService interface {
	CountImDefaultCirculationAll(companyGroup int) (count int64, err error)
	FindImDefaultCirculations(companyGroupID int) (businessunitOutput []model.ImDefaultCirculation, err error)
	FindImDefaultCirculationsOffset(limit int, offset int, order string, dir string, companyGroup int) (businessunitOutput []model.SelectImDefaultCirculationParameter, err error)
	SearchImDefaultCirculation(limit int, offset int, order string, dir string, search string, companyGroup int) (businessunitOutput []model.SelectImDefaultCirculationParameter, err error)
	CountSearchImDefaultCirculation(search string, companyGroup int) (count int64, err error)
	FindImDefaultCirculationById(id uint) (imDefaultCirculationOutput model.ImDefaultCirculation, err error)
	FindExcImDefaultCirculation(id uint) (imDefaultCirculationOutput []model.ImDefaultCirculation, err error)
	InsertImDefaultCirculation(imDefaultCirculation model.CreateImDefaultCirculationParameter) (imDefaultCirculationOutput model.ImDefaultCirculation, err error)
	UpdateImDefaultCirculation(imDefaultCirculation model.ImDefaultCirculation, id uint) (imDefaultCirculationOutput model.ImDefaultCirculation, err error)
	DeleteImDefaultCirculation(imDefaultCirculation model.ImDefaultCirculation, id uint) (imDefaultCirculationOutput model.ImDefaultCirculation, err error)
}

type imDefaultCirculationService struct {
	imDefaultCirculationRepository repository.ImDefaultCirculationRepository
}

func NewImDefaultCirculationService(imDefaultCirculationRep repository.ImDefaultCirculationRepository) ImDefaultCirculationService {
	return &imDefaultCirculationService{
		imDefaultCirculationRepository: imDefaultCirculationRep,
	}
}

func (service *imDefaultCirculationService) CountImDefaultCirculationAll(companyGroup int) (count int64, err error) {
	res, err := service.imDefaultCirculationRepository.CountImDefaultCirculationAll(companyGroup)
	return res, err
}

func (service *imDefaultCirculationService) FindImDefaultCirculations(companyGroupID int) (businessunitOutput []model.ImDefaultCirculation, err error) {
	res, err := service.imDefaultCirculationRepository.FindImDefaultCirculations(companyGroupID)
	return res, err
}

func (service *imDefaultCirculationService) FindImDefaultCirculationsOffset(limit int, offset int, order string, dir string, companyGroup int) (businessunitOutput []model.SelectImDefaultCirculationParameter, err error) {
	res, err := service.imDefaultCirculationRepository.FindImDefaultCirculationsOffset(limit, offset, order, dir, companyGroup)
	return res, err
}

func (service *imDefaultCirculationService) SearchImDefaultCirculation(limit int, offset int, order string, dir string, search string, companyGroup int) (businessunitOutput []model.SelectImDefaultCirculationParameter, err error) {
	res, err := service.imDefaultCirculationRepository.SearchImDefaultCirculation(limit, offset, order, dir, search, companyGroup)
	return res, err
}

func (service *imDefaultCirculationService) CountSearchImDefaultCirculation(search string, companyGroup int) (count int64, err error) {
	res, err := service.imDefaultCirculationRepository.CountSearchImDefaultCirculation(search, companyGroup)
	return res, err
}

func (service *imDefaultCirculationService) FindImDefaultCirculationById(id uint) (imDefaultCirculationOutput model.ImDefaultCirculation, err error) {
	return service.imDefaultCirculationRepository.FindImDefaultCirculationById(id)
}

func (service *imDefaultCirculationService) FindExcImDefaultCirculation(id uint) (imDefaultCirculationOutput []model.ImDefaultCirculation, err error) {
	return service.imDefaultCirculationRepository.FindExcImDefaultCirculation(id)
}
func (service *imDefaultCirculationService) InsertImDefaultCirculation(imDefaultCirculation model.CreateImDefaultCirculationParameter) (model.ImDefaultCirculation, error) {
	newImDefaultCirculation := model.ImDefaultCirculation{}
	err1 := smapping.FillStruct(&newImDefaultCirculation, smapping.MapFields(&imDefaultCirculation))

	if err1 != nil {
		return newImDefaultCirculation, err1
	}

	return service.imDefaultCirculationRepository.InsertImDefaultCirculation(newImDefaultCirculation)
}

func (service *imDefaultCirculationService) UpdateImDefaultCirculation(imDefaultCirculation model.ImDefaultCirculation, id uint) (imDefaultCirculationOutput model.ImDefaultCirculation, err error) {
	newImDefaultCirculation := model.ImDefaultCirculation{}
	err1 := smapping.FillStruct(&newImDefaultCirculation, smapping.MapFields(&imDefaultCirculation))

	if err1 != nil {
		return newImDefaultCirculation, err1
	}

	return service.imDefaultCirculationRepository.UpdateImDefaultCirculation(newImDefaultCirculation, id)
}

func (service *imDefaultCirculationService) DeleteImDefaultCirculation(imDefaultCirculation model.ImDefaultCirculation, id uint) (imDefaultCirculationOutput model.ImDefaultCirculation, err error) {
	newImDefaultCirculation := model.ImDefaultCirculation{}
	err1 := smapping.FillStruct(&newImDefaultCirculation, smapping.MapFields(&imDefaultCirculation))

	if err1 != nil {
		return newImDefaultCirculation, err1
	}

	return service.imDefaultCirculationRepository.UpdateImDefaultCirculation(newImDefaultCirculation, id)
}
