package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type DepartmentService interface {
	CountDepartmentAll() (count int64, err error)
	FindDepartments() (departmentOutput []model.SelectDepartmentParameter, err error)
	FindDepartmentsOffset(limit int, offset int, order string, dir string) (departmentOutput []model.SelectDepartmentParameter, err error)
	SearchDepartment(limit int, offset int, order string, dir string, search string) (departmentOutput []model.SelectDepartmentParameter, err error)
	CountSearchDepartment(search string) (count int64, err error)
	FindDepartmentById(id uint) (departmentOutput model.SelectDepartmentParameter, err error)
	FindExcDepartment(divId uint, id uint) (departmentOutput []model.SelectDepartmentParameter, err error)
	FindDepartmentByDivId(divId uint) (departmentOutput []model.SelectDepartmentParameter, err error)
	InsertDepartment(department model.CreateDepartmentParameter) (departmentOutput model.Department, err error)
	UpdateDepartment(department model.Department, id uint) (departmentOutput model.Department, err error)
	DeleteDepartment(department model.Department, id uint) (departmentOutput model.Department, err error)
}

type departmentService struct {
	departmentRepository repository.DepartmentRepository
}

func NewDepartmentService(departmentRep repository.DepartmentRepository) DepartmentService {
	return &departmentService{
		departmentRepository: departmentRep,
	}
}

func (service *departmentService) CountDepartmentAll() (count int64, err error) {
	res, err := service.departmentRepository.CountDepartmentAll()
	return res, err
}

func (service *departmentService) FindDepartments() (departmentOutput []model.SelectDepartmentParameter, err error) {
	res, err := service.departmentRepository.FindDepartments()
	return res, err
}

func (service *departmentService) FindDepartmentsOffset(limit int, offset int, order string, dir string) (departmentOutput []model.SelectDepartmentParameter, err error) {
	res, err := service.departmentRepository.FindDepartmentsOffset(limit, offset, order, dir)
	return res, err
}

func (service *departmentService) SearchDepartment(limit int, offset int, order string, dir string, search string) (departmentOutput []model.SelectDepartmentParameter, err error) {
	res, err := service.departmentRepository.SearchDepartment(limit, offset, order, dir, search)
	return res, err
}

func (service *departmentService) CountSearchDepartment(search string) (count int64, err error) {
	res, err := service.departmentRepository.CountSearchDepartment(search)
	return res, err
}

func (service *departmentService) FindDepartmentById(id uint) (departmentOutput model.SelectDepartmentParameter, err error) {
	res, err := service.departmentRepository.FindDepartmentById(id)
	return res, err
}

func (service *departmentService) FindExcDepartment(divId uint, id uint) (departmentOutput []model.SelectDepartmentParameter, err error) {
	res, err := service.departmentRepository.FindExcDepartment(divId, id)
	return res, err
}

func (service *departmentService) FindDepartmentByDivId(divId uint) (departmentOutput []model.SelectDepartmentParameter, err error) {
	res, err := service.departmentRepository.FindDepartmentByDivId(divId)
	return res, err
}

func (service *departmentService) InsertDepartment(department model.CreateDepartmentParameter) (departmentOutput model.Department, err error) {
	newDepartment := model.Department{}
	err1 := smapping.FillStruct(&newDepartment, smapping.MapFields(&department))
	if err != nil {
		return newDepartment, err1
	}
	res, err := service.departmentRepository.InsertDepartment(newDepartment)
	return res, err
}

func (service *departmentService) UpdateDepartment(department model.Department, id uint) (departmentOutput model.Department, err error) {
	newDepartment := model.Department{}
	err1 := smapping.FillStruct(&newDepartment, smapping.MapFields(&department))
	if err != nil {
		return newDepartment, err1
	}
	res, err := service.departmentRepository.UpdateDepartment(newDepartment, id)
	return res, err
}

func (service *departmentService) DeleteDepartment(department model.Department, id uint) (departmentOutput model.Department, err error) {
	newDepartment := model.Department{}
	err1 := smapping.FillStruct(&newDepartment, smapping.MapFields(&department))
	if err != nil {
		return newDepartment, err1
	}
	res, err := service.departmentRepository.UpdateDepartment(newDepartment, id)
	return res, err
}
