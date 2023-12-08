package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type EmployeeLeaveRequestApprovalService interface {
	CountEmployeeLeaveRequestApprovalAll(elrId int) (count int64, err error)
	FindEmployeeLeaveRequestApprovals() (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error)
	FindEmployeeLeaveRequestApprovalsOffset(limit int, offset int, order string, dir string, elrId int) (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error)
	SearchEmployeeLeaveRequestApproval(limit int, offset int, order string, dir string, search string, elrId int) (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error)
	CountSearchEmployeeLeaveRequestApproval(search string, elrId int) (count int64, err error)
	FindEmployeeLeaveRequestApprovalById(id uint) (employeeLeaveRequestApprovalOutput model.SelectEmployeeLeaveRequestApprovalParameter, err error)
	// FindExcEmployeeLeaveRequestApproval(divId uint, id uint) (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error)
	FindEmployeeLeaveRequestApprovalApproved(elrid uint) (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error)
	FindEmployeeLeaveRequestApprovalOpenApproved(elrid uint) (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error)
	FindEmployeeLeaveRequestApprovalOverwrite(elrid uint) (employeeLeaveRequestApprovalOutput model.SelectEmployeeLeaveRequestApprovalParameter, err error)
	FindEmployeeLeaveRequestApprovalByERId(elrid uint) (employeeLeaveRequestApprovalOutput model.SelectEmployeeLeaveRequestApprovalParameter, err error)
	// CountEmployeeLeaveRequestApprovalName(search string) (count int64, err error)
	InsertEmployeeLeaveRequestApproval(employeeLeaveRequestApproval model.CreateEmployeeLeaveRequestApprovalParameter) (employeeLeaveRequestApprovalOutput model.EmployeeLeaveRequestApproval, err error)
	UpdateEmployeeLeaveRequestApproval(employeeLeaveRequestApproval model.EmployeeLeaveRequestApproval, id uint) (employeeLeaveRequestApprovalOutput model.EmployeeLeaveRequestApproval, err error)
	DeleteEmployeeLeaveRequestApproval(employeeLeaveRequestApproval model.EmployeeLeaveRequestApproval, id uint) (employeeLeaveRequestApprovalOutput model.EmployeeLeaveRequestApproval, err error)
}

type employeeLeaveRequestApprovalService struct {
	employeeLeaveRequestApprovalRepository repository.EmployeeLeaveRequestApprovalRepository
}

func NewEmployeeLeaveRequestApprovalService(employeeLeaveRequestApprovalRep repository.EmployeeLeaveRequestApprovalRepository) EmployeeLeaveRequestApprovalService {
	return &employeeLeaveRequestApprovalService{
		employeeLeaveRequestApprovalRepository: employeeLeaveRequestApprovalRep,
	}
}

func (service *employeeLeaveRequestApprovalService) CountEmployeeLeaveRequestApprovalAll(elrId int) (count int64, err error) {
	res, err := service.employeeLeaveRequestApprovalRepository.CountEmployeeLeaveRequestApprovalAll(elrId)
	return res, err
}

func (service *employeeLeaveRequestApprovalService) FindEmployeeLeaveRequestApprovals() (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error) {
	res, err := service.employeeLeaveRequestApprovalRepository.FindEmployeeLeaveRequestApprovals()
	return res, err
}

func (service *employeeLeaveRequestApprovalService) FindEmployeeLeaveRequestApprovalsOffset(limit int, offset int, order string, dir string, elrId int) (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error) {
	res, err := service.employeeLeaveRequestApprovalRepository.FindEmployeeLeaveRequestApprovalsOffset(limit, offset, order, dir, elrId)
	return res, err
}

func (service *employeeLeaveRequestApprovalService) SearchEmployeeLeaveRequestApproval(limit int, offset int, order string, dir string, search string, elrId int) (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error) {
	res, err := service.employeeLeaveRequestApprovalRepository.SearchEmployeeLeaveRequestApproval(limit, offset, order, dir, search, elrId)
	return res, err
}

func (service *employeeLeaveRequestApprovalService) CountSearchEmployeeLeaveRequestApproval(search string, elrId int) (count int64, err error) {
	res, err := service.employeeLeaveRequestApprovalRepository.CountSearchEmployeeLeaveRequestApproval(search, elrId)
	return res, err
}

