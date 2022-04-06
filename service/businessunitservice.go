package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"
	"log"
	"time"

	"github.com/mashingan/smapping"
)

type BusinessUnitService interface {
	FindBusinessUnits() []model.BusinessUnit
	FindBusinessUnitById(id uint) model.BusinessUnit
	InsertBusinessUnit(businessUnit model.CreateBusinessUnitParameter) model.BusinessUnit
	UpdateBusinessUnit(businessUnit model.BusinessUnit) model.BusinessUnit
	DeleteBusinessUnit(businessUnit model.BusinessUnit, userId uint) model.BusinessUnit
}

type businessUnitService struct {
	businessUnitRepository repository.BusinessUnitRepository
}

func NewBusinessUnitService(businessUnitRep repository.BusinessUnitRepository) BusinessUnitService {
	return &businessUnitService{
		businessUnitRepository: businessUnitRep,
	}
}

func (service *businessUnitService) FindBusinessUnits() []model.BusinessUnit {
	return service.businessUnitRepository.FindBusinessUnits()
}

func (service *businessUnitService) FindBusinessUnitById(id uint) model.BusinessUnit {
	return service.businessUnitRepository.FindBusinessUnitById(id)
}

func (service *businessUnitService) InsertBusinessUnit(businessUnit model.CreateBusinessUnitParameter) model.BusinessUnit {
	newBusinessUnit := model.BusinessUnit{}
	err := smapping.FillStruct(&newBusinessUnit, smapping.MapFields(&businessUnit))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.businessUnitRepository.InsertBusinessUnit(newBusinessUnit)
	return res
}

func (service *businessUnitService) UpdateBusinessUnit(businessUnit model.BusinessUnit) model.BusinessUnit {
	newBusinessUnit := model.BusinessUnit{}
	err := smapping.FillStruct(&newBusinessUnit, smapping.MapFields(&businessUnit))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.businessUnitRepository.UpdateBusinessUnit(newBusinessUnit)
	return res
}

func (service *businessUnitService) DeleteBusinessUnit(businessUnit model.BusinessUnit, userID uint) model.BusinessUnit {
	now := time.Now()
	deletedAt := now.Unix()

	businessUnit.DeletedUserID = userID
	businessUnit.DeletedAt = float64(deletedAt)

	res := service.businessUnitRepository.UpdateBusinessUnit(businessUnit)
	return res
}
