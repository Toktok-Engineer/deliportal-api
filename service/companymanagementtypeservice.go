package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyManagementTypeService interface {
	CountCompanyManagementTypeAll() (count int64, err error)
	FindCompanyManagementTypes() (companyManagementTypeOutput []model.CompanyManagementType, err error)
	FindCompanyManagementTypesOffset(limit int, offset int, order string, dir string) (companyManagementTypeOutput []model.CompanyManagementType, err error)
	SearchCompanyManagementType(limit int, offset int, order string, dir string, search string) (companyManagementTypeOutput []model.CompanyManagementType, err error)
	CountSearchCompanyManagementType(search string) (count int64, err error)
	FindCompanyManagementTypeById(id uint) (companyManagementTypeOutput model.CompanyManagementType, err error)
	FindExcCompanyManagementType(id uint) (companyManagementTypeOutput []model.CompanyManagementType, err error)
	InsertCompanyManagementType(companyManagementType model.CreateCompanyManagementTypeParameter) (companyManagementTypeOutput model.CompanyManagementType, err error)
	UpdateCompanyManagementType(companyManagementType model.CompanyManagementType, id uint) (companyManagementTypeOutput model.CompanyManagementType, err error)
	DeleteCompanyManagementType(companyManagementType model.CompanyManagementType, id uint) (companyManagementTypeOutput model.CompanyManagementType, err error)
}

type companyManagementTypeService struct {
	companyManagementTypeRepository repository.CompanyManagementTypeRepository
}

func NewCompanyManagementTypeService(companyManagementTypeRep repository.CompanyManagementTypeRepository) CompanyManagementTypeService {
	return &companyManagementTypeService{
		companyManagementTypeRepository: companyManagementTypeRep,
	}
}

func (service *companyManagementTypeService) CountCompanyManagementTypeAll() (count int64, err error) {
	res, err := service.companyManagementTypeRepository.CountCompanyManagementTypeAll()
	return res, err
}

func (service *companyManagementTypeService) FindCompanyManagementTypes() (companyManagementTypeOutput []model.CompanyManagementType, err error) {
	res, err := service.companyManagementTypeRepository.FindCompanyManagementTypes()
	return res, err
}

func (service *companyManagementTypeService) FindCompanyManagementTypesOffset(limit int, offset int, order string, dir string) (companyManagementTypeOutput []model.CompanyManagementType, err error) {
	res, err := service.companyManagementTypeRepository.FindCompanyManagementTypesOffset(limit, offset, order, dir)
	return res, err
}

func (service *companyManagementTypeService) SearchCompanyManagementType(limit int, offset int, order string, dir string, search string) (companyManagementTypeOutput []model.CompanyManagementType, err error) {
	res, err := service.companyManagementTypeRepository.SearchCompanyManagementType(limit, offset, order, dir, search)
	return res, err
}

func (service *companyManagementTypeService) CountSearchCompanyManagementType(search string) (count int64, err error) {
	res, err := service.companyManagementTypeRepository.CountSearchCompanyManagementType(search)
	return res, err
}
func (service *companyManagementTypeService) FindCompanyManagementTypeById(id uint) (companyManagementTypeOutput model.CompanyManagementType, err error) {
	res, err := service.companyManagementTypeRepository.FindCompanyManagementTypeById(id)
	return res, err
}

func (service *companyManagementTypeService) FindExcCompanyManagementType(id uint) (companyManagementTypeOutput []model.CompanyManagementType, err error) {
	res, err := service.companyManagementTypeRepository.FindExcCompanyManagementType(id)
	return res, err
}

func (service *companyManagementTypeService) InsertCompanyManagementType(companyManagementType model.CreateCompanyManagementTypeParameter) (companyManagementTypeOutput model.CompanyManagementType, err error) {
	newCompanyManagementType := model.CompanyManagementType{}
	err1 := smapping.FillStruct(&newCompanyManagementType, smapping.MapFields(&companyManagementType))
	if err != nil {
		return newCompanyManagementType, err1
	}
	res, err := service.companyManagementTypeRepository.InsertCompanyManagementType(newCompanyManagementType)
	return res, err
}

func (service *companyManagementTypeService) UpdateCompanyManagementType(companyManagementType model.CompanyManagementType, id uint) (companyManagementTypeOutput model.CompanyManagementType, err error) {
	newCompanyManagementType := model.CompanyManagementType{}
	err1 := smapping.FillStruct(&newCompanyManagementType, smapping.MapFields(&companyManagementType))
	if err != nil {
		return newCompanyManagementType, err1
	}
	res, err := service.companyManagementTypeRepository.UpdateCompanyManagementType(newCompanyManagementType, id)
	return res, err
}

func (service *companyManagementTypeService) DeleteCompanyManagementType(companyManagementType model.CompanyManagementType, id uint) (companyManagementTypeOutput model.CompanyManagementType, err error) {
	newCompanyManagementType := model.CompanyManagementType{}
	err1 := smapping.FillStruct(&newCompanyManagementType, smapping.MapFields(&companyManagementType))
	if err != nil {
		return newCompanyManagementType, err1
	}
	res, err := service.companyManagementTypeRepository.UpdateCompanyManagementType(newCompanyManagementType, id)
	return res, err
}
