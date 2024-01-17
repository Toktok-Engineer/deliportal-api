package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyShareholderRepository interface {
	CountCompanyShareholderAll(companyId int) (count int64, err error)
	FindCompanyShareholders(companyId int) (companyShareholderOutput []model.CompanyShareholder, err error)
	FindCompanyShareholdersOffset(limit int, offset int, order string, dir string, companyId int) (companyShareholderOutput []model.SelectCompanyShareholder, err error)
	SearchCompanyShareholder(limit int, offset int, order string, dir string, search string, companyId int) (companyShareholderOutput []model.SelectCompanyShareholder, err error)
	CountSearchCompanyShareholder(search string, companyId int) (count int64, err error)
	FindCompanyShareholderByCompanyId(id uint) (companyShareholderOutput []model.CompanyShareholder, err error)
	FindCompanyShareholderById(id uint) (companyShareholderOutput model.CompanyShareholder, err error)
	FindExcCompanyShareholder(id uint) (companyShareholderOutput []model.CompanyShareholder, err error)
	InsertCompanyShareholder(companyShareholder model.CompanyShareholder) (companyShareholderOutput model.CompanyShareholder, err error)
	UpdateCompanyShareholder(companyShareholder model.CompanyShareholder, id uint) (companyShareholderOutput model.CompanyShareholder, err error)
	ReportCompanyShareholder(year int, groupId int) (companyShareholderOutput []model.SelectCompanyShareholderReport, err error)
}

type CompanyShareholderConnection struct {
	connection *gorm.DB
}

func NewCompanyShareholderRepository(db *gorm.DB) CompanyShareholderRepository {
	return &CompanyShareholderConnection{
		connection: db,
	}
}

func (db *CompanyShareholderConnection) CountCompanyShareholderAll(companyId int) (count int64, err error) {
	res := db.connection.Debug().Table("company_shareholders").Where("company_id = ? AND deleted_at = 0", companyId).Count(&count)
	return count, res.Error
}

func (db *CompanyShareholderConnection) FindCompanyShareholders(companyId int) (companyShareholderOutput []model.CompanyShareholder, err error) {
	var (
		companyShareholders []model.CompanyShareholder
	)
	res := db.connection.Where("company_id = ? AND deleted_at = 0", companyId).Order("shareholder_name").Find(&companyShareholders)
	return companyShareholders, res.Error
}

