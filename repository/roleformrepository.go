package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type RoleFormRepository interface {
	CountRoleFormAll() (count int64, err error)
	FindRoleForms() (roleFormOutput []model.SelectRoleFormParameter, err error)
	FindRoleFormsOffset(limit int, offset int, order string, dir string) (roleFormOutput []model.SelectRoleFormParameter, err error)
	SearchRoleForm(limit int, offset int, order string, dir string, search string) (roleFormOutput []model.SelectRoleFormParameter, err error)
	CountSearchRoleForm(search string) (count int64, err error)
	FindRoleFormById(id uint) (roleFormOutput model.SelectRoleFormParameter, err error)
	FindRoleFormByFormId(fid uint, rid uint) (roleFormOutput model.SelectRoleFormParameter, err error)
	FindExcRoleForm(id uint, rid uint) (roleFormOutput []model.SelectRoleFormParameter, err error)
	FindExcRoleFormOnly(id uint) (roleFormOutput []model.SelectRoleFormParameter, err error)
	InsertRoleForm(roleForm model.RoleForm) (roleFormOutput model.RoleForm, err error)
	UpdateRoleForm(roleForm model.RoleForm, id uint) (roleFormOutput model.RoleForm, err error)
	DeleteRoleForm(roleForm model.RoleForm, id uint) (roleFormOutput model.RoleForm, err error)
}

type RoleFormConnection struct {
	connection *gorm.DB
}

func NewRoleFormRepository(db *gorm.DB) RoleFormRepository {
	return &RoleFormConnection{
		connection: db,
	}
}

func (db *RoleFormConnection) CountRoleFormAll() (count int64, err error) {
	res := db.connection.Debug().Table("role_forms").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *RoleFormConnection) FindRoleForms() (roleFormOutput []model.SelectRoleFormParameter, err error) {
	var (
		roleForms []model.SelectRoleFormParameter
	)
	res := db.connection.Debug().Table("role_forms").Select("role_forms.id, role_forms.role_id, roles.role_code, roles.role_description, role_forms.form_id, forms.form_code, forms.form_description, role_forms.create_flag, role_forms.read_flag, role_forms.update_flag, role_forms.delete_flag, role_forms.remark, role_forms.created_user_id, role_forms.updated_user_id, role_forms.deleted_user_id, role_forms.created_at, role_forms.updated_at, role_forms.deleted_at ").Joins("left join roles ON role_forms.role_id = roles.id").Joins("left join forms ON role_forms.form_id = forms.id").Where("role_forms.deleted_at = 0").Order("role_forms.roleForm_name").Find(&roleForms)
	return roleForms, res.Error
}

func (db *RoleFormConnection) FindRoleFormsOffset(limit int, offset int, order string, dir string) (roleFormOutput []model.SelectRoleFormParameter, err error) {
	var (
		orderDirection string
		roleForms      []model.SelectRoleFormParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("role_forms").Select("role_forms.id, role_forms.role_id, roles.role_code, roles.role_description, role_forms.form_id, forms.form_code, forms.form_description, role_forms.create_flag, role_forms.read_flag, role_forms.update_flag, role_forms.delete_flag, role_forms.remark, role_forms.created_user_id, role_forms.updated_user_id, role_forms.deleted_user_id, role_forms.created_at, role_forms.updated_at, role_forms.deleted_at ").Joins("left join roles ON role_forms.role_id = roles.id").Joins("left join forms ON role_forms.form_id = forms.id").Where("role_forms.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&roleForms)
	return roleForms, res.Error
}

func (db *RoleFormConnection) SearchRoleForm(limit int, offset int, order string, dir string, search string) (roleFormOutput []model.SelectRoleFormParameter, err error) {
	var (
		orderDirection string
		final          string
		roleForms      []model.SelectRoleFormParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("role_forms").Select("role_forms.id, role_forms.role_id, roles.role_code, roles.role_description, role_forms.form_id, forms.form_code, forms.form_description, role_forms.create_flag, role_forms.read_flag, role_forms.update_flag, role_forms.delete_flag, role_forms.remark, role_forms.created_user_id, role_forms.updated_user_id, role_forms.deleted_user_id, role_forms.created_at, role_forms.updated_at, role_forms.deleted_at ").Joins("left join roles ON role_forms.role_id = roles.id").Joins("left join forms ON role_forms.form_id = forms.id").Where("(lower(roles.role_code) LIKE ? OR lower(forms.form_code) LIKE ? OR lower(role_forms.remark) LIKE ?) AND role_forms.deleted_at = 0", final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&roleForms)
	return roleForms, res.Error
}

func (db *RoleFormConnection) CountSearchRoleForm(search string) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("role_forms").Select("role_forms.id, role_forms.role_id, roles.role_code, roles.role_description, role_forms.form_id, forms.form_code, forms.form_description, role_forms.create_flag, role_forms.read_flag, role_forms.update_flag, role_forms.delete_flag, role_forms.remark, role_forms.created_user_id, role_forms.updated_user_id, role_forms.deleted_user_id, role_forms.created_at, role_forms.updated_at, role_forms.deleted_at ").Joins("left join roles ON role_forms.role_id = roles.id").Joins("left join forms ON role_forms.form_id = forms.id").Where("(lower(roles.role_code) LIKE ? OR lower(forms.form_code) LIKE ? OR lower(role_forms.remark) LIKE ?) AND role_forms.deleted_at = 0", final, final, final).Count(&count)
	return count, res.Error
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
	var (
		role_form model.RoleForm
	)
	res := db.connection.Model(&role_form).Where("id=?", id).Updates(map[string]interface{}{"role_id": roleForm.RoleID, "form_id": roleForm.FormID, "create_flag": roleForm.CreateFlag, "read_flag": roleForm.ReadFlag, "update_flag": roleForm.UpdateFlag, "delete_flag": roleForm.DeleteFlag, "remark": roleForm.Remark, "updated_user_id": roleForm.UpdatedUserID, "updated_at": roleForm.UpdatedAt})
	return roleForm, res.Error
}

func (db *RoleFormConnection) DeleteRoleForm(roleForm model.RoleForm, id uint) (roleFormOutput model.RoleForm, err error) {
	var (
		role_form model.RoleForm
	)
	res := db.connection.Model(&role_form).Where("id=?", id).Updates(map[string]interface{}{"deleted_user_id": roleForm.DeletedUserID, "deleted_at": roleForm.DeletedAt})
	return roleForm, res.Error
}
