package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type BusinessUnitService interface {
	CountBusinessUnitAll() (count int64, err error)
	FindBusinessUnits() (businessunitOutput []model.BusinessUnit, err error)
	FindBusinessUnitsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.BusinessUnit, err error)
	SearchBusinessUnit(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.BusinessUnit, err error)
	CountSearchBusinessUnit(search string) (count int64, err error)
	FindBusinessUnitById(id uint) (businessUnitOutput model.BusinessUnit, err error)
	FindExcBusinessUnit(id uint) (businessUnitOutput []model.BusinessUnit, err error)
	InsertBusinessUnit(businessUnit model.CreateBusinessUnitParameter) (businessUnitOutput model.BusinessUnit, err error)
	UpdateBusinessUnit(businessUnit model.BusinessUnit, id uint) (businessUnitOutput model.BusinessUnit, err error)
	DeleteBusinessUnit(businessUnit model.BusinessUnit, id uint) (businessUnitOutput model.BusinessUnit, err error)
}

type businessUnitService struct {
	businessUnitRepository repository.BusinessUnitRepository
}

func NewBusinessUnitService(businessUnitRep repository.BusinessUnitRepository) BusinessUnitService {
	return &businessUnitService{
		businessUnitRepository: businessUnitRep,
	}
}

func (service *businessUnitService) CountBusinessUnitAll() (count int64, err error) {
	res, err := service.businessUnitRepository.CountBusinessUnitAll()
	return res, err
}

func (service *businessUnitService) FindBusinessUnits() (businessunitOutput []model.BusinessUnit, err error) {
	res, err := service.businessUnitRepository.FindBusinessUnits()
	return res, err
}

func (service *businessUnitService) FindBusinessUnitsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.BusinessUnit, err error) {
	res, err := service.businessUnitRepository.FindBusinessUnitsOffset(limit, offset, order, dir)
	return res, err
}

func (service *businessUnitService) SearchBusinessUnit(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.BusinessUnit, err error) {
	res, err := service.businessUnitRepository.SearchBusinessUnit(limit, offset, order, dir, search)
	return res, err
}

func (service *businessUnitService) CountSearchBusinessUnit(search string) (count int64, err error) {
	res, err := service.businessUnitRepository.CountSearchBusinessUnit(search)
	return res, err
}

func (service *businessUnitService) FindBusinessUnitById(id uint) (businessUnitOutput model.BusinessUnit, err error) {
	return service.businessUnitRepository.FindBusinessUnitById(id)
}

func (service *businessUnitService) FindExcBusinessUnit(id uint) (businessUnitOutput []model.BusinessUnit, err error) {
	return service.businessUnitRepository.FindExcBusinessUnit(id)
}
func (service *businessUnitService) InsertBusinessUnit(businessUnit model.CreateBusinessUnitParameter) (model.BusinessUnit, error) {
	newBusinessUnit := model.BusinessUnit{}
	err1 := smapping.FillStruct(&newBusinessUnit, smapping.MapFields(&businessUnit))

	if err1 != nil {
		return newBusinessUnit, err1
	}

	return service.businessUnitRepository.InsertBusinessUnit(newBusinessUnit)
}

func (service *businessUnitService) UpdateBusinessUnit(businessUnit model.BusinessUnit, id uint) (businessUnitOutput model.BusinessUnit, err error) {
	newBusinessUnit := model.BusinessUnit{}
	err1 := smapping.FillStruct(&newBusinessUnit, smapping.MapFields(&businessUnit))

	if err1 != nil {
		return newBusinessUnit, err1
	}

	return service.businessUnitRepository.UpdateBusinessUnit(newBusinessUnit, id)
}

func (service *businessUnitService) DeleteBusinessUnit(businessUnit model.BusinessUnit, id uint) (businessUnitOutput model.BusinessUnit, err error) {
	newBusinessUnit := model.BusinessUnit{}
	err1 := smapping.FillStruct(&newBusinessUnit, smapping.MapFields(&businessUnit))

	if err1 != nil {
		return newBusinessUnit, err1
	}

	return service.businessUnitRepository.UpdateBusinessUnit(newBusinessUnit, id)
}
