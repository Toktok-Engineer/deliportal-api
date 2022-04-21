package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type BusinessUnitService interface {
	FindBusinessUnits() (businessUnitOutput []model.BusinessUnit, err error)
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

func (service *businessUnitService) FindBusinessUnits() (businessUnitOutput []model.BusinessUnit, err error) {
	return service.businessUnitRepository.FindBusinessUnits()
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
