package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type UserCompanyRestrictionService interface {
	CountUserCompanyRestrictionAll(usernameID int) (count int64, err error)
	FindUserCompanyRestrictions() (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error)
	FindUserCompanyRestrictionsOffset(limit int, offset int, order string, dir string, usernameID int) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error)
	SearchUserCompanyRestriction(limit int, offset int, order string, dir string, search string, usernameID int) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error)
	CountSearchUserCompanyRestriction(search string, usernameID int) (count int64, err error)
	FindUserCompanyRestrictionById(id uint) (usercompanyrestrictionOutput model.SelectUserCompanyRestrictionParameter, err error)
	FindUserCompanyRestrictionByUserId(uid uint) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error)
	FindExcUserCompanyRestriction(id uint) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error)
	InsertUserCompanyRestriction(usercompanyrestriction model.CreateUserCompanyRestrictionParameter) (usercompanyrestrictionOutput model.UserCompanyRestriction, err error)
	UpdateUserCompanyRestriction(usercompanyrestriction model.UserCompanyRestriction, id uint) (usercompanyrestrictionOutput model.UserCompanyRestriction, err error)
	DeleteUserCompanyRestriction(usercompanyrestriction model.UserCompanyRestriction, id uint) (usercompanyrestrictionOutput model.UserCompanyRestriction, err error)
}

type usercompanyrestrictionService struct {
	usercompanyrestrictionRepository repository.UserCompanyRestrictionRepository
}

func NewUserCompanyRestrictionService(usercompanyrestrictionRep repository.UserCompanyRestrictionRepository) UserCompanyRestrictionService {
	return &usercompanyrestrictionService{
		usercompanyrestrictionRepository: usercompanyrestrictionRep,
	}
}

func (service *usercompanyrestrictionService) CountUserCompanyRestrictionAll(usernameID int) (count int64, err error) {
	res, err := service.usercompanyrestrictionRepository.CountUserCompanyRestrictionAll(usernameID)
	return res, err
}

func (service *usercompanyrestrictionService) FindUserCompanyRestrictions() (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error) {
	res, err := service.usercompanyrestrictionRepository.FindUserCompanyRestrictions()
	return res, err
}

func (service *usercompanyrestrictionService) FindUserCompanyRestrictionsOffset(limit int, offset int, order string, dir string, usernameID int) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error) {
	res, err := service.usercompanyrestrictionRepository.FindUserCompanyRestrictionsOffset(limit, offset, order, dir, usernameID)
	return res, err
}

func (service *usercompanyrestrictionService) SearchUserCompanyRestriction(limit int, offset int, order string, dir string, search string, usernameID int) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error) {
	res, err := service.usercompanyrestrictionRepository.SearchUserCompanyRestriction(limit, offset, order, dir, search, usernameID)
	return res, err
}

func (service *usercompanyrestrictionService) CountSearchUserCompanyRestriction(search string, usernameID int) (count int64, err error) {
	res, err := service.usercompanyrestrictionRepository.CountSearchUserCompanyRestriction(search, usernameID)
	return res, err
}

func (service *usercompanyrestrictionService) FindUserCompanyRestrictionById(id uint) (usercompanyrestrictionOutput model.SelectUserCompanyRestrictionParameter, err error) {
	res, err := service.usercompanyrestrictionRepository.FindUserCompanyRestrictionById(id)
	return res, err
}

func (service *usercompanyrestrictionService) FindUserCompanyRestrictionByUserId(uid uint) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error) {
	res, err := service.usercompanyrestrictionRepository.FindUserCompanyRestrictionByUserId(uid)
	return res, err
}

func (service *usercompanyrestrictionService) FindExcUserCompanyRestriction(id uint) (usercompanyrestrictionOutput []model.SelectUserCompanyRestrictionParameter, err error) {
	res, err := service.usercompanyrestrictionRepository.FindExcUserCompanyRestriction(id)
	return res, err
}

func (service *usercompanyrestrictionService) InsertUserCompanyRestriction(usercompanyrestriction model.CreateUserCompanyRestrictionParameter) (usercompanyrestrictionOutput model.UserCompanyRestriction, err error) {
	newUserCompanyRestriction := model.UserCompanyRestriction{}
	err1 := smapping.FillStruct(&newUserCompanyRestriction, smapping.MapFields(&usercompanyrestriction))
	if err != nil {
		return newUserCompanyRestriction, err1
	}
	res, err := service.usercompanyrestrictionRepository.InsertUserCompanyRestriction(newUserCompanyRestriction)
	return res, err
}

func (service *usercompanyrestrictionService) UpdateUserCompanyRestriction(usercompanyrestriction model.UserCompanyRestriction, id uint) (usercompanyrestrictionOutput model.UserCompanyRestriction, err error) {
	newUserCompanyRestriction := model.UserCompanyRestriction{}
	err1 := smapping.FillStruct(&newUserCompanyRestriction, smapping.MapFields(&usercompanyrestriction))
	if err != nil {
		return newUserCompanyRestriction, err1
	}
	res, err := service.usercompanyrestrictionRepository.UpdateUserCompanyRestriction(newUserCompanyRestriction, id)
	return res, err
}

func (service *usercompanyrestrictionService) DeleteUserCompanyRestriction(usercompanyrestriction model.UserCompanyRestriction, id uint) (usercompanyrestrictionOutput model.UserCompanyRestriction, err error) {
	newUserCompanyRestriction := model.UserCompanyRestriction{}
	err1 := smapping.FillStruct(&newUserCompanyRestriction, smapping.MapFields(&usercompanyrestriction))
	if err != nil {
		return newUserCompanyRestriction, err1
	}
	res, err := service.usercompanyrestrictionRepository.UpdateUserCompanyRestriction(newUserCompanyRestriction, id)
	return res, err
}
