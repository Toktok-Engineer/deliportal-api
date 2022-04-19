package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type FormRepository interface {
	FindForms() (formOutput []model.SelectFormParameter, err error)
	FindFormJoinRole(uId uint, fpId uint) (formOutput []model.SelectFormCRUDParameter, err error)
	FindFormByRole(uId uint) (formOutput []model.SelectFormCRUDParameter, err error)
	FindFormByType(tyId uint) (formOutput []model.SelectFormParameter, err error)
	FindExcFormByType(tyId uint, id uint) (formOutput []model.SelectFormParameter, err error)
	FindFormById(id uint) (formOutput model.SelectFormParameter, err error)
	FindFormByFormTypeId(ftId uint) (formOutput []model.SelectFormParameter, err error)
	FindExcForm(ftId uint, id uint) (formOutput []model.SelectFormParameter, err error)
	FindFormHead(ftId uint) (formOutput []model.SelectFormParameter, err error)
	FindFormHeadDetail(id uint) (formOutput model.SelectFormParameter, err error)
	FindExcFormHead(id uint) (formOutput []model.SelectFormParameter, err error)
	FindExcFormOnly(id uint) (formOutput []model.SelectFormParameter, err error)
	InsertForm(form model.Form) (formOutput model.Form, err error)
	UpdateForm(form model.Form, id uint) (formOutput model.Form, err error)
}

type FormConnection struct {
	connection *gorm.DB
}

func NewFormRepository(db *gorm.DB) FormRepository {
	return &FormConnection{
		connection: db,
	}
}

