package repository

import (
	"deliportal-api/model"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(userInput model.User) (userOutput model.User, err error)
	UpdateUser(user model.User) model.User
	VerifyCredential(username string) interface{}
	IsDuplicateUsername(username string) (tx *gorm.DB)
	FindByUsername(username string) model.User
	IsUserRegistered(id uint64, username string) (tx *gorm.DB)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(userInput model.User) (userOutput model.User, err error) {
	userInput.Password = hashAndSalt([]byte(userInput.Password))
	res := db.connection.Save(&userInput)

	if res.Error != nil {
		log.Println(res.Error.Error())
		return userInput, res.Error
	}

	return userInput, nil
}

func (db *userConnection) UpdateUser(user model.User) model.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser model.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	return user
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
