package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type PositionService interface {
	FindPositions() (positionOutput []model.Position, err error)
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

func (service *positionService) FindPositions() (positionOutput []model.Position, err error) {
	res, err := service.positionRepository.FindPositions()
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
	if err1 != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.positionRepository.InsertPosition(newPosition)
	return res, err
}

func (service *positionService) UpdatePosition(position model.Position, id uint) (positionOutput model.Position, err error) {
	newPosition := model.Position{}
	err1 := smapping.FillStruct(&newPosition, smapping.MapFields(&position))
	if err1 != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.positionRepository.UpdatePosition(newPosition, id)
	return res, err
}

func (service *positionService) DeletePosition(position model.Position, id uint) (positionOutput model.Position, err error) {
	newPosition := model.Position{}
	err1 := smapping.FillStruct(&newPosition, smapping.MapFields(&position))
	if err1 != nil {
		log.Fatalf("Failed map %v", err1)
	}
	res, err := service.positionRepository.UpdatePosition(newPosition, id)
	return res, err
}
