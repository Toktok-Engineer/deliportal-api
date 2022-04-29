package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyService interface {
	FindCompanys() (companyOutput []model.SelectCompanyParameter, err error)
	FindCompanyById(id uint) (companyOutput model.SelectCompanyParameter, err error)
	FindExcCompany(id uint) (companyOutput []model.SelectCompanyParameter, err error)
	InsertCompany(company model.CreateCompanyParameter) (companyOutput model.Company, err error)
	UpdateCompany(company model.Company, id uint) (companyOutput model.Company, err error)
	UpdateCompanyApprove(company model.Company, id uint) (companyOutput model.Company, err error)
	UpdateCompanyDeactive(company model.Company, id uint) (companyOutput model.Company, err error)
	DeleteCompany(company model.Company, id uint) (companyOutput model.Company, err error)
}

type companyService struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyService(companyRep repository.CompanyRepository) CompanyService {
	return &companyService{
		companyRepository: companyRep,
	}
}

func (service *companyService) FindCompanys() (companyOutput []model.SelectCompanyParameter, err error) {
	res, err := service.companyRepository.FindCompanys()
	return res, err
}

func (service *companyService) FindCompanyById(id uint) (companyOutput model.SelectCompanyParameter, err error) {
	res, err := service.companyRepository.FindCompanyById(id)
	return res, err
}

func (service *companyService) FindExcCompany(id uint) (companyOutput []model.SelectCompanyParameter, err error) {
	res, err := service.companyRepository.FindExcCompany(id)
	return res, err
}

func (service *companyService) InsertCompany(company model.CreateCompanyParameter) (companyOutput model.Company, err error) {
	newCompany := model.Company{}
	err1 := smapping.FillStruct(&newCompany, smapping.MapFields(&company))
	if err != nil {
		return newCompany, err1
	}
	res, err := service.companyRepository.InsertCompany(newCompany)
	return res, err
}

func (service *companyService) UpdateCompany(company model.Company, id uint) (companyOutput model.Company, err error) {
	newCompany := model.Company{}
	err1 := smapping.FillStruct(&newCompany, smapping.MapFields(&company))
	if err != nil {
		return newCompany, err1
	}
	res, err := service.companyRepository.UpdateCompany(newCompany, id)
	return res, err
}

func (service *companyService) UpdateCompanyApprove(company model.Company, id uint) (companyOutput model.Company, err error) {
	newCompany := model.Company{}
	err1 := smapping.FillStruct(&newCompany, smapping.MapFields(&company))
	if err != nil {
		return newCompany, err1
	}
	res, err := service.companyRepository.UpdateCompanyApprove(newCompany, id)
	return res, err
}

func (service *companyService) UpdateCompanyDeactive(company model.Company, id uint) (companyOutput model.Company, err error) {
	newCompany := model.Company{}
	err1 := smapping.FillStruct(&newCompany, smapping.MapFields(&company))
	if err != nil {
		return newCompany, err1
	}
	res, err := service.companyRepository.UpdateCompanyDeactive(newCompany, id)
	return res, err
}

func (service *companyService) DeleteCompany(company model.Company, id uint) (companyOutput model.Company, err error) {
	newCompany := model.Company{}
	err1 := smapping.FillStruct(&newCompany, smapping.MapFields(&company))
	if err != nil {
		return newCompany, err1
	}
	res, err := service.companyRepository.UpdateCompany(newCompany, id)
	return res, err
}
