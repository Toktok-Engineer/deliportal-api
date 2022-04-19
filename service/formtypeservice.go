package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type FormTypeService interface {
	FindFormTypes() (formTypeOutput []model.FormType, err error)
	FindFormTypeById(id uint) (formTypeOutput model.FormType, err error)
	FindExcFormType(id uint) (formTypeOutput []model.FormType, err error)
	InsertFormType(formType model.CreateFormTypeParameter) (formTypeOutput model.FormType, err error)
	UpdateFormType(formType model.FormType, id uint) (formTypeOutput model.FormType, err error)
	DeleteFormType(formType model.FormType, id uint) (formTypeOutput model.FormType, err error)
}

type formTypeService struct {
	formTypeRepository repository.FormTypeRepository
}

func NewFormTypeService(formTypeRep repository.FormTypeRepository) FormTypeService {
	return &formTypeService{
		formTypeRepository: formTypeRep,
	}
}

func (service *formTypeService) FindFormTypes() (formTypeOutput []model.FormType, err error) {
	res, err := service.formTypeRepository.FindFormTypes()
	return res, err
}

func (service *formTypeService) FindFormTypeById(id uint) (formTypeOutput model.FormType, err error) {
	res, err := service.formTypeRepository.FindFormTypeById(id)
	return res, err
}

func (service *formTypeService) FindExcFormType(id uint) (formTypeOutput []model.FormType, err error) {
	res, err := service.formTypeRepository.FindExcFormType(id)
	return res, err
}

func (service *formTypeService) InsertFormType(formType model.CreateFormTypeParameter) (formTypeOutput model.FormType, err error) {
	newFormType := model.FormType{}
	err1 := smapping.FillStruct(&newFormType, smapping.MapFields(&formType))
	if err != nil {
		return newFormType, err1
	}
	res, err := service.formTypeRepository.InsertFormType(newFormType)
	return res, err
}

func (service *formTypeService) UpdateFormType(formType model.FormType, id uint) (formTypeOutput model.FormType, err error) {
	newFormType := model.FormType{}
	err1 := smapping.FillStruct(&newFormType, smapping.MapFields(&formType))
	if err != nil {
		return newFormType, err1
	}
	res, err := service.formTypeRepository.UpdateFormType(newFormType, id)
	return res, err
}

func (service *formTypeService) DeleteFormType(formType model.FormType, id uint) (formTypeOutput model.FormType, err error) {
	newFormType := model.FormType{}
	err1 := smapping.FillStruct(&newFormType, smapping.MapFields(&formType))
	if err != nil {
		return newFormType, err1
	}
	res, err := service.formTypeRepository.UpdateFormType(newFormType, id)
	return res, err
}
