package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"
	"log"
	"time"

	"github.com/mashingan/smapping"
)

type DivisionService interface {
	FindDivisions() []model.Division
	FindDivisionById(id uint) model.Division
	FindExcDivision(id uint) []model.Division
	InsertDivision(division model.CreateDivisionParameter) model.Division
	UpdateDivision(division model.Division) model.Division
	DeleteDivision(division model.Division, userId uint) model.Division
}

type divisionService struct {
	divisionRepository repository.DivisionRepository
}

func NewDivisionService(divisionRep repository.DivisionRepository) DivisionService {
	return &divisionService{
		divisionRepository: divisionRep,
	}
}

func (service *divisionService) FindDivisions() []model.Division {
	return service.divisionRepository.FindDivisions()
}

func (service *divisionService) FindDivisionById(id uint) model.Division {
	return service.divisionRepository.FindDivisionById(id)
}

func (service *divisionService) FindExcDivision(id uint) []model.Division {
	return service.divisionRepository.FindExcDivision(id)
}

func (service *divisionService) InsertDivision(division model.CreateDivisionParameter) model.Division {
	newDivision := model.Division{}
	err := smapping.FillStruct(&newDivision, smapping.MapFields(&division))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.divisionRepository.InsertDivision(newDivision)
	return res
}

func (service *divisionService) UpdateDivision(division model.Division) model.Division {
	newDivision := model.Division{}
	err := smapping.FillStruct(&newDivision, smapping.MapFields(&division))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.divisionRepository.UpdateDivision(newDivision)
	return res
}

func (service *divisionService) DeleteDivision(division model.Division, userID uint) model.Division {
	now := time.Now()
	deletedAt := now.Unix()

	division.DeletedUserID = userID
	division.DeletedAt = float64(deletedAt)

	res := service.divisionRepository.UpdateDivision(division)
	return res
}
