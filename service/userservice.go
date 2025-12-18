package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	CountUserAll() (count int64, err error)
	FindUsersAll(id uint) (userOutput model.SelectUserParameter, err error)
	FindUsers() (userOutput []model.SelectUserParameter, err error)
	FindUsersOffset(limit int, offset int, order string, dir string) (userOutput []model.SelectUserParameter, err error)
	SearchUser(limit int, offset int, order string, dir string, search string) (userOutput []model.SelectUserParameter, err error)
	CountSearchUser(search string) (count int64, err error)
	FindUserById(id uint) (userOutput model.SelectUserParameter, err error)
	FindUserByUName(uName string) (userOutput model.SelectUserParameter, err error)
	FindByEmployeeID(employeeID uint) (userOutput model.SelectUserParameter, err error)
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

func (service *userService) CountUserAll() (count int64, err error) {
	res, err := service.userRepository.CountUserAll()
	return res, err
}

func (service *userService) FindUsersAll(id uint) (userOutput model.SelectUserParameter, err error) {
	res, err := service.userRepository.FindUsersAll(id)
	return res, err
}

func (service *userService) FindUsers() (userOutput []model.SelectUserParameter, err error) {
	res, err := service.userRepository.FindUsers()
	return res, err
}

func (service *userService) FindUsersOffset(limit int, offset int, order string, dir string) (userOutput []model.SelectUserParameter, err error) {
	res, err := service.userRepository.FindUsersOffset(limit, offset, order, dir)
	return res, err
}

func (service *userService) SearchUser(limit int, offset int, order string, dir string, search string) (userOutput []model.SelectUserParameter, err error) {
	res, err := service.userRepository.SearchUser(limit, offset, order, dir, search)
	return res, err
}

func (service *userService) CountSearchUser(search string) (count int64, err error) {
	res, err := service.userRepository.CountSearchUser(search)
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

func (service *userService) FindByEmployeeID(employeeID uint) (userOutput model.SelectUserParameter, err error) {
	res, err := service.userRepository.FindByEmployeeID(employeeID)
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
