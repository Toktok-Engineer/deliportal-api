package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type VehicleLicenseTypeRepository interface {
	CountVehicleLicenseTypeAll() (count int64, err error)
	FindVehicleLicenseTypes() (vehiclelicensetypeOutput []model.VehicleLicenseType, err error)
	FindVehicleLicenseTypesOffset(limit int, offset int, order string, dir string) (vehiclelicensetypeOutput []model.VehicleLicenseType, err error)
	SearchVehicleLicenseType(limit int, offset int, order string, dir string, search string) (vehiclelicensetypeOutput []model.VehicleLicenseType, err error)
	CountSearchVehicleLicenseType(search string) (count int64, err error)
	FindVehicleLicenseTypeById(id uint) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error)
	FindExcVehicleLicenseType(id uint) (vehicleLicenseTypeOutput []model.VehicleLicenseType, err error)
	InsertVehicleLicenseType(vehicleLicenseType model.VehicleLicenseType) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error)
	UpdateVehicleLicenseType(vehicleLicenseType model.VehicleLicenseType, id uint) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error)
}

type vehicleLicenseTypeConnection struct {
	connection *gorm.DB
}

func NewVehicleLicenseTypeRepository(db *gorm.DB) VehicleLicenseTypeRepository {
	return &vehicleLicenseTypeConnection{
		connection: db,
	}
}
func (db *vehicleLicenseTypeConnection) CountVehicleLicenseTypeAll() (count int64, err error) {
	res := db.connection.Table("vehicle_license_types").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *vehicleLicenseTypeConnection) FindVehicleLicenseTypes() (vehiclelicensetypeOutput []model.VehicleLicenseType, err error) {
	var (
		vehiclelicensetypes []model.VehicleLicenseType
	)
	res := db.connection.Where("deleted_at = 0").Order("vehicle_license_type_name").Find(&vehiclelicensetypes)
	return vehiclelicensetypes, res.Error
}

func (db *vehicleLicenseTypeConnection) FindVehicleLicenseTypesOffset(limit int, offset int, order string, dir string) (vehiclelicensetypeOutput []model.VehicleLicenseType, err error) {
	var (
		orderDirection      string
		vehiclelicensetypes []model.VehicleLicenseType
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&vehiclelicensetypes)
	return vehiclelicensetypes, res.Error
}

func (db *vehicleLicenseTypeConnection) SearchVehicleLicenseType(limit int, offset int, order string, dir string, search string) (vehiclelicensetypeOutput []model.VehicleLicenseType, err error) {
	var (
		orderDirection      string
		final               string
		vehiclelicensetypes []model.VehicleLicenseType
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(vehicle_license_type_name) LIKE ? OR lower(reminder_first_month::TEXT) LIKE ? OR lower(reminder_month::TEXT) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&vehiclelicensetypes)
	return vehiclelicensetypes, res.Error
}

func (db *vehicleLicenseTypeConnection) CountSearchVehicleLicenseType(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("vehicle_license_types").Where("(lower(vehicle_license_type_name) LIKE ? OR lower(reminder_first_month::TEXT) LIKE ? OR lower(reminder_month::TEXT) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final, final).Count(&count)
	return count, res.Error
}

func (db *vehicleLicenseTypeConnection) FindVehicleLicenseTypeById(id uint) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error) {
	var (
		vehicleLicenseType model.VehicleLicenseType
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&vehicleLicenseType)
	return vehicleLicenseType, res.Error
}

func (db *vehicleLicenseTypeConnection) FindExcVehicleLicenseType(id uint) (vehicleLicenseTypeOutput []model.VehicleLicenseType, err error) {
	var (
		vehicleLicenseTypes []model.VehicleLicenseType
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("vehicle_license_type_name").Find(&vehicleLicenseTypes)
	return vehicleLicenseTypes, res.Error
}

func (db *vehicleLicenseTypeConnection) InsertVehicleLicenseType(vehicleLicenseType model.VehicleLicenseType) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error) {
	res := db.connection.Save(&vehicleLicenseType)
	return vehicleLicenseType, res.Error
}

func (db *vehicleLicenseTypeConnection) UpdateVehicleLicenseType(vehicleLicenseType model.VehicleLicenseType, id uint) (vehicleLicenseTypeOutput model.VehicleLicenseType, err error) {
	var (
		VehicleLicenseType model.VehicleLicenseType
	)
	res := db.connection.Model(&VehicleLicenseType).Where("id=?", id).Updates(map[string]interface{}{"vehicle_license_type_name": vehicleLicenseType.VehicleLicenseTypeName, "reminder_first_month": vehicleLicenseType.ReminderFirstMonth, "reminder_month": vehicleLicenseType.ReminderMonth, "remark": vehicleLicenseType.Remark, "updated_user_id": vehicleLicenseType.UpdatedUserID, "updated_at": vehicleLicenseType.UpdatedAt})
	return vehicleLicenseType, res.Error
}
