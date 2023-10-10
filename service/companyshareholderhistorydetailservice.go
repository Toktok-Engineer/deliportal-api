package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyShareholderHistoryDetailService interface {
	CountCompanyShareholderHistoryDetailAll(companyShareholderHistoryDetailHistoryId int) (count int64, err error)
	FindCompanyShareholderHistoryDetails(companyShareholderHistoryDetailHistoryId int) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error)
	FindCompanyShareholderHistoryDetailsOffset(limit int, offset int, order string, dir string, companyShareholderHistoryDetailHistoryId int) (companyShareholderHistoryDetailOutput []model.SelectCompanyShareholderHistoryDetail, err error)
	SearchCompanyShareholderHistoryDetail(limit int, offset int, order string, dir string, search string, companyShareholderHistoryDetailHistoryId int) (companyShareholderHistoryDetailOutput []model.SelectCompanyShareholderHistoryDetail, err error)
	CountSearchCompanyShareholderHistoryDetail(search string, companyShareholderHistoryDetailHistoryId int) (count int64, err error)
	FindCompanyShareholderHistoryDetailByCompanyId(id uint) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error)
	FindCompanyShareholderHistoryDetailById(id uint) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error)
	FindExcCompanyShareholderHistoryDetail(id uint) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error)
	InsertCompanyShareholderHistoryDetail(companyShareholderHistoryDetail model.CreateCompanyShareholderHistoryDetailParameter) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error)
	UpdateCompanyShareholderHistoryDetail(companyShareholderHistoryDetail model.CompanyShareholderHistoryDetail, id uint) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error)
	DeleteCompanyShareholderHistoryDetail(companyShareholderHistoryDetail model.CompanyShareholderHistoryDetail, id uint) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error)
}

type companyShareholderHistoryDetailService struct {
	companyShareholderHistoryDetailRepository repository.CompanyShareholderHistoryDetailRepository
}

func NewCompanyShareholderHistoryDetailService(companyShareholderHistoryDetailRep repository.CompanyShareholderHistoryDetailRepository) CompanyShareholderHistoryDetailService {
	return &companyShareholderHistoryDetailService{
		companyShareholderHistoryDetailRepository: companyShareholderHistoryDetailRep,
	}
}

func (service *companyShareholderHistoryDetailService) CountCompanyShareholderHistoryDetailAll(companyShareholderHistoryDetailHistoryId int) (count int64, err error) {
	res, err := service.companyShareholderHistoryDetailRepository.CountCompanyShareholderHistoryDetailAll(companyShareholderHistoryDetailHistoryId)
	return res, err
}

func (service *companyShareholderHistoryDetailService) FindCompanyShareholderHistoryDetails(companyShareholderHistoryDetailHistoryId int) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error) {
	res, err := service.companyShareholderHistoryDetailRepository.FindCompanyShareholderHistoryDetails(companyShareholderHistoryDetailHistoryId)
	return res, err
}

func (service *companyShareholderHistoryDetailService) FindCompanyShareholderHistoryDetailsOffset(limit int, offset int, order string, dir string, companyShareholderHistoryDetailHistoryId int) (companyShareholderHistoryDetailOutput []model.SelectCompanyShareholderHistoryDetail, err error) {
	res, err := service.companyShareholderHistoryDetailRepository.FindCompanyShareholderHistoryDetailsOffset(limit, offset, order, dir, companyShareholderHistoryDetailHistoryId)
	return res, err
}

func (service *companyShareholderHistoryDetailService) SearchCompanyShareholderHistoryDetail(limit int, offset int, order string, dir string, search string, companyShareholderHistoryDetailHistoryId int) (companyShareholderHistoryDetailOutput []model.SelectCompanyShareholderHistoryDetail, err error) {
	res, err := service.companyShareholderHistoryDetailRepository.SearchCompanyShareholderHistoryDetail(limit, offset, order, dir, search, companyShareholderHistoryDetailHistoryId)
	return res, err
}

func (service *companyShareholderHistoryDetailService) CountSearchCompanyShareholderHistoryDetail(search string, companyShareholderHistoryDetailHistoryId int) (count int64, err error) {
	res, err := service.companyShareholderHistoryDetailRepository.CountSearchCompanyShareholderHistoryDetail(search, companyShareholderHistoryDetailHistoryId)
	return res, err
}

func (service *companyShareholderHistoryDetailService) FindCompanyShareholderHistoryDetailByCompanyId(id uint) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error) {
	res, err := service.companyShareholderHistoryDetailRepository.FindCompanyShareholderHistoryDetailByCompanyId(id)
	return res, err
}

func (service *companyShareholderHistoryDetailService) FindCompanyShareholderHistoryDetailById(id uint) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error) {
	res, err := service.companyShareholderHistoryDetailRepository.FindCompanyShareholderHistoryDetailById(id)
	return res, err
}

func (service *companyShareholderHistoryDetailService) FindExcCompanyShareholderHistoryDetail(id uint) (companyShareholderHistoryDetailOutput []model.CompanyShareholderHistoryDetail, err error) {
	res, err := service.companyShareholderHistoryDetailRepository.FindExcCompanyShareholderHistoryDetail(id)
	return res, err
}

func (service *companyShareholderHistoryDetailService) InsertCompanyShareholderHistoryDetail(companyShareholderHistoryDetail model.CreateCompanyShareholderHistoryDetailParameter) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error) {
	newCompanyShareholderHistoryDetail := model.CompanyShareholderHistoryDetail{}
	err1 := smapping.FillStruct(&newCompanyShareholderHistoryDetail, smapping.MapFields(&companyShareholderHistoryDetail))
	if err != nil {
		return newCompanyShareholderHistoryDetail, err1
	}
	res, err := service.companyShareholderHistoryDetailRepository.InsertCompanyShareholderHistoryDetail(newCompanyShareholderHistoryDetail)
	return res, err
}

func (service *companyShareholderHistoryDetailService) UpdateCompanyShareholderHistoryDetail(companyShareholderHistoryDetail model.CompanyShareholderHistoryDetail, id uint) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error) {
	newCompanyShareholderHistoryDetail := model.CompanyShareholderHistoryDetail{}
	err1 := smapping.FillStruct(&newCompanyShareholderHistoryDetail, smapping.MapFields(&companyShareholderHistoryDetail))
	if err != nil {
		return newCompanyShareholderHistoryDetail, err1
	}
	res, err := service.companyShareholderHistoryDetailRepository.UpdateCompanyShareholderHistoryDetail(newCompanyShareholderHistoryDetail, id)
	return res, err
}

func (service *companyShareholderHistoryDetailService) DeleteCompanyShareholderHistoryDetail(companyShareholderHistoryDetail model.CompanyShareholderHistoryDetail, id uint) (companyShareholderHistoryDetailOutput model.CompanyShareholderHistoryDetail, err error) {
	newCompanyShareholderHistoryDetail := model.CompanyShareholderHistoryDetail{}
	err1 := smapping.FillStruct(&newCompanyShareholderHistoryDetail, smapping.MapFields(&companyShareholderHistoryDetail))
	if err != nil {
		return newCompanyShareholderHistoryDetail, err1
	}
	res, err := service.companyShareholderHistoryDetailRepository.UpdateCompanyShareholderHistoryDetail(newCompanyShareholderHistoryDetail, id)
	return res, err
}
