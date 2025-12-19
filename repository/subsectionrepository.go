package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type SubSectionRepository interface {
	CountSubSectionAll() (count int64, err error)
	FindSubSections() (subSectionOutput []model.SelectSubSectionParameter, err error)
	FindSubSectionsOffset(limit int, offset int, order string, dir string) (subSectionOutput []model.SelectSubSectionParameter, err error)
	SearchSubSection(limit int, offset int, order string, dir string, search string) (subSectionOutput []model.SelectSubSectionParameter, err error)
	CountSearchSubSection(search string) (count int64, err error)
	FindSubSectionById(id uint) (subSectionOutput model.SelectSubSectionParameter, err error)
	FindExcSubSection(sectionId uint, id uint) (subSectionOutput []model.SelectSubSectionParameter, err error)
	FindSubSectionBySecId(sectionId uint) (subSectionOutput []model.SelectSubSectionParameter, err error)
	FindSubSectionByDepId(depId uint) (subSectionOutput []model.SelectSubSectionParameter, err error)
	FindSubSectionByDivisionID(divId uint) (subSectionOutput []model.SelectSubSectionParameter, err error)
	CountSubSectionName(search string) (count int64, err error)
	InsertSubSection(subSection model.SubSection) (subSectionOutput model.SubSection, err error)
	UpdateSubSection(subSection model.SubSection, id uint) (subSectionOutput model.SubSection, err error)
}

type SubSectionConnection struct {
	connection *gorm.DB
}

func NewSubSectionRepository(db *gorm.DB) SubSectionRepository {
	return &SubSectionConnection{
		connection: db,
	}
}

func (db *SubSectionConnection) CountSubSectionAll() (count int64, err error) {
	res := db.connection.Table("sub_sections").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *SubSectionConnection) FindSubSections() (subSectionOutput []model.SelectSubSectionParameter, err error) {
	var (
		subSections []model.SelectSubSectionParameter
	)
	res := db.connection.Table("sub_sections").Select("sub_sections.id, sub_sections.sub_section_name, sub_sections.division_id, divisions.division_name, sub_sections.department_id, departments.department_name, sub_sections.section_id, sections.section_name, sub_sections.remark, sub_sections.created_user_id, sub_sections.updated_user_id, sub_sections.deleted_user_id, sub_sections.created_at, sub_sections.updated_at, sub_sections.deleted_at").Joins("left join sections ON sub_sections.section_id = sections.id").Joins("left join departments ON sub_sections.department_id = departments.id").Joins("left join divisions ON sub_sections.division_id = divisions.id").Where("sub_sections.deleted_at = 0").Order("divisions.division_id").Find(&subSections)
	return subSections, res.Error
}

