package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type VehicleTypeService interface {
	CountVehicleTypeAll() (count int64, err error)
	FindVehicleTypes() (businessunitOutput []model.VehicleType, err error)
	FindVehicleTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.VehicleType, err error)
	SearchVehicleType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.VehicleType, err error)
	CountSearchVehicleType(search string) (count int64, err error)
	FindVehicleTypeById(id uint) (vehicleTypeOutput model.VehicleType, err error)
	FindExcVehicleType(id uint) (vehicleTypeOutput []model.VehicleType, err error)
	InsertVehicleType(vehicleType model.CreateVehicleTypeParameter) (vehicleTypeOutput model.VehicleType, err error)
	UpdateVehicleType(vehicleType model.VehicleType, id uint) (vehicleTypeOutput model.VehicleType, err error)
	DeleteVehicleType(vehicleType model.VehicleType, id uint) (vehicleTypeOutput model.VehicleType, err error)
}

type vehicleTypeService struct {
	vehicleTypeRepository repository.VehicleTypeRepository
}

func NewVehicleTypeService(vehicleTypeRep repository.VehicleTypeRepository) VehicleTypeService {
	return &vehicleTypeService{
		vehicleTypeRepository: vehicleTypeRep,
	}
}

func (service *vehicleTypeService) CountVehicleTypeAll() (count int64, err error) {
	res, err := service.vehicleTypeRepository.CountVehicleTypeAll()
	return res, err
}

func (service *vehicleTypeService) FindVehicleTypes() (businessunitOutput []model.VehicleType, err error) {
	res, err := service.vehicleTypeRepository.FindVehicleTypes()
	return res, err
}

func (service *vehicleTypeService) FindVehicleTypesOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.VehicleType, err error) {
	res, err := service.vehicleTypeRepository.FindVehicleTypesOffset(limit, offset, order, dir)
	return res, err
}

func (service *vehicleTypeService) SearchVehicleType(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.VehicleType, err error) {
	res, err := service.vehicleTypeRepository.SearchVehicleType(limit, offset, order, dir, search)
	return res, err
}

func (service *vehicleTypeService) CountSearchVehicleType(search string) (count int64, err error) {
	res, err := service.vehicleTypeRepository.CountSearchVehicleType(search)
	return res, err
}

func (service *vehicleTypeService) FindVehicleTypeById(id uint) (vehicleTypeOutput model.VehicleType, err error) {
	return service.vehicleTypeRepository.FindVehicleTypeById(id)
}

func (service *vehicleTypeService) FindExcVehicleType(id uint) (vehicleTypeOutput []model.VehicleType, err error) {
	return service.vehicleTypeRepository.FindExcVehicleType(id)
}
func (service *vehicleTypeService) InsertVehicleType(vehicleType model.CreateVehicleTypeParameter) (model.VehicleType, error) {
	newVehicleType := model.VehicleType{}
	err1 := smapping.FillStruct(&newVehicleType, smapping.MapFields(&vehicleType))

	if err1 != nil {
		return newVehicleType, err1
	}

	return service.vehicleTypeRepository.InsertVehicleType(newVehicleType)
}

func (service *vehicleTypeService) UpdateVehicleType(vehicleType model.VehicleType, id uint) (vehicleTypeOutput model.VehicleType, err error) {
	newVehicleType := model.VehicleType{}
	err1 := smapping.FillStruct(&newVehicleType, smapping.MapFields(&vehicleType))

	if err1 != nil {
		return newVehicleType, err1
	}

	return service.vehicleTypeRepository.UpdateVehicleType(newVehicleType, id)
}

func (service *vehicleTypeService) DeleteVehicleType(vehicleType model.VehicleType, id uint) (vehicleTypeOutput model.VehicleType, err error) {
	newVehicleType := model.VehicleType{}
	err1 := smapping.FillStruct(&newVehicleType, smapping.MapFields(&vehicleType))

	if err1 != nil {
		return newVehicleType, err1
	}

	return service.vehicleTypeRepository.UpdateVehicleType(newVehicleType, id)
}
