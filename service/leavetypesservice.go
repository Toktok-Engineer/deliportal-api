package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type LeaveTypesService interface {
	CountLeaveTypesAll() (count int64, err error)
	FindLeaveTypess() (leavetypesOutput []model.LeaveTypes, err error)
	FindLeaveTypessOffset(limit int, offset int, order string, dir string) (leavetypesOutput []model.LeaveTypes, err error)
	SearchLeaveTypes(limit int, offset int, order string, dir string, search string) (leavetypesOutput []model.LeaveTypes, err error)
	CountSearchLeaveTypes(search string) (count int64, err error)
	FindLeaveTypesById(id uint) (leavetypesOutput model.LeaveTypes, err error)
	FindExcLeaveTypes(id uint) (leavetypesOutput []model.LeaveTypes, err error)
	CountLeaveTypesName(search string) (count int64, err error)
	InsertLeaveTypes(leavetypes model.CreateLeaveTypesParameter) (leavetypesOutput model.LeaveTypes, err error)
	UpdateLeaveTypes(leavetypes model.LeaveTypes, id uint) (leavetypesOutput model.LeaveTypes, err error)
	DeleteLeaveTypes(leavetypes model.LeaveTypes, id uint) (leavetypesOutput model.LeaveTypes, err error)
}

type leavetypesService struct {
	leavetypesRepository repository.LeaveTypesRepository
}

func NewLeaveTypesService(leavetypesRep repository.LeaveTypesRepository) LeaveTypesService {
	return &leavetypesService{
		leavetypesRepository: leavetypesRep,
	}
}

func (service *leavetypesService) CountLeaveTypesAll() (count int64, err error) {
	res, err := service.leavetypesRepository.CountLeaveTypeAll()
	return res, err
}

func (service *leavetypesService) FindLeaveTypess() (leavetypesOutput []model.LeaveTypes, err error) {
	res, err := service.leavetypesRepository.FindLeaveTypes()
	return res, err
}

func (service *leavetypesService) FindLeaveTypessOffset(limit int, offset int, order string, dir string) (leavetypesOutput []model.LeaveTypes, err error) {
	res, err := service.leavetypesRepository.FindLeaveTypesOffset(limit, offset, order, dir)
	return res, err
}

func (service *leavetypesService) SearchLeaveTypes(limit int, offset int, order string, dir string, search string) (leavetypesOutput []model.LeaveTypes, err error) {
	res, err := service.leavetypesRepository.SearchLeaveType(limit, offset, order, dir, search)
	return res, err
}

func (service *leavetypesService) CountSearchLeaveTypes(search string) (count int64, err error) {
	res, err := service.leavetypesRepository.CountSearchLeaveType(search)
	return res, err
}

func (service *leavetypesService) FindLeaveTypesById(id uint) (leavetypesOutput model.LeaveTypes, err error) {
	res, err := service.leavetypesRepository.FindLeaveTypeById(id)
	return res, err
}

func (service *leavetypesService) CountLeaveTypesName(search string) (count int64, err error) {
	res, err := service.leavetypesRepository.CountLeaveTypeName(search)
	return res, err
}

func (service *leavetypesService) FindExcLeaveTypes(id uint) (leavetypesOutput []model.LeaveTypes, err error) {
	res, err := service.leavetypesRepository.FindExcLeaveType(id)
	return res, err
}

func (service *leavetypesService) InsertLeaveTypes(leavetypes model.CreateLeaveTypesParameter) (leavetypesOutput model.LeaveTypes, err error) {
	newLeaveTypes := model.LeaveTypes{}
	err1 := smapping.FillStruct(&newLeaveTypes, smapping.MapFields(&leavetypes))
	if err != nil {
		return newLeaveTypes, err1
	}
	res, err := service.leavetypesRepository.InsertLeaveType(newLeaveTypes)
	return res, err
}

func (service *leavetypesService) UpdateLeaveTypes(leavetypes model.LeaveTypes, id uint) (leavetypesOutput model.LeaveTypes, err error) {
	newLeaveTypes := model.LeaveTypes{}
	err1 := smapping.FillStruct(&newLeaveTypes, smapping.MapFields(&leavetypes))
	if err != nil {
		return newLeaveTypes, err1
	}
	res, err := service.leavetypesRepository.UpdateLeaveType(newLeaveTypes, id)
	return res, err
}

func (service *leavetypesService) DeleteLeaveTypes(leavetypes model.LeaveTypes, id uint) (leavetypesOutput model.LeaveTypes, err error) {
	newLeaveTypes := model.LeaveTypes{}
	err1 := smapping.FillStruct(&newLeaveTypes, smapping.MapFields(&leavetypes))
	if err != nil {
		return newLeaveTypes, err1
	}
	res, err := service.leavetypesRepository.UpdateLeaveType(newLeaveTypes, id)
	return res, err
}
