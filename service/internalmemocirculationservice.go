package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type InternalMemoCirculationService interface {
	CountInternalMemoCirculationAll(internalMemoId int) (count int64, err error)
	FindInternalMemoCirculations(internalMemoId int) (internalMemoCirculationOutput []model.InternalMemoCirculation, err error)
	FindInternalMemoCirculationStat(internalMemoId int, cirStat int) (internalMemoCirculationOutput model.SelectInternalMemoCirculationParameter, err error)
	FindInternalMemoCirculationListStat(internalMemoId int, cirStat int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationParameter, err error)
	FindInternalMemoCirculationSeq(internalMemoId int, employeeID int) (internalMemoCirculationOutput model.SelectInternalMemoCirculationParameter, err error)
	FindInternalMemoCirculationsOffset(limit int, offset int, order string, dir string, internalMemoId int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationParameter, err error)
	SearchInternalMemoCirculation(limit int, offset int, order string, dir string, search string, internalMemoId int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationParameter, err error)
	CountSearchInternalMemoCirculation(search string, internalMemoId int) (count int64, err error)
	FindInternalMemoCirculationById(id uint) (internalMemoCirculationOutput model.SelectInternalMemoCirculationParameter, err error)
	FindExcInternalMemoCirculation(id uint) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationParameter, err error)
	FindInternalMemoCirculationByCompanyId(id uint) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationParameter, err error)
	InsertInternalMemoCirculation(internalMemoCirculation model.CreateInternalMemoCirculationParameter) (internalMemoCirculationOutput model.InternalMemoCirculation, err error)
	UpdateInternalMemoCirculation(internalMemoCirculation model.InternalMemoCirculation, id uint) (internalMemoCirculationOutput model.InternalMemoCirculation, err error)
	DeleteInternalMemoCirculation(internalMemoCirculation model.InternalMemoCirculation, id uint) (internalMemoCirculationOutput model.InternalMemoCirculation, err error)
	UpdateInternalMemoCirculationApprove(internalMemoCirculation model.InternalMemoCirculation, id uint) (internalMemoCirculationOutput model.InternalMemoCirculation, err error)
	CountInternalMemoCirculationJoinIM(employeeID int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationJoinIMParameter, err error)
	FindInternalMemoCirculationsJoinIMOffset(limit int, offset int, order string, dir string, employeeID int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationJoinIMParameter, err error)
	SearchInternalMemoCirculationJoinIM(limit int, offset int, order string, dir string, search string, employeeID int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationJoinIMParameter, err error)
	CountSearchInternalMemoCirculationJoinIM(search string, employeeID int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationJoinIMParameter, err error)
}

type internalMemoCirculationService struct {
	internalMemoCirculationRepository repository.InternalMemoCirculationRepository
}

func NewInternalMemoCirculationService(internalMemoCirculationRep repository.InternalMemoCirculationRepository) InternalMemoCirculationService {
	return &internalMemoCirculationService{
		internalMemoCirculationRepository: internalMemoCirculationRep,
	}
}

func (service *internalMemoCirculationService) CountInternalMemoCirculationAll(internalMemoId int) (count int64, err error) {
	res, err := service.internalMemoCirculationRepository.CountInternalMemoCirculationAll(internalMemoId)
	return res, err
}

func (service *internalMemoCirculationService) FindInternalMemoCirculations(internalMemoId int) (internalMemoCirculationOutput []model.InternalMemoCirculation, err error) {
	res, err := service.internalMemoCirculationRepository.FindInternalMemoCirculations(internalMemoId)
	return res, err
}

func (service *internalMemoCirculationService) FindInternalMemoCirculationStat(internalMemoId int, cirStat int) (internalMemoCirculationOutput model.SelectInternalMemoCirculationParameter, err error) {
	res, err := service.internalMemoCirculationRepository.FindInternalMemoCirculationStat(internalMemoId, cirStat)
	return res, err
}

func (service *internalMemoCirculationService) FindInternalMemoCirculationListStat(internalMemoId int, cirStat int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationParameter, err error) {
	res, err := service.internalMemoCirculationRepository.FindInternalMemoCirculationListStat(internalMemoId, cirStat)
	return res, err
}

func (service *internalMemoCirculationService) FindInternalMemoCirculationSeq(internalMemoId int, employeeID int) (internalMemoCirculationOutput model.SelectInternalMemoCirculationParameter, err error) {
	res, err := service.internalMemoCirculationRepository.FindInternalMemoCirculationSeq(internalMemoId, employeeID)
	return res, err
}

func (service *internalMemoCirculationService) FindInternalMemoCirculationsOffset(limit int, offset int, order string, dir string, internalMemoId int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationParameter, err error) {
	res, err := service.internalMemoCirculationRepository.FindInternalMemoCirculationsOffset(limit, offset, order, dir, internalMemoId)
	return res, err
}

func (service *internalMemoCirculationService) SearchInternalMemoCirculation(limit int, offset int, order string, dir string, search string, internalMemoId int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationParameter, err error) {
	res, err := service.internalMemoCirculationRepository.SearchInternalMemoCirculation(limit, offset, order, dir, search, internalMemoId)
	return res, err
}

