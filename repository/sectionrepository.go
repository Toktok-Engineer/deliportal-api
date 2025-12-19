package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type SectionRepository interface {
	CountSectionAll() (count int64, err error)
	FindSections() (sectionOutput []model.SelectSectionParameter, err error)
	FindSectionsOffset(limit int, offset int, order string, dir string) (sectionOutput []model.SelectSectionParameter, err error)
	SearchSection(limit int, offset int, order string, dir string, search string) (sectionOutput []model.SelectSectionParameter, err error)
	CountSearchSection(search string) (count int64, err error)
	FindSectionById(id uint) (sectionOutput model.SelectSectionParameter, err error)
	FindExcSection(depId uint, id uint) (sectionOutput []model.SelectSectionParameter, err error)
	FindSectionByDepId(depId uint) (sectionOutput []model.SelectSectionParameter, err error)
	FindSectionByDivisionID(divId uint) (sectionOutput []model.SelectSectionParameter, err error)
	CountSectionName(search string) (count int64, err error)
	InsertSection(section model.Section) (sectionOutput model.Section, err error)
	UpdateSection(section model.Section, id uint) (sectionOutput model.Section, err error)
}

type SectionConnection struct {
	connection *gorm.DB
}

func NewSectionRepository(db *gorm.DB) SectionRepository {
	return &SectionConnection{
		connection: db,
	}
}

func (db *SectionConnection) CountSectionAll() (count int64, err error) {
	res := db.connection.Table("sections").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *SectionConnection) FindSections() (sectionOutput []model.SelectSectionParameter, err error) {
	var (
		sections []model.SelectSectionParameter
	)
	res := db.connection.Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("sections.deleted_at = 0").Order("departments.department_name").Find(&sections)
	return sections, res.Error
}

func (db *SectionConnection) FindSectionsOffset(limit int, offset int, order string, dir string) (sectionOutput []model.SelectSectionParameter, err error) {
	var (
		orderDirection string
		sections       []model.SelectSectionParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("sections.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&sections)
	return sections, res.Error
}

func (db *SectionConnection) SearchSection(limit int, offset int, order string, dir string, search string) (sectionOutput []model.SelectSectionParameter, err error) {
	var (
		orderDirection string
		final          string
		sections       []model.SelectSectionParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("(lower(sections.section_name) LIKE ? OR lower(divisions.division_name) LIKE ? OR lower(departments.department_name) LIKE ? OR lower(departments.remark) LIKE ?) AND sections.deleted_at = 0", final, final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&sections)
	return sections, res.Error
}

func (db *SectionConnection) CountSearchSection(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("(lower(sections.section_name) LIKE ? OR lower(divisions.division_name) LIKE ? OR lower(departments.department_name) LIKE ? OR lower(departments.remark) LIKE ?) AND sections.deleted_at = 0", final, final, final, final).Count(&count)
	return count, res.Error
}

func (db *SectionConnection) FindSectionById(id uint) (sectionOutput model.SelectSectionParameter, err error) {
	var (
		section model.SelectSectionParameter
	)

	res := db.connection.Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("sections.id=? AND sections.deleted_at = 0", id).Take(&section)
	return section, res.Error
}

func (db *SectionConnection) FindExcSection(depId uint, id uint) (sectionOutput []model.SelectSectionParameter, err error) {
	var (
		sections []model.SelectSectionParameter
	)

	res := db.connection.Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("sections.department_id=? AND sections.id != ? AND sections.deleted_at = 0", depId, id).Find(&sections)
	return sections, res.Error
}

func (db *SectionConnection) FindSectionByDepId(depId uint) (sectionOutput []model.SelectSectionParameter, err error) {
	var (
		sections []model.SelectSectionParameter
	)

	res := db.connection.Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("sections.department_id=? AND sections.deleted_at = 0", depId).Order("sections.section_name").Find(&sections)
	return sections, res.Error
}

func (db *SectionConnection) FindSectionByDivisionID(divId uint) (sectionOutput []model.SelectSectionParameter, err error) {
	var (
		sections []model.SelectSectionParameter
	)

	res := db.connection.Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("sections.division_id=? AND sections.deleted_at = 0", divId).Find(&sections)
	return sections, res.Error
}

func (db *SectionConnection) CountSectionName(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("lower(sections.section_name) LIKE ?", final).Count(&count)
	return count, res.Error
}

func (db *SectionConnection) InsertSection(section model.Section) (sectionOutput model.Section, err error) {
	res := db.connection.Save(&section)
	return section, res.Error
}

func (db *SectionConnection) UpdateSection(section model.Section, id uint) (sectionOutput model.Section, err error) {
	res := db.connection.Where("id=?", id).Updates(&section)
	return section, res.Error
}
