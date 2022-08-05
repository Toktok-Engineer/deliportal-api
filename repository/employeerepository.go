package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CountEmployeeAll() (count int64, err error)
	FindEmployees() (employeeOutput []model.SelectEmployeeParameter, err error)
	FindEmployeesOffset(limit int, offset int, order string, dir string) (employeeOutput []model.SelectEmployeeParameter, err error)
	SearchEmployee(limit int, offset int, order string, dir string, search string) (employeeOutput []model.SelectEmployeeParameter, err error)
	CountSearchEmployee(search string) (count int64, err error)
	FindEmployeeById(id uint) (employeeOutput model.SelectEmployeeParameter, err error)
	FindEmployeeByNik(nik uint) (employeeOutput model.SelectEmployeeParameter, err error)
	FindExcEmployee(id uint) (employeeOutput []model.SelectEmployeeParameter, err error)
	InsertEmployee(employee model.Employee) (employeeOutput model.Employee, err error)
	UpdateEmployee(employee model.Employee, id uint) (employeeOutput model.Employee, err error)
}

type EmployeeConnection struct {
	connection *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &EmployeeConnection{
		connection: db,
	}
}

func (db *EmployeeConnection) CountEmployeeAll() (count int64, err error) {
	res := db.connection.Debug().Table("employees").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *EmployeeConnection) FindEmployees() (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		employees []model.SelectEmployeeParameter
	)
	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.deleted_at = 0").Order("employees.firstname").Find(&employees)
	return employees, res.Error
}

func (db *EmployeeConnection) FindEmployeesOffset(limit int, offset int, order string, dir string) (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		orderDirection string
		employees      []model.SelectEmployeeParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&employees)
	return employees, res.Error
}

func (db *EmployeeConnection) SearchEmployee(limit int, offset int, order string, dir string, search string) (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		orderDirection string
		final          string
		employees      []model.SelectEmployeeParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("(lower(employees.nik) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(divisions.division_name) LIKE ? OR lower(departments.department_name) LIKE ? OR lower(sections.section_name) LIKE ? OR lower(positions.position_name) LIKE ? OR lower(locations.location_name) LIKE ? OR lower(employees.email) LIKE ? OR lower(employees.remark) LIKE ?) AND departments.deleted_at = 0", final, final, final, final, final, final, final, final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&employees)
	return employees, res.Error
}

func (db *EmployeeConnection) CountSearchEmployee(search string) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("(lower(employees.nik) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(divisions.division_name) LIKE ? OR lower(departments.department_name) LIKE ? OR lower(sections.section_name) LIKE ? OR lower(positions.position_name) LIKE ? OR lower(locations.location_name) LIKE ? OR lower(employees.email) LIKE ? OR lower(employees.remark) LIKE ?) AND departments.deleted_at = 0", final, final, final, final, final, final, final, final, final, final).Count(&count)
	return count, res.Error
}

func (db *EmployeeConnection) FindEmployeeById(id uint) (employeeOutput model.SelectEmployeeParameter, err error) {
	var (
		employee model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.id=? AND employees.deleted_at = 0", id).Take(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) FindEmployeeByNik(nik uint) (employeeOutput model.SelectEmployeeParameter, err error) {
	var (
		employee model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.nik=? AND employees.deleted_at = 0", nik).Take(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) FindExcEmployee(id uint) (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		employees []model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.id != ? AND employees.deleted_at = 0", id).Find(&employees)
	return employees, res.Error
}

func (db *EmployeeConnection) InsertEmployee(employee model.Employee) (employeeOutput model.Employee, err error) {
	res := db.connection.Save(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) UpdateEmployee(employee model.Employee, id uint) (employeeOutput model.Employee, err error) {
	res := db.connection.Where("id=?", id).Updates(&employee)
	return employee, res.Error
}
