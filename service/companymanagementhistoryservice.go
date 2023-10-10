package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyManagementHistoryService interface {
	CountCompanyManagementHistoryAll(companyId int) (count int64, err error)
	FindCompanyManagementHistorys(companyId int) (companyManagementHistoryOutput []model.CompanyManagementHistory, err error)
	FindCompanyManagementHistorysOffset(limit int, offset int, order string, dir string, companyId int) (companyManagementHistoryOutput []model.SelectCompanyManagementHistoryParameter, err error)
	SearchCompanyManagementHistory(limit int, offset int, order string, dir string, search string, companyId int) (companyManagementHistoryOutput []model.SelectCompanyManagementHistoryParameter, err error)
	CountSearchCompanyManagementHistory(search string, companyId int) (count int64, err error)
	FindCompanyManagementHistoryById(id uint) (companyManagementHistoryOutput model.SelectCompanyManagementHistoryParameter, err error)
	FindExcCompanyManagementHistory(id uint) (companyManagementHistoryOutput []model.SelectCompanyManagementHistoryParameter, err error)
	FindCompanyManagementHistoryByCompanyId(id uint) (companyManagementHistoryOutput []model.SelectCompanyManagementHistoryParameter, err error)
	InsertCompanyManagementHistory(companyManagementHistory model.CreateCompanyManagementHistoryParameter) (companyManagementHistoryOutput model.CompanyManagementHistory, err error)
	UpdateCompanyManagementHistory(companyManagementHistory model.CompanyManagementHistory, id uint) (companyManagementHistoryOutput model.CompanyManagementHistory, err error)
	DeleteCompanyManagementHistory(companyManagementHistory model.CompanyManagementHistory, id uint) (companyManagementHistoryOutput model.CompanyManagementHistory, err error)
}

type companyManagementHistoryService struct {
	companyManagementHistoryRepository repository.CompanyManagementHistoryRepository
}

func NewCompanyManagementHistoryService(companyManagementHistoryRep repository.CompanyManagementHistoryRepository) CompanyManagementHistoryService {
	return &companyManagementHistoryService{
		companyManagementHistoryRepository: companyManagementHistoryRep,
	}
}

func (service *companyManagementHistoryService) CountCompanyManagementHistoryAll(companyId int) (count int64, err error) {
	res, err := service.companyManagementHistoryRepository.CountCompanyManagementHistoryAll(companyId)
	return res, err
}

func (service *companyManagementHistoryService) FindCompanyManagementHistorys(companyId int) (companyManagementHistoryOutput []model.CompanyManagementHistory, err error) {
	res, err := service.companyManagementHistoryRepository.FindCompanyManagementHistorys(companyId)
	return res, err
}

func (service *companyManagementHistoryService) FindCompanyManagementHistorysOffset(limit int, offset int, order string, dir string, companyId int) (companyManagementHistoryOutput []model.SelectCompanyManagementHistoryParameter, err error) {
	res, err := service.companyManagementHistoryRepository.FindCompanyManagementHistorysOffset(limit, offset, order, dir, companyId)
	return res, err
}

func (service *companyManagementHistoryService) SearchCompanyManagementHistory(limit int, offset int, order string, dir string, search string, companyId int) (companyManagementHistoryOutput []model.SelectCompanyManagementHistoryParameter, err error) {
	res, err := service.companyManagementHistoryRepository.SearchCompanyManagementHistory(limit, offset, order, dir, search, companyId)
	return res, err
}

func (service *companyManagementHistoryService) CountSearchCompanyManagementHistory(search string, companyId int) (count int64, err error) {
	res, err := service.companyManagementHistoryRepository.CountSearchCompanyManagementHistory(search, companyId)
	return res, err
}

func (service *companyManagementHistoryService) FindCompanyManagementHistoryById(id uint) (companyManagementHistoryOutput model.SelectCompanyManagementHistoryParameter, err error) {
	res, err := service.companyManagementHistoryRepository.FindCompanyManagementHistoryById(id)
	return res, err
}

func (service *companyManagementHistoryService) FindExcCompanyManagementHistory(id uint) (companyManagementHistoryOutput []model.SelectCompanyManagementHistoryParameter, err error) {
	res, err := service.companyManagementHistoryRepository.FindExcCompanyManagementHistory(id)
	return res, err
}

func (service *companyManagementHistoryService) FindCompanyManagementHistoryByCompanyId(id uint) (companyManagementHistoryOutput []model.SelectCompanyManagementHistoryParameter, err error) {
	res, err := service.companyManagementHistoryRepository.FindCompanyManagementHistoryByCompanyId(id)
	return res, err
}

func (service *companyManagementHistoryService) InsertCompanyManagementHistory(companyManagementHistory model.CreateCompanyManagementHistoryParameter) (companyManagementHistoryOutput model.CompanyManagementHistory, err error) {
	newCompanyManagementHistory := model.CompanyManagementHistory{}
	err1 := smapping.FillStruct(&newCompanyManagementHistory, smapping.MapFields(&companyManagementHistory))
	if err != nil {
		return newCompanyManagementHistory, err1
	}
	res, err := service.companyManagementHistoryRepository.InsertCompanyManagementHistory(newCompanyManagementHistory)
	return res, err
}

func (service *companyManagementHistoryService) UpdateCompanyManagementHistory(companyManagementHistory model.CompanyManagementHistory, id uint) (companyManagementHistoryOutput model.CompanyManagementHistory, err error) {
	newCompanyManagementHistory := model.CompanyManagementHistory{}
	err1 := smapping.FillStruct(&newCompanyManagementHistory, smapping.MapFields(&companyManagementHistory))
	if err != nil {
		return newCompanyManagementHistory, err1
	}
	res, err := service.companyManagementHistoryRepository.UpdateCompanyManagementHistory(newCompanyManagementHistory, id)
	return res, err
}

func (service *companyManagementHistoryService) DeleteCompanyManagementHistory(companyManagementHistory model.CompanyManagementHistory, id uint) (companyManagementHistoryOutput model.CompanyManagementHistory, err error) {
	newCompanyManagementHistory := model.CompanyManagementHistory{}
	err1 := smapping.FillStruct(&newCompanyManagementHistory, smapping.MapFields(&companyManagementHistory))
	if err != nil {
		return newCompanyManagementHistory, err1
	}
	res, err := service.companyManagementHistoryRepository.UpdateCompanyManagementHistory(newCompanyManagementHistory, id)
	return res, err
}
