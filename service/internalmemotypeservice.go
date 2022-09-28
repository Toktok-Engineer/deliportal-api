package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type InternalMemoTypeService interface {
	CountInternalMemoTypeAll() (count int64, err error)
	FindInternalMemoTypes() (businessunitOutput []model.InternalMemoType, err error)
	FindInternalMemoTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.InternalMemoType, err error)
	SearchInternalMemoType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.InternalMemoType, err error)
	CountSearchInternalMemoType(search string) (count int64, err error)
	FindInternalMemoTypeById(id uint) (internalMemoTypeOutput model.InternalMemoType, err error)
	FindExcInternalMemoType(id uint) (internalMemoTypeOutput []model.InternalMemoType, err error)
	InsertInternalMemoType(internalMemoType model.CreateInternalMemoTypeParameter) (internalMemoTypeOutput model.InternalMemoType, err error)
	UpdateInternalMemoType(internalMemoType model.InternalMemoType, id uint) (internalMemoTypeOutput model.InternalMemoType, err error)
	DeleteInternalMemoType(internalMemoType model.InternalMemoType, id uint) (internalMemoTypeOutput model.InternalMemoType, err error)
}

type internalMemoTypeService struct {
	internalMemoTypeRepository repository.InternalMemoTypeRepository
}

func NewInternalMemoTypeService(internalMemoTypeRep repository.InternalMemoTypeRepository) InternalMemoTypeService {
	return &internalMemoTypeService{
		internalMemoTypeRepository: internalMemoTypeRep,
	}
}

func (service *internalMemoTypeService) CountInternalMemoTypeAll() (count int64, err error) {
	res, err := service.internalMemoTypeRepository.CountInternalMemoTypeAll()
	return res, err
}

func (service *internalMemoTypeService) FindInternalMemoTypes() (businessunitOutput []model.InternalMemoType, err error) {
	res, err := service.internalMemoTypeRepository.FindInternalMemoTypes()
	return res, err
}

func (service *internalMemoTypeService) FindInternalMemoTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.InternalMemoType, err error) {
	res, err := service.internalMemoTypeRepository.FindInternalMemoTypesOffset(limit, offset, order, dir)
	return res, err
}

func (service *internalMemoTypeService) SearchInternalMemoType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.InternalMemoType, err error) {
	res, err := service.internalMemoTypeRepository.SearchInternalMemoType(limit, offset, order, dir, search)
	return res, err
}

func (service *internalMemoTypeService) CountSearchInternalMemoType(search string) (count int64, err error) {
	res, err := service.internalMemoTypeRepository.CountSearchInternalMemoType(search)
	return res, err
}

func (service *internalMemoTypeService) FindInternalMemoTypeById(id uint) (internalMemoTypeOutput model.InternalMemoType, err error) {
	return service.internalMemoTypeRepository.FindInternalMemoTypeById(id)
}

func (service *internalMemoTypeService) FindExcInternalMemoType(id uint) (internalMemoTypeOutput []model.InternalMemoType, err error) {
	return service.internalMemoTypeRepository.FindExcInternalMemoType(id)
}
func (service *internalMemoTypeService) InsertInternalMemoType(internalMemoType model.CreateInternalMemoTypeParameter) (model.InternalMemoType, error) {
	newInternalMemoType := model.InternalMemoType{}
	err1 := smapping.FillStruct(&newInternalMemoType, smapping.MapFields(&internalMemoType))

	if err1 != nil {
		return newInternalMemoType, err1
	}

	return service.internalMemoTypeRepository.InsertInternalMemoType(newInternalMemoType)
}

func (service *internalMemoTypeService) UpdateInternalMemoType(internalMemoType model.InternalMemoType, id uint) (internalMemoTypeOutput model.InternalMemoType, err error) {
	newInternalMemoType := model.InternalMemoType{}
	err1 := smapping.FillStruct(&newInternalMemoType, smapping.MapFields(&internalMemoType))

	if err1 != nil {
		return newInternalMemoType, err1
	}

	return service.internalMemoTypeRepository.UpdateInternalMemoType(newInternalMemoType, id)
}

func (service *internalMemoTypeService) DeleteInternalMemoType(internalMemoType model.InternalMemoType, id uint) (internalMemoTypeOutput model.InternalMemoType, err error) {
	newInternalMemoType := model.InternalMemoType{}
	err1 := smapping.FillStruct(&newInternalMemoType, smapping.MapFields(&internalMemoType))

	if err1 != nil {
		return newInternalMemoType, err1
	}

	return service.internalMemoTypeRepository.UpdateInternalMemoType(newInternalMemoType, id)
}
