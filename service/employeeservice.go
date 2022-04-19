package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type EmployeeService interface {
	FindEmployees() (employeeOutput []model.SelectEmployeeParameter, err error)
	FindEmployeeById(id uint) (employeeOutput model.SelectEmployeeParameter, err error)
	FindEmployeeByNik(nik uint) (employeeOutput model.SelectEmployeeParameter, err error)
	FindExcEmployee(id uint) (employeeOutput []model.SelectEmployeeParameter, err error)
	InsertEmployee(employee model.CreateEmployeeParameter) (employeeOutput model.Employee, err error)
	UpdateEmployee(employee model.Employee, id uint) (employeeOutput model.Employee, err error)
	DeleteEmployee(employee model.Employee, id uint) (employeeOutput model.Employee, err error)
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRep repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		employeeRepository: employeeRep,
	}
}

func (service *employeeService) FindEmployees() (employeeOutput []model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployees()
	return res, err
}

func (service *employeeService) FindEmployeeById(id uint) (employeeOutput model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployeeById(id)
	return res, err
}

func (service *employeeService) FindEmployeeByNik(nik uint) (employeeOutput model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployeeByNik(nik)
	return res, err
}

func (service *employeeService) FindExcEmployee(id uint) (employeeOutput []model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindExcEmployee(id)
	return res, err
}

func (service *employeeService) InsertEmployee(employee model.CreateEmployeeParameter) (employeeOutput model.Employee, err error) {
	newEmployee := model.Employee{}
	err1 := smapping.FillStruct(&newEmployee, smapping.MapFields(&employee))
	if err != nil {
		return newEmployee, err1
	}
	res, err := service.employeeRepository.InsertEmployee(newEmployee)
	return res, err
}

func (service *employeeService) UpdateEmployee(employee model.Employee, id uint) (employeeOutput model.Employee, err error) {
	newEmployee := model.Employee{}
	err1 := smapping.FillStruct(&newEmployee, smapping.MapFields(&employee))
	if err != nil {
		return newEmployee, err1
	}
	res, err := service.employeeRepository.UpdateEmployee(newEmployee, id)
	return res, err
}

func (service *employeeService) DeleteEmployee(employee model.Employee, id uint) (employeeOutput model.Employee, err error) {
	newEmployee := model.Employee{}
	err1 := smapping.FillStruct(&newEmployee, smapping.MapFields(&employee))
	if err != nil {
		return newEmployee, err1
	}
	res, err := service.employeeRepository.UpdateEmployee(newEmployee, id)
	return res, err
}
