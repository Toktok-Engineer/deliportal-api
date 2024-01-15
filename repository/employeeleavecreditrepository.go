package repository

import (
	"deliportal-api/model"
	"strings"
	"time"

	"gorm.io/gorm"
)

type EmployeeLeaveCreditRepository interface {
	CountEmployeeLeaveCreditAll(year int) (count int64, err error)
	FindEmployeeLeaveCredits(year int, empId int) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	FindEmployeeLeaveCreditsOffset(limit int, offset int, order string, dir string, year int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountSearchEmployeeLeaveCreditAll(year int, search string) (count int64, err error)
	FindSearchEmployeeLeaveCreditsOffset(limit int, offset int, order string, dir string, year int, search string) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountEmployeeLeaveCreditEmpID(year int, empId int) (count int64, err error)
	FindEmployeeLeaveCreditsOffsetByYear(limit int, offset int, order string, dir string, year int, empId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountSearchEmployeeLeaveCreditByEmp(year int, empId int, search string) (count int64, err error)
	FindSearchEmployeeLeaveCreditsOffsetByEmp(limit int, offset int, order string, dir string, year int, empId int, search string) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountSearchEmployeeLeaveCreditByDepId(year int, deptId int, positionId int, search string, groupId int) (count int64, err error)
	FindSearchEmployeeLeaveCreditsOffsetByDept(limit int, offset int, order string, dir string, year int, deptId int, positionId int, search string, groupId int, employeeid int, sectionid int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountEmployeeLeaveCreditByDepId(year int, deptId int, positionId int, groupId int) (count int64, err error)
	FindEmployeeLeaveCreditsOffsetByDept(limit int, offset int, order string, dir string, year int, deptId int, positionId int, groupId int, employeeid int, sectionid int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountEmployeeLeaveCreditByDivAll(year int, divId int, positionId int, groupId int) (count int64, err error)
	FindEmployeeLeaveCreditsOffsetByDivAll(limit int, offset int, order string, dir string, year int, divId int, positionId int, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountSearchEmployeeLeaveCreditByDivIdAll(year int, divId int, positionId int, search string, groupId int) (count int64, err error)
	FindSearchEmployeeLeaveCreditsOffsetByDivIdAll(limit int, offset int, order string, dir string, year int, divId int, positionId int, search string, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	FindEmployeeLeaveCreditsExcEmp(year int, empId string) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	FindEmployeeLeaveCreditsAll(year int) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	CountEmployeeLeaveCreditByDepAll(year int, depId int, positionId int, groupId int) (count int64, err error)
	FindEmployeeLeaveCreditsOffsetByDepAll(limit int, offset int, order string, dir string, year int, depId int, positionId int, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	CountSearchEmployeeLeaveCreditByDepIdAll(year int, depId int, positionId int, search string, groupId int) (count int64, err error)
	FindSearchEmployeeLeaveCreditsOffsetByDepIdAll(limit int, offset int, order string, dir string, year int, depId int, positionId int, search string, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error)
	// SearchEmployeeLeaveCredit(limit int, offset int, order string, dir string, search string) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	// CountSearchEmployeeLeaveCredit(search string) (count int64, err error)
	FindEmployeeLeaveCreditById(id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error)
	// CountEmployeeLeaveCreditName(search string) (count int64, err error)
	FindEmployeeLeaveCreditByEmpId(empid uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error)
	FindAllEmployeeLeaveCreditByEmpId(empid uint) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	FindExcEmployeeLeaveCredit(id uint) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error)
	InsertEmployeeLeaveCredit(employeeLeaveCredit model.EmployeeLeaveCredit) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error)
	UpdateEmployeeLeaveCredit(employeeLeaveCredit model.EmployeeLeaveCredit, id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error)
	DeleteEmployeeLeaveCredit(employeeLeaveCredit model.EmployeeLeaveCredit, id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error)
}

type EmployeeLeaveCreditConnection struct {
	connection *gorm.DB
}

func NewEmployeeLeaveCreditRepository(db *gorm.DB) EmployeeLeaveCreditRepository {
	return &EmployeeLeaveCreditConnection{
		connection: db,
	}
}

func (db *EmployeeLeaveCreditConnection) CountEmployeeLeaveCreditAll(year int) (count int64, err error) {
	res := db.connection.Debug().Table("employee_leave_credits").Where("period_year = ? AND deleted_at = 0", year).Count(&count)
	return count, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindEmployeeLeaveCredits(year int, empId int) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
	var (
		employee_leave_credits []model.EmployeeLeaveCredit
	)
	res := db.connection.Where("period_year = ? AND employee_id = ? AND deleted_at = 0", year, empId).Order("period_year").Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindEmployeeLeaveCreditsOffset(limit int, offset int, order string, dir string, year int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	var (
		orderDirection         string
		employee_leave_credits []model.SelectEmployeeLeaveCredit
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Where("employee_leave_credits.period_year = ? AND employee_leave_credits.deleted_at = 0", year).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) CountSearchEmployeeLeaveCreditAll(year int, search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Where("(lower(employee_leave_credits.period_year::varchar(50)) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE ?) AND employee_leave_credits.period_year = ? AND employee_leave_credits.deleted_at = 0", final, final, final, final, year).Count(&count)
	return count, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindSearchEmployeeLeaveCreditsOffset(limit int, offset int, order string, dir string, year int, search string) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	var (
		orderDirection         string
		final                  string
		employee_leave_credits []model.SelectEmployeeLeaveCredit
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Where("(lower(employee_leave_credits.period_year::varchar(50)) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE ?) AND employee_leave_credits.period_year = ? AND employee_leave_credits.deleted_at = 0", final, final, final, final, year).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) CountEmployeeLeaveCreditEmpID(year int, empId int) (count int64, err error) {
	res := db.connection.Debug().Table("employee_leave_credits").Where("employee_id = ? AND period_year = ? AND deleted_at = 0", empId, year).Count(&count)
	return count, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindEmployeeLeaveCreditsOffsetByYear(limit int, offset int, order string, dir string, year int, empId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	var (
		orderDirection         string
		employee_leave_credits []model.SelectEmployeeLeaveCredit
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Where("employee_leave_credits.employee_id = ? AND employee_leave_credits.period_year = ? AND employee_leave_credits.deleted_at = 0", empId, year).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) CountSearchEmployeeLeaveCreditByEmp(year int, empId int, search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Where("(lower(employee_leave_credits.period_year::varchar(50)) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE ?) AND employee_id = ? AND employee_leave_credits.period_year = ? AND employee_leave_credits.deleted_at = 0", final, final, final, final, empId, year).Count(&count)
	return count, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindSearchEmployeeLeaveCreditsOffsetByDept(limit int, offset int, order string, dir string, year int, deptId int, positionId int, search string, groupId int, employeeid int, sectionid int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	var (
		orderDirection         string
		final                  string
		employee_leave_credits []model.SelectEmployeeLeaveCredit
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Raw("SELECT employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at FROM employee_leave_credits left join employees ON employee_leave_credits.employee_id = employees.id left join departments ON employees.department_id = departments.id left join users ON employees.id = users.employee_id WHERE (lower(employee_leave_credits.period_year::varchar(50)) LIKE @final OR lower(employees.firstname) LIKE @final OR lower(employees.lastname) LIKE @final OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE @final) AND departments.id = @departmentid AND employee_leave_credits.period_year = @periodyear AND employees.id = @employeeid AND users.company_group_id = @companygroupid AND employee_leave_credits.deleted_at = 0 GROUP BY employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at UNION SELECT employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, 	employee_leave_credits.updated_at, employee_leave_credits.deleted_at FROM employee_leave_credits left join employees ON employee_leave_credits.employee_id = employees.id left join departments ON employees.department_id = departments.id left join users ON employees.id = users.employee_id join sections ON employees.section_id = sections.id WHERE  (lower(employee_leave_credits.period_year::varchar(50)) LIKE @final OR lower(employees.firstname) LIKE @final OR lower(employees.lastname) LIKE @final OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE @final) AND departments.id = @departmentid AND employee_leave_credits.period_year = @periodyear AND employees.position_id < @positionid AND users.company_group_id = @companygroupid AND sections.id = @section AND employee_leave_credits.deleted_at = 0 GROUP BY employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at",
		map[string]interface{}{"departmentid": deptId, "periodyear": year, "employeeid": employeeid, "companygroupid": groupId, "positionid": positionId, "section": sectionid, "final": final}).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) CountSearchEmployeeLeaveCreditByDepId(year int, deptId int, positionId int, search string, groupId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join users ON employees.id = users.employee_id").Where("(lower(employee_leave_credits.period_year::varchar(50)) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE ?) AND departments.id = ? AND employee_leave_credits.period_year = ? AND employees.position_id <= ? AND users.company_group_id = ? AND employee_leave_credits.deleted_at = 0", final, final, final, final, deptId, year, positionId, groupId).Count(&count)
	return count, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindSearchEmployeeLeaveCreditsOffsetByEmp(limit int, offset int, order string, dir string, year int, empId int, search string) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	var (
		orderDirection         string
		final                  string
		employee_leave_credits []model.SelectEmployeeLeaveCredit
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Where("(lower(employee_leave_credits.period_year::varchar(50)) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE ?) AND employee_id = ? AND employee_leave_credits.period_year = ? AND employee_leave_credits.deleted_at = 0", final, final, final, final, empId, year).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) CountEmployeeLeaveCreditByDepId(year int, deptId int, positionId int, groupId int) (count int64, err error) {
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join users ON employees.id = users.employee_id").Where("departments.id = ? AND employee_leave_credits.period_year = ? AND employees.position_id <= ? AND users.company_group_id = ? AND employee_leave_credits.deleted_at = 0", deptId, year, positionId, groupId).Count(&count)
	return count, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindEmployeeLeaveCreditsOffsetByDept(limit int, offset int, order string, dir string, year int, deptId int, positionId int, groupId int, employeeid int, sectionid int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	var (
		orderDirection         string
		employee_leave_credits []model.SelectEmployeeLeaveCredit
	)
	orderDirection = order + " " + dir
	res := db.connection.Raw("SELECT employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at FROM employee_leave_credits left join employees ON employee_leave_credits.employee_id = employees.id left join departments ON employees.department_id = departments.id left join users ON employees.id = users.employee_id WHERE departments.id = @departmentid AND employee_leave_credits.period_year = @periodyear AND employees.id = @employeeid AND users.company_group_id = @companygroupid AND employee_leave_credits.deleted_at = 0 GROUP BY employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at UNION SELECT employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, 	employee_leave_credits.updated_at, employee_leave_credits.deleted_at FROM employee_leave_credits left join employees ON employee_leave_credits.employee_id = employees.id left join departments ON employees.department_id = departments.id left join users ON employees.id = users.employee_id join sections ON employees.section_id = sections.id WHERE departments.id = @departmentid AND employee_leave_credits.period_year = @periodyear AND employees.position_id < @positionid AND users.company_group_id = @companygroupid AND sections.id = @section AND employee_leave_credits.deleted_at = 0 GROUP BY employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at",
		map[string]interface{}{"departmentid": deptId, "periodyear": year, "employeeid": employeeid, "companygroupid": groupId, "positionid": positionId, "section": sectionid}).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) CountEmployeeLeaveCreditByDivAll(year int, divId int, positionId int, groupId int) (count int64, err error) {
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join users ON employees.id = users.employee_id").Where("divisions.id = ? AND employee_leave_credits.period_year = ? AND employees.position_id <= ? AND users.company_group_id = ? AND employee_leave_credits.deleted_at = 0", divId, year, positionId, groupId).Count(&count)
	return count, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindEmployeeLeaveCreditsOffsetByDivAll(limit int, offset int, order string, dir string, year int, divId int, positionId int, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	var (
		orderDirection         string
		employee_leave_credits []model.SelectEmployeeLeaveCredit
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join users ON employees.id = users.employee_id").Where("divisions.id = ? AND employee_leave_credits.period_year = ? AND employees.position_id <= ? AND users.company_group_id = ? AND employee_leave_credits.deleted_at = 0", divId, year, positionId, groupId).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindSearchEmployeeLeaveCreditsOffsetByDivIdAll(limit int, offset int, order string, dir string, year int, divId int, positionId int, search string, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	var (
		orderDirection         string
		final                  string
		employee_leave_credits []model.SelectEmployeeLeaveCredit
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join users ON employees.id = users.employee_id").Where("(lower(employee_leave_credits.period_year::varchar(50)) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE ?) AND divisions.id = ? AND employee_leave_credits.period_year = ? AND employees.position_id <= ? AND users.company_group_id = ? AND employee_leave_credits.deleted_at = 0", final, final, final, final, divId, year, positionId, groupId).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) CountSearchEmployeeLeaveCreditByDivIdAll(year int, divId int, positionId int, search string, groupId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join users ON employees.id = users.employee_id").Where("(lower(employee_leave_credits.period_year::varchar(50)) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE ?) AND divisions.id = ? AND employee_leave_credits.period_year = ? AND employees.position_id <= ? AND users.company_group_id = ? AND employee_leave_credits.deleted_at = 0", final, final, final, final, divId, year, positionId, groupId).Count(&count)
	return count, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindEmployeeLeaveCreditsExcEmp(year int, empId string) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
	var (
		employee_leave_credits []model.EmployeeLeaveCredit
	)
	empid := strings.Split(empId, ",")
	res := db.connection.Where("period_year = ? AND employee_id NOT IN (?) AND deleted_at = 0", year, empid).Order("period_year").Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindEmployeeLeaveCreditsAll(year int) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
	var (
		employee_leave_credits []model.EmployeeLeaveCredit
	)
	res := db.connection.Where("period_year = ? AND deleted_at = 0", year).Order("period_year").Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) CountEmployeeLeaveCreditByDepAll(year int, depId int, positionId int, groupId int) (count int64, err error) {
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join users ON employees.id = users.employee_id").Where("departments.id = ? AND employee_leave_credits.period_year = ? AND employees.position_id <= ? AND users.company_group_id = ? AND employee_leave_credits.deleted_at = 0", depId, year, positionId, groupId).Count(&count)
	return count, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindEmployeeLeaveCreditsOffsetByDepAll(limit int, offset int, order string, dir string, year int, depId int, positionId int, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	var (
		orderDirection         string
		employee_leave_credits []model.SelectEmployeeLeaveCredit
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join users ON employees.id = users.employee_id").Where("departments.id = ? AND employee_leave_credits.period_year = ? AND employees.position_id <= ? AND users.company_group_id = ? AND employee_leave_credits.deleted_at = 0", depId, year, positionId, groupId).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindSearchEmployeeLeaveCreditsOffsetByDepIdAll(limit int, offset int, order string, dir string, year int, depId int, positionId int, search string, groupId int) (employeeLeaveCreditOutput []model.SelectEmployeeLeaveCredit, err error) {
	var (
		orderDirection         string
		final                  string
		employee_leave_credits []model.SelectEmployeeLeaveCredit
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join users ON employees.id = users.employee_id").Where("(lower(employee_leave_credits.period_year::varchar(50)) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE ?) AND departments.id = ? AND employee_leave_credits.period_year = ? AND employees.position_id <= ? AND users.company_group_id = ? AND employee_leave_credits.deleted_at = 0", final, final, final, final, depId, year, positionId, groupId).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) CountSearchEmployeeLeaveCreditByDepIdAll(year int, depId int, positionId int, search string, groupId int) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employee_leave_credits").Select("employee_leave_credits.id, employee_leave_credits.employee_id, employees.firstname, employees.lastname, employee_leave_credits.period_year, employee_leave_credits.annual_leave_limit, employee_leave_credits.annual_leave_used, employee_leave_credits.annual_leave_credit, employee_leave_credits.remark, employee_leave_credits.created_user_id, employee_leave_credits.updated_user_id, employee_leave_credits.deleted_user_id, employee_leave_credits.created_at, employee_leave_credits.updated_at, employee_leave_credits.deleted_at").Joins("left join employees ON employee_leave_credits.employee_id = employees.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join users ON employees.id = users.employee_id").Where("(lower(employee_leave_credits.period_year::varchar(50)) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(employee_leave_credits.annual_leave_credit::varchar(50)) LIKE ?) AND departments.id = ? AND employee_leave_credits.period_year = ? AND employees.position_id <= ? AND users.company_group_id = ? AND employee_leave_credits.deleted_at = 0", final, final, final, final, depId, year, positionId, groupId).Count(&count)
	return count, res.Error
}

// func (db *EmployeeLeaveCreditConnection) SearchEmployeeLeaveCredit(limit int, offset int, order string, dir string, search string) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
// 	var (
// 		orderDirection         string
// 		final                  string
// 		employee_leave_credits []model.EmployeeLeaveCredit
// 	)
// 	orderDirection = order + " " + dir
// 	final = "%" + strings.ToLower(search) + "%"
// 	res := db.connection.Where("(lower(period_year) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&employee_leave_credits)
// 	return employee_leave_credits, res.Error
// }

// func (db *EmployeeLeaveCreditConnection) CountSearchEmployeeLeaveCredit(search string) (count int64, err error) {
// 	final := "%" + strings.ToLower(search) + "%"
// 	res := db.connection.Debug().Table("employee_leave_credits").Where("(lower(period_year) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final).Count(&count)
// 	return count, res.Error
// }

func (db *EmployeeLeaveCreditConnection) FindEmployeeLeaveCreditById(id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error) {
	var (
		employeeLeaveCredit model.EmployeeLeaveCredit
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&employeeLeaveCredit)
	return employeeLeaveCredit, res.Error
}

// func (db *EmployeeLeaveCreditConnection) CountEmployeeLeaveCreditName(search string) (count int64, err error) {
// 	final := "%" + strings.ToLower(search) + "%"
// 	res := db.connection.Debug().Table("employee_leave_credits").Where("lower(period_year) LIKE ?", final).Count(&count)
// 	return count, res.Error
// }

func (db *EmployeeLeaveCreditConnection) FindEmployeeLeaveCreditByEmpId(empid uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error) {
	var (
		employee_leave_credit model.EmployeeLeaveCredit
	)
	res := db.connection.Where("employee_id=? AND period_year = ? AND deleted_at = 0 ", empid, time.Now().Year()).Order("period_year").Take(&employee_leave_credit)
	return employee_leave_credit, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindAllEmployeeLeaveCreditByEmpId(empid uint) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
	var (
		employee_leave_credit []model.EmployeeLeaveCredit
	)
	res := db.connection.Where("employee_id=? AND deleted_at = 0", empid).Order("period_year").Find(&employee_leave_credit)
	return employee_leave_credit, res.Error
}

func (db *EmployeeLeaveCreditConnection) FindExcEmployeeLeaveCredit(id uint) (employeeLeaveCreditOutput []model.EmployeeLeaveCredit, err error) {
	var (
		employee_leave_credits []model.EmployeeLeaveCredit
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("period_year").Find(&employee_leave_credits)
	return employee_leave_credits, res.Error
}

func (db *EmployeeLeaveCreditConnection) InsertEmployeeLeaveCredit(employeeLeaveCredit model.EmployeeLeaveCredit) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error) {
	res := db.connection.Save(&employeeLeaveCredit)
	return employeeLeaveCredit, res.Error
}

func (db *EmployeeLeaveCreditConnection) UpdateEmployeeLeaveCredit(employeeLeaveCredit model.EmployeeLeaveCredit, id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error) {
	var (
		employee_leave_credits model.EmployeeLeaveCredit
	)
	res := db.connection.Model(&employee_leave_credits).Where("id=?", id).Updates(map[string]interface{}{"employee_id": employeeLeaveCredit.EmployeeID, "period_year": employeeLeaveCredit.PeriodYear, "annual_leave_limit": employeeLeaveCredit.AnnualLeaveLimit, "annual_leave_used": employeeLeaveCredit.AnnualLeaveUsed, "annual_leave_credit": employeeLeaveCredit.AnnualLeaveCredit, "remark": employeeLeaveCredit.Remark, "updated_user_id": employeeLeaveCredit.UpdatedUserID, "updated_at": employeeLeaveCredit.UpdatedAt})
	return employeeLeaveCredit, res.Error
}

func (db *EmployeeLeaveCreditConnection) DeleteEmployeeLeaveCredit(employeeLeaveCredit model.EmployeeLeaveCredit, id uint) (employeeLeaveCreditOutput model.EmployeeLeaveCredit, err error) {
	var (
		employee_leave_credits model.EmployeeLeaveCredit
	)
	res := db.connection.Model(&employee_leave_credits).Where("id=?", id).Updates(map[string]interface{}{"deleted_user_id": employeeLeaveCredit.DeletedUserID, "deleted_at": employeeLeaveCredit.DeletedAt})
	return employeeLeaveCredit, res.Error
}
