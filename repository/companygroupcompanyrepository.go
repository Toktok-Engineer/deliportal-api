package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyGroupCompanyRepository interface {
	CountCompanyGroupCompanyAll(companyGroupID int) (count int64, err error)
	FindCompanyGroupCompanys() (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error)
	FindCompanyGroupCompanysOffset(limit int, offset int, order string, dir string, companyGroupID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error)
	SearchCompanyGroupCompany(limit int, offset int, order string, dir string, search string, companyGroupID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error)
	CountSearchCompanyGroupCompany(search string, companyGroupID int) (count int64, err error)
	FindCompanyGroupCompanyById(id uint) (companyGroupCompanyOutput model.SelectCompanyGroupCompanyParameter, err error)
	InsertCompanyGroupCompany(companyGroupCompany model.CompanyGroupCompany) (companyGroupCompanyOutput model.CompanyGroupCompany, err error)
	UpdateCompanyGroupCompany(companyGroupCompany model.CompanyGroupCompany, id uint) (companyGroupCompanyOutput model.CompanyGroupCompany, err error)
}

type CompanyGroupCompanyConnection struct {
	connection *gorm.DB
}

func NewCompanyGroupCompanyRepository(db *gorm.DB) CompanyGroupCompanyRepository {
	return &CompanyGroupCompanyConnection{
		connection: db,
	}
}

func (db *CompanyGroupCompanyConnection) CountCompanyGroupCompanyAll(companyGroupID int) (count int64, err error) {
	res := db.connection.Debug().Table("company_group_companies").Where("company_group_id = ? AND deleted_at = 0", companyGroupID).Count(&count)
	return count, res.Error
}

func (db *CompanyGroupCompanyConnection) FindCompanyGroupCompanys() (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error) {
	var (
		company_group_companies []model.SelectCompanyGroupCompanyParameter
	)
	res := db.connection.Debug().Table("company_group_companies").Select("company_group_companies.id, company_group_companies.company_group_id, company_groups.company_group_name, company_group_companies.company_id, companies.company_name, company_group_companies.remark, company_group_companies.created_user_id, company_group_companies.updated_user_id, company_group_companies.deleted_user_id, company_group_companies.created_at, company_group_companies.updated_at, company_group_companies.deleted_at").Joins("left join company_groups ON company_group_companies.company_group_id = company_groups.id").Joins("left join companies ON company_group_companies.company_id = companies.id").Where("company_group_companies.deleted_at = 0").Order("companies.company_name").Find(&company_group_companies)
	return company_group_companies, res.Error
}

func (db *CompanyGroupCompanyConnection) FindCompanyGroupCompanysOffset(limit int, offset int, order string, dir string, companyGroupID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error) {
	var (
		orderDirection          string
		company_group_companies []model.SelectCompanyGroupCompanyParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("company_group_companies").Select("company_group_companies.id, company_group_companies.company_group_id, company_groups.company_group_name, company_group_companies.company_id, companies.company_name, company_group_companies.remark, company_group_companies.created_user_id, company_group_companies.updated_user_id, company_group_companies.deleted_user_id, company_group_companies.created_at, company_group_companies.updated_at, company_group_companies.deleted_at").Joins("left join company_groups ON company_group_companies.company_group_id = company_groups.id").Joins("left join companies ON company_group_companies.company_id = companies.id").Where("company_group_companies.company_group_id = ? AND company_group_companies.deleted_at = 0", companyGroupID).Order(orderDirection).Limit(limit).Offset(offset).Find(&company_group_companies)
	return company_group_companies, res.Error
}

func (db *CompanyGroupCompanyConnection) SearchCompanyGroupCompany(limit int, offset int, order string, dir string, search string, companyGroupID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error) {
	var (
		orderDirection          string
		final                   string
		company_group_companies []model.SelectCompanyGroupCompanyParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_group_companies").Select("company_group_companies.id, company_group_companies.company_group_id, company_groups.company_group_name, company_group_companies.company_id, companies.company_name, company_group_companies.remark, company_group_companies.created_user_id, company_group_companies.updated_user_id, company_group_companies.deleted_user_id, company_group_companies.created_at, company_group_companies.updated_at, company_group_companies.deleted_at").Joins("left join company_groups ON company_group_companies.company_group_id = company_groups.id").Joins("left join companies ON company_group_companies.company_id = companies.id").Where("(lower(company_groups.company_group_name) LIKE ? OR lower(companies.company_name) LIKE ? OR lower(company_group_companies.remark) LIKE ?) AND company_group_companies.company_group_id = ? AND company_group_companies.deleted_at = 0", final, final, final, companyGroupID).Order(orderDirection).Limit(limit).Offset(offset).Find(&company_group_companies)
	return company_group_companies, res.Error
}

func (db *CompanyGroupCompanyConnection) CountSearchCompanyGroupCompany(search string, companyGroupID int) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_group_companies").Select("company_group_companies.id, company_group_companies.company_group_id, company_groups.company_group_name, company_group_companies.company_id, companies.company_name, company_group_companies.remark, company_group_companies.created_user_id, company_group_companies.updated_user_id, company_group_companies.deleted_user_id, company_group_companies.created_at, company_group_companies.updated_at, company_group_companies.deleted_at").Joins("left join company_groups ON company_group_companies.company_group_id = company_groups.id").Joins("left join companies ON company_group_companies.company_id = companies.id").Where("(lower(company_groups.company_group_name) LIKE ? OR lower(companies.company_name) LIKE ? OR lower(company_group_companies.remark) LIKE ?) AND company_group_companies.company_group_id = ? AND company_group_companies.deleted_at = 0", final, final, final, companyGroupID).Count(&count)
	return count, res.Error
}

func (db *CompanyGroupCompanyConnection) FindCompanyGroupCompanyById(id uint) (companyGroupCompanyOutput model.SelectCompanyGroupCompanyParameter, err error) {
	var (
		companyGroupCompany model.SelectCompanyGroupCompanyParameter
	)

	res := db.connection.Debug().Table("company_group_companies").Select("company_group_companies.id, company_group_companies.company_group_id, company_groups.company_group_name, company_group_companies.company_id, companies.company_name, company_group_companies.remark, company_group_companies.created_user_id, company_group_companies.updated_user_id, company_group_companies.deleted_user_id, company_group_companies.created_at, company_group_companies.updated_at, company_group_companies.deleted_at").Joins("left join company_groups ON company_group_companies.company_group_id = company_groups.id").Joins("left join companies ON company_group_companies.company_id = companies.id").Where("company_group_companies.id=? AND company_group_companies.deleted_at = 0", id).Take(&companyGroupCompany)
	return companyGroupCompany, res.Error
}

func (db *CompanyGroupCompanyConnection) InsertCompanyGroupCompany(companyGroupCompany model.CompanyGroupCompany) (companyGroupCompanyOutput model.CompanyGroupCompany, err error) {
	res := db.connection.Save(&companyGroupCompany)
	return companyGroupCompany, res.Error
}

func (db *CompanyGroupCompanyConnection) UpdateCompanyGroupCompany(companyGroupCompany model.CompanyGroupCompany, id uint) (companyGroupCompanyOutput model.CompanyGroupCompany, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyGroupCompany)
	return companyGroupCompany, res.Error
}
