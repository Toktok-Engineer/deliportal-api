package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyManagementService interface {
	CountCompanyManagementAll(companyId int) (count int64, err error)
	FindCompanyManagements(companyId int) (companyManagementOutput []model.CompanyManagement, err error)
	FindCompanyManagementsOffset(limit int, offset int, order string, dir string, companyId int) (companyManagementOutput []model.SelectCompanyManagementParameter, err error)
	SearchCompanyManagement(limit int, offset int, order string, dir string, search string, companyId int) (companyManagementOutput []model.SelectCompanyManagementParameter, err error)
	CountSearchCompanyManagement(search string, companyId int) (count int64, err error)
	FindCompanyManagementById(id uint) (companyManagementOutput model.SelectCompanyManagementParameter, err error)
	FindExcCompanyManagement(id uint) (companyManagementOutput []model.SelectCompanyManagementParameter, err error)
	FindCompanyManagementByCompanyId(id uint) (companyManagementOutput []model.SelectCompanyManagementParameter, err error)
	InsertCompanyManagement(companyManagement model.CreateCompanyManagementParameter) (companyManagementOutput model.CompanyManagement, err error)
	UpdateCompanyManagement(companyManagement model.CompanyManagement, id uint) (companyManagementOutput model.CompanyManagement, err error)
	DeleteCompanyManagement(companyManagement model.CompanyManagement, id uint) (companyManagementOutput model.CompanyManagement, err error)
}

type companyManagementService struct {
	companyManagementRepository repository.CompanyManagementRepository
}

func NewCompanyManagementService(companyManagementRep repository.CompanyManagementRepository) CompanyManagementService {
	return &companyManagementService{
		companyManagementRepository: companyManagementRep,
	}
}

func (service *companyManagementService) CountCompanyManagementAll(companyId int) (count int64, err error) {
	res, err := service.companyManagementRepository.CountCompanyManagementAll(companyId)
	return res, err
}

func (service *companyManagementService) FindCompanyManagements(companyId int) (companyManagementOutput []model.CompanyManagement, err error) {
	res, err := service.companyManagementRepository.FindCompanyManagements(companyId)
	return res, err
}

func (service *companyManagementService) FindCompanyManagementsOffset(limit int, offset int, order string, dir string, companyId int) (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	res, err := service.companyManagementRepository.FindCompanyManagementsOffset(limit, offset, order, dir, companyId)
	return res, err
}

func (service *companyManagementService) SearchCompanyManagement(limit int, offset int, order string, dir string, search string, companyId int) (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	res, err := service.companyManagementRepository.SearchCompanyManagement(limit, offset, order, dir, search, companyId)
	return res, err
}

func (service *companyManagementService) CountSearchCompanyManagement(search string, companyId int) (count int64, err error) {
	res, err := service.companyManagementRepository.CountSearchCompanyManagement(search, companyId)
	return res, err
}

func (service *companyManagementService) FindCompanyManagementById(id uint) (companyManagementOutput model.SelectCompanyManagementParameter, err error) {
	res, err := service.companyManagementRepository.FindCompanyManagementById(id)
	return res, err
}

func (service *companyManagementService) FindExcCompanyManagement(id uint) (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	res, err := service.companyManagementRepository.FindExcCompanyManagement(id)
	return res, err
}

func (service *companyManagementService) FindCompanyManagementByCompanyId(id uint) (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	res, err := service.companyManagementRepository.FindCompanyManagementByCompanyId(id)
	return res, err
}

func (service *companyManagementService) InsertCompanyManagement(companyManagement model.CreateCompanyManagementParameter) (companyManagementOutput model.CompanyManagement, err error) {
	newCompanyManagement := model.CompanyManagement{}
	err1 := smapping.FillStruct(&newCompanyManagement, smapping.MapFields(&companyManagement))
	if err != nil {
		return newCompanyManagement, err1
	}
	res, err := service.companyManagementRepository.InsertCompanyManagement(newCompanyManagement)
	return res, err
}

func (service *companyManagementService) UpdateCompanyManagement(companyManagement model.CompanyManagement, id uint) (companyManagementOutput model.CompanyManagement, err error) {
	newCompanyManagement := model.CompanyManagement{}
	err1 := smapping.FillStruct(&newCompanyManagement, smapping.MapFields(&companyManagement))
	if err != nil {
		return newCompanyManagement, err1
	}
	res, err := service.companyManagementRepository.UpdateCompanyManagement(newCompanyManagement, id)
	return res, err
}

func (service *companyManagementService) DeleteCompanyManagement(companyManagement model.CompanyManagement, id uint) (companyManagementOutput model.CompanyManagement, err error) {
	newCompanyManagement := model.CompanyManagement{}
	err1 := smapping.FillStruct(&newCompanyManagement, smapping.MapFields(&companyManagement))
	if err != nil {
		return newCompanyManagement, err1
	}
	res, err := service.companyManagementRepository.UpdateCompanyManagement(newCompanyManagement, id)
	return res, err
}
