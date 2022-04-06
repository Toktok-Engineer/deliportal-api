package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type BusinessUnitRepository interface {
	FindBusinessUnits() []model.BusinessUnit
	FindBusinessUnitById(id uint) model.BusinessUnit
	InsertBusinessUnit(businessUnit model.BusinessUnit) model.BusinessUnit
	UpdateBusinessUnit(businessUnit model.BusinessUnit) model.BusinessUnit
}

type businessUnitConnection struct {
	connection *gorm.DB
}

func NewBusinessUnitRepository(db *gorm.DB) BusinessUnitRepository {
	return &businessUnitConnection{
		connection: db,
	}
}

func (db *businessUnitConnection) FindBusinessUnits() []model.BusinessUnit {
	var businessUnits []model.BusinessUnit
	db.connection.Order("id").Find(&businessUnits)
	return businessUnits
}

func (db *businessUnitConnection) FindBusinessUnitById(id uint) model.BusinessUnit {
	var businessUnit model.BusinessUnit
	db.connection.Where("id=?", id).Take(&businessUnit)
	return businessUnit
}

func (db *businessUnitConnection) InsertBusinessUnit(businessUnit model.BusinessUnit) model.BusinessUnit {
	db.connection.Save(&businessUnit)
	return businessUnit
}

func (db *businessUnitConnection) UpdateBusinessUnit(businessUnit model.BusinessUnit) model.BusinessUnit {
	db.connection.Save(&businessUnit)
	return businessUnit
}
