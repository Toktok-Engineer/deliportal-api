package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type AsuransiService interface {
	CountAsuransiAll() (count int64, err error)
	FindAsuransis() (asuransiOutput []model.Asuransi, err error)
	FindAsuransisOffset(limit int, offset int, order string, dir string) (asuransiOutput []model.Asuransi, err error)
	SearchAsuransi(limit int, offset int, order string, dir string, search string) (asuransiOutput []model.Asuransi, err error)
	CountSearchAsuransi(search string) (count int64, err error)
	FindAsuransiById(id uint) (asuransiOutput model.Asuransi, err error)
	FindExcAsuransi(id uint) (asuransiOutput []model.Asuransi, err error)
	InsertAsuransi(asuransi model.CreateAsuransiParameter) (asuransiOutput model.Asuransi, err error)
	UpdateAsuransi(asuransi model.Asuransi, id uint) (asuransiOutput model.Asuransi, err error)
	DeleteAsuransi(asuransi model.Asuransi, id uint) (asuransiOutput model.Asuransi, err error)
}

type asuransiService struct {
	asuransiRepository repository.AsuransiRepository
}

func NewAsuransiService(asuransiRep repository.AsuransiRepository) AsuransiService {
	return &asuransiService{
		asuransiRepository: asuransiRep,
	}
}

func (service *asuransiService) CountAsuransiAll() (count int64, err error) {
	res, err := service.asuransiRepository.CountAsuransiAll()
	return res, err
}

func (service *asuransiService) FindAsuransis() (asuransiOutput []model.Asuransi, err error) {
	res, err := service.asuransiRepository.FindAsuransis()
	return res, err
}

func (service *asuransiService) FindAsuransisOffset(limit int, offset int, order string, dir string) (asuransiOutput []model.Asuransi, err error) {
	res, err := service.asuransiRepository.FindAsuransisOffset(limit, offset, order, dir)
	return res, err
}

func (service *asuransiService) SearchAsuransi(limit int, offset int, order string, dir string, search string) (asuransiOutput []model.Asuransi, err error) {
	res, err := service.asuransiRepository.SearchAsuransi(limit, offset, order, dir, search)
	return res, err
}

func (service *asuransiService) CountSearchAsuransi(search string) (count int64, err error) {
	res, err := service.asuransiRepository.CountSearchAsuransi(search)
	return res, err
}

func (service *asuransiService) FindAsuransiById(id uint) (asuransiOutput model.Asuransi, err error) {
	return service.asuransiRepository.FindAsuransiById(id)
}

func (service *asuransiService) FindExcAsuransi(id uint) (asuransiOutput []model.Asuransi, err error) {
	return service.asuransiRepository.FindExcAsuransi(id)
}
func (service *asuransiService) InsertAsuransi(asuransi model.CreateAsuransiParameter) (model.Asuransi, error) {
	newAsuransi := model.Asuransi{}
	err1 := smapping.FillStruct(&newAsuransi, smapping.MapFields(&asuransi))

	if err1 != nil {
		return newAsuransi, err1
	}

	return service.asuransiRepository.InsertAsuransi(newAsuransi)
}

func (service *asuransiService) UpdateAsuransi(asuransi model.Asuransi, id uint) (asuransiOutput model.Asuransi, err error) {
	newAsuransi := model.Asuransi{}
	err1 := smapping.FillStruct(&newAsuransi, smapping.MapFields(&asuransi))

	if err1 != nil {
		return newAsuransi, err1
	}

	return service.asuransiRepository.UpdateAsuransi(newAsuransi, id)
}

func (service *asuransiService) DeleteAsuransi(asuransi model.Asuransi, id uint) (asuransiOutput model.Asuransi, err error) {
	newAsuransi := model.Asuransi{}
	err1 := smapping.FillStruct(&newAsuransi, smapping.MapFields(&asuransi))

	if err1 != nil {
		return newAsuransi, err1
	}

	return service.asuransiRepository.UpdateAsuransi(newAsuransi, id)
}
