package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type BusinessUnitRepository interface {
	FindBusinessUnits() ([]model.BusinessUnit, error)
	FindBusinessUnitById(id uint) (model.BusinessUnit, error)
	InsertBusinessUnit(businessUnit model.BusinessUnit) (model.BusinessUnit, error)
	UpdateBusinessUnit(businessUnit model.BusinessUnit) (model.BusinessUnit, error)
}

type businessUnitConnection struct {
	connection *gorm.DB
}

func NewBusinessUnitRepository(db *gorm.DB) BusinessUnitRepository {
	return &businessUnitConnection{
		connection: db,
	}
}

func (db *businessUnitConnection) FindBusinessUnits() ([]model.BusinessUnit, error) {
	var businessUnits []model.BusinessUnit
	res := db.connection.Order("id").Find(&businessUnits)
	return businessUnits, res.Error
}

func (db *businessUnitConnection) FindBusinessUnitById(id uint) (model.BusinessUnit, error) {
	var businessUnit model.BusinessUnit
	res := db.connection.Where("id=?", id).Take(&businessUnit)
	return businessUnit, res.Error
}

func (db *businessUnitConnection) InsertBusinessUnit(businessUnit model.BusinessUnit) (model.BusinessUnit, error) {
	res := db.connection.Save(&businessUnit)
	return businessUnit, res.Error
}

func (db *businessUnitConnection) UpdateBusinessUnit(businessUnit model.BusinessUnit) (model.BusinessUnit, error) {
	res := db.connection.Where("id = ?", businessUnit.ID).Updates(&businessUnit)
	return businessUnit, res.Error
}
