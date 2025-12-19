package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type InternalMemoAttachmentRepository interface {
	CountInternalMemoAttachmentAll(internalMemoId int) (count int64, err error)
	FindInternalMemoAttachments(internalMemoId int) (internalMemoAttachmentOutput []model.InternalMemoAttachment, err error)
	FindInternalMemoAttachmentsOffset(limit int, offset int, order string, dir string, internalMemoId int) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error)
	SearchInternalMemoAttachment(limit int, offset int, order string, dir string, search string, internalMemoId int) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error)
	CountSearchInternalMemoAttachment(search string, internalMemoId int) (count int64, err error)
	FindInternalMemoAttachmentById(id uint) (internalMemoAttachmentOutput model.SelectInternalMemoAttachmentParameter, err error)
	FindExcInternalMemoAttachment(id uint) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error)
	FindInternalMemoAttachmentByCompanyId(id uint) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error)
	InsertInternalMemoAttachment(internalMemoAttachment model.InternalMemoAttachment) (internalMemoAttachmentOutput model.InternalMemoAttachment, err error)
	UpdateInternalMemoAttachment(internalMemoAttachment model.InternalMemoAttachment, id uint) (internalMemoAttachmentOutput model.InternalMemoAttachment, err error)
}

type InternalMemoAttachmentConnection struct {
	connection *gorm.DB
}

func NewInternalMemoAttachmentRepository(db *gorm.DB) InternalMemoAttachmentRepository {
	return &InternalMemoAttachmentConnection{
		connection: db,
	}
}

func (db *InternalMemoAttachmentConnection) CountInternalMemoAttachmentAll(internalMemoId int) (count int64, err error) {
	res := db.connection.Table("internal_memo_attachments").Where("internal_memo_id = ? AND deleted_at = 0", internalMemoId).Count(&count)
	return count, res.Error
}

func (db *InternalMemoAttachmentConnection) FindInternalMemoAttachments(internalMemoId int) (internalMemoAttachmentOutput []model.InternalMemoAttachment, err error) {
	var (
		internalMemoAttachments []model.InternalMemoAttachment
	)
	res := db.connection.Where("internal_memo_id = ? AND deleted_at = 0", internalMemoId).Order("file_url").Find(&internalMemoAttachments)
	return internalMemoAttachments, res.Error
}

