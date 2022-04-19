package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type DivisionRepository interface {
	FindDivisions() (divisionOutput []model.Division, err error)
	FindDivisionById(id uint) (divisionOutput model.Division, err error)
	FindExcDivision(id uint) (divisionOutput []model.Division, err error)
	InsertDivision(division model.Division) (divisionOutput model.Division, err error)
	UpdateDivision(division model.Division, id uint) (divisionOutput model.Division, err error)
}

type DivisionConnection struct {
	connection *gorm.DB
}

func NewDivisionRepository(db *gorm.DB) DivisionRepository {
	return &DivisionConnection{
		connection: db,
	}
}

func (db *DivisionConnection) FindDivisions() (divisionOutput []model.Division, err error) {
	var (
		divisions []model.Division
	)
	res := db.connection.Where("deleted_at = 0").Order("division_name").Find(&divisions)
	return divisions, res.Error
}

func (db *DivisionConnection) FindDivisionById(id uint) (divisionOutput model.Division, err error) {
	var (
		division model.Division
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&division)
	return division, res.Error
}

func (db *DivisionConnection) FindExcDivision(id uint) (divisionOutput []model.Division, err error) {
	var (
		divisions []model.Division
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("division_name").Find(&divisions)
	return divisions, res.Error
}

func (db *DivisionConnection) InsertDivision(division model.Division) (divisionOutput model.Division, err error) {
	res := db.connection.Save(&division)
	return division, res.Error
}

func (db *DivisionConnection) UpdateDivision(division model.Division, id uint) (divisionOutput model.Division, err error) {
	res := db.connection.Where("id=?", id).Updates(&division)
	return division, res.Error
}
