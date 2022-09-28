package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type InternalMemoService interface {
	CountInternalMemoAll(employeeID int) (internalMemoOutput []model.SelectInternalMemoParameter, err error)
	FindInternalMemos() (internalMemoOutput []model.SelectInternalMemoParameter, err error)
	FindInternalMemosOffset(limit int, offset int, order string, dir string, employeeID int) (internalMemoOutput []model.SelectInternalMemoParameter, err error)
	SearchInternalMemo(limit int, offset int, order string, dir string, search string, employeeID int) (internalMemoOutput []model.SelectInternalMemoParameter, err error)
	CountSearchInternalMemo(search string, employeeID int) (internalMemoOutput []model.SelectInternalMemoParameter, err error)
	CountInternalMemoByDept(deptName string) (count int64, err error)
	FindInternalMemoById(id uint) (internalMemoOutput model.SelectInternalMemoParameter, err error)
	FindExcInternalMemo(id uint) (internalMemoOutput []model.SelectInternalMemoParameter, err error)
	InsertInternalMemo(internalMemo model.CreateInternalMemoParameter) (internalMemoOutput model.InternalMemo, err error)
	UpdateInternalMemo(internalMemo model.InternalMemo, id uint) (internalMemoOutput model.InternalMemo, err error)
	DeleteInternalMemo(internalMemo model.InternalMemo, id uint) (internalMemoOutput model.InternalMemo, err error)
	UpdateInternalMemoApprove(internalMemo model.InternalMemo, id uint) (internalMemoOutput model.InternalMemo, err error)
	UploadIMDocument(internalMemo model.InternalMemo, id uint) (internalMemoOutput model.InternalMemo, err error)
}

type internalMemoService struct {
	internalMemoRepository repository.InternalMemoRepository
}

func NewInternalMemoService(internalMemoRep repository.InternalMemoRepository) InternalMemoService {
	return &internalMemoService{
		internalMemoRepository: internalMemoRep,
	}
}

func (service *internalMemoService) CountInternalMemoAll(employeeID int) (internalMemoOutput []model.SelectInternalMemoParameter, err error) {
	res, err := service.internalMemoRepository.CountInternalMemoAll(employeeID)
	return res, err
}

func (service *internalMemoService) FindInternalMemos() (internalMemoOutput []model.SelectInternalMemoParameter, err error) {
	res, err := service.internalMemoRepository.FindInternalMemos()
	return res, err
}

func (service *internalMemoService) FindInternalMemosOffset(limit int, offset int, order string, dir string, employeeID int) (internalMemoOutput []model.SelectInternalMemoParameter, err error) {
	res, err := service.internalMemoRepository.FindInternalMemosOffset(limit, offset, order, dir, employeeID)
	return res, err
}

func (service *internalMemoService) SearchInternalMemo(limit int, offset int, order string, dir string, search string, employeeID int) (internalMemoOutput []model.SelectInternalMemoParameter, err error) {
	res, err := service.internalMemoRepository.SearchInternalMemo(limit, offset, order, dir, search, employeeID)
	return res, err
}

func (service *internalMemoService) CountSearchInternalMemo(search string, employeeID int) (internalMemoOutput []model.SelectInternalMemoParameter, err error) {
	res, err := service.internalMemoRepository.CountSearchInternalMemo(search, employeeID)
	return res, err
}

func (service *internalMemoService) CountInternalMemoByDept(deptName string) (count int64, err error) {
	res, err := service.internalMemoRepository.CountInternalMemoByDept(deptName)
	return res, err
}

func (service *internalMemoService) FindInternalMemoById(id uint) (internalMemoOutput model.SelectInternalMemoParameter, err error) {
	res, err := service.internalMemoRepository.FindInternalMemoById(id)
	return res, err
}

func (service *internalMemoService) FindExcInternalMemo(id uint) (internalMemoOutput []model.SelectInternalMemoParameter, err error) {
	res, err := service.internalMemoRepository.FindExcInternalMemo(id)
	return res, err
}

func (service *internalMemoService) InsertInternalMemo(internalMemo model.CreateInternalMemoParameter) (internalMemoOutput model.InternalMemo, err error) {
	newInternalMemo := model.InternalMemo{}
	err1 := smapping.FillStruct(&newInternalMemo, smapping.MapFields(&internalMemo))
	if err != nil {
		return newInternalMemo, err1
	}
	res, err := service.internalMemoRepository.InsertInternalMemo(newInternalMemo)
	return res, err
}

func (service *internalMemoService) UpdateInternalMemo(internalMemo model.InternalMemo, id uint) (internalMemoOutput model.InternalMemo, err error) {
	newInternalMemo := model.InternalMemo{}
	err1 := smapping.FillStruct(&newInternalMemo, smapping.MapFields(&internalMemo))
	if err != nil {
		return newInternalMemo, err1
	}
	res, err := service.internalMemoRepository.UpdateInternalMemo(newInternalMemo, id)
	return res, err
}

func (service *internalMemoService) DeleteInternalMemo(internalMemo model.InternalMemo, id uint) (internalMemoOutput model.InternalMemo, err error) {
	newInternalMemo := model.InternalMemo{}
	err1 := smapping.FillStruct(&newInternalMemo, smapping.MapFields(&internalMemo))
	if err != nil {
		return newInternalMemo, err1
	}
	res, err := service.internalMemoRepository.UpdateInternalMemo(newInternalMemo, id)
	return res, err
}

func (service *internalMemoService) UpdateInternalMemoApprove(internalMemo model.InternalMemo, id uint) (internalMemoOutput model.InternalMemo, err error) {
	newInternalMemo := model.InternalMemo{}
	err1 := smapping.FillStruct(&newInternalMemo, smapping.MapFields(&internalMemo))
	if err != nil {
		return newInternalMemo, err1
	}
	res, err := service.internalMemoRepository.UpdateInternalMemoApprove(newInternalMemo, id)
	return res, err
}

func (service *internalMemoService) UploadIMDocument(internalMemo model.InternalMemo, id uint) (internalMemoOutput model.InternalMemo, err error) {
	newInternalMemo := model.InternalMemo{}
	err1 := smapping.FillStruct(&newInternalMemo, smapping.MapFields(&internalMemo))
	if err != nil {
		return newInternalMemo, err1
	}
	res, err := service.internalMemoRepository.UploadIMDocument(newInternalMemo, id)
	return res, err
}
