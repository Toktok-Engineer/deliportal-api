package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyManagementService interface {
	FindCompanyManagements() (companyManagementOutput []model.SelectCompanyManagementParameter, err error)
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

func (service *companyManagementService) FindCompanyManagements() (companyManagementOutput []model.SelectCompanyManagementParameter, err error) {
	res, err := service.companyManagementRepository.FindCompanyManagements()
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
