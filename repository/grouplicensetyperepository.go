package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type GroupLicenseTypeRepository interface {
	CountGroupLicenseTypeAll() (count int64, err error)
	FindGroupLicenseTypes() (businessunitOutput []model.GroupLicenseType, err error)
	FindGroupLicenseTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.GroupLicenseType, err error)
	SearchGroupLicenseType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.GroupLicenseType, err error)
	CountSearchGroupLicenseType(search string) (count int64, err error)
	FindGroupLicenseTypeById(id uint) (groupLicenseTypeOutput model.GroupLicenseType, err error)
	FindExcGroupLicenseType(id uint) (groupLicenseTypeOutput []model.GroupLicenseType, err error)
	InsertGroupLicenseType(groupLicenseType model.GroupLicenseType) (groupLicenseTypeOutput model.GroupLicenseType, err error)
	UpdateGroupLicenseType(groupLicenseType model.GroupLicenseType, id uint) (groupLicenseTypeOutput model.GroupLicenseType, err error)
}

type groupLicenseTypeConnection struct {
	connection *gorm.DB
}

func NewGroupLicenseTypeRepository(db *gorm.DB) GroupLicenseTypeRepository {
	return &groupLicenseTypeConnection{
		connection: db,
	}
}
func (db *groupLicenseTypeConnection) CountGroupLicenseTypeAll() (count int64, err error) {
	res := db.connection.Debug().Table("group_license_types").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *groupLicenseTypeConnection) FindGroupLicenseTypes() (businessunitOutput []model.GroupLicenseType, err error) {
	var (
		businessunits []model.GroupLicenseType
	)
	res := db.connection.Where("deleted_at = 0").Order("group_license_type_name").Find(&businessunits)
	return businessunits, res.Error
}

func (db *groupLicenseTypeConnection) FindGroupLicenseTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.GroupLicenseType, err error) {
	var (
		orderDirection string
		businessunits  []model.GroupLicenseType
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&businessunits)
	return businessunits, res.Error
}

func (db *groupLicenseTypeConnection) SearchGroupLicenseType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.GroupLicenseType, err error) {
	var (
		orderDirection string
		final          string
		businessunits  []model.GroupLicenseType
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(group_license_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&businessunits)
	return businessunits, res.Error
}

func (db *groupLicenseTypeConnection) CountSearchGroupLicenseType(search string) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("group_license_types").Where("(lower(group_license_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
}

func (db *groupLicenseTypeConnection) FindGroupLicenseTypeById(id uint) (groupLicenseTypeOutput model.GroupLicenseType, err error) {
	var (
		groupLicenseType model.GroupLicenseType
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&groupLicenseType)
	return groupLicenseType, res.Error
}

func (db *groupLicenseTypeConnection) FindExcGroupLicenseType(id uint) (groupLicenseTypeOutput []model.GroupLicenseType, err error) {
	var (
		groupLicenseTypes []model.GroupLicenseType
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("group_license_type_name").Find(&groupLicenseTypes)
	return groupLicenseTypes, res.Error
}

func (db *groupLicenseTypeConnection) InsertGroupLicenseType(groupLicenseType model.GroupLicenseType) (groupLicenseTypeOutput model.GroupLicenseType, err error) {
	res := db.connection.Save(&groupLicenseType)
	return groupLicenseType, res.Error
}

func (db *groupLicenseTypeConnection) UpdateGroupLicenseType(groupLicenseType model.GroupLicenseType, id uint) (groupLicenseTypeOutput model.GroupLicenseType, err error) {
	res := db.connection.Where("id=?", id).Updates(&groupLicenseType)
	return groupLicenseType, res.Error
}
