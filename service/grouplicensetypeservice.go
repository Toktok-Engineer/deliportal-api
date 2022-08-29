package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type GroupLicenseTypeService interface {
	CountGroupLicenseTypeAll() (count int64, err error)
	FindGroupLicenseTypes() (businessunitOutput []model.GroupLicenseType, err error)
	FindGroupLicenseTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.GroupLicenseType, err error)
	SearchGroupLicenseType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.GroupLicenseType, err error)
	CountSearchGroupLicenseType(search string) (count int64, err error)
	FindGroupLicenseTypeById(id uint) (groupLicenseTypeOutput model.GroupLicenseType, err error)
	FindExcGroupLicenseType(id uint) (groupLicenseTypeOutput []model.GroupLicenseType, err error)
	InsertGroupLicenseType(groupLicenseType model.CreateGroupLicenseTypeParameter) (groupLicenseTypeOutput model.GroupLicenseType, err error)
	UpdateGroupLicenseType(groupLicenseType model.GroupLicenseType, id uint) (groupLicenseTypeOutput model.GroupLicenseType, err error)
	DeleteGroupLicenseType(groupLicenseType model.GroupLicenseType, id uint) (groupLicenseTypeOutput model.GroupLicenseType, err error)
}

type groupLicenseTypeService struct {
	groupLicenseTypeRepository repository.GroupLicenseTypeRepository
}

func NewGroupLicenseTypeService(groupLicenseTypeRep repository.GroupLicenseTypeRepository) GroupLicenseTypeService {
	return &groupLicenseTypeService{
		groupLicenseTypeRepository: groupLicenseTypeRep,
	}
}

func (service *groupLicenseTypeService) CountGroupLicenseTypeAll() (count int64, err error) {
	res, err := service.groupLicenseTypeRepository.CountGroupLicenseTypeAll()
	return res, err
}

func (service *groupLicenseTypeService) FindGroupLicenseTypes() (businessunitOutput []model.GroupLicenseType, err error) {
	res, err := service.groupLicenseTypeRepository.FindGroupLicenseTypes()
	return res, err
}

func (service *groupLicenseTypeService) FindGroupLicenseTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.GroupLicenseType, err error) {
	res, err := service.groupLicenseTypeRepository.FindGroupLicenseTypesOffset(limit, offset, order, dir)
	return res, err
}

func (service *groupLicenseTypeService) SearchGroupLicenseType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.GroupLicenseType, err error) {
	res, err := service.groupLicenseTypeRepository.SearchGroupLicenseType(limit, offset, order, dir, search)
	return res, err
}

func (service *groupLicenseTypeService) CountSearchGroupLicenseType(search string) (count int64, err error) {
	res, err := service.groupLicenseTypeRepository.CountSearchGroupLicenseType(search)
	return res, err
}

func (service *groupLicenseTypeService) FindGroupLicenseTypeById(id uint) (groupLicenseTypeOutput model.GroupLicenseType, err error) {
	return service.groupLicenseTypeRepository.FindGroupLicenseTypeById(id)
}

func (service *groupLicenseTypeService) FindExcGroupLicenseType(id uint) (groupLicenseTypeOutput []model.GroupLicenseType, err error) {
	return service.groupLicenseTypeRepository.FindExcGroupLicenseType(id)
}
func (service *groupLicenseTypeService) InsertGroupLicenseType(groupLicenseType model.CreateGroupLicenseTypeParameter) (model.GroupLicenseType, error) {
	newGroupLicenseType := model.GroupLicenseType{}
	err1 := smapping.FillStruct(&newGroupLicenseType, smapping.MapFields(&groupLicenseType))

	if err1 != nil {
		return newGroupLicenseType, err1
	}

	return service.groupLicenseTypeRepository.InsertGroupLicenseType(newGroupLicenseType)
}

func (service *groupLicenseTypeService) UpdateGroupLicenseType(groupLicenseType model.GroupLicenseType, id uint) (groupLicenseTypeOutput model.GroupLicenseType, err error) {
	newGroupLicenseType := model.GroupLicenseType{}
	err1 := smapping.FillStruct(&newGroupLicenseType, smapping.MapFields(&groupLicenseType))

	if err1 != nil {
		return newGroupLicenseType, err1
	}

	return service.groupLicenseTypeRepository.UpdateGroupLicenseType(newGroupLicenseType, id)
}

func (service *groupLicenseTypeService) DeleteGroupLicenseType(groupLicenseType model.GroupLicenseType, id uint) (groupLicenseTypeOutput model.GroupLicenseType, err error) {
	newGroupLicenseType := model.GroupLicenseType{}
	err1 := smapping.FillStruct(&newGroupLicenseType, smapping.MapFields(&groupLicenseType))

	if err1 != nil {
		return newGroupLicenseType, err1
	}

	return service.groupLicenseTypeRepository.UpdateGroupLicenseType(newGroupLicenseType, id)
}
