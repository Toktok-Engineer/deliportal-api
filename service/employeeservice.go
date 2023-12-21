package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type EmployeeService interface {
	CountEmployeeAll() (count int64, err error)
	FindEmployees() (employeeOutput []model.SelectEmployeeParameter, err error)
	FindEmployeesOffset(limit int, offset int, order string, dir string) (employeeOutput []model.SelectEmployeeParameter, err error)
	SearchEmployee(limit int, offset int, order string, dir string, search string) (employeeOutput []model.SelectEmployeeParameter, err error)
	CountSearchEmployee(search string) (count int64, err error)
	FindEmployeeById(id uint) (employeeOutput model.SelectEmployeeParameter, err error)
	FindEmployeeByNik(nik string) (employeeOutput model.SelectEmployeeParameter, err error)
	FindExcEmployee(id uint) (employeeOutput []model.SelectEmployeeParameter, err error)
	FindEmployeeByDepartment(department int) (employeeOutput []model.SelectEmployeeParameter, err error)
	FindEmployeeByPosition(group int, division string, department string, position string) (employeeOutput model.SelectEmployeeParameter, err error)
	FindEmployeeByDivIdAndDepId(Divid uint, DepId uint) (employeeOutput []model.SelectEmployeeParameter, err error)
	FindEmployeeByDate(date float64) (employeeOutput []model.SelectEmployeeParameter, err error)
	InsertEmployee(employee model.CreateEmployeeParameter) (employeeOutput model.Employee, err error)
	FindEmployeeCuti(group int, section int, division int, department int, position int) (employeeOutput []model.SelectEmployeeCuti, err error)
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

func (service *employeeService) CountEmployeeAll() (count int64, err error) {
	res, err := service.employeeRepository.CountEmployeeAll()
	return res, err
}

func (service *employeeService) FindEmployees() (employeeOutput []model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployees()
	return res, err
}

func (service *employeeService) FindEmployeesOffset(limit int, offset int, order string, dir string) (employeeOutput []model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployeesOffset(limit, offset, order, dir)
	return res, err
}

func (service *employeeService) SearchEmployee(limit int, offset int, order string, dir string, search string) (employeeOutput []model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.SearchEmployee(limit, offset, order, dir, search)
	return res, err
}

func (service *employeeService) CountSearchEmployee(search string) (count int64, err error) {
	res, err := service.employeeRepository.CountSearchEmployee(search)
	return res, err
}

func (service *employeeService) FindEmployeeById(id uint) (employeeOutput model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployeeById(id)
	return res, err
}

func (service *employeeService) FindEmployeeByNik(nik string) (employeeOutput model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployeeByNik(nik)
	return res, err
}

func (service *employeeService) FindExcEmployee(id uint) (employeeOutput []model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindExcEmployee(id)
	return res, err
}

func (service *employeeService) FindEmployeeByDepartment(department int) (employeeOutput []model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployeeByDepartment(department)
	return res, err
}

func (service *employeeService) FindEmployeeByPosition(group int, division string, department string, position string) (employeeOutput model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployeeByPosition(group, division, department, position)
	return res, err
}

func (service *employeeService) FindEmployeeByDivIdAndDepId(Divid uint, DepId uint) (employeeOutput []model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployeeByDivIdAndDepId(Divid, DepId)
	return res, err
}

func (service *employeeService) FindEmployeeCuti(group int, section int, division int, department int, position int) (employeeOutput []model.SelectEmployeeCuti, err error) {
	res, err := service.employeeRepository.FindEmployeeCuti(group, section, division, department, position)
	return res, err
}

func (service *employeeService) FindEmployeeByDate(date float64) (employeeOutput []model.SelectEmployeeParameter, err error) {
	res, err := service.employeeRepository.FindEmployeeByDate(date)
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
