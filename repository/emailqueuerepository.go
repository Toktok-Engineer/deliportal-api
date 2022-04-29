package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type EmailQueueRepository interface {
	FindEmailQueues() (emailQueueOutput []model.SelectEmailQueueParameter, err error)
	FindEmailQueueById(id uint) (emailQueueOutput model.SelectEmailQueueParameter, err error)
	FindExcEmailQueue(id uint) (emailQueueOutput []model.SelectEmailQueueParameter, err error)
	FindEmailQueueByStatus(status uint) (emailQueueOutput []model.SelectEmailQueueParameter, err error)
	InsertEmailQueue(emailQueue model.EmailQueue) (emailQueueOutput model.EmailQueue, err error)
	UpdateEmailQueue(emailQueue model.EmailQueue, id uint) (emailQueueOutput model.EmailQueue, err error)
}

type EmailQueueConnection struct {
	connection *gorm.DB
}

func NewEmailQueueRepository(db *gorm.DB) EmailQueueRepository {
	return &EmailQueueConnection{
		connection: db,
	}
}

func (db *EmailQueueConnection) FindEmailQueues() (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	var (
		emailQueues []model.SelectEmailQueueParameter
	)
	res := db.connection.Debug().Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, email_queues.created_at, email_queues.updated_at, email_queues.deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("email_queues.deleted_at = 0").Order("email_queues.id").Find(&emailQueues)
	return emailQueues, res.Error
}

func (db *EmailQueueConnection) FindEmailQueueById(id uint) (emailQueueOutput model.SelectEmailQueueParameter, err error) {
	var (
		emailQueue model.SelectEmailQueueParameter
	)

	res := db.connection.Debug().Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, email_queues.created_at, email_queues.updated_at, email_queues.deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("email_queues.id=? AND email_queues.deleted_at = 0", id).Take(&emailQueue)
	return emailQueue, res.Error
}

func (db *EmailQueueConnection) FindExcEmailQueue(id uint) (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	var (
		emailQueues []model.SelectEmailQueueParameter
	)

	res := db.connection.Debug().Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, email_queues.created_at, email_queues.updated_at, email_queues.deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("email_queues.id != ? AND email_queues.deleted_at = 0", id).Order("email_queues.id").Find(&emailQueues)
	return emailQueues, res.Error
}

func (db *EmailQueueConnection) FindEmailQueueByStatus(status uint) (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	var (
		emailQueues []model.SelectEmailQueueParameter
	)

	res := db.connection.Debug().Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, email_queues.created_at, email_queues.updated_at, email_queues.deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("email_queues.status=? AND email_queues.deleted_at = 0", status).Order("email_queues.id").Find(&emailQueues)
	return emailQueues, res.Error
}

func (db *EmailQueueConnection) InsertEmailQueue(emailQueue model.EmailQueue) (emailQueueOutput model.EmailQueue, err error) {
	res := db.connection.Save(&emailQueue)
	return emailQueue, res.Error
}

func (db *EmailQueueConnection) UpdateEmailQueue(emailQueue model.EmailQueue, id uint) (emailQueueOutput model.EmailQueue, err error) {
	res := db.connection.Where("id=?", id).Updates(&emailQueue)
	return emailQueue, res.Error
}