func (service *internalMemoCirculationService) CountSearchInternalMemoCirculation(search string, internalMemoId int) (count int64, err error) {
	res, err := service.internalMemoCirculationRepository.CountSearchInternalMemoCirculation(search, internalMemoId)
	return res, err
}

func (service *internalMemoCirculationService) FindInternalMemoCirculationById(id uint) (internalMemoCirculationOutput model.SelectInternalMemoCirculationParameter, err error) {
	res, err := service.internalMemoCirculationRepository.FindInternalMemoCirculationById(id)
	return res, err
}

func (service *internalMemoCirculationService) FindExcInternalMemoCirculation(id uint) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationParameter, err error) {
	res, err := service.internalMemoCirculationRepository.FindExcInternalMemoCirculation(id)
	return res, err
}

func (service *internalMemoCirculationService) FindInternalMemoCirculationByCompanyId(id uint) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationParameter, err error) {
	res, err := service.internalMemoCirculationRepository.FindInternalMemoCirculationByCompanyId(id)
	return res, err
}

func (service *internalMemoCirculationService) InsertInternalMemoCirculation(internalMemoCirculation model.CreateInternalMemoCirculationParameter) (internalMemoCirculationOutput model.InternalMemoCirculation, err error) {
	newInternalMemoCirculation := model.InternalMemoCirculation{}
	err1 := smapping.FillStruct(&newInternalMemoCirculation, smapping.MapFields(&internalMemoCirculation))
	if err != nil {
		return newInternalMemoCirculation, err1
	}
	res, err := service.internalMemoCirculationRepository.InsertInternalMemoCirculation(newInternalMemoCirculation)
	return res, err
}

func (service *internalMemoCirculationService) UpdateInternalMemoCirculation(internalMemoCirculation model.InternalMemoCirculation, id uint) (internalMemoCirculationOutput model.InternalMemoCirculation, err error) {
	newInternalMemoCirculation := model.InternalMemoCirculation{}
	err1 := smapping.FillStruct(&newInternalMemoCirculation, smapping.MapFields(&internalMemoCirculation))
	if err != nil {
		return newInternalMemoCirculation, err1
	}
	res, err := service.internalMemoCirculationRepository.UpdateInternalMemoCirculation(newInternalMemoCirculation, id)
	return res, err
}

func (service *internalMemoCirculationService) DeleteInternalMemoCirculation(internalMemoCirculation model.InternalMemoCirculation, id uint) (internalMemoCirculationOutput model.InternalMemoCirculation, err error) {
	newInternalMemoCirculation := model.InternalMemoCirculation{}
	err1 := smapping.FillStruct(&newInternalMemoCirculation, smapping.MapFields(&internalMemoCirculation))
	if err != nil {
		return newInternalMemoCirculation, err1
	}
	res, err := service.internalMemoCirculationRepository.UpdateInternalMemoCirculation(newInternalMemoCirculation, id)
	return res, err
}

func (service *internalMemoCirculationService) UpdateInternalMemoCirculationApprove(internalMemoCirculation model.InternalMemoCirculation, id uint) (internalMemoCirculationOutput model.InternalMemoCirculation, err error) {
	newInternalMemoCirculation := model.InternalMemoCirculation{}
	err1 := smapping.FillStruct(&newInternalMemoCirculation, smapping.MapFields(&internalMemoCirculation))
	if err != nil {
		return newInternalMemoCirculation, err1
	}
	res, err := service.internalMemoCirculationRepository.UpdateInternalMemoCirculationApprove(newInternalMemoCirculation, id)
	return res, err
}

func (service *internalMemoCirculationService) CountInternalMemoCirculationJoinIM(employeeID int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationJoinIMParameter, err error) {
	res, err := service.internalMemoCirculationRepository.CountInternalMemoCirculationJoinIM(employeeID)
	return res, err
}

func (service *internalMemoCirculationService) FindInternalMemoCirculationsJoinIMOffset(limit int, offset int, order string, dir string, employeeID int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationJoinIMParameter, err error) {
	res, err := service.internalMemoCirculationRepository.FindInternalMemoCirculationsJoinIMOffset(limit, offset, order, dir, employeeID)
	return res, err
}

func (service *internalMemoCirculationService) SearchInternalMemoCirculationJoinIM(limit int, offset int, order string, dir string, search string, employeeID int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationJoinIMParameter, err error) {
	res, err := service.internalMemoCirculationRepository.SearchInternalMemoCirculationJoinIM(limit, offset, order, dir, search, employeeID)
	return res, err
}

func (service *internalMemoCirculationService) CountSearchInternalMemoCirculationJoinIM(search string, employeeID int) (internalMemoCirculationOutput []model.SelectInternalMemoCirculationJoinIMParameter, err error) {
	res, err := service.internalMemoCirculationRepository.CountSearchInternalMemoCirculationJoinIM(search, employeeID)
	return res, err
}
