package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyShareholderService interface {
	FindCompanyShareholders() (companyShareholderOutput []model.CompanyShareholder, err error)
	FindCompanyShareholderByCompanyId(id uint) (companyShareholderOutput []model.CompanyShareholder, err error)
	FindCompanyShareholderById(id uint) (companyShareholderOutput model.CompanyShareholder, err error)
	FindExcCompanyShareholder(id uint) (companyShareholderOutput []model.CompanyShareholder, err error)
	InsertCompanyShareholder(companyShareholder model.CreateCompanyShareholderParameter) (companyShareholderOutput model.CompanyShareholder, err error)
	UpdateCompanyShareholder(companyShareholder model.CompanyShareholder, id uint) (companyShareholderOutput model.CompanyShareholder, err error)
	DeleteCompanyShareholder(companyShareholder model.CompanyShareholder, id uint) (companyShareholderOutput model.CompanyShareholder, err error)
}

type companyShareholderService struct {
	companyShareholderRepository repository.CompanyShareholderRepository
}

func NewCompanyShareholderService(companyShareholderRep repository.CompanyShareholderRepository) CompanyShareholderService {
	return &companyShareholderService{
		companyShareholderRepository: companyShareholderRep,
	}
}

func (service *companyShareholderService) FindCompanyShareholders() (companyShareholderOutput []model.CompanyShareholder, err error) {
	res, err := service.companyShareholderRepository.FindCompanyShareholders()
	return res, err
}

func (service *companyShareholderService) FindCompanyShareholderByCompanyId(id uint) (companyShareholderOutput []model.CompanyShareholder, err error) {
	res, err := service.companyShareholderRepository.FindCompanyShareholderByCompanyId(id)
	return res, err
}

func (service *companyShareholderService) FindCompanyShareholderById(id uint) (companyShareholderOutput model.CompanyShareholder, err error) {
	res, err := service.companyShareholderRepository.FindCompanyShareholderById(id)
	return res, err
}

func (service *companyShareholderService) FindExcCompanyShareholder(id uint) (companyShareholderOutput []model.CompanyShareholder, err error) {
	res, err := service.companyShareholderRepository.FindExcCompanyShareholder(id)
	return res, err
}

func (service *companyShareholderService) InsertCompanyShareholder(companyShareholder model.CreateCompanyShareholderParameter) (companyShareholderOutput model.CompanyShareholder, err error) {
	newCompanyShareholder := model.CompanyShareholder{}
	err1 := smapping.FillStruct(&newCompanyShareholder, smapping.MapFields(&companyShareholder))
	if err != nil {
		return newCompanyShareholder, err1
	}
	res, err := service.companyShareholderRepository.InsertCompanyShareholder(newCompanyShareholder)
	return res, err
}

func (service *companyShareholderService) UpdateCompanyShareholder(companyShareholder model.CompanyShareholder, id uint) (companyShareholderOutput model.CompanyShareholder, err error) {
	newCompanyShareholder := model.CompanyShareholder{}
	err1 := smapping.FillStruct(&newCompanyShareholder, smapping.MapFields(&companyShareholder))
	if err != nil {
		return newCompanyShareholder, err1
	}
	res, err := service.companyShareholderRepository.UpdateCompanyShareholder(newCompanyShareholder, id)
	return res, err
}

func (service *companyShareholderService) DeleteCompanyShareholder(companyShareholder model.CompanyShareholder, id uint) (companyShareholderOutput model.CompanyShareholder, err error) {
	newCompanyShareholder := model.CompanyShareholder{}
	err1 := smapping.FillStruct(&newCompanyShareholder, smapping.MapFields(&companyShareholder))
	if err != nil {
		return newCompanyShareholder, err1
	}
	res, err := service.companyShareholderRepository.UpdateCompanyShareholder(newCompanyShareholder, id)
	return res, err
}
