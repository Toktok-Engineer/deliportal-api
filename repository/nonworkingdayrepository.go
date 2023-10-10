package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type NonWorkingDayRepository interface {
	CountNonWorkingDayAll() (count int64, err error)
	FindNonWorkingDays() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	FindNonWorkingDaysCuti() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	FindNonWorkingDaysAllDate() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	FindNonWorkingDaysOffset(limit int, offset int, order string, dir string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	SearchNonWorkingDay(limit int, offset int, order string, dir string, search string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	CountSearchNonWorkingDay(search string) (count int64, err error)
	FindNonWorkingDayById(id uint) (nonWorkingDayOutput model.SelectNonWorkingDayParameter, err error)
	FindExcNonWorkingDay(nwtId uint, id uint) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	FindNonWorkingDayByNWTId(nwtId uint) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	CountNonWorkingDayName(search string) (count int64, err error)
	FindNonWorkingDaybyDate(date float64) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error)
	InsertNonWorkingDay(nonWorkingDay model.NonWorkingDay) (nonWorkingDayOutput model.NonWorkingDay, err error)
	UpdateNonWorkingDay(nonWorkingDay model.NonWorkingDay, id uint) (nonWorkingDayOutput model.NonWorkingDay, err error)
}

type NonWorkingDayConnection struct {
	connection *gorm.DB
}

func NewNonWorkingDayRepository(db *gorm.DB) NonWorkingDayRepository {
	return &NonWorkingDayConnection{
		connection: db,
	}
}

func (db *NonWorkingDayConnection) CountNonWorkingDayAll() (count int64, err error) {
	res := db.connection.Debug().Table("non_working_days").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *NonWorkingDayConnection) FindNonWorkingDays() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	var (
		non_working_days []model.SelectNonWorkingDayParameter
	)
	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY') as effective_from, non_working_days.effective_from as effective_from_unix, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("non_working_days.deleted_at = 0").Order("non_working_days.id").Find(&non_working_days)
	return non_working_days, res.Error
}

func (db *NonWorkingDayConnection) FindNonWorkingDaysCuti() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	var (
		non_working_days []model.SelectNonWorkingDayParameter
	)
	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY') as effective_from, non_working_days.effective_from as effective_from_unix, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("non_working_days.status = 3 AND non_working_types.deduct_leave = true AND non_working_days.deleted_at = 0").Order("non_working_days.id").Find(&non_working_days)
	return non_working_days, res.Error
}

func (db *NonWorkingDayConnection) FindNonWorkingDaysAllDate() (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	var (
		non_working_days []model.SelectNonWorkingDayParameter
	)
	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD-Mon-YYYY') as effective_from, non_working_days.effective_from as effective_from_unix, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD-Mon-YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("non_working_days.status = 3 AND non_working_days.deleted_at = 0").Order("non_working_days.id").Find(&non_working_days)
	return non_working_days, res.Error
}

func (db *NonWorkingDayConnection) FindNonWorkingDaysOffset(limit int, offset int, order string, dir string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	var (
		orderDirection   string
		non_working_days []model.SelectNonWorkingDayParameter
	)
	orderDirection = order + " " + dir
	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY') as effective_from, non_working_days.effective_from as effective_from_unix, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("non_working_days.deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&non_working_days)
	return non_working_days, res.Error
}

func (db *NonWorkingDayConnection) SearchNonWorkingDay(limit int, offset int, order string, dir string, search string) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	var (
		orderDirection   string
		final            string
		non_working_days []model.SelectNonWorkingDayParameter
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY') as effective_from, non_working_days.effective_from as effective_from_unix, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("(lower(non_working_types.non_working_type_name) LIKE ? OR lower(non_working_days.period_year::varchar(50)) LIKE ? OR lower(non_working_days.description) LIKE ? OR lower(to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY')) LIKE ? OR lower(to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY')) LIKE ? OR lower(non_working_days.total::varchar(50)) LIKE ? OR lower(CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END) LIKE ? OR lower(non_working_days.remark) LIKE ?) AND non_working_days.deleted_at = 0", final, final, final, final, final, final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&non_working_days)
	return non_working_days, res.Error
}

