package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyShareholderHistoryDetailRepository interface {
	CountCompanyShareholderHistoryDetailAll(companyShareholderHistoryId int) (count int64, err error)
	FindCompanyShareholderHistoryDetails(companyShareholderHistoryId int) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error)
	FindCompanyShareholderHistoryDetailsOffset(limit int, offset int, order string, dir string, companyShareholderHistoryId int) (companyShareholderHistoryDetailOutput []model.SelectCompanyShareholderHistoryDetail, err error)
	SearchCompanyShareholderHistoryDetail(limit int, offset int, order string, dir string, search string, companyShareholderHistoryId int) (companyShareholderHistoryDetailOutput []model.SelectCompanyShareholderHistoryDetail, err error)
	CountSearchCompanyShareholderHistoryDetail(search string, companyShareholderHistoryId int) (count int64, err error)
	FindCompanyShareholderHistoryDetailByCompanyId(id uint) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error)
	FindCompanyShareholderHistoryDetailById(id uint) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error)
	FindExcCompanyShareholderHistoryDetail(id uint) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error)
	InsertCompanyShareholderHistoryDetail(companyShareholderHistoryDetail model.CompanyShareholderHistoryDetail) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error)
	UpdateCompanyShareholderHistoryDetail(companyShareholderHistoryDetail model.CompanyShareholderHistoryDetail, id uint) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error)
}

type CompanyShareholderHistoryDetailConnection struct {
	connection *gorm.DB
}

func NewCompanyShareholderHistoryDetailRepository(db *gorm.DB) CompanyShareholderHistoryDetailRepository {
	return &CompanyShareholderHistoryDetailConnection{
		connection: db,
	}
}

func (db *CompanyShareholderHistoryDetailConnection) CountCompanyShareholderHistoryDetailAll(companyShareholderHistoryId int) (count int64, err error) {
	res := db.connection.Debug().Table("company_shareholder_history_details").Where("company_shareholder_history_id = ? AND deleted_at = 0", companyShareholderHistoryId).Count(&count)
	return count, res.Error
}

func (db *CompanyShareholderHistoryDetailConnection) FindCompanyShareholderHistoryDetails(companyShareholderHistoryId int) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error) {
	var (
		companyShareholderHistoryDetails []model.CompanyShareholderHistoryDetail
	)
	res := db.connection.Where("company_shareholder_history_id = ? AND deleted_at = 0", companyShareholderHistoryId).Order("shareholder_name").Find(&companyShareholderHistoryDetails)
	return companyShareholderHistoryDetails, res.Error
}

