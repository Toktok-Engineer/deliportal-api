package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyGroupService interface {
	CountCompanyGroupAll() (count int64, err error)
	FindCompanyGroups() (businessunitOutput []model.CompanyGroup, err error)
	FindCompanyGroupsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.CompanyGroup, err error)
	SearchCompanyGroup(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.CompanyGroup, err error)
	CountSearchCompanyGroup(search string) (count int64, err error)
	FindCompanyGroupById(id uint) (companyGroupOutput model.CompanyGroup, err error)
	FindExcCompanyGroup(id uint) (companyGroupOutput []model.CompanyGroup, err error)
	InsertCompanyGroup(companyGroup model.CreateCompanyGroupParameter) (companyGroupOutput model.CompanyGroup, err error)
	UpdateCompanyGroup(companyGroup model.CompanyGroup, id uint) (companyGroupOutput model.CompanyGroup, err error)
	DeleteCompanyGroup(companyGroup model.CompanyGroup, id uint) (companyGroupOutput model.CompanyGroup, err error)
}

type companyGroupService struct {
	companyGroupRepository repository.CompanyGroupRepository
}

func NewCompanyGroupService(companyGroupRep repository.CompanyGroupRepository) CompanyGroupService {
	return &companyGroupService{
		companyGroupRepository: companyGroupRep,
	}
}

func (service *companyGroupService) CountCompanyGroupAll() (count int64, err error) {
	res, err := service.companyGroupRepository.CountCompanyGroupAll()
	return res, err
}

func (service *companyGroupService) FindCompanyGroups() (businessunitOutput []model.CompanyGroup, err error) {
	res, err := service.companyGroupRepository.FindCompanyGroups()
	return res, err
}

func (service *companyGroupService) FindCompanyGroupsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.CompanyGroup, err error) {
	res, err := service.companyGroupRepository.FindCompanyGroupsOffset(limit, offset, order, dir)
	return res, err
}

func (service *companyGroupService) SearchCompanyGroup(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.CompanyGroup, err error) {
	res, err := service.companyGroupRepository.SearchCompanyGroup(limit, offset, order, dir, search)
	return res, err
}

func (service *companyGroupService) CountSearchCompanyGroup(search string) (count int64, err error) {
	res, err := service.companyGroupRepository.CountSearchCompanyGroup(search)
	return res, err
}

func (service *companyGroupService) FindCompanyGroupById(id uint) (companyGroupOutput model.CompanyGroup, err error) {
	return service.companyGroupRepository.FindCompanyGroupById(id)
}

func (service *companyGroupService) FindExcCompanyGroup(id uint) (companyGroupOutput []model.CompanyGroup, err error) {
	return service.companyGroupRepository.FindExcCompanyGroup(id)
}
func (service *companyGroupService) InsertCompanyGroup(companyGroup model.CreateCompanyGroupParameter) (model.CompanyGroup, error) {
	newCompanyGroup := model.CompanyGroup{}
	err1 := smapping.FillStruct(&newCompanyGroup, smapping.MapFields(&companyGroup))

	if err1 != nil {
		return newCompanyGroup, err1
	}

	return service.companyGroupRepository.InsertCompanyGroup(newCompanyGroup)
}

func (service *companyGroupService) UpdateCompanyGroup(companyGroup model.CompanyGroup, id uint) (companyGroupOutput model.CompanyGroup, err error) {
	newCompanyGroup := model.CompanyGroup{}
	err1 := smapping.FillStruct(&newCompanyGroup, smapping.MapFields(&companyGroup))

	if err1 != nil {
		return newCompanyGroup, err1
	}

	return service.companyGroupRepository.UpdateCompanyGroup(newCompanyGroup, id)
}

func (service *companyGroupService) DeleteCompanyGroup(companyGroup model.CompanyGroup, id uint) (companyGroupOutput model.CompanyGroup, err error) {
	newCompanyGroup := model.CompanyGroup{}
	err1 := smapping.FillStruct(&newCompanyGroup, smapping.MapFields(&companyGroup))

	if err1 != nil {
		return newCompanyGroup, err1
	}

	return service.companyGroupRepository.UpdateCompanyGroup(newCompanyGroup, id)
}
