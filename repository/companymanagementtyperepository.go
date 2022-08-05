package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyManagementTypeRepository interface {
	CountCompanyManagementTypeAll() (count int64, err error)
	FindCompanyManagementTypes() (companyManagementTypeOutput []model.CompanyManagementType, err error)
	FindCompanyManagementTypesOffset(limit int, offset int, order string, dir string) (companyManagementTypeOutput []model.CompanyManagementType, err error)
	SearchCompanyManagementType(limit int, offset int, order string, dir string, search string) (companyManagementTypeOutput []model.CompanyManagementType, err error)
	CountSearchCompanyManagementType(search string) (count int64, err error)
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

func (db *CompanyManagementTypeConnection) CountCompanyManagementTypeAll() (count int64, err error) {
	res := db.connection.Debug().Table("company_management_types").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *CompanyManagementTypeConnection) FindCompanyManagementTypes() (companyManagementTypeOutput []model.CompanyManagementType, err error) {
	var (
		companyManagementTypes []model.CompanyManagementType
	)
	res := db.connection.Where("deleted_at = 0").Order("company_management_type_name").Find(&companyManagementTypes)
	return companyManagementTypes, res.Error
}

func (db *CompanyManagementTypeConnection) FindCompanyManagementTypesOffset(limit int, offset int, order string, dir string) (companyManagementTypeOutput []model.CompanyManagementType, err error) {
	var (
		orderDirection         string
		companyManagementTypes []model.CompanyManagementType
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&companyManagementTypes)
	return companyManagementTypes, res.Error
}

func (db *CompanyManagementTypeConnection) SearchCompanyManagementType(limit int, offset int, order string, dir string, search string) (companyManagementTypeOutput []model.CompanyManagementType, err error) {
	var (
		orderDirection         string
		final                  string
		companyManagementTypes []model.CompanyManagementType
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(company_management_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyManagementTypes)
	return companyManagementTypes, res.Error
}

func (db *CompanyManagementTypeConnection) CountSearchCompanyManagementType(search string) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_management_types").Where("(lower(company_management_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
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
