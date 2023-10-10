package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyShareholderHistoryRepository interface {
	CountCompanyShareholderHistoryAll(companyId int) (count int64, err error)
	FindCompanyShareholderHistorys(companyId int) (companyManagementOutput []model.CompanyShareholderHistory, err error)
	FindCompanyShareholderHistorysOffset(limit int, offset int, order string, dir string, companyId int) (companyManagementOutput []model.SelectCompanyShareholderHistoryParameter, err error)
	SearchCompanyShareholderHistory(limit int, offset int, order string, dir string, search string, companyId int) (companyManagementOutput []model.SelectCompanyShareholderHistoryParameter, err error)
	CountSearchCompanyShareholderHistory(search string, companyId int) (count int64, err error)
	FindCompanyShareholderHistoryById(id uint) (companyManagementOutput model.SelectCompanyShareholderHistoryParameter, err error)
	FindExcCompanyShareholderHistory(id uint) (companyManagementOutput []model.SelectCompanyShareholderHistoryParameter, err error)
	FindCompanyShareholderHistoryByCompanyId(id uint) (companyManagementOutput []model.SelectCompanyShareholderHistoryParameter, err error)
	InsertCompanyShareholderHistory(companyManagement model.CompanyShareholderHistory) (companyManagementOutput model.CompanyShareholderHistory, err error)
	UpdateCompanyShareholderHistory(companyManagement model.CompanyShareholderHistory, id uint) (companyManagementOutput model.CompanyShareholderHistory, err error)
}

type CompanyShareholderHistoryConnection struct {
	connection *gorm.DB
}

func NewCompanyShareholderHistoryRepository(db *gorm.DB) CompanyShareholderHistoryRepository {
	return &CompanyShareholderHistoryConnection{
		connection: db,
	}
}

func (db *CompanyShareholderHistoryConnection) CountCompanyShareholderHistoryAll(companyId int) (count int64, err error) {
	res := db.connection.Debug().Table("company_shareholder_histories").Where("company_id = ? AND deleted_at = 0", companyId).Count(&count)
	return count, res.Error
}

