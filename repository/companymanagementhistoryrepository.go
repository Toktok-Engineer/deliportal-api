package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyManagementHistoryRepository interface {
	CountCompanyManagementHistoryAll(companyId int) (count int64, err error)
	FindCompanyManagementHistorys(companyId int) (companyManagementOutput []model.CompanyManagementHistory, err error)
	FindCompanyManagementHistorysOffset(limit int, offset int, order string, dir string, companyId int) (companyManagementOutput []model.SelectCompanyManagementHistoryParameter, err error)
	SearchCompanyManagementHistory(limit int, offset int, order string, dir string, search string, companyId int) (companyManagementOutput []model.SelectCompanyManagementHistoryParameter, err error)
	CountSearchCompanyManagementHistory(search string, companyId int) (count int64, err error)
	FindCompanyManagementHistoryById(id uint) (companyManagementOutput model.SelectCompanyManagementHistoryParameter, err error)
	FindExcCompanyManagementHistory(id uint) (companyManagementOutput []model.SelectCompanyManagementHistoryParameter, err error)
	FindCompanyManagementHistoryByCompanyId(id uint) (companyManagementOutput []model.SelectCompanyManagementHistoryParameter, err error)
	InsertCompanyManagementHistory(companyManagement model.CompanyManagementHistory) (companyManagementOutput model.CompanyManagementHistory, err error)
	UpdateCompanyManagementHistory(companyManagement model.CompanyManagementHistory, id uint) (companyManagementOutput model.CompanyManagementHistory, err error)
}

type CompanyManagementHistoryConnection struct {
	connection *gorm.DB
}

func NewCompanyManagementHistoryRepository(db *gorm.DB) CompanyManagementHistoryRepository {
	return &CompanyManagementHistoryConnection{
		connection: db,
	}
}

func (db *CompanyManagementHistoryConnection) CountCompanyManagementHistoryAll(companyId int) (count int64, err error) {
	res := db.connection.Debug().Table("company_management_histories").Where("company_id = ? AND deleted_at = 0", companyId).Count(&count)
	return count, res.Error
}

