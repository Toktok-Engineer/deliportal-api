package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type DivisionService interface {
	FindDivisions() (divisionOutput []model.Division, err error)
	FindDivisionById(id uint) (divisionOutput model.Division, err error)
	FindExcDivision(id uint) (divisionOutput []model.Division, err error)
	InsertDivision(division model.CreateDivisionParameter) (divisionOutput model.Division, err error)
	UpdateDivision(division model.Division, id uint) (divisionOutput model.Division, err error)
	DeleteDivision(division model.Division, id uint) (divisionOutput model.Division, err error)
}

type divisionService struct {
	divisionRepository repository.DivisionRepository
}

func NewDivisionService(divisionRep repository.DivisionRepository) DivisionService {
	return &divisionService{
		divisionRepository: divisionRep,
	}
}

func (service *divisionService) FindDivisions() (divisionOutput []model.Division, err error) {
	res, err := service.divisionRepository.FindDivisions()
	return res, err
}

func (service *divisionService) FindDivisionById(id uint) (divisionOutput model.Division, err error) {
	res, err := service.divisionRepository.FindDivisionById(id)
	return res, err
}

func (service *divisionService) FindExcDivision(id uint) (divisionOutput []model.Division, err error) {
	res, err := service.divisionRepository.FindExcDivision(id)
	return res, err
}

func (service *divisionService) InsertDivision(division model.CreateDivisionParameter) (divisionOutput model.Division, err error) {
	newDivision := model.Division{}
	err1 := smapping.FillStruct(&newDivision, smapping.MapFields(&division))
	if err != nil {
		return newDivision, err1
	}
	res, err := service.divisionRepository.InsertDivision(newDivision)
	return res, err
}

func (service *divisionService) UpdateDivision(division model.Division, id uint) (divisionOutput model.Division, err error) {
	newDivision := model.Division{}
	err1 := smapping.FillStruct(&newDivision, smapping.MapFields(&division))
	if err != nil {
		return newDivision, err1
	}
	res, err := service.divisionRepository.UpdateDivision(newDivision, id)
	return res, err
}

func (service *divisionService) DeleteDivision(division model.Division, id uint) (divisionOutput model.Division, err error) {
	newDivision := model.Division{}
	err1 := smapping.FillStruct(&newDivision, smapping.MapFields(&division))
	if err != nil {
		return newDivision, err1
	}
	res, err := service.divisionRepository.UpdateDivision(newDivision, id)
	return res, err
}
