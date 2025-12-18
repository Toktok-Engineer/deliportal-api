package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type AsuransiReminderService interface {
	CountAsuransiReminderAll(vehicleID int) (count int64, err error)
	FindAsuransiReminders() (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error)
	FindAsuransiRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error)
	SearchAsuransiReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error)
	CountSearchAsuransiReminder(vehicleID int, search string) (count int64, err error)
	FindAsuransiReminderById(id uint) (asuransiReminderOutput model.SelectAsuransiReminderParameter, err error)
	FindExcAsuransiReminder(id uint) (asuransiReminderOutput []model.SelectAsuransiReminderParameter, err error)
	InsertAsuransiReminder(asuransiReminder model.CreateAsuransiReminderParameter) (model.AsuransiReminder, error)
	UpdateAsuransiReminder(asuransiReminder model.AsuransiReminder, id uint) (asuransiReminderOutput model.AsuransiReminder, err error)
	DeleteAsuransiReminder(asuransiReminder model.AsuransiReminder, id uint) (asuransiReminderOutput model.AsuransiReminder, err error)
}

type asuransiReminderService struct {
	asuransiReminderRepository repository.AsuransiReminderRepository
}

func NewAsuransiReminderService(asuransiReminderRep repository.AsuransiReminderRepository) AsuransiReminderService {
	return &asuransiReminderService{
		asuransiReminderRepository: asuransiReminderRep,
	}
}

func (service *asuransiReminderService) CountAsuransiReminderAll(vehicleID int) (count int64, err error) {
	res, err := service.asuransiReminderRepository.CountAsuransiReminderAll(vehicleID)
	return res, err
}

func (service *asuransiReminderService) FindAsuransiReminders() (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error) {
	res, err := service.asuransiReminderRepository.FindAsuransiReminders()
	return res, err
}

func (service *asuransiReminderService) FindAsuransiRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error) {
	res, err := service.asuransiReminderRepository.FindAsuransiRemindersOffset(vehicleID, limit, offset, order, dir)
	return res, err
}

func (service *asuransiReminderService) SearchAsuransiReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (asuransireminderOutput []model.SelectAsuransiReminderParameter, err error) {
	res, err := service.asuransiReminderRepository.SearchAsuransiReminder(vehicleID, limit, offset, order, dir, search)
	return res, err
}

func (service *asuransiReminderService) CountSearchAsuransiReminder(vehicleID int, search string) (count int64, err error) {
	res, err := service.asuransiReminderRepository.CountSearchAsuransiReminder(vehicleID, search)
	return res, err
}

func (service *asuransiReminderService) FindAsuransiReminderById(id uint) (asuransiReminderOutput model.SelectAsuransiReminderParameter, err error) {
	return service.asuransiReminderRepository.FindAsuransiReminderById(id)
}

func (service *asuransiReminderService) FindExcAsuransiReminder(id uint) (asuransiReminderOutput []model.SelectAsuransiReminderParameter, err error) {
	return service.asuransiReminderRepository.FindExcAsuransiReminder(id)
}

func (service *asuransiReminderService) InsertAsuransiReminder(asuransiReminder model.CreateAsuransiReminderParameter) (model.AsuransiReminder, error) {
	newAsuransiReminder := model.AsuransiReminder{}
	err1 := smapping.FillStruct(&newAsuransiReminder, smapping.MapFields(&asuransiReminder))

	if err1 != nil {
		return newAsuransiReminder, err1
	}

	return service.asuransiReminderRepository.InsertAsuransiReminder(newAsuransiReminder)
}

func (service *asuransiReminderService) UpdateAsuransiReminder(asuransiReminder model.AsuransiReminder, id uint) (asuransiReminderOutput model.AsuransiReminder, err error) {
	newAsuransiReminder := model.AsuransiReminder{}
	err1 := smapping.FillStruct(&newAsuransiReminder, smapping.MapFields(&asuransiReminder))

	if err1 != nil {
		return newAsuransiReminder, err1
	}

	return service.asuransiReminderRepository.UpdateAsuransiReminder(newAsuransiReminder, id)
}

func (service *asuransiReminderService) DeleteAsuransiReminder(asuransiReminder model.AsuransiReminder, id uint) (asuransiReminderOutput model.AsuransiReminder, err error) {
	newAsuransiReminder := model.AsuransiReminder{}
	err1 := smapping.FillStruct(&newAsuransiReminder, smapping.MapFields(&asuransiReminder))

	if err1 != nil {
		return newAsuransiReminder, err1
	}

	return service.asuransiReminderRepository.UpdateAsuransiReminder(newAsuransiReminder, id)
}
