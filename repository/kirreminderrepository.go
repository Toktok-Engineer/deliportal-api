package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type KIRReminderRepository interface {
	CountKIRReminderAll(vehicleID int) (count int64, err error)
	FindKIRReminders() (kirreminderOutput []model.SelectKIRReminderParameter, err error)
	FindKIRRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (kirreminderOutput []model.SelectKIRReminderParameter, err error)
	SearchKIRReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (kirreminderOutput []model.SelectKIRReminderParameter, err error)
	CountSearchKIRReminder(vehicleID int, search string) (count int64, err error)
	FindKIRReminderById(id uint) (kirReminderOutput model.SelectKIRReminderParameter, err error)
	FindExcKIRReminder(id uint) (kirReminderOutput []model.SelectKIRReminderParameter, err error)
	InsertKIRReminder(kirReminder model.KIRReminder) (kirReminderOutput model.KIRReminder, err error)
	UpdateKIRReminder(kirReminder model.KIRReminder, id uint) (kirReminderOutput model.KIRReminder, err error)
}

type kirReminderConnection struct {
	connection *gorm.DB
}

func NewKIRReminderRepository(db *gorm.DB) KIRReminderRepository {
	return &kirReminderConnection{
		connection: db,
	}
}
func (db *kirReminderConnection) CountKIRReminderAll(vehicleID int) (count int64, err error) {
	res := db.connection.Debug().Table("kir_reminders").Where("deleted_at = 0 AND vehicle_id = ?", vehicleID).Count(&count)
	return count, res.Error
}

func (db *kirReminderConnection) FindKIRReminders() (kirreminderOutput []model.SelectKIRReminderParameter, err error) {
	var (
		kirreminders []model.SelectKIRReminderParameter
	)
	res := db.connection.Table("kir_reminders").Select("id, vehicle_id, to_char(to_timestamp(kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, status, bukti_bayar").Where("deleted_at = 0").Order("kir_end_date").Find(&kirreminders)
	return kirreminders, res.Error
}

func (db *kirReminderConnection) FindKIRRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (kirreminderOutput []model.SelectKIRReminderParameter, err error) {
	var (
		orderDirection string
		kirreminders   []model.SelectKIRReminderParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("kir_reminders").Select("id, vehicle_id, to_char(to_timestamp(kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, status, bukti_bayar").Where("deleted_at = 0 AND vehicle_id = ?", vehicleID).Order(orderDirection).Limit(limit).Offset(offset).Find(&kirreminders)
	return kirreminders, res.Error
}

func (db *kirReminderConnection) SearchKIRReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (kirreminderOutput []model.SelectKIRReminderParameter, err error) {
	var (
		orderDirection string
		final          string
		kirreminders   []model.SelectKIRReminderParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("kir_reminders").Select("id, vehicle_id, to_char(to_timestamp(kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, status, bukti_bayar").Where("(lower(to_char(to_timestamp(kir_end_date::numeric), 'DD-Mon-YYYY')) LIKE ?) AND deleted_at = 0 AND vehicle_id = ?", final, vehicleID).Order(orderDirection).Limit(limit).Offset(offset).Find(&kirreminders)
	return kirreminders, res.Error
}

func (db *kirReminderConnection) CountSearchKIRReminder(vehicleID int, search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("kir_reminders").Where("(lower(to_char(to_timestamp(kir_end_date::numeric), 'DD-Mon-YYYY')) LIKE ?) AND deleted_at = 0 AND vehicle_id = ?", final, vehicleID).Count(&count)
	return count, res.Error
}

func (db *kirReminderConnection) FindKIRReminderById(id uint) (kirReminderOutput model.SelectKIRReminderParameter, err error) {
	var (
		kirReminder model.SelectKIRReminderParameter
	)
	res := db.connection.Table("kir_reminders").Select("id, vehicle_id, to_char(to_timestamp(kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, status, bukti_bayar").Where("id=? AND deleted_at = 0", id).Take(&kirReminder)
	return kirReminder, res.Error
}

func (db *kirReminderConnection) FindExcKIRReminder(id uint) (kirReminderOutput []model.SelectKIRReminderParameter, err error) {
	var (
		kirReminders []model.SelectKIRReminderParameter
	)
	res := db.connection.Table("kir_reminders").Select("id, vehicle_id, to_char(to_timestamp(kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, status, bukti_bayar").Where("id!=? AND deleted_at = 0", id).Order("kir_end_date").Find(&kirReminders)
	return kirReminders, res.Error
}

func (db *kirReminderConnection) InsertKIRReminder(kirReminder model.KIRReminder) (kirReminderOutput model.KIRReminder, err error) {
	res := db.connection.Save(&kirReminder)
	return kirReminder, res.Error
}

func (db *kirReminderConnection) UpdateKIRReminder(kirReminder model.KIRReminder, id uint) (kirReminderOutput model.KIRReminder, err error) {
	res := db.connection.Where("id=?", id).Updates(&kirReminder)
	return kirReminder, res.Error
}
