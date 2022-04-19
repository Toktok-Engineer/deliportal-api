package repository

import (
	"deliportal-api/model"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(userInput model.User) (model.User, error)
	UpdateUser(user model.User, id uint) (model.User, error)
	VerifyCredential(username string) interface{}
	IsDuplicateUsername(username string) (tx *gorm.DB)
	FindByUsername(username string) model.User
	IsUserRegistered(id uint64, username string) (tx *gorm.DB)
	FindUsers() (userOutput []model.SelectUserParameter, err error)
	FindUserById(id uint) (userOutput model.SelectUserParameter, err error)
	FindUserByUName(uName string) (userOutput model.SelectUserParameter, err error)
	FindExcUser(id uint) (userOutput []model.SelectUserParameter, err error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(userInput model.User) (model.User, error) {
	userInput.Password = hashAndSalt([]byte(userInput.Password))
	res := db.connection.Save(&userInput)
	return userInput, res.Error
}

func (db *userConnection) UpdateUser(user model.User, id uint) (model.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser model.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	res := db.connection.Where("id=?", id).Updates(&user)
	return user, res.Error
}

func (db *userConnection) VerifyCredential(username string) interface{} {
	var user model.User
	res := db.connection.Where("username = ? AND deleted_at = 0", username).Take(&user)
	if res.Error == nil {
		return user
	}

	fmt.Print(res.Error.Error())
	return user
}

func (db *userConnection) IsDuplicateUsername(username string) (tx *gorm.DB) {
	var user model.User
	return db.connection.Where("username = ?", username).Take(&user)
}

func (db *userConnection) FindByUsername(username string) model.User {
	var user model.User
	db.connection.Where("username = ?", username).Take(&user)
	return user
}

func (db *userConnection) IsUserRegistered(id uint64, username string) (tx *gorm.DB) {
	var user model.User
	return db.connection.Where("id = ? AND username = ?", id, username).Take(&user)
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func (db *userConnection) FindUsers() (userOutput []model.SelectUserParameter, err error) {
	var (
		users []model.SelectUserParameter
	)

	res := db.connection.Debug().Table("users").Select("users.id, users.username, users.password, users.employee_id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, users.email, users.remark, users.created_user_id, users.updated_user_id, users.deleted_user_id, users.created_at, users.updated_at, users.deleted_at").Joins("left join employees ON users.employee_id = employees.id").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("users.deleted_at = 0").Order("users.username").Find(&users)
	return users, res.Error
}

func (db *userConnection) FindUserById(id uint) (userOutput model.SelectUserParameter, err error) {
	var (
		user model.SelectUserParameter
	)

	res := db.connection.Debug().Table("users").Select("users.id, users.username, users.password, users.employee_id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, users.email, users.remark, users.created_user_id, users.updated_user_id, users.deleted_user_id, users.created_at, users.updated_at, users.deleted_at").Joins("left join employees ON users.employee_id = employees.id").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("users.id = ? AND users.deleted_at = 0", id).Take(&user)
	return user, res.Error
}

// FindUserByUName implements UserRepository
func (db *userConnection) FindUserByUName(uName string) (userOutput model.SelectUserParameter, err error) {
	var (
		user model.SelectUserParameter
	)

	res := db.connection.Debug().Table("users").Select("users.id, users.username, users.password, users.employee_id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, users.email, users.remark, users.created_user_id, users.updated_user_id, users.deleted_user_id, users.created_at, users.updated_at, users.deleted_at").Joins("left join employees ON users.employee_id = employees.id").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("users.username = ? AND users.deleted_at = 0", uName).Take(&user)
	return user, res.Error
}

// FindExcUser implements UserRepository
func (db *userConnection) FindExcUser(id uint) (userOutput []model.SelectUserParameter, err error) {
	var (
		users []model.SelectUserParameter
	)

	res := db.connection.Debug().Table("users").Select("users.id, users.username, users.password, users.employee_id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, users.email, users.remark, users.created_user_id, users.updated_user_id, users.deleted_user_id, users.created_at, users.updated_at, users.deleted_at").Joins("left join employees ON users.employee_id = employees.id").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("users.id != ? AND users.deleted_at = 0", id).Order("users.username").Find(&users)
	return users, res.Error
}
