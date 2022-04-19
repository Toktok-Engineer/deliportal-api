package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	FindDepartments() (departmentOutput []model.SelectDepartmentParameter, err error)
	FindDepartmentById(id uint) (departmentOutput model.SelectDepartmentParameter, err error)
	FindExcDepartment(divId uint, id uint) (departmentOutput []model.SelectDepartmentParameter, err error)
	FindDepartmentByDivId(divId uint) (departmentOutput []model.SelectDepartmentParameter, err error)
	InsertDepartment(department model.Department) (departmentOutput model.Department, err error)
	UpdateDepartment(department model.Department, id uint) (departmentOutput model.Department, err error)
}

type DepartmentConnection struct {
	connection *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &DepartmentConnection{
		connection: db,
	}
}

func (db *DepartmentConnection) FindDepartments() (departmentOutput []model.SelectDepartmentParameter, err error) {
	var (
		departments []model.SelectDepartmentParameter
	)

	res := db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("departments.deleted_at = 0").Order("departments.department_name").Find(&departments)
	return departments, res.Error
}

func (db *DepartmentConnection) FindDepartmentById(id uint) (departmentOutput model.SelectDepartmentParameter, err error) {
	var (
		department model.SelectDepartmentParameter
	)

	res := db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("departments.id=? AND departments.deleted_at = 0", id).Take(&department)
	return department, res.Error
}

func (db *DepartmentConnection) FindExcDepartment(divId uint, id uint) (departmentOutput []model.SelectDepartmentParameter, err error) {
	var (
		departments []model.SelectDepartmentParameter
	)

	res := db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("departments.division_id = ? AND departments.id!=? AND departments.deleted_at = 0", divId, id).Order("departments.department_name").Find(&departments)
	return departments, res.Error
}

func (db *DepartmentConnection) FindDepartmentByDivId(divId uint) (departmentOutput []model.SelectDepartmentParameter, err error) {
	var (
		departments []model.SelectDepartmentParameter
	)

	res := db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("departments.division_id=? AND departments.deleted_at = 0", divId).Order("departments.department_name").Find(&departments)
	return departments, res.Error
}

func (db *DepartmentConnection) InsertDepartment(department model.Department) (departmentOutput model.Department, err error) {
	res := db.connection.Save(&department)
	return department, res.Error
}

func (db *DepartmentConnection) UpdateDepartment(department model.Department, id uint) (departmentOutput model.Department, err error) {
	res := db.connection.Where("id=?", id).Updates(&department)
	return department, res.Error
}
