package model

type EmployeeLeaveRequest struct {
	ID                        uint    `json:"employee_leave_request_id"`
	EmployeeID                uint    `json:"employee_id"`
	EmployeeLeaveRequestNo    string  `gorm:"type:varchar(50)" json:"employee_leave_request_no"`
	EmployeeLeaveRequestDate  uint    `gorm:"type:bigint;typedefault:null" json:"employee_leave_request_date"`
	LeaveTypeID               uint    `json:"leave_type_id"`
	LeaveCredit               *int    `json:"leave_credit"`
	LeaveUsed                 *int    `json:"leave_used"`
	LeaveRemaining            *int    `json:"leave_remaining"`
	LeaveFrom                 uint    `gorm:"type:bigint;typedefault:null" json:"leave_from"`
	LeaveTo                   uint    `gorm:"type:bigint;typedefault:null" json:"leave_to"`
	Address                   string  `gorm:"type:varchar(200)" json:"address"`
	PhoneNumber               string  `gorm:"type:varchar(20)" json:"phone_number"`
	CurrentApprovalSequenceNo uint    `json:"current_approval_sequence_no"`
	CurrentApprovalEmployeeId uint    `json:"current_approval_employee_id"`
	Status                    uint    `gorm:"type:bigint;not null" json:"status"`
	Remark                    string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID             uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID             uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID             uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                 float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                 float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                 float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
	SecondaryPIC              string  `gorm:"type:varchar(100)" json:"secondary_pic"`
}

type SelectEmployeeLeaveRequestParameter struct {
	ID                        uint    `json:"employee_leave_request_id"`
	EmployeeID                uint    `json:"employee_id"`
	Firstname                 string  `gorm:"type:varchar(30);not null" json:"first_name"`
	Lastname                  string  `gorm:"type:varchar(30)" json:"last_name"`
	EmployeeLeaveRequestNo    string  `gorm:"type:varchar(50)" json:"employee_leave_request_no"`
	EmployeeLeaveRequestDate  string  `gorm:"type:varchar(100);not null" json:"employee_leave_request_date"`
	LeaveTypeID               uint    `json:"leave_type_id"`
	LeaveTypeName             string  `gorm:"type:varchar(150);not null;unique" json:"leave_type_name"`
	LeaveCredit               *int    `json:"leave_credit"`
	LeaveUsed                 *int    `json:"leave_used"`
	LeaveRemaining            *int    `json:"leave_remaining"`
	LeaveFrom                 string  `gorm:"type:varchar(100);not null"  json:"leave_from"`
	LeaveTo                   string  `gorm:"type:varchar(100);not null" json:"leave_to"`
	Address                   string  `gorm:"type:varchar(200)" json:"address"`
	PhoneNumber               string  `gorm:"type:varchar(20)" json:"phone_number"`
	CurrentApprovalSequenceNo uint    `json:"current_approval_sequence_no"`
	CurrentApprovalEmployeeId uint    `json:"current_approval_employee_id"`
	Status                    uint    `gorm:"type:bigint;not null" json:"status"`
	StatusName                string  `gorm:"type:varchar(30);not null" json:"status_name"`
	Remark                    string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID             uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID             uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID             uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                 float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                 float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                 float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
	SecondaryPIC              string  `gorm:"type:varchar(100)" json:"secondary_pic"`
}

type SelectEmployeeLeaveRequestPICParameter struct {
	ID                             uint    `json:"employee_leave_request_id"`
	EmployeeID                     uint    `json:"employee_id"`
	EmployeeLeaveRequestApprovalID uint    `json:"employee_leave_request_approval_id"`
	Firstname                      string  `gorm:"type:varchar(30);not null" json:"first_name"`
	Lastname                       string  `gorm:"type:varchar(30)" json:"last_name"`
	EmployeeLeaveRequestNo         string  `gorm:"type:varchar(50)" json:"employee_leave_request_no"`
	EmployeeLeaveRequestDate       string  `gorm:"type:varchar(100);not null" json:"employee_leave_request_date"`
	LeaveTypeID                    uint    `json:"leave_type_id"`
	LeaveTypeName                  string  `gorm:"type:varchar(150);not null;unique" json:"leave_type_name"`
	LeaveCredit                    *int    `json:"leave_credit"`
	LeaveUsed                      *int    `json:"leave_used"`
	LeaveRemaining                 *int    `json:"leave_remaining"`
	LeaveFrom                      string  `gorm:"type:varchar(100);not null"  json:"leave_from"`
	LeaveTo                        string  `gorm:"type:varchar(100);not null" json:"leave_to"`
	Address                        string  `gorm:"type:varchar(200)" json:"address"`
	PhoneNumber                    string  `gorm:"type:varchar(20)" json:"phone_number"`
	CurrentApprovalSequenceNo      uint    `json:"current_approval_sequence_no"`
	CurrentApprovalEmployeeId      uint    `json:"current_approval_employee_id"`
	CurrentApprovalEmployeeName    string  `gorm:"type:varchar(100)" json:"current_approval_employee_name"`
	Status                         uint    `gorm:"type:bigint;not null" json:"status"`
	StatusName                     string  `gorm:"type:varchar(30);not null" json:"status_name"`
	Remark                         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID                  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID                  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID                  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
	SecondaryPIC                   string  `gorm:"type:varchar(100)" json:"secondary_pic"`
}

type CreateEmployeeLeaveRequestParameter struct {
	EmployeeID                uint    `json:"employee_id"`
	EmployeeLeaveRequestNo    string  `gorm:"type:varchar(50)" json:"employee_leave_request_no"`
	EmployeeLeaveRequestDate  uint    `gorm:"type:bigint;typedefault:null" json:"employee_leave_request_date"`
	LeaveTypeID               uint    `json:"leave_type_id"`
	LeaveCredit               *int    `json:"leave_credit"`
	LeaveUsed                 *int    `json:"leave_used"`
	LeaveRemaining            *int    `json:"leave_remaining"`
	LeaveFrom                 uint    `gorm:"type:bigint;typedefault:null" json:"leave_from"`
	LeaveTo                   uint    `gorm:"type:bigint;typedefault:null" json:"leave_to"`
	Address                   string  `gorm:"type:varchar(200)" json:"address"`
	PhoneNumber               string  `gorm:"type:varchar(20)" json:"phone_number"`
	CurrentApprovalSequenceNo uint    `json:"current_approval_sequence_no"`
	CurrentApprovalEmployeeId uint    `json:"current_approval_employee_id"`
	Status                    uint    `gorm:"type:bigint;not null" json:"status"`
	Remark                    string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID             uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID             uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID             uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                 float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                 float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                 float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
	SecondaryPIC              string  `gorm:"type:varchar(100)" json:"secondary_pic"`
}
