package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type EmployeeLeaveRequestTracingRepository interface {
	CountEmployeeLeaveRequestTracingAll(employeeLeaveRequestId int) (count int64, err error)
	FindEmployeeLeaveRequestTracings(employeeLeaveRequestId int) (internalmemotracingOutput []model.EmployeeLeaveRequestTracing, err error)
	FindEmployeeLeaveRequestTracingById(id uint) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error)
	FindExcEmployeeLeaveRequestTracing(id uint) (employeeLeaveRequestTracingOutput []model.EmployeeLeaveRequestTracing, err error)
	InsertEmployeeLeaveRequestTracing(employeeLeaveRequestTracing model.EmployeeLeaveRequestTracing) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error)
	UpdateEmployeeLeaveRequestTracing(employeeLeaveRequestTracing model.EmployeeLeaveRequestTracing, id uint) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error)
}

type employeeLeaveRequestTracingConnection struct {
	connection *gorm.DB
}

func NewEmployeeLeaveRequestTracingRepository(db *gorm.DB) EmployeeLeaveRequestTracingRepository {
	return &employeeLeaveRequestTracingConnection{
		connection: db,
	}
}

func (db *employeeLeaveRequestTracingConnection) CountEmployeeLeaveRequestTracingAll(employeeLeaveRequestId int) (count int64, err error) {
	res := db.connection.Table("employee_leave_request_tracings").Where("employee_leave_request_id = ? AND deleted_at = 0", employeeLeaveRequestId).Count(&count)
	return count, res.Error
}

func (db *employeeLeaveRequestTracingConnection) FindEmployeeLeaveRequestTracings(employeeLeaveRequestId int) (internalmemotracingOutput []model.EmployeeLeaveRequestTracing, err error) {
	var (
		internalmemotracings []model.EmployeeLeaveRequestTracing
	)
	res := db.connection.Where("employee_leave_request_id = ? AND deleted_at = 0", employeeLeaveRequestId).Order("sequence_no DESC").Find(&internalmemotracings)
	return internalmemotracings, res.Error
}

func (db *employeeLeaveRequestTracingConnection) FindEmployeeLeaveRequestTracingById(id uint) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error) {
	var (
		employeeLeaveRequestTracing model.EmployeeLeaveRequestTracing
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&employeeLeaveRequestTracing)
	return employeeLeaveRequestTracing, res.Error
}

func (db *employeeLeaveRequestTracingConnection) FindExcEmployeeLeaveRequestTracing(id uint) (employeeLeaveRequestTracingOutput []model.EmployeeLeaveRequestTracing, err error) {
	var (
		employeeLeaveRequestTracings []model.EmployeeLeaveRequestTracing
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("sequence_no").Find(&employeeLeaveRequestTracings)
	return employeeLeaveRequestTracings, res.Error
}

func (db *employeeLeaveRequestTracingConnection) InsertEmployeeLeaveRequestTracing(employeeLeaveRequestTracing model.EmployeeLeaveRequestTracing) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error) {
	res := db.connection.Save(&employeeLeaveRequestTracing)
	return employeeLeaveRequestTracing, res.Error
}

func (db *employeeLeaveRequestTracingConnection) UpdateEmployeeLeaveRequestTracing(employeeLeaveRequestTracing model.EmployeeLeaveRequestTracing, id uint) (employeeLeaveRequestTracingOutput model.EmployeeLeaveRequestTracing, err error) {
	res := db.connection.Where("id=?", id).Updates(&employeeLeaveRequestTracing)
	return employeeLeaveRequestTracing, res.Error
}
