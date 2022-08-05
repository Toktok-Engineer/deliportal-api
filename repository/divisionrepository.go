package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type DivisionRepository interface {
	CountDivisionAll() (count int64, err error)
	FindDivisions() (divisionOutput []model.Division, err error)
	FindDivisionsOffset(limit int, offset int, order string, dir string) (divisionOutput []model.Division, err error)
	SearchDivision(limit int, offset int, order string, dir string, search string) (divisionOutput []model.Division, err error)
	CountSearchDivision(search string) (count int64, err error)
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

func (db *DivisionConnection) CountDivisionAll() (count int64, err error) {
	res := db.connection.Debug().Table("divisions").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *DivisionConnection) FindDivisions() (divisionOutput []model.Division, err error) {
	var (
		divisions []model.Division
	)
	res := db.connection.Where("deleted_at = 0").Order("division_name").Find(&divisions)
	return divisions, res.Error
}

func (db *DivisionConnection) FindDivisionsOffset(limit int, offset int, order string, dir string) (divisionOutput []model.Division, err error) {
	var (
		orderDirection string
		divisions      []model.Division
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&divisions)
	return divisions, res.Error
}

func (db *DivisionConnection) SearchDivision(limit int, offset int, order string, dir string, search string) (divisionOutput []model.Division, err error) {
	var (
		orderDirection string
		final          string
		divisions      []model.Division
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(division_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&divisions)
	return divisions, res.Error
}

func (db *DivisionConnection) CountSearchDivision(search string) (count int64, err error) {
	var (
		final string
	)
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("divisions").Where("(lower(division_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
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
