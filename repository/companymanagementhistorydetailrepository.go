package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyManagementHistoryDetailRepository interface {
	CountCompanyManagementHistoryDetailAll(companyManagementHistoryId int) (count int64, err error)
	FindCompanyManagementHistoryDetails(companyManagementHistoryId int) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error)
	FindCompanyManagementHistoryDetailsOffset(limit int, offset int, order string, dir string, companyManagementHistoryId int) (companyManagementHistoryDetailOutput []model.SelectCompanyManagementHistoryDetail, err error)
	SearchCompanyManagementHistoryDetail(limit int, offset int, order string, dir string, search string, companyManagementHistoryId int) (companyManagementHistoryDetailOutput []model.SelectCompanyManagementHistoryDetail, err error)
	CountSearchCompanyManagementHistoryDetail(search string, companyManagementHistoryId int) (count int64, err error)
	FindCompanyManagementHistoryDetailByCompanyId(id uint) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error)
	FindCompanyManagementHistoryDetailById(id uint) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error)
	FindExcCompanyManagementHistoryDetail(id uint) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error)
	InsertCompanyManagementHistoryDetail(companyManagementHistoryDetail model.CompanyManagementHistoryDetail) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error)
	UpdateCompanyManagementHistoryDetail(companyManagementHistoryDetail model.CompanyManagementHistoryDetail, id uint) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error)
}

type CompanyManagementHistoryDetailConnection struct {
	connection *gorm.DB
}

func NewCompanyManagementHistoryDetailRepository(db *gorm.DB) CompanyManagementHistoryDetailRepository {
	return &CompanyManagementHistoryDetailConnection{
		connection: db,
	}
}

func (db *CompanyManagementHistoryDetailConnection) CountCompanyManagementHistoryDetailAll(companyManagementHistoryId int) (count int64, err error) {
	res := db.connection.Table("company_management_history_details").Where("company_management_history_id = ? AND deleted_at = 0", companyManagementHistoryId).Count(&count)
	return count, res.Error
}

func (db *CompanyManagementHistoryDetailConnection) FindCompanyManagementHistoryDetails(companyManagementHistoryId int) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error) {
	var (
		companyManagementHistoryDetails []model.CompanyManagementHistoryDetail
	)
	res := db.connection.Where("company_management_history_id = ? AND deleted_at = 0", companyManagementHistoryId).Order("management_name").Find(&companyManagementHistoryDetails)
	return companyManagementHistoryDetails, res.Error
}

