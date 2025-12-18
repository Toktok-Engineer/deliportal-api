package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type AsuransiRekeningService interface {
	CountAsuransiRekeningAll(AsuransiID int) (count int64, err error)
	FindAsuransiRekenings() (asuransiRekeningOutput []model.AsuransiRekening, err error)
	FindAsuransiRekeningsOffset(AsuransiID int, limit int, offset int, order string, dir string) (asuransiRekeningOutput []model.AsuransiRekening, err error)
	SearchAsuransiRekening(AsuransiID int, limit int, offset int, order string, dir string, search string) (asuransiRekeningOutput []model.AsuransiRekening, err error)
	CountSearchAsuransiRekening(AsuransiID int, search string) (count int64, err error)
	FindAsuransiRekeningById(id uint) (asuransiRekeningOutput model.AsuransiRekening, err error)
	FindExcAsuransiRekening(id uint) (asuransiRekeningOutput []model.AsuransiRekening, err error)
	InsertAsuransiRekening(asuransiRekening model.CreateAsuransiRekeningParameter) (asuransiRekeningOutput model.AsuransiRekening, err error)
	UpdateAsuransiRekening(asuransiRekening model.AsuransiRekening, id uint) (asuransiRekeningOutput model.AsuransiRekening, err error)
	DeleteAsuransiRekening(asuransiRekening model.AsuransiRekening, id uint) (asuransiRekeningOutput model.AsuransiRekening, err error)
}

type asuransiRekeningService struct {
	asuransiRekeningRepository repository.AsuransiRekeningRepository
}

func NewAsuransiRekeningService(asuransiRekeningRep repository.AsuransiRekeningRepository) AsuransiRekeningService {
	return &asuransiRekeningService{
		asuransiRekeningRepository: asuransiRekeningRep,
	}
}

func (service *asuransiRekeningService) CountAsuransiRekeningAll(AsuransiID int) (count int64, err error) {
	res, err := service.asuransiRekeningRepository.CountAsuransiRekeningAll(AsuransiID)
	return res, err
}

func (service *asuransiRekeningService) FindAsuransiRekenings() (asuransiRekeningOutput []model.AsuransiRekening, err error) {
	res, err := service.asuransiRekeningRepository.FindAsuransiRekenings()
	return res, err
}

func (service *asuransiRekeningService) FindAsuransiRekeningsOffset(AsuransiID int, limit int, offset int, order string, dir string) (asuransiRekeningOutput []model.AsuransiRekening, err error) {
	res, err := service.asuransiRekeningRepository.FindAsuransiRekeningsOffset(AsuransiID, limit, offset, order, dir)
	return res, err
}

func (service *asuransiRekeningService) SearchAsuransiRekening(AsuransiID int, limit int, offset int, order string, dir string, search string) (asuransiRekeningOutput []model.AsuransiRekening, err error) {
	res, err := service.asuransiRekeningRepository.SearchAsuransiRekening(AsuransiID, limit, offset, order, dir, search)
	return res, err
}

func (service *asuransiRekeningService) CountSearchAsuransiRekening(AsuransiID int, search string) (count int64, err error) {
	res, err := service.asuransiRekeningRepository.CountSearchAsuransiRekening(AsuransiID, search)
	return res, err
}

func (service *asuransiRekeningService) FindAsuransiRekeningById(id uint) (asuransiRekeningOutput model.AsuransiRekening, err error) {
	return service.asuransiRekeningRepository.FindAsuransiRekeningById(id)
}

func (service *asuransiRekeningService) FindExcAsuransiRekening(id uint) (asuransiRekeningOutput []model.AsuransiRekening, err error) {
	return service.asuransiRekeningRepository.FindExcAsuransiRekening(id)
}
func (service *asuransiRekeningService) InsertAsuransiRekening(asuransiRekening model.CreateAsuransiRekeningParameter) (model.AsuransiRekening, error) {
	newAsuransiRekening := model.AsuransiRekening{}
	err1 := smapping.FillStruct(&newAsuransiRekening, smapping.MapFields(&asuransiRekening))

	if err1 != nil {
		return newAsuransiRekening, err1
	}

	return service.asuransiRekeningRepository.InsertAsuransiRekening(newAsuransiRekening)
}

func (service *asuransiRekeningService) UpdateAsuransiRekening(asuransiRekening model.AsuransiRekening, id uint) (asuransiRekeningOutput model.AsuransiRekening, err error) {
	newAsuransiRekening := model.AsuransiRekening{}
	err1 := smapping.FillStruct(&newAsuransiRekening, smapping.MapFields(&asuransiRekening))

	if err1 != nil {
		return newAsuransiRekening, err1
	}

	return service.asuransiRekeningRepository.UpdateAsuransiRekening(newAsuransiRekening, id)
}

func (service *asuransiRekeningService) DeleteAsuransiRekening(asuransiRekening model.AsuransiRekening, id uint) (asuransiRekeningOutput model.AsuransiRekening, err error) {
	newAsuransiRekening := model.AsuransiRekening{}
	err1 := smapping.FillStruct(&newAsuransiRekening, smapping.MapFields(&asuransiRekening))

	if err1 != nil {
		return newAsuransiRekening, err1
	}

	return service.asuransiRekeningRepository.UpdateAsuransiRekening(newAsuransiRekening, id)
}
