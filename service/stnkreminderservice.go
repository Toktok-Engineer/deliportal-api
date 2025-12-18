package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type STNKReminderService interface {
	CountSTNKReminderAll(vehicleID int) (count int64, err error)
	FindSTNKReminders() (stnkreminderOutput []model.SelectSTNKReminderParameter, err error)
	FindSTNKRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (stnkreminderOutput []model.SelectSTNKReminderParameter, err error)
	SearchSTNKReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (stnkreminderOutput []model.SelectSTNKReminderParameter, err error)
	CountSearchSTNKReminder(vehicleID int, search string) (count int64, err error)
	FindSTNKReminderById(id uint) (stnkReminderOutput model.SelectSTNKReminderParameter, err error)
	FindExcSTNKReminder(id uint) (stnkReminderOutput []model.SelectSTNKReminderParameter, err error)
	InsertSTNKReminder(stnkReminder model.CreateSTNKReminderParameter) (model.STNKReminder, error)
	UpdateSTNKReminder(stnkReminder model.STNKReminder, id uint) (stnkReminderOutput model.STNKReminder, err error)
	DeleteSTNKReminder(stnkReminder model.STNKReminder, id uint) (stnkReminderOutput model.STNKReminder, err error)
}

type stnkReminderService struct {
	stnkReminderRepository repository.STNKReminderRepository
}

func NewSTNKReminderService(stnkReminderRep repository.STNKReminderRepository) STNKReminderService {
	return &stnkReminderService{
		stnkReminderRepository: stnkReminderRep,
	}
}

func (service *stnkReminderService) CountSTNKReminderAll(vehicleID int) (count int64, err error) {
	res, err := service.stnkReminderRepository.CountSTNKReminderAll(vehicleID)
	return res, err
}

func (service *stnkReminderService) FindSTNKReminders() (stnkreminderOutput []model.SelectSTNKReminderParameter, err error) {
	res, err := service.stnkReminderRepository.FindSTNKReminders()
	return res, err
}

func (service *stnkReminderService) FindSTNKRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (stnkreminderOutput []model.SelectSTNKReminderParameter, err error) {
	res, err := service.stnkReminderRepository.FindSTNKRemindersOffset(vehicleID, limit, offset, order, dir)
	return res, err
}

func (service *stnkReminderService) SearchSTNKReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (stnkreminderOutput []model.SelectSTNKReminderParameter, err error) {
	res, err := service.stnkReminderRepository.SearchSTNKReminder(vehicleID, limit, offset, order, dir, search)
	return res, err
}

func (service *stnkReminderService) CountSearchSTNKReminder(vehicleID int, search string) (count int64, err error) {
	res, err := service.stnkReminderRepository.CountSearchSTNKReminder(vehicleID, search)
	return res, err
}

func (service *stnkReminderService) FindSTNKReminderById(id uint) (stnkReminderOutput model.SelectSTNKReminderParameter, err error) {
	return service.stnkReminderRepository.FindSTNKReminderById(id)
}

func (service *stnkReminderService) FindExcSTNKReminder(id uint) (stnkReminderOutput []model.SelectSTNKReminderParameter, err error) {
	return service.stnkReminderRepository.FindExcSTNKReminder(id)
}

func (service *stnkReminderService) InsertSTNKReminder(stnkReminder model.CreateSTNKReminderParameter) (model.STNKReminder, error) {
	newSTNKReminder := model.STNKReminder{}
	err1 := smapping.FillStruct(&newSTNKReminder, smapping.MapFields(&stnkReminder))

	if err1 != nil {
		return newSTNKReminder, err1
	}

	return service.stnkReminderRepository.InsertSTNKReminder(newSTNKReminder)
}

func (service *stnkReminderService) UpdateSTNKReminder(stnkReminder model.STNKReminder, id uint) (stnkReminderOutput model.STNKReminder, err error) {
	newSTNKReminder := model.STNKReminder{}
	err1 := smapping.FillStruct(&newSTNKReminder, smapping.MapFields(&stnkReminder))

	if err1 != nil {
		return newSTNKReminder, err1
	}

	return service.stnkReminderRepository.UpdateSTNKReminder(newSTNKReminder, id)
}

func (service *stnkReminderService) DeleteSTNKReminder(stnkReminder model.STNKReminder, id uint) (stnkReminderOutput model.STNKReminder, err error) {
	newSTNKReminder := model.STNKReminder{}
	err1 := smapping.FillStruct(&newSTNKReminder, smapping.MapFields(&stnkReminder))

	if err1 != nil {
		return newSTNKReminder, err1
	}

	return service.stnkReminderRepository.UpdateSTNKReminder(newSTNKReminder, id)
}
