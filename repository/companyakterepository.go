package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type CompanyAkteRepository interface {
	CountCompanyAkteAll(companyId int) (count int64, err error)
	FindCompanyAktes(companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	FindCompanyAktesOffset(limit int, offset int, order string, dir string, companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	SearchCompanyAkte(limit int, offset int, order string, dir string, search string, companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	CountSearchCompanyAkte(search string, companyId int) (count int64, err error)
	FindCompanyAkteById(id uint) (companyAkteOutput model.SelectCompanyAkteParameter, err error)
	FindCompanyAkteByYear(companyId uint, year uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	FindExcCompanyAkteByYear(companyId uint, year uint, companyAkteId uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	FindExcCompanyAkte(id uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	InsertCompanyAkte(companyAkte model.CompanyAkte) (companyAkteOutput model.CompanyAkte, err error)
	UpdateCompanyAkte(companyAkte model.CompanyAkte, id uint) (companyAkteOutput model.CompanyAkte, err error)
}

type CompanyAkteConnection struct {
	connection *gorm.DB
}

func NewCompanyAkteRepository(db *gorm.DB) CompanyAkteRepository {
	return &CompanyAkteConnection{
		connection: db,
	}
}

func (db *CompanyAkteConnection) CountCompanyAkteAll(companyId int) (count int64, err error) {
	res := db.connection.Debug().Table("company_aktes").Where("company_id = ? AND deleted_at = 0 ", companyId).Count(&count)
	return count, res.Error
}

func (db *CompanyAkteConnection) FindCompanyAktes(companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	var (
		companyAktes []model.SelectCompanyAkteParameter
	)
	res := db.connection.Debug().Table("company_aktes").Select("company_aktes.id, company_aktes.company_id, companies.company_name, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_aktes.skno1, company_aktes.skno2, company_aktes.skno3, company_aktes.file_name, company_aktes.file_url, company_aktes.remark, company_aktes.created_user_id, createdUID.username AS created_user, company_aktes.updated_user_id, updatedUID.username AS updated_user, company_aktes.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_aktes.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_aktes.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_aktes.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies ON company_aktes.company_id = companies.id").Joins("left join users createdUID on company_aktes.created_user_id = createdUID.id").Joins("left join users updatedUID on company_aktes.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_aktes.deleted_user_id = deletedUID.id").Where("company_aktes.company_id = ? AND company_aktes.deleted_at = 0", companyId).Order("company_aktes.year").Find(&companyAktes)
	return companyAktes, res.Error
}

func (db *CompanyAkteConnection) FindCompanyAktesOffset(limit int, offset int, order string, dir string, companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	var (
		orderDirection string
		companyAktes   []model.SelectCompanyAkteParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("company_aktes").Select("company_aktes.id, company_aktes.company_id, companies.company_name, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_aktes.skno1, company_aktes.skno2, company_aktes.skno3, company_aktes.file_name, company_aktes.file_url, company_aktes.remark, company_aktes.created_user_id, createdUID.username AS created_user, company_aktes.updated_user_id, updatedUID.username AS updated_user, company_aktes.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_aktes.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_aktes.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_aktes.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies ON company_aktes.company_id = companies.id").Joins("left join users createdUID on company_aktes.created_user_id = createdUID.id").Joins("left join users updatedUID on company_aktes.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_aktes.deleted_user_id = deletedUID.id").Where("company_aktes.company_id = ? AND company_aktes.deleted_at = 0", companyId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyAktes)
	return companyAktes, res.Error
}

func (db *CompanyAkteConnection) SearchCompanyAkte(limit int, offset int, order string, dir string, search string, companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	var (
		orderDirection string
		final          string
		companyAktes   []model.SelectCompanyAkteParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_aktes").Select("company_aktes.id, company_aktes.company_id, companies.company_name, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_aktes.skno1, company_aktes.skno2, company_aktes.skno3, company_aktes.file_name, company_aktes.file_url, company_aktes.remark, company_aktes.created_user_id, createdUID.username AS created_user, company_aktes.updated_user_id, updatedUID.username AS updated_user, company_aktes.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_aktes.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_aktes.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_aktes.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies ON company_aktes.company_id = companies.id").Joins("left join users createdUID on company_aktes.created_user_id = createdUID.id").Joins("left join users updatedUID on company_aktes.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_aktes.deleted_user_id = deletedUID.id").Where("company_aktes.year::text LIKE ? OR lower(to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(company_aktes.akte_no) LIKE ? OR lower(company_aktes.remark) LIKE ? AND company_aktes.company_id = ? AND company_aktes.deleted_at = 0", final, final, final, final, companyId).Order(orderDirection).Limit(limit).Offset(offset).Find(&companyAktes)
	return companyAktes, res.Error
}

func (db *CompanyAkteConnection) CountSearchCompanyAkte(search string, companyId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("company_aktes").Select("company_aktes.id, company_aktes.company_id, companies.company_name, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_aktes.skno1, company_aktes.skno2, company_aktes.skno3, company_aktes.file_name, company_aktes.file_url, company_aktes.remark, company_aktes.created_user_id, createdUID.username AS created_user, company_aktes.updated_user_id, updatedUID.username AS updated_user, company_aktes.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_aktes.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_aktes.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_aktes.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies ON company_aktes.company_id = companies.id").Joins("left join users createdUID on company_aktes.created_user_id = createdUID.id").Joins("left join users updatedUID on company_aktes.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_aktes.deleted_user_id = deletedUID.id").Where("company_aktes.year::text LIKE ? OR lower(to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(company_aktes.akte_no) LIKE ? OR lower(company_aktes.remark) LIKE ? AND company_aktes.company_id = ? AND company_aktes.deleted_at = 0", final, final, final, final, companyId).Count(&count)
	return count, res.Error
}

func (db *CompanyAkteConnection) FindCompanyAkteById(id uint) (companyAkteOutput model.SelectCompanyAkteParameter, err error) {
	var (
		companyAkte model.SelectCompanyAkteParameter
	)
	res := db.connection.Debug().Table("company_aktes").Select("company_aktes.id, company_aktes.company_id, companies.company_name, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_aktes.skno1, company_aktes.skno2, company_aktes.skno3, company_aktes.file_name, company_aktes.file_url, company_aktes.remark, company_aktes.created_user_id, createdUID.username AS created_user, company_aktes.updated_user_id, updatedUID.username AS updated_user, company_aktes.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_aktes.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_aktes.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_aktes.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies ON company_aktes.company_id = companies.id").Joins("left join users createdUID on company_aktes.created_user_id = createdUID.id").Joins("left join users updatedUID on company_aktes.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_aktes.deleted_user_id = deletedUID.id").Where("company_aktes.id=? AND company_aktes.deleted_at = 0", id).Take(&companyAkte)
	return companyAkte, res.Error
}

func (db *CompanyAkteConnection) FindCompanyAkteByYear(companyId uint, year uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	var (
		companyAktes []model.SelectCompanyAkteParameter
	)
	res := db.connection.Debug().Table("company_aktes").Select("company_aktes.id, company_aktes.company_id, companies.company_name, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_aktes.skno1, company_aktes.skno2, company_aktes.skno3, company_aktes.file_name, company_aktes.file_url, company_aktes.remark, company_aktes.created_user_id, createdUID.username AS created_user, company_aktes.updated_user_id, updatedUID.username AS updated_user, company_aktes.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_aktes.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_aktes.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_aktes.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies ON company_aktes.company_id = companies.id").Joins("left join users createdUID on company_aktes.created_user_id = createdUID.id").Joins("left join users updatedUID on company_aktes.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_aktes.deleted_user_id = deletedUID.id").Where("company_aktes.company_id = ? AND company_aktes.year=? AND company_aktes.deleted_at = 0", companyId, year).Find(&companyAktes)
	return companyAktes, res.Error
}

func (db *CompanyAkteConnection) FindExcCompanyAkteByYear(companyId uint, year uint, companyAkteId uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	var (
		companyAktes []model.SelectCompanyAkteParameter
	)
	res := db.connection.Debug().Table("company_aktes").Select("company_aktes.id, company_aktes.company_id, companies.company_name, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_aktes.skno1, company_aktes.skno2, company_aktes.skno3, company_aktes.file_name, company_aktes.file_url, company_aktes.remark, company_aktes.created_user_id, createdUID.username AS created_user, company_aktes.updated_user_id, updatedUID.username AS updated_user, company_aktes.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_aktes.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_aktes.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_aktes.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies ON company_aktes.company_id = companies.id").Joins("left join users createdUID on company_aktes.created_user_id = createdUID.id").Joins("left join users updatedUID on company_aktes.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_aktes.deleted_user_id = deletedUID.id").Where("company_aktes.company_id = ? AND company_aktes.year=? AND company_aktes.id != ? AND company_aktes.deleted_at = 0", companyId, year, companyAkteId).Find(&companyAktes)
	return companyAktes, res.Error
}

func (db *CompanyAkteConnection) FindExcCompanyAkte(id uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	var (
		companyAktes []model.SelectCompanyAkteParameter
	)
	res := db.connection.Debug().Table("company_aktes").Select("company_aktes.id, company_aktes.company_id, companies.company_name, company_aktes.akte_no, to_char(to_timestamp(company_aktes.akte_date::numeric), 'DD-Mon-YYYY') as akte_date, company_aktes.year, company_aktes.skno1, company_aktes.skno2, company_aktes.skno3, company_aktes.file_name, company_aktes.file_url, company_aktes.remark, company_aktes.created_user_id, createdUID.username AS created_user, company_aktes.updated_user_id, updatedUID.username AS updated_user, company_aktes.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(company_aktes.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(company_aktes.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(company_aktes.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join companies ON company_aktes.company_id = companies.id").Joins("left join users createdUID on company_aktes.created_user_id = createdUID.id").Joins("left join users updatedUID on company_aktes.updated_user_id = updatedUID.id").Joins("left join users deletedUID on company_aktes.deleted_user_id = deletedUID.id").Where("company_aktes.id!=? AND company_aktes.deleted_at = 0", id).Order("company_aktes.year").Find(&companyAktes)
	return companyAktes, res.Error
}

func (db *CompanyAkteConnection) InsertCompanyAkte(companyAkte model.CompanyAkte) (companyAkteOutput model.CompanyAkte, err error) {
	res := db.connection.Save(&companyAkte)
	return companyAkte, res.Error
}

func (db *CompanyAkteConnection) UpdateCompanyAkte(companyAkte model.CompanyAkte, id uint) (companyAkteOutput model.CompanyAkte, err error) {
	res := db.connection.Where("id=?", id).Updates(&companyAkte)
	return companyAkte, res.Error
}