func (db *CompanyShareholderConnection) FindCompanyShareholdersOffset(limit int, offset int, order string, dir string, companyId int) (companyShareholderOutput []model.SelectCompanyShareholder, err error) {
	var (
		orderDirection      string
		companyShareholders []model.SelectCompanyShareholder
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("company_shareholders").Select("company_shareholders.id, company_shareholders.company_id, company_shareholders.shareholder_name,company_shareholders.number_of_share, company_shareholders.percentage_of_share, company_shareholders.share_amount, company_shareholders.remark, company_shareholders.created_user_id, createdUID.username AS created_user, company_shareholders.updated_user_id, updatedUID.username AS updated_user, company_shareholders.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholders.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholders.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholders.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on company_shareholders.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholders.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholders.deleted_user_id = deletedUID.id").Where("company_shareholders.company_id = ? AND company_shareholders.deleted_at = 0", companyId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyShareholders)
	return companyShareholders, res.Error
}

func (db *CompanyShareholderConnection) SearchCompanyShareholder(limit int, offset int, order string, dir string, search string, companyId int) (companyShareholderOutput []model.SelectCompanyShareholder, err error) {
	var (
		orderDirection      string
		final               string
		companyShareholders []model.SelectCompanyShareholder
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_shareholders").Select("company_shareholders.id, company_shareholders.company_id, company_shareholders.shareholder_name,company_shareholders.number_of_share, company_shareholders.percentage_of_share, company_shareholders.share_amount, company_shareholders.remark, company_shareholders.created_user_id, createdUID.username AS created_user, company_shareholders.updated_user_id, updatedUID.username AS updated_user, company_shareholders.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholders.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholders.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholders.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on company_shareholders.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholders.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholders.deleted_user_id = deletedUID.id").Where("(lower(company_shareholders.shareholder_name) LIKE ? OR lower(company_shareholders.number_of_share::varchar(50)) LIKE ? OR lower(company_shareholders.percentage_of_share::varchar(50)) LIKE ? OR lower(company_shareholders.share_amount::varchar(50)) LIKE ? OR lower(to_char(to_timestamp(company_shareholders.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_shareholders.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(company_shareholders.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? ) AND company_shareholders.company_id = ? AND company_shareholders.deleted_at = 0", final, final, final, final, final, final, final, final, final, companyId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyShareholders)
	return companyShareholders, res.Error
}

func (db *CompanyShareholderConnection) CountSearchCompanyShareholder(search string, companyId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_shareholders").Select("company_shareholders.id, company_shareholders.company_id, company_shareholders.shareholder_name,company_shareholders.number_of_share, company_shareholders.percentage_of_share, company_shareholders.share_amount, company_shareholders.remark, company_shareholders.created_user_id, createdUID.username AS created_user, company_shareholders.updated_user_id, updatedUID.username AS updated_user, company_shareholders.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_shareholders.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_shareholders.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_shareholders.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join users createdUID on company_shareholders.created_user_id = createdUID.id").Joins("left join users updatedUID on company_shareholders.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_shareholders.deleted_user_id = deletedUID.id").Where("(lower(company_shareholders.shareholder_name) LIKE ? OR lower(company_shareholders.number_of_share::varchar(50)) LIKE ? OR lower(company_shareholders.percentage_of_share::varchar(50)) LIKE ? OR lower(company_shareholders.share_amount::varchar(50)) LIKE ? OR lower(to_char(to_timestamp(company_shareholders.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(company_shareholders.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(company_shareholders.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? ) AND company_shareholders.company_id = ? AND company_shareholders.deleted_at = 0", final, final, final, final, final, final, final, final, final, companyId).Count(&count)
	return count, res.Error
}

func (db *CompanyShareholderConnection) FindCompanyShareholderByCompanyId(id uint) (companyShareholderOutput []model.CompanyShareholder, err error) {
	var (
		companyShareholders []model.CompanyShareholder
	)
	res := db.connection.Where("company_id=? AND deleted_at = 0", id).Order("shareholder_name").Find(&companyShareholders)
	return companyShareholders, res.Error
}

func (db *CompanyShareholderConnection) FindCompanyShareholderById(id uint) (companyShareholderOutput model.CompanyShareholder, err error) {
	var (
		companyShareholder model.CompanyShareholder
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&companyShareholder)
	return companyShareholder, res.Error
}

func (db *CompanyShareholderConnection) FindExcCompanyShareholder(id uint) (companyShareholderOutput []model.CompanyShareholder, err error) {
	var (
		companyShareholders []model.CompanyShareholder
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("shareholder_name").Find(&companyShareholders)
	return companyShareholders, res.Error
}

func (db *CompanyShareholderConnection) InsertCompanyShareholder(companyShareholder model.CompanyShareholder) (companyShareholderOutput model.CompanyShareholder, err error) {
	res := db.connection.Save(&companyShareholder)
	return companyShareholder, res.Error
}

func (db *CompanyShareholderConnection) UpdateCompanyShareholder(companyShareholder model.CompanyShareholder, id uint) (companyShareholderOutput model.CompanyShareholder, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyShareholder)
	return companyShareholder, res.Error
}

func (db *CompanyShareholderConnection) ReportCompanyShareholder(year int, groupId int) (companyShareholderOutput []model.SelectCompanyShareholderReport, err error) {
	var (
		companyShareholder []model.SelectCompanyShareholderReport
	)
	res := db.connection.Raw("SELECT * FROM public.function_generate_shareholder_excel($1, $2) ORDER BY company_shareholder_id", year, groupId).Find(&companyShareholder)
	return companyShareholder, res.Error
}
