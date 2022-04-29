package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type EmailQueueService interface {
	FindEmailQueues() (emailQueueOutput []model.SelectEmailQueueParameter, err error)
	FindEmailQueueById(id uint) (emailQueueOutput model.SelectEmailQueueParameter, err error)
	FindExcEmailQueue(id uint) (emailQueueOutput []model.SelectEmailQueueParameter, err error)
	FindEmailQueueByStatus(status uint) (emailQueueOutput []model.SelectEmailQueueParameter, err error)
	InsertEmailQueue(emailQueue model.CreateEmailQueueParameter) (emailQueueOutput model.EmailQueue, err error)
	UpdateEmailQueue(emailQueue model.EmailQueue, id uint) (emailQueueOutput model.EmailQueue, err error)
	DeleteEmailQueue(emailQueue model.EmailQueue, id uint) (emailQueueOutput model.EmailQueue, err error)
}

type emailQueueService struct {
	emailQueueRepository repository.EmailQueueRepository
}

func NewEmailQueueService(emailQueueRep repository.EmailQueueRepository) EmailQueueService {
	return &emailQueueService{
		emailQueueRepository: emailQueueRep,
	}
}

func (service *emailQueueService) FindEmailQueues() (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	res, err := service.emailQueueRepository.FindEmailQueues()
	return res, err
}

func (service *emailQueueService) FindEmailQueueById(id uint) (emailQueueOutput model.SelectEmailQueueParameter, err error) {
	res, err := service.emailQueueRepository.FindEmailQueueById(id)
	return res, err
}

func (service *emailQueueService) FindExcEmailQueue(id uint) (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	res, err := service.emailQueueRepository.FindExcEmailQueue(id)
	return res, err
}

func (service *emailQueueService) FindEmailQueueByStatus(status uint) (emailQueueOutput []model.SelectEmailQueueParameter, err error) {
	res, err := service.emailQueueRepository.FindEmailQueueByStatus(status)
	return res, err
}

func (service *emailQueueService) InsertEmailQueue(emailQueue model.CreateEmailQueueParameter) (emailQueueOutput model.EmailQueue, err error) {
	newEmailQueue := model.EmailQueue{}
	err1 := smapping.FillStruct(&newEmailQueue, smapping.MapFields(&emailQueue))
	if err != nil {
		return newEmailQueue, err1
	}
	res, err := service.emailQueueRepository.InsertEmailQueue(newEmailQueue)
	return res, err
}

func (service *emailQueueService) UpdateEmailQueue(emailQueue model.EmailQueue, id uint) (emailQueueOutput model.EmailQueue, err error) {
	newEmailQueue := model.EmailQueue{}
	err1 := smapping.FillStruct(&newEmailQueue, smapping.MapFields(&emailQueue))
	if err != nil {
		return newEmailQueue, err1
	}
	res, err := service.emailQueueRepository.UpdateEmailQueue(newEmailQueue, id)
	return res, err
}

func (service *emailQueueService) DeleteEmailQueue(emailQueue model.EmailQueue, id uint) (emailQueueOutput model.EmailQueue, err error) {
	newEmailQueue := model.EmailQueue{}
	err1 := smapping.FillStruct(&newEmailQueue, smapping.MapFields(&emailQueue))
	if err != nil {
		return newEmailQueue, err1
	}
	res, err := service.emailQueueRepository.UpdateEmailQueue(newEmailQueue, id)
	return res, err
}
