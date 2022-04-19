package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type FormService interface {
	FindForms() (formOutput []model.SelectFormParameter, err error)
	FindFormJoinRole(uId uint, fpId uint) (formOutput []model.SelectFormCRUDParameter, err error)
	FindFormByRole(uId uint) (formOutput []model.SelectFormCRUDParameter, err error)
	FindFormByType(tyId uint) (formOutput []model.SelectFormParameter, err error)
	FindExcFormByType(tyId uint, id uint) (formOutput []model.SelectFormParameter, err error)
	FindFormById(id uint) (formOutput model.SelectFormParameter, err error)
	FindFormByFormTypeId(ftId uint) (formOutput []model.SelectFormParameter, err error)
	FindExcForm(ftId uint, id uint) (formOutput []model.SelectFormParameter, err error)
	FindFormHead(ftId uint) (formOutput []model.SelectFormParameter, err error)
	FindFormHeadDetail(id uint) (formOutput model.SelectFormParameter, err error)
	FindExcFormHead(id uint) (formOutput []model.SelectFormParameter, err error)
	FindExcFormOnly(id uint) (formOutput []model.SelectFormParameter, err error)
	InsertForm(form model.CreateFormParameter) (formOutput model.Form, err error)
	UpdateForm(form model.Form, id uint) (formOutput model.Form, err error)
	DeleteForm(form model.Form, id uint) (formOutput model.Form, err error)
}

type formService struct {
	formRepository repository.FormRepository
}

func NewFormService(formRep repository.FormRepository) FormService {
	return &formService{
		formRepository: formRep,
	}
}

func (service *formService) FindForms() (formOutput []model.SelectFormParameter, err error) {
	res, err := service.formRepository.FindForms()
	return res, err
}

func (service *formService) FindFormJoinRole(uId uint, fpId uint) (formOutput []model.SelectFormCRUDParameter, err error) {
	res, err := service.formRepository.FindFormJoinRole(uId, fpId)
	return res, err
}

func (service *formService) FindFormByRole(uId uint) (formOutput []model.SelectFormCRUDParameter, err error) {
	res, err := service.formRepository.FindFormByRole(uId)
	return res, err
}

func (service *formService) FindFormByType(tyId uint) (formOutput []model.SelectFormParameter, err error) {
	res, err := service.formRepository.FindFormByType(tyId)
	return res, err
}

func (service *formService) FindExcFormByType(tyId uint, id uint) (formOutput []model.SelectFormParameter, err error) {
	res, err := service.formRepository.FindExcFormByType(tyId, id)
	return res, err
}

func (service *formService) FindFormById(id uint) (formOutput model.SelectFormParameter, err error) {
	res, err := service.formRepository.FindFormById(id)
	return res, err
}

func (service *formService) FindFormByFormTypeId(ftId uint) (formOutput []model.SelectFormParameter, err error) {
	res, err := service.formRepository.FindFormByFormTypeId(ftId)
	return res, err
}

func (service *formService) FindExcForm(ftId uint, id uint) (formOutput []model.SelectFormParameter, err error) {
	res, err := service.formRepository.FindExcForm(ftId, id)
	return res, err
}

func (service *formService) FindFormHeadDetail(id uint) (formOutput model.SelectFormParameter, err error) {
	res, err := service.formRepository.FindFormHeadDetail(id)
	return res, err
}

func (service *formService) FindFormHead(ftId uint) (formOutput []model.SelectFormParameter, err error) {
	res, err := service.formRepository.FindFormHead(ftId)
	return res, err
}

func (service *formService) FindExcFormHead(id uint) (formOutput []model.SelectFormParameter, err error) {
	res, err := service.formRepository.FindExcFormHead(id)
	return res, err
}

func (service *formService) FindExcFormOnly(id uint) (formOutput []model.SelectFormParameter, err error) {
	res, err := service.formRepository.FindExcFormOnly(id)
	return res, err
}

func (service *formService) InsertForm(form model.CreateFormParameter) (formOutput model.Form, err error) {
	newForm := model.Form{}
	err1 := smapping.FillStruct(&newForm, smapping.MapFields(&form))
	if err != nil {
		return newForm, err1
	}
	res, err := service.formRepository.InsertForm(newForm)
	return res, err
}

func (service *formService) UpdateForm(form model.Form, id uint) (formOutput model.Form, err error) {
	newForm := model.Form{}
	err1 := smapping.FillStruct(&newForm, smapping.MapFields(&form))
	if err != nil {
		return newForm, err1
	}
	res, err := service.formRepository.UpdateForm(newForm, id)
	return res, err
}

func (service *formService) DeleteForm(form model.Form, id uint) (formOutput model.Form, err error) {
	newForm := model.Form{}
	err1 := smapping.FillStruct(&newForm, smapping.MapFields(&form))
	if err != nil {
		return newForm, err1
	}
	res, err := service.formRepository.UpdateForm(newForm, id)
	return res, err
}
