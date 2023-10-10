package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type VehicleMerkService interface {
	CountVehicleMerkAll() (count int64, err error)
	FindVehicleMerks() (businessunitOutput []model.VehicleMerk, err error)
	FindVehicleMerksOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.VehicleMerk, err error)
	SearchVehicleMerk(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.VehicleMerk, err error)
	CountSearchVehicleMerk(search string) (count int64, err error)
	FindVehicleMerkById(id uint) (vehicleMerkOutput model.VehicleMerk, err error)
	FindExcVehicleMerk(id uint) (vehicleMerkOutput []model.VehicleMerk, err error)
	InsertVehicleMerk(vehicleMerk model.CreateVehicleMerkParameter) (vehicleMerkOutput model.VehicleMerk, err error)
	UpdateVehicleMerk(vehicleMerk model.VehicleMerk, id uint) (vehicleMerkOutput model.VehicleMerk, err error)
	DeleteVehicleMerk(vehicleMerk model.VehicleMerk, id uint) (vehicleMerkOutput model.VehicleMerk, err error)
}

type vehicleMerkService struct {
	vehicleMerkRepository repository.VehicleMerkRepository
}

func NewVehicleMerkService(vehicleMerkRep repository.VehicleMerkRepository) VehicleMerkService {
	return &vehicleMerkService{
		vehicleMerkRepository: vehicleMerkRep,
	}
}

func (service *vehicleMerkService) CountVehicleMerkAll() (count int64, err error) {
	res, err := service.vehicleMerkRepository.CountVehicleMerkAll()
	return res, err
}

func (service *vehicleMerkService) FindVehicleMerks() (businessunitOutput []model.VehicleMerk, err error) {
	res, err := service.vehicleMerkRepository.FindVehicleMerks()
	return res, err
}

func (service *vehicleMerkService) FindVehicleMerksOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.VehicleMerk, err error) {
	res, err := service.vehicleMerkRepository.FindVehicleMerksOffset(limit, offset, order, dir)
	return res, err
}

func (service *vehicleMerkService) SearchVehicleMerk(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.VehicleMerk, err error) {
	res, err := service.vehicleMerkRepository.SearchVehicleMerk(limit, offset, order, dir, search)
	return res, err
}

func (service *vehicleMerkService) CountSearchVehicleMerk(search string) (count int64, err error) {
	res, err := service.vehicleMerkRepository.CountSearchVehicleMerk(search)
	return res, err
}

func (service *vehicleMerkService) FindVehicleMerkById(id uint) (vehicleMerkOutput model.VehicleMerk, err error) {
	return service.vehicleMerkRepository.FindVehicleMerkById(id)
}

func (service *vehicleMerkService) FindExcVehicleMerk(id uint) (vehicleMerkOutput []model.VehicleMerk, err error) {
	return service.vehicleMerkRepository.FindExcVehicleMerk(id)
}
func (service *vehicleMerkService) InsertVehicleMerk(vehicleMerk model.CreateVehicleMerkParameter) (model.VehicleMerk, error) {
	newVehicleMerk := model.VehicleMerk{}
	err1 := smapping.FillStruct(&newVehicleMerk, smapping.MapFields(&vehicleMerk))

	if err1 != nil {
		return newVehicleMerk, err1
	}

	return service.vehicleMerkRepository.InsertVehicleMerk(newVehicleMerk)
}

func (service *vehicleMerkService) UpdateVehicleMerk(vehicleMerk model.VehicleMerk, id uint) (vehicleMerkOutput model.VehicleMerk, err error) {
	newVehicleMerk := model.VehicleMerk{}
	err1 := smapping.FillStruct(&newVehicleMerk, smapping.MapFields(&vehicleMerk))

	if err1 != nil {
		return newVehicleMerk, err1
	}

	return service.vehicleMerkRepository.UpdateVehicleMerk(newVehicleMerk, id)
}

func (service *vehicleMerkService) DeleteVehicleMerk(vehicleMerk model.VehicleMerk, id uint) (vehicleMerkOutput model.VehicleMerk, err error) {
	newVehicleMerk := model.VehicleMerk{}
	err1 := smapping.FillStruct(&newVehicleMerk, smapping.MapFields(&vehicleMerk))

	if err1 != nil {
		return newVehicleMerk, err1
	}

	return service.vehicleMerkRepository.UpdateVehicleMerk(newVehicleMerk, id)
}
