package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type CompanyLicenseRepository interface {
	FindCompanyLicenses() (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	FindCompanyLicenseApp() (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	FindExpCompanyLicenses() (companyLicenseOutput []model.SelectCompanyLicenseExpiredParameter, err error)
	FindCompanyLicenseById(id uint) (companyLicenseOutput model.SelectCompanyLicenseParameter, err error)
	FindExcCompanyLicense(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	FindCompanyLicenseByCompanyId(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	InsertCompanyLicense(companyLicense model.CompanyLicense) (companyLicenseOutput model.CompanyLicense, err error)
	UpdateCompanyLicense(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
	UpdateCompanyLicenseStatus(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
	UpdateCompanyLicenseDeactive(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
	UpdateCompanyLicenseApprovedRenewalStatus(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
	DeleteCompanyLicense(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
	UpdateCompanyRemark(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
}

type CompanyLicenseConnection struct {
	connection *gorm.DB
}

func NewCompanyLicenseRepository(db *gorm.DB) CompanyLicenseRepository {
	return &CompanyLicenseConnection{
		connection: db,
	}
}

func (db *CompanyLicenseConnection) FindCompanyLicenses() (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	var (
		companyLicenses []model.SelectCompanyLicenseParameter
	)

	res := db.connection.Debug().Table("company_licenses").Select("company_licenses.id, company_licenses.parent_license_id, company_licenses.license_no, company_licenses.license_type_id, license_types.license_type_name, company_licenses.company_id, company_licenses.renewable, company_licenses.reminder_counter, company_licenses.issued_by, company_licenses.issued_date, company_licenses.expired_date, company_licenses.earliest_renewal_date, company_licenses.last_renewal_date, company_licenses.status, company_licenses.renewal_status, company_licenses.approved_user_id, company_licenses.renewal_approved_user_id, company_licenses.approved_date, company_licenses.renewal_approved_date, company_licenses.remark, company_licenses.created_user_id, company_licenses.updated_user_id, company_licenses.deleted_user_id, company_licenses.created_at, company_licenses.updated_at, company_licenses.deleted_at").Joins("left join license_types ON company_licenses.license_type_id = license_types.id").Where("company_licenses.deleted_at = 0").Order("company_licenses.id").Find(&companyLicenses)
	return companyLicenses, res.Error
}

func (db *CompanyLicenseConnection) FindCompanyLicenseApp() (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	var (
		companyLicenses []model.SelectCompanyLicenseParameter
	)

	res := db.connection.Debug().Table("company_licenses").Select("company_licenses.id, company_licenses.parent_license_id, company_licenses.license_no, company_licenses.license_type_id, license_types.license_type_name, company_licenses.company_id, company_licenses.renewable, company_licenses.reminder_counter, company_licenses.issued_by, company_licenses.issued_date, company_licenses.expired_date, company_licenses.earliest_renewal_date, company_licenses.last_renewal_date, company_licenses.status, company_licenses.renewal_status, company_licenses.approved_user_id, company_licenses.renewal_approved_user_id, company_licenses.approved_date, company_licenses.renewal_approved_date, company_licenses.remark, company_licenses.created_user_id, company_licenses.updated_user_id, company_licenses.deleted_user_id, company_licenses.created_at, company_licenses.updated_at, company_licenses.deleted_at").Joins("left join license_types ON company_licenses.license_type_id = license_types.id").Where("company_licenses.status = 2 OR company_licenses.renewal_status = 3 OR company_licenses.renewal_status = 4 AND company_licenses.deleted_at = 0").Order("company_licenses.id").Find(&companyLicenses)
	return companyLicenses, res.Error
}

func (db *CompanyLicenseConnection) FindExpCompanyLicenses() (companyLicenseOutput []model.SelectCompanyLicenseExpiredParameter, err error) {
	var (
		companyLicenses []model.SelectCompanyLicenseExpiredParameter
	)

	res := db.connection.Debug().Table("company_licenses").Select("company_licenses.id, license_types.license_type_name, company_licenses.license_no, companies.company_name, company_licenses.expired_date, company_licenses.earliest_renewal_date, company_licenses.last_renewal_date, company_licenses.renewal_status, company_licenses.status, company_licenses.reminder_counter, company_licenses.remark, company_licenses.created_user_id, company_licenses.updated_user_id, company_licenses.deleted_user_id, company_licenses.created_at, company_licenses.updated_at, company_licenses.deleted_at").Joins("inner join license_types ON company_licenses.license_type_id = license_types.id").Joins("inner join companies ON company_licenses.company_id = companies.id").Where("company_licenses.renewable = true AND company_licenses.renewal_status not in (5,6) AND company_licenses.status = 3 AND date(now()) >= date((timezone('Asia/Jakarta', to_timestamp(company_licenses.earliest_renewal_date)) - (license_types.reminder_before_month::text || ' month')::interval)) AND company_licenses.deleted_at = 0").Order("company_licenses.expired_date").Find(&companyLicenses)
	return companyLicenses, res.Error
}

func (db *CompanyLicenseConnection) FindCompanyLicenseById(id uint) (companyLicenseOutput model.SelectCompanyLicenseParameter, err error) {
	var (
		companyLicense model.SelectCompanyLicenseParameter
	)

	res := db.connection.Debug().Table("company_licenses").Select("company_licenses.id, company_licenses.parent_license_id, company_licenses.license_no, company_licenses.license_type_id, license_types.license_type_name, company_licenses.company_id, company_licenses.renewable, company_licenses.reminder_counter, company_licenses.issued_by, company_licenses.issued_date, company_licenses.expired_date, company_licenses.earliest_renewal_date, company_licenses.last_renewal_date, company_licenses.status, company_licenses.renewal_status, company_licenses.approved_user_id, company_licenses.renewal_approved_user_id, company_licenses.approved_date, company_licenses.renewal_approved_date, company_licenses.remark, company_licenses.created_user_id, company_licenses.updated_user_id, company_licenses.deleted_user_id, company_licenses.created_at, company_licenses.updated_at, company_licenses.deleted_at").Joins("left join license_types ON company_licenses.license_type_id = license_types.id").Where("company_licenses.id=? AND company_licenses.deleted_at = 0", id).Take(&companyLicense)
	return companyLicense, res.Error
}

func (db *CompanyLicenseConnection) FindExcCompanyLicense(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	var (
		companyLicenses []model.SelectCompanyLicenseParameter
	)

	res := db.connection.Debug().Table("company_licenses").Select("company_licenses.id, company_licenses.parent_license_id, company_licenses.license_no, company_licenses.license_type_id, license_types.license_type_name, company_licenses.company_id, company_licenses.renewable, company_licenses.reminder_counter, company_licenses.issued_by, company_licenses.issued_date, company_licenses.expired_date, company_licenses.earliest_renewal_date, company_licenses.last_renewal_date, company_licenses.status, company_licenses.renewal_status, company_licenses.approved_user_id, company_licenses.renewal_approved_user_id, company_licenses.approved_date, company_licenses.renewal_approved_date, company_licenses.remark, company_licenses.created_user_id, company_licenses.updated_user_id, company_licenses.deleted_user_id, company_licenses.created_at, company_licenses.updated_at, company_licenses.deleted_at").Joins("left join license_types ON company_licenses.license_type_id = license_types.id").Where("company_licenses.id!=? AND company_licenses.deleted_at = 0", id).Order("company_licenses.id").Find(&companyLicenses)
	return companyLicenses, res.Error
}

func (db *CompanyLicenseConnection) FindCompanyLicenseByCompanyId(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	var (
		companyLicenses []model.SelectCompanyLicenseParameter
	)

	res := db.connection.Debug().Table("company_licenses").Select("company_licenses.id, company_licenses.parent_license_id, company_licenses.license_no, company_licenses.license_type_id, license_types.license_type_name, company_licenses.company_id, company_licenses.renewable, company_licenses.reminder_counter, company_licenses.issued_by, company_licenses.issued_date, company_licenses.expired_date, company_licenses.earliest_renewal_date, company_licenses.last_renewal_date, company_licenses.status, company_licenses.renewal_status, company_licenses.approved_user_id, company_licenses.renewal_approved_user_id, company_licenses.approved_date, company_licenses.renewal_approved_date, company_licenses.remark, company_licenses.created_user_id, company_licenses.updated_user_id, company_licenses.deleted_user_id, company_licenses.created_at, company_licenses.updated_at, company_licenses.deleted_at").Joins("left join license_types ON company_licenses.license_type_id = license_types.id").Where("company_licenses.company_id=? AND company_licenses.deleted_at = 0", id).Order("company_licenses.id").Find(&companyLicenses)
	return companyLicenses, res.Error
}

func (db *CompanyLicenseConnection) InsertCompanyLicense(companyLicense model.CompanyLicense) (companyLicenseOutput model.CompanyLicense, err error) {
	res := db.connection.Save(&companyLicense)
	return companyLicense, res.Error
}

func (db *CompanyLicenseConnection) UpdateCompanyLicense(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	var (
		company_license model.CompanyLicense
	)
	res := db.connection.Model(&company_license).Where("id=?", id).Updates(map[string]interface{}{"parent_license_id": companyLicense.ParentLicenseID, "license_no": companyLicense.LicenseNo, "license_type_id": companyLicense.LicenseTypeID, "company_id": companyLicense.CompanyID, "renewable": companyLicense.Renewable, "reminder_counter": companyLicense.ReminderCounter, "issued_by": companyLicense.IssuedBy, "issued_date": companyLicense.IssuedDate, "expired_date": companyLicense.ExpiredDate, "earliest_renewal_date": companyLicense.EarliestRenewalDate, "last_renewal_date": companyLicense.LastRenewalDate, "status": companyLicense.Status, "renewal_status": companyLicense.RenewalStatus, "remark": companyLicense.Remark, "updated_user_id": companyLicense.UpdatedUserID, "updated_at": companyLicense.UpdatedAt})
	return companyLicense, res.Error
}

func (db *CompanyLicenseConnection) UpdateCompanyLicenseStatus(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	var (
		company_license model.CompanyLicense
	)
	res := db.connection.Model(&company_license).Where("id=?", id).Updates(map[string]interface{}{"status": companyLicense.Status, "approved_user_id": companyLicense.ApprovedUserID, "approved_date": companyLicense.ApprovedDate})
	return companyLicense, res.Error
}

func (db *CompanyLicenseConnection) UpdateCompanyLicenseDeactive(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	var (
		company_license model.CompanyLicense
	)
	res := db.connection.Model(&company_license).Where("id=?", id).Updates(map[string]interface{}{"status": companyLicense.Status})
	return companyLicense, res.Error
}

func (db *CompanyLicenseConnection) UpdateCompanyLicenseApprovedRenewalStatus(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	var (
		company_license model.CompanyLicense
	)
	res := db.connection.Model(&company_license).Where("id=?", id).Updates(map[string]interface{}{"renewal_status": companyLicense.RenewalStatus, "renewal_approved_user_id": companyLicense.RenewalApprovedUserID, "renewal_approved_date": companyLicense.RenewalApprovedDate})
	return companyLicense, res.Error
}

func (db *CompanyLicenseConnection) DeleteCompanyLicense(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	var (
		company_license model.CompanyLicense
	)
	res := db.connection.Model(&company_license).Where("id=?", id).Updates(map[string]interface{}{"license_no": companyLicense.LicenseNo, "deleted_user_id": companyLicense.DeletedUserID, "deleted_at": companyLicense.DeletedAt})
	return companyLicense, res.Error
}

func (db *CompanyLicenseConnection) UpdateCompanyRemark(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	res := db.connection.Model(&companyLicense).Where("id=?", id).Updates(map[string]interface{}{"remark": companyLicense.Remark})
	return companyLicense, res.Error
}
