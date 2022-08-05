package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type RoleService interface {
	CountRoleAll() (count int64, err error)
	FindRoles() (roleOutput []model.Role, err error)
	FindRolesOffset(limit int, offset int, order string, dir string) (roleOutput []model.Role, err error)
	SearchRole(limit int, offset int, order string, dir string, search string) (roleOutput []model.Role, err error)
	CountSearchRole(search string) (count int64, err error)
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

func (service *roleService) CountRoleAll() (count int64, err error) {
	res, err := service.roleRepository.CountRoleAll()
	return res, err
}

func (service *roleService) FindRoles() (roleOutput []model.Role, err error) {
	res, err := service.roleRepository.FindRoles()
	return res, err
}

func (service *roleService) FindRolesOffset(limit int, offset int, order string, dir string) (roleOutput []model.Role, err error) {
	res, err := service.roleRepository.FindRolesOffset(limit, offset, order, dir)
	return res, err
}

func (service *roleService) SearchRole(limit int, offset int, order string, dir string, search string) (roleOutput []model.Role, err error) {
	res, err := service.roleRepository.SearchRole(limit, offset, order, dir, search)
	return res, err
}

func (service *roleService) CountSearchRole(search string) (count int64, err error) {
	res, err := service.roleRepository.CountSearchRole(search)
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
