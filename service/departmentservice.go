package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type DepartmentService interface {
	FindDepartments() []model.SelectDepartmentParameter
	FindDepartmentById(id uint) model.SelectDepartmentParameter
	FindExcDepartment(divId uint, id uint) []model.SelectDepartmentParameter
	FindDepartmentByDivId(divId uint) []model.SelectDepartmentParameter
	InsertDepartment(department model.CreateDepartmentParameter) model.Department
	UpdateDepartment(department model.Department, id uint) model.Department
	DeleteDepartment(department model.Department, id uint) model.Department
}

type departmentService struct {
	departmentRepository repository.DepartmentRepository
}

func NewDepartmentService(departmentRep repository.DepartmentRepository) DepartmentService {
	return &departmentService{
		departmentRepository: departmentRep,
	}
}

func (service *departmentService) FindDepartments() []model.SelectDepartmentParameter {
	return service.departmentRepository.FindDepartments()
}

func (service *departmentService) FindDepartmentById(id uint) model.SelectDepartmentParameter {
	return service.departmentRepository.FindDepartmentById(id)
}

func (service *departmentService) FindExcDepartment(divId uint, id uint) []model.SelectDepartmentParameter {
	return service.departmentRepository.FindExcDepartment(divId, id)
}

func (service *departmentService) FindDepartmentByDivId(divId uint) []model.SelectDepartmentParameter {
	return service.departmentRepository.FindDepartmentByDivId(divId)
}

func (service *departmentService) InsertDepartment(department model.CreateDepartmentParameter) model.Department {
	newDepartment := model.Department{}
	err := smapping.FillStruct(&newDepartment, smapping.MapFields(&department))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.departmentRepository.InsertDepartment(newDepartment)
	return res
}

func (service *departmentService) UpdateDepartment(department model.Department, id uint) model.Department {
	newDepartment := model.Department{}
	err := smapping.FillStruct(&newDepartment, smapping.MapFields(&department))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.departmentRepository.UpdateDepartment(newDepartment, id)
	return res
}

func (service *departmentService) DeleteDepartment(department model.Department, id uint) model.Department {
	newDepartment := model.Department{}
	err := smapping.FillStruct(&newDepartment, smapping.MapFields(&department))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.departmentRepository.UpdateDepartment(newDepartment, id)
	return res
}
