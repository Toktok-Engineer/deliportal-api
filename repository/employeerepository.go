package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CountEmployeeAll() (count int64, err error)
	FindEmployees() (employeeOutput []model.SelectEmployeeParameter, err error)
	FindEmployeesOffset(limit int, offset int, order string, dir string) (employeeOutput []model.SelectEmployeeParameter, err error)
	SearchEmployee(limit int, offset int, order string, dir string, search string) (employeeOutput []model.SelectEmployeeParameter, err error)
	CountSearchEmployee(search string) (count int64, err error)
	FindEmployeeById(id uint) (employeeOutput model.SelectEmployeeParameter, err error)
	FindEmployeeByNik(nik string) (employeeOutput model.SelectEmployeeParameter, err error)
	FindExcEmployee(id uint) (employeeOutput []model.SelectEmployeeParameter, err error)
	FindEmployeeByDepartment(department int) (employeeOutput []model.SelectEmployeeParameter, err error)
	FindEmployeeByPosition(group int, division string, department string, position string) (employeeOutput model.SelectEmployeeParameter, err error)
	FindEmployeeByDivIdAndDepId(Divid uint, DepId uint) (employeeOutput []model.SelectEmployeeParameter, err error)
	FindEmployeeCuti(group int, subsection int, section int, division int, department int, position int) (employeeOutput []model.SelectEmployeeCuti, err error)
	FindEmployeeByDate(date float64) (employeeOutput []model.SelectEmployeeParameter, err error)
	InsertEmployee(employee model.Employee) (employeeOutput model.Employee, err error)
	UpdateEmployee(employee model.Employee, id uint) (employeeOutput model.Employee, err error)
}

type EmployeeConnection struct {
	connection *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &EmployeeConnection{
		connection: db,
	}
}

func (db *EmployeeConnection) CountEmployeeAll() (count int64, err error) {
	res := db.connection.Debug().Table("employees").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *EmployeeConnection) FindEmployees() (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		employees []model.SelectEmployeeParameter
	)
	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.deleted_at = 0").Order("employees.firstname").Find(&employees)
	return employees, res.Error
}

func (db *EmployeeConnection) FindEmployeesOffset(limit int, offset int, order string, dir string) (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		orderDirection string
		employees      []model.SelectEmployeeParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&employees)
	return employees, res.Error
}

func (db *EmployeeConnection) SearchEmployee(limit int, offset int, order string, dir string, search string) (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		orderDirection string
		final          string
		employees      []model.SelectEmployeeParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("(lower(employees.nik) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(divisions.division_name) LIKE ? OR lower(departments.department_name) LIKE ? OR lower(sections.section_name) LIKE ? OR lower(positions.position_name) LIKE ? OR lower(locations.location_name) LIKE ? OR lower(employees.email) LIKE ? OR lower(employees.remark) LIKE ?) AND departments.deleted_at = 0 AND employees.deleted_at = 0", final, final, final, final, final, final, final, final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&employees)
	return employees, res.Error
}

func (db *EmployeeConnection) CountSearchEmployee(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("(lower(employees.nik) LIKE ? OR lower(employees.firstname) LIKE ? OR lower(employees.lastname) LIKE ? OR lower(divisions.division_name) LIKE ? OR lower(departments.department_name) LIKE ? OR lower(sections.section_name) LIKE ? OR lower(positions.position_name) LIKE ? OR lower(locations.location_name) LIKE ? OR lower(employees.email) LIKE ? OR lower(employees.remark) LIKE ?) AND departments.deleted_at = 0 AND employees.deleted_at = 0", final, final, final, final, final, final, final, final, final, final).Count(&count)
	return count, res.Error
}

