package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type LicenseTypeRepository interface {
	CountLicenseTypeAll() (count int64, err error)
	FindLicenseTypes() (licenseTypeOutput []model.LicenseType, err error)
	FindLicenseTypesOffset(limit int, offset int, order string, dir string) (licenseTypeOutput []model.LicenseType, err error)
	SearchLicenseType(limit int, offset int, order string, dir string, search string) (licenseTypeOutput []model.LicenseType, err error)
	CountSearchLicenseType(search string) (count int64, err error)
	FindLicenseTypeById(id uint) (licenseTypeOutput model.LicenseType, err error)
	FindExcLicenseType(id uint) (licenseTypeOutput []model.LicenseType, err error)
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

func (db *LicenseTypeConnection) FindLicenseTypes() (licenseTypeOutput []model.LicenseType, err error) {
	var (
		licenseTypes []model.LicenseType
	)
	res := db.connection.Where("deleted_at = 0").Order("license_type_name").Find(&licenseTypes)
	return licenseTypes, res.Error
}

func (db *LicenseTypeConnection) FindLicenseTypesOffset(limit int, offset int, order string, dir string) (licenseTypeOutput []model.LicenseType, err error) {
	var (
		orderDirection string
		licenseTypes   []model.LicenseType
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&licenseTypes)
	return licenseTypes, res.Error
}

func (db *LicenseTypeConnection) SearchLicenseType(limit int, offset int, order string, dir string, search string) (licenseTypeOutput []model.LicenseType, err error) {
	var (
		orderDirection string
		final          string
		licenseTypes   []model.LicenseType
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(license_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&licenseTypes)
	return licenseTypes, res.Error
}

func (db *LicenseTypeConnection) CountSearchLicenseType(search string) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("license_types").Where("(lower(license_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
}

func (db *LicenseTypeConnection) FindLicenseTypeById(id uint) (licenseTypeOutput model.LicenseType, err error) {
	var (
		licenseType model.LicenseType
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&licenseType)
	return licenseType, res.Error
}

func (db *LicenseTypeConnection) FindExcLicenseType(id uint) (licenseTypeOutput []model.LicenseType, err error) {
	var (
		licenseTypes []model.LicenseType
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("license_type_name").Find(&licenseTypes)
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
