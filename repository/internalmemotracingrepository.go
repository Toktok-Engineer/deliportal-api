package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type InternalMemoTracingRepository interface {
	CountInternalMemoTracingAll(internalMemoId int) (count int64, err error)
	FindInternalMemoTracings(internalMemoId int) (internalmemotracingOutput []model.InternalMemoTracing, err error)
	FindInternalMemoTracingById(id uint) (internalMemoTracingOutput model.InternalMemoTracing, err error)
	FindExcInternalMemoTracing(id uint) (internalMemoTracingOutput []model.InternalMemoTracing, err error)
	InsertInternalMemoTracing(internalMemoTracing model.InternalMemoTracing) (internalMemoTracingOutput model.InternalMemoTracing, err error)
	UpdateInternalMemoTracing(internalMemoTracing model.InternalMemoTracing, id uint) (internalMemoTracingOutput model.InternalMemoTracing, err error)
}

type internalMemoTracingConnection struct {
	connection *gorm.DB
}

func NewInternalMemoTracingRepository(db *gorm.DB) InternalMemoTracingRepository {
	return &internalMemoTracingConnection{
		connection: db,
	}
}

func (db *internalMemoTracingConnection) CountInternalMemoTracingAll(internalMemoId int) (count int64, err error) {
	res := db.connection.Debug().Table("internal_memo_tracings").Where("internal_memo_id = ? AND deleted_at = 0", internalMemoId).Count(&count)
	return count, res.Error
}

func (db *internalMemoTracingConnection) FindInternalMemoTracings(internalMemoId int) (internalmemotracingOutput []model.InternalMemoTracing, err error) {
	var (
		internalmemotracings []model.InternalMemoTracing
	)
	res := db.connection.Where("internal_memo_id = ? AND deleted_at = 0", internalMemoId).Order("sequence_no DESC").Find(&internalmemotracings)
	return internalmemotracings, res.Error
}

func (db *internalMemoTracingConnection) FindInternalMemoTracingById(id uint) (internalMemoTracingOutput model.InternalMemoTracing, err error) {
	var (
		internalMemoTracing model.InternalMemoTracing
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&internalMemoTracing)
	return internalMemoTracing, res.Error
}

func (db *internalMemoTracingConnection) FindExcInternalMemoTracing(id uint) (internalMemoTracingOutput []model.InternalMemoTracing, err error) {
	var (
		internalMemoTracings []model.InternalMemoTracing
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("sequence_no").Find(&internalMemoTracings)
	return internalMemoTracings, res.Error
}

func (db *internalMemoTracingConnection) InsertInternalMemoTracing(internalMemoTracing model.InternalMemoTracing) (internalMemoTracingOutput model.InternalMemoTracing, err error) {
	res := db.connection.Save(&internalMemoTracing)
	return internalMemoTracing, res.Error
}

func (db *internalMemoTracingConnection) UpdateInternalMemoTracing(internalMemoTracing model.InternalMemoTracing, id uint) (internalMemoTracingOutput model.InternalMemoTracing, err error) {
	res := db.connection.Where("id=?", id).Updates(&internalMemoTracing)
	return internalMemoTracing, res.Error
}
