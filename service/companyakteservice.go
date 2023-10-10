package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyAkteService interface {
	CountCompanyAkteAll(companyId int) (count int64, err error)
	FindCompanyAktes(companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	FindCompanyAktesOffset(limit int, offset int, order string, dir string, companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	SearchCompanyAkte(limit int, offset int, order string, dir string, search string, companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	CountSearchCompanyAkte(search string, companyId int) (count int64, err error)
	FindCompanyAkteById(id uint) (companyAkteOutput model.SelectCompanyAkteParameter, err error)
	FindCompanyAkteByYear(companyId uint, year uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	FindExcCompanyAkteByYear(companyId uint, year uint, companyAkteId uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	FindExcCompanyAkte(id uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error)
	InsertCompanyAkte(companyAkte model.CreateCompanyAkteParameter) (companyAkteOutput model.CompanyAkte, err error)
	UpdateCompanyAkte(companyAkte model.CompanyAkte, id uint) (companyAkteOutput model.CompanyAkte, err error)
	DeleteCompanyAkte(companyAkte model.CompanyAkte, id uint) (companyAkteOutput model.CompanyAkte, err error)
}

type companyAkteService struct {
	companyAkteRepository repository.CompanyAkteRepository
}

func NewCompanyAkteService(companyAkteRep repository.CompanyAkteRepository) CompanyAkteService {
	return &companyAkteService{
		companyAkteRepository: companyAkteRep,
	}
}

func (service *companyAkteService) CountCompanyAkteAll(companyId int) (count int64, err error) {
	res, err := service.companyAkteRepository.CountCompanyAkteAll(companyId)
	return res, err
}

func (service *companyAkteService) FindCompanyAktes(companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	res, err := service.companyAkteRepository.FindCompanyAktes(companyId)
	return res, err
}

func (service *companyAkteService) FindCompanyAktesOffset(limit int, offset int, order string, dir string, companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	res, err := service.companyAkteRepository.FindCompanyAktesOffset(limit, offset, order, dir, companyId)
	return res, err
}

func (service *companyAkteService) SearchCompanyAkte(limit int, offset int, order string, dir string, search string, companyId int) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	res, err := service.companyAkteRepository.SearchCompanyAkte(limit, offset, order, dir, search, companyId)
	return res, err
}

func (service *companyAkteService) CountSearchCompanyAkte(search string, companyId int) (count int64, err error) {
	res, err := service.companyAkteRepository.CountSearchCompanyAkte(search, companyId)
	return res, err
}

func (service *companyAkteService) FindCompanyAkteById(id uint) (companyAkteOutput model.SelectCompanyAkteParameter, err error) {
	res, err := service.companyAkteRepository.FindCompanyAkteById(id)
	return res, err
}

func (service *companyAkteService) FindCompanyAkteByYear(companyId uint, year uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	res, err := service.companyAkteRepository.FindCompanyAkteByYear(companyId, year)
	return res, err
}

func (service *companyAkteService) FindExcCompanyAkteByYear(companyId uint, year uint, companyAkteId uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	res, err := service.companyAkteRepository.FindExcCompanyAkteByYear(companyId, year, companyAkteId)
	return res, err
}

func (service *companyAkteService) FindExcCompanyAkte(id uint) (companyAkteOutput []model.SelectCompanyAkteParameter, err error) {
	res, err := service.companyAkteRepository.FindExcCompanyAkte(id)
	return res, err
}

func (service *companyAkteService) InsertCompanyAkte(companyAkte model.CreateCompanyAkteParameter) (companyAkteOutput model.CompanyAkte, err error) {
	newCompanyAkte := model.CompanyAkte{}
	err1 := smapping.FillStruct(&newCompanyAkte, smapping.MapFields(&companyAkte))
	if err != nil {
		return newCompanyAkte, err1
	}
	res, err := service.companyAkteRepository.InsertCompanyAkte(newCompanyAkte)
	return res, err
}

func (service *companyAkteService) UpdateCompanyAkte(companyAkte model.CompanyAkte, id uint) (companyAkteOutput model.CompanyAkte, err error) {
	newCompanyAkte := model.CompanyAkte{}
	err1 := smapping.FillStruct(&newCompanyAkte, smapping.MapFields(&companyAkte))
	if err != nil {
		return newCompanyAkte, err1
	}
	res, err := service.companyAkteRepository.UpdateCompanyAkte(newCompanyAkte, id)
	return res, err
}

func (service *companyAkteService) DeleteCompanyAkte(companyAkte model.CompanyAkte, id uint) (companyAkteOutput model.CompanyAkte, err error) {
	newCompanyAkte := model.CompanyAkte{}
	err1 := smapping.FillStruct(&newCompanyAkte, smapping.MapFields(&companyAkte))
	if err != nil {
		return newCompanyAkte, err1
	}
	res, err := service.companyAkteRepository.UpdateCompanyAkte(newCompanyAkte, id)
	return res, err
}
