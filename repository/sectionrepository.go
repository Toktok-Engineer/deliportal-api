package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type SectionRepository interface {
	FindSections() (sectionOutput []model.SelectSectionParameter, err error)
	FindSectionById(id uint) (sectionOutput model.SelectSectionParameter, err error)
	FindExcSection(depId uint, id uint) (sectionOutput []model.SelectSectionParameter, err error)
	FindSectionByDepId(depId uint) (sectionOutput []model.SelectSectionParameter, err error)
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

func (db *SectionConnection) FindSections() (sectionOutput []model.SelectSectionParameter, err error) {
	var (
		sections []model.SelectSectionParameter
	)

	res := db.connection.Debug().Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("sections.deleted_at = 0").Order("sections.section_name").Find(&sections)
	return sections, res.Error
}

func (db *SectionConnection) FindSectionById(id uint) (sectionOutput model.SelectSectionParameter, err error) {
	var (
		section model.SelectSectionParameter
	)

	res := db.connection.Debug().Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("sections.id=? AND sections.deleted_at = 0", id).Take(&section)
	return section, res.Error
}

func (db *SectionConnection) FindExcSection(depId uint, id uint) (sectionOutput []model.SelectSectionParameter, err error) {
	var (
		sections []model.SelectSectionParameter
	)

	res := db.connection.Debug().Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("sections.department_id=? AND sections.id != ? AND sections.deleted_at = 0", depId, id).Find(&sections)
	return sections, res.Error
}

func (db *SectionConnection) FindSectionByDepId(depId uint) (sectionOutput []model.SelectSectionParameter, err error) {
	var (
		sections []model.SelectSectionParameter
	)

	res := db.connection.Debug().Table("sections").Select("sections.id, sections.section_name, sections.division_id, divisions.division_name, sections.department_id, departments.department_name, sections.remark, sections.created_user_id, sections.updated_user_id, sections.deleted_user_id, sections.created_at, sections.updated_at, sections.deleted_at").Joins("left join departments ON sections.department_id = departments.id").Joins("left join divisions ON sections.division_id = divisions.id").Where("sections.department_id=? AND sections.deleted_at = 0", depId).Order("sections.section_name").Find(&sections)
	return sections, res.Error
}

func (db *SectionConnection) InsertSection(section model.Section) (sectionOutput model.Section, err error) {
	res := db.connection.Save(&section)
	return section, res.Error
}

func (db *SectionConnection) UpdateSection(section model.Section, id uint) (sectionOutput model.Section, err error) {
	res := db.connection.Where("id=?", id).Updates(&section)
	return section, res.Error
}
