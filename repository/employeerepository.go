package repository

import (
	"deliportal-api/model"
	"log"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindEmployees() (employeeOutput []model.SelectEmployeeParameter, err error)
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

func (db *EmployeeConnection) FindEmployees() (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		employees []model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.deleted_at = 0").Order("employees.firstname").Find(&employees)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return employees, res.Error
	}
	return employees, nil
}

func (db *EmployeeConnection) FindEmployeeById(id uint) (employeeOutput model.SelectEmployeeParameter, err error) {
	var (
		employee model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.id=? AND employees.deleted_at = 0", id).Take(&employee)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return employee, res.Error
	}
	return employee, nil
}

func (db *EmployeeConnection) FindEmployeeByNik(nik uint) (employeeOutput model.SelectEmployeeParameter, err error) {
	var (
		employee model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.nik=? AND employees.deleted_at = 0", nik).Take(&employee)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return employee, res.Error
	}
	return employee, nil
}

func (db *EmployeeConnection) FindExcEmployee(id uint) (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		employees []model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.id != ? AND employees.deleted_at = 0", id).Find(&employees)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return employees, res.Error
	}
	return employees, nil
}

func (db *EmployeeConnection) InsertEmployee(employee model.Employee) (employeeOutput model.Employee, err error) {
	res := db.connection.Save(&employee)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return employee, res.Error
	}
	return employee, nil
}

func (db *EmployeeConnection) UpdateEmployee(employee model.Employee, id uint) (employeeOutput model.Employee, err error) {
	res := db.connection.Where("id=?", id).Updates(&employee)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return employee, res.Error
	}
	return employee, nil
}
