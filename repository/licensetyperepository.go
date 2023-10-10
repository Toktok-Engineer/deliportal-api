package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type LicenseTypeRepository interface {
	CountLicenseTypeAll() (count int64, err error)
	FindLicenseTypes() (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	FindLicenseTypesOffset(limit int, offset int, order string, dir string) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	SearchLicenseType(limit int, offset int, order string, dir string, search string) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	CountSearchLicenseType(search string) (count int64, err error)
	FindLicenseTypeById(id uint) (licenseTypeOutput model.SelectLicenseTypeParameter, err error)
	FindExcLicenseType(id uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	FindExcCompleteLicenseType(groupLT uint, id uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	FindLicenseTypeByGroupLT(groupLT uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	InsertLicenseType(licenseType model.LicenseType) (licenseTypeOutput model.LicenseType, err error)
	UpdateLicenseType(licenseType model.LicenseType, id uint) (licenseTypeOutput model.LicenseType, err error)
}

type LicenseTypeConnection struct {
	connection *gorm.DB
}

func NewLicenseTypeRepository(db *gorm.DB) LicenseTypeRepository {
	return &LicenseTypeConnection{
		connection: db,
	}
}

func (db *LicenseTypeConnection) CountLicenseTypeAll() (count int64, err error) {
	res := db.connection.Debug().Table("license_types").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *LicenseTypeConnection) FindLicenseTypes() (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	var (
		licenseTypes []model.SelectLicenseTypeParameter
	)
	res := db.connection.Debug().Table("license_types").Select("license_types.id, license_types.license_type_name, license_types.reminder_before_month, license_types.management_reminder_before_month, license_types.reminder_frequency_day, license_types.management_reminder_frequency_day, license_types.group_license_type_id, group_license_types.group_license_type_name, license_types.remark, license_types.created_user_id, license_types.updated_user_id, license_types.deleted_user_id, license_types.created_at, license_types.updated_at, license_types.deleted_at").Joins("left join group_license_types ON license_types.group_license_type_id = group_license_types.id").Where("license_types.deleted_at = 0").Order("license_types.license_type_name").Find(&licenseTypes)
	return licenseTypes, res.Error
}

func (db *LicenseTypeConnection) FindLicenseTypesOffset(limit int, offset int, order string, dir string) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	var (
		orderDirection string
		licenseTypes   []model.SelectLicenseTypeParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("license_types").Select("license_types.id, license_types.license_type_name, license_types.reminder_before_month, license_types.management_reminder_before_month, license_types.reminder_frequency_day, license_types.management_reminder_frequency_day, license_types.group_license_type_id, group_license_types.group_license_type_name, license_types.remark, license_types.created_user_id, license_types.updated_user_id, license_types.deleted_user_id, license_types.created_at, license_types.updated_at, license_types.deleted_at").Joins("left join group_license_types ON license_types.group_license_type_id = group_license_types.id").Where("license_types.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&licenseTypes)
	return licenseTypes, res.Error
}

func (db *LicenseTypeConnection) SearchLicenseType(limit int, offset int, order string, dir string, search string) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	var (
		orderDirection string
		final          string
		licenseTypes   []model.SelectLicenseTypeParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("license_types").Select("license_types.id, license_types.license_type_name, license_types.reminder_before_month, license_types.management_reminder_before_month, license_types.reminder_frequency_day, license_types.management_reminder_frequency_day, license_types.group_license_type_id, group_license_types.group_license_type_name, license_types.remark, license_types.created_user_id, license_types.updated_user_id, license_types.deleted_user_id, license_types.created_at, license_types.updated_at, license_types.deleted_at").Joins("left join group_license_types ON license_types.group_license_type_id = group_license_types.id").Where("(lower(group_license_types.group_license_type_name) LIKE ? OR lower(license_types.license_type_name) LIKE ? OR lower(license_types.remark) LIKE ?) AND license_types.deleted_at = 0", final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&licenseTypes)
	return licenseTypes, res.Error
}

func (db *LicenseTypeConnection) CountSearchLicenseType(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("license_types").Select("license_types.id, license_types.license_type_name, license_types.reminder_before_month, license_types.management_reminder_before_month, license_types.reminder_frequency_day, license_types.management_reminder_frequency_day, license_types.group_license_type_id, group_license_types.group_license_type_name, license_types.remark, license_types.created_user_id, license_types.updated_user_id, license_types.deleted_user_id, license_types.created_at, license_types.updated_at, license_types.deleted_at").Joins("left join group_license_types ON license_types.group_license_type_id = group_license_types.id").Where("(lower(group_license_types.group_license_type_name) LIKE ? OR lower(license_types.license_type_name) LIKE ? OR lower(license_types.remark) LIKE ?) AND license_types.deleted_at = 0", final, final, final).Count(&count)
	return count, res.Error
}

func (db *LicenseTypeConnection) FindLicenseTypeById(id uint) (licenseTypeOutput model.SelectLicenseTypeParameter, err error) {
	var (
		licenseType model.SelectLicenseTypeParameter
	)
	res := db.connection.Debug().Table("license_types").Select("license_types.id, license_types.license_type_name, license_types.reminder_before_month, license_types.management_reminder_before_month, license_types.reminder_frequency_day, license_types.management_reminder_frequency_day, license_types.group_license_type_id, group_license_types.group_license_type_name, license_types.remark, license_types.created_user_id, license_types.updated_user_id, license_types.deleted_user_id, license_types.created_at, license_types.updated_at, license_types.deleted_at").Joins("left join group_license_types ON license_types.group_license_type_id = group_license_types.id").Where("license_types.id=? AND license_types.deleted_at = 0", id).Take(&licenseType)
	return licenseType, res.Error
}

func (db *LicenseTypeConnection) FindExcLicenseType(id uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	var (
		licenseTypes []model.SelectLicenseTypeParameter
	)
	res := db.connection.Debug().Table("license_types").Select("license_types.id, license_types.license_type_name, license_types.reminder_before_month, license_types.management_reminder_before_month, license_types.reminder_frequency_day, license_types.management_reminder_frequency_day, license_types.group_license_type_id, group_license_types.group_license_type_name, license_types.remark, license_types.created_user_id, license_types.updated_user_id, license_types.deleted_user_id, license_types.created_at, license_types.updated_at, license_types.deleted_at").Joins("left join group_license_types ON license_types.group_license_type_id = group_license_types.id").Where("license_types.id!=? AND license_types.deleted_at = 0", id).Order("license_types.license_type_name").Find(&licenseTypes)
	return licenseTypes, res.Error
}

func (db *LicenseTypeConnection) FindExcCompleteLicenseType(groupLT uint, id uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	var (
		licenseTypes []model.SelectLicenseTypeParameter
	)

	res := db.connection.Debug().Table("license_types").Select("license_types.id, license_types.license_type_name, license_types.reminder_before_month, license_types.management_reminder_before_month, license_types.reminder_frequency_day, license_types.management_reminder_frequency_day, license_types.group_license_type_id, group_license_types.group_license_type_name, license_types.remark, license_types.created_user_id, license_types.updated_user_id, license_types.deleted_user_id, license_types.created_at, license_types.updated_at, license_types.deleted_at").Joins("left join group_license_types ON license_types.group_license_type_id = group_license_types.id").Where("license_types.group_license_type_id = ? AND license_types.id!=? AND license_types.deleted_at = 0", groupLT, id).Order("license_types.license_type_name").Find(&licenseTypes)
	return licenseTypes, res.Error
}

func (db *LicenseTypeConnection) FindLicenseTypeByGroupLT(groupLT uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	var (
		licenseTypes []model.SelectLicenseTypeParameter
	)

	res := db.connection.Debug().Table("license_types").Select("license_types.id, license_types.license_type_name, license_types.reminder_before_month, license_types.management_reminder_before_month, license_types.reminder_frequency_day, license_types.management_reminder_frequency_day, license_types.group_license_type_id, group_license_types.group_license_type_name, license_types.remark, license_types.created_user_id, license_types.updated_user_id, license_types.deleted_user_id, license_types.created_at, license_types.updated_at, license_types.deleted_at").Joins("left join group_license_types ON license_types.group_license_type_id = group_license_types.id").Where("license_types.group_license_type_id=? AND license_types.deleted_at = 0", groupLT).Order("license_types.license_type_name").Find(&licenseTypes)
	return licenseTypes, res.Error
}

func (db *LicenseTypeConnection) InsertLicenseType(licenseType model.LicenseType) (licenseTypeOutput model.LicenseType, err error) {
	res := db.connection.Save(&licenseType)
	return licenseType, res.Error
}

func (db *LicenseTypeConnection) UpdateLicenseType(licenseType model.LicenseType, id uint) (licenseTypeOutput model.LicenseType, err error) {
	res := db.connection.Where("id=?", id).Updates(&licenseType)
	return licenseType, res.Error
}