func (db *CompanyShareholderHistoryConnection) FindCompanyShareholderHistorys(companyId int) (companyManagementOutput []model.CompanyShareholderHistory, err error) {
	var (
		companyManagements []model.CompanyShareholderHistory
	)
	res := db.connection.Where("company_id = ? AND deleted_at = 0", companyId).Order("company_akte_id").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyShareholderHistoryConnection) FindCompanyShareholderHistorysOffset(limit int, offset int, order string, dir string, companyId int) (companyManagementOutput []model.SelectCompanyShareholderHistoryParameter, err error) {
	var (
		orderDirection     string
		companyManagements []model.SelectCompanyShareholderHistoryParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("company_shareholder_histories").Select("company_shareholder_histories.id, company_shareholder_histories.company_id, companies.company_name, company_shareholder_histories.company_akte_id, company_aktes.akte_no, string_agg(CASE WHEN company_shareholder_history_details.shareholder_name LIKE '%-DELETE' THEN NULL ELSE CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(company_shareholder_history_details.shareholder_name, ' : '), TO_CHAR(company_shareholder_history_details.number_of_share, 'fm999G999G999G999')),' saham ('), company_shareholder_history_details.percentage_of_share),'% equals to '), TO_CHAR(company_shareholder_history_details.share_amount, 'fm999G999G999G999')), ')') END, '<br/> ') AS change_information, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_shareholder_histories.remark, company_shareholder_histories.created_user_id, createdUID.username AS created_user, company_shareholder_histories.updated_user_id, updatedUID.username AS updated_user, company_shareholder_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholder_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholder_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholder_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join company_shareholder_history_details on company_shareholder_histories.id = company_shareholder_history_details.company_shareholder_history_id").Joins("left join companies on company_shareholder_histories.company_id = companies.id").Joins("left join company_aktes on company_shareholder_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_shareholder_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholder_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholder_histories.deleted_user_id = deletedUID.id").Where("company_shareholder_histories.company_id = ? AND company_shareholder_histories.deleted_at = 0", companyId).Group("company_shareholder_histories.id, company_shareholder_histories.company_id, companies.company_name, company_shareholder_histories.company_akte_id, company_aktes.akte_no, company_aktes.akte_date, company_aktes.year, company_shareholder_histories.remark, company_shareholder_histories.created_user_id, createdUID.username, company_shareholder_histories.updated_user_id, updatedUID.username, company_shareholder_histories.deleted_user_id, deletedUID.username, company_shareholder_histories.created_at, company_shareholder_histories.updated_at, company_shareholder_histories.deleted_at").Order(orderDirection).Limit(limit).Offset(offset).Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyShareholderHistoryConnection) SearchCompanyShareholderHistory(limit int, offset int, order string, dir string, search string, companyId int) (companyManagementOutput []model.SelectCompanyShareholderHistoryParameter, err error) {
	var (
		orderDirection     string
		final              string
		companyManagements []model.SelectCompanyShareholderHistoryParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_shareholder_histories").Select("company_shareholder_histories.id, company_shareholder_histories.company_id, companies.company_name, company_shareholder_histories.company_akte_id, company_aktes.akte_no, string_agg(CASE WHEN company_shareholder_history_details.shareholder_name LIKE '%-DELETE' THEN NULL ELSE CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(company_shareholder_history_details.shareholder_name, ' : '), TO_CHAR(company_shareholder_history_details.number_of_share, 'fm999G999G999G999')),' saham ('), company_shareholder_history_details.percentage_of_share),'% equals to '), TO_CHAR(company_shareholder_history_details.share_amount, 'fm999G999G999G999')), ')') END, '<br/> ') AS change_information, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_shareholder_histories.remark, company_shareholder_histories.created_user_id, createdUID.username AS created_user, company_shareholder_histories.updated_user_id, updatedUID.username AS updated_user, company_shareholder_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholder_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholder_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholder_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join company_shareholder_history_details on company_shareholder_histories.id = company_shareholder_history_details.company_shareholder_history_id").Joins("left join company_aktes on company_shareholder_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_shareholder_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholder_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholder_histories.deleted_user_id = deletedUID.id").Joins("left join companies on company_shareholder_histories.company_id = companies.id").Having("(company_aktes.year::text LIKE ? OR lower(to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(string_agg(CASE WHEN company_shareholder_history_details.shareholder_name LIKE '%-DELETE' THEN NULL ELSE CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(company_shareholder_history_details.shareholder_name, ' : '), TO_CHAR(company_shareholder_history_details.number_of_share, 'fm999G999G999G999')),' saham ('), company_shareholder_history_details.percentage_of_share),'% equals to '), TO_CHAR(company_shareholder_history_details.share_amount, 'fm999G999G999G999')), ')') END, '<br/> ')) LIKE ? OR lower(company_aktes.akte_no) LIKE ? OR lower(company_shareholder_histories.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(company_shareholder_histories.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_shareholder_histories.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND company_shareholder_histories.company_id = ? AND company_shareholder_histories.deleted_at = 0", final, final, final, final, final, final, final, final, final, companyId).Group("company_shareholder_histories.id, company_shareholder_histories.company_id, companies.company_name, company_shareholder_histories.company_akte_id, company_aktes.akte_no, company_aktes.akte_date, company_aktes.year, company_shareholder_histories.remark, company_shareholder_histories.created_user_id, createdUID.username, company_shareholder_histories.updated_user_id, updatedUID.username, company_shareholder_histories.deleted_user_id, deletedUID.username, company_shareholder_histories.created_at, company_shareholder_histories.updated_at, company_shareholder_histories.deleted_at").Order(orderDirection).Limit(limit).Offset(offset).Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyShareholderHistoryConnection) CountSearchCompanyShareholderHistory(search string, companyId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_shareholder_histories").Select("company_shareholder_histories.id, company_shareholder_histories.company_id, companies.company_name, company_shareholder_histories.company_akte_id, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_shareholder_histories.remark, company_shareholder_histories.created_user_id, createdUID.username AS created_user, company_shareholder_histories.updated_user_id, updatedUID.username AS updated_user, company_shareholder_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholder_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholder_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholder_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join company_shareholder_history_details on company_shareholder_histories.id = company_shareholder_history_details.company_shareholder_history_id").Joins("left join companies on company_shareholder_histories.company_id = companies.id").Joins("left join company_aktes on company_shareholder_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_shareholder_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholder_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholder_histories.deleted_user_id = deletedUID.id").Having("(company_aktes.year::text LIKE ? OR lower(to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(company_aktes.akte_no) LIKE ? OR lower(string_agg(CASE WHEN company_shareholder_history_details.shareholder_name LIKE '%-DELETE' THEN NULL ELSE CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(CONCAT(company_shareholder_history_details.shareholder_name, ' : '), TO_CHAR(company_shareholder_history_details.number_of_share, 'fm999G999G999G999')),' saham ('), company_shareholder_history_details.percentage_of_share),'% equals to '), TO_CHAR(company_shareholder_history_details.share_amount, 'fm999G999G999G999')), ')') END, '<br/> ')) LIKE ? OR lower(company_shareholder_histories.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(company_shareholder_histories.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_shareholder_histories.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND company_shareholder_histories.company_id = ? AND company_shareholder_histories.deleted_at = 0", final, final, final, final, final, final, final, final, final, companyId).Group("company_shareholder_histories.id, company_shareholder_histories.company_id, companies.company_name, company_shareholder_histories.company_akte_id, company_aktes.akte_no, company_aktes.akte_date, company_aktes.year, company_shareholder_histories.remark, company_shareholder_histories.created_user_id, createdUID.username, company_shareholder_histories.updated_user_id, updatedUID.username, company_shareholder_histories.deleted_user_id, deletedUID.username, company_shareholder_histories.created_at, company_shareholder_histories.updated_at, company_shareholder_histories.deleted_at").Count(&count)
	return count, res.Error
}

