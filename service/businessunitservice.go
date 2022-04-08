package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"
	"time"

	"github.com/mashingan/smapping"
)

type BusinessUnitService interface {
	FindBusinessUnits() ([]model.BusinessUnit, error)
	FindBusinessUnitById(id uint) (model.BusinessUnit, error)
	InsertBusinessUnit(businessUnit model.CreateBusinessUnitParameter) (model.BusinessUnit, error)
	UpdateBusinessUnit(businessUnit model.BusinessUnit) (model.BusinessUnit, error)
	DeleteBusinessUnit(businessUnit model.BusinessUnit, userId uint) (model.BusinessUnit, error)
}

type businessUnitService struct {
	businessUnitRepository repository.BusinessUnitRepository
}

func NewBusinessUnitService(businessUnitRep repository.BusinessUnitRepository) BusinessUnitService {
	return &businessUnitService{
		businessUnitRepository: businessUnitRep,
	}
}

func (service *businessUnitService) FindBusinessUnits() ([]model.BusinessUnit, error) {
	return service.businessUnitRepository.FindBusinessUnits()
}

func (service *businessUnitService) FindBusinessUnitById(id uint) (model.BusinessUnit, error) {
	return service.businessUnitRepository.FindBusinessUnitById(id)
}

func (service *businessUnitService) InsertBusinessUnit(businessUnit model.CreateBusinessUnitParameter) (model.BusinessUnit, error) {
	newBusinessUnit := model.BusinessUnit{}
	err := smapping.FillStruct(&newBusinessUnit, smapping.MapFields(&businessUnit))

	if err != nil {
		return newBusinessUnit, err
	}

	return service.businessUnitRepository.InsertBusinessUnit(newBusinessUnit)
}

func (service *businessUnitService) UpdateBusinessUnit(businessUnit model.BusinessUnit) (model.BusinessUnit, error) {
	newBusinessUnit := model.BusinessUnit{}
	err := smapping.FillStruct(&newBusinessUnit, smapping.MapFields(&businessUnit))

	if err != nil {
		return newBusinessUnit, err
	}

	return service.businessUnitRepository.UpdateBusinessUnit(newBusinessUnit)
}

func (service *businessUnitService) DeleteBusinessUnit(businessUnit model.BusinessUnit, userID uint) (model.BusinessUnit, error) {
	now := time.Now()
	deletedAt := now.Unix()

	businessUnit.DeletedUserID = userID
	businessUnit.DeletedAt = float64(deletedAt)

	return service.businessUnitRepository.UpdateBusinessUnit(businessUnit)
}
