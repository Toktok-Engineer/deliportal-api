package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type FormTypeRepository interface {
	FindFormTypes() (formTypeOutput []model.FormType, err error)
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

func (db *FormTypeConnection) FindFormTypes() (formTypeOutput []model.FormType, err error) {
	var (
		formTypes []model.FormType
	)
	res := db.connection.Where("deleted_at = 0").Order("form_type_code").Find(&formTypes)
	return formTypes, res.Error
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
