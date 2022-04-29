package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type LicenseTypeService interface {
	FindLicenseTypes() (licenseTypeOutput []model.LicenseType, err error)
	FindLicenseTypeById(id uint) (licenseTypeOutput model.LicenseType, err error)
	FindExcLicenseType(id uint) (licenseTypeOutput []model.LicenseType, err error)
	InsertLicenseType(licenseType model.CreateLicenseTypeParameter) (licenseTypeOutput model.LicenseType, err error)
	UpdateLicenseType(licenseType model.LicenseType, id uint) (licenseTypeOutput model.LicenseType, err error)
	DeleteLicenseType(licenseType model.LicenseType, id uint) (licenseTypeOutput model.LicenseType, err error)
}

type licenseTypeService struct {
	licenseTypeRepository repository.LicenseTypeRepository
}

func NewLicenseTypeService(licenseTypeRep repository.LicenseTypeRepository) LicenseTypeService {
	return &licenseTypeService{
		licenseTypeRepository: licenseTypeRep,
	}
}

func (service *licenseTypeService) FindLicenseTypes() (licenseTypeOutput []model.LicenseType, err error) {
	res, err := service.licenseTypeRepository.FindLicenseTypes()
	return res, err
}

func (service *licenseTypeService) FindLicenseTypeById(id uint) (licenseTypeOutput model.LicenseType, err error) {
	res, err := service.licenseTypeRepository.FindLicenseTypeById(id)
	return res, err
}

func (service *licenseTypeService) FindExcLicenseType(id uint) (licenseTypeOutput []model.LicenseType, err error) {
	res, err := service.licenseTypeRepository.FindExcLicenseType(id)
	return res, err
}

func (service *licenseTypeService) InsertLicenseType(licenseType model.CreateLicenseTypeParameter) (licenseTypeOutput model.LicenseType, err error) {
	newLicenseType := model.LicenseType{}
	err1 := smapping.FillStruct(&newLicenseType, smapping.MapFields(&licenseType))
	if err != nil {
		return newLicenseType, err1
	}
	res, err := service.licenseTypeRepository.InsertLicenseType(newLicenseType)
	return res, err
}

func (service *licenseTypeService) UpdateLicenseType(licenseType model.LicenseType, id uint) (licenseTypeOutput model.LicenseType, err error) {
	newLicenseType := model.LicenseType{}
	err1 := smapping.FillStruct(&newLicenseType, smapping.MapFields(&licenseType))
	if err != nil {
		return newLicenseType, err1
	}
	res, err := service.licenseTypeRepository.UpdateLicenseType(newLicenseType, id)
	return res, err
}

func (service *licenseTypeService) DeleteLicenseType(licenseType model.LicenseType, id uint) (licenseTypeOutput model.LicenseType, err error) {
	newLicenseType := model.LicenseType{}
	err1 := smapping.FillStruct(&newLicenseType, smapping.MapFields(&licenseType))
	if err != nil {
		return newLicenseType, err1
	}
	res, err := service.licenseTypeRepository.UpdateLicenseType(newLicenseType, id)
	return res, err
}
