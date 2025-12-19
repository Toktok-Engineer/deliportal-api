package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type LeaveTypesRepository interface {
	CountLeaveTypeAll() (count int64, err error)
	FindLeaveTypes() (leavetypesOutput []model.LeaveTypes, err error)
	FindLeaveTypesOffset(limit int, offset int, order string, dir string) (leavetypesOutput []model.LeaveTypes, err error)
	SearchLeaveType(limit int, offset int, order string, dir string, search string) (leavetypesOutput []model.LeaveTypes, err error)
	CountSearchLeaveType(search string) (count int64, err error)
	FindLeaveTypeById(id uint) (leavetypesOutput model.LeaveTypes, err error)
	CountLeaveTypeName(search string) (count int64, err error)
	FindExcLeaveType(id uint) (leavetypesOutput []model.LeaveTypes, err error)
	InsertLeaveType(leavetypes model.LeaveTypes) (leavetypesOutput model.LeaveTypes, err error)
	UpdateLeaveType(leavetypes model.LeaveTypes, id uint) (leavetypesOutput model.LeaveTypes, err error)
}

type LeaveTypeConnection struct {
	connection *gorm.DB
}

func NewLeaveTypeRepository(db *gorm.DB) LeaveTypesRepository {
	return &LeaveTypeConnection{
		connection: db,
	}
}

func (db *LeaveTypeConnection) CountLeaveTypeAll() (count int64, err error) {
	res := db.connection.Table("leave_types").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *LeaveTypeConnection) FindLeaveTypes() (leavetypesOutput []model.LeaveTypes, err error) {
	var (
		leavetypes []model.LeaveTypes
	)
	res := db.connection.Where("deleted_at = 0").Order("leave_type_name").Find(&leavetypes)
	return leavetypes, res.Error
}

func (db *LeaveTypeConnection) FindLeaveTypesOffset(limit int, offset int, order string, dir string) (leavetypesOutput []model.LeaveTypes, err error) {
	var (
		orderDirection string
		leavetypes     []model.LeaveTypes
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&leavetypes)
	return leavetypes, res.Error
}

func (db *LeaveTypeConnection) SearchLeaveType(limit int, offset int, order string, dir string, search string) (leavetypesOutput []model.LeaveTypes, err error) {
	var (
		orderDirection string
		final          string
		leavetypes     []model.LeaveTypes
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(leave_type_name) LIKE ? OR lower(limits::varchar(50)) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&leavetypes)
	return leavetypes, res.Error
}

func (db *LeaveTypeConnection) CountSearchLeaveType(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("leave_types").Where("(lower(leave_type_name) LIKE ? OR lower(limits::varchar(50)) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final).Count(&count)
	return count, res.Error
}

func (db *LeaveTypeConnection) FindLeaveTypeById(id uint) (leavetypesOutput model.LeaveTypes, err error) {
	var (
		leavetypes model.LeaveTypes
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&leavetypes)
	return leavetypes, res.Error
}

func (db *LeaveTypeConnection) CountLeaveTypeName(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("leave_types").Where("lower(leave_type_name) LIKE ?", final).Count(&count)
	return count, res.Error
}

func (db *LeaveTypeConnection) FindExcLeaveType(id uint) (leavetypesOutput []model.LeaveTypes, err error) {
	var (
		leavetypes []model.LeaveTypes
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("leave_type_name").Find(&leavetypes)
	return leavetypes, res.Error
}

func (db *LeaveTypeConnection) InsertLeaveType(leavetypes model.LeaveTypes) (leavetypesOutput model.LeaveTypes, err error) {
	res := db.connection.Save(&leavetypes)
	return leavetypes, res.Error
}

func (db *LeaveTypeConnection) UpdateLeaveType(leavetypes model.LeaveTypes, id uint) (leavetypesOutput model.LeaveTypes, err error) {
	res := db.connection.Where("id=?", id).Updates(&leavetypes)
	return leavetypes, res.Error
}
