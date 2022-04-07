package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type LocationService interface {
	FindLocations() (locationOutput []model.Location, err error)
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

func (service *locationService) FindLocations() (locationOutput []model.Location, err error) {
	res, err := service.locationRepository.FindLocations()
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
	if err1 != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.locationRepository.InsertLocation(newLocation)
	return res, err
}

func (service *locationService) UpdateLocation(location model.Location, id uint) (locationOutput model.Location, err error) {
	newLocation := model.Location{}
	err1 := smapping.FillStruct(&newLocation, smapping.MapFields(&location))
	if err1 != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.locationRepository.UpdateLocation(newLocation, id)
	return res, err
}

func (service *locationService) DeleteLocation(location model.Location, id uint) (locationOutput model.Location, err error) {
	newLocation := model.Location{}
	err1 := smapping.FillStruct(&newLocation, smapping.MapFields(&location))
	if err1 != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.locationRepository.UpdateLocation(newLocation, id)
	return res, err
}
