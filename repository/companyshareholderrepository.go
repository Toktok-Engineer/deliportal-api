package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type CompanyShareholderRepository interface {
	FindCompanyShareholders() (companyShareholderOutput []model.CompanyShareholder, err error)
	FindCompanyShareholderByCompanyId(id uint) (companyShareholderOutput []model.CompanyShareholder, err error)
	FindCompanyShareholderById(id uint) (companyShareholderOutput model.CompanyShareholder, err error)
	FindExcCompanyShareholder(id uint) (companyShareholderOutput []model.CompanyShareholder, err error)
	InsertCompanyShareholder(companyShareholder model.CompanyShareholder) (companyShareholderOutput model.CompanyShareholder, err error)
	UpdateCompanyShareholder(companyShareholder model.CompanyShareholder, id uint) (companyShareholderOutput model.CompanyShareholder, err error)
}

type CompanyShareholderConnection struct {
	connection *gorm.DB
}

func NewCompanyShareholderRepository(db *gorm.DB) CompanyShareholderRepository {
	return &CompanyShareholderConnection{
		connection: db,
	}
}

func (db *CompanyShareholderConnection) FindCompanyShareholders() (companyShareholderOutput []model.CompanyShareholder, err error) {
	var (
		companyShareholders []model.CompanyShareholder
	)
	res := db.connection.Where("deleted_at = 0").Order("shareholder_name").Find(&companyShareholders)
	return companyShareholders, res.Error
}

func (db *CompanyShareholderConnection) FindCompanyShareholderByCompanyId(id uint) (companyShareholderOutput []model.CompanyShareholder, err error) {
	var (
		companyShareholders []model.CompanyShareholder
	)
	res := db.connection.Where("company_id=? AND deleted_at = 0", id).Order("shareholder_name").Find(&companyShareholders)
	return companyShareholders, res.Error
}

func (db *CompanyShareholderConnection) FindCompanyShareholderById(id uint) (companyShareholderOutput model.CompanyShareholder, err error) {
	var (
		companyShareholder model.CompanyShareholder
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&companyShareholder)
	return companyShareholder, res.Error
}

func (db *CompanyShareholderConnection) FindExcCompanyShareholder(id uint) (companyShareholderOutput []model.CompanyShareholder, err error) {
	var (
		companyShareholders []model.CompanyShareholder
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("shareholder_name").Find(&companyShareholders)
	return companyShareholders, res.Error
}

func (db *CompanyShareholderConnection) InsertCompanyShareholder(companyShareholder model.CompanyShareholder) (companyShareholderOutput model.CompanyShareholder, err error) {
	res := db.connection.Save(&companyShareholder)
	return companyShareholder, res.Error
}

func (db *CompanyShareholderConnection) UpdateCompanyShareholder(companyShareholder model.CompanyShareholder, id uint) (companyShareholderOutput model.CompanyShareholder, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyShareholder)
	return companyShareholder, res.Error
}
