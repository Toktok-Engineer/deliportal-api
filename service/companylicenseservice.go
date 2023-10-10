package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyLicenseService interface {
	FindCompanyLicenseByCompanyId(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	CountCompanyLicenseAll(companyId int, licenseGroup uint) (count int64, err error)
	FindCompanyLicenses() (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	FindCompanyLicensesOffset(limit int, offset int, order string, dir string, companyId int, licenseGroup uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	SearchCompanyLicense(limit int, offset int, order string, dir string, search string, companyId int, licenseGroup uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	CountSearchCompanyLicense(search string, companyId int, licenseGroup uint) (count int64, err error)
	CountCompanyLicenseAllMobile(companyId int) (count int64, err error)
	FindCompanyLicensesOffsetMobile(limit int, offset int, order string, dir string, companyId int) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	SearchCompanyLicenseMobile(limit int, offset int, order string, dir string, search string, companyId int) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	CountSearchCompanyLicenseMobile(search string, companyId int) (count int64, err error)
	CountCompanyLicenseApp(companyID string) (count int64, err error)
	FindCompanyLicensesApp(limit int, offset int, order string, dir string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	SearchCompanyLicenseApp(limit int, offset int, order string, dir string, search string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	CountSearchCompanyLicenseApp(search string, companyID string) (count int64, err error)
	CountExpCompanyLicense(companyID string) (count int64, err error)
	FindExpCompanyLicenses(limit int, offset int, order string, dir string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	CountSearchExpCompanyLicense(search string, companyID string) (count int64, err error)
	SearchExpCompanyLicense(limit int, offset int, order string, dir string, search string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	FindCompanyLicenseById(id uint) (companyLicenseOutput model.SelectCompanyLicenseParameter, err error)
	FindExcCompanyLicense(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	CountCompanyLicenseFull(companyID string) (count int64, err error)
	FindCompanyLicensesOffsetFull(limit int, offset int, order string, dir string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	SearchCompanyLicenseFull(limit int, offset int, order string, dir string, search string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error)
	CountSearchCompanyLicenseFull(search string, companyID string) (count int64, err error)
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

func (service *companyLicenseService) FindCompanyLicenseByCompanyId(id uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindCompanyLicenseByCompanyId(id)
	return res, err
}

func (service *companyLicenseService) CountCompanyLicenseAll(companyId int, licenseGroup uint) (count int64, err error) {
	res, err := service.companyLicenseRepository.CountCompanyLicenseAll(companyId, licenseGroup)
	return res, err
}

func (service *companyLicenseService) FindCompanyLicenses() (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindCompanyLicenses()
	return res, err
}

func (service *companyLicenseService) FindCompanyLicensesOffset(limit int, offset int, order string, dir string, companyId int, licenseGroup uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindCompanyLicensesOffset(limit, offset, order, dir, companyId, licenseGroup)
	return res, err
}

func (service *companyLicenseService) SearchCompanyLicense(limit int, offset int, order string, dir string, search string, companyId int, licenseGroup uint) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.SearchCompanyLicense(limit, offset, order, dir, search, companyId, licenseGroup)
	return res, err
}

func (service *companyLicenseService) CountSearchCompanyLicense(search string, companyId int, licenseGroup uint) (count int64, err error) {
	res, err := service.companyLicenseRepository.CountSearchCompanyLicense(search, companyId, licenseGroup)
	return res, err
}

func (service *companyLicenseService) CountCompanyLicenseAllMobile(companyId int) (count int64, err error) {
	res, err := service.companyLicenseRepository.CountCompanyLicenseAllMobile(companyId)
	return res, err
}

func (service *companyLicenseService) FindCompanyLicensesOffsetMobile(limit int, offset int, order string, dir string, companyId int) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindCompanyLicensesOffsetMobile(limit, offset, order, dir, companyId)
	return res, err
}

func (service *companyLicenseService) SearchCompanyLicenseMobile(limit int, offset int, order string, dir string, search string, companyId int) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.SearchCompanyLicenseMobile(limit, offset, order, dir, search, companyId)
	return res, err
}

func (service *companyLicenseService) CountSearchCompanyLicenseMobile(search string, companyId int) (count int64, err error) {
	res, err := service.companyLicenseRepository.CountSearchCompanyLicenseMobile(search, companyId)
	return res, err
}

func (service *companyLicenseService) CountCompanyLicenseApp(companyID string) (count int64, err error) {
	res, err := service.companyLicenseRepository.CountCompanyLicenseApp(companyID)
	return res, err
}

func (service *companyLicenseService) FindCompanyLicensesApp(limit int, offset int, order string, dir string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindCompanyLicensesApp(limit, offset, order, dir, companyID)
	return res, err
}

func (service *companyLicenseService) SearchCompanyLicenseApp(limit int, offset int, order string, dir string, search string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.SearchCompanyLicenseApp(limit, offset, order, dir, search, companyID)
	return res, err
}

func (service *companyLicenseService) CountSearchCompanyLicenseApp(search string, companyID string) (count int64, err error) {
	res, err := service.companyLicenseRepository.CountSearchCompanyLicenseApp(search, companyID)
	return res, err
}

func (service *companyLicenseService) CountExpCompanyLicense(companyID string) (count int64, err error) {
	res, err := service.companyLicenseRepository.CountExpCompanyLicense(companyID)
	return res, err
}

func (service *companyLicenseService) FindExpCompanyLicenses(limit int, offset int, order string, dir string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindExpCompanyLicenses(limit, offset, order, dir, companyID)
	return res, err
}

func (service *companyLicenseService) SearchExpCompanyLicense(limit int, offset int, order string, dir string, search string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.SearchExpCompanyLicense(limit, offset, order, dir, search, companyID)
	return res, err
}

func (service *companyLicenseService) CountSearchExpCompanyLicense(search string, companyID string) (count int64, err error) {
	res, err := service.companyLicenseRepository.CountSearchExpCompanyLicense(search, companyID)
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

func (service *companyLicenseService) CountCompanyLicenseFull(companyID string) (count int64, err error) {
	res, err := service.companyLicenseRepository.CountCompanyLicenseFull(companyID)
	return res, err
}

func (service *companyLicenseService) FindCompanyLicensesOffsetFull(limit int, offset int, order string, dir string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.FindCompanyLicensesOffsetFull(limit, offset, order, dir, companyID)
	return res, err
}

func (service *companyLicenseService) SearchCompanyLicenseFull(limit int, offset int, order string, dir string, search string, companyID string) (companyLicenseOutput []model.SelectCompanyLicenseParameter, err error) {
	res, err := service.companyLicenseRepository.SearchCompanyLicenseFull(limit, offset, order, dir, search, companyID)
	return res, err
}

func (service *companyLicenseService) CountSearchCompanyLicenseFull(search string, companyID string) (count int64, err error) {
	res, err := service.companyLicenseRepository.CountSearchCompanyLicenseFull(search, companyID)
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
