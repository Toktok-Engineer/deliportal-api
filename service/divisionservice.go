package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type DivisionService interface {
	CountDivisionAll() (count int64, err error)
	FindDivisions() (divisionOutput []model.Division, err error)
	FindDivisionsOffset(limit int, offset int, order string, dir string) (divisionOutput []model.Division, err error)
	SearchDivision(limit int, offset int, order string, dir string, search string) (divisionOutput []model.Division, err error)
	CountSearchDivision(search string) (count int64, err error)
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

func (service *divisionService) CountDivisionAll() (count int64, err error) {
	res, err := service.divisionRepository.CountDivisionAll()
	return res, err
}

func (service *divisionService) FindDivisions() (divisionOutput []model.Division, err error) {
	res, err := service.divisionRepository.FindDivisions()
	return res, err
}

func (service *divisionService) FindDivisionsOffset(limit int, offset int, order string, dir string) (divisionOutput []model.Division, err error) {
	res, err := service.divisionRepository.FindDivisionsOffset(limit, offset, order, dir)
	return res, err
}

func (service *divisionService) SearchDivision(limit int, offset int, order string, dir string, search string) (divisionOutput []model.Division, err error) {
	res, err := service.divisionRepository.SearchDivision(limit, offset, order, dir, search)
	return res, err
}

func (service *divisionService) CountSearchDivision(search string) (count int64, err error) {
	res, err := service.divisionRepository.CountSearchDivision(search)
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
