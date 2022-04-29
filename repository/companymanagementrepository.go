package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type CompanyManagementRepository interface {
	FindCompanyManagements() (companyManagementOutput []model.SelectCompanyManagementParameter, err error)
	FindCompanyManagementById(id uint) (companyManagementOutput model.SelectCompanyManagementParameter, err error)
	FindExcCompanyManagement(id uint) (companyManagementOutput []model.SelectCompanyManagementParameter, err error)
	FindCompanyManagementByCompanyId(id uint) (companyManagementOutput []model.SelectCompanyManagementParameter, err error)
	InsertCompanyManagement(companyManagement model.CompanyManagement) (companyManagementOutput model.CompanyManagement, err error)
	UpdateCompanyManagement(companyManagement model.CompanyManagement, id uint) (companyManagementOutput model.CompanyManagement, err error)
}

type CompanyManagementConnection struct {
	connection *gorm.DB
}

func NewCompanyManagementRepository(db *gorm.DB) CompanyManagementRepository {
	return &CompanyManagementConnection{
		connection: db,
	}
}

func (db *CompanyManagementConnection) FindCompanyManagements() (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
	)

	res := db.connection.Debug().Table("company_managements").Select("company_managements.id, company_managements.company_id, companies.company_name, company_managements.company_management_type_id, company_management_types.company_management_type_name, company_managements.management_name, company_managements.remark, company_managements.created_user_id, company_managements.updated_user_id, company_managements.deleted_user_id, company_managements.created_at, company_managements.updated_at, company_managements.deleted_at").Joins("left join companies ON company_managements.company_id = companies.id").Joins("left join company_management_types ON company_managements.company_management_type_id = company_management_types.id").Where("company_managements.deleted_at = 0").Order("company_managements.id").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementConnection) FindCompanyManagementById(id uint) (companyManagementOutput model.SelectCompanyManagementParameter, err error) {
	var (
		companyManagement model.SelectCompanyManagementParameter
	)

	res := db.connection.Debug().Table("company_managements").Select("company_managements.id, company_managements.company_id, companies.company_name, company_managements.company_management_type_id, company_management_types.company_management_type_name, company_managements.management_name, company_managements.remark, company_managements.created_user_id, company_managements.updated_user_id, company_managements.deleted_user_id, company_managements.created_at, company_managements.updated_at, company_managements.deleted_at").Joins("left join companies ON company_managements.company_id = companies.id").Joins("left join company_management_types ON company_managements.company_management_type_id = company_management_types.id").Where("company_managements.id=? AND company_managements.deleted_at = 0", id).Take(&companyManagement)
	return companyManagement, res.Error
}

func (db *CompanyManagementConnection) FindExcCompanyManagement(id uint) (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
	)

	res := db.connection.Debug().Table("company_managements").Select("company_managements.id, company_managements.company_id, companies.company_name, company_managements.company_management_type_id, company_management_types.company_management_type_name, company_managements.management_name, company_managements.remark, company_managements.created_user_id, company_managements.updated_user_id, company_managements.deleted_user_id, company_managements.created_at, company_managements.updated_at, company_managements.deleted_at").Joins("left join companies ON company_managements.company_id = companies.id").Joins("left join company_management_types ON company_managements.company_management_type_id = company_management_types.id").Where(" company_managements.id!=? AND company_managements.deleted_at = 0", id).Order("company_managements.id").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementConnection) FindCompanyManagementByCompanyId(id uint) (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
	)

	res := db.connection.Debug().Table("company_managements").Select("company_managements.id, company_managements.company_id, companies.company_name, company_managements.company_management_type_id, company_management_types.company_management_type_name, company_managements.management_name, company_managements.remark, company_managements.created_user_id, company_managements.updated_user_id, company_managements.deleted_user_id, company_managements.created_at, company_managements.updated_at, company_managements.deleted_at").Joins("left join companies ON company_managements.company_id = companies.id").Joins("left join company_management_types ON company_managements.company_management_type_id = company_management_types.id").Where("company_managements.company_id=? AND company_managements.deleted_at = 0", id).Order("company_managements.id").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementConnection) InsertCompanyManagement(companyManagement model.CompanyManagement) (companyManagementOutput model.CompanyManagement, err error) {
	res := db.connection.Save(&companyManagement)
	return companyManagement, res.Error
}

func (db *CompanyManagementConnection) UpdateCompanyManagement(companyManagement model.CompanyManagement, id uint) (companyManagementOutput model.CompanyManagement, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyManagement)
	return companyManagement, res.Error
}
