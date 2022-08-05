package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type PositionService interface {
	CountPositionAll() (count int64, err error)
	FindPositions() (positionOutput []model.Position, err error)
	FindPositionsOffset(limit int, offset int, order string, dir string) (positionOutput []model.Position, err error)
	SearchPosition(limit int, offset int, order string, dir string, search string) (positionOutput []model.Position, err error)
	CountSearchPosition(search string) (count int64, err error)
	FindPositionById(id uint) (positionOutput model.Position, err error)
	FindExcPosition(id uint) (positionOutput []model.Position, err error)
	InsertPosition(position model.CreatePositionParameter) (positionOutput model.Position, err error)
	UpdatePosition(position model.Position, id uint) (positionOutput model.Position, err error)
	DeletePosition(position model.Position, id uint) (positionOutput model.Position, err error)
}

type positionService struct {
	positionRepository repository.PositionRepository
}

func NewPositionService(positionRep repository.PositionRepository) PositionService {
	return &positionService{
		positionRepository: positionRep,
	}
}

func (service *positionService) CountPositionAll() (count int64, err error) {
	res, err := service.positionRepository.CountPositionAll()
	return res, err
}

func (service *positionService) FindPositions() (positionOutput []model.Position, err error) {
	res, err := service.positionRepository.FindPositions()
	return res, err
}

func (service *positionService) FindPositionsOffset(limit int, offset int, order string, dir string) (positionOutput []model.Position, err error) {
	res, err := service.positionRepository.FindPositionsOffset(limit, offset, order, dir)
	return res, err
}

func (service *positionService) SearchPosition(limit int, offset int, order string, dir string, search string) (positionOutput []model.Position, err error) {
	res, err := service.positionRepository.SearchPosition(limit, offset, order, dir, search)
	return res, err
}

func (service *positionService) CountSearchPosition(search string) (count int64, err error) {
	res, err := service.positionRepository.CountSearchPosition(search)
	return res, err
}

func (service *positionService) FindPositionById(id uint) (positionOutput model.Position, err error) {
	res, err := service.positionRepository.FindPositionById(id)
	return res, err
}

func (service *positionService) FindExcPosition(id uint) (positionOutput []model.Position, err error) {
	res, err := service.positionRepository.FindExcPosition(id)
	return res, err
}

func (service *positionService) InsertPosition(position model.CreatePositionParameter) (positionOutput model.Position, err error) {
	newPosition := model.Position{}
	err1 := smapping.FillStruct(&newPosition, smapping.MapFields(&position))
	if err != nil {
		return newPosition, err1
	}
	res, err := service.positionRepository.InsertPosition(newPosition)
	return res, err
}

func (service *positionService) UpdatePosition(position model.Position, id uint) (positionOutput model.Position, err error) {
	newPosition := model.Position{}
	err1 := smapping.FillStruct(&newPosition, smapping.MapFields(&position))
	if err != nil {
		return newPosition, err1
	}
	res, err := service.positionRepository.UpdatePosition(newPosition, id)
	return res, err
}

func (service *positionService) DeletePosition(position model.Position, id uint) (positionOutput model.Position, err error) {
	newPosition := model.Position{}
	err1 := smapping.FillStruct(&newPosition, smapping.MapFields(&position))
	if err != nil {
		return newPosition, err1
	}
	res, err := service.positionRepository.UpdatePosition(newPosition, id)
	return res, err
}
