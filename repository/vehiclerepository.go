package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type VehicleRepository interface {
	CountVehicleAll(companyId int) (count int64, err error)
	FindVehicles(companyId int) (vehicleOutput []model.Vehicle, err error)
	FindVehiclesOffset(limit int, offset int, order string, dir string, companyId int) (vehicleOutput []model.SelectVehicleParameter, err error)
	SearchVehicle(limit int, offset int, order string, dir string, search string, companyId int) (vehicleOutput []model.SelectVehicleParameter, err error)
	CountSearchVehicle(search string, companyId int) (count int64, err error)
	FindVehicleById(id uint) (vehicleOutput model.SelectVehicleParameter, err error)
	FindExcVehicle(id uint) (vehicleOutput []model.SelectVehicleParameter, err error)
	FindVehicleByCompanyId(id uint) (vehicleOutput []model.SelectVehicleParameter, err error)
	InsertVehicle(vehicle model.Vehicle) (vehicleOutput model.Vehicle, err error)
	UpdateVehicle(vehicle model.Vehicle, id uint) (vehicleOutput model.Vehicle, err error)
	CountVehicleFull() (count int64, err error)
	FindVehiclesOffsetFull(limit int, offset int, order string, dir string) (vehicleOutput []model.SelectVehicleParameter, err error)
	SearchVehicleFull(limit int, offset int, order string, dir string, search string) (vehicleOutput []model.SelectVehicleParameter, err error)
	CountSearchVehicleFull(search string) (count int64, err error)
}

type VehicleConnection struct {
	connection *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
	return &VehicleConnection{
		connection: db,
	}
}

func (db *VehicleConnection) CountVehicleAll(companyId int) (count int64, err error) {
	res := db.connection.Table("vehicles").Where("company_id = ? AND deleted_at = 0", companyId).Count(&count)
	return count, res.Error
}

func (db *VehicleConnection) FindVehicles(companyId int) (vehicleOutput []model.Vehicle, err error) {
	var (
		vehicles []model.Vehicle
	)
	res := db.connection.Where("company_id = ? AND deleted_at = 0", companyId).Order("company_id").Find(&vehicles)
	return vehicles, res.Error
}

