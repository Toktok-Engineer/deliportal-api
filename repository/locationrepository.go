package repository

import (
	"deliportal-api/model"
	"log"

	"gorm.io/gorm"
)

type LocationRepository interface {
	FindLocations() (locationOutput []model.Location, err error)
	FindLocationById(id uint) (locationOutput model.Location, err error)
	FindExcLocation(id uint) (locationOutput []model.Location, err error)
	InsertLocation(location model.Location) (locationOutput model.Location, err error)
	UpdateLocation(location model.Location, id uint) (locationOutput model.Location, err error)
}

type LocationConnection struct {
	connection *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &LocationConnection{
		connection: db,
	}
}

func (db *LocationConnection) FindLocations() (locationOutput []model.Location, err error) {
	var (
		locations []model.Location
	)
	res := db.connection.Where("deleted_at = 0").Order("location_name").Find(&locations)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return locations, res.Error
	}
	return locations, nil
}

func (db *LocationConnection) FindLocationById(id uint) (locationOutput model.Location, err error) {
	var (
		location model.Location
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&location)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return location, res.Error
	}
	return location, nil
}

func (db *LocationConnection) FindExcLocation(id uint) (locationOutput []model.Location, err error) {
	var (
		locations []model.Location
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("location_name").Find(&locations)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return locations, res.Error
	}
	return locations, nil
}

func (db *LocationConnection) InsertLocation(location model.Location) (locationOutput model.Location, err error) {
	res := db.connection.Save(&location)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return location, res.Error
	}
	return location, nil
}

func (db *LocationConnection) UpdateLocation(location model.Location, id uint) (locationOutput model.Location, err error) {
	res := db.connection.Where("id=?", id).Updates(&location)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return location, res.Error
	}
	return location, nil
}
