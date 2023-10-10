package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type BusinessUnitRepository interface {
	CountBusinessUnitAll() (count int64, err error)
	FindBusinessUnits() (businessunitOutput []model.BusinessUnit, err error)
	FindBusinessUnitsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.BusinessUnit, err error)
	SearchBusinessUnit(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.BusinessUnit, err error)
	CountSearchBusinessUnit(search string) (count int64, err error)
	FindBusinessUnitById(id uint) (businessUnitOutput model.BusinessUnit, err error)
	FindExcBusinessUnit(id uint) (businessUnitOutput []model.BusinessUnit, err error)
	InsertBusinessUnit(businessUnit model.BusinessUnit) (businessUnitOutput model.BusinessUnit, err error)
	UpdateBusinessUnit(businessUnit model.BusinessUnit, id uint) (businessUnitOutput model.BusinessUnit, err error)
}

type businessUnitConnection struct {
	connection *gorm.DB
}

func NewBusinessUnitRepository(db *gorm.DB) BusinessUnitRepository {
	return &businessUnitConnection{
		connection: db,
	}
}
func (db *businessUnitConnection) CountBusinessUnitAll() (count int64, err error) {
	res := db.connection.Debug().Table("business_units").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *businessUnitConnection) FindBusinessUnits() (businessunitOutput []model.BusinessUnit, err error) {
	var (
		businessunits []model.BusinessUnit
	)
	res := db.connection.Where("deleted_at = 0").Order("business_unit_name").Find(&businessunits)
	return businessunits, res.Error
}

func (db *businessUnitConnection) FindBusinessUnitsOffset(limit int, offset int, order string, dir string) (businessunitOutput []model.BusinessUnit, err error) {
	var (
		orderDirection string
		businessunits  []model.BusinessUnit
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&businessunits)
	return businessunits, res.Error
}

func (db *businessUnitConnection) SearchBusinessUnit(limit int, offset int, order string, dir string, search string) (businessunitOutput []model.BusinessUnit, err error) {
	var (
		orderDirection string
		final          string
		businessunits  []model.BusinessUnit
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(business_unit_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&businessunits)
	return businessunits, res.Error
}

func (db *businessUnitConnection) CountSearchBusinessUnit(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("business_units").Where("(lower(business_unit_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
}

func (db *businessUnitConnection) FindBusinessUnitById(id uint) (businessUnitOutput model.BusinessUnit, err error) {
	var (
		businessUnit model.BusinessUnit
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&businessUnit)
	return businessUnit, res.Error
}

func (db *businessUnitConnection) FindExcBusinessUnit(id uint) (businessUnitOutput []model.BusinessUnit, err error) {
	var (
		businessUnits []model.BusinessUnit
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("business_unit_name").Find(&businessUnits)
	return businessUnits, res.Error
}

func (db *businessUnitConnection) InsertBusinessUnit(businessUnit model.BusinessUnit) (businessUnitOutput model.BusinessUnit, err error) {
	res := db.connection.Save(&businessUnit)
	return businessUnit, res.Error
}

func (db *businessUnitConnection) UpdateBusinessUnit(businessUnit model.BusinessUnit, id uint) (businessUnitOutput model.BusinessUnit, err error) {
	res := db.connection.Where("id=?", id).Updates(&businessUnit)
	return businessUnit, res.Error
}