func (db *EmployeeConnection) FindEmployeeById(id uint) (employeeOutput model.SelectEmployeeParameter, err error) {
	var (
		employee model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.id=? AND employees.deleted_at = 0", id).Take(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) FindEmployeeByNik(nik string) (employeeOutput model.SelectEmployeeParameter, err error) {
	var (
		employee model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.nik=? AND employees.deleted_at = 0", nik).Take(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) FindExcEmployee(id uint) (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		employees []model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.id != ? AND employees.deleted_at = 0", id).Find(&employees)
	return employees, res.Error
}

func (db *EmployeeConnection) FindEmployeeByDepartment(department int) (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		employee []model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.department_id=? AND employees.deleted_at = 0", department).Find(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) FindEmployeeByPosition(group int, division string, department string, position string) (employeeOutput model.SelectEmployeeParameter, err error) {
	var (
		employee model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join users on employees.id = users.employee_id").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("users.company_group_id = ? AND divisions.division_name LIKE ? AND departments.department_name=? AND positions.position_name=? AND employees.deleted_at = 0", group, division, department, position).Find(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) FindEmployeeByDivIdAndDepId(Divid uint, DepId uint) (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		employee []model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("employees.division_id=? AND employees.department_id = ? AND employees.deleted_at = 0", Divid, DepId).Find(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) FindEmployeeCuti(group int, subsection int, section int, division int, department int, position int) (employeeOutput []model.SelectEmployeeCuti, err error) {
	var (
		employee []model.SelectEmployeeCuti
	)

	res := db.connection.Raw("SELECT MIN(CONCAT(employees.firstname, ' ', employees.lastname, ',', employees.id))::varchar(150) AS fullname, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name FROM employees JOIN users ON employees.id = users.employee_id JOIN sections ON employees.section_id = sections.id JOIN sub_sections ON employees.sub_section_id = sub_sections.id JOIN divisions ON employees.division_id = divisions.id JOIN departments ON employees.department_id = departments.id JOIN positions ON employees.position_id = positions.id WHERE employees.division_id = @divisionid AND employees.department_id = @departmentid AND employees.section_id = @sectionid AND employees.sub_section_id = @subsectionid AND employees.position_id > @positionid AND users.company_group_id = @groupid AND employees.deleted_at = 0 GROUP BY employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name UNION SELECT MIN(CONCAT(employees.firstname, ' ', employees.lastname, ',', employees.id))::varchar(150) AS fullname, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name FROM employees JOIN users ON employees.id = users.employee_id JOIN sections ON employees.section_id = sections.id JOIN sub_sections ON employees.sub_section_id = sub_sections.id JOIN divisions ON employees.division_id = divisions.id JOIN departments ON employees.department_id = departments.id JOIN positions ON employees.position_id = positions.id WHERE employees.division_id = @divisionid AND employees.department_id = @departmentid AND employees.section_id = @sectionid AND sub_sections.sub_section_name = 'All' AND employees.position_id > @positionid AND users.company_group_id = @groupid AND employees.deleted_at = 0 GROUP BY employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name UNION SELECT MIN(CONCAT(employees.firstname, ' ', employees.lastname, ',', employees.id))::varchar(150) AS fullname, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name FROM employees JOIN users ON employees.id = users.employee_id JOIN sections ON employees.section_id = sections.id JOIN sub_sections ON employees.sub_section_id = sub_sections.id JOIN divisions ON employees.division_id = divisions.id JOIN departments ON employees.department_id = departments.id JOIN positions ON employees.position_id = positions.id WHERE employees.division_id = @divisionid AND employees.department_id = @departmentid AND sections.section_name = 'All' AND employees.position_id > @positionid AND users.company_group_id = @groupid AND employees.deleted_at = 0 GROUP BY employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name UNION SELECT MIN(CONCAT(employees.firstname, ' ', employees.lastname, ',', employees.id))::varchar(150) AS fullname, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name FROM employees JOIN users ON employees.id = users.employee_id JOIN sections ON employees.section_id = sections.id JOIN sub_sections ON employees.sub_section_id = sub_sections.id JOIN divisions ON employees.division_id = divisions.id JOIN departments ON employees.department_id = departments.id JOIN positions ON employees.position_id = positions.id WHERE employees.division_id = @divisionid AND departments.department_name = 'All' AND employees.position_id > @positionid AND users.company_group_id = @groupid AND employees.deleted_at = 0 GROUP BY employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.position_id, positions.position_name ORDER BY position_id",
		map[string]interface{}{"subsectionid": subsection, "divisionid": division, "departmentid": department, "sectionid": section, "positionid": position, "groupid": group}).Find(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) FindEmployeeByDate(date float64) (employeeOutput []model.SelectEmployeeParameter, err error) {
	var (
		employee []model.SelectEmployeeParameter
	)

	res := db.connection.Debug().Table("employees").Select("employees.id, employees.nik, employees.firstname, employees.lastname, employees.initials, employees.signature, employees.division_id, divisions.division_name, employees.department_id, departments.department_name, employees.section_id, sections.section_name, employees.sub_section_id, sub_sections.sub_section_name, employees.position_id, positions.position_name, employees.location_id, locations.location_name, employees.email, employees.remark, employees.created_user_id, employees.updated_user_id, employees.deleted_user_id, employees.created_at, employees.updated_at, employees.deleted_at, to_char(to_timestamp(employees.joined_at::numeric), 'DD-Mon-YYYY') as joined_at, employees.jenis_kelamin").Joins("left join divisions ON employees.division_id = divisions.id").Joins("left join departments ON employees.department_id = departments.id").Joins("left join sections ON employees.section_id = sections.id").Joins("left join sub_sections ON employees.sub_section_id = sub_sections.id").Joins("left join positions ON employees.position_id = positions.id").Joins("left join locations ON employees.location_id = locations.id").Where("to_timestamp(employees.joined_at) + interval '1 year' > CURRENT_DATE AND to_timestamp(joined_at) > to_timestamp(?) - interval '1 year' AND employees.deleted_at = 0", date).Find(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) InsertEmployee(employee model.Employee) (employeeOutput model.Employee, err error) {
	res := db.connection.Save(&employee)
	return employee, res.Error
}

func (db *EmployeeConnection) UpdateEmployee(employee model.Employee, id uint) (employeeOutput model.Employee, err error) {
	res := db.connection.Where("id=?", id).Updates(&employee)
	return employee, res.Error
}
