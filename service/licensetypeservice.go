package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type LicenseTypeService interface {
	CountLicenseTypeAll() (count int64, err error)
	FindLicenseTypes() (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	FindLicenseTypesOffset(limit int, offset int, order string, dir string) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	SearchLicenseType(limit int, offset int, order string, dir string, search string) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	CountSearchLicenseType(search string) (count int64, err error)
	FindLicenseTypeById(id uint) (licenseTypeOutput model.SelectLicenseTypeParameter, err error)
	FindExcLicenseType(id uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	FindExcCompleteLicenseType(groupLT uint, id uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
	FindLicenseTypeByGroupLT(groupLT uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error)
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

func (service *licenseTypeService) CountLicenseTypeAll() (count int64, err error) {
	res, err := service.licenseTypeRepository.CountLicenseTypeAll()
	return res, err
}

func (service *licenseTypeService) FindLicenseTypes() (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	res, err := service.licenseTypeRepository.FindLicenseTypes()
	return res, err
}

func (service *licenseTypeService) FindLicenseTypesOffset(limit int, offset int, order string, dir string) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	res, err := service.licenseTypeRepository.FindLicenseTypesOffset(limit, offset, order, dir)
	return res, err
}

func (service *licenseTypeService) SearchLicenseType(limit int, offset int, order string, dir string, search string) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	res, err := service.licenseTypeRepository.SearchLicenseType(limit, offset, order, dir, search)
	return res, err
}

func (service *licenseTypeService) CountSearchLicenseType(search string) (count int64, err error) {
	res, err := service.licenseTypeRepository.CountSearchLicenseType(search)
	return res, err
}

func (service *licenseTypeService) FindLicenseTypeById(id uint) (licenseTypeOutput model.SelectLicenseTypeParameter, err error) {
	res, err := service.licenseTypeRepository.FindLicenseTypeById(id)
	return res, err
}

func (service *licenseTypeService) FindExcLicenseType(id uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	res, err := service.licenseTypeRepository.FindExcLicenseType(id)
	return res, err
}

func (service *licenseTypeService) FindExcCompleteLicenseType(groupLT uint, id uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	res, err := service.licenseTypeRepository.FindExcCompleteLicenseType(groupLT, id)
	return res, err
}

func (service *licenseTypeService) FindLicenseTypeByGroupLT(groupLT uint) (licenseTypeOutput []model.SelectLicenseTypeParameter, err error) {
	res, err := service.licenseTypeRepository.FindLicenseTypeByGroupLT(groupLT)
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
