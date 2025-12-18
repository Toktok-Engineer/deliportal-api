package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type KIRReminderService interface {
	CountKIRReminderAll(vehicleID int) (count int64, err error)
	FindKIRReminders() (kirreminderOutput []model.SelectKIRReminderParameter, err error)
	FindKIRRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (kirreminderOutput []model.SelectKIRReminderParameter, err error)
	SearchKIRReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (kirreminderOutput []model.SelectKIRReminderParameter, err error)
	CountSearchKIRReminder(vehicleID int, search string) (count int64, err error)
	FindKIRReminderById(id uint) (kirReminderOutput model.SelectKIRReminderParameter, err error)
	FindExcKIRReminder(id uint) (kirReminderOutput []model.SelectKIRReminderParameter, err error)
	InsertKIRReminder(kirReminder model.CreateKIRReminderParameter) (model.KIRReminder, error)
	UpdateKIRReminder(kirReminder model.KIRReminder, id uint) (kirReminderOutput model.KIRReminder, err error)
	DeleteKIRReminder(kirReminder model.KIRReminder, id uint) (kirReminderOutput model.KIRReminder, err error)
}

type kirReminderService struct {
	kirReminderRepository repository.KIRReminderRepository
}

func NewKIRReminderService(kirReminderRep repository.KIRReminderRepository) KIRReminderService {
	return &kirReminderService{
		kirReminderRepository: kirReminderRep,
	}
}

func (service *kirReminderService) CountKIRReminderAll(vehicleID int) (count int64, err error) {
	res, err := service.kirReminderRepository.CountKIRReminderAll(vehicleID)
	return res, err
}

func (service *kirReminderService) FindKIRReminders() (kirreminderOutput []model.SelectKIRReminderParameter, err error) {
	res, err := service.kirReminderRepository.FindKIRReminders()
	return res, err
}

func (service *kirReminderService) FindKIRRemindersOffset(vehicleID int, limit int, offset int, order string, dir string) (kirreminderOutput []model.SelectKIRReminderParameter, err error) {
	res, err := service.kirReminderRepository.FindKIRRemindersOffset(vehicleID, limit, offset, order, dir)
	return res, err
}

func (service *kirReminderService) SearchKIRReminder(vehicleID int, limit int, offset int, order string, dir string, search string) (kirreminderOutput []model.SelectKIRReminderParameter, err error) {
	res, err := service.kirReminderRepository.SearchKIRReminder(vehicleID, limit, offset, order, dir, search)
	return res, err
}

func (service *kirReminderService) CountSearchKIRReminder(vehicleID int, search string) (count int64, err error) {
	res, err := service.kirReminderRepository.CountSearchKIRReminder(vehicleID, search)
	return res, err
}

func (service *kirReminderService) FindKIRReminderById(id uint) (kirReminderOutput model.SelectKIRReminderParameter, err error) {
	return service.kirReminderRepository.FindKIRReminderById(id)
}

func (service *kirReminderService) FindExcKIRReminder(id uint) (kirReminderOutput []model.SelectKIRReminderParameter, err error) {
	return service.kirReminderRepository.FindExcKIRReminder(id)
}

func (service *kirReminderService) InsertKIRReminder(kirReminder model.CreateKIRReminderParameter) (model.KIRReminder, error) {
	newKIRReminder := model.KIRReminder{}
	err1 := smapping.FillStruct(&newKIRReminder, smapping.MapFields(&kirReminder))

	if err1 != nil {
		return newKIRReminder, err1
	}

	return service.kirReminderRepository.InsertKIRReminder(newKIRReminder)
}

func (service *kirReminderService) UpdateKIRReminder(kirReminder model.KIRReminder, id uint) (kirReminderOutput model.KIRReminder, err error) {
	newKIRReminder := model.KIRReminder{}
	err1 := smapping.FillStruct(&newKIRReminder, smapping.MapFields(&kirReminder))

	if err1 != nil {
		return newKIRReminder, err1
	}

	return service.kirReminderRepository.UpdateKIRReminder(newKIRReminder, id)
}

func (service *kirReminderService) DeleteKIRReminder(kirReminder model.KIRReminder, id uint) (kirReminderOutput model.KIRReminder, err error) {
	newKIRReminder := model.KIRReminder{}
	err1 := smapping.FillStruct(&newKIRReminder, smapping.MapFields(&kirReminder))

	if err1 != nil {
		return newKIRReminder, err1
	}

	return service.kirReminderRepository.UpdateKIRReminder(newKIRReminder, id)
}