func (db *CompanyManagementHistoryConnection) FindCompanyManagementHistorys(companyId int) (companyManagementOutput []model.CompanyManagementHistory, err error) {
	var (
		companyManagements []model.CompanyManagementHistory
	)
	res := db.connection.Where("company_id = ? AND deleted_at = 0", companyId).Order("company_akte_id").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementHistoryConnection) FindCompanyManagementHistorysOffset(limit int, offset int, order string, dir string, companyId int) (companyManagementOutput []model.SelectCompanyManagementHistoryParameter, err error) {
	var (
		orderDirection     string
		companyManagements []model.SelectCompanyManagementHistoryParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("company_management_histories").Select("company_management_histories.id, company_management_histories.company_id, companies.company_name, company_management_histories.company_akte_id, company_aktes.akte_no, string_agg(CASE WHEN company_management_history_details.management_name LIKE '%-DELETE' THEN NULL ELSE CONCAT(CONCAT(company_management_types.company_management_type_name, ' : '), company_management_history_details.management_name) END, '<br/> ') AS change_information, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_management_histories.remark, company_management_histories.created_user_id, createdUID.username AS created_user, company_management_histories.updated_user_id, updatedUID.username AS updated_user, company_management_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_management_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_management_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_management_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join company_management_history_details ON company_management_history_details.company_management_history_id = company_management_histories.id").Joins("left join company_management_types ON company_management_history_details.company_management_type_id = company_management_types.id").Joins("left join companies on company_management_histories.company_id = companies.id").Joins("left join company_aktes on company_management_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_management_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_management_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_management_histories.deleted_user_id = deletedUID.id").Where("company_management_histories.company_id = ? AND company_management_histories.deleted_at = 0", companyId).Group("company_management_histories.id, company_management_histories.company_id, companies.company_name, company_management_histories.company_akte_id, company_aktes.akte_no, company_aktes.akte_date, company_aktes.year, company_management_histories.remark, company_management_histories.created_user_id, createdUID.username, company_management_histories.updated_user_id, updatedUID.username, company_management_histories.deleted_user_id, deletedUID.username, company_management_histories.created_at, company_management_histories.updated_at, company_management_histories.deleted_at").Order(orderDirection).Limit(limit).Offset(offset).Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementHistoryConnection) SearchCompanyManagementHistory(limit int, offset int, order string, dir string, search string, companyId int) (companyManagementOutput []model.SelectCompanyManagementHistoryParameter, err error) {
	var (
		orderDirection     string
		final              string
		companyManagements []model.SelectCompanyManagementHistoryParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_management_histories").Select("company_management_histories.id, company_management_histories.company_id, companies.company_name, company_management_histories.company_akte_id, company_aktes.akte_no, string_agg(CASE WHEN company_management_history_details.management_name LIKE '%-DELETE' THEN NULL ELSE CONCAT(CONCAT(company_management_types.company_management_type_name, ' : '), company_management_history_details.management_name) END, '<br/> ') AS change_information, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_management_histories.remark, company_management_histories.created_user_id, createdUID.username AS created_user, company_management_histories.updated_user_id, updatedUID.username AS updated_user, company_management_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_management_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_management_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_management_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join company_management_history_details on company_management_histories.id = company_management_history_details.company_management_history_id").Joins("left join company_management_types ON company_management_history_details.company_management_type_id = company_management_types.id").Joins("left join companies on company_management_histories.company_id = companies.id").Joins("left join company_aktes on company_management_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_management_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_management_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_management_histories.deleted_user_id = deletedUID.id").Having("(company_aktes.year::text LIKE ? OR lower(to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(string_agg(CASE WHEN company_management_history_details.management_name LIKE '%-DELETE' THEN NULL ELSE CONCAT(CONCAT(company_management_types.company_management_type_name, ' : '), company_management_history_details.management_name) END, '<br/> ')) LIKE ? OR lower(company_aktes.akte_no) LIKE ? OR lower(company_management_histories.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(company_management_histories.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_management_histories.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND company_management_histories.company_id = ? AND company_management_histories.deleted_at = 0", final, final, final, final, final, final, final, final, final, companyId).Group("company_management_histories.id, company_management_histories.company_id, companies.company_name, company_management_histories.company_akte_id, company_aktes.akte_no, company_aktes.akte_date, company_aktes.year, company_management_histories.remark, company_management_histories.created_user_id, createdUID.username, company_management_histories.updated_user_id, updatedUID.username, company_management_histories.deleted_user_id, deletedUID.username, company_management_histories.created_at, company_management_histories.updated_at, company_management_histories.deleted_at").Order(orderDirection).Limit(limit).Offset(offset).Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementHistoryConnection) CountSearchCompanyManagementHistory(search string, companyId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_management_histories").Select("company_management_histories.id, company_management_histories.company_id, companies.company_name, company_management_histories.company_akte_id, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_management_histories.remark, company_management_histories.created_user_id, createdUID.username AS created_user, company_management_histories.updated_user_id, updatedUID.username AS updated_user, company_management_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_management_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_management_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_management_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join company_management_history_details on company_management_histories.id = company_management_history_details.company_management_history_id").Joins("left join company_management_types ON company_management_history_details.company_management_type_id = company_management_types.id").Joins("left join companies on company_management_histories.company_id = companies.id").Joins("left join company_aktes on company_management_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_management_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_management_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_management_histories.deleted_user_id = deletedUID.id").Having("(company_aktes.year::text LIKE ? OR lower(to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(company_aktes.akte_no) LIKE ? OR lower(string_agg(CASE WHEN company_management_history_details.management_name LIKE '%-DELETE' THEN NULL ELSE CONCAT(CONCAT(company_management_types.company_management_type_name, ' : '), company_management_history_details.management_name) END, '<br/> ')) LIKE ? OR lower(company_management_histories.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(company_management_histories.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_management_histories.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND company_management_histories.company_id = ? AND company_management_histories.deleted_at = 0", final, final, final, final, final, final, final, final, final, companyId).Group("company_management_histories.id, company_management_histories.company_id, companies.company_name, company_management_histories.company_akte_id, company_aktes.akte_no, company_aktes.akte_date, company_aktes.year, company_management_histories.remark, company_management_histories.created_user_id, createdUID.username, company_management_histories.updated_user_id, updatedUID.username, company_management_histories.deleted_user_id, deletedUID.username, company_management_histories.created_at, company_management_histories.updated_at, company_management_histories.deleted_at").Count(&count)
	return count, res.Error
}

func (db *CompanyManagementHistoryConnection) FindCompanyManagementHistoryById(id uint) (companyManagementOutput model.SelectCompanyManagementHistoryParameter, err error) {
	var (
		companyManagement model.SelectCompanyManagementHistoryParameter
	)

	res := db.connection.Debug().Table("company_management_histories").Select("company_management_histories.id, company_management_histories.company_id, companies.company_name, company_management_histories.company_akte_id, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_management_histories.remark, company_management_histories.created_user_id, createdUID.username AS created_user, company_management_histories.updated_user_id, updatedUID.username AS updated_user, company_management_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_management_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_management_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_management_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_management_histories.company_id = companies.id").Joins("left join company_aktes on company_management_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_management_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_management_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_management_histories.deleted_user_id = deletedUID.id").Where("company_management_histories.id=? AND company_management_histories.deleted_at = 0", id).Take(&companyManagement)
	return companyManagement, res.Error
}

func (db *CompanyManagementHistoryConnection) FindExcCompanyManagementHistory(id uint) (companyManagementOutput []model.SelectCompanyManagementHistoryParameter, err error) {
	var (
		companyManagements []model.SelectCompanyManagementHistoryParameter
	)

	res := db.connection.Debug().Table("company_management_histories").Select("company_management_histories.id, company_management_histories.company_id, companies.company_name, company_management_histories.company_akte_id, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_management_histories.remark, company_management_histories.created_user_id, createdUID.username AS created_user, company_management_histories.updated_user_id, updatedUID.username AS updated_user, company_management_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_management_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_management_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_management_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_management_histories.company_id = companies.id").Joins("left join company_aktes on company_management_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_management_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_management_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_management_histories.deleted_user_id = deletedUID.id").Where(" company_management_histories.id!=? AND company_management_histories.deleted_at = 0", id).Order("company_management_histories.id").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementHistoryConnection) FindCompanyManagementHistoryByCompanyId(id uint) (companyManagementOutput []model.SelectCompanyManagementHistoryParameter, err error) {
	var (
		companyManagements []model.SelectCompanyManagementHistoryParameter
	)

	res := db.connection.Debug().Table("company_management_histories").Select("company_management_histories.id, company_management_histories.company_id, companies.company_name, company_management_histories.company_akte_id, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_management_histories.remark, company_management_histories.created_user_id, createdUID.username AS created_user, company_management_histories.updated_user_id, updatedUID.username AS updated_user, company_management_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_management_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_management_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_management_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_management_histories.company_id = companies.id").Joins("left join company_aktes on company_management_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_management_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_management_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_management_histories.deleted_user_id = deletedUID.id").Where("company_management_histories.company_id=? AND company_management_histories.deleted_at = 0", id).Order("company_management_histories.id").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyManagementHistoryConnection) InsertCompanyManagementHistory(companyManagement model.CompanyManagementHistory) (companyManagementOutput model.CompanyManagementHistory, err error) {
	res := db.connection.Save(&companyManagement)
	return companyManagement, res.Error
}

func (db *CompanyManagementHistoryConnection) UpdateCompanyManagementHistory(companyManagement model.CompanyManagementHistory, id uint) (companyManagementOutput model.CompanyManagementHistory, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyManagement)
	return companyManagement, res.Error
}
