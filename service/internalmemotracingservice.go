package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type InternalMemoTracingService interface {
	CountInternalMemoTracingAll(internalMemoId int) (count int64, err error)
	FindInternalMemoTracings(internalMemoId int) (internalmemotracingOutput []model.InternalMemoTracing, err error)
	FindInternalMemoTracingById(id uint) (internalMemoTracingOutput model.InternalMemoTracing, err error)
	FindExcInternalMemoTracing(id uint) (internalMemoTracingOutput []model.InternalMemoTracing, err error)
	InsertInternalMemoTracing(internalMemoTracing model.CreateInternalMemoTracingParameter) (internalMemoTracingOutput model.InternalMemoTracing, err error)
	UpdateInternalMemoTracing(internalMemoTracing model.InternalMemoTracing, id uint) (internalMemoTracingOutput model.InternalMemoTracing, err error)
	DeleteInternalMemoTracing(internalMemoTracing model.InternalMemoTracing, id uint) (internalMemoTracingOutput model.InternalMemoTracing, err error)
}

type internalMemoTracingService struct {
	internalMemoTracingRepository repository.InternalMemoTracingRepository
}

func NewInternalMemoTracingService(internalMemoTracingRep repository.InternalMemoTracingRepository) InternalMemoTracingService {
	return &internalMemoTracingService{
		internalMemoTracingRepository: internalMemoTracingRep,
	}
}

func (service *internalMemoTracingService) CountInternalMemoTracingAll(internalMemoId int) (count int64, err error) {
	res, err := service.internalMemoTracingRepository.CountInternalMemoTracingAll(internalMemoId)
	return res, err
}

func (service *internalMemoTracingService) FindInternalMemoTracings(internalMemoId int) (internalmemotracingOutput []model.InternalMemoTracing, err error) {
	res, err := service.internalMemoTracingRepository.FindInternalMemoTracings(internalMemoId)
	return res, err
}

func (service *internalMemoTracingService) FindInternalMemoTracingById(id uint) (internalMemoTracingOutput model.InternalMemoTracing, err error) {
	return service.internalMemoTracingRepository.FindInternalMemoTracingById(id)
}

func (service *internalMemoTracingService) FindExcInternalMemoTracing(id uint) (internalMemoTracingOutput []model.InternalMemoTracing, err error) {
	return service.internalMemoTracingRepository.FindExcInternalMemoTracing(id)
}
func (service *internalMemoTracingService) InsertInternalMemoTracing(internalMemoTracing model.CreateInternalMemoTracingParameter) (model.InternalMemoTracing, error) {
	newInternalMemoTracing := model.InternalMemoTracing{}
	err1 := smapping.FillStruct(&newInternalMemoTracing, smapping.MapFields(&internalMemoTracing))

	if err1 != nil {
		return newInternalMemoTracing, err1
	}

	return service.internalMemoTracingRepository.InsertInternalMemoTracing(newInternalMemoTracing)
}

func (service *internalMemoTracingService) UpdateInternalMemoTracing(internalMemoTracing model.InternalMemoTracing, id uint) (internalMemoTracingOutput model.InternalMemoTracing, err error) {
	newInternalMemoTracing := model.InternalMemoTracing{}
	err1 := smapping.FillStruct(&newInternalMemoTracing, smapping.MapFields(&internalMemoTracing))

	if err1 != nil {
		return newInternalMemoTracing, err1
	}

	return service.internalMemoTracingRepository.UpdateInternalMemoTracing(newInternalMemoTracing, id)
}

func (service *internalMemoTracingService) DeleteInternalMemoTracing(internalMemoTracing model.InternalMemoTracing, id uint) (internalMemoTracingOutput model.InternalMemoTracing, err error) {
	newInternalMemoTracing := model.InternalMemoTracing{}
	err1 := smapping.FillStruct(&newInternalMemoTracing, smapping.MapFields(&internalMemoTracing))

	if err1 != nil {
		return newInternalMemoTracing, err1
	}

	return service.internalMemoTracingRepository.UpdateInternalMemoTracing(newInternalMemoTracing, id)
}