func (db *VehicleConnection) FindVehiclesOffset(limit int, offset int, order string, dir string, companyId int) (vehicleOutput []model.SelectVehicleParameter, err error) {
	var (
		orderDirection string
		vehicles       []model.SelectVehicleParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("vehicles").Select("vehicles.id, vehicles.vehicle_hull_number, vehicles.vehicle_plate_name, vehicles.vehicle_category_id, vehicle_categories.vehicle_category_name, vehicles.vehicle_merk_id, vehicle_merks.vehicle_merk_name, vehicles.vehicle_type_id, vehicle_types.vehicle_type_name, vehicles.vehicle_registration_name, vehicles.nama_pengguna, vehicles.production_year, vehicles.company_id, companies.company_name, vehicles.chassis_number, vehicles.machine_number, vehicles.original_bpkb_storage, vehicles.bpkb_number, to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY') as tax_stnk_end_date, to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, vehicles.location_id, locations.location_name, vehicles.remark, vehicles.created_user_id, createdUID.username AS created_user, vehicles.updated_user_id, updatedUID.username AS updated_user, vehicles.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(vehicles.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at, vehicles.price").Joins("left join companies on vehicles.company_id = companies.id").Joins("left join vehicle_types ON vehicles.vehicle_type_id = vehicle_types.id").Joins("left join vehicle_categories ON vehicles.vehicle_category_id = vehicle_categories.id").Joins("left join vehicle_merks ON vehicles.vehicle_merk_id = vehicle_merks.id").Joins("left join locations ON vehicles.location_id = locations.id").Joins("left join users createdUID on vehicles.created_user_id = createdUID.id").Joins("left join users updatedUID on vehicles.updated_user_id = updatedUID.id").Joins("left join users deletedUID on vehicles.deleted_user_id = deletedUID.id").Where("vehicles.company_id = ? AND vehicles.deleted_at = 0", companyId).Order(orderDirection).Limit(limit).Offset(offset).Find(&vehicles)
	return vehicles, res.Error
}

func (db *VehicleConnection) SearchVehicle(limit int, offset int, order string, dir string, search string, companyId int) (vehicleOutput []model.SelectVehicleParameter, err error) {
	var (
		orderDirection string
		final          string
		vehicles       []model.SelectVehicleParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("vehicles").Select("vehicles.id, vehicles.vehicle_hull_number, vehicles.vehicle_plate_name, vehicles.vehicle_category_id, vehicle_categories.vehicle_category_name, vehicles.vehicle_merk_id, vehicle_merks.vehicle_merk_name, vehicles.vehicle_type_id, vehicle_types.vehicle_type_name, vehicles.vehicle_registration_name, vehicles.nama_pengguna, vehicles.production_year, vehicles.company_id, companies.company_name, vehicles.chassis_number, vehicles.machine_number, vehicles.original_bpkb_storage, vehicles.bpkb_number, to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY') as tax_stnk_end_date, to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, vehicles.location_id, locations.location_name, vehicles.remark, vehicles.created_user_id, createdUID.username AS created_user, vehicles.updated_user_id, updatedUID.username AS updated_user, vehicles.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(vehicles.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at, vehicles.price").Joins("left join companies on vehicles.company_id = companies.id").Joins("left join vehicle_types ON vehicles.vehicle_type_id = vehicle_types.id").Joins("left join vehicle_categories ON vehicles.vehicle_category_id = vehicle_categories.id").Joins("left join vehicle_merks ON vehicles.vehicle_merk_id = vehicle_merks.id").Joins("left join locations ON vehicles.location_id = locations.id").Joins("left join users createdUID on vehicles.created_user_id = createdUID.id").Joins("left join users updatedUID on vehicles.updated_user_id = updatedUID.id").Joins("left join users deletedUID on vehicles.deleted_user_id = deletedUID.id").Where("(lower(vehicles.vehicle_hull_number) LIKE ? OR lower(vehicles.vehicle_plate_name) LIKE ? OR lower(vehicle_categories.vehicle_category_name) LIKE ? OR lower(vehicle_merks.vehicle_merk_name) LIKE ? OR lower(vehicle_types.vehicle_type_name) LIKE ? OR lower(vehicles.vehicle_registration_name) LIKE ? OR lower(vehicles.production_year::TEXT) LIKE ? OR lower(companies.company_name) LIKE ? OR lower(vehicles.chassis_number) LIKE ? OR lower(vehicles.machine_number) LIKE ? OR lower(vehicles.original_bpkb_storage) LIKE ? OR lower(vehicles.bpkb_number) LIKE ? OR lower(to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(locations.location_name) LIKE ? OR lower(vehicles.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(vehicles.price::TEXT) LIKE ? OR lower(vehicles.nama_pengguna) LIKE ? ) AND vehicles.company_id = ? AND vehicles.deleted_at = 0", final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, companyId).Order(orderDirection).Limit(limit).Offset(offset).Find(&vehicles)
	return vehicles, res.Error
}

func (db *VehicleConnection) CountSearchVehicle(search string, companyId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("vehicles").Table("vehicles").Select("vehicles.id, vehicles.vehicle_hull_number, vehicles.vehicle_plate_name, vehicles.vehicle_category_id, vehicle_categories.vehicle_category_name, vehicles.vehicle_merk_id, vehicle_merks.vehicle_merk_name, vehicles.vehicle_type_id, vehicle_types.vehicle_type_name, vehicles.vehicle_registration_name, vehicles.nama_pengguna, vehicles.production_year, vehicles.company_id, companies.company_name, vehicles.chassis_number, vehicles.machine_number, vehicles.original_bpkb_storage, vehicles.bpkb_number, to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY') as tax_stnk_end_date, to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, vehicles.location_id, locations.location_name, vehicles.remark, vehicles.created_user_id, createdUID.username AS created_user, vehicles.updated_user_id, updatedUID.username AS updated_user, vehicles.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(vehicles.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at, vehicles.price").Joins("left join companies on vehicles.company_id = companies.id").Joins("left join vehicle_types ON vehicles.vehicle_type_id = vehicle_types.id").Joins("left join vehicle_categories ON vehicles.vehicle_category_id = vehicle_categories.id").Joins("left join vehicle_merks ON vehicles.vehicle_merk_id = vehicle_merks.id").Joins("left join locations ON vehicles.location_id = locations.id").Joins("left join users createdUID on vehicles.created_user_id = createdUID.id").Joins("left join users updatedUID on vehicles.updated_user_id = updatedUID.id").Joins("left join users deletedUID on vehicles.deleted_user_id = deletedUID.id").Where("(lower(vehicles.vehicle_hull_number) LIKE ? OR lower(vehicles.vehicle_plate_name) LIKE ? OR lower(vehicle_categories.vehicle_category_name) LIKE ? OR lower(vehicle_merks.vehicle_merk_name) LIKE ? OR lower(vehicle_types.vehicle_type_name) LIKE ? OR lower(vehicles.vehicle_registration_name) LIKE ? OR lower(vehicles.production_year::TEXT) LIKE ? OR lower(companies.company_name) LIKE ? OR lower(vehicles.chassis_number) LIKE ? OR lower(vehicles.machine_number) LIKE ? OR lower(vehicles.original_bpkb_storage) LIKE ? OR lower(vehicles.bpkb_number) LIKE ? OR lower(to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(locations.location_name) LIKE ? OR lower(vehicles.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(vehicles.price::TEXT) LIKE ? OR lower(vehicles.nama_pengguna) LIKE ? ) AND vehicles.company_id = ? AND vehicles.deleted_at = 0", final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, companyId).Count(&count)
	return count, res.Error
}

func (db *VehicleConnection) FindVehicleById(id uint) (vehicleOutput model.SelectVehicleParameter, err error) {
	var (
		vehicle model.SelectVehicleParameter
	)

	res := db.connection.Table("vehicles").Select("vehicles.id, vehicles.vehicle_hull_number, vehicles.vehicle_plate_name, vehicles.vehicle_category_id, vehicle_categories.vehicle_category_name, vehicles.vehicle_merk_id, vehicle_merks.vehicle_merk_name, vehicles.vehicle_type_id, vehicle_types.vehicle_type_name, vehicles.vehicle_registration_name, vehicles.nama_pengguna, vehicles.production_year, vehicles.company_id, companies.company_name, vehicles.chassis_number, vehicles.machine_number, vehicles.original_bpkb_storage, vehicles.bpkb_number, to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY') as tax_stnk_end_date, to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, vehicles.location_id, locations.location_name, vehicles.remark, vehicles.created_user_id, createdUID.username AS created_user, vehicles.updated_user_id, updatedUID.username AS updated_user, vehicles.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(vehicles.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at, vehicles.price").Joins("left join companies on vehicles.company_id = companies.id").Joins("left join vehicle_types ON vehicles.vehicle_type_id = vehicle_types.id").Joins("left join vehicle_categories ON vehicles.vehicle_category_id = vehicle_categories.id").Joins("left join vehicle_merks ON vehicles.vehicle_merk_id = vehicle_merks.id").Joins("left join locations ON vehicles.location_id = locations.id").Joins("left join users createdUID on vehicles.created_user_id = createdUID.id").Joins("left join users updatedUID on vehicles.updated_user_id = updatedUID.id").Joins("left join users deletedUID on vehicles.deleted_user_id = deletedUID.id").Where("vehicles.id=? AND vehicles.deleted_at = 0", id).Take(&vehicle)
	return vehicle, res.Error
}

func (db *VehicleConnection) FindExcVehicle(id uint) (vehicleOutput []model.SelectVehicleParameter, err error) {
	var (
		vehicles []model.SelectVehicleParameter
	)

	res := db.connection.Table("vehicles").Select("vehicles.id, vehicles.vehicle_hull_number, vehicles.vehicle_plate_name, vehicles.vehicle_category_id, vehicle_categories.vehicle_category_name, vehicles.vehicle_merk_id, vehicle_merks.vehicle_merk_name, vehicles.vehicle_type_id, vehicle_types.vehicle_type_name, vehicles.vehicle_registration_name, vehicles.nama_pengguna, vehicles.production_year, vehicles.company_id, companies.company_name, vehicles.chassis_number, vehicles.machine_number, vehicles.original_bpkb_storage, vehicles.bpkb_number, to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY') as tax_stnk_end_date, to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, vehicles.location_id, locations.location_name, vehicles.remark, vehicles.created_user_id, createdUID.username AS created_user, vehicles.updated_user_id, updatedUID.username AS updated_user, vehicles.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(vehicles.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at, vehicles.price").Joins("left join companies on vehicles.company_id = companies.id").Joins("left join vehicle_types ON vehicles.vehicle_type_id = vehicle_types.id").Joins("left join vehicle_categories ON vehicles.vehicle_category_id = vehicle_categories.id").Joins("left join vehicle_merks ON vehicles.vehicle_merk_id = vehicle_merks.id").Joins("left join locations ON vehicles.location_id = locations.id").Joins("left join users createdUID on vehicles.created_user_id = createdUID.id").Joins("left join users updatedUID on vehicles.updated_user_id = updatedUID.id").Joins("left join users deletedUID on vehicles.deleted_user_id = deletedUID.id").Where(" vehicles.id!=? AND vehicles.deleted_at = 0", id).Order("vehicles.id").Find(&vehicles)
	return vehicles, res.Error
}

func (db *VehicleConnection) FindVehicleByCompanyId(id uint) (vehicleOutput []model.SelectVehicleParameter, err error) {
	var (
		vehicles []model.SelectVehicleParameter
	)

	res := db.connection.Table("vehicles").Select("vehicles.id, vehicles.vehicle_hull_number, vehicles.vehicle_plate_name, vehicles.vehicle_category_id, vehicle_categories.vehicle_category_name, vehicles.vehicle_merk_id, vehicle_merks.vehicle_merk_name, vehicles.vehicle_type_id, vehicle_types.vehicle_type_name, vehicles.vehicle_registration_name, vehicles.nama_pengguna, vehicles.production_year, vehicles.company_id, companies.company_name, vehicles.chassis_number, vehicles.machine_number, vehicles.original_bpkb_storage, vehicles.bpkb_number, to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY') as tax_stnk_end_date, to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, vehicles.location_id, locations.location_name, vehicles.remark, vehicles.created_user_id, createdUID.username AS created_user, vehicles.updated_user_id, updatedUID.username AS updated_user, vehicles.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(vehicles.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at, vehicles.price").Joins("left join companies on vehicles.company_id = companies.id").Joins("left join vehicle_types ON vehicles.vehicle_type_id = vehicle_types.id").Joins("left join vehicle_categories ON vehicles.vehicle_category_id = vehicle_categories.id").Joins("left join vehicle_merks ON vehicles.vehicle_merk_id = vehicle_merks.id").Joins("left join locations ON vehicles.location_id = locations.id").Joins("left join users createdUID on vehicles.created_user_id = createdUID.id").Joins("left join users updatedUID on vehicles.updated_user_id = updatedUID.id").Joins("left join users deletedUID on vehicles.deleted_user_id = deletedUID.id").Where("vehicles.company_id=? AND vehicles.deleted_at = 0", id).Order("vehicles.id").Find(&vehicles)
	return vehicles, res.Error
}

func (db *VehicleConnection) InsertVehicle(vehicle model.Vehicle) (vehicleOutput model.Vehicle, err error) {
	res := db.connection.Save(&vehicle)
	return vehicle, res.Error
}

func (db *VehicleConnection) UpdateVehicle(vehicle model.Vehicle, id uint) (vehicleOutput model.Vehicle, err error) {
	res := db.connection.Where("id=?", id).Updates(&vehicle)
	return vehicle, res.Error
}

func (db *VehicleConnection) CountVehicleFull() (count int64, err error) {
	res := db.connection.Table("vehicles").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *VehicleConnection) FindVehiclesOffsetFull(limit int, offset int, order string, dir string) (vehicleOutput []model.SelectVehicleParameter, err error) {
	var (
		orderDirection string
		vehicles       []model.SelectVehicleParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("vehicles").Select("vehicles.id, vehicles.vehicle_hull_number, vehicles.vehicle_plate_name, vehicles.vehicle_category_id, vehicle_categories.vehicle_category_name, vehicles.vehicle_merk_id, vehicle_merks.vehicle_merk_name, vehicles.vehicle_type_id, vehicle_types.vehicle_type_name, vehicles.vehicle_registration_name, vehicles.nama_pengguna, vehicles.production_year, vehicles.company_id, companies.company_name, vehicles.chassis_number, vehicles.machine_number, vehicles.original_bpkb_storage, vehicles.bpkb_number, to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY') as tax_stnk_end_date, to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, vehicles.location_id, locations.location_name, vehicles.remark, vehicles.created_user_id, createdUID.username AS created_user, vehicles.updated_user_id, updatedUID.username AS updated_user, vehicles.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(vehicles.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at, vehicles.price").Joins("left join companies on vehicles.company_id = companies.id").Joins("left join vehicle_types ON vehicles.vehicle_type_id = vehicle_types.id").Joins("left join vehicle_categories ON vehicles.vehicle_category_id = vehicle_categories.id").Joins("left join vehicle_merks ON vehicles.vehicle_merk_id = vehicle_merks.id").Joins("left join locations ON vehicles.location_id = locations.id").Joins("left join users createdUID on vehicles.created_user_id = createdUID.id").Joins("left join users updatedUID on vehicles.updated_user_id = updatedUID.id").Joins("left join users deletedUID on vehicles.deleted_user_id = deletedUID.id").Where("vehicles.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&vehicles)
	return vehicles, res.Error
}

func (db *VehicleConnection) SearchVehicleFull(limit int, offset int, order string, dir string, search string) (vehicleOutput []model.SelectVehicleParameter, err error) {
	var (
		orderDirection string
		final          string
		vehicles       []model.SelectVehicleParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("vehicles").Select("vehicles.id, vehicles.vehicle_hull_number, vehicles.vehicle_plate_name, vehicles.vehicle_category_id, vehicle_categories.vehicle_category_name, vehicles.vehicle_merk_id, vehicle_merks.vehicle_merk_name, vehicles.vehicle_type_id, vehicle_types.vehicle_type_name, vehicles.vehicle_registration_name, vehicles.nama_pengguna, vehicles.production_year, vehicles.company_id, companies.company_name, vehicles.chassis_number, vehicles.machine_number, vehicles.original_bpkb_storage, vehicles.bpkb_number, to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY') as tax_stnk_end_date, to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, vehicles.location_id, locations.location_name, vehicles.remark, vehicles.created_user_id, createdUID.username AS created_user, vehicles.updated_user_id, updatedUID.username AS updated_user, vehicles.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(vehicles.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at, vehicles.price").Joins("left join companies on vehicles.company_id = companies.id").Joins("left join vehicle_types ON vehicles.vehicle_type_id = vehicle_types.id").Joins("left join vehicle_categories ON vehicles.vehicle_category_id = vehicle_categories.id").Joins("left join vehicle_merks ON vehicles.vehicle_merk_id = vehicle_merks.id").Joins("left join locations ON vehicles.location_id = locations.id").Joins("left join users createdUID on vehicles.created_user_id = createdUID.id").Joins("left join users updatedUID on vehicles.updated_user_id = updatedUID.id").Joins("left join users deletedUID on vehicles.deleted_user_id = deletedUID.id").Where("(lower(vehicles.vehicle_hull_number) LIKE ? OR lower(vehicles.vehicle_plate_name) LIKE ? OR lower(vehicle_categories.vehicle_category_name) LIKE ? OR lower(vehicle_merks.vehicle_merk_name) LIKE ? OR lower(vehicle_types.vehicle_type_name) LIKE ? OR lower(vehicles.vehicle_registration_name) LIKE ? OR lower(vehicles.production_year::TEXT) LIKE ? OR lower(companies.company_name) LIKE ? OR lower(vehicles.chassis_number) LIKE ? OR lower(vehicles.machine_number) LIKE ? OR lower(vehicles.original_bpkb_storage) LIKE ? OR lower(vehicles.bpkb_number) LIKE ? OR lower(to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(locations.location_name) LIKE ? OR lower(vehicles.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(vehicles.price::TEXT) LIKE ? OR lower(vehicles.nama_pengguna) LIKE ? ) AND vehicles.deleted_at = 0", final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&vehicles)
	return vehicles, res.Error
}

func (db *VehicleConnection) CountSearchVehicleFull(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("vehicles").Table("vehicles").Select("vehicles.id, vehicles.vehicle_hull_number, vehicles.vehicle_plate_name, vehicles.vehicle_category_id, vehicle_categories.vehicle_category_name, vehicles.vehicle_merk_id, vehicle_merks.vehicle_merk_name, vehicles.vehicle_type_id, vehicle_types.vehicle_type_name, vehicles.vehicle_registration_name, vehicles.nama_pengguna, vehicles.production_year, vehicles.company_id, companies.company_name, vehicles.chassis_number, vehicles.machine_number, vehicles.original_bpkb_storage, vehicles.bpkb_number, to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY') as stnk_end_date, to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY') as tax_stnk_end_date, to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY') as kir_end_date, vehicles.location_id, locations.location_name, vehicles.remark, vehicles.created_user_id, createdUID.username AS created_user, vehicles.updated_user_id, updatedUID.username AS updated_user, vehicles.deleted_user_id, deletedUID.username AS deleted_user, to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY') as created_at, to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY') as updated_at, to_char(to_timestamp(vehicles.deleted_at::numeric), 'DD-Mon-YYYY') as deleted_at, vehicles.price").Joins("left join companies on vehicles.company_id = companies.id").Joins("left join vehicle_types ON vehicles.vehicle_type_id = vehicle_types.id").Joins("left join vehicle_categories ON vehicles.vehicle_category_id = vehicle_categories.id").Joins("left join vehicle_merks ON vehicles.vehicle_merk_id = vehicle_merks.id").Joins("left join locations ON vehicles.location_id = locations.id").Joins("left join users createdUID on vehicles.created_user_id = createdUID.id").Joins("left join users updatedUID on vehicles.updated_user_id = updatedUID.id").Joins("left join users deletedUID on vehicles.deleted_user_id = deletedUID.id").Where("(lower(vehicles.vehicle_hull_number) LIKE ? OR lower(vehicles.vehicle_plate_name) LIKE ? OR lower(vehicle_categories.vehicle_category_name) LIKE ? OR lower(vehicle_merks.vehicle_merk_name) LIKE ? OR lower(vehicle_types.vehicle_type_name) LIKE ? OR lower(vehicles.vehicle_registration_name) LIKE ? OR lower(vehicles.production_year::TEXT) LIKE ? OR lower(companies.company_name) LIKE ? OR lower(vehicles.chassis_number) LIKE ? OR lower(vehicles.machine_number) LIKE ? OR lower(vehicles.original_bpkb_storage) LIKE ? OR lower(vehicles.bpkb_number) LIKE ? OR lower(to_char(to_timestamp(vehicles.stnk_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.tax_stnk_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.kir_end_date::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(locations.location_name) LIKE ? OR lower(vehicles.remark) LIKE ? OR lower(createdUID.username) LIKE ?  OR lower(updatedUID.username) LIKE ? OR lower(to_char(to_timestamp(vehicles.created_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(to_char(to_timestamp(vehicles.updated_at::numeric), 'DD-Mon-YYYY')) LIKE ? OR lower(vehicles.price::TEXT) LIKE ? OR lower(vehicles.nama_pengguna) LIKE ? ) AND vehicles.deleted_at = 0", final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final, final).Count(&count)
	return count, res.Error
}
