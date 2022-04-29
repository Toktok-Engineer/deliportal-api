package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type CompanyManagementTypeService interface {
	FindCompanyManagementTypes() (companyManagementTypeOutput []model.CompanyManagementType, err error)
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

func (service *companyManagementTypeService) FindCompanyManagementTypes() (companyManagementTypeOutput []model.CompanyManagementType, err error) {
	res, err := service.companyManagementTypeRepository.FindCompanyManagementTypes()
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
