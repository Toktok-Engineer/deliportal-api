package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type RoleRepository interface {
	CountRoleAll() (count int64, err error)
	FindRoles() (roleOutput []model.Role, err error)
	FindRolesOffset(limit int, offset int, order string, dir string) (roleOutput []model.Role, err error)
	SearchRole(limit int, offset int, order string, dir string, search string) (roleOutput []model.Role, err error)
	CountSearchRole(search string) (count int64, err error)
	FindRoleById(id uint) (roleOutput model.Role, err error)
	FindExcRole(id uint) (roleOutput []model.Role, err error)
	InsertRole(role model.Role) (roleOutput model.Role, err error)
	UpdateRole(role model.Role, id uint) (roleOutput model.Role, err error)
}

type RoleConnection struct {
	connection *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleConnection{
		connection: db,
	}
}

func (db *RoleConnection) CountRoleAll() (count int64, err error) {
	res := db.connection.Debug().Table("roles").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *RoleConnection) FindRoles() (roleOutput []model.Role, err error) {
	var (
		roles []model.Role
	)
	res := db.connection.Where("deleted_at = 0").Order("role_code").Find(&roles)
	return roles, res.Error
}

func (db *RoleConnection) FindRolesOffset(limit int, offset int, order string, dir string) (roleOutput []model.Role, err error) {
	var (
		orderDirection string
		roles          []model.Role
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&roles)
	return roles, res.Error
}

func (db *RoleConnection) SearchRole(limit int, offset int, order string, dir string, search string) (roleOutput []model.Role, err error) {
	var (
		orderDirection string
		final          string
		roles          []model.Role
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(role_code) LIKE ? OR lower(role_description) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&roles)
	return roles, res.Error
}

func (db *RoleConnection) CountSearchRole(search string) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("roles").Where("(lower(role_code) LIKE ? OR lower(role_description) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final).Count(&count)
	return count, res.Error
}

func (db *RoleConnection) FindRoleById(id uint) (roleOutput model.Role, err error) {
	var (
		role model.Role
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&role)
	return role, res.Error
}

func (db *RoleConnection) FindExcRole(id uint) (roleOutput []model.Role, err error) {
	var (
		roles []model.Role
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("role_code").Find(&roles)
	return roles, res.Error
}

func (db *RoleConnection) InsertRole(role model.Role) (roleOutput model.Role, err error) {
	res := db.connection.Save(&role)
	return role, res.Error
}

func (db *RoleConnection) UpdateRole(role model.Role, id uint) (roleOutput model.Role, err error) {
	res := db.connection.Where("id=?", id).Updates(&role)
	return role, res.Error
}
