package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	FindDepartments() []model.SelectDepartmentParameter
	FindDepartmentById(id uint) model.SelectDepartmentParameter
	FindExcDepartment(divId uint, id uint) []model.SelectDepartmentParameter
	FindDepartmentByDivId(divId uint) []model.SelectDepartmentParameter
	InsertDepartment(department model.Department) model.Department
	UpdateDepartment(department model.Department, id uint) model.Department
}

type DepartmentConnection struct {
	connection *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &DepartmentConnection{
		connection: db,
	}
}

func (db *DepartmentConnection) FindDepartments() []model.SelectDepartmentParameter {
	var departments []model.SelectDepartmentParameter

	db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("departments.deleted_at = 0").Order("departments.department_name").Find(&departments)

	return departments
}

func (db *DepartmentConnection) FindDepartmentById(id uint) model.SelectDepartmentParameter {
	var department model.SelectDepartmentParameter

	db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("departments.id=? AND departments.deleted_at = 0", id).Take(&department)

	return department
}

func (db *DepartmentConnection) FindExcDepartment(divId uint, id uint) []model.SelectDepartmentParameter {
	var departments []model.SelectDepartmentParameter

	db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("departments.division_id = ? AND departments.id!=? AND departments.deleted_at = 0", divId, id).Order("departments.department_name").Find(&departments)

	return departments
}

func (db *DepartmentConnection) FindDepartmentByDivId(divId uint) []model.SelectDepartmentParameter {
	var departments []model.SelectDepartmentParameter

	db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("departments.division_id=? AND departments.deleted_at = 0", divId).Order("departments.department_name").Find(&departments)

	return departments
}

func (db *DepartmentConnection) InsertDepartment(department model.Department) model.Department {
	db.connection.Save(&department)
	return department
}

func (db *DepartmentConnection) UpdateDepartment(department model.Department, id uint) model.Department {
	db.connection.Where("id=?", id).Updates(&department)
	return department
}
