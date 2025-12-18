package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type STNKReminderRepository interface {
	CountSTNKReminderAll(vehicleID int) (count int64, err error)
	FindSTNKReminders() (stnkreminderOutput []model.SelectSTNKReminderParameter, err error)
	FindSTNKRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (stnkreminderOutput []model.SelectSTNKReminderParameter, err error)
	SearchSTNKReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (stnkreminderOutput []model.SelectSTNKReminderParameter, err error)
	CountSearchSTNKReminder(vehicleID int, search string) (count int64, err error)
	FindSTNKReminderById(id uint) (stnkReminderOutput model.SelectSTNKReminderParameter, err error)
	FindExcSTNKReminder(id uint) (stnkReminderOutput []model.SelectSTNKReminderParameter, err error)
	InsertSTNKReminder(stnkReminder model.STNKReminder) (stnkReminderOutput model.STNKReminder, err error)
	UpdateSTNKReminder(stnkReminder model.STNKReminder, id uint) (stnkReminderOutput model.STNKReminder, err error)
}

type stnkReminderConnection struct {
	connection *gorm.DB
}

func NewSTNKReminderRepository(db *gorm.DB) STNKReminderRepository {
	return &stnkReminderConnection{
		connection: db,
	}
}
func (db *stnkReminderConnection) CountSTNKReminderAll(vehicleID int) (count int64, err error) {
	res := db.connection.Debug().Table("stnk_reminders").Where("deleted_at = 0 AND vehicle_id = ?", vehicleID).Count(&count)
	return count, res.Error
}

func (db *stnkReminderConnection) FindSTNKReminders() (stnkreminderOutput []model.SelectSTNKReminderParameter, err error) {
	var (
		stnkreminders []model.SelectSTNKReminderParameter
	)
	res := db.connection.Table("stnk_reminders").Select("id, vehicle_id, to_char(to_timestamp(stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, status").Where("deleted_at = 0").Order("stnk_end_date").Find(&stnkreminders)
	return stnkreminders, res.Error
}

func (db *stnkReminderConnection) FindSTNKRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (stnkreminderOutput []model.SelectSTNKReminderParameter, err error) {
	var (
		orderDirection string
		stnkreminders  []model.SelectSTNKReminderParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("stnk_reminders").Select("id, vehicle_id, to_char(to_timestamp(stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, status").Where("deleted_at = 0 AND vehicle_id = ?", vehicleID).Order(orderDirection).Limit(limit).Offset(offset).Find(&stnkreminders)
	return stnkreminders, res.Error
}

func (db *stnkReminderConnection) SearchSTNKReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (stnkreminderOutput []model.SelectSTNKReminderParameter, err error) {
	var (
		orderDirection string
		final          string
		stnkreminders  []model.SelectSTNKReminderParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("stnk_reminders").Select("id, vehicle_id, to_char(to_timestamp(stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, status").Where("(lower(to_char(to_timestamp(stnk_end_date::numeric), 'DD-Mon-YYYY')) LIKE ?) AND deleted_at = 0 AND vehicle_id = ?", final, vehicleID).Order(orderDirection).Limit(limit).Offset(offset).Find(&stnkreminders)
	return stnkreminders, res.Error
}

func (db *stnkReminderConnection) CountSearchSTNKReminder(vehicleID int, search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("stnk_reminders").Where("(lower(to_char(to_timestamp(stnk_end_date::numeric), 'DD-Mon-YYYY')) LIKE ?) AND deleted_at = 0 AND vehicle_id = ?", final, vehicleID).Count(&count)
	return count, res.Error
}

func (db *stnkReminderConnection) FindSTNKReminderById(id uint) (stnkReminderOutput model.SelectSTNKReminderParameter, err error) {
	var (
		stnkReminder model.SelectSTNKReminderParameter
	)
	res := db.connection.Table("stnk_reminders").Select("id, vehicle_id, to_char(to_timestamp(stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, status").Where("id=? AND deleted_at = 0", id).Take(&stnkReminder)
	return stnkReminder, res.Error
}

func (db *stnkReminderConnection) FindExcSTNKReminder(id uint) (stnkReminderOutput []model.SelectSTNKReminderParameter, err error) {
	var (
		stnkReminders []model.SelectSTNKReminderParameter
	)
	res := db.connection.Table("stnk_reminders").Select("id, vehicle_id, to_char(to_timestamp(stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, status").Where("id!=? AND deleted_at = 0", id).Order("stnk_end_date").Find(&stnkReminders)
	return stnkReminders, res.Error
}

func (db *stnkReminderConnection) InsertSTNKReminder(stnkReminder model.STNKReminder) (stnkReminderOutput model.STNKReminder, err error) {
	res := db.connection.Save(&stnkReminder)
	return stnkReminder, res.Error
}

func (db *stnkReminderConnection) UpdateSTNKReminder(stnkReminder model.STNKReminder, id uint) (stnkReminderOutput model.STNKReminder, err error) {
	res := db.connection.Where("id=?", id).Updates(&stnkReminder)
	return stnkReminder, res.Error
}
