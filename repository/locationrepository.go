package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type LocationRepository interface {
	CountLocationAll() (count int64, err error)
	FindLocations() (locationOutput []model.Location, err error)
	FindLocationsOffset(limit int, offset int, order string, dir string) (locationOutput []model.Location, err error)
	SearchLocation(limit int, offset int, order string, dir string, search string) (locationOutput []model.Location, err error)
	CountSearchLocation(search string) (count int64, err error)
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

func (db *LocationConnection) CountLocationAll() (count int64, err error) {
	res := db.connection.Debug().Table("locations").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *LocationConnection) FindLocations() (locationOutput []model.Location, err error) {
	var (
		locations []model.Location
	)
	res := db.connection.Where("deleted_at = 0").Order("location_name").Find(&locations)
	return locations, res.Error
}

func (db *LocationConnection) FindLocationsOffset(limit int, offset int, order string, dir string) (locationOutput []model.Location, err error) {
	var (
		orderDirection string
		locations      []model.Location
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&locations)
	return locations, res.Error
}

func (db *LocationConnection) SearchLocation(limit int, offset int, order string, dir string, search string) (locationOutput []model.Location, err error) {
	var (
		orderDirection string
		final          string
		locations      []model.Location
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(location_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&locations)
	return locations, res.Error
}

func (db *LocationConnection) CountSearchLocation(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("locations").Where("(lower(location_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
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
