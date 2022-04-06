package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type DivisionRepository interface {
	FindDivisions() []model.Division
	FindDivisionById(id uint) model.Division
	FindExcDivision(id uint) []model.Division
	InsertDivision(division model.Division) model.Division
	UpdateDivision(division model.Division) model.Division
}

type DivisionConnection struct {
	connection *gorm.DB
}

func NewDivisionRepository(db *gorm.DB) DivisionRepository {
	return &DivisionConnection{
		connection: db,
	}
}

func (db *DivisionConnection) FindDivisions() []model.Division {
	var divisions []model.Division
	db.connection.Where("deleted_at = 0").Order("division_name").Find(&divisions)
	return divisions
}

func (db *DivisionConnection) FindDivisionById(id uint) model.Division {
	var division model.Division
	db.connection.Where("id=? AND deleted_at = 0", id).Take(&division)
	return division
}

func (db *DivisionConnection) FindExcDivision(id uint) []model.Division {
	var divisions []model.Division
	db.connection.Where("id!=? AND deleted_at = 0", id).Order("division_name").Find(&divisions)
	return divisions
}

func (db *DivisionConnection) InsertDivision(division model.Division) model.Division {
	db.connection.Save(&division)
	return division
}

func (db *DivisionConnection) UpdateDivision(division model.Division) model.Division {
	db.connection.Save(&division)
	return division
}
