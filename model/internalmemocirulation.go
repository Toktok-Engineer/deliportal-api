package model

type InternalMemoCirculation struct {
	ID                uint    `json:"internal_memo_circulation_id"`
	InternalMemoID    uint    `json:"internal_memo_id"`
	EmployeeID        uint    `json:"employee_id"`
	SequenceNo        uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	ApprovalDate      float64 `gorm:"type:double precision;not null" json:"approval_date"`
	CirculationStatus uint    `gorm:"type:bigint;not null" json:"circulation_status"`
	Remark            string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID     uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID     uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt         float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt         float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt         float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type SelectInternalMemoCirculationParameter struct {
	ID                    uint   `json:"internal_memo_circulation_id"`
	InternalMemoID        uint   `json:"internal_memo_id"`
	EmployeeID            uint   `json:"employee_id"`
	Firstname             string `gorm:"type:varchar(30);not null" json:"first_name"`
	Lastname              string `gorm:"type:varchar(30)" json:"last_name"`
	DivisionName          string `gorm:"type:varchar(50);not null;unique" json:"division_name"`
	DepartmentName        string `gorm:"type:varchar(50);not null;unique" json:"department_name"`
	SequenceNo            uint   `gorm:"type:bigint;not null" json:"sequence_no"`
	ApprovalDate          string `gorm:"type:varchar(100);typedefault:null" json:"approval_date"`
	CirculationStatus     uint   `gorm:"type:bigint;not null" json:"circulation_status"`
	CirculationStatusName string `gorm:"type:varchar(50);" json:"circulation_status_name"`
	Remark                string `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID         uint   `gorm:"type:bigint;not null" json:"created_user_id"`
	CreatedUser           string `gorm:"type:varchar(100);" json:"created_user"`
	UpdatedUserID         uint   `gorm:"type:bigint;not null" json:"updated_user_id"`
	UpdatedUser           string `gorm:"type:varchar(100);" json:"updated_user"`
	DeletedUserID         uint   `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	DeletedUser           string `gorm:"type:varchar(100);" json:"deleted_user"`
	CreatedAt             string `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt             string `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt             string `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type SelectInternalMemoCirculationJoinIMParameter struct {
	ID                          uint    `json:"internal_memo_id"`
	InternalMemoNo              string  `gorm:"type:varchar(50);not null;unique" json:"internal_memo_no" binding:"required"`
	InternalMemoDate            string  `gorm:"type:varchar(100);typedefault:null" json:"internal_memo_date"`
	InternalMemoTypeID          uint    `json:"internal_memo_type_id"`
	InternalMemoTypeName        string  `gorm:"type:varchar(150);not null;unique" json:"internal_memo_type_name"`
	ReturnToEmployeeID          uint    `json:"return_to_employee_id"`
	ReturnToEmployeeName        string  `gorm:"type:varchar(100);" json:"return_to_employee_name"`
	DocumentNo                  string  `gorm:"type:varchar(50);" json:"document_no"`
	CompanyID                   uint    `json:"company_id"`
	CompanyName                 string  `gorm:"type:varchar(50);not null;unique" json:"company_name"`
	RelatedParty                string  `gorm:"type:varchar(150);" json:"related_party"`
	Amount                      float64 `gorm:"type:double precision;" json:"amount"`
	AmountNote                  string  `gorm:"type:varchar(200)" json:"amount_note"`
	InternalMemoDescription     string  `gorm:"type:text" json:"internal_memo_description"`
	InternalMemoResult          string  `gorm:"type:text" json:"internal_memo_result"`
	LastDate                    string  `gorm:"type:varchar(100);typedefault:null" json:"last_date"`
	TotalCirculation            uint    `gorm:"type:bigint;not null" json:"total_circulation"`
	LastCirculationSequenceNo   uint    `gorm:"type:bigint;not null" json:"last_circulation_sequence_no"`
	LastCirculationEmployeeID   uint    `gorm:"type:bigint;not null" json:"last_circulation_employee_id"`
	LastCirculationEmployeeName string  `gorm:"type:varchar(100);" json:"last_circulation_employee_name"`
	Status                      uint    `gorm:"type:bigint;not null" json:"status"`
	StatusName                  string  `gorm:"type:varchar(50);" json:"status_name"`
	Remark                      string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID               uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	CreatedUser                 string  `gorm:"type:varchar(100);" json:"created_user"`
	UpdatedUserID               uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	UpdatedUser                 string  `gorm:"type:varchar(100);" json:"updated_user"`
	DeletedUserID               uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	DeletedUser                 string  `gorm:"type:varchar(100);" json:"deleted_user"`
	CreatedAt                   string  `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt                   string  `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt                   string  `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
	FileName                    string  `gorm:"type:varchar(200)" json:"file_name"`
	FileUrl                     string  `gorm:"type:varchar(200)" json:"file_url"`
}

type CreateInternalMemoCirculationParameter struct {
	InternalMemoID    uint    `json:"internal_memo_id"`
	EmployeeID        uint    `json:"employee_id"`
	SequenceNo        uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	ApprovalDate      float64 `gorm:"type:double precision;not null" json:"approval_date"`
	CirculationStatus uint    `gorm:"type:bigint;not null" json:"circulation_status"`
	Remark            string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID     uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID     uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt         float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt         float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt         float64 `gorm:"type:double precision;" json:"deleted_at"`
}
