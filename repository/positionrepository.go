package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type PositionRepository interface {
	CountPositionAll() (count int64, err error)
	FindPositions() (positionOutput []model.Position, err error)
	FindPositionsOffset(limit int, offset int, order string, dir string) (positionOutput []model.Position, err error)
	SearchPosition(limit int, offset int, order string, dir string, search string) (positionOutput []model.Position, err error)
	CountSearchPosition(search string) (count int64, err error)
	FindPositionById(id uint) (positionOutput model.Position, err error)
	FindExcPosition(id uint) (positionOutput []model.Position, err error)
	InsertPosition(position model.Position) (positionOutput model.Position, err error)
	UpdatePosition(position model.Position, id uint) (positionOutput model.Position, err error)
}

type PositionConnection struct {
	connection *gorm.DB
}

func NewPositionRepository(db *gorm.DB) PositionRepository {
	return &PositionConnection{
		connection: db,
	}
}

func (db *PositionConnection) CountPositionAll() (count int64, err error) {
	res := db.connection.Table("positions").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *PositionConnection) FindPositions() (positionOutput []model.Position, err error) {
	var (
		positions []model.Position
	)
	res := db.connection.Where("deleted_at = 0").Order("position_name").Find(&positions)
	return positions, res.Error
}

func (db *PositionConnection) FindPositionsOffset(limit int, offset int, order string, dir string) (positionOutput []model.Position, err error) {
	var (
		orderDirection string
		positions      []model.Position
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&positions)
	return positions, res.Error
}

func (db *PositionConnection) SearchPosition(limit int, offset int, order string, dir string, search string) (positionOutput []model.Position, err error) {
	var (
		orderDirection string
		final          string
		positions      []model.Position
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(position_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&positions)
	return positions, res.Error
}

func (db *PositionConnection) CountSearchPosition(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("positions").Where("(lower(position_name) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
	return count, res.Error
}

func (db *PositionConnection) FindPositionById(id uint) (positionOutput model.Position, err error) {
	var (
		position model.Position
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&position)
	return position, res.Error
}

func (db *PositionConnection) FindExcPosition(id uint) (positionOutput []model.Position, err error) {
	var (
		positions []model.Position
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("position_name").Find(&positions)
	return positions, res.Error
}

func (db *PositionConnection) InsertPosition(position model.Position) (positionOutput model.Position, err error) {
	res := db.connection.Save(&position)
	return position, res.Error
}

func (db *PositionConnection) UpdatePosition(position model.Position, id uint) (positionOutput model.Position, err error) {
	res := db.connection.Where("id=?", id).Updates(&position)
	return position, res.Error
}
