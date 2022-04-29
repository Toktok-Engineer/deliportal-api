package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type EmailQueueTypeService interface {
	FindEmailQueueTypes() (emailQueueTypeOutput []model.EmailQueueType, err error)
	FindEmailQueueTypeById(id uint) (emailQueueTypeOutput model.EmailQueueType, err error)
	FindExcEmailQueueType(id uint) (emailQueueTypeOutput []model.EmailQueueType, err error)
	InsertEmailQueueType(emailQueueType model.CreateEmailQueueTypeParameter) (emailQueueTypeOutput model.EmailQueueType, err error)
	UpdateEmailQueueType(emailQueueType model.EmailQueueType, id uint) (emailQueueTypeOutput model.EmailQueueType, err error)
	DeleteEmailQueueType(emailQueueType model.EmailQueueType, id uint) (emailQueueTypeOutput model.EmailQueueType, err error)
}

type emailQueueTypeService struct {
	emailQueueTypeRepository repository.EmailQueueTypeRepository
}

func NewEmailQueueTypeService(emailQueueTypeRep repository.EmailQueueTypeRepository) EmailQueueTypeService {
	return &emailQueueTypeService{
		emailQueueTypeRepository: emailQueueTypeRep,
	}
}

func (service *emailQueueTypeService) FindEmailQueueTypes() (emailQueueTypeOutput []model.EmailQueueType, err error) {
	res, err := service.emailQueueTypeRepository.FindEmailQueueTypes()
	return res, err
}

func (service *emailQueueTypeService) FindEmailQueueTypeById(id uint) (emailQueueTypeOutput model.EmailQueueType, err error) {
	res, err := service.emailQueueTypeRepository.FindEmailQueueTypeById(id)
	return res, err
}

func (service *emailQueueTypeService) FindExcEmailQueueType(id uint) (emailQueueTypeOutput []model.EmailQueueType, err error) {
	res, err := service.emailQueueTypeRepository.FindExcEmailQueueType(id)
	return res, err
}

func (service *emailQueueTypeService) InsertEmailQueueType(emailQueueType model.CreateEmailQueueTypeParameter) (emailQueueTypeOutput model.EmailQueueType, err error) {
	newEmailQueueType := model.EmailQueueType{}
	err1 := smapping.FillStruct(&newEmailQueueType, smapping.MapFields(&emailQueueType))
	if err != nil {
		return newEmailQueueType, err1
	}
	res, err := service.emailQueueTypeRepository.InsertEmailQueueType(newEmailQueueType)
	return res, err
}

func (service *emailQueueTypeService) UpdateEmailQueueType(emailQueueType model.EmailQueueType, id uint) (emailQueueTypeOutput model.EmailQueueType, err error) {
	newEmailQueueType := model.EmailQueueType{}
	err1 := smapping.FillStruct(&newEmailQueueType, smapping.MapFields(&emailQueueType))
	if err != nil {
		return newEmailQueueType, err1
	}
	res, err := service.emailQueueTypeRepository.UpdateEmailQueueType(newEmailQueueType, id)
	return res, err
}

func (service *emailQueueTypeService) DeleteEmailQueueType(emailQueueType model.EmailQueueType, id uint) (emailQueueTypeOutput model.EmailQueueType, err error) {
	newEmailQueueType := model.EmailQueueType{}
	err1 := smapping.FillStruct(&newEmailQueueType, smapping.MapFields(&emailQueueType))
	if err != nil {
		return newEmailQueueType, err1
	}
	res, err := service.emailQueueTypeRepository.UpdateEmailQueueType(newEmailQueueType, id)
	return res, err
}
