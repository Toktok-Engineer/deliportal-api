package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type EmailQueueReferenceRepository interface {
	FindEmailQueueReferences() (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error)
	FindEmailQueueReferenceById(id uint) (emailQueueReferenceOutput model.SelectEmailQueueReferenceParameter, err error)
	FindExcEmailQueueReference(id uint) (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error)
	InsertEmailQueueReference(emailQueueReference model.EmailQueueReference) (emailQueueReferenceOutput model.EmailQueueReference, err error)
	UpdateEmailQueueReference(emailQueueReference model.EmailQueueReference, id uint) (emailQueueReferenceOutput model.EmailQueueReference, err error)
}

type EmailQueueReferenceConnection struct {
	connection *gorm.DB
}

func NewEmailQueueReferenceRepository(db *gorm.DB) EmailQueueReferenceRepository {
	return &EmailQueueReferenceConnection{
		connection: db,
	}
}

func (db *EmailQueueReferenceConnection) FindEmailQueueReferences() (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error) {
	var (
		emailQueueReferences []model.SelectEmailQueueReferenceParameter
	)
	res := db.connection.Debug().Table("email_queue_references").Select("email_queue_references.id, email_queue_references.email_queue_id, email_queue_references.email_queue_type_id, email_queue_types.email_queue_type_name, email_queue_references.reference_id, email_queue_references.remark, email_queue_references.created_user_id, email_queue_references.updated_user_id, email_queue_references.deleted_user_id, email_queue_references.created_at, email_queue_references.updated_at, email_queue_references.deleted_at").Joins("left join email_queue_types ON email_queue_references.email_queue_type_id = email_queue_types.id").Where("email_queue_references.deleted_at = 0").Order("email_queue_references.id").Find(&emailQueueReferences)
	return emailQueueReferences, res.Error
}

func (db *EmailQueueReferenceConnection) FindEmailQueueReferenceById(id uint) (emailQueueReferenceOutput model.SelectEmailQueueReferenceParameter, err error) {
	var (
		emailQueueReference model.SelectEmailQueueReferenceParameter
	)
	res := db.connection.Debug().Table("email_queue_references").Select("email_queue_references.id, email_queue_references.email_queue_id, email_queue_references.email_queue_type_id, email_queue_types.email_queue_type_name, email_queue_references.reference_id, email_queue_references.remark, email_queue_references.created_user_id, email_queue_references.updated_user_id, email_queue_references.deleted_user_id, email_queue_references.created_at, email_queue_references.updated_at, email_queue_references.deleted_at").Joins("left join email_queue_types ON email_queue_references.email_queue_type_id = email_queue_types.id").Where("email_queue_references.id=? AND email_queue_references.deleted_at = 0", id).Take(&emailQueueReference)
	return emailQueueReference, res.Error
}

func (db *EmailQueueReferenceConnection) FindExcEmailQueueReference(id uint) (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error) {
	var (
		emailQueueReferences []model.SelectEmailQueueReferenceParameter
	)
	res := db.connection.Debug().Table("email_queue_references").Select("email_queue_references.id, email_queue_references.email_queue_id, email_queue_references.email_queue_type_id, email_queue_types.email_queue_type_name, email_queue_references.reference_id, email_queue_references.remark, email_queue_references.created_user_id, email_queue_references.updated_user_id, email_queue_references.deleted_user_id, email_queue_references.created_at, email_queue_references.updated_at, email_queue_references.deleted_at").Joins("left join email_queue_types ON email_queue_references.email_queue_type_id = email_queue_types.id").Where("email_queue_references.id!=? AND email_queue_references.deleted_at = 0", id).Order("email_queue_references.id").Find(&emailQueueReferences)
	return emailQueueReferences, res.Error
}

func (db *EmailQueueReferenceConnection) InsertEmailQueueReference(emailQueueReference model.EmailQueueReference) (emailQueueReferenceOutput model.EmailQueueReference, err error) {
	res := db.connection.Save(&emailQueueReference)
	return emailQueueReference, res.Error
}

func (db *EmailQueueReferenceConnection) UpdateEmailQueueReference(emailQueueReference model.EmailQueueReference, id uint) (emailQueueReferenceOutput model.EmailQueueReference, err error) {
	res := db.connection.Where("id=?", id).Updates(&emailQueueReference)
	return emailQueueReference, res.Error
}
