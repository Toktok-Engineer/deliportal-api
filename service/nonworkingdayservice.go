package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type NonWorkingDayService interface {
	CountNonWorkingDayAll() (count int64, err error)
	FindNonWorkingDays() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	FindNonWorkingDaysCuti() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	FindNonWorkingDaysAllDate() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	FindNonWorkingDaysOffset(limit int, offset int, order string, dir string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	SearchNonWorkingDay(limit int, offset int, order string, dir string, search string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	CountSearchNonWorkingDay(search string) (count int64, err error)
	CountNonWorkingDayExPersonalia() (count int64, err error)
	FindNonWorkingDaysExPersonaliaOffset(limit int, offset int, order string, dir string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	SearchNonWorkingDayExPersonalia(limit int, offset int, order string, dir string, search string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	CountSearchNonWorkingDayExPersonalia(search string) (count int64, err error)
	FindNonWorkingDayById(id uint) (nonWorkingDayOutput model.SelectNonWorkingDayParameter, err error)
	FindExcNonWorkingDay(nwtId uint, id uint) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	FindNonWorkingDayByNWTId(nwtId uint) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	CountNonWorkingDayName(search string) (count int64, err error)
	FindNonWorkingDaybyDate(date float64) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	InsertNonWorkingDay(nonWorkingDay model.CreateNonWorkingDayParameter) (nonWorkingDayOutput model.NonWorkingDay, err error)
	UpdateNonWorkingDay(nonWorkingDay model.NonWorkingDay, id uint) (nonWorkingDayOutput model.NonWorkingDay, err error)
	DeleteNonWorkingDay(nonWorkingDay model.NonWorkingDay, id uint) (nonWorkingDayOutput model.NonWorkingDay, err error)
}

type nonWorkingDayService struct {
	nonWorkingDayRepository repository.NonWorkingDayRepository
}

func NewNonWorkingDayService(nonWorkingDayRep repository.NonWorkingDayRepository) NonWorkingDayService {
	return &nonWorkingDayService{
		nonWorkingDayRepository: nonWorkingDayRep,
	}
}

func (service *nonWorkingDayService) CountNonWorkingDayAll() (count int64, err error) {
	res, err := service.nonWorkingDayRepository.CountNonWorkingDayAll()
	return res, err
}

func (service *nonWorkingDayService) FindNonWorkingDays() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.FindNonWorkingDays()
	return res, err
}

func (service *nonWorkingDayService) FindNonWorkingDaysCuti() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.FindNonWorkingDaysCuti()
	return res, err
}

func (service *nonWorkingDayService) FindNonWorkingDaysAllDate() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.FindNonWorkingDaysAllDate()
	return res, err
}

func (service *nonWorkingDayService) FindNonWorkingDaysOffset(limit int, offset int, order string, dir string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.FindNonWorkingDaysOffset(limit, offset, order, dir)
	return res, err
}

func (service *nonWorkingDayService) SearchNonWorkingDay(limit int, offset int, order string, dir string, search string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.SearchNonWorkingDay(limit, offset, order, dir, search)
	return res, err
}

func (service *nonWorkingDayService) CountSearchNonWorkingDay(search string) (count int64, err error) {
	res, err := service.nonWorkingDayRepository.CountSearchNonWorkingDay(search)
	return res, err
}

func (service *nonWorkingDayService) CountNonWorkingDayExPersonalia() (count int64, err error) {
	res, err := service.nonWorkingDayRepository.CountNonWorkingDayExPersonalia()
	return res, err
}
func (service *nonWorkingDayService) FindNonWorkingDaysExPersonaliaOffset(limit int, offset int, order string, dir string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.FindNonWorkingDaysExPersonaliaOffset(limit, offset, order, dir)
	return res, err
}

func (service *nonWorkingDayService) SearchNonWorkingDayExPersonalia(limit int, offset int, order string, dir string, search string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.SearchNonWorkingDayExPersonalia(limit, offset, order, dir, search)
	return res, err
}

func (service *nonWorkingDayService) CountSearchNonWorkingDayExPersonalia(search string) (count int64, err error) {
	res, err := service.nonWorkingDayRepository.CountSearchNonWorkingDayExPersonalia(search)
	return res, err
}

func (service *nonWorkingDayService) FindNonWorkingDayById(id uint) (nonWorkingDayOutput model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.FindNonWorkingDayById(id)
	return res, err
}

func (service *nonWorkingDayService) FindExcNonWorkingDay(nwtId uint, id uint) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.FindExcNonWorkingDay(nwtId, id)
	return res, err
}

func (service *nonWorkingDayService) FindNonWorkingDayByNWTId(nwtId uint) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.FindNonWorkingDayByNWTId(nwtId)
	return res, err
}

func (service *nonWorkingDayService) CountNonWorkingDayName(search string) (count int64, err error) {
	res, err := service.nonWorkingDayRepository.CountNonWorkingDayName(search)
	return res, err
}

func (service *nonWorkingDayService) FindNonWorkingDaybyDate(date float64) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	res, err := service.nonWorkingDayRepository.FindNonWorkingDaybyDate(date)
	return res, err
}

func (service *nonWorkingDayService) InsertNonWorkingDay(nonWorkingDay model.CreateNonWorkingDayParameter) (nonWorkingDayOutput model.NonWorkingDay, err error) {
	newNonWorkingDay := model.NonWorkingDay{}
	err1 := smapping.FillStruct(&newNonWorkingDay, smapping.MapFields(&nonWorkingDay))
	if err != nil {
		return newNonWorkingDay, err1
	}
	res, err := service.nonWorkingDayRepository.InsertNonWorkingDay(newNonWorkingDay)
	return res, err
}

func (service *nonWorkingDayService) UpdateNonWorkingDay(nonWorkingDay model.NonWorkingDay, id uint) (nonWorkingDayOutput model.NonWorkingDay, err error) {
	newNonWorkingDay := model.NonWorkingDay{}
	err1 := smapping.FillStruct(&newNonWorkingDay, smapping.MapFields(&nonWorkingDay))
	if err != nil {
		return newNonWorkingDay, err1
	}
	res, err := service.nonWorkingDayRepository.UpdateNonWorkingDay(newNonWorkingDay, id)
	return res, err
}

func (service *nonWorkingDayService) DeleteNonWorkingDay(nonWorkingDay model.NonWorkingDay, id uint) (nonWorkingDayOutput model.NonWorkingDay, err error) {
	newNonWorkingDay := model.NonWorkingDay{}
	err1 := smapping.FillStruct(&newNonWorkingDay, smapping.MapFields(&nonWorkingDay))
	if err != nil {
		return newNonWorkingDay, err1
	}
	res, err := service.nonWorkingDayRepository.UpdateNonWorkingDay(newNonWorkingDay, id)
	return res, err
}
