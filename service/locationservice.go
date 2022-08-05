package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type LocationService interface {
	CountLocationAll() (count int64, err error)
	FindLocations() (locationOutput []model.Location, err error)
	FindLocationsOffset(limit int, offset int, order string, dir string) (locationOutput []model.Location, err error)
	SearchLocation(limit int, offset int, order string, dir string, search string) (locationOutput []model.Location, err error)
	CountSearchLocation(search string) (count int64, err error)
	FindLocationById(id uint) (locationOutput model.Location, err error)
	FindExcLocation(id uint) (locationOutput []model.Location, err error)
	InsertLocation(location model.CreateLocationParameter) (locationOutput model.Location, err error)
	UpdateLocation(location model.Location, id uint) (locationOutput model.Location, err error)
	DeleteLocation(location model.Location, id uint) (locationOutput model.Location, err error)
}

type locationService struct {
	locationRepository repository.LocationRepository
}

func NewLocationService(locationRep repository.LocationRepository) LocationService {
	return &locationService{
		locationRepository: locationRep,
	}
}

func (service *locationService) CountLocationAll() (count int64, err error) {
	res, err := service.locationRepository.CountLocationAll()
	return res, err
}

func (service *locationService) FindLocations() (locationOutput []model.Location, err error) {
	res, err := service.locationRepository.FindLocations()
	return res, err
}

func (service *locationService) FindLocationsOffset(limit int, offset int, order string, dir string) (locationOutput []model.Location, err error) {
	res, err := service.locationRepository.FindLocationsOffset(limit, offset, order, dir)
	return res, err
}

func (service *locationService) SearchLocation(limit int, offset int, order string, dir string, search string) (locationOutput []model.Location, err error) {
	res, err := service.locationRepository.SearchLocation(limit, offset, order, dir, search)
	return res, err
}

func (service *locationService) CountSearchLocation(search string) (count int64, err error) {
	res, err := service.locationRepository.CountSearchLocation(search)
	return res, err
}

func (service *locationService) FindLocationById(id uint) (locationOutput model.Location, err error) {
	res, err := service.locationRepository.FindLocationById(id)
	return res, err
}

func (service *locationService) FindExcLocation(id uint) (locationOutput []model.Location, err error) {
	res, err := service.locationRepository.FindExcLocation(id)
	return res, err
}

func (service *locationService) InsertLocation(location model.CreateLocationParameter) (locationOutput model.Location, err error) {
	newLocation := model.Location{}
	err1 := smapping.FillStruct(&newLocation, smapping.MapFields(&location))
	if err != nil {
		return newLocation, err1
	}
	res, err := service.locationRepository.InsertLocation(newLocation)
	return res, err
}

func (service *locationService) UpdateLocation(location model.Location, id uint) (locationOutput model.Location, err error) {
	newLocation := model.Location{}
	err1 := smapping.FillStruct(&newLocation, smapping.MapFields(&location))
	if err != nil {
		return newLocation, err1
	}
	res, err := service.locationRepository.UpdateLocation(newLocation, id)
	return res, err
}

func (service *locationService) DeleteLocation(location model.Location, id uint) (locationOutput model.Location, err error) {
	newLocation := model.Location{}
	err1 := smapping.FillStruct(&newLocation, smapping.MapFields(&location))
	if err != nil {
		return newLocation, err1
	}
	res, err := service.locationRepository.UpdateLocation(newLocation, id)
	return res, err
}
