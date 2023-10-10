package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyGroupRepository interface {
	CountCompanyGroupAll() (count int64, err error)
	FindCompanyGroups() (businessunitOutput []model.CompanyGroup, err error)
	FindCompanyGroupsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.CompanyGroup, err error)
	SearchCompanyGroup(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.CompanyGroup, err error)
	CountSearchCompanyGroup(search string) (count int64, err error)
	FindCompanyGroupById(id uint) (companyGroupOutput model.CompanyGroup, err error)
	FindExcCompanyGroup(id uint) (companyGroupOutput []model.CompanyGroup, err error)
	InsertCompanyGroup(companyGroup model.CompanyGroup) (companyGroupOutput model.CompanyGroup, err error)
	UpdateCompanyGroup(companyGroup model.CompanyGroup, id uint) (companyGroupOutput model.CompanyGroup, err error)
}

type companyGroupConnection struct {
	connection *gorm.DB
}

func NewCompanyGroupRepository(db *gorm.DB) CompanyGroupRepository {
	return &companyGroupConnection{
		connection: db,
	}
}
func (db *companyGroupConnection) CountCompanyGroupAll() (count int64, err error) {
	res := db.connection.Debug().Table("company_groups").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *companyGroupConnection) FindCompanyGroups() (businessunitOutput []model.CompanyGroup, err error) {
	var (
		businessunits []model.CompanyGroup
	)
	res := db.connection.Where("deleted_at = 0").Order("company_group_name").Find(&businessunits)
	return businessunits, res.Error
}

func (db *companyGroupConnection) FindCompanyGroupsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.CompanyGroup, err error) {
	var (
		orderDirection string
		businessunits  []model.CompanyGroup
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&businessunits)
	return businessunits, res.Error
}

func (db *companyGroupConnection) SearchCompanyGroup(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.CompanyGroup, err error) {
	var (
		orderDirection string
		final          string
		businessunits  []model.CompanyGroup
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(company_group_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&businessunits)
	return businessunits, res.Error
}

func (db *companyGroupConnection) CountSearchCompanyGroup(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_groups").Where("(lower(company_group_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
}

func (db *companyGroupConnection) FindCompanyGroupById(id uint) (companyGroupOutput model.CompanyGroup, err error) {
	var (
		companyGroup model.CompanyGroup
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&companyGroup)
	return companyGroup, res.Error
}

func (db *companyGroupConnection) FindExcCompanyGroup(id uint) (companyGroupOutput []model.CompanyGroup, err error) {
	var (
		companyGroups []model.CompanyGroup
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("company_group_name").Find(&companyGroups)
	return companyGroups, res.Error
}

func (db *companyGroupConnection) InsertCompanyGroup(companyGroup model.CompanyGroup) (companyGroupOutput model.CompanyGroup, err error) {
	res := db.connection.Save(&companyGroup)
	return companyGroup, res.Error
}

func (db *companyGroupConnection) UpdateCompanyGroup(companyGroup model.CompanyGroup, id uint) (companyGroupOutput model.CompanyGroup, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyGroup)
	return companyGroup, res.Error
}
