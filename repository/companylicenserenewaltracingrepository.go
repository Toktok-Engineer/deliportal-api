package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type CompanyLicenseRenewalTracingRepository interface {
	FindCompanyLicenseRenewalTracings() (companyLicenseRenewalTracingOutput []model.CompanyLicenseRenewalTracing, err error)
	FindCompanyLicenseRenewalTracingById(id uint) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error)
	FindExcCompanyLicenseRenewalTracing(id uint) (companyLicenseRenewalTracingOutput []model.CompanyLicenseRenewalTracing, err error)
	InsertCompanyLicenseRenewalTracing(companyLicenseRenewalTracing model.CompanyLicenseRenewalTracing) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error)
	UpdateCompanyLicenseRenewalTracing(companyLicenseRenewalTracing model.CompanyLicenseRenewalTracing, id uint) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error)
}

type CompanyLicenseRenewalTracingConnection struct {
	connection *gorm.DB
}

func NewCompanyLicenseRenewalTracingRepository(db *gorm.DB) CompanyLicenseRenewalTracingRepository {
	return &CompanyLicenseRenewalTracingConnection{
		connection: db,
	}
}

func (db *CompanyLicenseRenewalTracingConnection) FindCompanyLicenseRenewalTracings() (companyLicenseRenewalTracingOutput []model.CompanyLicenseRenewalTracing, err error) {
	var (
		companyLicenseRenewalTracings []model.CompanyLicenseRenewalTracing
	)
	res := db.connection.Where("deleted_at = 0").Order("company_license_id").Find(&companyLicenseRenewalTracings)
	return companyLicenseRenewalTracings, res.Error
}

func (db *CompanyLicenseRenewalTracingConnection) FindCompanyLicenseRenewalTracingById(id uint) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error) {
	var (
		companyLicenseRenewalTracing model.CompanyLicenseRenewalTracing
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&companyLicenseRenewalTracing)
	return companyLicenseRenewalTracing, res.Error
}

func (db *CompanyLicenseRenewalTracingConnection) FindExcCompanyLicenseRenewalTracing(id uint) (companyLicenseRenewalTracingOutput []model.CompanyLicenseRenewalTracing, err error) {
	var (
		companyLicenseRenewalTracings []model.CompanyLicenseRenewalTracing
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("company_license_id").Find(&companyLicenseRenewalTracings)
	return companyLicenseRenewalTracings, res.Error
}

func (db *CompanyLicenseRenewalTracingConnection) InsertCompanyLicenseRenewalTracing(companyLicenseRenewalTracing model.CompanyLicenseRenewalTracing) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error) {
	res := db.connection.Save(&companyLicenseRenewalTracing)
	return companyLicenseRenewalTracing, res.Error
}

func (db *CompanyLicenseRenewalTracingConnection) UpdateCompanyLicenseRenewalTracing(companyLicenseRenewalTracing model.CompanyLicenseRenewalTracing, id uint) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyLicenseRenewalTracing)
	return companyLicenseRenewalTracing, res.Error
}