func (db *CompanyShareholderHistoryConnection) FindCompanyShareholderHistoryById(id uint) (companyManagementOutput model.SelectCompanyShareholderHistoryParameter, err error) {
	var (
		companyManagement model.SelectCompanyShareholderHistoryParameter
	)

	res := db.connection.Debug().Table("company_shareholder_histories").Select("company_shareholder_histories.id, company_shareholder_histories.company_id, companies.company_name, company_shareholder_histories.company_akte_id, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_shareholder_histories.remark, company_shareholder_histories.created_user_id, createdUID.username AS created_user, company_shareholder_histories.updated_user_id, updatedUID.username AS updated_user, company_shareholder_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholder_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholder_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholder_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_shareholder_histories.company_id = companies.id").Joins("left join company_aktes on company_shareholder_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_shareholder_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholder_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholder_histories.deleted_user_id = deletedUID.id").Where("company_shareholder_histories.id=? AND company_shareholder_histories.deleted_at = 0", id).Take(&companyManagement)
	return companyManagement, res.Error
}

func (db *CompanyShareholderHistoryConnection) FindExcCompanyShareholderHistory(id uint) (companyManagementOutput []model.SelectCompanyShareholderHistoryParameter, err error) {
	var (
		companyManagements []model.SelectCompanyShareholderHistoryParameter
	)

	res := db.connection.Debug().Table("company_shareholder_histories").Select("company_shareholder_histories.id, company_shareholder_histories.company_id, companies.company_name, company_shareholder_histories.company_akte_id, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_shareholder_histories.remark, company_shareholder_histories.created_user_id, createdUID.username AS created_user, company_shareholder_histories.updated_user_id, updatedUID.username AS updated_user, company_shareholder_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholder_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholder_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholder_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_shareholder_histories.company_id = companies.id").Joins("left join company_aktes on company_shareholder_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_shareholder_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholder_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholder_histories.deleted_user_id = deletedUID.id").Where(" company_shareholder_histories.id!=? AND company_shareholder_histories.deleted_at = 0", id).Order("company_shareholder_histories.id").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyShareholderHistoryConnection) FindCompanyShareholderHistoryByCompanyId(id uint) (companyManagementOutput []model.SelectCompanyShareholderHistoryParameter, err error) {
	var (
		companyManagements []model.SelectCompanyShareholderHistoryParameter
	)

	res := db.connection.Debug().Table("company_shareholder_histories").Select("company_shareholder_histories.id, company_shareholder_histories.company_id, companies.company_name, company_shareholder_histories.company_akte_id, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_shareholder_histories.remark, company_shareholder_histories.created_user_id, createdUID.username AS created_user, company_shareholder_histories.updated_user_id, updatedUID.username AS updated_user, company_shareholder_histories.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholder_histories.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholder_histories.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholder_histories.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies on company_shareholder_histories.company_id = companies.id").Joins("left join company_aktes on company_shareholder_histories.company_akte_id = company_aktes.id").Joins("left join users createdUID on company_shareholder_histories.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholder_histories.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholder_histories.deleted_user_id = deletedUID.id").Where("company_shareholder_histories.company_id=? AND company_shareholder_histories.deleted_at = 0", id).Order("company_shareholder_histories.id").Find(&companyManagements)
	return companyManagements, res.Error
}

func (db *CompanyShareholderHistoryConnection) InsertCompanyShareholderHistory(companyManagement model.CompanyShareholderHistory) (companyManagementOutput model.CompanyShareholderHistory, err error) {
	res := db.connection.Save(&companyManagement)
	return companyManagement, res.Error
}

func (db *CompanyShareholderHistoryConnection) UpdateCompanyShareholderHistory(companyManagement model.CompanyShareholderHistory, id uint) (companyManagementOutput model.CompanyShareholderHistory, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyManagement)
	return companyManagement, res.Error
}
