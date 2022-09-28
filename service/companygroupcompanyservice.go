package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyGroupCompanyService interface {
	CountCompanyGroupCompanyAll(companyGroupID int) (count int64, err error)
	FindCompanyGroupCompanysByCompanyID(companyID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error)
	FindCompanyGroupCompanys(companyGroupID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error)
	FindCompanyGroupCompanysOffset(limit int, offset int, order string, dir string, companyGroupID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error)
	SearchCompanyGroupCompany(limit int, offset int, order string, dir string, search string, companyGroupID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error)
	CountSearchCompanyGroupCompany(search string, companyGroupID int) (count int64, err error)
	FindCompanyGroupCompanyById(id uint) (companyGroupCompanyOutput model.SelectCompanyGroupCompanyParameter, err error)
	InsertCompanyGroupCompany(companyGroupCompany model.CreateCompanyGroupCompanyParameter) (companyGroupCompanyOutput model.CompanyGroupCompany, err error)
	UpdateCompanyGroupCompany(companyGroupCompany model.CompanyGroupCompany, id uint) (companyGroupCompanyOutput model.CompanyGroupCompany, err error)
	DeleteCompanyGroupCompany(companyGroupCompany model.CompanyGroupCompany, id uint) (companyGroupCompanyOutput model.CompanyGroupCompany, err error)
}

type companyGroupCompanyService struct {
	companyGroupCompanyRepository repository.CompanyGroupCompanyRepository
}

func NewCompanyGroupCompanyService(companyGroupCompanyRep repository.CompanyGroupCompanyRepository) CompanyGroupCompanyService {
	return &companyGroupCompanyService{
		companyGroupCompanyRepository: companyGroupCompanyRep,
	}
}

func (service *companyGroupCompanyService) CountCompanyGroupCompanyAll(companyGroupID int) (count int64, err error) {
	res, err := service.companyGroupCompanyRepository.CountCompanyGroupCompanyAll(companyGroupID)
	return res, err
}

func (service *companyGroupCompanyService) FindCompanyGroupCompanysByCompanyID(companyID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error) {
	res, err := service.companyGroupCompanyRepository.FindCompanyGroupCompanysByCompanyID(companyID)
	return res, err
}

func (service *companyGroupCompanyService) FindCompanyGroupCompanys(companyGroupID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error) {
	res, err := service.companyGroupCompanyRepository.FindCompanyGroupCompanys(companyGroupID)
	return res, err
}

func (service *companyGroupCompanyService) FindCompanyGroupCompanysOffset(limit int, offset int, order string, dir string, companyGroupID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error) {
	res, err := service.companyGroupCompanyRepository.FindCompanyGroupCompanysOffset(limit, offset, order, dir, companyGroupID)
	return res, err
}

func (service *companyGroupCompanyService) SearchCompanyGroupCompany(limit int, offset int, order string, dir string, search string, companyGroupID int) (companyGroupCompanyOutput []model.SelectCompanyGroupCompanyParameter, err error) {
	res, err := service.companyGroupCompanyRepository.SearchCompanyGroupCompany(limit, offset, order, dir, search, companyGroupID)
	return res, err
}

func (service *companyGroupCompanyService) CountSearchCompanyGroupCompany(search string, companyGroupID int) (count int64, err error) {
	res, err := service.companyGroupCompanyRepository.CountSearchCompanyGroupCompany(search, companyGroupID)
	return res, err
}

func (service *companyGroupCompanyService) FindCompanyGroupCompanyById(id uint) (companyGroupCompanyOutput model.SelectCompanyGroupCompanyParameter, err error) {
	res, err := service.companyGroupCompanyRepository.FindCompanyGroupCompanyById(id)
	return res, err
}

func (service *companyGroupCompanyService) InsertCompanyGroupCompany(companyGroupCompany model.CreateCompanyGroupCompanyParameter) (companyGroupCompanyOutput model.CompanyGroupCompany, err error) {
	newCompanyGroupCompany := model.CompanyGroupCompany{}
	err1 := smapping.FillStruct(&newCompanyGroupCompany, smapping.MapFields(&companyGroupCompany))
	if err != nil {
		return newCompanyGroupCompany, err1
	}
	res, err := service.companyGroupCompanyRepository.InsertCompanyGroupCompany(newCompanyGroupCompany)
	return res, err
}

func (service *companyGroupCompanyService) UpdateCompanyGroupCompany(companyGroupCompany model.CompanyGroupCompany, id uint) (companyGroupCompanyOutput model.CompanyGroupCompany, err error) {
	newCompanyGroupCompany := model.CompanyGroupCompany{}
	err1 := smapping.FillStruct(&newCompanyGroupCompany, smapping.MapFields(&companyGroupCompany))
	if err != nil {
		return newCompanyGroupCompany, err1
	}
	res, err := service.companyGroupCompanyRepository.UpdateCompanyGroupCompany(newCompanyGroupCompany, id)
	return res, err
}

func (service *companyGroupCompanyService) DeleteCompanyGroupCompany(companyGroupCompany model.CompanyGroupCompany, id uint) (companyGroupCompanyOutput model.CompanyGroupCompany, err error) {
	newCompanyGroupCompany := model.CompanyGroupCompany{}
	err1 := smapping.FillStruct(&newCompanyGroupCompany, smapping.MapFields(&companyGroupCompany))
	if err != nil {
		return newCompanyGroupCompany, err1
	}
	res, err := service.companyGroupCompanyRepository.UpdateCompanyGroupCompany(newCompanyGroupCompany, id)
	return res, err
}
