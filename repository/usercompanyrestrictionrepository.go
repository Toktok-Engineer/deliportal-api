package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type UserCompanyRestrictionRepository interface {
	CountUserCompanyRestrictionAll(usernameID int) (count int64, err error)
	FindUserCompanyRestrictions() (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error)
	FindUserCompanyRestrictionsOffset(limit int, offset int, order string, dir string, usernameID int) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error)
	SearchUserCompanyRestriction(limit int, offset int, order string, dir string, search string, usernameID int) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error)
	CountSearchUserCompanyRestriction(search string, usernameID int) (count int64, err error)
	FindUserCompanyRestrictionById(id uint) (usercompanyrestrictionOutput model.SelectUserCompanyRestrictionParameter, err error)
	FindUserCompanyRestrictionByUserId(uid uint) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error)
	FindExcUserCompanyRestriction(id uint) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error)
	InsertUserCompanyRestriction(usercompanyrestriction model.UserCompanyRestriction) (usercompanyrestrictionOutput model.UserCompanyRestriction, err error)
	UpdateUserCompanyRestriction(usercompanyrestriction model.UserCompanyRestriction, id uint) (usercompanyrestrictionOutput model.UserCompanyRestriction, err error)
}

type UserCompanyRestrictionConnection struct {
	connection *gorm.DB
}

func NewUserCompanyRestrictionRepository(db *gorm.DB) UserCompanyRestrictionRepository {
	return &UserCompanyRestrictionConnection{
		connection: db,
	}
}

func (db *UserCompanyRestrictionConnection) CountUserCompanyRestrictionAll(usernameID int) (count int64, err error) {
	res := db.connection.Debug().Table("user_company_restrictions").Where("user_id = ? AND deleted_at = 0", usernameID).Count(&count)
	return count, res.Error
}

func (db *UserCompanyRestrictionConnection) FindUserCompanyRestrictions() (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error) {
	var (
		usercompanyrestrictions []model.SelectUserCompanyRestrictionParameter
	)
	res := db.connection.Debug().Table("user_company_restrictions").Select("user_company_restrictions.id, user_company_restrictions.user_id, users.username, user_company_restrictions.company_id, companies.company_name, user_company_restrictions.remark, user_company_restrictions.created_user_id, user_company_restrictions.updated_user_id, user_company_restrictions.deleted_user_id, user_company_restrictions.created_at, user_company_restrictions.updated_at, user_company_restrictions.deleted_at").Joins("left join users ON user_company_restrictions.user_id = users.id").Joins("left join companies ON user_company_restrictions.company_id = companies.id").Where("user_company_restrictions.deleted_at = 0").Order("users.username").Find(&usercompanyrestrictions)
	return usercompanyrestrictions, res.Error
}

func (db *UserCompanyRestrictionConnection) FindUserCompanyRestrictionsOffset(limit int, offset int, order string, dir string, usernameID int) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error) {
	var (
		orderDirection          string
		usercompanyrestrictions []model.SelectUserCompanyRestrictionParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("user_company_restrictions").Select("user_company_restrictions.id, user_company_restrictions.user_id, users.username, user_company_restrictions.company_id, companies.company_name, user_company_restrictions.remark, user_company_restrictions.created_user_id, user_company_restrictions.updated_user_id, user_company_restrictions.deleted_user_id, user_company_restrictions.created_at, user_company_restrictions.updated_at, user_company_restrictions.deleted_at").Joins("left join users ON user_company_restrictions.user_id = users.id").Joins("left join companies ON user_company_restrictions.company_id = companies.id").Where("user_company_restrictions.user_id = ? AND user_company_restrictions.deleted_at = 0", usernameID).Order(orderDirection).Limit(limit).Offset(offset).Find(&usercompanyrestrictions)
	return usercompanyrestrictions, res.Error
}

func (db *UserCompanyRestrictionConnection) SearchUserCompanyRestriction(limit int, offset int, order string, dir string, search string, usernameID int) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error) {
	var (
		orderDirection          string
		final                   string
		usercompanyrestrictions []model.SelectUserCompanyRestrictionParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("user_company_restrictions").Select("user_company_restrictions.id, user_company_restrictions.user_id, users.username, user_company_restrictions.company_id, companies.company_name, user_company_restrictions.remark, user_company_restrictions.created_user_id, user_company_restrictions.updated_user_id, user_company_restrictions.deleted_user_id, user_company_restrictions.created_at, user_company_restrictions.updated_at, user_company_restrictions.deleted_at").Joins("left join users ON user_company_restrictions.user_id = users.id").Joins("left join companies ON user_company_restrictions.company_id = companies.id").Where("(lower(users.username) LIKE ? OR lower(companies.company_name) LIKE ? OR lower(user_company_restrictions.remark) LIKE ?) AND user_company_restrictions.user_id = ? AND user_company_restrictions.deleted_at = 0", final, final, final, usernameID).Order(orderDirection).Limit(limit).Offset(offset).Find(&usercompanyrestrictions)
	return usercompanyrestrictions, res.Error
}

func (db *UserCompanyRestrictionConnection) CountSearchUserCompanyRestriction(search string, usernameID int) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("user_company_restrictions").Select("user_company_restrictions.id, user_company_restrictions.user_id, users.username, user_company_restrictions.company_id, companies.company_name, user_company_restrictions.remark, user_company_restrictions.created_user_id, user_company_restrictions.updated_user_id, user_company_restrictions.deleted_user_id, user_company_restrictions.created_at, user_company_restrictions.updated_at, user_company_restrictions.deleted_at").Joins("left join users ON user_company_restrictions.user_id = users.id").Joins("left join companies ON user_company_restrictions.company_id = companies.id").Where("(lower(users.username) LIKE ? OR lower(companies.company_name) LIKE ? OR lower(user_company_restrictions.remark) LIKE ?) AND user_company_restrictions.user_id = ? AND user_company_restrictions.deleted_at = 0", final, final, final, usernameID).Count(&count)
	return count, res.Error
}

