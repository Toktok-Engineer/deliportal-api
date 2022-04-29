package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type EmailQueueReferenceService interface {
	FindEmailQueueReferences() (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error)
	FindEmailQueueReferenceById(id uint) (emailQueueReferenceOutput model.SelectEmailQueueReferenceParameter, err error)
	FindExcEmailQueueReference(id uint) (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error)
	InsertEmailQueueReference(emailQueueReference model.CreateEmailQueueReferenceParameter) (emailQueueReferenceOutput model.EmailQueueReference, err error)
	UpdateEmailQueueReference(emailQueueReference model.EmailQueueReference, id uint) (emailQueueReferenceOutput model.EmailQueueReference, err error)
	DeleteEmailQueueReference(emailQueueReference model.EmailQueueReference, id uint) (emailQueueReferenceOutput model.EmailQueueReference, err error)
}

type emailQueueReferenceService struct {
	emailQueueReferenceRepository repository.EmailQueueReferenceRepository
}

func NewEmailQueueReferenceService(emailQueueReferenceRep repository.EmailQueueReferenceRepository) EmailQueueReferenceService {
	return &emailQueueReferenceService{
		emailQueueReferenceRepository: emailQueueReferenceRep,
	}
}

func (service *emailQueueReferenceService) FindEmailQueueReferences() (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error) {
	res, err := service.emailQueueReferenceRepository.FindEmailQueueReferences()
	return res, err
}

func (service *emailQueueReferenceService) FindEmailQueueReferenceById(id uint) (emailQueueReferenceOutput model.SelectEmailQueueReferenceParameter, err error) {
	res, err := service.emailQueueReferenceRepository.FindEmailQueueReferenceById(id)
	return res, err
}

func (service *emailQueueReferenceService) FindExcEmailQueueReference(id uint) (emailQueueReferenceOutput []model.SelectEmailQueueReferenceParameter, err error) {
	res, err := service.emailQueueReferenceRepository.FindExcEmailQueueReference(id)
	return res, err
}

func (service *emailQueueReferenceService) InsertEmailQueueReference(emailQueueReference model.CreateEmailQueueReferenceParameter) (emailQueueReferenceOutput model.EmailQueueReference, err error) {
	newEmailQueueReference := model.EmailQueueReference{}
	err1 := smapping.FillStruct(&newEmailQueueReference, smapping.MapFields(&emailQueueReference))
	if err != nil {
		return newEmailQueueReference, err1
	}
	res, err := service.emailQueueReferenceRepository.InsertEmailQueueReference(newEmailQueueReference)
	return res, err
}

func (service *emailQueueReferenceService) UpdateEmailQueueReference(emailQueueReference model.EmailQueueReference, id uint) (emailQueueReferenceOutput model.EmailQueueReference, err error) {
	newEmailQueueReference := model.EmailQueueReference{}
	err1 := smapping.FillStruct(&newEmailQueueReference, smapping.MapFields(&emailQueueReference))
	if err != nil {
		return newEmailQueueReference, err1
	}
	res, err := service.emailQueueReferenceRepository.UpdateEmailQueueReference(newEmailQueueReference, id)
	return res, err
}

func (service *emailQueueReferenceService) DeleteEmailQueueReference(emailQueueReference model.EmailQueueReference, id uint) (emailQueueReferenceOutput model.EmailQueueReference, err error) {
	newEmailQueueReference := model.EmailQueueReference{}
	err1 := smapping.FillStruct(&newEmailQueueReference, smapping.MapFields(&emailQueueReference))
	if err != nil {
		return newEmailQueueReference, err1
	}
	res, err := service.emailQueueReferenceRepository.UpdateEmailQueueReference(newEmailQueueReference, id)
	return res, err
}