func (db *CompanyShareholderHistoryDetailConnection) FindCompanyShareholderHistoryDetailsOffset(limit int, offset int, order string, dir string, companyShareholderHistoryId int) (companyShareholderHistoryDetailOutput []model.SelectCompanyShareholderHistoryDetail, err error) {
	var (
		orderDirection                   string
		companyShareholderHistoryDetails []model.SelectCompanyShareholderHistoryDetail
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("company_shareholder_history_details").Select("company_shareholder_history_details.id, company_shareholder_history_details.company_shareholder_history_id, company_shareholder_history_details.shareholder_name,company_shareholder_history_details.number_of_share, company_shareholder_history_details.percentage_of_share, company_shareholder_history_details.share_amount, company_shareholder_history_details.remark, company_shareholder_history_details.created_user_id, createdUID.username AS created_user, company_shareholder_history_details.updated_user_id, updatedUID.username AS updated_user, company_shareholder_history_details.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholder_history_details.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholder_history_details.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholder_history_details.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on company_shareholder_history_details.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholder_history_details.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholder_history_details.deleted_user_id = deletedUID.id").Where("company_shareholder_history_details.company_shareholder_history_id = ? AND company_shareholder_history_details.deleted_at = 0", companyShareholderHistoryId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyShareholderHistoryDetails)
	return companyShareholderHistoryDetails, res.Error
}

func (db *CompanyShareholderHistoryDetailConnection) SearchCompanyShareholderHistoryDetail(limit int, offset int, order string, dir string, search string, companyShareholderHistoryId int) (companyShareholderHistoryDetailOutput []model.SelectCompanyShareholderHistoryDetail, err error) {
	var (
		orderDirection                   string
		final                            string
		companyShareholderHistoryDetails []model.SelectCompanyShareholderHistoryDetail
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_shareholder_history_details").Select("company_shareholder_history_details.id, company_shareholder_history_details.company_shareholder_history_id, company_shareholder_history_details.shareholder_name,company_shareholder_history_details.number_of_share, company_shareholder_history_details.percentage_of_share, company_shareholder_history_details.share_amount, company_shareholder_history_details.remark, company_shareholder_history_details.created_user_id, createdUID.username AS created_user, company_shareholder_history_details.updated_user_id, updatedUID.username AS updated_user, company_shareholder_history_details.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholder_history_details.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholder_history_details.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholder_history_details.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on company_shareholder_history_details.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholder_history_details.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholder_history_details.deleted_user_id = deletedUID.id").Where("(lower(company_shareholder_history_details.shareholder_name) LIKE ? OR lower(company_shareholder_history_details.number_of_share::varchar(50)) LIKE ? OR lower(company_shareholder_history_details.percentage_of_share::varchar(50)) LIKE ? OR lower(company_shareholder_history_details.share_amount::varchar(50)) LIKE ? OR lower(to_char(to_timestamp(company_shareholder_history_details.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_shareholder_history_details.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(company_shareholder_history_details.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? ) AND company_shareholder_history_details.company_shareholder_history_id = ? AND company_shareholder_history_details.deleted_at = 0", final, final, final, final, final, final, final, final, final, companyShareholderHistoryId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyShareholderHistoryDetails)
	return companyShareholderHistoryDetails, res.Error
}

func (db *CompanyShareholderHistoryDetailConnection) CountSearchCompanyShareholderHistoryDetail(search string, companyShareholderHistoryId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_shareholder_history_details").Select("company_shareholder_history_details.id, company_shareholder_history_details.company_shareholder_history_id, company_shareholder_history_details.shareholder_name,company_shareholder_history_details.number_of_share, company_shareholder_history_details.percentage_of_share, company_shareholder_history_details.share_amount, company_shareholder_history_details.remark, company_shareholder_history_details.created_user_id, createdUID.username AS created_user, company_shareholder_history_details.updated_user_id, updatedUID.username AS updated_user, company_shareholder_history_details.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholder_history_details.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholder_history_details.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholder_history_details.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on company_shareholder_history_details.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholder_history_details.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholder_history_details.deleted_user_id = deletedUID.id").Where("(lower(company_shareholder_history_details.shareholder_name) LIKE ? OR lower(company_shareholder_history_details.number_of_share::varchar(50)) LIKE ? OR lower(company_shareholder_history_details.percentage_of_share::varchar(50)) LIKE ? OR lower(company_shareholder_history_details.share_amount::varchar(50)) LIKE ? OR lower(to_char(to_timestamp(company_shareholder_history_details.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_shareholder_history_details.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(company_shareholder_history_details.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? ) AND company_shareholder_history_details.company_shareholder_history_id = ? AND company_shareholder_history_details.deleted_at = 0", final, final, final, final, final, final, final, final, final, companyShareholderHistoryId).Count(&count)
	return count, res.Error
}

func (db *CompanyShareholderHistoryDetailConnection) FindCompanyShareholderHistoryDetailByCompanyId(id uint) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error) {
	var (
		companyShareholderHistoryDetails []model.CompanyShareholderHistoryDetail
	)
	res := db.connection.Where("company_shareholder_history_id=? AND deleted_at = 0", id).Order("shareholder_name").Find(&companyShareholderHistoryDetails)
	return companyShareholderHistoryDetails, res.Error
}

func (db *CompanyShareholderHistoryDetailConnection) FindCompanyShareholderHistoryDetailById(id uint) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error) {
	var (
		companyShareholderHistoryDetail model.CompanyShareholderHistoryDetail
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&companyShareholderHistoryDetail)
	return companyShareholderHistoryDetail, res.Error
}

func (db *CompanyShareholderHistoryDetailConnection) FindExcCompanyShareholderHistoryDetail(id uint) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error) {
	var (
		companyShareholderHistoryDetails []model.CompanyShareholderHistoryDetail
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("shareholder_name").Find(&companyShareholderHistoryDetails)
	return companyShareholderHistoryDetails, res.Error
}

func (db *CompanyShareholderHistoryDetailConnection) InsertCompanyShareholderHistoryDetail(companyShareholderHistoryDetail model.CompanyShareholderHistoryDetail) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error) {
	res := db.connection.Save(&companyShareholderHistoryDetail)
	return companyShareholderHistoryDetail, res.Error
}

func (db *CompanyShareholderHistoryDetailConnection) UpdateCompanyShareholderHistoryDetail(companyShareholderHistoryDetail model.CompanyShareholderHistoryDetail, id uint) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyShareholderHistoryDetail)
	return companyShareholderHistoryDetail, res.Error
}
