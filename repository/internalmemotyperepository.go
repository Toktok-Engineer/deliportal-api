package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type InternalMemoTypeRepository interface {
	CountInternalMemoTypeAll() (count int64, err error)
	FindInternalMemoTypes() (businessunitOutput []model.InternalMemoType, err error)
	FindInternalMemoTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.InternalMemoType, err error)
	SearchInternalMemoType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.InternalMemoType, err error)
	CountSearchInternalMemoType(search string) (count int64, err error)
	FindInternalMemoTypeById(id uint) (internalMemoTypeOutput model.InternalMemoType, err error)
	FindExcInternalMemoType(id uint) (internalMemoTypeOutput []model.InternalMemoType, err error)
	InsertInternalMemoType(internalMemoType model.InternalMemoType) (internalMemoTypeOutput model.InternalMemoType, err error)
	UpdateInternalMemoType(internalMemoType model.InternalMemoType, id uint) (internalMemoTypeOutput model.InternalMemoType, err error)
}

type internalMemoTypeConnection struct {
	connection *gorm.DB
}

func NewInternalMemoTypeRepository(db *gorm.DB) InternalMemoTypeRepository {
	return &internalMemoTypeConnection{
		connection: db,
	}
}
func (db *internalMemoTypeConnection) CountInternalMemoTypeAll() (count int64, err error) {
	res := db.connection.Debug().Table("internal_memo_types").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *internalMemoTypeConnection) FindInternalMemoTypes() (businessunitOutput []model.InternalMemoType, err error) {
	var (
		businessunits []model.InternalMemoType
	)
	res := db.connection.Where("deleted_at = 0").Order("internal_memo_type_name").Find(&businessunits)
	return businessunits, res.Error
}

func (db *internalMemoTypeConnection) FindInternalMemoTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.InternalMemoType, err error) {
	var (
		orderDirection string
		businessunits  []model.InternalMemoType
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&businessunits)
	return businessunits, res.Error
}

func (db *internalMemoTypeConnection) SearchInternalMemoType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.InternalMemoType, err error) {
	var (
		orderDirection string
		final          string
		businessunits  []model.InternalMemoType
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(internal_memo_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&businessunits)
	return businessunits, res.Error
}

func (db *internalMemoTypeConnection) CountSearchInternalMemoType(search string) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("internal_memo_types").Where("(lower(internal_memo_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
}

func (db *internalMemoTypeConnection) FindInternalMemoTypeById(id uint) (internalMemoTypeOutput model.InternalMemoType, err error) {
	var (
		internalMemoType model.InternalMemoType
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&internalMemoType)
	return internalMemoType, res.Error
}

func (db *internalMemoTypeConnection) FindExcInternalMemoType(id uint) (internalMemoTypeOutput []model.InternalMemoType, err error) {
	var (
		internalMemoTypes []model.InternalMemoType
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("internal_memo_type_name").Find(&internalMemoTypes)
	return internalMemoTypes, res.Error
}

func (db *internalMemoTypeConnection) InsertInternalMemoType(internalMemoType model.InternalMemoType) (internalMemoTypeOutput model.InternalMemoType, err error) {
	res := db.connection.Save(&internalMemoType)
	return internalMemoType, res.Error
}

func (db *internalMemoTypeConnection) UpdateInternalMemoType(internalMemoType model.InternalMemoType, id uint) (internalMemoTypeOutput model.InternalMemoType, err error) {
	res := db.connection.Where("id=?", id).Updates(&internalMemoType)
	return internalMemoType, res.Error
}
