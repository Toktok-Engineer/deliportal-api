package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type UserRoleService interface {
	CountUserRoleAll(usernameID int) (count int64, err error)
	FindUserRoles() (userRoleOutput []model.SelectUserRoleParameter, err error)
	FindUserRolesOffset(limit int, offset int, order string, dir string, usernameID int) (userRoleOutput []model.SelectUserRoleParameter, err error)
	SearchUserRole(limit int, offset int, order string, dir string, search string, usernameID int) (userRoleOutput []model.SelectUserRoleParameter, err error)
	CountSearchUserRole(search string, usernameID int) (count int64, err error)
	FindUserRoleById(id uint) (userRoleOutput model.SelectUserRoleParameter, err error)
	FindUserRoleByUserId(uid uint) (userRoleOutput model.SelectUserRoleParameter, err error)
	FindExcUserRole(id uint, uid uint) (userRoleOutput []model.SelectUserRoleParameter, err error)
	FindExcUserRoleOnly(id uint) (userRoleOutput []model.SelectUserRoleParameter, err error)
	InsertUserRole(userRole model.CreateUserRoleParameter) (userRoleOutput model.UserRole, err error)
	UpdateUserRole(userRole model.UserRole, id uint) (userRoleOutput model.UserRole, err error)
	DeleteUserRole(userRole model.UserRole, id uint) (userRoleOutput model.UserRole, err error)
}

type userRoleService struct {
	userRoleRepository repository.UserRoleRepository
}

func NewUserRoleService(userRoleRep repository.UserRoleRepository) UserRoleService {
	return &userRoleService{
		userRoleRepository: userRoleRep,
	}
}

func (service *userRoleService) CountUserRoleAll(usernameID int) (count int64, err error) {
	res, err := service.userRoleRepository.CountUserRoleAll(usernameID)
	return res, err
}

func (service *userRoleService) FindUserRoles() (userRoleOutput []model.SelectUserRoleParameter, err error) {
	res, err := service.userRoleRepository.FindUserRoles()
	return res, err
}

func (service *userRoleService) FindUserRolesOffset(limit int, offset int, order string, dir string, usernameID int) (userRoleOutput []model.SelectUserRoleParameter, err error) {
	res, err := service.userRoleRepository.FindUserRolesOffset(limit, offset, order, dir, usernameID)
	return res, err
}

func (service *userRoleService) SearchUserRole(limit int, offset int, order string, dir string, search string, usernameID int) (userRoleOutput []model.SelectUserRoleParameter, err error) {
	res, err := service.userRoleRepository.SearchUserRole(limit, offset, order, dir, search, usernameID)
	return res, err
}

func (service *userRoleService) CountSearchUserRole(search string, usernameID int) (count int64, err error) {
	res, err := service.userRoleRepository.CountSearchUserRole(search, usernameID)
	return res, err
}

func (service *userRoleService) FindUserRoleById(id uint) (userRoleOutput model.SelectUserRoleParameter, err error) {
	res, err := service.userRoleRepository.FindUserRoleById(id)
	return res, err
}

func (service *userRoleService) FindUserRoleByUserId(uid uint) (userRoleOutput model.SelectUserRoleParameter, err error) {
	res, err := service.userRoleRepository.FindUserRoleByUserId(uid)
	return res, err
}

func (service *userRoleService) FindExcUserRole(id uint, uid uint) (userRoleOutput []model.SelectUserRoleParameter, err error) {
	res, err := service.userRoleRepository.FindExcUserRole(id, uid)
	return res, err
}

func (service *userRoleService) FindExcUserRoleOnly(id uint) (userRoleOutput []model.SelectUserRoleParameter, err error) {
	res, err := service.userRoleRepository.FindExcUserRoleOnly(id)
	return res, err
}

func (service *userRoleService) InsertUserRole(userRole model.CreateUserRoleParameter) (userRoleOutput model.UserRole, err error) {
	newUserRole := model.UserRole{}
	err1 := smapping.FillStruct(&newUserRole, smapping.MapFields(&userRole))
	if err != nil {
		return newUserRole, err1
	}
	res, err := service.userRoleRepository.InsertUserRole(newUserRole)
	return res, err
}

func (service *userRoleService) UpdateUserRole(userRole model.UserRole, id uint) (userRoleOutput model.UserRole, err error) {
	newUserRole := model.UserRole{}
	err1 := smapping.FillStruct(&newUserRole, smapping.MapFields(&userRole))
	if err != nil {
		return newUserRole, err1
	}
	res, err := service.userRoleRepository.UpdateUserRole(newUserRole, id)
	return res, err
}

func (service *userRoleService) DeleteUserRole(userRole model.UserRole, id uint) (userRoleOutput model.UserRole, err error) {
	newUserRole := model.UserRole{}
	err1 := smapping.FillStruct(&newUserRole, smapping.MapFields(&userRole))
	if err != nil {
		return newUserRole, err1
	}
	res, err := service.userRoleRepository.UpdateUserRole(newUserRole, id)
	return res, err
}
