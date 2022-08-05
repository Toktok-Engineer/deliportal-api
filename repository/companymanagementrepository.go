package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyManagementRepository interface {
	CountCompanyManagementAll(companyId int) (count int64, err error)
	FindCompanyManagements(companyId int) (companyManagementOutput []model.CompanyManagement, err error)
	FindCompanyManagementsOffset(limit int, offset int, order string, dir string, companyId int) (companyManagementOutput []model.SelectCompanyManagementParameter, err error)
	SearchCompanyManagement(limit int, offset int, order string, dir string, search string, companyId int) (companyManagementOutput []model.SelectCompanyManagementParameter, err error)
	CountSearchCompanyManagement(search string, companyId int) (count int64, err error)
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

func (db *CompanyManagementConnection) CountCompanyManagementAll(companyId int) (count int64, err error) {
	res := db.connection.Debug().Table("company_managements").Where("company_id = ? AND deleted_at = 0", companyId).Count(&count)
	return count, res.Error
}

func (db *CompanyManagementConnection) FindCompanyManagements(companyId int) (companyManagementOutput []model.CompanyManagement, err error) {
	var (
		companyManagements []model.CompanyManagement
	)
	res := db.connection.Where("company_id = ? AND deleted_at = 0", companyId).Order("shareholder_name").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementConnection) FindCompanyManagementsOffset(limit int, offset int, order string, dir string, companyId int) (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	var (
		orderDirection     string
		companyManagements []model.SelectCompanyManagementParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("company_managements").Select("company_managements.id, company_managements.company_id, companies.company_name, company_managements.company_management_type_id, company_management_types.company_management_type_name, company_managements.management_name, company_managements.remark, company_managements.created_user_id, createdUID.username AS created_user, company_managements.updated_user_id, updatedUID.username AS updated_user, company_managements.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_managements.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_managements.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_managements.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_managements.company_id = companies.id").Joins("left join company_management_types ON company_managements.company_management_type_id = company_management_types.id").Joins("left join users createdUID on company_managements.created_user_id = createdUID.id").Joins("left join users updatedUID on company_managements.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_managements.deleted_user_id = deletedUID.id").Where("company_managements.company_id = ? AND company_managements.deleted_at = 0", companyId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementConnection) SearchCompanyManagement(limit int, offset int, order string, dir string, search string, companyId int) (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	var (
		orderDirection     string
		final              string
		companyManagements []model.SelectCompanyManagementParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_managements").Select("company_managements.id, company_managements.company_id, companies.company_name, company_managements.company_management_type_id, company_management_types.company_management_type_name, company_managements.management_name, company_managements.remark, company_managements.created_user_id, createdUID.username AS created_user, company_managements.updated_user_id, updatedUID.username AS updated_user, company_managements.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_managements.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_managements.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_managements.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join company_management_types ON company_managements.company_management_type_id = company_management_types.id").Joins("left join users createdUID on company_managements.created_user_id = createdUID.id").Joins("left join users updatedUID on company_managements.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_managements.deleted_user_id = deletedUID.id").Joins("left join companies on company_managements.company_id = companies.id").Where("(lower(company_management_types.company_management_type_name) LIKE ? OR lower(company_managements.management_name) LIKE ? OR lower(company_managements.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(company_managements.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_managements.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND company_managements.company_id = ? AND company_managements.deleted_at = 0", final, final, final, final, final, final, final, companyId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementConnection) CountSearchCompanyManagement(search string, companyId int) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_managements").Select("company_managements.id, company_managements.company_id, companies.company_name, company_managements.company_management_type_id, company_management_types.company_management_type_name, company_managements.management_name, company_managements.remark, company_managements.created_user_id, createdUID.username AS created_user, company_managements.updated_user_id, updatedUID.username AS updated_user, company_managements.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_managements.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_managements.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_managements.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_managements.company_id = companies.id").Joins("left join company_management_types ON company_managements.company_management_type_id = company_management_types.id").Joins("left join users createdUID on company_managements.created_user_id = createdUID.id").Joins("left join users updatedUID on company_managements.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_managements.deleted_user_id = deletedUID.id").Where("(lower(company_management_types.company_management_type_name) LIKE ? OR lower(company_managements.management_name) LIKE ? OR lower(company_managements.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(company_managements.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_managements.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND company_managements.company_id = ? AND company_managements.deleted_at = 0", final, final, final, final, final, final, final, companyId).Count(&count)
	return count, res.Error
}

func (db *CompanyManagementConnection) FindCompanyManagementById(id uint) (companyManagementOutput model.SelectCompanyManagementParameter, err error) {
	var (
		companyManagement model.SelectCompanyManagementParameter
	)

	res := db.connection.Debug().Table("company_managements").Select("company_managements.id, company_managements.company_id, companies.company_name, company_managements.company_management_type_id, company_management_types.company_management_type_name, company_managements.management_name, company_managements.remark, company_managements.created_user_id, createdUID.username AS created_user, company_managements.updated_user_id, updatedUID.username AS updated_user, company_managements.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_managements.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_managements.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_managements.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_managements.company_id = companies.id").Joins("left join company_management_types ON company_managements.company_management_type_id = company_management_types.id").Joins("left join users createdUID on company_managements.created_user_id = createdUID.id").Joins("left join users updatedUID on company_managements.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_managements.deleted_user_id = deletedUID.id").Where("company_managements.id=? AND company_managements.deleted_at = 0", id).Take(&companyManagement)
	return companyManagement, res.Error
}

func (db *CompanyManagementConnection) FindExcCompanyManagement(id uint) (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
	)

	res := db.connection.Debug().Table("company_managements").Select("company_managements.id, company_managements.company_id, companies.company_name, company_managements.company_management_type_id, company_management_types.company_management_type_name, company_managements.management_name, company_managements.remark, company_managements.created_user_id, createdUID.username AS created_user, company_managements.updated_user_id, updatedUID.username AS updated_user, company_managements.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_managements.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_managements.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_managements.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_managements.company_id = companies.id").Joins("left join company_management_types ON company_managements.company_management_type_id = company_management_types.id").Joins("left join users createdUID on company_managements.created_user_id = createdUID.id").Joins("left join users updatedUID on company_managements.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_managements.deleted_user_id = deletedUID.id").Where(" company_managements.id!=? AND company_managements.deleted_at = 0", id).Order("company_managements.id").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementConnection) FindCompanyManagementByCompanyId(id uint) (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	var (
		companyManagements []model.SelectCompanyManagementParameter
	)

	res := db.connection.Debug().Table("company_managements").Select("company_managements.id, company_managements.company_id, companies.company_name, company_managements.company_management_type_id, company_management_types.company_management_type_name, company_managements.management_name, company_managements.remark, company_managements.created_user_id, createdUID.username AS created_user, company_managements.updated_user_id, updatedUID.username AS updated_user, company_managements.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_managements.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_managements.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_managements.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_managements.company_id = companies.id").Joins("left join company_management_types ON company_managements.company_management_type_id = company_management_types.id").Joins("left join users createdUID on company_managements.created_user_id = createdUID.id").Joins("left join users updatedUID on company_managements.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_managements.deleted_user_id = deletedUID.id").Where("company_managements.company_id=? AND company_managements.deleted_at = 0", id).Order("company_managements.id").Find(&companyManagements)
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
