package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type AsuransiReminderRepository interface {
	CountAsuransiReminderAll(vehicleID int) (count int64, err error)
	FindAsuransiReminders() (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error)
	FindAsuransiRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error)
	SearchAsuransiReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error)
	CountSearchAsuransiReminder(vehicleID int, search string) (count int64, err error)
	FindAsuransiReminderById(id uint) (asuransiReminderOutput model.SelectAsuransiReminderParameter, err error)
	FindExcAsuransiReminder(id uint) (asuransiReminderOutput []model.SelectAsuransiReminderParameter, err error)
	InsertAsuransiReminder(asuransiReminder model.AsuransiReminder) (asuransiReminderOutput model.AsuransiReminder, err error)
	UpdateAsuransiReminder(asuransiReminder model.AsuransiReminder, id uint) (asuransiReminderOutput model.AsuransiReminder, err error)
}

type asuransiReminderConnection struct {
	connection *gorm.DB
}

func NewAsuransiReminderRepository(db *gorm.DB) AsuransiReminderRepository {
	return &asuransiReminderConnection{
		connection: db,
	}
}
func (db *asuransiReminderConnection) CountAsuransiReminderAll(vehicleID int) (count int64, err error) {
	res := db.connection.Debug().Table("asuransi_reminders").Where("asuransi_reminders.deleted_at = 0 AND asuransi_reminders.vehicle_id = ?", vehicleID).Count(&count)
	return count, res.Error
}

func (db *asuransiReminderConnection) FindAsuransiReminders() (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error) {
	var (
		asuransireminders []model.SelectAsuransiReminderParameter
	)
	res := db.connection.Table("asuransi_reminders").Select("asuransi_reminders.id, asuransi_reminders.vehicle_id, to_char(to_timestamp(asuransi_end_date::numeric), 'DD-Mon-YYYY') as asuransi_end_date, asuransi_reminders.asuransi_id, asuransis.asuransi_name, asuransi_reminders.nomor_polis, asuransi_reminders.status, asuransi_reminders.price , asuransi_reminders.bukti_bayar").Joins("left join asuransis ON asuransis.id = asuransi_reminders.asuransi_id").Where("asuransi_reminders.deleted_at = 0").Order("asuransi_reminders.asuransi_end_date").Find(&asuransireminders)
	return asuransireminders, res.Error
}

