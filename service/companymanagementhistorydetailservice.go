package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyManagementHistoryDetailService interface {
	CountCompanyManagementHistoryDetailAll(companyManagementHistoryDetailHistoryId int) (count int64, err error)
	FindCompanyManagementHistoryDetails(companyManagementHistoryDetailHistoryId int) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error)
	FindCompanyManagementHistoryDetailsOffset(limit int, offset int, order string, dir string, companyManagementHistoryDetailHistoryId int) (companyManagementHistoryDetailOutput []model.SelectCompanyManagementHistoryDetail, err error)
	SearchCompanyManagementHistoryDetail(limit int, offset int, order string, dir string, search string, companyManagementHistoryDetailHistoryId int) (companyManagementHistoryDetailOutput []model.SelectCompanyManagementHistoryDetail, err error)
	CountSearchCompanyManagementHistoryDetail(search string, companyManagementHistoryDetailHistoryId int) (count int64, err error)
	FindCompanyManagementHistoryDetailByCompanyId(id uint) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error)
	FindCompanyManagementHistoryDetailById(id uint) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error)
	FindExcCompanyManagementHistoryDetail(id uint) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error)
	InsertCompanyManagementHistoryDetail(companyManagementHistoryDetail model.CreateCompanyManagementHistoryDetailParameter) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error)
	UpdateCompanyManagementHistoryDetail(companyManagementHistoryDetail model.CompanyManagementHistoryDetail, id uint) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error)
	DeleteCompanyManagementHistoryDetail(companyManagementHistoryDetail model.CompanyManagementHistoryDetail, id uint) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error)
}

type companyManagementHistoryDetailService struct {
	companyManagementHistoryDetailRepository repository.CompanyManagementHistoryDetailRepository
}

func NewCompanyManagementHistoryDetailService(companyManagementHistoryDetailRep repository.CompanyManagementHistoryDetailRepository) CompanyManagementHistoryDetailService {
	return &companyManagementHistoryDetailService{
		companyManagementHistoryDetailRepository: companyManagementHistoryDetailRep,
	}
}

func (service *companyManagementHistoryDetailService) CountCompanyManagementHistoryDetailAll(companyManagementHistoryDetailHistoryId int) (count int64, err error) {
	res, err := service.companyManagementHistoryDetailRepository.CountCompanyManagementHistoryDetailAll(companyManagementHistoryDetailHistoryId)
	return res, err
}

func (service *companyManagementHistoryDetailService) FindCompanyManagementHistoryDetails(companyManagementHistoryDetailHistoryId int) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error) {
	res, err := service.companyManagementHistoryDetailRepository.FindCompanyManagementHistoryDetails(companyManagementHistoryDetailHistoryId)
	return res, err
}

func (service *companyManagementHistoryDetailService) FindCompanyManagementHistoryDetailsOffset(limit int, offset int, order string, dir string, companyManagementHistoryDetailHistoryId int) (companyManagementHistoryDetailOutput []model.SelectCompanyManagementHistoryDetail, err error) {
	res, err := service.companyManagementHistoryDetailRepository.FindCompanyManagementHistoryDetailsOffset(limit, offset, order, dir, companyManagementHistoryDetailHistoryId)
	return res, err
}

func (service *companyManagementHistoryDetailService) SearchCompanyManagementHistoryDetail(limit int, offset int, order string, dir string, search string, companyManagementHistoryDetailHistoryId int) (companyManagementHistoryDetailOutput []model.SelectCompanyManagementHistoryDetail, err error) {
	res, err := service.companyManagementHistoryDetailRepository.SearchCompanyManagementHistoryDetail(limit, offset, order, dir, search, companyManagementHistoryDetailHistoryId)
	return res, err
}

func (service *companyManagementHistoryDetailService) CountSearchCompanyManagementHistoryDetail(search string, companyManagementHistoryDetailHistoryId int) (count int64, err error) {
	res, err := service.companyManagementHistoryDetailRepository.CountSearchCompanyManagementHistoryDetail(search, companyManagementHistoryDetailHistoryId)
	return res, err
}

func (service *companyManagementHistoryDetailService) FindCompanyManagementHistoryDetailByCompanyId(id uint) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error) {
	res, err := service.companyManagementHistoryDetailRepository.FindCompanyManagementHistoryDetailByCompanyId(id)
	return res, err
}

func (service *companyManagementHistoryDetailService) FindCompanyManagementHistoryDetailById(id uint) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error) {
	res, err := service.companyManagementHistoryDetailRepository.FindCompanyManagementHistoryDetailById(id)
	return res, err
}

func (service *companyManagementHistoryDetailService) FindExcCompanyManagementHistoryDetail(id uint) (companyManagementHistoryDetailOutput []model.CompanyManagementHistoryDetail, err error) {
	res, err := service.companyManagementHistoryDetailRepository.FindExcCompanyManagementHistoryDetail(id)
	return res, err
}

func (service *companyManagementHistoryDetailService) InsertCompanyManagementHistoryDetail(companyManagementHistoryDetail model.CreateCompanyManagementHistoryDetailParameter) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error) {
	newCompanyManagementHistoryDetail := model.CompanyManagementHistoryDetail{}
	err1 := smapping.FillStruct(&newCompanyManagementHistoryDetail, smapping.MapFields(&companyManagementHistoryDetail))
	if err != nil {
		return newCompanyManagementHistoryDetail, err1
	}
	res, err := service.companyManagementHistoryDetailRepository.InsertCompanyManagementHistoryDetail(newCompanyManagementHistoryDetail)
	return res, err
}

func (service *companyManagementHistoryDetailService) UpdateCompanyManagementHistoryDetail(companyManagementHistoryDetail model.CompanyManagementHistoryDetail, id uint) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error) {
	newCompanyManagementHistoryDetail := model.CompanyManagementHistoryDetail{}
	err1 := smapping.FillStruct(&newCompanyManagementHistoryDetail, smapping.MapFields(&companyManagementHistoryDetail))
	if err != nil {
		return newCompanyManagementHistoryDetail, err1
	}
	res, err := service.companyManagementHistoryDetailRepository.UpdateCompanyManagementHistoryDetail(newCompanyManagementHistoryDetail, id)
	return res, err
}

func (service *companyManagementHistoryDetailService) DeleteCompanyManagementHistoryDetail(companyManagementHistoryDetail model.CompanyManagementHistoryDetail, id uint) (companyManagementHistoryDetailOutput model.CompanyManagementHistoryDetail, err error) {
	newCompanyManagementHistoryDetail := model.CompanyManagementHistoryDetail{}
	err1 := smapping.FillStruct(&newCompanyManagementHistoryDetail, smapping.MapFields(&companyManagementHistoryDetail))
	if err != nil {
		return newCompanyManagementHistoryDetail, err1
	}
	res, err := service.companyManagementHistoryDetailRepository.UpdateCompanyManagementHistoryDetail(newCompanyManagementHistoryDetail, id)
	return res, err
}