func (db *UserCompanyRestrictionConnection) FindUserCompanyRestrictionById(id uint) (usercompanyrestrictionOutput model.SelectUserCompanyRestrictionParameter, err error) {
	var (
		user_company_restriction model.SelectUserCompanyRestrictionParameter
	)

	res := db.connection.Debug().Table("user_company_restrictions").Select("user_company_restrictions.id, user_company_restrictions.user_id, users.username, user_company_restrictions.company_id, companies.company_name, user_company_restrictions.remark, user_company_restrictions.created_user_id, user_company_restrictions.updated_user_id, user_company_restrictions.deleted_user_id, user_company_restrictions.created_at, user_company_restrictions.updated_at, user_company_restrictions.deleted_at").Joins("left join users ON user_company_restrictions.user_id = users.id").Joins("left join companies ON user_company_restrictions.company_id = companies.id").Where("user_company_restrictions.id=? AND user_company_restrictions.deleted_at = 0", id).Take(&user_company_restriction)
	return user_company_restriction, res.Error
}

func (db *UserCompanyRestrictionConnection) FindUserCompanyRestrictionByUserId(uid uint) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error) {
	var (
		user_company_restrictions []model.SelectUserCompanyRestrictionParameter
	)

	res := db.connection.Debug().Table("user_company_restrictions").Select("user_company_restrictions.id, user_company_restrictions.user_id, users.username, user_company_restrictions.company_id, companies.company_name, user_company_restrictions.remark, user_company_restrictions.created_user_id, user_company_restrictions.updated_user_id, user_company_restrictions.deleted_user_id, user_company_restrictions.created_at, user_company_restrictions.updated_at, user_company_restrictions.deleted_at").Joins("left join users ON user_company_restrictions.user_id = users.id").Joins("left join companies ON user_company_restrictions.company_id = companies.id").Where("user_company_restrictions.user_id=? AND user_company_restrictions.deleted_at = 0", uid).Order("user_company_restrictions.id").Find(&user_company_restrictions)
	return user_company_restrictions, res.Error
}

func (db *UserCompanyRestrictionConnection) FindExcUserCompanyRestriction(id uint) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error) {
	var (
		user_company_restrictions []model.SelectUserCompanyRestrictionParameter
	)

	res := db.connection.Debug().Table("user_company_restrictions").Select("user_company_restrictions.id, user_company_restrictions.user_id, users.username, user_company_restrictions.company_id, companies.company_name, user_company_restrictions.remark, user_company_restrictions.created_user_id, user_company_restrictions.updated_user_id, user_company_restrictions.deleted_user_id, user_company_restrictions.created_at, user_company_restrictions.updated_at, user_company_restrictions.deleted_at").Joins("left join users ON user_company_restrictions.user_id = users.id").Joins("left join companies ON user_company_restrictions.company_id = companies.id").Where("user_company_restrictions.id != ? AND user_company_restrictions.deleted_at = 0", id).Find(&user_company_restrictions)
	return user_company_restrictions, res.Error
}

func (db *UserCompanyRestrictionConnection) InsertUserCompanyRestriction(usercompanyrestriction model.UserCompanyRestriction) (usercompanyrestrictionOutput model.UserCompanyRestriction, err error) {
	res := db.connection.Save(&usercompanyrestriction)
	return usercompanyrestriction, res.Error
}

func (db *UserCompanyRestrictionConnection) UpdateUserCompanyRestriction(usercompanyrestriction model.UserCompanyRestriction, id uint) (usercompanyrestrictionOutput model.UserCompanyRestriction, err error) {
	res := db.connection.Where("id=?", id).Updates(&usercompanyrestriction)
	return usercompanyrestriction, res.Error
}
