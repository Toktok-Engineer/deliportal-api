package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type LicenseTypeRepository interface {
	FindLicenseTypes() (licenseTypeOutput []model.LicenseType, err error)
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

func (db *LicenseTypeConnection) FindLicenseTypes() (licenseTypeOutput []model.LicenseType, err error) {
	var (
		licenseTypes []model.LicenseType
	)
	res := db.connection.Where("deleted_at = 0").Order("license_type_name").Find(&licenseTypes)
	return licenseTypes, res.Error
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
