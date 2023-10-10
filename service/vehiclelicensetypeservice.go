package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type VehicleLicenseTypeService interface {
	CountVehicleLicenseTypeAll() (count int64, err error)
	FindVehicleLicenseTypes() (businessunitOutput []model.VehicleLicenseType, err error)
	FindVehicleLicenseTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.VehicleLicenseType, err error)
	SearchVehicleLicenseType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.VehicleLicenseType, err error)
	CountSearchVehicleLicenseType(search string) (count int64, err error)
	FindVehicleLicenseTypeById(id uint) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error)
	FindExcVehicleLicenseType(id uint) (vehicleLicenseTypeOutput []model.VehicleLicenseType, err error)
	InsertVehicleLicenseType(vehicleLicenseType model.CreateVehicleLicenseTypeParameter) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error)
	UpdateVehicleLicenseType(vehicleLicenseType model.VehicleLicenseType, id uint) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error)
	DeleteVehicleLicenseType(vehicleLicenseType model.VehicleLicenseType, id uint) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error)
}

type vehicleLicenseTypeService struct {
	vehicleLicenseTypeRepository repository.VehicleLicenseTypeRepository
}

func NewVehicleLicenseTypeService(vehicleLicenseTypeRep repository.VehicleLicenseTypeRepository) VehicleLicenseTypeService {
	return &vehicleLicenseTypeService{
		vehicleLicenseTypeRepository: vehicleLicenseTypeRep,
	}
}

func (service *vehicleLicenseTypeService) CountVehicleLicenseTypeAll() (count int64, err error) {
	res, err := service.vehicleLicenseTypeRepository.CountVehicleLicenseTypeAll()
	return res, err
}

func (service *vehicleLicenseTypeService) FindVehicleLicenseTypes() (businessunitOutput []model.VehicleLicenseType, err error) {
	res, err := service.vehicleLicenseTypeRepository.FindVehicleLicenseTypes()
	return res, err
}

func (service *vehicleLicenseTypeService) FindVehicleLicenseTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.VehicleLicenseType, err error) {
	res, err := service.vehicleLicenseTypeRepository.FindVehicleLicenseTypesOffset(limit, offset, order, dir)
	return res, err
}

func (service *vehicleLicenseTypeService) SearchVehicleLicenseType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.VehicleLicenseType, err error) {
	res, err := service.vehicleLicenseTypeRepository.SearchVehicleLicenseType(limit, offset, order, dir, search)
	return res, err
}

func (service *vehicleLicenseTypeService) CountSearchVehicleLicenseType(search string) (count int64, err error) {
	res, err := service.vehicleLicenseTypeRepository.CountSearchVehicleLicenseType(search)
	return res, err
}

func (service *vehicleLicenseTypeService) FindVehicleLicenseTypeById(id uint) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error) {
	return service.vehicleLicenseTypeRepository.FindVehicleLicenseTypeById(id)
}

func (service *vehicleLicenseTypeService) FindExcVehicleLicenseType(id uint) (vehicleLicenseTypeOutput []model.VehicleLicenseType, err error) {
	return service.vehicleLicenseTypeRepository.FindExcVehicleLicenseType(id)
}
func (service *vehicleLicenseTypeService) InsertVehicleLicenseType(vehicleLicenseType model.CreateVehicleLicenseTypeParameter) (model.VehicleLicenseType, error) {
	newVehicleLicenseType := model.VehicleLicenseType{}
	err1 := smapping.FillStruct(&newVehicleLicenseType, smapping.MapFields(&vehicleLicenseType))

	if err1 != nil {
		return newVehicleLicenseType, err1
	}

	return service.vehicleLicenseTypeRepository.InsertVehicleLicenseType(newVehicleLicenseType)
}

func (service *vehicleLicenseTypeService) UpdateVehicleLicenseType(vehicleLicenseType model.VehicleLicenseType, id uint) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error) {
	newVehicleLicenseType := model.VehicleLicenseType{}
	err1 := smapping.FillStruct(&newVehicleLicenseType, smapping.MapFields(&vehicleLicenseType))

	if err1 != nil {
		return newVehicleLicenseType, err1
	}

	return service.vehicleLicenseTypeRepository.UpdateVehicleLicenseType(newVehicleLicenseType, id)
}

func (service *vehicleLicenseTypeService) DeleteVehicleLicenseType(vehicleLicenseType model.VehicleLicenseType, id uint) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error) {
	newVehicleLicenseType := model.VehicleLicenseType{}
	err1 := smapping.FillStruct(&newVehicleLicenseType, smapping.MapFields(&vehicleLicenseType))

	if err1 != nil {
		return newVehicleLicenseType, err1
	}

	return service.vehicleLicenseTypeRepository.UpdateVehicleLicenseType(newVehicleLicenseType, id)
}