func (db *InternalMemoAttachmentConnection) FindInternalMemoAttachmentsOffset(limit int, offset int, order string, dir string, internalMemoId int) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error) {
	var (
		orderDirection          string
		internalMemoAttachments []model.SelectInternalMemoAttachmentParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("internal_memo_attachments").Select("internal_memo_attachments.id, internal_memo_attachments.internal_memo_id, internal_memo_attachments.file_name, internal_memo_attachments.description, internal_memo_attachments.file_url, internal_memo_attachments.remark, internal_memo_attachments.created_user_id, createdUID.username AS created_user, internal_memo_attachments.updated_user_id, updatedUID.username AS updated_user, internal_memo_attachments.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(internal_memo_attachments.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(internal_memo_attachments.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(internal_memo_attachments.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on internal_memo_attachments.created_user_id = createdUID.id").Joins("left join users updatedUID on internal_memo_attachments.updated_user_id = updatedUID.id").Joins("left join users deletedUID on internal_memo_attachments.deleted_user_id = deletedUID.id").Where("internal_memo_attachments.internal_memo_id = ? AND internal_memo_attachments.deleted_at = 0", internalMemoId).Order(orderDirection).Limit(limit).Offset(offset).Find(&internalMemoAttachments)
	return internalMemoAttachments, res.Error
}

func (db *InternalMemoAttachmentConnection) SearchInternalMemoAttachment(limit int, offset int, order string, dir string, search string, internalMemoId int) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error) {
	var (
		orderDirection          string
		final                   string
		internalMemoAttachments []model.SelectInternalMemoAttachmentParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("internal_memo_attachments").Select("internal_memo_attachments.id, internal_memo_attachments.internal_memo_id, internal_memo_attachments.file_name, internal_memo_attachments.description, internal_memo_attachments.file_url, internal_memo_attachments.remark, internal_memo_attachments.created_user_id, createdUID.username AS created_user, internal_memo_attachments.updated_user_id, updatedUID.username AS updated_user, internal_memo_attachments.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(internal_memo_attachments.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(internal_memo_attachments.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(internal_memo_attachments.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on internal_memo_attachments.created_user_id = createdUID.id").Joins("left join users updatedUID on internal_memo_attachments.updated_user_id = updatedUID.id").Joins("left join users deletedUID on internal_memo_attachments.deleted_user_id = deletedUID.id").Where("(lower(internal_memo_attachments.file_url) LIKE ? OR lower(internal_memo_attachments.description) LIKE ? OR lower(to_char(to_timestamp(internal_memo_attachments.created_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND internal_memo_attachments.internal_memo_id = ? AND internal_memo_attachments.deleted_at = 0", final, final, final, internalMemoId).Order(orderDirection).Limit(limit).Offset(offset).Find(&internalMemoAttachments)

	return internalMemoAttachments, res.Error
}

func (db *InternalMemoAttachmentConnection) CountSearchInternalMemoAttachment(search string, internalMemoId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("internal_memo_attachments").Select("internal_memo_attachments.id, internal_memo_attachments.internal_memo_id, internal_memo_attachments.file_name, internal_memo_attachments.description, internal_memo_attachments.file_url, internal_memo_attachments.remark, internal_memo_attachments.created_user_id, createdUID.username AS created_user, internal_memo_attachments.updated_user_id, updatedUID.username AS updated_user, internal_memo_attachments.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(internal_memo_attachments.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(internal_memo_attachments.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(internal_memo_attachments.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on internal_memo_attachments.created_user_id = createdUID.id").Joins("left join users updatedUID on internal_memo_attachments.updated_user_id = updatedUID.id").Joins("left join users deletedUID on internal_memo_attachments.deleted_user_id = deletedUID.id").Where("(lower(internal_memo_attachments.file_url) LIKE ? OR lower(internal_memo_attachments.description) LIKE ? OR lower(to_char(to_timestamp(internal_memo_attachments.created_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND internal_memo_attachments.internal_memo_id = ? AND internal_memo_attachments.deleted_at = 0", final, final, final, internalMemoId).Count(&count)
	return count, res.Error
}

func (db *InternalMemoAttachmentConnection) FindInternalMemoAttachmentById(id uint) (internalMemoAttachmentOutput model.SelectInternalMemoAttachmentParameter, err error) {
	var (
		internalMemoAttachment model.SelectInternalMemoAttachmentParameter
	)

	res := db.connection.Table("internal_memo_attachments").Select("internal_memo_attachments.id, internal_memo_attachments.internal_memo_id, internal_memo_attachments.file_name, internal_memo_attachments.description, internal_memo_attachments.file_url, internal_memo_attachments.remark, internal_memo_attachments.created_user_id, createdUID.username AS created_user, internal_memo_attachments.updated_user_id, updatedUID.username AS updated_user, internal_memo_attachments.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(internal_memo_attachments.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(internal_memo_attachments.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(internal_memo_attachments.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on internal_memo_attachments.created_user_id = createdUID.id").Joins("left join users updatedUID on internal_memo_attachments.updated_user_id = updatedUID.id").Joins("left join users deletedUID on internal_memo_attachments.deleted_user_id = deletedUID.id").Where("internal_memo_attachments.id=? AND internal_memo_attachments.deleted_at = 0", id).Take(&internalMemoAttachment)
	return internalMemoAttachment, res.Error
}

func (db *InternalMemoAttachmentConnection) FindExcInternalMemoAttachment(id uint) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error) {
	var (
		internalMemoAttachments []model.SelectInternalMemoAttachmentParameter
	)

	res := db.connection.Table("internal_memo_attachments").Select("internal_memo_attachments.id, internal_memo_attachments.internal_memo_id, internal_memo_attachments.file_name, internal_memo_attachments.description, internal_memo_attachments.file_url, internal_memo_attachments.remark, internal_memo_attachments.created_user_id, createdUID.username AS created_user, internal_memo_attachments.updated_user_id, updatedUID.username AS updated_user, internal_memo_attachments.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(internal_memo_attachments.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(internal_memo_attachments.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(internal_memo_attachments.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on internal_memo_attachments.created_user_id = createdUID.id").Joins("left join users updatedUID on internal_memo_attachments.updated_user_id = updatedUID.id").Joins("left join users deletedUID on internal_memo_attachments.deleted_user_id = deletedUID.id").Where(" internal_memo_attachments.id!=? AND internal_memo_attachments.deleted_at = 0", id).Order("internal_memo_attachments.file_url").Find(&internalMemoAttachments)
	return internalMemoAttachments, res.Error
}

func (db *InternalMemoAttachmentConnection) FindInternalMemoAttachmentByCompanyId(id uint) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error) {
	var (
		internalMemoAttachments []model.SelectInternalMemoAttachmentParameter
	)

	res := db.connection.Table("internal_memo_attachments").Select("internal_memo_attachments.id, internal_memo_attachments.internal_memo_id, internal_memo_attachments.file_name, internal_memo_attachments.description, internal_memo_attachments.file_url, internal_memo_attachments.remark, internal_memo_attachments.created_user_id, createdUID.username AS created_user, internal_memo_attachments.updated_user_id, updatedUID.username AS updated_user, internal_memo_attachments.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(internal_memo_attachments.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(internal_memo_attachments.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(internal_memo_attachments.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on internal_memo_attachments.created_user_id = createdUID.id").Joins("left join users updatedUID on internal_memo_attachments.updated_user_id = updatedUID.id").Joins("left join users deletedUID on internal_memo_attachments.deleted_user_id = deletedUID.id").Where("internal_memo_attachments.internal_memo_id=? AND internal_memo_attachments.deleted_at = 0", id).Order("internal_memo_attachments.file_url").Find(&internalMemoAttachments)
	return internalMemoAttachments, res.Error
}

func (db *InternalMemoAttachmentConnection) InsertInternalMemoAttachment(internalMemoAttachment model.InternalMemoAttachment) (internalMemoAttachmentOutput model.InternalMemoAttachment, err error) {
	res := db.connection.Save(&internalMemoAttachment)
	return internalMemoAttachment, res.Error
}

func (db *InternalMemoAttachmentConnection) UpdateInternalMemoAttachment(internalMemoAttachment model.InternalMemoAttachment, id uint) (internalMemoAttachmentOutput model.InternalMemoAttachment, err error) {
	res := db.connection.Where("id=?", id).Updates(&internalMemoAttachment)
	return internalMemoAttachment, res.Error
}
