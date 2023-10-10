package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type NonWorkingTypeService interface {
	CountNonWorkingTypeAll() (count int64, err error)
	FindNonWorkingTypes() (nonworkingtypeOutput []model.NonWorkingType, err error)
	FindNonWorkingTypesOffset(limit int, offset int, order string, dir string) (nonworkingtypeOutput []model.NonWorkingType, err error)
	SearchNonWorkingType(limit int, offset int, order string, dir string, search string) (nonworkingtypeOutput []model.NonWorkingType, err error)
	CountSearchNonWorkingType(search string) (count int64, err error)
	FindNonWorkingTypeById(id uint) (nonworkingtypeOutput model.NonWorkingType, err error)
	FindExcNonWorkingType(id uint) (nonworkingtypeOutput []model.NonWorkingType, err error)
	CountNonWorkingTypeName(search string) (count int64, err error)
	InsertNonWorkingType(nonworkingtype model.CreateNonWorkingTypeParameter) (nonworkingtypeOutput model.NonWorkingType, err error)
	UpdateNonWorkingType(nonworkingtype model.NonWorkingType, id uint) (nonworkingtypeOutput model.NonWorkingType, err error)
	DeleteNonWorkingType(nonworkingtype model.NonWorkingType, id uint) (nonworkingtypeOutput model.NonWorkingType, err error)
}

type nonworkingtypeService struct {
	nonworkingtypeRepository repository.NonWorkingTypeRepository
}

func NewNonWorkingTypeService(nonworkingtypeRep repository.NonWorkingTypeRepository) NonWorkingTypeService {
	return &nonworkingtypeService{
		nonworkingtypeRepository: nonworkingtypeRep,
	}
}

func (service *nonworkingtypeService) CountNonWorkingTypeAll() (count int64, err error) {
	res, err := service.nonworkingtypeRepository.CountNonWorkingTypeAll()
	return res, err
}

func (service *nonworkingtypeService) FindNonWorkingTypes() (nonworkingtypeOutput []model.NonWorkingType, err error) {
	res, err := service.nonworkingtypeRepository.FindNonWorkingType()
	return res, err
}

func (service *nonworkingtypeService) FindNonWorkingTypesOffset(limit int, offset int, order string, dir string) (nonworkingtypeOutput []model.NonWorkingType, err error) {
	res, err := service.nonworkingtypeRepository.FindNonWorkingTypeOffset(limit, offset, order, dir)
	return res, err
}

func (service *nonworkingtypeService) SearchNonWorkingType(limit int, offset int, order string, dir string, search string) (nonworkingtypeOutput []model.NonWorkingType, err error) {
	res, err := service.nonworkingtypeRepository.SearchNonWorkingType(limit, offset, order, dir, search)
	return res, err
}

func (service *nonworkingtypeService) CountSearchNonWorkingType(search string) (count int64, err error) {
	res, err := service.nonworkingtypeRepository.CountSearchNonWorkingType(search)
	return res, err
}

func (service *nonworkingtypeService) FindNonWorkingTypeById(id uint) (nonworkingtypeOutput model.NonWorkingType, err error) {
	res, err := service.nonworkingtypeRepository.FindNonWorkingTypeById(id)
	return res, err
}

func (service *nonworkingtypeService) CountNonWorkingTypeName(search string) (count int64, err error) {
	res, err := service.nonworkingtypeRepository.CountNonWorkingTypeName(search)
	return res, err
}

func (service *nonworkingtypeService) FindExcNonWorkingType(id uint) (nonworkingtypeOutput []model.NonWorkingType, err error) {
	res, err := service.nonworkingtypeRepository.FindExcNonWorkingType(id)
	return res, err
}

func (service *nonworkingtypeService) InsertNonWorkingType(nonworkingtype model.CreateNonWorkingTypeParameter) (nonworkingtypeOutput model.NonWorkingType, err error) {
	newNonWorkingType := model.NonWorkingType{}
	err1 := smapping.FillStruct(&newNonWorkingType, smapping.MapFields(&nonworkingtype))
	if err != nil {
		return newNonWorkingType, err1
	}
	res, err := service.nonworkingtypeRepository.InsertNonWorkingType(newNonWorkingType)
	return res, err
}

func (service *nonworkingtypeService) UpdateNonWorkingType(nonworkingtype model.NonWorkingType, id uint) (nonworkingtypeOutput model.NonWorkingType, err error) {
	newNonWorkingType := model.NonWorkingType{}
	err1 := smapping.FillStruct(&newNonWorkingType, smapping.MapFields(&nonworkingtype))
	if err != nil {
		return newNonWorkingType, err1
	}
	res, err := service.nonworkingtypeRepository.UpdateNonWorkingType(newNonWorkingType, id)
	return res, err
}

func (service *nonworkingtypeService) DeleteNonWorkingType(nonworkingtype model.NonWorkingType, id uint) (nonworkingtypeOutput model.NonWorkingType, err error) {
	newNonWorkingType := model.NonWorkingType{}
	err1 := smapping.FillStruct(&newNonWorkingType, smapping.MapFields(&nonworkingtype))
	if err != nil {
		return newNonWorkingType, err1
	}
	res, err := service.nonworkingtypeRepository.DeleteNonWorkingType(newNonWorkingType, id)
	return res, err
}