func (db *FormConnection) FindForms() (formOutput []model.SelectFormParameter, err error) {
	var (
		forms []model.SelectFormParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at").Joins("left join form_types ON forms.form_type_id = form_types.id").Where("forms.deleted_at = 0").Order("forms.id").Find(&forms)
	return forms, res.Error
}

func (db *FormConnection) FindFormJoinRole(uId uint, fpId uint) (formOutput []model.SelectFormCRUDParameter, err error) {
	var (
		forms []model.SelectFormCRUDParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at, bool_or(role_forms.create_flag) AS \"create_flag\", bool_or(role_forms.read_flag) AS \"read_flag\", bool_or(role_forms.update_flag) AS \"update_flag\", bool_or(role_forms.delete_flag) AS \"delete_flag\"").Joins("left join form_types ON forms.form_type_id = form_types.id").Joins("left join role_forms ON forms.id = role_forms.form_id").Joins("left join user_roles ON role_forms.role_id = user_roles.role_id").Where("user_roles.user_id = ? AND forms.form_parent_id = ? AND role_forms.deleted_at = 0", uId, fpId).Group("forms.id, form_types.Form_type_code, form_types.form_type_description").Order("forms.sequence_no").Find(&forms)
	return forms, res.Error
}

func (db *FormConnection) FindFormByRole(uId uint) (formOutput []model.SelectFormCRUDParameter, err error) {
	var (
		forms []model.SelectFormCRUDParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at, bool_or(role_forms.create_flag) AS \"create_flag\", bool_or(role_forms.read_flag) AS \"read_flag\", bool_or(role_forms.update_flag) AS \"update_flag\", bool_or(role_forms.delete_flag) AS \"delete_flag\"").Joins("left join form_types ON forms.form_type_id = form_types.id").Joins("left join role_forms ON forms.id = role_forms.form_id").Joins("left join user_roles ON role_forms.role_id = user_roles.role_id").Where("user_roles.user_id = ? AND role_forms.deleted_at = 0", uId).Group("forms.id, form_types.Form_type_code, form_types.form_type_description").Order("forms.sequence_no").Find(&forms)
	return forms, res.Error
}

func (db *FormConnection) FindFormByType(tyId uint) (formOutput []model.SelectFormParameter, err error) {
	var (
		forms []model.SelectFormParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at").Joins("left join form_types ON forms.form_type_id = form_types.id").Where("forms.form_type_id BETWEEN ? AND ? AND forms.deleted_at = 0", tyId, tyId).Order("forms.sequence_no").Find(&forms)
	return forms, res.Error
}

func (db *FormConnection) FindExcFormByType(tyId uint, id uint) (formOutput []model.SelectFormParameter, err error) {
	var (
		forms []model.SelectFormParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at").Joins("left join form_types ON forms.form_type_id = form_types.id").Where("forms.form_type_id BETWEEN ? AND ? AND forms.id!=? AND forms.deleted_at = 0", tyId, tyId, id).Order("forms.sequence_no").Find(&forms)
	return forms, res.Error
}

func (db *FormConnection) FindFormById(id uint) (formOutput model.SelectFormParameter, err error) {
	var (
		form model.SelectFormParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at").Joins("left join form_types ON forms.form_type_id = form_types.id").Where("forms.id=? AND forms.deleted_at = 0", id).Order("forms.id").Take(&form)
	return form, res.Error
}

func (db *FormConnection) FindFormByFormTypeId(ftId uint) (formOutput []model.SelectFormParameter, err error) {
	var (
		forms []model.SelectFormParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at").Joins("left join form_types ON forms.form_type_id = form_types.id").Where("forms.form_type_id=? AND forms.deleted_at = 0", ftId).Order("forms.id").Find(&forms)
	return forms, res.Error
}

func (db *FormConnection) FindExcForm(ftId uint, id uint) (formOutput []model.SelectFormParameter, err error) {
	var (
		forms []model.SelectFormParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at").Joins("left join form_types ON forms.form_type_id = form_types.id").Where("forms.form_type_id = ? AND forms.id!=? AND forms.deleted_at = 0", ftId, id).Order("forms.form_code").Find(&forms)
	return forms, res.Error
}

func (db *FormConnection) FindFormHeadDetail(id uint) (formOutput model.SelectFormParameter, err error) {
	var (
		form model.SelectFormParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at").Joins("left join form_types ON forms.form_type_id = form_types.id").Where("forms.id=? AND forms.deleted_at = 0", id).Order("forms.form_code").Take(&form)
	return form, res.Error
}

func (db *FormConnection) FindFormHead(ftId uint) (formOutput []model.SelectFormParameter, err error) {
	var (
		forms []model.SelectFormParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at").Joins("left join form_types ON forms.form_type_id = form_types.id").Where("forms.form_type_id=? AND forms.deleted_at = 0", ftId).Order("forms.form_code").Find(&forms)
	return forms, res.Error
}

func (db *FormConnection) FindExcFormHead(id uint) (formOutput []model.SelectFormParameter, err error) {
	var (
		forms []model.SelectFormParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at").Joins("left join form_types ON forms.form_type_id = form_types.id").Where("forms.id!=? AND forms.form_type_id = 1 AND forms.deleted_at = 0", id).Order("forms.form_code").Find(&forms)
	return forms, res.Error
}

func (db *FormConnection) FindExcFormOnly(id uint) (formOutput []model.SelectFormParameter, err error) {
	var (
		forms []model.SelectFormParameter
	)

	res := db.connection.Debug().Table("forms").Select("forms.id, forms.form_php, forms.form_code, forms.form_description, forms.form_type_id, form_types.Form_type_code, form_types.form_type_description, forms.form_parent_id, forms.sequence_no, forms.class_tag, forms.remark, forms.created_user_id, forms.updated_user_id, forms.deleted_user_id, forms.created_at, forms.updated_at, forms.deleted_at").Joins("left join form_types ON forms.form_type_id = form_types.id").Where("forms.id!=? AND forms.deleted_at = 0", id).Order("forms.form_code").Find(&forms)
	return forms, res.Error
}

func (db *FormConnection) InsertForm(form model.Form) (formOutput model.Form, err error) {
	res := db.connection.Save(&form)
	return form, res.Error
}

func (db *FormConnection) UpdateForm(form model.Form, id uint) (formOutput model.Form, err error) {
	res := db.connection.Where("id=?", id).Updates(&form)
	return form, res.Error
}
