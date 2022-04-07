package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type DepartmentService interface {
	FindDepartments() (departmentOutput []model.SelectDepartmentParameter, err error)
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

func (service *departmentService) FindDepartments() (departmentOutput []model.SelectDepartmentParameter, err error) {
	res, err := service.departmentRepository.FindDepartments()
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
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.departmentRepository.InsertDepartment(newDepartment)
	return res, err
}

func (service *departmentService) UpdateDepartment(department model.Department, id uint) (departmentOutput model.Department, err error) {
	newDepartment := model.Department{}
	err1 := smapping.FillStruct(&newDepartment, smapping.MapFields(&department))
	if err != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.departmentRepository.UpdateDepartment(newDepartment, id)
	return res, err
}

func (service *departmentService) DeleteDepartment(department model.Department, id uint) (departmentOutput model.Department, err error) {
	newDepartment := model.Department{}
	err1 := smapping.FillStruct(&newDepartment, smapping.MapFields(&department))
	if err1 != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.departmentRepository.UpdateDepartment(newDepartment, id)
	return res, err
}
