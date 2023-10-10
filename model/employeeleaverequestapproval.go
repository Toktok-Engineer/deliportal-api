package model

type EmployeeLeaveRequestApproval struct {
	ID                     uint    `json:"employee_leave_request_approval_id"`
	EmployeeLeaveRequestID uint    `json:"employee_leave_request_id"`
	EmployeeID             uint    `json:"employee_id"`
	SequenceNo             uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	DivisionID             uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	DepartmentID           uint    `gorm:"type:bigint;foreign_key;index:" json:"department_id"`
	SectionID              uint    `gorm:"type:bigint;foreign_key;index:" json:"section_id"`
	PositionID             uint    `gorm:"type:bigint;foreign_key;index:" json:"position_id"`
	ApprovedAt             float64 `gorm:"type:double precision;not null" json:"approved_at"`
	ApprovalStatus         uint    `gorm:"type:bigint;not null" json:"approval_status"`
	Remark                 string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID          uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID          uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID          uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt              float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt              float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt              float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type SelectEmployeeLeaveRequestApprovalParameter struct {
	ID                     uint    `json:"employee_leave_request_approval_id"`
	EmployeeLeaveRequestID uint    `json:"employee_leave_request_id"`
	EmployeeID             uint    `json:"employee_id"`
	Firstname              string  `gorm:"type:varchar(30);not null" json:"first_name"`
	Lastname               string  `gorm:"type:varchar(30)" json:"last_name"`
	SequenceNo             uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	DivisionID             uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	DivisionName           string  `gorm:"type:varchar(50);not null;unique" json:"division_name"`
	DepartmentID           uint    `gorm:"type:bigint;foreign_key;index:" json:"department_id"`
	DepartmentName         string  `gorm:"type:varchar(50);not null;unique" json:"department_name"`
	SectionID              uint    `gorm:"type:bigint;foreign_key;index:" json:"section_id"`
	SectionName            string  `gorm:"type:varchar(50);not null;unique" json:"section_name"`
	PositionID             uint    `gorm:"type:bigint;foreign_key;index:" json:"position_id"`
	PositionName           string  `gorm:"type:varchar(50);not null;unique" json:"position_name"`
	ApprovedAt             string  `gorm:"type:varchar(100);not null" json:"approved_at"`
	ApprovalStatus         uint    `gorm:"type:bigint;not null" json:"approval_status"`
	ApprovalStatusName     string  `gorm:"type:varchar(50);not null;unique" json:"approval_status_name"`
	Remark                 string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID          uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID          uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID          uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt              float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt              float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt              float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type CreateEmployeeLeaveRequestApprovalParameter struct {
	EmployeeLeaveRequestID uint    `json:"employee_leave_request_id"`
	EmployeeID             uint    `json:"employee_id"`
	SequenceNo             uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	DivisionID             uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	DepartmentID           uint    `gorm:"type:bigint;foreign_key;index:" json:"department_id"`
	SectionID              uint    `gorm:"type:bigint;foreign_key;index:" json:"section_id"`
	PositionID             uint    `gorm:"type:bigint;foreign_key;index:" json:"position_id"`
	ApprovedAt             float64 `gorm:"type:double precision;not null" json:"approved_at"`
	ApprovalStatus         uint    `gorm:"type:bigint;not null" json:"approval_status"`
	Remark                 string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID          uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID          uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID          uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt              float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt              float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt              float64 `gorm:"type:double precision;" json:"deleted_at"`
}
