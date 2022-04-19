package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type RoleService interface {
	FindRoles() (roleOutput []model.Role, err error)
	FindRoleById(id uint) (roleOutput model.Role, err error)
	FindExcRole(id uint) (roleOutput []model.Role, err error)
	InsertRole(role model.CreateRoleParameter) (roleOutput model.Role, err error)
	UpdateRole(role model.Role, id uint) (roleOutput model.Role, err error)
	DeleteRole(role model.Role, id uint) (roleOutput model.Role, err error)
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRep repository.RoleRepository) RoleService {
	return &roleService{
		roleRepository: roleRep,
	}
}

func (service *roleService) FindRoles() (roleOutput []model.Role, err error) {
	res, err := service.roleRepository.FindRoles()
	return res, err
}

func (service *roleService) FindRoleById(id uint) (roleOutput model.Role, err error) {
	res, err := service.roleRepository.FindRoleById(id)
	return res, err
}

func (service *roleService) FindExcRole(id uint) (roleOutput []model.Role, err error) {
	res, err := service.roleRepository.FindExcRole(id)
	return res, err
}

func (service *roleService) InsertRole(role model.CreateRoleParameter) (roleOutput model.Role, err error) {
	newRole := model.Role{}
	err1 := smapping.FillStruct(&newRole, smapping.MapFields(&role))
	if err != nil {
		return newRole, err1
	}
	res, err := service.roleRepository.InsertRole(newRole)
	return res, err
}

func (service *roleService) UpdateRole(role model.Role, id uint) (roleOutput model.Role, err error) {
	newRole := model.Role{}
	err1 := smapping.FillStruct(&newRole, smapping.MapFields(&role))
	if err != nil {
		return newRole, err1
	}
	res, err := service.roleRepository.UpdateRole(newRole, id)
	return res, err
}

func (service *roleService) DeleteRole(role model.Role, id uint) (roleOutput model.Role, err error) {
	newRole := model.Role{}
	err1 := smapping.FillStruct(&newRole, smapping.MapFields(&role))
	if err != nil {
		return newRole, err1
	}
	res, err := service.roleRepository.UpdateRole(newRole, id)
	return res, err
}