func (db *CompanyManagementHistoryDetailConnection) FindCompanyManagementHistoryDetailsOffset(limit int, offset int, order string, dir string, companyManagementHistoryId int) (companyManagementHistoryDetailOutput []model.SelectCompanyManagementHistoryDetail, err error) {
	var (
		orderDirection                  string
		companyManagementHistoryDetails []model.SelectCompanyManagementHistoryDetail
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("company_management_history_details").Select("company_management_history_details.id, company_management_history_details.company_management_history_id, company_management_history_details.company_management_type_id, company_management_types.company_management_type_name, company_management_history_details.management_name, company_management_history_details.remark, company_management_history_details.created_user_id, createdUID.username AS created_user, company_management_history_details.updated_user_id, updatedUID.username AS updated_user, company_management_history_details.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_management_history_details.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_management_history_details.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_management_history_details.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join company_management_types ON company_management_history_details.company_management_type_id = company_management_types.id").Joins("left join users createdUID on company_management_history_details.created_user_id = createdUID.id").Joins("left join users updatedUID on company_management_history_details.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_management_history_details.deleted_user_id = deletedUID.id").Where("company_management_history_details.company_management_history_id = ? AND company_management_history_details.deleted_at = 0", companyManagementHistoryId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyManagementHistoryDetails)
	return companyManagementHistoryDetails, res.Error
}

func (db *CompanyManagementHistoryDetailConnection) SearchCompanyManagementHistoryDetail(limit int, offset int, order string, dir string, search string, companyManagementHistoryId int) (companyManagementHistoryDetailOutput []model.SelectCompanyManagementHistoryDetail, err error) {
	var (
		orderDirection                  string
		final                           string
		companyManagementHistoryDetails []model.SelectCompanyManagementHistoryDetail
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("company_management_history_details").Select("company_management_history_details.id, company_management_history_details.company_management_history_id, company_management_history_details.company_management_type_id, company_management_types.company_management_type_name, company_management_history_details.management_name, company_management_history_details.remark, company_management_history_details.created_user_id, createdUID.username AS created_user, company_management_history_details.updated_user_id, updatedUID.username AS updated_user, company_management_history_details.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_management_history_details.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_management_history_details.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_management_history_details.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join company_management_types ON company_management_history_details.company_management_type_id = company_management_types.id").Joins("left join users createdUID on company_management_history_details.created_user_id = createdUID.id").Joins("left join users updatedUID on company_management_history_details.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_management_history_details.deleted_user_id = deletedUID.id").Where("(lower(company_management_history_details.management_name) LIKE ? OR lower(company_management_types.company_management_type_name) LIKE ? OR lower(to_char(to_timestamp(company_management_history_details.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_management_history_details.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(company_management_history_details.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? ) AND company_management_history_details.company_management_history_id = ? AND company_management_history_details.deleted_at = 0", final, final, final, final, final, final, final, companyManagementHistoryId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyManagementHistoryDetails)
	return companyManagementHistoryDetails, res.Error
}

func (db *CompanyManagementHistoryDetailConnection) CountSearchCompanyManagementHistoryDetail(search string, companyManagementHistoryId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("company_management_history_details").Select("company_management_history_details").Select("company_management_history_details.id, company_management_history_details.company_management_history_id, company_management_history_details.company_management_type_id, company_management_types.company_management_type_name, company_management_history_details.management_name, company_management_history_details.remark, company_management_history_details.created_user_id, createdUID.username AS created_user, company_management_history_details.updated_user_id, updatedUID.username AS updated_user, company_management_history_details.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_management_history_details.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_management_history_details.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_management_history_details.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join company_management_types ON company_management_history_details.company_management_type_id = company_management_types.id").Joins("left join users createdUID on company_management_history_details.created_user_id = createdUID.id").Joins("left join users updatedUID on company_management_history_details.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_management_history_details.deleted_user_id = deletedUID.id").Where("(lower(company_management_history_details.management_name) LIKE ? OR lower(company_management_types.company_management_type_name) LIKE ? OR lower(to_char(to_timestamp(company_management_history_details.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_management_history_details.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(company_management_history_details.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? ) AND company_management_history_details.company_management_history_id = ? AND company_management_history_details.deleted_at = 0", final, final, final, final, final, final, final, companyManagementHistoryId).Count(&count)
	return count, res.Error
}

func (db *CompanyManagementHistoryDetailConnection) FindCompanyManagementHistoryDetailByCompanyId(id uint) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error) {
	var (
		companyManagementHistoryDetails []model.CompanyManagementHistoryDetail
	)
	res := db.connection.Where("company_management_history_id=? AND deleted_at = 0", id).Order("management_name").Find(&companyManagementHistoryDetails)
	return companyManagementHistoryDetails, res.Error
}

func (db *CompanyManagementHistoryDetailConnection) FindCompanyManagementHistoryDetailById(id uint) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error) {
	var (
		companyManagementHistoryDetail model.CompanyManagementHistoryDetail
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&companyManagementHistoryDetail)
	return companyManagementHistoryDetail, res.Error
}

func (db *CompanyManagementHistoryDetailConnection) FindExcCompanyManagementHistoryDetail(id uint) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error) {
	var (
		companyManagementHistoryDetails []model.CompanyManagementHistoryDetail
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("management_name").Find(&companyManagementHistoryDetails)
	return companyManagementHistoryDetails, res.Error
}

func (db *CompanyManagementHistoryDetailConnection) InsertCompanyManagementHistoryDetail(companyManagementHistoryDetail model.CompanyManagementHistoryDetail) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error) {
	res := db.connection.Save(&companyManagementHistoryDetail)
	return companyManagementHistoryDetail, res.Error
}

func (db *CompanyManagementHistoryDetailConnection) UpdateCompanyManagementHistoryDetail(companyManagementHistoryDetail model.CompanyManagementHistoryDetail, id uint) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyManagementHistoryDetail)
	return companyManagementHistoryDetail, res.Error
}