func (db *SubSectionConnection) FindSubSectionsOffset(limit int, offset int, order string, dir string) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	var (
		orderDirection string
		subSections    []model.SelectSubSectionParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("sub_sections").Select("sub_sections.id, sub_sections.sub_section_name, sub_sections.division_id, divisions.division_name, sub_sections.department_id, departments.department_name, sub_sections.section_id, sections.section_name, sub_sections.remark, sub_sections.created_user_id, sub_sections.updated_user_id, sub_sections.deleted_user_id, sub_sections.created_at, sub_sections.updated_at, sub_sections.deleted_at").Joins("left join sections ON sub_sections.section_id = sections.id").Joins("left join departments ON sub_sections.department_id = departments.id").Joins("left join divisions ON sub_sections.division_id = divisions.id").Where("sub_sections.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&subSections)
	return subSections, res.Error
}

func (db *SubSectionConnection) SearchSubSection(limit int, offset int, order string, dir string, search string) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	var (
		orderDirection string
		final          string
		subSections    []model.SelectSubSectionParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("sub_sections").Select("sub_sections.id, sub_sections.sub_section_name, sub_sections.division_id, divisions.division_name, sub_sections.department_id, departments.department_name, sub_sections.section_id, sections.section_name, sub_sections.remark, sub_sections.created_user_id, sub_sections.updated_user_id, sub_sections.deleted_user_id, sub_sections.created_at, sub_sections.updated_at, sub_sections.deleted_at").Joins("left join sections ON sub_sections.section_id = sections.id").Joins("left join departments ON sub_sections.department_id = departments.id").Joins("left join divisions ON sub_sections.division_id = divisions.id").Where("(lower(sub_sections.sub_section_name) LIKE ? OR lower(sections.section_name) LIKE ? OR lower(divisions.division_name) LIKE ? OR lower(departments.department_name) LIKE ? OR lower(departments.remark) LIKE ?) AND sub_sections.deleted_at = 0", final, final, final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&subSections)
	return subSections, res.Error
}

func (db *SubSectionConnection) CountSearchSubSection(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("sub_sections").Select("sub_sections.id, sub_sections.sub_section_name, sub_sections.division_id, divisions.division_name, sub_sections.department_id, departments.department_name, sub_sections.section_id, sections.section_name, sub_sections.remark, sub_sections.created_user_id, sub_sections.updated_user_id, sub_sections.deleted_user_id, sub_sections.created_at, sub_sections.updated_at, sub_sections.deleted_at").Joins("left join sections ON sub_sections.section_id = sections.id").Joins("left join departments ON sub_sections.department_id = departments.id").Joins("left join divisions ON sub_sections.division_id = divisions.id").Where("(lower(sub_sections.sub_section_name) LIKE ? OR lower(sections.section_name) LIKE ? OR lower(divisions.division_name) LIKE ? OR lower(departments.department_name) LIKE ? OR lower(departments.remark) LIKE ?) AND sub_sections.deleted_at = 0", final, final, final, final, final).Count(&count)
	return count, res.Error
}

func (db *SubSectionConnection) FindSubSectionById(id uint) (subSectionOutput model.SelectSubSectionParameter, err error) {
	var (
		subSection model.SelectSubSectionParameter
	)

	res := db.connection.Table("sub_sections").Select("sub_sections.id, sub_sections.sub_section_name, sub_sections.division_id, divisions.division_name, sub_sections.department_id, departments.department_name, sub_sections.section_id, sections.section_name, sub_sections.remark, sub_sections.created_user_id, sub_sections.updated_user_id, sub_sections.deleted_user_id, sub_sections.created_at, sub_sections.updated_at, sub_sections.deleted_at").Joins("left join sections ON sub_sections.section_id = sections.id").Joins("left join departments ON sub_sections.department_id = departments.id").Joins("left join divisions ON sub_sections.division_id = divisions.id").Where("sub_sections.id=? AND sub_sections.deleted_at = 0", id).Take(&subSection)
	return subSection, res.Error
}

func (db *SubSectionConnection) FindExcSubSection(sectionId uint, id uint) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	var (
		subSections []model.SelectSubSectionParameter
	)

	res := db.connection.Table("sub_sections").Select("sub_sections.id, sub_sections.sub_section_name, sub_sections.division_id, divisions.division_name, sub_sections.department_id, departments.department_name, sub_sections.section_id, sections.section_name, sub_sections.remark, sub_sections.created_user_id, sub_sections.updated_user_id, sub_sections.deleted_user_id, sub_sections.created_at, sub_sections.updated_at, sub_sections.deleted_at").Joins("left join sections ON sub_sections.section_id = sections.id").Joins("left join departments ON sub_sections.department_id = departments.id").Joins("left join divisions ON sub_sections.division_id = divisions.id").Where("sub_sections.section_id=? AND sub_sections.id != ? AND sub_sections.deleted_at = 0", sectionId, id).Find(&subSections)
	return subSections, res.Error
}

