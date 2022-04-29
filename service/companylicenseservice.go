package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyLicenseService interface {
	FindCompanyLicenses() (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	FindCompanyLicenseById(id uint) (companyLicenseOutput model.SelectCompanyLicenseParameter, err error)
	FindExcCompanyLicense(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	FindCompanyLicenseByCompanyId(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	InsertCompanyLicense(companyLicense model.CreateCompanyLicenseParameter) (companyLicenseOutput model.CompanyLicense, err error)
	UpdateCompanyLicense(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
	UpdateCompanyLicenseStatus(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
	UpdateCompanyLicenseDeactive(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
	UpdateCompanyLicenseApprovedRenewalStatus(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
	DeleteCompanyLicense(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
	UpdateCompanyRemark(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error)
}

type companyLicenseService struct {
	companyLicenseRepository repository.CompanyLicenseRepository
}

func NewCompanyLicenseService(companyLicenseRep repository.CompanyLicenseRepository) CompanyLicenseService {
	return &companyLicenseService{
		companyLicenseRepository: companyLicenseRep,
	}
}

func (service *companyLicenseService) FindCompanyLicenses() (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindCompanyLicenses()
	return res, err
}

func (service *companyLicenseService) FindCompanyLicenseById(id uint) (companyLicenseOutput model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindCompanyLicenseById(id)
	return res, err
}

func (service *companyLicenseService) FindExcCompanyLicense(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindExcCompanyLicense(id)
	return res, err
}

func (service *companyLicenseService) FindCompanyLicenseByCompanyId(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindCompanyLicenseByCompanyId(id)
	return res, err
}

func (service *companyLicenseService) InsertCompanyLicense(companyLicense model.CreateCompanyLicenseParameter) (companyLicenseOutput model.CompanyLicense, err error) {
	newCompanyLicense := model.CompanyLicense{}
	err1 := smapping.FillStruct(&newCompanyLicense, smapping.MapFields(&companyLicense))
	if err != nil {
		return newCompanyLicense, err1
	}
	res, err := service.companyLicenseRepository.InsertCompanyLicense(newCompanyLicense)
	return res, err
}

func (service *companyLicenseService) UpdateCompanyLicense(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	newCompanyLicense := model.CompanyLicense{}
	err1 := smapping.FillStruct(&newCompanyLicense, smapping.MapFields(&companyLicense))
	if err != nil {
		return newCompanyLicense, err1
	}
	res, err := service.companyLicenseRepository.UpdateCompanyLicense(newCompanyLicense, id)
	return res, err
}

func (service *companyLicenseService) UpdateCompanyLicenseStatus(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	newCompanyLicense := model.CompanyLicense{}
	err1 := smapping.FillStruct(&newCompanyLicense, smapping.MapFields(&companyLicense))
	if err != nil {
		return newCompanyLicense, err1
	}
	res, err := service.companyLicenseRepository.UpdateCompanyLicenseStatus(newCompanyLicense, id)
	return res, err
}

func (service *companyLicenseService) UpdateCompanyLicenseApprovedRenewalStatus(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	newCompanyLicense := model.CompanyLicense{}
	err1 := smapping.FillStruct(&newCompanyLicense, smapping.MapFields(&companyLicense))
	if err != nil {
		return newCompanyLicense, err1
	}
	res, err := service.companyLicenseRepository.UpdateCompanyLicenseApprovedRenewalStatus(newCompanyLicense, id)
	return res, err
}

func (service *companyLicenseService) UpdateCompanyLicenseDeactive(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	newCompanyLicense := model.CompanyLicense{}
	err1 := smapping.FillStruct(&newCompanyLicense, smapping.MapFields(&companyLicense))
	if err != nil {
		return newCompanyLicense, err1
	}
	res, err := service.companyLicenseRepository.UpdateCompanyLicenseDeactive(newCompanyLicense, id)
	return res, err
}

func (service *companyLicenseService) DeleteCompanyLicense(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	newCompanyLicense := model.CompanyLicense{}
	err1 := smapping.FillStruct(&newCompanyLicense, smapping.MapFields(&companyLicense))
	if err != nil {
		return newCompanyLicense, err1
	}
	res, err := service.companyLicenseRepository.DeleteCompanyLicense(newCompanyLicense, id)
	return res, err
}

func (service *companyLicenseService) UpdateCompanyRemark(companyLicense model.CompanyLicense, id uint) (companyLicenseOutput model.CompanyLicense, err error) {
	newCompanyLicense := model.CompanyLicense{}
	err1 := smapping.FillStruct(&newCompanyLicense, smapping.MapFields(&companyLicense))
	if err != nil {
		return newCompanyLicense, err1
	}
	res, err := service.companyLicenseRepository.UpdateCompanyRemark(newCompanyLicense, id)
	return res, err
}
