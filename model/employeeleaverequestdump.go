package model

type EmployeeLeaveRequestDump struct {
	ID                            uint    `json:"employee_leave_request_dump_id"`
	EmployeeLeaveRequestID        uint    `json:"employee_leave_request_id"`
	CurrentEmployeeID             uint    `json:"current_employee_id"`
	EmployeeLeaveRequestTracingID uint    `json:"employee_leave_request_tracing_id"`
	LeaveUsed                     *int    `json:"leave_used"`
	LeaveRemaining                *int    `json:"leave_remaining"`
	LeaveFrom                     uint    `gorm:"type:bigint;typedefault:null" json:"leave_from"`
	LeaveTo                       uint    `gorm:"type:bigint;typedefault:null" json:"leave_to"`
	Status                        uint    `gorm:"type:bigint;not null" json:"status"`
	Remark                        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID                 uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID                 uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID                 uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateEmployeeLeaveRequestDumpParameter struct {
	EmployeeLeaveRequestID        uint    `json:"employee_leave_request_id"`
	CurrentEmployeeID             uint    `json:"current_employee_id"`
	EmployeeLeaveRequestTracingID uint    `json:"employee_leave_request_tracing_id"`
	LeaveUsed                     *int    `json:"leave_used"`
	LeaveRemaining                *int    `json:"leave_remaining"`
	LeaveFrom                     uint    `gorm:"type:bigint;typedefault:null" json:"leave_from"`
	LeaveTo                       uint    `gorm:"type:bigint;typedefault:null" json:"leave_to"`
	Status                        uint    `gorm:"type:bigint;not null" json:"status"`
	Remark                        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID                 uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID                 uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID                 uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
