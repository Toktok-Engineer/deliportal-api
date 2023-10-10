package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type EmployeeLeaveCreditService interface {
	CountEmployeeLeaveCreditAll(year int) (count int64, err error)
	FindEmployeeLeaveCredits(year int, empId int) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	FindEmployeeLeaveCreditsOffset(limit int, offset int, order string, dir string, year int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountSearchEmployeeLeaveCreditAll(year int, search string) (count int64, err error)
	FindSearchEmployeeLeaveCreditsOffset(limit int, offset int, order string, dir string, year int, search string) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountEmployeeLeaveCreditEmpID(year int, empId int) (count int64, err error)
	FindEmployeeLeaveCreditsOffsetByYear(limit int, offset int, order string, dir string, year int, empId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountSearchEmployeeLeaveCreditByEmp(year int, empId int, search string) (count int64, err error)
	FindSearchEmployeeLeaveCreditsOffsetByEmp(limit int, offset int, order string, dir string, year int, empId int, search string) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountSearchEmployeeLeaveCreditByDepId(year int, deptId int, positionId int, search string, groupId int) (count int64, err error)
	FindSearchEmployeeLeaveCreditsOffsetByDept(limit int, offset int, order string, dir string, year int, deptId int, positionId int, search string, groupId int, employeeid int, sectionid int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountEmployeeLeaveCreditByDepId(year int, deptId int, positionId int, groupId int) (count int64, err error)
	FindEmployeeLeaveCreditsOffsetByDept(limit int, offset int, order string, dir string, year int, deptId int, positionId int, groupId int, employeeid int, sectionid int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountEmployeeLeaveCreditByDivAll(year int, divId int, positionId int, groupId int) (count int64, err error)
	FindEmployeeLeaveCreditsOffsetByDivAll(limit int, offset int, order string, dir string, year int, divId int, positionId int, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountSearchEmployeeLeaveCreditByDivIdAll(year int, divId int, positionId int, search string, groupId int) (count int64, err error)
	FindSearchEmployeeLeaveCreditsOffsetByDivIdAll(limit int, offset int, order string, dir string, year int, divId int, positionId int, search string, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	FindEmployeeLeaveCreditsExcEmp(year int, empId string) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	FindEmployeeLeaveCreditsAll(year int) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	CountEmployeeLeaveCreditByDepAll(year int, depId int, positionId int, groupId int) (count int64, err error)
	FindEmployeeLeaveCreditsOffsetByDepAll(limit int, offset int, order string, dir string, year int, depId int, positionId int, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountSearchEmployeeLeaveCreditByDepIdAll(year int, depId int, positionId int, search string, groupId int) (count int64, err error)
	FindSearchEmployeeLeaveCreditsOffsetByDepIdAll(limit int, offset int, order string, dir string, year int, depId int, positionId int, search string, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	// SearchEmployeeLeaveCredit(limit int, offset int, order string, dir string, search string) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	// CountSearchEmployeeLeaveCredit(search string) (count int64, err error)
	FindEmployeeLeaveCreditById(id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error)
	FindExcEmployeeLeaveCredit(id uint) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	// CountEmployeeLeaveCreditName(search string) (count int64, err error)
	FindEmployeeLeaveCreditByEmpId(empid uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error)
	FindAllEmployeeLeaveCreditByEmpId(empid uint) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	InsertEmployeeLeaveCredit(employeeLeaveCredit model.CreateEmployeeLeaveCreditParameter) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error)
	UpdateEmployeeLeaveCredit(employeeLeaveCredit model.EmployeeLeaveCredit, id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error)
	DeleteEmployeeLeaveCredit(employeeLeaveCredit model.EmployeeLeaveCredit, id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error)
}

type employeeLeaveCreditService struct {
	employeeLeaveCreditRepository repository.EmployeeLeaveCreditRepository
}

func NewEmployeeLeaveCreditService(employeeLeaveCreditRep repository.EmployeeLeaveCreditRepository) EmployeeLeaveCreditService {
	return &employeeLeaveCreditService{
		employeeLeaveCreditRepository: employeeLeaveCreditRep,
	}
}

func (service *employeeLeaveCreditService) CountEmployeeLeaveCreditAll(year int) (count int64, err error) {
	res, err := service.employeeLeaveCreditRepository.CountEmployeeLeaveCreditAll(year)
	return res, err
}

func (service *employeeLeaveCreditService) FindEmployeeLeaveCredits(year int, empId int) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindEmployeeLeaveCredits(year, empId)
	return res, err
}

func (service *employeeLeaveCreditService) FindEmployeeLeaveCreditsOffset(limit int, offset int, order string, dir string, year int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindEmployeeLeaveCreditsOffset(limit, offset, order, dir, year)
	return res, err
}

func (service *employeeLeaveCreditService) CountSearchEmployeeLeaveCreditAll(year int, search string) (count int64, err error) {
	res, err := service.employeeLeaveCreditRepository.CountSearchEmployeeLeaveCreditAll(year, search)
	return res, err
}

func (service *employeeLeaveCreditService) FindSearchEmployeeLeaveCreditsOffset(limit int, offset int, order string, dir string, year int, search string) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindSearchEmployeeLeaveCreditsOffset(limit, offset, order, dir, year, search)
	return res, err
}

func (service *employeeLeaveCreditService) CountEmployeeLeaveCreditEmpID(year int, empId int) (count int64, err error) {
	res, err := service.employeeLeaveCreditRepository.CountEmployeeLeaveCreditEmpID(year, empId)
	return res, err
}
func (service *employeeLeaveCreditService) FindEmployeeLeaveCreditsOffsetByYear(limit int, offset int, order string, dir string, year int, empId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindEmployeeLeaveCreditsOffsetByYear(limit, offset, order, dir, year, empId)
	return res, err
}

func (service *employeeLeaveCreditService) CountSearchEmployeeLeaveCreditByEmp(year int, empId int, search string) (count int64, err error) {
	res, err := service.employeeLeaveCreditRepository.CountSearchEmployeeLeaveCreditByEmp(year, empId, search)
	return res, err
}

func (service *employeeLeaveCreditService) FindSearchEmployeeLeaveCreditsOffsetByEmp(limit int, offset int, order string, dir string, year int, empId int, search string) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindSearchEmployeeLeaveCreditsOffsetByEmp(limit, offset, order, dir, year, empId, search)
	return res, err
}

func (service *employeeLeaveCreditService) CountSearchEmployeeLeaveCreditByDepId(year int, deptId int, positionId int, search string, groupId int) (count int64, err error) {
	res, err := service.employeeLeaveCreditRepository.CountSearchEmployeeLeaveCreditByDepId(year, deptId, positionId, search, groupId)
	return res, err
}

func (service *employeeLeaveCreditService) FindSearchEmployeeLeaveCreditsOffsetByDept(limit int, offset int, order string, dir string, year int, deptId int, positionId int, search string, groupId int, employeeid int, sectionid int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindSearchEmployeeLeaveCreditsOffsetByDept(limit, offset, order, dir, year, deptId, positionId, search, groupId, employeeid, sectionid)
	return res, err
}

func (service *employeeLeaveCreditService) CountEmployeeLeaveCreditByDepId(year int, deptId int, positionId int, groupId int) (count int64, err error) {
	res, err := service.employeeLeaveCreditRepository.CountEmployeeLeaveCreditByDepId(year, deptId, positionId, groupId)
	return res, err
}

func (service *employeeLeaveCreditService) FindEmployeeLeaveCreditsOffsetByDept(limit int, offset int, order string, dir string, year int, deptId int, positionId int, groupId int, employeeid int, sectionid int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindEmployeeLeaveCreditsOffsetByDept(limit, offset, order, dir, year, deptId, positionId, groupId, employeeid, sectionid)
	return res, err
}

func (service *employeeLeaveCreditService) CountEmployeeLeaveCreditByDivAll(year int, divId int, positionId int, groupId int) (count int64, err error) {
	res, err := service.employeeLeaveCreditRepository.CountEmployeeLeaveCreditByDivAll(year, divId, positionId, groupId)
	return res, err
}

func (service *employeeLeaveCreditService) FindEmployeeLeaveCreditsOffsetByDivAll(limit int, offset int, order string, dir string, year int, divId int, positionId int, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindEmployeeLeaveCreditsOffsetByDivAll(limit, offset, order, dir, year, divId, positionId, groupId)
	return res, err
}

func (service *employeeLeaveCreditService) CountSearchEmployeeLeaveCreditByDivIdAll(year int, divId int, positionId int, search string, groupId int) (count int64, err error) {
	res, err := service.employeeLeaveCreditRepository.CountSearchEmployeeLeaveCreditByDivIdAll(year, divId, positionId, search, groupId)
	return res, err
}

func (service *employeeLeaveCreditService) FindSearchEmployeeLeaveCreditsOffsetByDivIdAll(limit int, offset int, order string, dir string, year int, divId int, positionId int, search string, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindSearchEmployeeLeaveCreditsOffsetByDivIdAll(limit, offset, order, dir, year, divId, positionId, search, groupId)
	return res, err
}

func (service *employeeLeaveCreditService) FindEmployeeLeaveCreditsExcEmp(year int, empId string) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindEmployeeLeaveCreditsExcEmp(year, empId)
	return res, err
}
func (service *employeeLeaveCreditService) FindEmployeeLeaveCreditsAll(year int) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindEmployeeLeaveCreditsAll(year)
	return res, err
}

func (service *employeeLeaveCreditService) CountEmployeeLeaveCreditByDepAll(year int, depId int, positionId int, groupId int) (count int64, err error) {
	res, err := service.employeeLeaveCreditRepository.CountEmployeeLeaveCreditByDepAll(year, depId, positionId, groupId)
	return res, err
}

func (service *employeeLeaveCreditService) FindEmployeeLeaveCreditsOffsetByDepAll(limit int, offset int, order string, dir string, year int, depId int, positionId int, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindEmployeeLeaveCreditsOffsetByDepAll(limit, offset, order, dir, year, depId, positionId, groupId)
	return res, err
}

func (service *employeeLeaveCreditService) CountSearchEmployeeLeaveCreditByDepIdAll(year int, depId int, positionId int, search string, groupId int) (count int64, err error) {
	res, err := service.employeeLeaveCreditRepository.CountSearchEmployeeLeaveCreditByDepIdAll(year, depId, positionId, search, groupId)
	return res, err
}

func (service *employeeLeaveCreditService) FindSearchEmployeeLeaveCreditsOffsetByDepIdAll(limit int, offset int, order string, dir string, year int, depId int, positionId int, search string, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindSearchEmployeeLeaveCreditsOffsetByDepIdAll(limit, offset, order, dir, year, depId, positionId, search, groupId)
	return res, err
}

// func (service *employeeLeaveCreditService) SearchEmployeeLeaveCredit(limit int, offset int, order string, dir string, search string) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
// 	res, err := service.employeeLeaveCreditRepository.SearchEmployeeLeaveCredit(limit, offset, order, dir, search)
// 	return res, err
// }

// func (service *employeeLeaveCreditService) CountSearchEmployeeLeaveCredit(search string) (count int64, err error) {
// 	res, err := service.employeeLeaveCreditRepository.CountSearchEmployeeLeaveCredit(search)
// 	return res, err
// }

func (service *employeeLeaveCreditService) FindEmployeeLeaveCreditById(id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindEmployeeLeaveCreditById(id)
	return res, err
}

// func (service *employeeLeaveCreditService) CountEmployeeLeaveCreditName(search string) (count int64, err error) {
// 	res, err := service.employeeLeaveCreditRepository.CountEmployeeLeaveCreditName(search)
// 	return res, err
// }

func (service *employeeLeaveCreditService) FindEmployeeLeaveCreditByEmpId(empid uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindEmployeeLeaveCreditByEmpId(empid)
	return res, err
}

func (service *employeeLeaveCreditService) FindAllEmployeeLeaveCreditByEmpId(empid uint) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindAllEmployeeLeaveCreditByEmpId(empid)
	return res, err
}

func (service *employeeLeaveCreditService) FindExcEmployeeLeaveCredit(id uint) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
	res, err := service.employeeLeaveCreditRepository.FindExcEmployeeLeaveCredit(id)
	return res, err
}

