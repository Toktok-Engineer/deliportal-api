package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyLicenseRenewalTracingService interface {
	FindCompanyLicenseRenewalTracings() (companyLicenseRenewalTracingOutput []model.CompanyLicenseRenewalTracing, err error)
	FindCompanyLicenseRenewalTracingById(id uint) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error)
	FindExcCompanyLicenseRenewalTracing(id uint) (companyLicenseRenewalTracingOutput []model.CompanyLicenseRenewalTracing, err error)
	InsertCompanyLicenseRenewalTracing(companyLicenseRenewalTracing model.CreateCompanyLicenseRenewalTracingParameter) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error)
	UpdateCompanyLicenseRenewalTracing(companyLicenseRenewalTracing model.CompanyLicenseRenewalTracing, id uint) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error)
	DeleteCompanyLicenseRenewalTracing(companyLicenseRenewalTracing model.CompanyLicenseRenewalTracing, id uint) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error)
}

type companyLicenseRenewalTracingService struct {
	companyLicenseRenewalTracingRepository repository.CompanyLicenseRenewalTracingRepository
}

func NewCompanyLicenseRenewalTracingService(companyLicenseRenewalTracingRep repository.CompanyLicenseRenewalTracingRepository) CompanyLicenseRenewalTracingService {
	return &companyLicenseRenewalTracingService{
		companyLicenseRenewalTracingRepository: companyLicenseRenewalTracingRep,
	}
}

func (service *companyLicenseRenewalTracingService) FindCompanyLicenseRenewalTracings() (companyLicenseRenewalTracingOutput []model.CompanyLicenseRenewalTracing, err error) {
	res, err := service.companyLicenseRenewalTracingRepository.FindCompanyLicenseRenewalTracings()
	return res, err
}

func (service *companyLicenseRenewalTracingService) FindCompanyLicenseRenewalTracingById(id uint) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error) {
	res, err := service.companyLicenseRenewalTracingRepository.FindCompanyLicenseRenewalTracingById(id)
	return res, err
}

func (service *companyLicenseRenewalTracingService) FindExcCompanyLicenseRenewalTracing(id uint) (companyLicenseRenewalTracingOutput []model.CompanyLicenseRenewalTracing, err error) {
	res, err := service.companyLicenseRenewalTracingRepository.FindExcCompanyLicenseRenewalTracing(id)
	return res, err
}

func (service *companyLicenseRenewalTracingService) InsertCompanyLicenseRenewalTracing(companyLicenseRenewalTracing model.CreateCompanyLicenseRenewalTracingParameter) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error) {
	newCompanyLicenseRenewalTracing := model.CompanyLicenseRenewalTracing{}
	err1 := smapping.FillStruct(&newCompanyLicenseRenewalTracing, smapping.MapFields(&companyLicenseRenewalTracing))
	if err != nil {
		return newCompanyLicenseRenewalTracing, err1
	}
	res, err := service.companyLicenseRenewalTracingRepository.InsertCompanyLicenseRenewalTracing(newCompanyLicenseRenewalTracing)
	return res, err
}

func (service *companyLicenseRenewalTracingService) UpdateCompanyLicenseRenewalTracing(companyLicenseRenewalTracing model.CompanyLicenseRenewalTracing, id uint) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error) {
	newCompanyLicenseRenewalTracing := model.CompanyLicenseRenewalTracing{}
	err1 := smapping.FillStruct(&newCompanyLicenseRenewalTracing, smapping.MapFields(&companyLicenseRenewalTracing))
	if err != nil {
		return newCompanyLicenseRenewalTracing, err1
	}
	res, err := service.companyLicenseRenewalTracingRepository.UpdateCompanyLicenseRenewalTracing(newCompanyLicenseRenewalTracing, id)
	return res, err
}

func (service *companyLicenseRenewalTracingService) DeleteCompanyLicenseRenewalTracing(companyLicenseRenewalTracing model.CompanyLicenseRenewalTracing, id uint) (companyLicenseRenewalTracingOutput model.CompanyLicenseRenewalTracing, err error) {
	newCompanyLicenseRenewalTracing := model.CompanyLicenseRenewalTracing{}
	err1 := smapping.FillStruct(&newCompanyLicenseRenewalTracing, smapping.MapFields(&companyLicenseRenewalTracing))
	if err != nil {
		return newCompanyLicenseRenewalTracing, err1
	}
	res, err := service.companyLicenseRenewalTracingRepository.UpdateCompanyLicenseRenewalTracing(newCompanyLicenseRenewalTracing, id)
	return res, err
}
