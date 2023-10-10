package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type EmployeeLeaveRequestTracingService interface {
	CountEmployeeLeaveRequestTracingAll(employeeLeaveRequestId int) (count int64, err error)
	FindEmployeeLeaveRequestTracings(employeeLeaveRequestId int) (internalmemotracingOutput []model.EmployeeLeaveRequestTracing, err error)
	FindEmployeeLeaveRequestTracingById(id uint) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error)
	FindExcEmployeeLeaveRequestTracing(id uint) (employeeLeaveRequestTracingOutput []model.EmployeeLeaveRequestTracing, err error)
	InsertEmployeeLeaveRequestTracing(employeeLeaveRequestTracing model.CreateEmployeeLeaveRequestTracingParameter) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error)
	UpdateEmployeeLeaveRequestTracing(employeeLeaveRequestTracing model.EmployeeLeaveRequestTracing, id uint) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error)
	DeleteEmployeeLeaveRequestTracing(employeeLeaveRequestTracing model.EmployeeLeaveRequestTracing, id uint) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error)
}

type employeeLeaveRequestTracingService struct {
	employeeLeaveRequestTracingRepository repository.EmployeeLeaveRequestTracingRepository
}

func NewEmployeeLeaveRequestTracingService(employeeLeaveRequestTracingRep repository.EmployeeLeaveRequestTracingRepository) EmployeeLeaveRequestTracingService {
	return &employeeLeaveRequestTracingService{
		employeeLeaveRequestTracingRepository: employeeLeaveRequestTracingRep,
	}
}

func (service *employeeLeaveRequestTracingService) CountEmployeeLeaveRequestTracingAll(employeeLeaveRequestId int) (count int64, err error) {
	res, err := service.employeeLeaveRequestTracingRepository.CountEmployeeLeaveRequestTracingAll(employeeLeaveRequestId)
	return res, err
}

func (service *employeeLeaveRequestTracingService) FindEmployeeLeaveRequestTracings(employeeLeaveRequestId int) (internalmemotracingOutput []model.EmployeeLeaveRequestTracing, err error) {
	res, err := service.employeeLeaveRequestTracingRepository.FindEmployeeLeaveRequestTracings(employeeLeaveRequestId)
	return res, err
}

func (service *employeeLeaveRequestTracingService) FindEmployeeLeaveRequestTracingById(id uint) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error) {
	return service.employeeLeaveRequestTracingRepository.FindEmployeeLeaveRequestTracingById(id)
}

func (service *employeeLeaveRequestTracingService) FindExcEmployeeLeaveRequestTracing(id uint) (employeeLeaveRequestTracingOutput []model.EmployeeLeaveRequestTracing, err error) {
	return service.employeeLeaveRequestTracingRepository.FindExcEmployeeLeaveRequestTracing(id)
}
func (service *employeeLeaveRequestTracingService) InsertEmployeeLeaveRequestTracing(employeeLeaveRequestTracing model.CreateEmployeeLeaveRequestTracingParameter) (model.EmployeeLeaveRequestTracing, error) {
	newEmployeeLeaveRequestTracing := model.EmployeeLeaveRequestTracing{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequestTracing, smapping.MapFields(&employeeLeaveRequestTracing))

	if err1 != nil {
		return newEmployeeLeaveRequestTracing, err1
	}

	return service.employeeLeaveRequestTracingRepository.InsertEmployeeLeaveRequestTracing(newEmployeeLeaveRequestTracing)
}

func (service *employeeLeaveRequestTracingService) UpdateEmployeeLeaveRequestTracing(employeeLeaveRequestTracing model.EmployeeLeaveRequestTracing, id uint) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error) {
	newEmployeeLeaveRequestTracing := model.EmployeeLeaveRequestTracing{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequestTracing, smapping.MapFields(&employeeLeaveRequestTracing))

	if err1 != nil {
		return newEmployeeLeaveRequestTracing, err1
	}

	return service.employeeLeaveRequestTracingRepository.UpdateEmployeeLeaveRequestTracing(newEmployeeLeaveRequestTracing, id)
}

func (service *employeeLeaveRequestTracingService) DeleteEmployeeLeaveRequestTracing(employeeLeaveRequestTracing model.EmployeeLeaveRequestTracing, id uint) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error) {
	newEmployeeLeaveRequestTracing := model.EmployeeLeaveRequestTracing{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequestTracing, smapping.MapFields(&employeeLeaveRequestTracing))

	if err1 != nil {
		return newEmployeeLeaveRequestTracing, err1
	}

	return service.employeeLeaveRequestTracingRepository.UpdateEmployeeLeaveRequestTracing(newEmployeeLeaveRequestTracing, id)
}
