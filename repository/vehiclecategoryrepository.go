package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type VehicleCategoryRepository interface {
	CountVehicleCategoryAll() (count int64, err error)
	FindVehicleCategorys() (vehiclecategoryOutput []model.VehicleCategory, err error)
	FindVehicleCategorysOffset(limit int, offset int, order string, dir string) (vehiclecategoryOutput []model.VehicleCategory, err error)
	SearchVehicleCategory(limit int, offset int, order string, dir string, search string) (vehiclecategoryOutput []model.VehicleCategory, err error)
	CountSearchVehicleCategory(search string) (count int64, err error)
	FindVehicleCategoryById(id uint) (vehicleCategoryOutput model.VehicleCategory, err error)
	FindExcVehicleCategory(id uint) (vehicleCategoryOutput []model.VehicleCategory, err error)
	InsertVehicleCategory(vehicleCategory model.VehicleCategory) (vehicleCategoryOutput model.VehicleCategory, err error)
	UpdateVehicleCategory(vehicleCategory model.VehicleCategory, id uint) (vehicleCategoryOutput model.VehicleCategory, err error)
}

type vehicleCategoryConnection struct {
	connection *gorm.DB
}

func NewVehicleCategoryRepository(db *gorm.DB) VehicleCategoryRepository {
	return &vehicleCategoryConnection{
		connection: db,
	}
}
func (db *vehicleCategoryConnection) CountVehicleCategoryAll() (count int64, err error) {
	res := db.connection.Table("vehicle_categories").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *vehicleCategoryConnection) FindVehicleCategorys() (vehiclecategoryOutput []model.VehicleCategory, err error) {
	var (
		vehiclecategorys []model.VehicleCategory
	)
	res := db.connection.Where("deleted_at = 0").Order("vehicle_category_name").Find(&vehiclecategorys)
	return vehiclecategorys, res.Error
}

func (db *vehicleCategoryConnection) FindVehicleCategorysOffset(limit int, offset int, order string, dir string) (vehiclecategoryOutput []model.VehicleCategory, err error) {
	var (
		orderDirection   string
		vehiclecategorys []model.VehicleCategory
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&vehiclecategorys)
	return vehiclecategorys, res.Error
}

func (db *vehicleCategoryConnection) SearchVehicleCategory(limit int, offset int, order string, dir string, search string) (vehiclecategoryOutput []model.VehicleCategory, err error) {
	var (
		orderDirection   string
		final            string
		vehiclecategorys []model.VehicleCategory
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(vehicle_category_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&vehiclecategorys)
	return vehiclecategorys, res.Error
}

func (db *vehicleCategoryConnection) CountSearchVehicleCategory(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("vehicle_categories").Where("(lower(vehicle_category_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
}

func (db *vehicleCategoryConnection) FindVehicleCategoryById(id uint) (vehicleCategoryOutput model.VehicleCategory, err error) {
	var (
		vehicleCategory model.VehicleCategory
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&vehicleCategory)
	return vehicleCategory, res.Error
}

func (db *vehicleCategoryConnection) FindExcVehicleCategory(id uint) (vehicleCategoryOutput []model.VehicleCategory, err error) {
	var (
		vehicleCategorys []model.VehicleCategory
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("vehicle_category_name").Find(&vehicleCategorys)
	return vehicleCategorys, res.Error
}

func (db *vehicleCategoryConnection) InsertVehicleCategory(vehicleCategory model.VehicleCategory) (vehicleCategoryOutput model.VehicleCategory, err error) {
	res := db.connection.Save(&vehicleCategory)
	return vehicleCategory, res.Error
}

func (db *vehicleCategoryConnection) UpdateVehicleCategory(vehicleCategory model.VehicleCategory, id uint) (vehicleCategoryOutput model.VehicleCategory, err error) {
	res := db.connection.Where("id=?", id).Updates(&vehicleCategory)
	return vehicleCategory, res.Error
}