func (db *asuransiReminderConnection) FindAsuransiRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error) {
	var (
		orderDirection    string
		asuransireminders []model.SelectAsuransiReminderParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("asuransi_reminders").Select("asuransi_reminders.id, asuransi_reminders.vehicle_id, to_char(to_timestamp(asuransi_end_date::numeric), 'DD-Mon-YYYY') as asuransi_end_date, asuransi_reminders.asuransi_id, asuransis.asuransi_name, asuransi_reminders.nomor_polis, asuransi_reminders.status, asuransi_reminders.price , asuransi_reminders.bukti_bayar").Joins("left join asuransis ON asuransis.id = asuransi_reminders.asuransi_id").Where("asuransi_reminders.deleted_at = 0 AND asuransi_reminders.vehicle_id = ?", vehicleID).Order(orderDirection).Limit(limit).Offset(offset).Find(&asuransireminders)
	return asuransireminders, res.Error
}

func (db *asuransiReminderConnection) SearchAsuransiReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error) {
	var (
		orderDirection    string
		final             string
		asuransireminders []model.SelectAsuransiReminderParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("asuransi_reminders").Select("asuransi_reminders.id, asuransi_reminders.vehicle_id, to_char(to_timestamp(asuransi_end_date::numeric), 'DD-Mon-YYYY') as asuransi_end_date, asuransi_reminders.asuransi_id, asuransis.asuransi_name, asuransi_reminders.nomor_polis, asuransi_reminders.status, asuransi_reminders.price , asuransi_reminders.bukti_bayar").Joins("left join asuransis ON asuransis.id = asuransi_reminders.asuransi_id").Where("(lower(to_char(to_timestamp(asuransi_reminders.asuransi_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(asuransi_reminders.price::text) LIKE ? OR lower(asuransis.asuransi_name) LIKE ? OR lower(asuransi_reminders.nomor_polis) LIKE ?) AND asuransi_reminders.deleted_at = 0 AND asuransi_reminders.vehicle_id = ?", final, final, final, final, vehicleID).Order(orderDirection).Limit(limit).Offset(offset).Find(&asuransireminders)
	return asuransireminders, res.Error
}

func (db *asuransiReminderConnection) CountSearchAsuransiReminder(vehicleID int, search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("asuransi_reminders").Select("asuransi_reminders.id, asuransi_reminders.vehicle_id, to_char(to_timestamp(asuransi_end_date::numeric), 'DD-Mon-YYYY') as asuransi_end_date, asuransi_reminders.asuransi_id, asuransis.asuransi_name, asuransi_reminders.nomor_polis, asuransi_reminders.status, asuransi_reminders.price , asuransi_reminders.bukti_bayar").Joins("left join asuransis ON asuransis.id = asuransi_reminders.asuransi_id").Where("(lower(to_char(to_timestamp(asuransi_reminders.asuransi_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(asuransi_reminders.price::text) LIKE ? OR lower(asuransis.asuransi_name) LIKE ? OR lower(asuransi_reminders.nomor_polis) LIKE ?) AND asuransi_reminders.deleted_at = 0 AND asuransi_reminders.vehicle_id = ?", final, final, final, final, vehicleID).Count(&count)
	return count, res.Error
}

func (db *asuransiReminderConnection) FindAsuransiReminderById(id uint) (asuransiReminderOutput model.SelectAsuransiReminderParameter, err error) {
	var (
		asuransiReminder model.SelectAsuransiReminderParameter
	)
	res := db.connection.Table("asuransi_reminders").Select("asuransi_reminders.id, asuransi_reminders.vehicle_id, to_char(to_timestamp(asuransi_end_date::numeric), 'DD-Mon-YYYY') as asuransi_end_date, asuransi_reminders.asuransi_id, asuransis.asuransi_name, asuransi_reminders.nomor_polis, asuransi_reminders.status, asuransi_reminders.price , asuransi_reminders.bukti_bayar").Joins("left join asuransis ON asuransis.id = asuransi_reminders.asuransi_id").Where("asuransi_reminders.id=? AND asuransi_reminders.deleted_at = 0", id).Take(&asuransiReminder)
	return asuransiReminder, res.Error
}

func (db *asuransiReminderConnection) FindExcAsuransiReminder(id uint) (asuransiReminderOutput []model.SelectAsuransiReminderParameter, err error) {
	var (
		asuransiReminders []model.SelectAsuransiReminderParameter
	)
	res := db.connection.Table("asuransi_reminders").Select("asuransi_reminders.id, asuransi_reminders.vehicle_id, to_char(to_timestamp(asuransi_end_date::numeric), 'DD-Mon-YYYY') as asuransi_end_date, asuransi_reminders.asuransi_id, asuransis.asuransi_name, asuransi_reminders.nomor_polis, asuransi_reminders.status, asuransi_reminders.price , asuransi_reminders.bukti_bayar").Joins("left join asuransis ON asuransis.id = asuransi_reminders.asuransi_id").Where("asuransi_reminders.id!=? AND asuransi_reminders.deleted_at = 0", id).Order("asuransi_reminders.asuransi_end_date").Find(&asuransiReminders)
	return asuransiReminders, res.Error
}

func (db *asuransiReminderConnection) InsertAsuransiReminder(asuransiReminder model.AsuransiReminder) (asuransiReminderOutput model.AsuransiReminder, err error) {
	res := db.connection.Save(&asuransiReminder)
	return asuransiReminder, res.Error
}

func (db *asuransiReminderConnection) UpdateAsuransiReminder(asuransiReminder model.AsuransiReminder, id uint) (asuransiReminderOutput model.AsuransiReminder, err error) {
	res := db.connection.Where("id=?", id).Updates(&asuransiReminder)
	return asuransiReminder, res.Error
}
