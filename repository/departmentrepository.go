package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	CountDepartmentAll() (count int64, err error)
	FindDepartments() (departmentOutput []model.SelectDepartmentParameter, err error)
	FindDepartmentsOffset(limit int, offset int, order string, dir string) (departmentOutput []model.SelectDepartmentParameter, err error)
	SearchDepartment(limit int, offset int, order string, dir string, search string) (departmentOutput []model.SelectDepartmentParameter, err error)
	CountSearchDepartment(search string) (count int64, err error)
	FindDepartmentById(id uint) (departmentOutput model.SelectDepartmentParameter, err error)
	FindExcDepartment(divId uint, id uint) (departmentOutput []model.SelectDepartmentParameter, err error)
	FindDepartmentByDivId(divId uint) (departmentOutput []model.SelectDepartmentParameter, err error)
	CountDepartmentName(search string) (count int64, err error)
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

func (db *DepartmentConnection) CountDepartmentAll() (count int64, err error) {
	res := db.connection.Debug().Table("departments").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *DepartmentConnection) FindDepartments() (departmentOutput []model.SelectDepartmentParameter, err error) {
	var (
		departments []model.SelectDepartmentParameter
	)
	res := db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("departments.deleted_at = 0").Order("departments.department_name").Find(&departments)
	return departments, res.Error
}

func (db *DepartmentConnection) FindDepartmentsOffset(limit int, offset int, order string, dir string) (departmentOutput []model.SelectDepartmentParameter, err error) {
	var (
		orderDirection string
		departments    []model.SelectDepartmentParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("departments.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&departments)
	return departments, res.Error
}

func (db *DepartmentConnection) SearchDepartment(limit int, offset int, order string, dir string, search string) (departmentOutput []model.SelectDepartmentParameter, err error) {
	var (
		orderDirection string
		final          string
		departments    []model.SelectDepartmentParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("(lower(divisions.division_name) LIKE ? OR lower(departments.department_name) LIKE ? OR lower(departments.remark) LIKE ?) AND departments.deleted_at = 0", final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&departments)
	return departments, res.Error
}

func (db *DepartmentConnection) CountSearchDepartment(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("(lower(divisions.division_name) LIKE ? OR lower(departments.department_name) LIKE ? OR lower(departments.remark) LIKE ?) AND departments.deleted_at = 0", final, final, final).Count(&count)
	return count, res.Error
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

func (db *DepartmentConnection) CountDepartmentName(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("departments").Select("departments.id, departments.department_name, departments.division_id, divisions.division_name, departments.remark, departments.created_user_id, departments.updated_user_id, departments.deleted_user_id, departments.created_at, departments.updated_at, departments.deleted_at").Joins("left join divisions ON departments.division_id = divisions.id").Where("lower(departments.department_name) LIKE ?", final).Count(&count)
	return count, res.Error
}

func (db *DepartmentConnection) InsertDepartment(department model.Department) (departmentOutput model.Department, err error) {
	res := db.connection.Save(&department)
	return department, res.Error
}

func (db *DepartmentConnection) UpdateDepartment(department model.Department, id uint) (departmentOutput model.Department, err error) {
	res := db.connection.Where("id=?", id).Updates(&department)
	return department, res.Error
}
