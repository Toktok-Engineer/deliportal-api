package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type BusinessUnitRepository interface {
	FindBusinessUnits() (businessUnitOutput []model.BusinessUnit, err error)
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

func (db *businessUnitConnection) FindBusinessUnits() (businessUnitOutput []model.BusinessUnit, err error) {
	var (
		businessUnits []model.BusinessUnit
	)
	res := db.connection.Where("deleted_at = 0").Order("business_unit_name").Find(&businessUnits)
	return businessUnits, res.Error
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