func (db *NonWorkingDayConnection) CountSearchNonWorkingDay(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY') as effective_from, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("(lower(non_working_types.non_working_type_name) LIKE ? OR lower(non_working_days.period_year::varchar(50)) LIKE ? OR lower(non_working_days.description) LIKE ? OR lower(to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY')) LIKE ? OR lower(to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY')) LIKE ? OR lower(non_working_days.total::varchar(50)) LIKE ? OR lower(CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END) LIKE ? OR lower(non_working_days.remark) LIKE ?) AND non_working_days.deleted_at = 0", final, final, final, final, final, final, final, final).Count(&count)
	return count, res.Error
}

func (db *NonWorkingDayConnection) FindNonWorkingDayById(id uint) (nonWorkingDayOutput model.SelectNonWorkingDayParameter, err error) {
	var (
		nonWorkingDay model.SelectNonWorkingDayParameter
	)

	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY') as effective_from, non_working_days.effective_from as effective_from_unix, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("non_working_days.id=? AND non_working_days.deleted_at = 0", id).Take(&nonWorkingDay)
	return nonWorkingDay, res.Error
}

func (db *NonWorkingDayConnection) FindExcNonWorkingDay(nwtId uint, id uint) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	var (
		non_working_days []model.SelectNonWorkingDayParameter
	)

	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY') as effective_from, non_working_days.effective_from as effective_from_unix, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("non_working_days.non_working_type_id = ? AND non_working_days.id!=? AND non_working_days.deleted_at = 0", nwtId, id).Order("non_working_days.id").Find(&non_working_days)
	return non_working_days, res.Error
}

func (db *NonWorkingDayConnection) FindNonWorkingDayByNWTId(nwtId uint) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	var (
		non_working_days []model.SelectNonWorkingDayParameter
	)

	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY') as effective_from, non_working_days.effective_from as effective_from_unix, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("non_working_days.non_working_type_id=? AND non_working_days.deleted_at = 0", nwtId).Order("non_working_days.id").Find(&non_working_days)
	return non_working_days, res.Error
}

func (db *NonWorkingDayConnection) CountNonWorkingDayName(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY') as effective_from, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("lower(non_working_types.non_working_type_name) LIKE ?", final).Count(&count)
	return count, res.Error
}

func (db *NonWorkingDayConnection) FindNonWorkingDaybyDate(date float64) (nonWorkingDayOutput []model.SelectNonWorkingDayParameter, err error) {
	var (
		non_working_days []model.SelectNonWorkingDayParameter
	)

	res := db.connection.Debug().Table("non_working_days").Select("non_working_days.id, non_working_days.non_working_type_id, non_working_types.non_working_type_name, non_working_days.period_year, non_working_days.description, to_char(to_timestamp(non_working_days.effective_from::numeric), 'DD Mon YYYY') as effective_from, non_working_days.effective_from as effective_from_unix, to_char(to_timestamp(non_working_days.effective_to::numeric), 'DD Mon YYYY') as effective_to, non_working_days.total, non_working_days.status, CASE WHEN non_working_days.status = 1 THEN 'Draft' WHEN non_working_days.status = 2 THEN 'Ask For Approval' WHEN non_working_days.status = 3 THEN 'Approved' WHEN non_working_days.status = 4 THEN 'Rejected' WHEN non_working_days.status = 5 THEN 'Cancelled' END AS status_name, non_working_days.remark, non_working_days.created_user_id, non_working_days.updated_user_id, non_working_days.deleted_user_id, non_working_days.created_at, non_working_days.updated_at, non_working_days.deleted_at").Joins("left join non_working_types ON non_working_days.non_working_type_id = non_working_types.id").Where("to_timestamp(non_working_days.effective_from) >= to_timestamp(?) + interval '1 year' AND non_working_days.status = 3 AND non_working_types.deduct_leave = true AND non_working_days.deleted_at = 0", date).Order("non_working_days.id").Find(&non_working_days)
	return non_working_days, res.Error
}

func (db *NonWorkingDayConnection) InsertNonWorkingDay(nonWorkingDay model.NonWorkingDay) (nonWorkingDayOutput model.NonWorkingDay, err error) {
	res := db.connection.Save(&nonWorkingDay)
	return nonWorkingDay, res.Error
}

func (db *NonWorkingDayConnection) UpdateNonWorkingDay(nonWorkingDay model.NonWorkingDay, id uint) (nonWorkingDayOutput model.NonWorkingDay, err error) {
	res := db.connection.Where("id=?", id).Updates(&nonWorkingDay)
	return nonWorkingDay, res.Error
}
