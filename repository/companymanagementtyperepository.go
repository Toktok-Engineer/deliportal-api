package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type CompanyManagementTypeRepository interface {
	FindCompanyManagementTypes() (companyManagementTypeOutput []model.CompanyManagementType, err error)
	FindCompanyManagementTypeById(id uint) (companyManagementTypeOutput model.CompanyManagementType, err error)
	FindExcCompanyManagementType(id uint) (companyManagementTypeOutput []model.CompanyManagementType, err error)
	InsertCompanyManagementType(companyManagementType model.CompanyManagementType) (companyManagementTypeOutput model.CompanyManagementType, err error)
	UpdateCompanyManagementType(companyManagementType model.CompanyManagementType, id uint) (companyManagementTypeOutput model.CompanyManagementType, err error)
}

type CompanyManagementTypeConnection struct {
	connection *gorm.DB
}

func NewCompanyManagementTypeRepository(db *gorm.DB) CompanyManagementTypeRepository {
	return &CompanyManagementTypeConnection{
		connection: db,
	}
}

func (db *CompanyManagementTypeConnection) FindCompanyManagementTypes() (companyManagementTypeOutput []model.CompanyManagementType, err error) {
	var (
		companyManagementTypes []model.CompanyManagementType
	)
	res := db.connection.Where("deleted_at = 0").Order("company_management_type_name").Find(&companyManagementTypes)
	return companyManagementTypes, res.Error
}

func (db *CompanyManagementTypeConnection) FindCompanyManagementTypeById(id uint) (companyManagementTypeOutput model.CompanyManagementType, err error) {
	var (
		companyManagementType model.CompanyManagementType
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&companyManagementType)
	return companyManagementType, res.Error
}

func (db *CompanyManagementTypeConnection) FindExcCompanyManagementType(id uint) (companyManagementTypeOutput []model.CompanyManagementType, err error) {
	var (
		companyManagementTypes []model.CompanyManagementType
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("company_management_type_name").Find(&companyManagementTypes)
	return companyManagementTypes, res.Error
}

func (db *CompanyManagementTypeConnection) InsertCompanyManagementType(companyManagementType model.CompanyManagementType) (companyManagementTypeOutput model.CompanyManagementType, err error) {
	res := db.connection.Save(&companyManagementType)
	return companyManagementType, res.Error
}

func (db *CompanyManagementTypeConnection) UpdateCompanyManagementType(companyManagementType model.CompanyManagementType, id uint) (companyManagementTypeOutput model.CompanyManagementType, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyManagementType)
	return companyManagementType, res.Error
}