func (service *employeeLeaveRequestApprovalService) FindEmployeeLeaveRequestApprovalById(id uint) (employeeLeaveRequestApprovalOutput model.SelectEmployeeLeaveRequestApprovalParameter, err error) {
	res, err := service.employeeLeaveRequestApprovalRepository.FindEmployeeLeaveRequestApprovalById(id)
	return res, err
}

// func (service *employeeLeaveRequestApprovalService) FindExcEmployeeLeaveRequestApproval(divId uint, id uint) (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error) {
// 	res, err := service.employeeLeaveRequestApprovalRepository.FindExcEmployeeLeaveRequestApproval(divId, id)
// 	return res, err
// }

func (service *employeeLeaveRequestApprovalService) FindEmployeeLeaveRequestApprovalApproved(erlid uint) (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error) {
	res, err := service.employeeLeaveRequestApprovalRepository.FindEmployeeLeaveRequestApprovalApproved(erlid)
	return res, err
}

func (service *employeeLeaveRequestApprovalService) FindEmployeeLeaveRequestApprovalOpenApproved(erlid uint) (employeeLeaveRequestApprovalOutput []model.SelectEmployeeLeaveRequestApprovalParameter, err error) {
	res, err := service.employeeLeaveRequestApprovalRepository.FindEmployeeLeaveRequestApprovalOpenApproved(erlid)
	return res, err
}

func (service *employeeLeaveRequestApprovalService) FindEmployeeLeaveRequestApprovalOverwrite(elrid uint) (employeeLeaveRequestApprovalOutput model.SelectEmployeeLeaveRequestApprovalParameter, err error) {
	res, err := service.employeeLeaveRequestApprovalRepository.FindEmployeeLeaveRequestApprovalOverwrite(elrid)
	return res, err
}
func (service *employeeLeaveRequestApprovalService) FindEmployeeLeaveRequestApprovalByERId(erlid uint) (employeeLeaveRequestApprovalOutput model.SelectEmployeeLeaveRequestApprovalParameter, err error) {
	res, err := service.employeeLeaveRequestApprovalRepository.FindEmployeeLeaveRequestApprovalByERId(erlid)
	return res, err
}

// func (service *employeeLeaveRequestApprovalService) CountEmployeeLeaveRequestApprovalName(search string) (count int64, err error) {
// 	res, err := service.employeeLeaveRequestApprovalRepository.CountEmployeeLeaveRequestApprovalName(search)
// 	return res, err
// }

func (service *employeeLeaveRequestApprovalService) InsertEmployeeLeaveRequestApproval(employeeLeaveRequestApproval model.CreateEmployeeLeaveRequestApprovalParameter) (employeeLeaveRequestApprovalOutput model.EmployeeLeaveRequestApproval, err error) {
	newEmployeeLeaveRequestApproval := model.EmployeeLeaveRequestApproval{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequestApproval, smapping.MapFields(&employeeLeaveRequestApproval))
	if err != nil {
		return newEmployeeLeaveRequestApproval, err1
	}
	res, err := service.employeeLeaveRequestApprovalRepository.InsertEmployeeLeaveRequestApproval(newEmployeeLeaveRequestApproval)
	return res, err
}

func (service *employeeLeaveRequestApprovalService) UpdateEmployeeLeaveRequestApproval(employeeLeaveRequestApproval model.EmployeeLeaveRequestApproval, id uint) (employeeLeaveRequestApprovalOutput model.EmployeeLeaveRequestApproval, err error) {
	newEmployeeLeaveRequestApproval := model.EmployeeLeaveRequestApproval{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequestApproval, smapping.MapFields(&employeeLeaveRequestApproval))
	if err != nil {
		return newEmployeeLeaveRequestApproval, err1
	}
	res, err := service.employeeLeaveRequestApprovalRepository.UpdateEmployeeLeaveRequestApproval(newEmployeeLeaveRequestApproval, id)
	return res, err
}

func (service *employeeLeaveRequestApprovalService) DeleteEmployeeLeaveRequestApproval(employeeLeaveRequestApproval model.EmployeeLeaveRequestApproval, id uint) (employeeLeaveRequestApprovalOutput model.EmployeeLeaveRequestApproval, err error) {
	newEmployeeLeaveRequestApproval := model.EmployeeLeaveRequestApproval{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequestApproval, smapping.MapFields(&employeeLeaveRequestApproval))
	if err != nil {
		return newEmployeeLeaveRequestApproval, err1
	}
	res, err := service.employeeLeaveRequestApprovalRepository.UpdateEmployeeLeaveRequestApproval(newEmployeeLeaveRequestApproval, id)
	return res, err
}
