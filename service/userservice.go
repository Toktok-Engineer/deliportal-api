package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	FindUsers() (userOutput []model.SelectUserParameter, err error)
	FindUserById(id uint) (userOutput model.SelectUserParameter, err error)
	FindUserByUName(uName string) (userOutput model.SelectUserParameter, err error)
	FindExcUser(id uint) (userOutput []model.SelectUserParameter, err error)
	InsertUser(user model.CreateUserParameter) (userOutput model.User, err error)
	UpdateUser(user model.User, id uint) (userOutput model.User, err error)
	DeleteUser(user model.User, id uint) (userOutput model.User, err error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRep repository.UserRepository) UserService {
	return &userService{
		userRepository: userRep,
	}
}

func (service *userService) FindUsers() (userOutput []model.SelectUserParameter, err error) {
	res, err := service.userRepository.FindUsers()
	return res, err
}

func (service *userService) FindUserById(id uint) (userOutput model.SelectUserParameter, err error) {
	res, err := service.userRepository.FindUserById(id)
	return res, err
}

func (service *userService) FindUserByUName(uName string) (userOutput model.SelectUserParameter, err error) {
	res, err := service.userRepository.FindUserByUName(uName)
	return res, err
}

func (service *userService) FindExcUser(id uint) (userOutput []model.SelectUserParameter, err error) {
	res, err := service.userRepository.FindExcUser(id)
	return res, err
}

func (service *userService) InsertUser(user model.CreateUserParameter) (userOutput model.User, err error) {
	newUser := model.User{}
	err1 := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err != nil {
		return newUser, err1
	}
	res, err := service.userRepository.InsertUser(newUser)
	return res, err
}

func (service *userService) UpdateUser(user model.User, id uint) (userOutput model.User, err error) {
	newUser := model.User{}
	err1 := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err != nil {
		return newUser, err1
	}
	res, err := service.userRepository.UpdateUser(newUser, id)
	return res, err
}

func (service *userService) DeleteUser(user model.User, id uint) (userOutput model.User, err error) {
	newUser := model.User{}
	err1 := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err != nil {
		return newUser, err1
	}
	res, err := service.userRepository.UpdateUser(newUser, id)
	return res, err
}
