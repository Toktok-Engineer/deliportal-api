package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type EmailQueueReferenceRepository interface {
	CountEmailQueueReferenceAll() (count int64, err error)
	FindEmailQueueReferences() (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error)
	FindEmailQueueReferencesOffset(limit int, offset int, order string, dir string) (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error)
	SearchEmailQueueReference(limit int, offset int, order string, dir string, search string) (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error)
	CountSearchEmailQueueReference(search string) (count int64, err error)
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

func (db *EmailQueueReferenceConnection) CountEmailQueueReferenceAll() (count int64, err error) {
	res := db.connection.Table("email_queue_references").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *EmailQueueReferenceConnection) FindEmailQueueReferences() (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error) {
	var (
		emailQueueReferences []model.SelectEmailQueueReferenceParameter
	)
	res := db.connection.Table("email_queue_references").Select("email_queue_references.id, email_queue_references.email_queue_id, email_queue_references.email_queue_type_id, email_queue_types.email_queue_type_name, email_queue_references.reference_id, email_queue_references.remark, email_queue_references.created_user_id, email_queue_references.updated_user_id, email_queue_references.deleted_user_id, email_queue_references.created_at, email_queue_references.updated_at, email_queue_references.deleted_at").Joins("left join email_queue_types ON email_queue_references.email_queue_type_id = email_queue_types.id").Where("email_queue_references.deleted_at = 0").Order("email_queue_references.email_queue_id").Find(&emailQueueReferences)
	return emailQueueReferences, res.Error
}

func (db *EmailQueueReferenceConnection) FindEmailQueueReferencesOffset(limit int, offset int, order string, dir string) (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error) {
	var (
		orderDirection       string
		emailQueueReferences []model.SelectEmailQueueReferenceParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("email_queue_references").Select("email_queue_references.id, email_queue_references.email_queue_id, email_queue_references.email_queue_type_id, email_queue_types.email_queue_type_name, email_queue_references.reference_id, email_queue_references.remark, email_queue_references.created_user_id, email_queue_references.updated_user_id, email_queue_references.deleted_user_id, email_queue_references.created_at, email_queue_references.updated_at, email_queue_references.deleted_at").Joins("left join email_queue_types ON email_queue_references.email_queue_type_id = email_queue_types.id").Where("email_queue_references.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&emailQueueReferences)
	return emailQueueReferences, res.Error
}

func (db *EmailQueueReferenceConnection) SearchEmailQueueReference(limit int, offset int, order string, dir string, search string) (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error) {
	var (
		orderDirection       string
		final                string
		emailQueueReferences []model.SelectEmailQueueReferenceParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("email_queue_references").Select("email_queue_references.id, email_queue_references.email_queue_id, email_queue_references.email_queue_type_id, email_queue_types.email_queue_type_name, email_queue_references.reference_id, email_queue_references.remark, email_queue_references.created_user_id, email_queue_references.updated_user_id, email_queue_references.deleted_user_id, email_queue_references.created_at, email_queue_references.updated_at, email_queue_references.deleted_at").Joins("left join email_queue_types ON email_queue_references.email_queue_type_id = email_queue_types.id").Where("lower(email_queue_types.email_queue_type_name) LIKE ? AND email_queue_references.deleted_at = 0", final).Order(orderDirection).Limit(limit).Offset(offset).Find(&emailQueueReferences)
	return emailQueueReferences, res.Error
}

func (db *EmailQueueReferenceConnection) CountSearchEmailQueueReference(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("email_queue_references").Select("email_queue_references.id, email_queue_references.email_queue_id, email_queue_references.email_queue_type_id, email_queue_types.email_queue_type_name, email_queue_references.reference_id, email_queue_references.remark, email_queue_references.created_user_id, email_queue_references.updated_user_id, email_queue_references.deleted_user_id, email_queue_references.created_at, email_queue_references.updated_at, email_queue_references.deleted_at").Joins("left join email_queue_types ON email_queue_references.email_queue_type_id = email_queue_types.id").Where("lower(email_queue_types.email_queue_type_name) LIKE ? AND email_queue_references.deleted_at = 0", final).Count(&count)
	return count, res.Error
}

func (db *EmailQueueReferenceConnection) FindEmailQueueReferenceById(id uint) (emailQueueReferenceOutput model.SelectEmailQueueReferenceParameter, err error) {
	var (
		emailQueueReference model.SelectEmailQueueReferenceParameter
	)
	res := db.connection.Table("email_queue_references").Select("email_queue_references.id, email_queue_references.email_queue_id, email_queue_references.email_queue_type_id, email_queue_types.email_queue_type_name, email_queue_references.reference_id, email_queue_references.remark, email_queue_references.created_user_id, email_queue_references.updated_user_id, email_queue_references.deleted_user_id, email_queue_references.created_at, email_queue_references.updated_at, email_queue_references.deleted_at").Joins("left join email_queue_types ON email_queue_references.email_queue_type_id = email_queue_types.id").Where("email_queue_references.id=? AND email_queue_references.deleted_at = 0", id).Take(&emailQueueReference)
	return emailQueueReference, res.Error
}

func (db *EmailQueueReferenceConnection) FindExcEmailQueueReference(id uint) (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error) {
	var (
		emailQueueReferences []model.SelectEmailQueueReferenceParameter
	)
	res := db.connection.Table("email_queue_references").Select("email_queue_references.id, email_queue_references.email_queue_id, email_queue_references.email_queue_type_id, email_queue_types.email_queue_type_name, email_queue_references.reference_id, email_queue_references.remark, email_queue_references.created_user_id, email_queue_references.updated_user_id, email_queue_references.deleted_user_id, email_queue_references.created_at, email_queue_references.updated_at, email_queue_references.deleted_at").Joins("left join email_queue_types ON email_queue_references.email_queue_type_id = email_queue_types.id").Where("email_queue_references.id!=? AND email_queue_references.deleted_at = 0", id).Order("email_queue_references.id").Find(&emailQueueReferences)
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
