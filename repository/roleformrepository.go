package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type RoleFormRepository interface {
	FindRoleForms() (roleFormOutput []model.SelectRoleFormParameter, err error)
	FindRoleFormById(id uint) (roleFormOutput model.SelectRoleFormParameter, err error)
	FindRoleFormByFormId(fid uint, rid uint) (roleFormOutput model.SelectRoleFormParameter, err error)
	FindExcRoleForm(id uint, rid uint) (roleFormOutput []model.SelectRoleFormParameter, err error)
	FindExcRoleFormOnly(id uint) (roleFormOutput []model.SelectRoleFormParameter, err error)
	InsertRoleForm(roleForm model.RoleForm) (roleFormOutput model.RoleForm, err error)
	UpdateRoleForm(roleForm model.RoleForm, id uint) (roleFormOutput model.RoleForm, err error)
}

type RoleFormConnection struct {
	connection *gorm.DB
}

func NewRoleFormRepository(db *gorm.DB) RoleFormRepository {
	return &RoleFormConnection{
		connection: db,
	}
}

func (db *RoleFormConnection) FindRoleForms() (roleFormOutput []model.SelectRoleFormParameter, err error) {
	var (
		role_forms []model.SelectRoleFormParameter
	)

	res := db.connection.Debug().Table("role_forms").Select("role_forms.id, role_forms.role_id, roles.role_code, roles.role_description, role_forms.form_id, forms.form_code, forms.form_description, role_forms.create_flag, role_forms.read_flag, role_forms.update_flag, role_forms.delete_flag, role_forms.remark, role_forms.created_user_id, role_forms.updated_user_id, role_forms.deleted_user_id, role_forms.created_at, role_forms.updated_at, role_forms.deleted_at ").Joins("left join roles ON role_forms.role_id = roles.id").Joins("left join forms ON role_forms.form_id = forms.id").Where("role_forms.deleted_at = 0").Order("role_forms.id").Find(&role_forms)
	return role_forms, res.Error
}

func (db *RoleFormConnection) FindRoleFormById(id uint) (roleFormOutput model.SelectRoleFormParameter, err error) {
	var (
		roleForm model.SelectRoleFormParameter
	)

	res := db.connection.Debug().Table("role_forms").Select("role_forms.id, role_forms.role_id, roles.role_code, roles.role_description, role_forms.form_id, forms.form_code, forms.form_description, role_forms.create_flag, role_forms.read_flag, role_forms.update_flag, role_forms.delete_flag, role_forms.remark, role_forms.created_user_id, role_forms.updated_user_id, role_forms.deleted_user_id, role_forms.created_at, role_forms.updated_at, role_forms.deleted_at ").Joins("left join roles ON role_forms.role_id = roles.id").Joins("left join forms ON role_forms.form_id = forms.id").Where("role_forms.id=? AND role_forms.deleted_at = 0", id).Take(&roleForm)
	return roleForm, res.Error
}

func (db *RoleFormConnection) FindRoleFormByFormId(fid uint, rid uint) (roleFormOutput model.SelectRoleFormParameter, err error) {
	var (
		roleForm model.SelectRoleFormParameter
	)

	res := db.connection.Debug().Table("role_forms").Select("role_forms.id, role_forms.role_id, roles.role_code, roles.role_description, role_forms.form_id, forms.form_code, forms.form_description, role_forms.create_flag, role_forms.read_flag, role_forms.update_flag, role_forms.delete_flag, role_forms.remark, role_forms.created_user_id, role_forms.updated_user_id, role_forms.deleted_user_id, role_forms.created_at, role_forms.updated_at, role_forms.deleted_at ").Joins("left join roles ON role_forms.role_id = roles.id").Joins("left join forms ON role_forms.form_id = forms.id").Where("role_forms.form_id=? AND role_forms.role_id=? AND role_forms.deleted_at = 0", fid, rid).Take(&roleForm)
	return roleForm, res.Error
}

func (db *RoleFormConnection) FindExcRoleForm(id uint, rid uint) (roleFormOutput []model.SelectRoleFormParameter, err error) {
	var (
		role_forms []model.SelectRoleFormParameter
	)

	res := db.connection.Debug().Table("role_forms").Select("role_forms.id, role_forms.role_id, roles.role_code, roles.role_description, role_forms.form_id, forms.form_code, forms.form_description, role_forms.create_flag, role_forms.read_flag, role_forms.update_flag, role_forms.delete_flag, role_forms.remark, role_forms.created_user_id, role_forms.updated_user_id, role_forms.deleted_user_id, role_forms.created_at, role_forms.updated_at, role_forms.deleted_at ").Joins("left join roles ON role_forms.role_id = roles.id").Joins("left join forms ON role_forms.form_id = forms.id").Where("role_forms.id!=? AND role_forms.role_id=? AND role_forms.deleted_at = 0", id, rid).Find(&role_forms)
	return role_forms, res.Error
}

func (db *RoleFormConnection) FindExcRoleFormOnly(id uint) (roleFormOutput []model.SelectRoleFormParameter, err error) {
	var (
		role_forms []model.SelectRoleFormParameter
	)

	res := db.connection.Debug().Table("role_forms").Select("role_forms.id, role_forms.role_id, roles.role_code, roles.role_description, role_forms.form_id, forms.form_code, forms.form_description, role_forms.create_flag, role_forms.read_flag, role_forms.update_flag, role_forms.delete_flag, role_forms.remark, role_forms.created_user_id, role_forms.updated_user_id, role_forms.deleted_user_id, role_forms.created_at, role_forms.updated_at, role_forms.deleted_at ").Joins("left join roles ON role_forms.role_id = roles.id").Joins("left join forms ON role_forms.form_id = forms.id").Where("role_forms.id!=? AND role_forms.deleted_at = 0", id).Find(&role_forms)
	return role_forms, res.Error
}

func (db *RoleFormConnection) InsertRoleForm(roleForm model.RoleForm) (roleFormOutput model.RoleForm, err error) {
	res := db.connection.Save(&roleForm)
	return roleForm, res.Error
}

func (db *RoleFormConnection) UpdateRoleForm(roleForm model.RoleForm, id uint) (roleFormOutput model.RoleForm, err error) {
	res := db.connection.Where("id=?", id).Updates(&roleForm)
	return roleForm, res.Error
}
