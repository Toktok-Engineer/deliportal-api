package repository

import (
	"deliportal-api/model"

	"gorm.io/gorm"
)

type PositionRepository interface {
	FindPositions() (positionOutput []model.Position, err error)
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

func (db *PositionConnection) FindPositions() (positionOutput []model.Position, err error) {
	var (
		positions []model.Position
	)
	res := db.connection.Where("deleted_at = 0").Order("position_name").Find(&positions)
	return positions, res.Error
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
