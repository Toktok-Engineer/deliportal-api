package service

import (
	"deliportal-api/model"
	"deliportal-api/repository"

	"github.com/mashingan/smapping"
)

type EmployeeLeaveRequestService interface {
	CountEmployeeLeaveRequestAll(empId int) (count int64, err error)
	CountEmployeeLeaveRequestApprovalAllDt(year string) (count int64, err error)
	FindEmployeeLeaveRequests() (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	FindEmployeeLeaveRequestsOffset(limit int, offset int, order string, dir string, empId int) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	SearchEmployeeLeaveRequest(limit int, offset int, order string, dir string, search string, empId int) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	CountSearchEmployeeLeaveRequest(search string, empId int) (count int64, err error)
	FindEmployeeLeaveRequestById(id uint) (employeeLeaveRequestOutput model.SelectEmployeeLeaveRequestParameter, err error)
	CountEmployeeLeaveRequestAllPIC(empId int) (count int64, err error)
	FindEmployeeLeaveRequestsPICOffset(limit int, offset int, order string, dir string, empId int) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestPICParameter, err error)
	CountSearchEmployeeLeaveRequestPIC(search string, empId int) (count int64, err error)
	SearchEmployeeLeaveRequestPIC(limit int, offset int, order string, dir string, search string, empId int) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestPICParameter, err error)
	FindEmployeeLeaveRequestPICById(id uint) (employeeLeaveRequestOutput model.SelectEmployeeLeaveRequestPICParameter, err error)
	CountEmployeeLeaveRequestDaftarCuti() (count int64, err error)
	FindEmployeeLeaveRequestsDaftarCutiOffset(limit int, offset int, order string, dir string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	SearchEmployeeLeaveRequestDaftarCuti(limit int, offset int, order string, dir string, search string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	CountSearchEmployeeLeaveRequestDaftarCuti(search string) (count int64, err error)
	FindEmployeeLeaveRequestByEmpId(empId uint, year string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	FindEmployeeLeaveRequestByEmpIdDate(empId uint, year string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	FindEmployeeLeaveRequestByEmpIdExcDate(elrId uint, empId uint, year string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	FindEmployeeLeaveRequestDraftOrAFA(empId uint) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	FindEmployeeLeaveRequestPeriodDate(date float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	CountEmployeeLeaveRequestPersonaliaAllPIC() (count int64, err error)
	FindEmployeeLeaveRequestsPersonaliaPICOffset(limit int, offset int, order string, dir string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestPICParameter, err error)
	CountSearchEmployeeLeaveRequestPersonaliaPIC(search string) (count int64, err error)
	SearchEmployeeLeaveRequestPersonaliaPIC(limit int, offset int, order string, dir string, search string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestPICParameter, err error)
	FindAllEmployeeLeaveRequestByEmpId(empId uint, year string, date float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	FindEmployeeLeaveRequestFromDate(date float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	FindEmployeeLeaveRequestToDate(date float64, date2 float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	FindEmployeeLeaveRequestDate(date float64, date2 float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	FindAllEmployeeLeaveRequestByEmpIdDate(empId uint, year string, date float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	// FindExcEmployeeLeaveRequest(nwtId uint, id uint) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	// FindEmployeeLeaveRequestByNWTId(nwtId uint) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error)
	// CountEmployeeLeaveRequestName(search string) (count int64, err error)
	InsertEmployeeLeaveRequest(employeeLeaveRequest model.CreateEmployeeLeaveRequestParameter) (employeeLeaveRequestOutput model.EmployeeLeaveRequest, err error)
	UpdateEmployeeLeaveRequest(employeeLeaveRequest model.EmployeeLeaveRequest, id uint) (employeeLeaveRequestOutput model.EmployeeLeaveRequest, err error)
	DeleteEmployeeLeaveRequest(employeeLeaveRequest model.EmployeeLeaveRequest, id uint) (employeeLeaveRequestOutput model.EmployeeLeaveRequest, err error)
}

type employeeLeaveRequestService struct {
	employeeLeaveRequestRepository repository.EmployeeLeaveRequestRepository
}

func NewEmployeeLeaveRequestService(employeeLeaveRequestRep repository.EmployeeLeaveRequestRepository) EmployeeLeaveRequestService {
	return &employeeLeaveRequestService{
		employeeLeaveRequestRepository: employeeLeaveRequestRep,
	}
}

func (service *employeeLeaveRequestService) CountEmployeeLeaveRequestAll(empId int) (count int64, err error) {
	res, err := service.employeeLeaveRequestRepository.CountEmployeeLeaveRequestAll(empId)
	return res, err
}

func (service *employeeLeaveRequestService) CountEmployeeLeaveRequestApprovalAllDt(year string) (count int64, err error) {
	res, err := service.employeeLeaveRequestRepository.CountEmployeeLeaveRequestApprovalAllDt(year)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequests() (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequests()
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestsOffset(limit int, offset int, order string, dir string, empId int) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestsOffset(limit, offset, order, dir, empId)
	return res, err
}

func (service *employeeLeaveRequestService) SearchEmployeeLeaveRequest(limit int, offset int, order string, dir string, search string, empId int) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.SearchEmployeeLeaveRequest(limit, offset, order, dir, search, empId)
	return res, err
}

func (service *employeeLeaveRequestService) CountSearchEmployeeLeaveRequest(search string, empId int) (count int64, err error) {
	res, err := service.employeeLeaveRequestRepository.CountSearchEmployeeLeaveRequest(search, empId)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestById(id uint) (employeeLeaveRequestOutput model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestById(id)
	return res, err
}

func (service *employeeLeaveRequestService) CountEmployeeLeaveRequestAllPIC(empId int) (count int64, err error) {
	res, err := service.employeeLeaveRequestRepository.CountEmployeeLeaveRequestAllPIC(empId)
	return res, err
}
func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestsPICOffset(limit int, offset int, order string, dir string, empId int) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestPICParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestsPICOffset(limit, offset, order, dir, empId)
	return res, err
}
func (service *employeeLeaveRequestService) CountSearchEmployeeLeaveRequestPIC(search string, empId int) (count int64, err error) {
	res, err := service.employeeLeaveRequestRepository.CountSearchEmployeeLeaveRequestPIC(search, empId)
	return res, err
}
func (service *employeeLeaveRequestService) SearchEmployeeLeaveRequestPIC(limit int, offset int, order string, dir string, search string, empId int) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestPICParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.SearchEmployeeLeaveRequestPIC(limit, offset, order, dir, search, empId)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestPICById(id uint) (employeeLeaveRequestOutput model.SelectEmployeeLeaveRequestPICParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestPICById(id)
	return res, err
}

func (service *employeeLeaveRequestService) CountEmployeeLeaveRequestDaftarCuti() (count int64, err error) {
	res, err := service.employeeLeaveRequestRepository.CountEmployeeLeaveRequestDaftarCuti()
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestsDaftarCutiOffset(limit int, offset int, order string, dir string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestsDaftarCutiOffset(limit, offset, order, dir)
	return res, err
}

func (service *employeeLeaveRequestService) CountSearchEmployeeLeaveRequestDaftarCuti(search string) (count int64, err error) {
	res, err := service.employeeLeaveRequestRepository.CountSearchEmployeeLeaveRequestDaftarCuti(search)
	return res, err
}

func (service *employeeLeaveRequestService) SearchEmployeeLeaveRequestDaftarCuti(limit int, offset int, order string, dir string, search string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.SearchEmployeeLeaveRequestDaftarCuti(limit, offset, order, dir, search)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestByEmpId(empId uint, year string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestByEmpId(empId, year)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestByEmpIdDate(empId uint, year string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestByEmpIdDate(empId, year)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestByEmpIdExcDate(elrId uint, empId uint, year string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestByEmpIdExcDate(elrId, empId, year)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestDraftOrAFA(empId uint) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestDraftOrAFA(empId)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestPeriodDate(date float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestPeriodDate(date)
	return res, err
}

func (service *employeeLeaveRequestService) CountEmployeeLeaveRequestPersonaliaAllPIC() (count int64, err error) {
	res, err := service.employeeLeaveRequestRepository.CountEmployeeLeaveRequestPersonaliaAllPIC()
	return res, err
}
func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestsPersonaliaPICOffset(limit int, offset int, order string, dir string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestPICParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestsPersonaliaPICOffset(limit, offset, order, dir)
	return res, err
}
func (service *employeeLeaveRequestService) CountSearchEmployeeLeaveRequestPersonaliaPIC(search string) (count int64, err error) {
	res, err := service.employeeLeaveRequestRepository.CountSearchEmployeeLeaveRequestPersonaliaPIC(search)
	return res, err
}
func (service *employeeLeaveRequestService) SearchEmployeeLeaveRequestPersonaliaPIC(limit int, offset int, order string, dir string, search string) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestPICParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.SearchEmployeeLeaveRequestPersonaliaPIC(limit, offset, order, dir, search)
	return res, err
}

func (service *employeeLeaveRequestService) FindAllEmployeeLeaveRequestByEmpId(empId uint, year string, date float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindAllEmployeeLeaveRequestByEmpId(empId, year, date)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestFromDate(date float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestFromDate(date)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestToDate(date float64, date2 float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestToDate(date, date2)
	return res, err
}

func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestDate(date float64, date2 float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestDate(date, date2)
	return res, err
}

func (service *employeeLeaveRequestService) FindAllEmployeeLeaveRequestByEmpIdDate(empId uint, year string, date float64) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
	res, err := service.employeeLeaveRequestRepository.FindAllEmployeeLeaveRequestByEmpIdDate(empId, year, date)
	return res, err
}

// func (service *employeeLeaveRequestService) FindExcEmployeeLeaveRequest(nwtId uint, id uint) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
// 	res, err := service.employeeLeaveRequestRepository.FindExcEmployeeLeaveRequest(nwtId, id)
// 	return res, err
// }

// func (service *employeeLeaveRequestService) FindEmployeeLeaveRequestByNWTId(nwtId uint) (employeeLeaveRequestOutput []model.SelectEmployeeLeaveRequestParameter, err error) {
// 	res, err := service.employeeLeaveRequestRepository.FindEmployeeLeaveRequestByNWTId(nwtId)
// 	return res, err
// }

// func (service *employeeLeaveRequestService) CountEmployeeLeaveRequestName(search string) (count int64, err error) {
// 	res, err := service.employeeLeaveRequestRepository.CountEmployeeLeaveRequestName(search)
// 	return res, err
// }

func (service *employeeLeaveRequestService) InsertEmployeeLeaveRequest(employeeLeaveRequest model.CreateEmployeeLeaveRequestParameter) (employeeLeaveRequestOutput model.EmployeeLeaveRequest, err error) {
	newEmployeeLeaveRequest := model.EmployeeLeaveRequest{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequest, smapping.MapFields(&employeeLeaveRequest))
	if err != nil {
		return newEmployeeLeaveRequest, err1
	}
	res, err := service.employeeLeaveRequestRepository.InsertEmployeeLeaveRequest(newEmployeeLeaveRequest)
	return res, err
}

func (service *employeeLeaveRequestService) UpdateEmployeeLeaveRequest(employeeLeaveRequest model.EmployeeLeaveRequest, id uint) (employeeLeaveRequestOutput model.EmployeeLeaveRequest, err error) {
	newEmployeeLeaveRequest := model.EmployeeLeaveRequest{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequest, smapping.MapFields(&employeeLeaveRequest))
	if err != nil {
		return newEmployeeLeaveRequest, err1
	}
	res, err := service.employeeLeaveRequestRepository.UpdateEmployeeLeaveRequest(newEmployeeLeaveRequest, id)
	return res, err
}

func (service *employeeLeaveRequestService) DeleteEmployeeLeaveRequest(employeeLeaveRequest model.EmployeeLeaveRequest, id uint) (employeeLeaveRequestOutput model.EmployeeLeaveRequest, err error) {
	newEmployeeLeaveRequest := model.EmployeeLeaveRequest{}
	err1 := smapping.FillStruct(&newEmployeeLeaveRequest, smapping.MapFields(&employeeLeaveRequest))
	if err != nil {
		return newEmployeeLeaveRequest, err1
	}
	res, err := service.employeeLeaveRequestRepository.UpdateEmployeeLeaveRequest(newEmployeeLeaveRequest, id)
	return res, err
}
