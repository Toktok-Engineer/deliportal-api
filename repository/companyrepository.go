package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	CountCompanyAll() (count int64, err error)
	FindCompanys() (companyOutput []model.SelectCompanyParameter, err error)
	FindCompanyFilters(companyID string) (companyOutput []model.SelectCompanyParameter, err error)
	FindCompanysOffset(limit int, offset int, order string, dir string, companyID string) (companyOutput []model.SelectCompanyParameter, err error)
	SearchCompany(limit int, offset int, order string, dir string, search string, companyID string) (companyOutput []model.SelectCompanyParameter, err error)
	CountSearchCompany(search string, companyID string) (count int64, err error)
	CountCompanyApprove(companyID string) (count int64, err error)
	FindCompanyApprove(limit int, offset int, order string, dir string, companyID string) (companyOutput []model.SelectCompanyParameter, err error)
	SearchCompanyApprove(limit int, offset int, order string, dir string, search string, companyID string) (companyOutput []model.SelectCompanyParameter, err error)
	CountSearchCompanyApprove(search string, companyID string) (count int64, err error)
	FindCompanyById(id uint) (companyOutput model.SelectCompanyParameter, err error)
	FindExcCompany(id uint) (companyOutput []model.SelectCompanyParameter, err error)
	InsertCompany(company model.Company) (companyOutput model.Company, err error)
	UpdateCompany(company model.Company, id uint) (companyOutput model.Company, err error)
	UpdateCompanyApprove(company model.Company, id uint) (companyOutput model.Company, err error)
	UpdateCompanyDeactive(company model.Company, id uint) (companyOutput model.Company, err error)
}

type CompanyConnection struct {
	connection *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &CompanyConnection{
		connection: db,
	}
}

