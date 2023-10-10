package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type NonWorkingTypeRepository interface {
	CountNonWorkingTypeAll() (count int64, err error)
	FindNonWorkingType() (nonworkingtypeOutput []model.NonWorkingType, err error)
	FindNonWorkingTypeOffset(limit int, offset int, order string, dir string) (nonworkingtypeOutput []model.NonWorkingType, err error)
	SearchNonWorkingType(limit int, offset int, order string, dir string, search string) (nonworkingtypeOutput []model.NonWorkingType, err error)
	CountSearchNonWorkingType(search string) (count int64, err error)
	FindNonWorkingTypeById(id uint) (nonworkingtypeOutput model.NonWorkingType, err error)
	CountNonWorkingTypeName(search string) (count int64, err error)
	FindExcNonWorkingType(id uint) (nonworkingtypeOutput []model.NonWorkingType, err error)
	InsertNonWorkingType(nonworkingtype model.NonWorkingType) (nonworkingtypeOutput model.NonWorkingType, err error)
	UpdateNonWorkingType(nonworkingtype model.NonWorkingType, id uint) (nonworkingtypeOutput model.NonWorkingType, err error)
	DeleteNonWorkingType(nonworkingtype model.NonWorkingType, id uint) (nonworkingtypeOutput model.NonWorkingType, err error)
}

type NonWorkingTypeConnection struct {
	connection *gorm.DB
}

func NewNonWorkingTypeRepository(db *gorm.DB) NonWorkingTypeRepository {
	return &NonWorkingTypeConnection{
		connection: db,
	}
}

func (db *NonWorkingTypeConnection) CountNonWorkingTypeAll() (count int64, err error) {
	res := db.connection.Debug().Table("non_working_types").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *NonWorkingTypeConnection) FindNonWorkingType() (nonworkingtypeOutput []model.NonWorkingType, err error) {
	var (
		nonworkingtype []model.NonWorkingType
	)
	res := db.connection.Where("deleted_at = 0").Order("non_working_type_name").Find(&nonworkingtype)
	return nonworkingtype, res.Error
}

func (db *NonWorkingTypeConnection) FindNonWorkingTypeOffset(limit int, offset int, order string, dir string) (nonworkingtypeOutput []model.NonWorkingType, err error) {
	var (
		orderDirection string
		nonworkingtype []model.NonWorkingType
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&nonworkingtype)
	return nonworkingtype, res.Error
}

func (db *NonWorkingTypeConnection) SearchNonWorkingType(limit int, offset int, order string, dir string, search string) (nonworkingtypeOutput []model.NonWorkingType, err error) {
	var (
		orderDirection string
		final          string
		nonworkingtype []model.NonWorkingType
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(non_working_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&nonworkingtype)
	return nonworkingtype, res.Error
}

func (db *NonWorkingTypeConnection) CountSearchNonWorkingType(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("non_working_types").Where("(lower(non_working_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
}

func (db *NonWorkingTypeConnection) FindNonWorkingTypeById(id uint) (nonworkingtypeOutput model.NonWorkingType, err error) {
	var (
		nonworkingtype model.NonWorkingType
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&nonworkingtype)
	return nonworkingtype, res.Error
}

func (db *NonWorkingTypeConnection) CountNonWorkingTypeName(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("non_working_types").Where("lower(non_working_type_name) LIKE ?", final).Count(&count)
	return count, res.Error
}

func (db *NonWorkingTypeConnection) FindExcNonWorkingType(id uint) (nonworkingtypeOutput []model.NonWorkingType, err error) {
	var (
		nonworkingtype []model.NonWorkingType
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("non_working_type_name").Find(&nonworkingtype)
	return nonworkingtype, res.Error
}

func (db *NonWorkingTypeConnection) InsertNonWorkingType(nonworkingtype model.NonWorkingType) (nonworkingtypeOutput model.NonWorkingType, err error) {
	res := db.connection.Save(&nonworkingtype)
	return nonworkingtype, res.Error
}

func (db *NonWorkingTypeConnection) UpdateNonWorkingType(nonworkingtype model.NonWorkingType, id uint) (nonworkingtypeOutput model.NonWorkingType, err error) {
	var (
		nonworking_type model.NonWorkingType
	)
	res := db.connection.Model(&nonworking_type).Where("id=?", id).Updates(map[string]interface{}{"non_working_type_name": nonworkingtype.NonWorkingTypeName, "deduct_leave": nonworkingtype.DeductLeave, "remark": nonworkingtype.Remark, "updated_user_id": nonworkingtype.UpdatedUserID, "updated_at": nonworkingtype.UpdatedAt})
	return nonworkingtype, res.Error
}

func (db *NonWorkingTypeConnection) DeleteNonWorkingType(nonworkingtype model.NonWorkingType, id uint) (nonworkingtypeOutput model.NonWorkingType, err error) {
	var (
		nonworking_type model.NonWorkingType
	)
	res := db.connection.Model(&nonworking_type).Where("id=?", id).Updates(map[string]interface{}{"non_working_type_name": nonworkingtype.NonWorkingTypeName, "deleted_user_id": nonworkingtype.DeletedUserID, "deleted_at": nonworkingtype.DeletedAt})
	return nonworkingtype, res.Error
}
