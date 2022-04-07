package repository

import (
	"deliportal-api/model"
	"log"

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
	var divisions []model.Division
	res := db.connection.Where("deleted_at = 0").Order("division_name").Find(&divisions)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return divisions, res.Error
	}
	return divisions, nil
}

func (db *DivisionConnection) FindDivisionById(id uint) (divisionOutput model.Division, err error) {
	var division model.Division
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&division)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return division, res.Error
	}
	return division, nil
}

func (db *DivisionConnection) FindExcDivision(id uint) (divisionOutput []model.Division, err error) {
	var divisions []model.Division
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("division_name").Find(&divisions)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return divisions, res.Error
	}
	return divisions, nil
}

func (db *DivisionConnection) InsertDivision(division model.Division) (divisionOutput model.Division, err error) {
	res := db.connection.Save(&division)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return division, res.Error
	}
	return division, nil
}

func (db *DivisionConnection) UpdateDivision(division model.Division, id uint) (divisionOutput model.Division, err error) {
	res := db.connection.Where("id=?", id).Updates(&division)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return division, res.Error
	}
	return division, nil
}
