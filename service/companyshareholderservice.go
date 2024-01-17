package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyShareholderService interface {
	CountCompanyShareholderAll(companyId int) (count int64, err error)
	FindCompanyShareholders(companyId int) (companyShareholderOutput []model.CompanyShareholder, err error)
	FindCompanyShareholdersOffset(limit int, offset int, order string, dir string, companyId int) (companyShareholderOutput []model.SelectCompanyShareholder, err error)
	SearchCompanyShareholder(limit int, offset int, order string, dir string, search string, companyId int) (companyShareholderOutput []model.SelectCompanyShareholder, err error)
	CountSearchCompanyShareholder(search string, companyId int) (count int64, err error)
	FindCompanyShareholderByCompanyId(id uint) (companyShareholderOutput []model.CompanyShareholder, err error)
	FindCompanyShareholderById(id uint) (companyShareholderOutput model.CompanyShareholder, err error)
	FindExcCompanyShareholder(id uint) (companyShareholderOutput []model.CompanyShareholder, err error)
	InsertCompanyShareholder(companyShareholder model.CreateCompanyShareholderParameter) (companyShareholderOutput model.CompanyShareholder, err error)
	UpdateCompanyShareholder(companyShareholder model.CompanyShareholder, id uint) (companyShareholderOutput model.CompanyShareholder, err error)
	DeleteCompanyShareholder(companyShareholder model.CompanyShareholder, id uint) (companyShareholderOutput model.CompanyShareholder, err error)
	ReportCompanyShareholder(year int, groupId int) (companyShareholderOutput []model.SelectCompanyShareholderReport, err error)
}

type companyShareholderService struct {
	companyShareholderRepository repository.CompanyShareholderRepository
}

func NewCompanyShareholderService(companyShareholderRep repository.CompanyShareholderRepository) CompanyShareholderService {
	return &companyShareholderService{
		companyShareholderRepository: companyShareholderRep,
	}
}

func (service *companyShareholderService) CountCompanyShareholderAll(companyId int) (count int64, err error) {
	res, err := service.companyShareholderRepository.CountCompanyShareholderAll(companyId)
	return res, err
}

func (service *companyShareholderService) FindCompanyShareholders(companyId int) (companyShareholderOutput []model.CompanyShareholder, err error) {
	res, err := service.companyShareholderRepository.FindCompanyShareholders(companyId)
	return res, err
}

func (service *companyShareholderService) FindCompanyShareholdersOffset(limit int, offset int, order string, dir string, companyId int) (companyShareholderOutput []model.SelectCompanyShareholder, err error) {
	res, err := service.companyShareholderRepository.FindCompanyShareholdersOffset(limit, offset, order, dir, companyId)
	return res, err
}

func (service *companyShareholderService) SearchCompanyShareholder(limit int, offset int, order string, dir string, search string, companyId int) (companyShareholderOutput []model.SelectCompanyShareholder, err error) {
	res, err := service.companyShareholderRepository.SearchCompanyShareholder(limit, offset, order, dir, search, companyId)
	return res, err
}

func (service *companyShareholderService) CountSearchCompanyShareholder(search string, companyId int) (count int64, err error) {
	res, err := service.companyShareholderRepository.CountSearchCompanyShareholder(search, companyId)
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

func (service *companyShareholderService) ReportCompanyShareholder(year int, groupId int) (companyShareholderOutput []model.SelectCompanyShareholderReport, err error) {
	res, err := service.companyShareholderRepository.ReportCompanyShareholder(year, groupId)
	return res, err
}
