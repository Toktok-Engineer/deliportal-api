package model

type InternalMemo struct {
	ID                        uint    `json:"internal_memo_id"`
	InternalMemoNo            string  `gorm:"type:varchar(50);not null;unique" json:"internal_memo_no" binding:"required"`
	InternalMemoDate          float64 `gorm:"type:double precision;not null" json:"internal_memo_date"`
	InternalMemoTypeID        uint    `json:"internal_memo_type_id"`
	PerihalOthers             string  `gorm:"type:varchar(100);" json:"perihal"`
	ReturnToEmployeeID        uint    `json:"return_to_employee_id"`
	DocumentNo                string  `gorm:"type:varchar(50);" json:"document_no"`
	CompanyID                 uint    `json:"company_id"`
	RelatedParty              string  `gorm:"type:varchar(150);" json:"related_party"`
	Amount                    float64 `gorm:"type:double precision;" json:"amount"`
	AmountNote                string  `gorm:"type:varchar(200)" json:"amount_note"`
	InternalMemoDescription   string  `gorm:"type:text" json:"internal_memo_description"`
	InternalMemoResult        string  `gorm:"type:text" json:"internal_memo_result"`
	TotalCirculation          uint    `gorm:"type:bigint;not null" json:"total_circulation"`
	LastCirculationSequenceNo uint    `gorm:"type:bigint;not null" json:"last_circulation_sequence_no"`
	LastCirculationEmployeeID uint    `gorm:"type:bigint;not null" json:"last_circulation_employee_id"`
	Status                    int     `json:"status"`
	Remark                    string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID             uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID             uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID             uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                 float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                 float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                 float64 `gorm:"type:double precision;" json:"deleted_at"`
	FileName                  string  `gorm:"type:varchar(200)" json:"file_name"`
	FileUrl                   string  `gorm:"type:varchar(200)" json:"file_url"`
}

type SelectInternalMemoParameter struct {
	ID                          uint    `json:"internal_memo_id"`
	InternalMemoNo              string  `gorm:"type:varchar(50);not null;unique" json:"internal_memo_no" binding:"required"`
	InternalMemoDate            string  `gorm:"type:varchar(100);typedefault:null" json:"internal_memo_date"`
	InternalMemoTypeID          uint    `json:"internal_memo_type_id"`
	InternalMemoTypeName        string  `gorm:"type:varchar(150);not null;unique" json:"internal_memo_type_name"`
	PerihalOthers               string  `gorm:"type:varchar(100);" json:"perihal"`
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

type CreateInternalMemoParameter struct {
	InternalMemoNo            string  `gorm:"type:varchar(50);not null;unique" json:"internal_memo_no" binding:"required"`
	InternalMemoDate          float64 `gorm:"type:double precision;not null" json:"internal_memo_date"`
	InternalMemoTypeID        uint    `json:"internal_memo_type_id"`
	PerihalOthers             string  `gorm:"type:varchar(100);" json:"perihal"`
	ReturnToEmployeeID        uint    `json:"return_to_employee_id"`
	DocumentNo                string  `gorm:"type:varchar(50);" json:"document_no"`
	CompanyID                 uint    `json:"company_id"`
	RelatedParty              string  `gorm:"type:varchar(150);" json:"related_party"`
	Amount                    float64 `gorm:"type:double precision;" json:"amount"`
	AmountNote                string  `gorm:"type:varchar(200)" json:"amount_note"`
	InternalMemoDescription   string  `gorm:"type:text" json:"internal_memo_description"`
	InternalMemoResult        string  `gorm:"type:text" json:"internal_memo_result"`
	TotalCirculation          uint    `gorm:"type:bigint;not null" json:"total_circulation"`
	LastCirculationSequenceNo uint    `gorm:"type:bigint;not null" json:"last_circulation_sequence_no"`
	LastCirculationEmployeeID uint    `gorm:"type:bigint;not null" json:"last_circulation_employee_id"`
	Status                    int     `json:"status"`
	Remark                    string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID             uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID             uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID             uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                 float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                 float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                 float64 `gorm:"type:double precision;" json:"deleted_at"`
	FileName                  string  `gorm:"type:varchar(200)" json:"file_name"`
	FileUrl                   string  `gorm:"type:varchar(200)" json:"file_url"`
}
