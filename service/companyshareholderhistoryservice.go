package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyShareholderHistoryService interface {
	CountCompanyShareholderHistoryAll(companyId int) (count int64, err error)
	FindCompanyShareholderHistorys(companyId int) (companyShareholderHistoryOutput []model.CompanyShareholderHistory, err error)
	FindCompanyShareholderHistorysOffset(limit int, offset int, order string, dir string, companyId int) (companyShareholderHistoryOutput []model.SelectCompanyShareholderHistoryParameter, err error)
	SearchCompanyShareholderHistory(limit int, offset int, order string, dir string, search string, companyId int) (companyShareholderHistoryOutput []model.SelectCompanyShareholderHistoryParameter, err error)
	CountSearchCompanyShareholderHistory(search string, companyId int) (count int64, err error)
	FindCompanyShareholderHistoryById(id uint) (companyShareholderHistoryOutput model.SelectCompanyShareholderHistoryParameter, err error)
	FindExcCompanyShareholderHistory(id uint) (companyShareholderHistoryOutput []model.SelectCompanyShareholderHistoryParameter, err error)
	FindCompanyShareholderHistoryByCompanyId(id uint) (companyShareholderHistoryOutput []model.SelectCompanyShareholderHistoryParameter, err error)
	InsertCompanyShareholderHistory(companyShareholderHistory model.CreateCompanyShareholderHistoryParameter) (companyShareholderHistoryOutput model.CompanyShareholderHistory, err error)
	UpdateCompanyShareholderHistory(companyShareholderHistory model.CompanyShareholderHistory, id uint) (companyShareholderHistoryOutput model.CompanyShareholderHistory, err error)
	DeleteCompanyShareholderHistory(companyShareholderHistory model.CompanyShareholderHistory, id uint) (companyShareholderHistoryOutput model.CompanyShareholderHistory, err error)
}

type companyShareholderHistoryService struct {
	companyShareholderHistoryRepository repository.CompanyShareholderHistoryRepository
}

func NewCompanyShareholderHistoryService(companyShareholderHistoryRep repository.CompanyShareholderHistoryRepository) CompanyShareholderHistoryService {
	return &companyShareholderHistoryService{
		companyShareholderHistoryRepository: companyShareholderHistoryRep,
	}
}

func (service *companyShareholderHistoryService) CountCompanyShareholderHistoryAll(companyId int) (count int64, err error) {
	res, err := service.companyShareholderHistoryRepository.CountCompanyShareholderHistoryAll(companyId)
	return res, err
}

func (service *companyShareholderHistoryService) FindCompanyShareholderHistorys(companyId int) (companyShareholderHistoryOutput []model.CompanyShareholderHistory, err error) {
	res, err := service.companyShareholderHistoryRepository.FindCompanyShareholderHistorys(companyId)
	return res, err
}

func (service *companyShareholderHistoryService) FindCompanyShareholderHistorysOffset(limit int, offset int, order string, dir string, companyId int) (companyShareholderHistoryOutput []model.SelectCompanyShareholderHistoryParameter, err error) {
	res, err := service.companyShareholderHistoryRepository.FindCompanyShareholderHistorysOffset(limit, offset, order, dir, companyId)
	return res, err
}

func (service *companyShareholderHistoryService) SearchCompanyShareholderHistory(limit int, offset int, order string, dir string, search string, companyId int) (companyShareholderHistoryOutput []model.SelectCompanyShareholderHistoryParameter, err error) {
	res, err := service.companyShareholderHistoryRepository.SearchCompanyShareholderHistory(limit, offset, order, dir, search, companyId)
	return res, err
}

func (service *companyShareholderHistoryService) CountSearchCompanyShareholderHistory(search string, companyId int) (count int64, err error) {
	res, err := service.companyShareholderHistoryRepository.CountSearchCompanyShareholderHistory(search, companyId)
	return res, err
}

func (service *companyShareholderHistoryService) FindCompanyShareholderHistoryById(id uint) (companyShareholderHistoryOutput model.SelectCompanyShareholderHistoryParameter, err error) {
	res, err := service.companyShareholderHistoryRepository.FindCompanyShareholderHistoryById(id)
	return res, err
}

func (service *companyShareholderHistoryService) FindExcCompanyShareholderHistory(id uint) (companyShareholderHistoryOutput []model.SelectCompanyShareholderHistoryParameter, err error) {
	res, err := service.companyShareholderHistoryRepository.FindExcCompanyShareholderHistory(id)
	return res, err
}

func (service *companyShareholderHistoryService) FindCompanyShareholderHistoryByCompanyId(id uint) (companyShareholderHistoryOutput []model.SelectCompanyShareholderHistoryParameter, err error) {
	res, err := service.companyShareholderHistoryRepository.FindCompanyShareholderHistoryByCompanyId(id)
	return res, err
}

func (service *companyShareholderHistoryService) InsertCompanyShareholderHistory(companyShareholderHistory model.CreateCompanyShareholderHistoryParameter) (companyShareholderHistoryOutput model.CompanyShareholderHistory, err error) {
	newCompanyShareholderHistory := model.CompanyShareholderHistory{}
	err1 := smapping.FillStruct(&newCompanyShareholderHistory, smapping.MapFields(&companyShareholderHistory))
	if err != nil {
		return newCompanyShareholderHistory, err1
	}
	res, err := service.companyShareholderHistoryRepository.InsertCompanyShareholderHistory(newCompanyShareholderHistory)
	return res, err
}

func (service *companyShareholderHistoryService) UpdateCompanyShareholderHistory(companyShareholderHistory model.CompanyShareholderHistory, id uint) (companyShareholderHistoryOutput model.CompanyShareholderHistory, err error) {
	newCompanyShareholderHistory := model.CompanyShareholderHistory{}
	err1 := smapping.FillStruct(&newCompanyShareholderHistory, smapping.MapFields(&companyShareholderHistory))
	if err != nil {
		return newCompanyShareholderHistory, err1
	}
	res, err := service.companyShareholderHistoryRepository.UpdateCompanyShareholderHistory(newCompanyShareholderHistory, id)
	return res, err
}

func (service *companyShareholderHistoryService) DeleteCompanyShareholderHistory(companyShareholderHistory model.CompanyShareholderHistory, id uint) (companyShareholderHistoryOutput model.CompanyShareholderHistory, err error) {
	newCompanyShareholderHistory := model.CompanyShareholderHistory{}
	err1 := smapping.FillStruct(&newCompanyShareholderHistory, smapping.MapFields(&companyShareholderHistory))
	if err != nil {
		return newCompanyShareholderHistory, err1
	}
	res, err := service.companyShareholderHistoryRepository.UpdateCompanyShareholderHistory(newCompanyShareholderHistory, id)
	return res, err
}
