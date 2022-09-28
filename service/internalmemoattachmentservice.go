package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type InternalMemoAttachmentService interface {
	CountInternalMemoAttachmentAll(internalMemoId int) (count int64, err error)
	FindInternalMemoAttachments(internalMemoId int) (internalMemoAttachmentOutput []model.InternalMemoAttachment, err error)
	FindInternalMemoAttachmentsOffset(limit int, offset int, order string, dir string, internalMemoId int) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error)
	SearchInternalMemoAttachment(limit int, offset int, order string, dir string, search string, internalMemoId int) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error)
	CountSearchInternalMemoAttachment(search string, internalMemoId int) (count int64, err error)
	FindInternalMemoAttachmentById(id uint) (internalMemoAttachmentOutput model.SelectInternalMemoAttachmentParameter, err error)
	FindExcInternalMemoAttachment(id uint) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error)
	FindInternalMemoAttachmentByCompanyId(id uint) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error)
	InsertInternalMemoAttachment(internalMemoAttachment model.CreateInternalMemoAttachmentParameter) (internalMemoAttachmentOutput model.InternalMemoAttachment, err error)
	UpdateInternalMemoAttachment(internalMemoAttachment model.InternalMemoAttachment, id uint) (internalMemoAttachmentOutput model.InternalMemoAttachment, err error)
	DeleteInternalMemoAttachment(internalMemoAttachment model.InternalMemoAttachment, id uint) (internalMemoAttachmentOutput model.InternalMemoAttachment, err error)
}

type internalMemoAttachmentService struct {
	internalMemoAttachmentRepository repository.InternalMemoAttachmentRepository
}

func NewInternalMemoAttachmentService(internalMemoAttachmentRep repository.InternalMemoAttachmentRepository) InternalMemoAttachmentService {
	return &internalMemoAttachmentService{
		internalMemoAttachmentRepository: internalMemoAttachmentRep,
	}
}

func (service *internalMemoAttachmentService) CountInternalMemoAttachmentAll(internalMemoId int) (count int64, err error) {
	res, err := service.internalMemoAttachmentRepository.CountInternalMemoAttachmentAll(internalMemoId)
	return res, err
}

func (service *internalMemoAttachmentService) FindInternalMemoAttachments(internalMemoId int) (internalMemoAttachmentOutput []model.InternalMemoAttachment, err error) {
	res, err := service.internalMemoAttachmentRepository.FindInternalMemoAttachments(internalMemoId)
	return res, err
}

func (service *internalMemoAttachmentService) FindInternalMemoAttachmentsOffset(limit int, offset int, order string, dir string, internalMemoId int) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error) {
	res, err := service.internalMemoAttachmentRepository.FindInternalMemoAttachmentsOffset(limit, offset, order, dir, internalMemoId)
	return res, err
}

func (service *internalMemoAttachmentService) SearchInternalMemoAttachment(limit int, offset int, order string, dir string, search string, internalMemoId int) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error) {
	res, err := service.internalMemoAttachmentRepository.SearchInternalMemoAttachment(limit, offset, order, dir, search, internalMemoId)
	return res, err
}

func (service *internalMemoAttachmentService) CountSearchInternalMemoAttachment(search string, internalMemoId int) (count int64, err error) {
	res, err := service.internalMemoAttachmentRepository.CountSearchInternalMemoAttachment(search, internalMemoId)
	return res, err
}

func (service *internalMemoAttachmentService) FindInternalMemoAttachmentById(id uint) (internalMemoAttachmentOutput model.SelectInternalMemoAttachmentParameter, err error) {
	res, err := service.internalMemoAttachmentRepository.FindInternalMemoAttachmentById(id)
	return res, err
}

func (service *internalMemoAttachmentService) FindExcInternalMemoAttachment(id uint) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error) {
	res, err := service.internalMemoAttachmentRepository.FindExcInternalMemoAttachment(id)
	return res, err
}

func (service *internalMemoAttachmentService) FindInternalMemoAttachmentByCompanyId(id uint) (internalMemoAttachmentOutput []model.SelectInternalMemoAttachmentParameter, err error) {
	res, err := service.internalMemoAttachmentRepository.FindInternalMemoAttachmentByCompanyId(id)
	return res, err
}

func (service *internalMemoAttachmentService) InsertInternalMemoAttachment(internalMemoAttachment model.CreateInternalMemoAttachmentParameter) (internalMemoAttachmentOutput model.InternalMemoAttachment, err error) {
	newInternalMemoAttachment := model.InternalMemoAttachment{}
	err1 := smapping.FillStruct(&newInternalMemoAttachment, smapping.MapFields(&internalMemoAttachment))
	if err != nil {
		return newInternalMemoAttachment, err1
	}
	res, err := service.internalMemoAttachmentRepository.InsertInternalMemoAttachment(newInternalMemoAttachment)
	return res, err
}

func (service *internalMemoAttachmentService) UpdateInternalMemoAttachment(internalMemoAttachment model.InternalMemoAttachment, id uint) (internalMemoAttachmentOutput model.InternalMemoAttachment, err error) {
	newInternalMemoAttachment := model.InternalMemoAttachment{}
	err1 := smapping.FillStruct(&newInternalMemoAttachment, smapping.MapFields(&internalMemoAttachment))
	if err != nil {
		return newInternalMemoAttachment, err1
	}
	res, err := service.internalMemoAttachmentRepository.UpdateInternalMemoAttachment(newInternalMemoAttachment, id)
	return res, err
}

func (service *internalMemoAttachmentService) DeleteInternalMemoAttachment(internalMemoAttachment model.InternalMemoAttachment, id uint) (internalMemoAttachmentOutput model.InternalMemoAttachment, err error) {
	newInternalMemoAttachment := model.InternalMemoAttachment{}
	err1 := smapping.FillStruct(&newInternalMemoAttachment, smapping.MapFields(&internalMemoAttachment))
	if err != nil {
		return newInternalMemoAttachment, err1
	}
	res, err := service.internalMemoAttachmentRepository.UpdateInternalMemoAttachment(newInternalMemoAttachment, id)
	return res, err
}
