package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type EmailQueueTypeRepository interface {
	FindEmailQueueTypes() (emailQueueTypeOutput []model.EmailQueueType, err error)
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

func (db *emailQueueTypeConnection) FindEmailQueueTypes() (emailQueueTypeOutput []model.EmailQueueType, err error) {
	var (
		emailQueueTypes []model.EmailQueueType
	)
	res := db.connection.Where("deleted_at = 0").Order("email_queue_type_name").Find(&emailQueueTypes)
	return emailQueueTypes, res.Error
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