func (db *CompanyConnection) CountCompanyAll() (count int64, err error) {
	res := db.connection.Table("companies").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *CompanyConnection) FindCompanys() (companyOutput []model.SelectCompanyParameter, err error) {
	var (
		companys []model.SelectCompanyParameter
	)
	res := db.connection.Table("companies").Select("companies.id, companies.company_name, companies.previous_company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, CASE WHEN companies.status = 1 THEN 'Draft' WHEN companies.status = 2 THEN 'Need Approval' WHEN companies.status = 3 THEN 'Active' WHEN companies.status = 4 THEN 'Not Active' END AS status_name, companies.approved_user_id, appUID.username AS approved_user, companies.deactived_user_id, deactiveUID.username AS deactivate_user, to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY') AS approved_date, to_char(to_timestamp(companies.deactived_date::numeric), 'DD-Mon-YYYY') AS deactived_date, companies.remark, companies.created_user_id, createdUID.username AS created_user, companies.updated_user_id, updatedUID.username AS updated_user, companies.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(companies.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Joins("left join users appUID on companies.approved_user_id = appUID.id").Joins("left join users deactiveUID on companies.deactived_user_id = deactiveUID.id").Joins("left join users createdUID on companies.created_user_id = createdUID.id").Joins("left join users updatedUID on companies.updated_user_id = updatedUID.id").Joins("left join users deletedUID on companies.deleted_user_id = deletedUID.id").Where("companies.deleted_at = 0").Order("companies.company_name").Find(&companys)
	return companys, res.Error
}

func (db *CompanyConnection) FindCompanyFilters(companyID string) (companyOutput []model.SelectCompanyParameter, err error) {
	var (
		companys []model.SelectCompanyParameter
	)
	compId := strings.Split(companyID, ",")
	res := db.connection.Table("companies").Select("companies.id, companies.company_name, companies.previous_company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, CASE WHEN companies.status = 1 THEN 'Draft' WHEN companies.status = 2 THEN 'Need Approval' WHEN companies.status = 3 THEN 'Active' WHEN companies.status = 4 THEN 'Not Active' END AS status_name, companies.approved_user_id, appUID.username AS approved_user, companies.deactived_user_id, deactiveUID.username AS deactivate_user, to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY') AS approved_date, to_char(to_timestamp(companies.deactived_date::numeric), 'DD-Mon-YYYY') AS deactived_date, companies.remark, companies.created_user_id, createdUID.username AS created_user, companies.updated_user_id, updatedUID.username AS updated_user, companies.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(companies.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Joins("left join users appUID on companies.approved_user_id = appUID.id").Joins("left join users deactiveUID on companies.deactived_user_id = deactiveUID.id").Joins("left join users createdUID on companies.created_user_id = createdUID.id").Joins("left join users updatedUID on companies.updated_user_id = updatedUID.id").Joins("left join users deletedUID on companies.deleted_user_id = deletedUID.id").Where("companies.id NOT IN (?) AND companies.deleted_at = 0", compId).Order("companies.company_name").Find(&companys)
	return companys, res.Error
}

func (db *CompanyConnection) FindCompanysOffset(limit int, offset int, order string, dir string, companyID string) (companyOutput []model.SelectCompanyParameter, err error) {
	var (
		orderDirection string
		companys       []model.SelectCompanyParameter
	)
	orderDirection = order + " " + dir
	compId := strings.Split(companyID, ",")
	res := db.connection.Table("companies").Select("companies.id, companies.company_name, companies.previous_company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, CASE WHEN companies.status = 1 THEN 'Draft' WHEN companies.status = 2 THEN 'Need Approval' WHEN companies.status = 3 THEN 'Active' WHEN companies.status = 4 THEN 'Not Active' END AS status_name, companies.approved_user_id, appUID.username AS approved_user, companies.deactived_user_id, deactiveUID.username AS deactivate_user, to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY') AS approved_date, to_char(to_timestamp(companies.deactived_date::numeric), 'DD-Mon-YYYY') AS deactived_date, companies.remark, companies.created_user_id, createdUID.username AS created_user, companies.updated_user_id, updatedUID.username AS updated_user, companies.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(companies.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Joins("left join users appUID on companies.approved_user_id = appUID.id").Joins("left join users deactiveUID on companies.deactived_user_id = deactiveUID.id").Joins("left join users createdUID on companies.created_user_id = createdUID.id").Joins("left join users updatedUID on companies.updated_user_id = updatedUID.id").Joins("left join users deletedUID on companies.deleted_user_id = deletedUID.id").Where("companies.id NOT IN (?) AND companies.deleted_at = 0", compId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companys)
	return companys, res.Error
}

func (db *CompanyConnection) SearchCompany(limit int, offset int, order string, dir string, search string, companyID string) (companyOutput []model.SelectCompanyParameter, err error) {
	var (
		orderDirection string
		final          string
		companys       []model.SelectCompanyParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	compId := strings.Split(companyID, ",")
	res := db.connection.Table("companies").Select("companies.id, companies.company_name, companies.previous_company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, CASE WHEN companies.status = 1 THEN 'Draft' WHEN companies.status = 2 THEN 'Need Approval' WHEN companies.status = 3 THEN 'Active' WHEN companies.status = 4 THEN 'Not Active' END AS status_name, companies.approved_user_id, appUID.username AS approved_user, companies.deactived_user_id, deactiveUID.username AS deactivate_user, to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY') AS approved_date, to_char(to_timestamp(companies.deactived_date::numeric), 'DD-Mon-YYYY') AS deactived_date, companies.remark, companies.created_user_id, createdUID.username AS created_user, companies.updated_user_id, updatedUID.username AS updated_user, companies.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(companies.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Joins("left join users appUID on companies.approved_user_id = appUID.id").Joins("left join users deactiveUID on companies.deactived_user_id = deactiveUID.id").Joins("left join users createdUID on companies.created_user_id = createdUID.id").Joins("left join users updatedUID on companies.updated_user_id = updatedUID.id").Joins("left join users deletedUID on companies.deleted_user_id = deletedUID.id").Where("(lower(companies.company_name) LIKE ? OR lower(business_units.business_unit_name) LIKE ? OR lower(companies.address) LIKE ? OR lower(companies.legal_license_file_url) LIKE ? OR (CASE WHEN companies.status = 1 THEN 'draft' WHEN companies.status = 2 THEN 'need approval' WHEN companies.status = 3 THEN 'active' WHEN companies.status = 4 THEN 'not active' END LIKE ?) OR lower(appUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(companies.remark) LIKE ? OR lower(createdUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND companies.id NOT IN (?) AND companies.deleted_at = 0", final, final, final, final, final, final, final, final, final, final, final, final, compId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companys)
	return companys, res.Error
}

func (db *CompanyConnection) CountSearchCompany(search string, companyID string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	compId := strings.Split(companyID, ",")
	res := db.connection.Table("companies").Select("companies.id, companies.company_name, companies.previous_company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, CASE WHEN companies.status = 1 THEN 'Draft' WHEN companies.status = 2 THEN 'Need Approval' WHEN companies.status = 3 THEN 'Active' WHEN companies.status = 4 THEN 'Not Active' END AS status_name, companies.approved_user_id, appUID.username AS approved_user, companies.deactived_user_id, deactiveUID.username AS deactivate_user, to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY') AS approved_date, to_char(to_timestamp(companies.deactived_date::numeric), 'DD-Mon-YYYY') AS deactived_date, companies.remark, companies.created_user_id, createdUID.username AS created_user, companies.updated_user_id, updatedUID.username AS updated_user, companies.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(companies.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Joins("left join users appUID on companies.approved_user_id = appUID.id").Joins("left join users deactiveUID on companies.deactived_user_id = deactiveUID.id").Joins("left join users createdUID on companies.created_user_id = createdUID.id").Joins("left join users updatedUID on companies.updated_user_id = updatedUID.id").Joins("left join users deletedUID on companies.deleted_user_id = deletedUID.id").Where("(lower(companies.company_name) LIKE ? OR lower(business_units.business_unit_name) LIKE ? OR lower(companies.address) LIKE ? OR lower(companies.legal_license_file_url) LIKE ? OR (CASE WHEN companies.status = 1 THEN 'draft' WHEN companies.status = 2 THEN 'need approval' WHEN companies.status = 3 THEN 'active' WHEN companies.status = 4 THEN 'not active' END LIKE ?) OR lower(appUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(companies.remark) LIKE ? OR lower(createdUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND companies.id NOT IN (?) AND companies.deleted_at = 0", final, final, final, final, final, final, final, final, final, final, final, final, compId).Count(&count)
	return count, res.Error
}

func (db *CompanyConnection) CountCompanyApprove(companyID string) (count int64, err error) {
	compId := strings.Split(companyID, ",")
	res := db.connection.Table("companies").Where("companies.id NOT IN (?) AND companies.status = 2 AND deleted_at = 0", compId).Count(&count)
	return count, res.Error
}

func (db *CompanyConnection) FindCompanyApprove(limit int, offset int, order string, dir string, companyID string) (companyOutput []model.SelectCompanyParameter, err error) {
	var (
		orderDirection string
		companys       []model.SelectCompanyParameter
	)
	orderDirection = order + " " + dir
	compId := strings.Split(companyID, ",")
	res := db.connection.Table("companies").Select("companies.id, companies.company_name, companies.previous_company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, CASE WHEN companies.status = 1 THEN 'Draft' WHEN companies.status = 2 THEN 'Need Approval' WHEN companies.status = 3 THEN 'Active' WHEN companies.status = 4 THEN 'Not Active' END AS status_name, companies.approved_user_id, appUID.username AS approved_user, companies.deactived_user_id, deactiveUID.username AS deactivate_user, to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY') AS approved_date, to_char(to_timestamp(companies.deactived_date::numeric), 'DD-Mon-YYYY') AS deactived_date, companies.remark, companies.created_user_id, createdUID.username AS created_user, companies.updated_user_id, updatedUID.username AS updated_user, companies.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(companies.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Joins("left join users appUID on companies.approved_user_id = appUID.id").Joins("left join users deactiveUID on companies.deactived_user_id = deactiveUID.id").Joins("left join users createdUID on companies.created_user_id = createdUID.id").Joins("left join users updatedUID on companies.updated_user_id = updatedUID.id").Joins("left join users deletedUID on companies.deleted_user_id = deletedUID.id").Where("companies.id NOT IN (?) AND companies.status = 2 AND companies.deleted_at = 0", compId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companys)
	return companys, res.Error
}

func (db *CompanyConnection) SearchCompanyApprove(limit int, offset int, order string, dir string, search string, companyID string) (companyOutput []model.SelectCompanyParameter, err error) {
	var (
		orderDirection string
		final          string
		companys       []model.SelectCompanyParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	compId := strings.Split(companyID, ",")
	res := db.connection.Table("companies").Select("companies.id, companies.company_name, companies.previous_company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, CASE WHEN companies.status = 1 THEN 'Draft' WHEN companies.status = 2 THEN 'Need Approval' WHEN companies.status = 3 THEN 'Active' WHEN companies.status = 4 THEN 'Not Active' END AS status_name, companies.approved_user_id, appUID.username AS approved_user, companies.deactived_user_id, deactiveUID.username AS deactivate_user, to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY') AS approved_date, to_char(to_timestamp(companies.deactived_date::numeric), 'DD-Mon-YYYY') AS deactived_date, companies.remark, companies.created_user_id, createdUID.username AS created_user, companies.updated_user_id, updatedUID.username AS updated_user, companies.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(companies.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Joins("left join users appUID on companies.approved_user_id = appUID.id").Joins("left join users deactiveUID on companies.deactived_user_id = deactiveUID.id").Joins("left join users createdUID on companies.created_user_id = createdUID.id").Joins("left join users updatedUID on companies.updated_user_id = updatedUID.id").Joins("left join users deletedUID on companies.deleted_user_id = deletedUID.id").Where("(lower(companies.company_name) LIKE ? OR lower(business_units.business_unit_name) LIKE ? OR lower(companies.address) LIKE ? OR lower(companies.legal_license_file_url) LIKE ? OR (CASE WHEN companies.status = 1 THEN 'draft' WHEN companies.status = 2 THEN 'need approval' WHEN companies.status = 3 THEN 'active' WHEN companies.status = 4 THEN 'not active' END LIKE ?) OR lower(appUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(companies.remark) LIKE ? OR lower(createdUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND companies.id NOT IN (?) AND companies.status = 2 AND companies.deleted_at = 0", final, final, final, final, final, final, final, final, final, final, final, final, compId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companys)
	return companys, res.Error
}

func (db *CompanyConnection) CountSearchCompanyApprove(search string, companyID string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	compId := strings.Split(companyID, ",")
	res := db.connection.Table("companies").Select("companies.id, companies.company_name, companies.previous_company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, CASE WHEN companies.status = 1 THEN 'Draft' WHEN companies.status = 2 THEN 'Need Approval' WHEN companies.status = 3 THEN 'Active' WHEN companies.status = 4 THEN 'Not Active' END AS status_name, companies.approved_user_id, appUID.username AS approved_user, companies.deactived_user_id, deactiveUID.username AS deactivate_user, to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY') AS approved_date, to_char(to_timestamp(companies.deactived_date::numeric), 'DD-Mon-YYYY') AS deactived_date, companies.remark, companies.created_user_id, createdUID.username AS created_user, companies.updated_user_id, updatedUID.username AS updated_user, companies.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(companies.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Joins("left join users appUID on companies.approved_user_id = appUID.id").Joins("left join users deactiveUID on companies.deactived_user_id = deactiveUID.id").Joins("left join users createdUID on companies.created_user_id = createdUID.id").Joins("left join users updatedUID on companies.updated_user_id = updatedUID.id").Joins("left join users deletedUID on companies.deleted_user_id = deletedUID.id").Where("(lower(companies.company_name) LIKE ? OR lower(business_units.business_unit_name) LIKE ? OR lower(companies.address) LIKE ? OR lower(companies.legal_license_file_url) LIKE ? OR (CASE WHEN companies.status = 1 THEN 'draft' WHEN companies.status = 2 THEN 'need approval' WHEN companies.status = 3 THEN 'active' WHEN companies.status = 4 THEN 'not active' END LIKE ?) OR lower(appUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(companies.remark) LIKE ? OR lower(createdUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ?) AND companies.id NOT IN (?) AND companies.status = 2 AND companies.deleted_at = 0", final, final, final, final, final, final, final, final, final, final, final, final, compId).Count(&count)
	return count, res.Error
}

func (db *CompanyConnection) FindCompanyById(id uint) (companyOutput model.SelectCompanyParameter, err error) {
	var (
		company model.SelectCompanyParameter
	)
	res := db.connection.Table("companies").Select("companies.id, companies.company_name, companies.previous_company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, CASE WHEN companies.status = 1 THEN 'Draft' WHEN companies.status = 2 THEN 'Need Approval' WHEN companies.status = 3 THEN 'Active' WHEN companies.status = 4 THEN 'Not Active' END AS status_name, companies.approved_user_id, appUID.username AS approved_user, companies.deactived_user_id, deactiveUID.username AS deactivate_user, to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY') AS approved_date, to_char(to_timestamp(companies.deactived_date::numeric), 'DD-Mon-YYYY') AS deactived_date, companies.remark, companies.created_user_id, createdUID.username AS created_user, companies.updated_user_id, updatedUID.username AS updated_user, companies.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(companies.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Joins("left join users appUID on companies.approved_user_id = appUID.id").Joins("left join users deactiveUID on companies.deactived_user_id = deactiveUID.id").Joins("left join users createdUID on companies.created_user_id = createdUID.id").Joins("left join users updatedUID on companies.updated_user_id = updatedUID.id").Joins("left join users deletedUID on companies.deleted_user_id = deletedUID.id").Where("companies.id=? AND companies.deleted_at = 0", id).Take(&company)
	return company, res.Error
}

func (db *CompanyConnection) FindExcCompany(id uint) (companyOutput []model.SelectCompanyParameter, err error) {
	var (
		companys []model.SelectCompanyParameter
	)
	res := db.connection.Table("companies").Select("companies.id, companies.company_name, companies.previous_company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, CASE WHEN companies.status = 1 THEN 'Draft' WHEN companies.status = 2 THEN 'Need Approval' WHEN companies.status = 3 THEN 'Active' WHEN companies.status = 4 THEN 'Not Active' END AS status_name, companies.approved_user_id, appUID.username AS approved_user, companies.deactived_user_id, deactiveUID.username AS deactivate_user, to_char(to_timestamp(companies.approved_date::numeric), 'DD-Mon-YYYY') AS approved_date, to_char(to_timestamp(companies.deactived_date::numeric), 'DD-Mon-YYYY') AS deactived_date, companies.remark, companies.created_user_id, createdUID.username AS created_user, companies.updated_user_id, updatedUID.username AS updated_user, companies.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(companies.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(companies.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(companies.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Joins("left join users appUID on companies.approved_user_id = appUID.id").Joins("left join users deactiveUID on companies.deactived_user_id = deactiveUID.id").Joins("left join users createdUID on companies.created_user_id = createdUID.id").Joins("left join users updatedUID on companies.updated_user_id = updatedUID.id").Joins("left join users deletedUID on companies.deleted_user_id = deletedUID.id").Where("companies.id!=? AND companies.deleted_at = 0", id).Order("companies.company_name").Find(&companys)
	return companys, res.Error
}

func (db *CompanyConnection) InsertCompany(company model.Company) (companyOutput model.Company, err error) {
	res := db.connection.Save(&company)
	return company, res.Error
}

func (db *CompanyConnection) UpdateCompany(company model.Company, id uint) (companyOutput model.Company, err error) {
	res := db.connection.Where("id=?", id).Updates(&company)
	return company, res.Error
}

func (db *CompanyConnection) UpdateCompanyApprove(company model.Company, id uint) (companyOutput model.Company, err error) {
	res := db.connection.Model(&company).Where("id=?", id).Updates(map[string]interface{}{"status": company.Status, "approved_user_id": company.ApprovedUserID, "approved_date": company.ApprovedDate})
	return company, res.Error
}

func (db *CompanyConnection) UpdateCompanyDeactive(company model.Company, id uint) (companyOutput model.Company, err error) {
	res := db.connection.Model(&company).Where("id=?", id).Updates(map[string]interface{}{"status": company.Status, "deactived_user_id": company.DeactivedUserID, "deactived_date": company.DeactivedDate})
	return company, res.Error
}
