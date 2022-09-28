package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type ImDefaultCirculationRepository interface {
	CountImDefaultCirculationAll(companyGroup int) (count int64, err error)
	FindImDefaultCirculations(companyGroupID int) (imdefaultcirculationOutput []model.ImDefaultCirculation, err error)
	FindImDefaultCirculationsOffset(limit int, offset int, order string, dir string, companyGroup int) (imdefaultcirculationOutput []model.SelectImDefaultCirculationParameter, err error)
	SearchImDefaultCirculation(limit int, offset int, order string, dir string, search string, companyGroup int) (imdefaultcirculationOutput []model.SelectImDefaultCirculationParameter, err error)
	CountSearchImDefaultCirculation(search string, companyGroup int) (count int64, err error)
	FindImDefaultCirculationById(id uint) (imDefaultCirculationOutput model.ImDefaultCirculation, err error)
	FindExcImDefaultCirculation(id uint) (imDefaultCirculationOutput []model.ImDefaultCirculation, err error)
	InsertImDefaultCirculation(imDefaultCirculation model.ImDefaultCirculation) (imDefaultCirculationOutput model.ImDefaultCirculation, err error)
	UpdateImDefaultCirculation(imDefaultCirculation model.ImDefaultCirculation, id uint) (imDefaultCirculationOutput model.ImDefaultCirculation, err error)
}

type imDefaultCirculationConnection struct {
	connection *gorm.DB
}

func NewImDefaultCirculationRepository(db *gorm.DB) ImDefaultCirculationRepository {
	return &imDefaultCirculationConnection{
		connection: db,
	}
}
func (db *imDefaultCirculationConnection) CountImDefaultCirculationAll(companyGroup int) (count int64, err error) {
	res := db.connection.Debug().Table("im_default_circulations").Where("company_group_id = ? AND deleted_at = 0", companyGroup).Count(&count)
	return count, res.Error
}
func (db *imDefaultCirculationConnection) FindImDefaultCirculations(companyGroupID int) (imdefaultcirculationOutput []model.ImDefaultCirculation, err error) {
	var (
		imdefaultcirculations []model.ImDefaultCirculation
	)
	res := db.connection.Where("company_group_id = ? AND deleted_at = 0", companyGroupID).Order("sequence_no").Find(&imdefaultcirculations)
	return imdefaultcirculations, res.Error
}

func (db *imDefaultCirculationConnection) FindImDefaultCirculationsOffset(limit int, offset int, order string, dir string, companyGroup int) (imdefaultcirculationOutput []model.SelectImDefaultCirculationParameter, err error) {
	var (
		orderDirection        string
		imdefaultcirculations []model.SelectImDefaultCirculationParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("im_default_circulations").Select("im_default_circulations.id, im_default_circulations.sequence_no, im_default_circulations.company_group_id, im_default_circulations.employee_id, employees.firstname, employees.lastname, im_default_circulations.remark, im_default_circulations.created_user_id, createdUID.username AS created_user, im_default_circulations.updated_user_id, updatedUID.username AS updated_user, im_default_circulations.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(im_default_circulations.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(im_default_circulations.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(im_default_circulations.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join employees on im_default_circulations.employee_id = employees.id").Joins("left join users createdUID on im_default_circulations.created_user_id = createdUID.id").Joins("left join users updatedUID on im_default_circulations.updated_user_id = updatedUID.id").Joins("left join users deletedUID on im_default_circulations.deleted_user_id = deletedUID.id").Where("im_default_circulations.company_group_id = ? AND im_default_circulations.deleted_at = 0", companyGroup).Order(orderDirection).Limit(limit).Offset(offset).Find(&imdefaultcirculations)
	return imdefaultcirculations, res.Error
}

func (db *imDefaultCirculationConnection) SearchImDefaultCirculation(limit int, offset int, order string, dir string, search string, companyGroup int) (imdefaultcirculationOutput []model.SelectImDefaultCirculationParameter, err error) {
	var (
		orderDirection        string
		final                 string
		imdefaultcirculations []model.SelectImDefaultCirculationParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("im_default_circulations").Select("im_default_circulations.id, im_default_circulations.sequence_no, im_default_circulations.company_group_id, im_default_circulations.employee_id, employees.firstname, employees.lastname, im_default_circulations.remark, im_default_circulations.created_user_id, createdUID.username AS created_user, im_default_circulations.updated_user_id, updatedUID.username AS updated_user, im_default_circulations.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(im_default_circulations.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(im_default_circulations.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(im_default_circulations.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join employees on im_default_circulations.employee_id = employees.id").Joins("left join users createdUID on im_default_circulations.created_user_id = createdUID.id").Joins("left join users updatedUID on im_default_circulations.updated_user_id = updatedUID.id").Joins("left join users deletedUID on im_default_circulations.deleted_user_id = deletedUID.id").Where("(lower(im_default_circulations.last_circulation_sequence_no::text) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(im_default_circulations.remark) LIKE ?) AND im_default_circulations.company_group_id = ? AND deleted_at = 0", final, final, final, final, companyGroup).Order(orderDirection).Limit(limit).Offset(offset).Find(&imdefaultcirculations)
	return imdefaultcirculations, res.Error
}

func (db *imDefaultCirculationConnection) CountSearchImDefaultCirculation(search string, companyGroup int) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("im_default_circulations").Select("im_default_circulations.id, im_default_circulations.sequence_no, im_default_circulations.company_group_id, im_default_circulations.employee_id, employees.firstname, employees.lastname, im_default_circulations.remark, im_default_circulations.created_user_id, createdUID.username AS created_user, im_default_circulations.updated_user_id, updatedUID.username AS updated_user, im_default_circulations.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(im_default_circulations.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(im_default_circulations.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(im_default_circulations.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at").Joins("left join employees on im_default_circulations.employee_id = employees.id").Joins("left join users createdUID on im_default_circulations.created_user_id = createdUID.id").Joins("left join users updatedUID on im_default_circulations.updated_user_id = updatedUID.id").Joins("left join users deletedUID on im_default_circulations.deleted_user_id = deletedUID.id").Where("(lower(im_default_circulations.last_circulation_sequence_no::text) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(im_default_circulations.remark) LIKE ?) AND im_default_circulations.company_group_id = ? AND deleted_at = 0", final, final, final, final, companyGroup).Count(&count)
	return count, res.Error
}

func (db *imDefaultCirculationConnection) FindImDefaultCirculationById(id uint) (imDefaultCirculationOutput model.ImDefaultCirculation, err error) {
	var (
		imDefaultCirculation model.ImDefaultCirculation
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&imDefaultCirculation)
	return imDefaultCirculation, res.Error
}

func (db *imDefaultCirculationConnection) FindExcImDefaultCirculation(id uint) (imDefaultCirculationOutput []model.ImDefaultCirculation, err error) {
	var (
		imDefaultCirculations []model.ImDefaultCirculation
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("sequence_no").Find(&imDefaultCirculations)
	return imDefaultCirculations, res.Error
}

func (db *imDefaultCirculationConnection) InsertImDefaultCirculation(imDefaultCirculation model.ImDefaultCirculation) (imDefaultCirculationOutput model.ImDefaultCirculation, err error) {
	res := db.connection.Save(&imDefaultCirculation)
	return imDefaultCirculation, res.Error
}

func (db *imDefaultCirculationConnection) UpdateImDefaultCirculation(imDefaultCirculation model.ImDefaultCirculation, id uint) (imDefaultCirculationOutput model.ImDefaultCirculation, err error) {
	res := db.connection.Where("id=?", id).Updates(&imDefaultCirculation)
	return imDefaultCirculation, res.Error
}