func (service *employeeLeaveCreditService) InsertEmployeeLeaveCredit(employeeLeaveCredit model.CreateEmployeeLeaveCreditParameter) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error) {
	newEmployeeLeaveCredit := model.EmployeeLeaveCredit{}
	err1 := smapping.FillStruct(&newEmployeeLeaveCredit, smapping.MapFields(&employeeLeaveCredit))
	if err != nil {
		return newEmployeeLeaveCredit, err1
	}
	res, err := service.employeeLeaveCreditRepository.InsertEmployeeLeaveCredit(newEmployeeLeaveCredit)
	return res, err
}

func (service *employeeLeaveCreditService) UpdateEmployeeLeaveCredit(employeeLeaveCredit model.EmployeeLeaveCredit, id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error) {
	newEmployeeLeaveCredit := model.EmployeeLeaveCredit{}
	err1 := smapping.FillStruct(&newEmployeeLeaveCredit, smapping.MapFields(&employeeLeaveCredit))
	if err != nil {
		return newEmployeeLeaveCredit, err1
	}
	res, err := service.employeeLeaveCreditRepository.UpdateEmployeeLeaveCredit(newEmployeeLeaveCredit, id)
	return res, err
}

func (service *employeeLeaveCreditService) DeleteEmployeeLeaveCredit(employeeLeaveCredit model.EmployeeLeaveCredit, id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error) {
	newEmployeeLeaveCredit := model.EmployeeLeaveCredit{}
	err1 := smapping.FillStruct(&newEmployeeLeaveCredit, smapping.MapFields(&employeeLeaveCredit))
	if err != nil {
		return newEmployeeLeaveCredit, err1
	}
	res, err := service.employeeLeaveCreditRepository.DeleteEmployeeLeaveCredit(newEmployeeLeaveCredit, id)
	return res, err
}
