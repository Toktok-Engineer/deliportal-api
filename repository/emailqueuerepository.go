package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type EmailQueueRepository interface {
	CountEmailQueueAll() (count int64, err error)
	FindEmailQueues() (emailQueueOutput []model.SelectEmailQueueParameter, err error)
	FindEmailQueuesOffset(limit int, offset int, order string, dir string) (emailQueueOutput []model.SelectEmailQueueParameter, err error)
	SearchEmailQueue(limit int, offset int, order string, dir string, search string) (emailQueueOutput []model.SelectEmailQueueParameter, err error)
	CountSearchEmailQueue(search string) (count int64, err error)
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

func (db *EmailQueueConnection) CountEmailQueueAll() (count int64, err error) {
	res := db.connection.Table("email_queues").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *EmailQueueConnection) FindEmailQueues() (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	var (
		emailQueues []model.SelectEmailQueueParameter
	)
	res := db.connection.Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, to_char(to_timestamp(email_queues.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(email_queues.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(email_queues.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("email_queues.deleted_at = 0").Order("email_queue_types.email_queue_type_name").Find(&emailQueues)
	return emailQueues, res.Error
}

func (db *EmailQueueConnection) FindEmailQueuesOffset(limit int, offset int, order string, dir string) (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	var (
		orderDirection string
		emailQueues    []model.SelectEmailQueueParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, to_char(to_timestamp(email_queues.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(email_queues.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(email_queues.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("email_queues.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&emailQueues)
	return emailQueues, res.Error
}

func (db *EmailQueueConnection) SearchEmailQueue(limit int, offset int, order string, dir string, search string) (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	var (
		orderDirection string
		final          string
		emailQueues    []model.SelectEmailQueueParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, to_char(to_timestamp(email_queues.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(email_queues.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(email_queues.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("(lower(email_queue_types.email_queue_type_name) LIKE ? OR lower(email_queues.email_recipient) LIKE ? OR lower(email_queues.email_cc) LIKE ? OR lower(email_queues.email_subject) LIKE ? OR lower(email_queues.email_body) LIKE ? OR lower(email_queues.error_message) LIKE ? OR lower(email_queues.remark) LIKE ?) AND email_queues.deleted_at = 0", final, final, final, final, final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&emailQueues)
	return emailQueues, res.Error
}

func (db *EmailQueueConnection) CountSearchEmailQueue(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, to_char(to_timestamp(email_queues.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(email_queues.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(email_queues.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("(lower(email_queue_types.email_queue_type_name) LIKE ? OR lower(email_queues.email_recipient) LIKE ? OR lower(email_queues.email_cc) LIKE ? OR lower(email_queues.email_subject) LIKE ? OR lower(email_queues.email_body) LIKE ? OR lower(email_queues.error_message) LIKE ? OR lower(email_queues.remark) LIKE ?) AND email_queues.deleted_at = 0", final, final, final, final, final, final, final).Count(&count)
	return count, res.Error
}

func (db *EmailQueueConnection) FindEmailQueueById(id uint) (emailQueueOutput model.SelectEmailQueueParameter, err error) {
	var (
		emailQueue model.SelectEmailQueueParameter
	)

	res := db.connection.Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, to_char(to_timestamp(email_queues.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(email_queues.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(email_queues.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("email_queues.id=? AND email_queues.deleted_at = 0", id).Take(&emailQueue)
	return emailQueue, res.Error
}

func (db *EmailQueueConnection) FindExcEmailQueue(id uint) (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	var (
		emailQueues []model.SelectEmailQueueParameter
	)

	res := db.connection.Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, to_char(to_timestamp(email_queues.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(email_queues.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(email_queues.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("email_queues.id != ? AND email_queues.deleted_at = 0", id).Order("email_queues.id").Find(&emailQueues)
	return emailQueues, res.Error
}

func (db *EmailQueueConnection) FindEmailQueueByStatus(status uint) (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	var (
		emailQueues []model.SelectEmailQueueParameter
	)

	res := db.connection.Table("email_queues").Select("email_queues.id, email_queues.email_queue_type_id, email_queue_types.email_queue_type_name, email_queues.email_recipient, email_queues.email_cc, email_queues.email_subject, email_queues.email_body, email_queues.status, email_queues.error_message, email_queues.remark, email_queues.created_user_id, email_queues.updated_user_id, email_queues.deleted_user_id, to_char(to_timestamp(email_queues.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(email_queues.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(email_queues.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join email_queue_types ON email_queues.email_queue_type_id = email_queue_types.id").Where("email_queues.status=? AND email_queues.deleted_at = 0", status).Order("email_queues.id").Find(&emailQueues)
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
