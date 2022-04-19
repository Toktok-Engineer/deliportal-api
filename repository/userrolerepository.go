package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type UserRoleRepository interface {
	FindUserRoles() (userRoleOutput []model.SelectUserRoleParameter, err error)
	FindUserRoleById(id uint) (userRoleOutput model.SelectUserRoleParameter, err error)
	FindUserRoleByUserId(uid uint) (userRoleOutput model.SelectUserRoleParameter, err error)
	FindExcUserRole(id uint, uid uint) (userRoleOutput []model.SelectUserRoleParameter, err error)
	FindExcUserRoleOnly(id uint) (userRoleOutput []model.SelectUserRoleParameter, err error)
	InsertUserRole(userRole model.UserRole) (userRoleOutput model.UserRole, err error)
	UpdateUserRole(userRole model.UserRole, id uint) (userRoleOutput model.UserRole, err error)
}

type UserRoleConnection struct {
	connection *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &UserRoleConnection{
		connection: db,
	}
}

func (db *UserRoleConnection) FindUserRoles() (userRoleOutput []model.SelectUserRoleParameter, err error) {
	var (
		user_roles []model.SelectUserRoleParameter
	)

	res := db.connection.Debug().Table("user_roles").Select("user_roles.id,	user_roles.user_id, users.username,	users.password,	users.employee_id, users.email, user_roles.role_id,	roles.role_code, roles.role_description, user_roles.remark,	user_roles.created_user_id, user_roles.updated_user_id,	user_roles.deleted_user_id, user_roles.created_at, user_roles.updated_at, user_roles.deleted_at").Joins("left join users ON user_roles.user_id = users.id").Joins("left join roles ON user_roles.role_id = roles.id").Where("user_roles.deleted_at = 0").Order("user_roles.id").Find(&user_roles)
	return user_roles, res.Error
}

func (db *UserRoleConnection) FindUserRoleById(id uint) (userRoleOutput model.SelectUserRoleParameter, err error) {
	var (
		userRole model.SelectUserRoleParameter
	)

	res := db.connection.Debug().Table("user_roles").Select("user_roles.id,	user_roles.user_id, users.username,	users.password,	users.employee_id, users.email, user_roles.role_id,	roles.role_code, roles.role_description, user_roles.remark,	user_roles.created_user_id, user_roles.updated_user_id,	user_roles.deleted_user_id, user_roles.created_at, user_roles.updated_at, user_roles.deleted_at").Joins("left join users ON user_roles.user_id = users.id").Joins("left join roles ON user_roles.role_id = roles.id").Where("user_roles.id=? AND user_roles.deleted_at = 0", id).Take(&userRole)
	return userRole, res.Error
}

func (db *UserRoleConnection) FindUserRoleByUserId(uid uint) (userRoleOutput model.SelectUserRoleParameter, err error) {
	var (
		userRole model.SelectUserRoleParameter
	)

	res := db.connection.Debug().Table("user_roles").Select("user_roles.id,	user_roles.user_id, users.username,	users.password,	users.employee_id, users.email, user_roles.role_id,	roles.role_code, roles.role_description, user_roles.remark,	user_roles.created_user_id, user_roles.updated_user_id,	user_roles.deleted_user_id, user_roles.created_at, user_roles.updated_at, user_roles.deleted_at").Joins("left join users ON user_roles.user_id = users.id").Joins("left join roles ON user_roles.role_id = roles.id").Where("user_roles.user_id=? AND user_roles.deleted_at = 0", uid).Take(&userRole)
	return userRole, res.Error
}

func (db *UserRoleConnection) FindExcUserRole(id uint, uid uint) (userRoleOutput []model.SelectUserRoleParameter, err error) {
	var (
		user_roles []model.SelectUserRoleParameter
	)

	res := db.connection.Debug().Table("user_roles").Select("user_roles.id,	user_roles.user_id, users.username,	users.password,	users.employee_id, users.email, user_roles.role_id,	roles.role_code, roles.role_description, user_roles.remark,	user_roles.created_user_id, user_roles.updated_user_id,	user_roles.deleted_user_id, user_roles.created_at, user_roles.updated_at, user_roles.deleted_at").Joins("left join users ON user_roles.user_id = users.id").Joins("left join roles ON user_roles.role_id = roles.id").Where("user_roles.user_id=? AND user_roles.user_id = ? AND user_roles.deleted_at = 0", id, uid).Find(&user_roles)
	return user_roles, res.Error
}

func (db *UserRoleConnection) FindExcUserRoleOnly(id uint) (userRoleOutput []model.SelectUserRoleParameter, err error) {
	var (
		user_roles []model.SelectUserRoleParameter
	)

	res := db.connection.Debug().Table("user_roles").Select("user_roles.id,	user_roles.user_id, users.username,	users.password,	users.employee_id, users.email, user_roles.role_id,	roles.role_code, roles.role_description, user_roles.remark,	user_roles.created_user_id, user_roles.updated_user_id,	user_roles.deleted_user_id, user_roles.created_at, user_roles.updated_at, user_roles.deleted_at").Joins("left join users ON user_roles.user_id = users.id").Joins("left join roles ON user_roles.role_id = roles.id").Where("user_roles.user_id=? AND user_roles.deleted_at = 0", id).Find(&user_roles)
	return user_roles, res.Error
}

func (db *UserRoleConnection) InsertUserRole(userRole model.UserRole) (userRoleOutput model.UserRole, err error) {
	res := db.connection.Save(&userRole)
	return userRole, res.Error
}

func (db *UserRoleConnection) UpdateUserRole(userRole model.UserRole, id uint) (userRoleOutput model.UserRole, err error) {
	res := db.connection.Where("id=?", id).Updates(&userRole)
	return userRole, res.Error
}
