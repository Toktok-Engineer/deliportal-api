package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type VehicleTypeRepository interface {
	CountVehicleTypeAll() (count int64, err error)
	FindVehicleTypes() (vehicletypeOutput []model.VehicleType, err error)
	FindVehicleTypesOffset(limit int, offset int, order string, dir string) (vehicletypeOutput []model.VehicleType, err error)
	SearchVehicleType(limit int, offset int, order string, dir string, search string) (vehicletypeOutput []model.VehicleType, err error)
	CountSearchVehicleType(search string) (count int64, err error)
	FindVehicleTypeById(id uint) (vehicleTypeOutput model.VehicleType, err error)
	FindExcVehicleType(id uint) (vehicleTypeOutput []model.VehicleType, err error)
	InsertVehicleType(vehicleType model.VehicleType) (vehicleTypeOutput model.VehicleType, err error)
	UpdateVehicleType(vehicleType model.VehicleType, id uint) (vehicleTypeOutput model.VehicleType, err error)
}

type vehicleTypeConnection struct {
	connection *gorm.DB
}

func NewVehicleTypeRepository(db *gorm.DB) VehicleTypeRepository {
	return &vehicleTypeConnection{
		connection: db,
	}
}
func (db *vehicleTypeConnection) CountVehicleTypeAll() (count int64, err error) {
	res := db.connection.Debug().Table("vehicle_types").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *vehicleTypeConnection) FindVehicleTypes() (vehicletypeOutput []model.VehicleType, err error) {
	var (
		vehicletypes []model.VehicleType
	)
	res := db.connection.Where("deleted_at = 0").Order("vehicle_type_name").Find(&vehicletypes)
	return vehicletypes, res.Error
}

func (db *vehicleTypeConnection) FindVehicleTypesOffset(limit int, offset int, order string, dir string) (vehicletypeOutput []model.VehicleType, err error) {
	var (
		orderDirection string
		vehicletypes   []model.VehicleType
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&vehicletypes)
	return vehicletypes, res.Error
}

func (db *vehicleTypeConnection) SearchVehicleType(limit int, offset int, order string, dir string, search string) (vehicletypeOutput []model.VehicleType, err error) {
	var (
		orderDirection string
		final          string
		vehicletypes   []model.VehicleType
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(vehicle_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&vehicletypes)
	return vehicletypes, res.Error
}

func (db *vehicleTypeConnection) CountSearchVehicleType(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("vehicle_types").Where("(lower(vehicle_type_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
}

func (db *vehicleTypeConnection) FindVehicleTypeById(id uint) (vehicleTypeOutput model.VehicleType, err error) {
	var (
		vehicleType model.VehicleType
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&vehicleType)
	return vehicleType, res.Error
}

func (db *vehicleTypeConnection) FindExcVehicleType(id uint) (vehicleTypeOutput []model.VehicleType, err error) {
	var (
		vehicleTypes []model.VehicleType
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("vehicle_type_name").Find(&vehicleTypes)
	return vehicleTypes, res.Error
}

func (db *vehicleTypeConnection) InsertVehicleType(vehicleType model.VehicleType) (vehicleTypeOutput model.VehicleType, err error) {
	res := db.connection.Save(&vehicleType)
	return vehicleType, res.Error
}

func (db *vehicleTypeConnection) UpdateVehicleType(vehicleType model.VehicleType, id uint) (vehicleTypeOutput model.VehicleType, err error) {
	res := db.connection.Where("id=?", id).Updates(&vehicleType)
	return vehicleType, res.Error
}
