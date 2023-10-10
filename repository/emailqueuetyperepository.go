package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type EmailQueueTypeRepository interface {
	CountEmailQueueTypeAll() (count int64, err error)
	FindEmailQueueTypes() (emailQueueTypeOutput []model.EmailQueueType, err error)
	FindEmailQueueTypesOffset(limit int, offset int, order string, dir string) (emailQueueTypeOutput []model.EmailQueueType, err error)
	SearchEmailQueueType(limit int, offset int, order string, dir string, search string) (emailQueueTypeOutput []model.EmailQueueType, err error)
	CountSearchEmailQueueType(search string) (count int64, err error)
	FindEmailQueueTypeById(id uint) (emailQueueTypeOutput model.EmailQueueType, err error)
	FindExcEmailQueueType(id uint) (emailQueueTypeOutput []model.EmailQueueType, err error)
	InsertEmailQueueType(emailQueueType model.EmailQueueType) (emailQueueTypeOutput model.EmailQueueType, err error)
	UpdateEmailQueueType(emailQueueType model.EmailQueueType, id uint) (emailQueueTypeOutput model.EmailQueueType, err error)
}

type emailQueueTypeConnection struct {
	connection *gorm.DB
}

func NewEmailQueueTypeRepository(db *gorm.DB) EmailQueueTypeRepository {
	return &emailQueueTypeConnection{
		connection: db,
	}
}

func (db *emailQueueTypeConnection) CountEmailQueueTypeAll() (count int64, err error) {
	res := db.connection.Debug().Table("email_queue_types").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *emailQueueTypeConnection) FindEmailQueueTypes() (emailQueueTypeOutput []model.EmailQueueType, err error) {
	var (
		emailQueueTypes []model.EmailQueueType
	)
	res := db.connection.Where("deleted_at = 0").Order("email_queue_type_name").Find(&emailQueueTypes)
	return emailQueueTypes, res.Error
}

func (db *emailQueueTypeConnection) FindEmailQueueTypesOffset(limit int, offset int, order string, dir string) (emailQueueTypeOutput []model.EmailQueueType, err error) {
	var (
		orderDirection  string
		emailQueueTypes []model.EmailQueueType
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&emailQueueTypes)
	return emailQueueTypes, res.Error
}

func (db *emailQueueTypeConnection) SearchEmailQueueType(limit int, offset int, order string, dir string, search string) (emailQueueTypeOutput []model.EmailQueueType, err error) {
	var (
		orderDirection  string
		final           string
		emailQueueTypes []model.EmailQueueType
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(email_queue_type_name) LIKE ? OR lower(table_reference) LIKE ? OR lower(field_reference) LIKE ? OR lower(email_recipient) LIKE ? OR lower(email_subject) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final, final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&emailQueueTypes)
	return emailQueueTypes, res.Error
}

func (db *emailQueueTypeConnection) CountSearchEmailQueueType(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("email_queue_types").Where("(lower(email_queue_type_name) LIKE ? OR lower(table_reference) LIKE ? OR lower(field_reference) LIKE ? OR lower(email_recipient) LIKE ? OR lower(email_subject) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final, final, final, final).Count(&count)
	return count, res.Error
}

func (db *emailQueueTypeConnection) FindEmailQueueTypeById(id uint) (emailQueueTypeOutput model.EmailQueueType, err error) {
	var (
		emailQueueType model.EmailQueueType
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&emailQueueType)
	return emailQueueType, res.Error
}

func (db *emailQueueTypeConnection) FindExcEmailQueueType(id uint) (emailQueueTypeOutput []model.EmailQueueType, err error) {
	var (
		emailQueueTypes []model.EmailQueueType
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("email_queue_type_name").Find(&emailQueueTypes)
	return emailQueueTypes, res.Error
}

func (db *emailQueueTypeConnection) InsertEmailQueueType(emailQueueType model.EmailQueueType) (emailQueueTypeOutput model.EmailQueueType, err error) {
	res := db.connection.Save(&emailQueueType)
	return emailQueueType, res.Error
}

func (db *emailQueueTypeConnection) UpdateEmailQueueType(emailQueueType model.EmailQueueType, id uint) (emailQueueTypeOutput model.EmailQueueType, err error) {
	res := db.connection.Where("id=?", id).Updates(&emailQueueType)
	return emailQueueType, res.Error
}
