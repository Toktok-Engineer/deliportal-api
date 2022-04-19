package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type RoleFormService interface {
	FindRoleForms() (roleFormOutput []model.SelectRoleFormParameter, err error)
	FindRoleFormById(id uint) (roleFormOutput model.SelectRoleFormParameter, err error)
	FindRoleFormByFormId(fid uint, rid uint) (roleFormOutput model.SelectRoleFormParameter, err error)
	FindExcRoleForm(id uint, rid uint) (roleFormOutput []model.SelectRoleFormParameter, err error)
	FindExcRoleFormOnly(id uint) (roleFormOutput []model.SelectRoleFormParameter, err error)
	InsertRoleForm(roleForm model.CreateRoleFormParameter) (roleFormOutput model.RoleForm, err error)
	UpdateRoleForm(roleForm model.RoleForm, id uint) (roleFormOutput model.RoleForm, err error)
	DeleteRoleForm(roleForm model.RoleForm, id uint) (roleFormOutput model.RoleForm, err error)
}

type roleFormService struct {
	roleFormRepository repository.RoleFormRepository
}

func NewRoleFormService(roleFormRep repository.RoleFormRepository) RoleFormService {
	return &roleFormService{
		roleFormRepository: roleFormRep,
	}
}

func (service *roleFormService) FindRoleForms() (roleFormOutput []model.SelectRoleFormParameter, err error) {
	res, err := service.roleFormRepository.FindRoleForms()
	return res, err
}

func (service *roleFormService) FindRoleFormById(id uint) (roleFormOutput model.SelectRoleFormParameter, err error) {
	res, err := service.roleFormRepository.FindRoleFormById(id)
	return res, err
}

func (service *roleFormService) FindRoleFormByFormId(fid uint, rid uint) (roleFormOutput model.SelectRoleFormParameter, err error) {
	res, err := service.roleFormRepository.FindRoleFormByFormId(fid, rid)
	return res, err
}

func (service *roleFormService) FindExcRoleForm(id uint, rid uint) (roleFormOutput []model.SelectRoleFormParameter, err error) {
	res, err := service.roleFormRepository.FindExcRoleForm(id, rid)
	return res, err
}

func (service *roleFormService) FindExcRoleFormOnly(id uint) (roleFormOutput []model.SelectRoleFormParameter, err error) {
	res, err := service.roleFormRepository.FindExcRoleFormOnly(id)
	return res, err
}

func (service *roleFormService) InsertRoleForm(roleForm model.CreateRoleFormParameter) (roleFormOutput model.RoleForm, err error) {
	newRoleForm := model.RoleForm{}
	err1 := smapping.FillStruct(&newRoleForm, smapping.MapFields(&roleForm))
	if err != nil {
		return newRoleForm, err1
	}
	res, err := service.roleFormRepository.InsertRoleForm(newRoleForm)
	return res, err
}

func (service *roleFormService) UpdateRoleForm(roleForm model.RoleForm, id uint) (roleFormOutput model.RoleForm, err error) {
	newRoleForm := model.RoleForm{}
	err1 := smapping.FillStruct(&newRoleForm, smapping.MapFields(&roleForm))
	if err != nil {
		return newRoleForm, err1
	}
	res, err := service.roleFormRepository.UpdateRoleForm(newRoleForm, id)
	return res, err
}

func (service *roleFormService) DeleteRoleForm(roleForm model.RoleForm, id uint) (roleFormOutput model.RoleForm, err error) {
	newRoleForm := model.RoleForm{}
	err1 := smapping.FillStruct(&newRoleForm, smapping.MapFields(&roleForm))
	if err != nil {
		return newRoleForm, err1
	}
	res, err := service.roleFormRepository.UpdateRoleForm(newRoleForm, id)
	return res, err
}
