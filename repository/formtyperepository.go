package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type FormTypeRepository interface {
	CountFormTypeAll() (count int64, err error)
	FindFormTypes() (formTypeOutput []model.FormType, err error)
	FindFormTypesOffset(limit int, offset int, order string, dir string) (formTypeOutput []model.FormType, err error)
	SearchFormType(limit int, offset int, order string, dir string, search string) (formTypeOutput []model.FormType, err error)
	CountSearchFormType(search string) (count int64, err error)
	FindFormTypeById(id uint) (formTypeOutput model.FormType, err error)
	FindExcFormType(id uint) (formTypeOutput []model.FormType, err error)
	InsertFormType(formType model.FormType) (formTypeOutput model.FormType, err error)
	UpdateFormType(formType model.FormType, id uint) (formTypeOutput model.FormType, err error)
}

type FormTypeConnection struct {
	connection *gorm.DB
}

func NewFormTypeRepository(db *gorm.DB) FormTypeRepository {
	return &FormTypeConnection{
		connection: db,
	}
}

func (db *FormTypeConnection) CountFormTypeAll() (count int64, err error) {
	res := db.connection.Debug().Table("form_types").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *FormTypeConnection) FindFormTypes() (formTypeOutput []model.FormType, err error) {
	var (
		formTypes []model.FormType
	)
	res := db.connection.Where("deleted_at = 0").Order("form_type_code").Find(&formTypes)
	return formTypes, res.Error
}

func (db *FormTypeConnection) FindFormTypesOffset(limit int, offset int, order string, dir string) (formTypeOutput []model.FormType, err error) {
	var (
		orderDirection string
		formTypes      []model.FormType
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&formTypes)
	return formTypes, res.Error
}

func (db *FormTypeConnection) SearchFormType(limit int, offset int, order string, dir string, search string) (formTypeOutput []model.FormType, err error) {
	var (
		orderDirection string
		final          string
		formTypes      []model.FormType
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(form_type_code) LIKE ? OR lower(form_type_description) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&formTypes)
	return formTypes, res.Error
}

func (db *FormTypeConnection) CountSearchFormType(search string) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("form_types").Where("(lower(form_type_code) LIKE ? OR lower(form_type_description) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final).Count(&count)
	return count, res.Error
}

func (db *FormTypeConnection) FindFormTypeById(id uint) (formTypeOutput model.FormType, err error) {
	var (
		formType model.FormType
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&formType)
	return formType, res.Error
}

func (db *FormTypeConnection) FindExcFormType(id uint) (formTypeOutput []model.FormType, err error) {
	var (
		formTypes []model.FormType
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("form_type_code").Find(&formTypes)
	return formTypes, res.Error
}

func (db *FormTypeConnection) InsertFormType(formType model.FormType) (formTypeOutput model.FormType, err error) {
	res := db.connection.Save(&formType)
	return formType, res.Error
}

func (db *FormTypeConnection) UpdateFormType(formType model.FormType, id uint) (formTypeOutput model.FormType, err error) {
	res := db.connection.Where("id=?", id).Updates(&formType)
	return formType, res.Error
}