func (db *SubSectionConnection) FindSubSectionBySecId(sectionId uint) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	var (
		subSections []model.SelectSubSectionParameter
	)

	res := db.connection.Table("sub_sections").Select("sub_sections.id, sub_sections.sub_section_name, sub_sections.division_id, divisions.division_name, sub_sections.department_id, departments.department_name, sub_sections.section_id, sections.section_name, sub_sections.remark, sub_sections.created_user_id, sub_sections.updated_user_id, sub_sections.deleted_user_id, sub_sections.created_at, sub_sections.updated_at, sub_sections.deleted_at").Joins("left join sections ON sub_sections.section_id = sections.id").Joins("left join departments ON sub_sections.department_id = departments.id").Joins("left join divisions ON sub_sections.division_id = divisions.id").Where("sub_sections.section_id=? AND sub_sections.deleted_at = 0", sectionId).Order("sub_sections.sub_section_name").Find(&subSections)
	return subSections, res.Error
}

func (db *SubSectionConnection) FindSubSectionByDepId(depId uint) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	var (
		subSections []model.SelectSubSectionParameter
	)

	res := db.connection.Table("sub_sections").Select("sub_sections.id, sub_sections.sub_section_name, sub_sections.division_id, divisions.division_name, sub_sections.department_id, departments.department_name, sub_sections.section_id, sections.section_name, sub_sections.remark, sub_sections.created_user_id, sub_sections.updated_user_id, sub_sections.deleted_user_id, sub_sections.created_at, sub_sections.updated_at, sub_sections.deleted_at").Joins("left join sections ON sub_sections.section_id = sections.id").Joins("left join departments ON sub_sections.department_id = departments.id").Joins("left join divisions ON sub_sections.division_id = divisions.id").Where("sub_sections.department_id=? AND sub_sections.deleted_at = 0", depId).Order("sub_sections.sub_section_name").Find(&subSections)
	return subSections, res.Error
}

func (db *SubSectionConnection) FindSubSectionByDivisionID(divId uint) (subSectionOutput []model.SelectSubSectionParameter, err error) {
	var (
		subSections []model.SelectSubSectionParameter
	)

	res := db.connection.Table("sub_sections").Select("sub_sections.id, sub_sections.sub_section_name, sub_sections.division_id, divisions.division_name, sub_sections.department_id, departments.department_name, sub_sections.section_id, sections.section_name, sub_sections.remark, sub_sections.created_user_id, sub_sections.updated_user_id, sub_sections.deleted_user_id, sub_sections.created_at, sub_sections.updated_at, sub_sections.deleted_at").Joins("left join sections ON sub_sections.section_id = sections.id").Joins("left join departments ON sub_sections.department_id = departments.id").Joins("left join divisions ON sub_sections.division_id = divisions.id").Where("sub_sections.division_id=? AND sub_sections.deleted_at = 0", divId).Find(&subSections)
	return subSections, res.Error
}

func (db *SubSectionConnection) CountSubSectionName(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("sub_sections").Select("sub_sections.id, sub_sections.sub_section_name, sub_sections.division_id, divisions.division_name, sub_sections.department_id, departments.department_name, sub_sections.section_id, sections.section_name, sub_sections.remark, sub_sections.created_user_id, sub_sections.updated_user_id, sub_sections.deleted_user_id, sub_sections.created_at, sub_sections.updated_at, sub_sections.deleted_at").Joins("left join sections ON sub_sections.section_id = sections.id").Joins("left join departments ON sub_sections.department_id = departments.id").Joins("left join divisions ON sub_sections.division_id = divisions.id").Where("lower(sub_sections.sub_section_name) LIKE ?", final).Count(&count)
	return count, res.Error
}

func (db *SubSectionConnection) InsertSubSection(subSection model.SubSection) (subSectionOutput model.SubSection, err error) {
	res := db.connection.Save(&subSection)
	return subSection, res.Error
}

func (db *SubSectionConnection) UpdateSubSection(subSection model.SubSection, id uint) (subSectionOutput model.SubSection, err error) {
	res := db.connection.Where("id=?", id).Updates(&subSection)
	return subSection, res.Error
}
