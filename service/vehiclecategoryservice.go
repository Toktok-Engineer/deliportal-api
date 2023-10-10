package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type VehicleCategoryService interface {
	CountVehicleCategoryAll() (count int64, err error)
	FindVehicleCategorys() (businessunitOutput []model.VehicleCategory, err error)
	FindVehicleCategorysOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.VehicleCategory, err error)
	SearchVehicleCategory(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.VehicleCategory, err error)
	CountSearchVehicleCategory(search string) (count int64, err error)
	FindVehicleCategoryById(id uint) (vehicleCategoryOutput model.VehicleCategory, err error)
	FindExcVehicleCategory(id uint) (vehicleCategoryOutput []model.VehicleCategory, err error)
	InsertVehicleCategory(vehicleCategory model.CreateVehicleCategoryParameter) (vehicleCategoryOutput model.VehicleCategory, err error)
	UpdateVehicleCategory(vehicleCategory model.VehicleCategory, id uint) (vehicleCategoryOutput model.VehicleCategory, err error)
	DeleteVehicleCategory(vehicleCategory model.VehicleCategory, id uint) (vehicleCategoryOutput model.VehicleCategory, err error)
}

type vehicleCategoryService struct {
	vehicleCategoryRepository repository.VehicleCategoryRepository
}

func NewVehicleCategoryService(vehicleCategoryRep repository.VehicleCategoryRepository) VehicleCategoryService {
	return &vehicleCategoryService{
		vehicleCategoryRepository: vehicleCategoryRep,
	}
}

func (service *vehicleCategoryService) CountVehicleCategoryAll() (count int64, err error) {
	res, err := service.vehicleCategoryRepository.CountVehicleCategoryAll()
	return res, err
}

func (service *vehicleCategoryService) FindVehicleCategorys() (businessunitOutput []model.VehicleCategory, err error) {
	res, err := service.vehicleCategoryRepository.FindVehicleCategorys()
	return res, err
}

func (service *vehicleCategoryService) FindVehicleCategorysOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.VehicleCategory, err error) {
	res, err := service.vehicleCategoryRepository.FindVehicleCategorysOffset(limit, offset, order, dir)
	return res, err
}

func (service *vehicleCategoryService) SearchVehicleCategory(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.VehicleCategory, err error) {
	res, err := service.vehicleCategoryRepository.SearchVehicleCategory(limit, offset, order, dir, search)
	return res, err
}

func (service *vehicleCategoryService) CountSearchVehicleCategory(search string) (count int64, err error) {
	res, err := service.vehicleCategoryRepository.CountSearchVehicleCategory(search)
	return res, err
}

func (service *vehicleCategoryService) FindVehicleCategoryById(id uint) (vehicleCategoryOutput model.VehicleCategory, err error) {
	return service.vehicleCategoryRepository.FindVehicleCategoryById(id)
}

func (service *vehicleCategoryService) FindExcVehicleCategory(id uint) (vehicleCategoryOutput []model.VehicleCategory, err error) {
	return service.vehicleCategoryRepository.FindExcVehicleCategory(id)
}
func (service *vehicleCategoryService) InsertVehicleCategory(vehicleCategory model.CreateVehicleCategoryParameter) (model.VehicleCategory, error) {
	newVehicleCategory := model.VehicleCategory{}
	err1 := smapping.FillStruct(&newVehicleCategory, smapping.MapFields(&vehicleCategory))

	if err1 != nil {
		return newVehicleCategory, err1
	}

	return service.vehicleCategoryRepository.InsertVehicleCategory(newVehicleCategory)
}

func (service *vehicleCategoryService) UpdateVehicleCategory(vehicleCategory model.VehicleCategory, id uint) (vehicleCategoryOutput model.VehicleCategory, err error) {
	newVehicleCategory := model.VehicleCategory{}
	err1 := smapping.FillStruct(&newVehicleCategory, smapping.MapFields(&vehicleCategory))

	if err1 != nil {
		return newVehicleCategory, err1
	}

	return service.vehicleCategoryRepository.UpdateVehicleCategory(newVehicleCategory, id)
}

func (service *vehicleCategoryService) DeleteVehicleCategory(vehicleCategory model.VehicleCategory, id uint) (vehicleCategoryOutput model.VehicleCategory, err error) {
	newVehicleCategory := model.VehicleCategory{}
	err1 := smapping.FillStruct(&newVehicleCategory, smapping.MapFields(&vehicleCategory))

	if err1 != nil {
		return newVehicleCategory, err1
	}

	return service.vehicleCategoryRepository.UpdateVehicleCategory(newVehicleCategory, id)
}
