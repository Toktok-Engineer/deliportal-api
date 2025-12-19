package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type EmployeeLeaveRequestDumpRepository interface {
	CountEmployeeLeaveRequestDumpAll() (count int64, err error)
	FindEmployeeLeaveRequestDumps() (businessunitOutput []model.EmployeeLeaveRequestDump, err error)
	FindEmployeeLeaveRequestDumpsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.EmployeeLeaveRequestDump, err error)
	// SearchEmployeeLeaveRequestDump(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.EmployeeLeaveRequestDump, err error)
	// CountSearchEmployeeLeaveRequestDump(search string) (count int64, err error)
	FindEmployeeLeaveRequestDumpById(id uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error)
	FindEmployeeLeaveRequestDumpByErlIdandEmp(id uint, tracingID uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error)
	FindExcEmployeeLeaveRequestDump(id uint) (employeeLeaveRequestDumpOutput []model.EmployeeLeaveRequestDump, err error)
	InsertEmployeeLeaveRequestDump(employeeLeaveRequestDump model.EmployeeLeaveRequestDump) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error)
	UpdateEmployeeLeaveRequestDump(employeeLeaveRequestDump model.EmployeeLeaveRequestDump, id uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error)
}

type employeeLeaveRequestDumpConnection struct {
	connection *gorm.DB
}

func NewEmployeeLeaveRequestDumpRepository(db *gorm.DB) EmployeeLeaveRequestDumpRepository {
	return &employeeLeaveRequestDumpConnection{
		connection: db,
	}
}
func (db *employeeLeaveRequestDumpConnection) CountEmployeeLeaveRequestDumpAll() (count int64, err error) {
	res := db.connection.Table("employee_leave_request_dumps").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *employeeLeaveRequestDumpConnection) FindEmployeeLeaveRequestDumps() (businessunitOutput []model.EmployeeLeaveRequestDump, err error) {
	var (
		businessunits []model.EmployeeLeaveRequestDump
	)
	res := db.connection.Where("deleted_at = 0").Order("id").Find(&businessunits)
	return businessunits, res.Error
}

func (db *employeeLeaveRequestDumpConnection) FindEmployeeLeaveRequestDumpsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.EmployeeLeaveRequestDump, err error) {
	var (
		orderDirection string
		businessunits  []model.EmployeeLeaveRequestDump
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&businessunits)
	return businessunits, res.Error
}

// func (db *employeeLeaveRequestDumpConnection) SearchEmployeeLeaveRequestDump(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.EmployeeLeaveRequestDump, err error) {
// 	var (
// 		orderDirection string
// 		final          string
// 		businessunits  []model.EmployeeLeaveRequestDump
// 	)
// 	orderDirection = order + " " + dir
// 	final = "%" + strings.ToLower(search) + "%"
// 	res := db.connection.Where("(lower(business_unit_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&businessunits)
// 	return businessunits, res.Error
// }

// func (db *employeeLeaveRequestDumpConnection) CountSearchEmployeeLeaveRequestDump(search string) (count int64, err error) {
// 	final := "%" + strings.ToLower(search) + "%"
// 	res := db.connection.Table("business_units").Where("(lower(business_unit_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
// 	return count, res.Error
// }

func (db *employeeLeaveRequestDumpConnection) FindEmployeeLeaveRequestDumpById(id uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error) {
	var (
		employeeLeaveRequestDump model.EmployeeLeaveRequestDump
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&employeeLeaveRequestDump)
	return employeeLeaveRequestDump, res.Error
}

func (db *employeeLeaveRequestDumpConnection) FindEmployeeLeaveRequestDumpByErlIdandEmp(id uint, tracingID uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error) {
	var (
		employeeLeaveRequestDump model.EmployeeLeaveRequestDump
	)
	res := db.connection.Where("employee_leave_request_id=? AND employee_leave_request_tracing_id=? AND deleted_at = 0", id, tracingID).Take(&employeeLeaveRequestDump)
	return employeeLeaveRequestDump, res.Error
}

func (db *employeeLeaveRequestDumpConnection) FindExcEmployeeLeaveRequestDump(id uint) (employeeLeaveRequestDumpOutput []model.EmployeeLeaveRequestDump, err error) {
	var (
		employeeLeaveRequestDumps []model.EmployeeLeaveRequestDump
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("id").Find(&employeeLeaveRequestDumps)
	return employeeLeaveRequestDumps, res.Error
}

func (db *employeeLeaveRequestDumpConnection) InsertEmployeeLeaveRequestDump(employeeLeaveRequestDump model.EmployeeLeaveRequestDump) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error) {
	res := db.connection.Save(&employeeLeaveRequestDump)
	return employeeLeaveRequestDump, res.Error
}

func (db *employeeLeaveRequestDumpConnection) UpdateEmployeeLeaveRequestDump(employeeLeaveRequestDump model.EmployeeLeaveRequestDump, id uint) (employeeLeaveRequestDumpOutput model.EmployeeLeaveRequestDump, err error) {
	res := db.connection.Where("id=?", id).Updates(&employeeLeaveRequestDump)
	return employeeLeaveRequestDump, res.Error
}
