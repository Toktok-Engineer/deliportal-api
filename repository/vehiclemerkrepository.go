package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type VehicleMerkRepository interface {
	CountVehicleMerkAll() (count int64, err error)
	FindVehicleMerks() (vehiclemerkOutput []model.VehicleMerk, err error)
	FindVehicleMerksOffset(limit int, offset int, order string, dir string) (vehiclemerkOutput []model.VehicleMerk, err error)
	SearchVehicleMerk(limit int, offset int, order string, dir string, search string) (vehiclemerkOutput []model.VehicleMerk, err error)
	CountSearchVehicleMerk(search string) (count int64, err error)
	FindVehicleMerkById(id uint) (vehicleMerkOutput model.VehicleMerk, err error)
	FindExcVehicleMerk(id uint) (vehicleMerkOutput []model.VehicleMerk, err error)
	InsertVehicleMerk(vehicleMerk model.VehicleMerk) (vehicleMerkOutput model.VehicleMerk, err error)
	UpdateVehicleMerk(vehicleMerk model.VehicleMerk, id uint) (vehicleMerkOutput model.VehicleMerk, err error)
}

type vehicleMerkConnection struct {
	connection *gorm.DB
}

func NewVehicleMerkRepository(db *gorm.DB) VehicleMerkRepository {
	return &vehicleMerkConnection{
		connection: db,
	}
}
func (db *vehicleMerkConnection) CountVehicleMerkAll() (count int64, err error) {
	res := db.connection.Table("vehicle_merks").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *vehicleMerkConnection) FindVehicleMerks() (vehiclemerkOutput []model.VehicleMerk, err error) {
	var (
		vehiclemerks []model.VehicleMerk
	)
	res := db.connection.Where("deleted_at = 0").Order("vehicle_merk_name").Find(&vehiclemerks)
	return vehiclemerks, res.Error
}

func (db *vehicleMerkConnection) FindVehicleMerksOffset(limit int, offset int, order string, dir string) (vehiclemerkOutput []model.VehicleMerk, err error) {
	var (
		orderDirection string
		vehiclemerks   []model.VehicleMerk
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&vehiclemerks)
	return vehiclemerks, res.Error
}

func (db *vehicleMerkConnection) SearchVehicleMerk(limit int, offset int, order string, dir string, search string) (vehiclemerkOutput []model.VehicleMerk, err error) {
	var (
		orderDirection string
		final          string
		vehiclemerks   []model.VehicleMerk
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(vehicle_merk_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&vehiclemerks)
	return vehiclemerks, res.Error
}

func (db *vehicleMerkConnection) CountSearchVehicleMerk(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("vehicle_merks").Where("(lower(vehicle_merk_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
}

func (db *vehicleMerkConnection) FindVehicleMerkById(id uint) (vehicleMerkOutput model.VehicleMerk, err error) {
	var (
		vehicleMerk model.VehicleMerk
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&vehicleMerk)
	return vehicleMerk, res.Error
}

func (db *vehicleMerkConnection) FindExcVehicleMerk(id uint) (vehicleMerkOutput []model.VehicleMerk, err error) {
	var (
		vehicleMerks []model.VehicleMerk
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("vehicle_merk_name").Find(&vehicleMerks)
	return vehicleMerks, res.Error
}

func (db *vehicleMerkConnection) InsertVehicleMerk(vehicleMerk model.VehicleMerk) (vehicleMerkOutput model.VehicleMerk, err error) {
	res := db.connection.Save(&vehicleMerk)
	return vehicleMerk, res.Error
}

func (db *vehicleMerkConnection) UpdateVehicleMerk(vehicleMerk model.VehicleMerk, id uint) (vehicleMerkOutput model.VehicleMerk, err error) {
	res := db.connection.Where("id=?", id).Updates(&vehicleMerk)
	return vehicleMerk, res.Error
}
