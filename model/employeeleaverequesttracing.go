package model

type EmployeeLeaveRequestTracing struct {
	ID                     uint    `json:"employee_leave_request_tracing_id"`
	EmployeeLeaveRequestID uint    `json:"employee_leave_request_id"`
	EmployeeID             uint    `json:"employee_id"`
	SequenceNo             uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	ApprovalStatus         uint    `gorm:"type:bigint;not null" json:"approval_status"`
	Remark                 string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID          uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID          uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID          uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt              float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt              float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt              float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type CreateEmployeeLeaveRequestTracingParameter struct {
	EmployeeLeaveRequestID uint    `json:"employee_leave_request_id"`
	EmployeeID             uint    `json:"employee_id"`
	SequenceNo             uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	ApprovalStatus         uint    `gorm:"type:bigint;not null" json:"approval_status"`
	Remark                 string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID          uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID          uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID          uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt              float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt              float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt              float64 `gorm:"type:double precision;" json:"deleted_at"`
}
