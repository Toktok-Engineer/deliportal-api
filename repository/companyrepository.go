package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	FindCompanys() (companyOutput []model.SelectCompanyParameter, err error)
	FindCompanyApprove() (companyOutput []model.SelectCompanyParameter, err error)
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

func (db *CompanyConnection) FindCompanys() (companyOutput []model.SelectCompanyParameter, err error) {
	var (
		companys []model.SelectCompanyParameter
	)
	res := db.connection.Debug().Table("companies").Select("companies.id, companies.company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, companies.approved_user_id, companies.deactived_user_id, companies.approved_date, companies.deactived_date, companies.remark, companies.created_user_id, companies.updated_user_id, companies.deleted_user_id, companies.created_at, companies.updated_at, companies.deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Where("companies.deleted_at = 0").Order("companies.company_name").Find(&companys)
	return companys, res.Error
}

func (db *CompanyConnection) FindCompanyApprove() (companyOutput []model.SelectCompanyParameter, err error) {
	var (
		companys []model.SelectCompanyParameter
	)
	res := db.connection.Debug().Table("companies").Select("companies.id, companies.company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, companies.approved_user_id, companies.deactived_user_id, companies.approved_date, companies.deactived_date, companies.remark, companies.created_user_id, companies.updated_user_id, companies.deleted_user_id, companies.created_at, companies.updated_at, companies.deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Where("companies.status = 2 AND companies.deleted_at = 0").Order("companies.company_name").Find(&companys)
	return companys, res.Error
}

func (db *CompanyConnection) FindCompanyById(id uint) (companyOutput model.SelectCompanyParameter, err error) {
	var (
		company model.SelectCompanyParameter
	)
	res := db.connection.Debug().Table("companies").Select("companies.id, companies.company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, companies.approved_user_id, companies.deactived_user_id, companies.approved_date, companies.deactived_date, companies.remark, companies.created_user_id, companies.updated_user_id, companies.deleted_user_id, companies.created_at, companies.updated_at, companies.deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Where("companies.id=? AND companies.deleted_at = 0", id).Take(&company)
	return company, res.Error
}

func (db *CompanyConnection) FindExcCompany(id uint) (companyOutput []model.SelectCompanyParameter, err error) {
	var (
		companys []model.SelectCompanyParameter
	)
	res := db.connection.Debug().Table("companies").Select("companies.id, companies.company_name, companies.business_unit_id, business_units.business_unit_name, companies.address, companies.legal_license_file_url, companies.status, companies.approved_user_id, companies.deactived_user_id, companies.approved_date, companies.deactived_date, companies.remark, companies.created_user_id, companies.updated_user_id, companies.deleted_user_id, companies.created_at, companies.updated_at, companies.deleted_at").Joins("left join business_units ON companies.business_unit_id = business_units.id").Where("companies.id!=? AND companies.deleted_at = 0", id).Order("companies.company_name").Find(&companys)
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
