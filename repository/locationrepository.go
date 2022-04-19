package repository

import (
	"deliportal-api/model"

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
	return locations, res.Error
}

func (db *LocationConnection) FindLocationById(id uint) (locationOutput model.Location, err error) {
	var (
		location model.Location
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&location)
	return location, res.Error
}

func (db *LocationConnection) FindExcLocation(id uint) (locationOutput []model.Location, err error) {
	var (
		locations []model.Location
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("location_name").Find(&locations)
	return locations, res.Error
}

func (db *LocationConnection) InsertLocation(location model.Location) (locationOutput model.Location, err error) {
	res := db.connection.Save(&location)
	return location, res.Error
}

func (db *LocationConnection) UpdateLocation(location model.Location, id uint) (locationOutput model.Location, err error) {
	res := db.connection.Where("id=?", id).Updates(&location)
	return location, res.Error
}
