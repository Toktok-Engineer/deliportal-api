package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type VehicleService interface {
	CountVehicleAll(companyId int) (count int64, err error)
	FindVehicles(companyId int) (vehicleOutput []model.Vehicle, err error)
	FindVehiclesOffset(limit int, offset int, order string, dir string, companyId int) (vehicleOutput []model.SelectVehicleParameter, err error)
	SearchVehicle(limit int, offset int, order string, dir string, search string, companyId int) (vehicleOutput []model.SelectVehicleParameter, err error)
	CountSearchVehicle(search string, companyId int) (count int64, err error)
	FindVehicleById(id uint) (vehicleOutput model.SelectVehicleParameter, err error)
	FindExcVehicle(id uint) (vehicleOutput []model.SelectVehicleParameter, err error)
	FindVehicleByCompanyId(id uint) (vehicleOutput []model.SelectVehicleParameter, err error)
	InsertVehicle(vehicle model.CreateVehicleParameter) (vehicleOutput model.Vehicle, err error)
	UpdateVehicle(vehicle model.Vehicle, id uint) (vehicleOutput model.Vehicle, err error)
	DeleteVehicle(vehicle model.Vehicle, id uint) (vehicleOutput model.Vehicle, err error)
	CountVehicleFull() (count int64, err error)
	FindVehiclesOffsetFull(limit int, offset int, order string, dir string) (vehicleOutput []model.SelectVehicleParameter, err error)
	SearchVehicleFull(limit int, offset int, order string, dir string, search string) (vehicleOutput []model.SelectVehicleParameter, err error)
	CountSearchVehicleFull(search string) (count int64, err error)
}

type vehicleService struct {
	vehicleRepository repository.VehicleRepository
}

func NewVehicleService(vehicleRep repository.VehicleRepository) VehicleService {
	return &vehicleService{
		vehicleRepository: vehicleRep,
	}
}

func (service *vehicleService) CountVehicleAll(companyId int) (count int64, err error) {
	res, err := service.vehicleRepository.CountVehicleAll(companyId)
	return res, err
}

func (service *vehicleService) FindVehicles(companyId int) (vehicleOutput []model.Vehicle, err error) {
	res, err := service.vehicleRepository.FindVehicles(companyId)
	return res, err
}

func (service *vehicleService) FindVehiclesOffset(limit int, offset int, order string, dir string, companyId int) (vehicleOutput []model.SelectVehicleParameter, err error) {
	res, err := service.vehicleRepository.FindVehiclesOffset(limit, offset, order, dir, companyId)
	return res, err
}

func (service *vehicleService) SearchVehicle(limit int, offset int, order string, dir string, search string, companyId int) (vehicleOutput []model.SelectVehicleParameter, err error) {
	res, err := service.vehicleRepository.SearchVehicle(limit, offset, order, dir, search, companyId)
	return res, err
}

func (service *vehicleService) CountSearchVehicle(search string, companyId int) (count int64, err error) {
	res, err := service.vehicleRepository.CountSearchVehicle(search, companyId)
	return res, err
}

func (service *vehicleService) FindVehicleById(id uint) (vehicleOutput model.SelectVehicleParameter, err error) {
	res, err := service.vehicleRepository.FindVehicleById(id)
	return res, err
}

func (service *vehicleService) FindExcVehicle(id uint) (vehicleOutput []model.SelectVehicleParameter, err error) {
	res, err := service.vehicleRepository.FindExcVehicle(id)
	return res, err
}

func (service *vehicleService) FindVehicleByCompanyId(id uint) (vehicleOutput []model.SelectVehicleParameter, err error) {
	res, err := service.vehicleRepository.FindVehicleByCompanyId(id)
	return res, err
}

func (service *vehicleService) InsertVehicle(vehicle model.CreateVehicleParameter) (vehicleOutput model.Vehicle, err error) {
	newVehicle := model.Vehicle{}
	err1 := smapping.FillStruct(&newVehicle, smapping.MapFields(&vehicle))
	if err != nil {
		return newVehicle, err1
	}
	res, err := service.vehicleRepository.InsertVehicle(newVehicle)
	return res, err
}

func (service *vehicleService) UpdateVehicle(vehicle model.Vehicle, id uint) (vehicleOutput model.Vehicle, err error) {
	newVehicle := model.Vehicle{}
	err1 := smapping.FillStruct(&newVehicle, smapping.MapFields(&vehicle))
	if err != nil {
		return newVehicle, err1
	}
	res, err := service.vehicleRepository.UpdateVehicle(newVehicle, id)
	return res, err
}

func (service *vehicleService) DeleteVehicle(vehicle model.Vehicle, id uint) (vehicleOutput model.Vehicle, err error) {
	newVehicle := model.Vehicle{}
	err1 := smapping.FillStruct(&newVehicle, smapping.MapFields(&vehicle))
	if err != nil {
		return newVehicle, err1
	}
	res, err := service.vehicleRepository.UpdateVehicle(newVehicle, id)
	return res, err
}

func (service *vehicleService) CountVehicleFull() (count int64, err error) {
	res, err := service.vehicleRepository.CountVehicleFull()
	return res, err
}

func (service *vehicleService) FindVehiclesOffsetFull(limit int, offset int, order string, dir string) (vehicleOutput []model.SelectVehicleParameter, err error) {
	res, err := service.vehicleRepository.FindVehiclesOffsetFull(limit, offset, order, dir)
	return res, err
}

func (service *vehicleService) SearchVehicleFull(limit int, offset int, order string, dir string, search string) (vehicleOutput []model.SelectVehicleParameter, err error) {
	res, err := service.vehicleRepository.SearchVehicleFull(limit, offset, order, dir, search)
	return res, err
}

func (service *vehicleService) CountSearchVehicleFull(search string) (count int64, err error) {
	res, err := service.vehicleRepository.CountSearchVehicleFull(search)
	return res, err
}
