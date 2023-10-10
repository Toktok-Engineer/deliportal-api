package model

type EmployeeLeaveCredit struct {
	ID                uint    `json:"employee_leave_credit_id"`
	EmployeeID        uint    `json:"employee_id"`
	PeriodYear        uint    `json:"period_year"`
	AnnualLeaveLimit  *int    `json:"annual_leave_limit"`
	AnnualLeaveUsed   *int    `json:"annual_leave_used"`
	AnnualLeaveCredit *int    `json:"annual_leave_credit"`
	Remark            string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID     uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID     uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt         float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt         float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt         float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectEmployeeLeaveCredit struct {
	ID                uint    `json:"employee_leave_credit_id"`
	EmployeeID        uint    `json:"employee_id"`
	Firstname         string  `gorm:"type:varchar(30);not null" json:"first_name"`
	Lastname          string  `gorm:"type:varchar(30)" json:"last_name"`
	PeriodYear        uint    `json:"period_year"`
	AnnualLeaveLimit  *int    `json:"annual_leave_limit"`
	AnnualLeaveUsed   *int    `json:"annual_leave_used"`
	AnnualLeaveCredit *int    `json:"annual_leave_credit"`
	Remark            string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID     uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID     uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt         float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt         float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt         float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateEmployeeLeaveCreditParameter struct {
	EmployeeID        uint    `json:"employee_id"`
	PeriodYear        uint    `json:"period_year"`
	AnnualLeaveLimit  *int    `json:"annual_leave_limit"`
	AnnualLeaveUsed   *int    `json:"annual_leave_used"`
	AnnualLeaveCredit *int    `json:"annual_leave_credit"`
	Remark            string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID     uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID     uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt         float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt         float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt         float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
