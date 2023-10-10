package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type EmployeeLeaveRequestDumpService interface {
	CountEmployeeLeaveRequestDumpAll() (count int64, err error)
	FindEmployeeLeaveRequestDumps() (businessunitOutput []model.EmployeeLeaveRequestDump, err error)
	FindEmployeeLeaveRequestDumpsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.EmployeeLeaveRequestDump, err error)
	// SearchEmployeeLeaveRequestDump(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.EmployeeLeaveRequestDump, err error)
	// CountSearchEmployeeLeaveRequestDump(search string) (count int64, err error)
	FindEmployeeLeaveRequestDumpById(id uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error)
	FindExcEmployeeLeaveRequestDump(id uint) (employeeLeaveRequestDumpOutput []model.EmployeeLeaveRequestDump, err error)
	InsertEmployeeLeaveRequestDump(employeeLeaveRequestDump model.CreateEmployeeLeaveRequestDumpParameter) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error)
	UpdateEmployeeLeaveRequestDump(employeeLeaveRequestDump model.EmployeeLeaveRequestDump, id uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error)
	DeleteEmployeeLeaveRequestDump(employeeLeaveRequestDump model.EmployeeLeaveRequestDump, id uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error)
}

type employeeLeaveRequestDumpService struct {
	employeeLeaveRequestDumpRepository repository.EmployeeLeaveRequestDumpRepository
}

func NewEmployeeLeaveRequestDumpService(employeeLeaveRequestDumpRep repository.EmployeeLeaveRequestDumpRepository) EmployeeLeaveRequestDumpService {
	return &employeeLeaveRequestDumpService{
		employeeLeaveRequestDumpRepository: employeeLeaveRequestDumpRep,
	}
}

func (service *employeeLeaveRequestDumpService) CountEmployeeLeaveRequestDumpAll() (count int64, err error) {
	res, err := service.employeeLeaveRequestDumpRepository.CountEmployeeLeaveRequestDumpAll()
	return res, err
}

func (service *employeeLeaveRequestDumpService) FindEmployeeLeaveRequestDumps() (businessunitOutput []model.EmployeeLeaveRequestDump, err error) {
	res, err := service.employeeLeaveRequestDumpRepository.FindEmployeeLeaveRequestDumps()
	return res, err
}

func (service *employeeLeaveRequestDumpService) FindEmployeeLeaveRequestDumpsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.EmployeeLeaveRequestDump, err error) {
	res, err := service.employeeLeaveRequestDumpRepository.FindEmployeeLeaveRequestDumpsOffset(limit, offset, order, dir)
	return res, err
}

// func (service *employeeLeaveRequestDumpService) SearchEmployeeLeaveRequestDump(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.EmployeeLeaveRequestDump, err error) {
// 	res, err := service.employeeLeaveRequestDumpRepository.SearchEmployeeLeaveRequestDump(limit, offset, order, dir, search)
// 	return res, err
// }

// func (service *employeeLeaveRequestDumpService) CountSearchEmployeeLeaveRequestDump(search string) (count int64, err error) {
// 	res, err := service.employeeLeaveRequestDumpRepository.CountSearchEmployeeLeaveRequestDump(search)
// 	return res, err
// }

func (service *employeeLeaveRequestDumpService) FindEmployeeLeaveRequestDumpById(id uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error) {
	return service.employeeLeaveRequestDumpRepository.FindEmployeeLeaveRequestDumpById(id)
}

func (service *employeeLeaveRequestDumpService) FindExcEmployeeLeaveRequestDump(id uint) (employeeLeaveRequestDumpOutput []model.EmployeeLeaveRequestDump, err error) {
	return service.employeeLeaveRequestDumpRepository.FindExcEmployeeLeaveRequestDump(id)
}
func (service *employeeLeaveRequestDumpService) InsertEmployeeLeaveRequestDump(employeeLeaveRequestDump model.CreateEmployeeLeaveRequestDumpParameter) (model.EmployeeLeaveRequestDump, error) {
	newEmployeeLeaveRequestDump := model.EmployeeLeaveRequestDump{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequestDump, smapping.MapFields(&employeeLeaveRequestDump))

	if err1 != nil {
		return newEmployeeLeaveRequestDump, err1
	}

	return service.employeeLeaveRequestDumpRepository.InsertEmployeeLeaveRequestDump(newEmployeeLeaveRequestDump)
}

func (service *employeeLeaveRequestDumpService) UpdateEmployeeLeaveRequestDump(employeeLeaveRequestDump model.EmployeeLeaveRequestDump, id uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error) {
	newEmployeeLeaveRequestDump := model.EmployeeLeaveRequestDump{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequestDump, smapping.MapFields(&employeeLeaveRequestDump))

	if err1 != nil {
		return newEmployeeLeaveRequestDump, err1
	}

	return service.employeeLeaveRequestDumpRepository.UpdateEmployeeLeaveRequestDump(newEmployeeLeaveRequestDump, id)
}

func (service *employeeLeaveRequestDumpService) DeleteEmployeeLeaveRequestDump(employeeLeaveRequestDump model.EmployeeLeaveRequestDump, id uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error) {
	newEmployeeLeaveRequestDump := model.EmployeeLeaveRequestDump{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequestDump, smapping.MapFields(&employeeLeaveRequestDump))

	if err1 != nil {
		return newEmployeeLeaveRequestDump, err1
	}

	return service.employeeLeaveRequestDumpRepository.UpdateEmployeeLeaveRequestDump(newEmployeeLeaveRequestDump, id)
}
