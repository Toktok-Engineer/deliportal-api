package model

type InternalMemoTracing struct {
	ID                uint    `json:"internal_memo_tracing_id"`
	InternalMemoID    uint    `json:"internal_memo_id"`
	EmployeeID        uint    `json:"employee_id"`
	SequenceNo        uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	CirculationStatus uint    `gorm:"type:bigint;not null" json:"circulation_status"`
	Remark            string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID     uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID     uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt         float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt         float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt         float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type CreateInternalMemoTracingParameter struct {
	InternalMemoID    uint    `json:"internal_memo_id"`
	EmployeeID        uint    `json:"employee_id"`
	SequenceNo        uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	CirculationStatus uint    `gorm:"type:bigint;not null" json:"circulation_status"`
	Remark            string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID     uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID     uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt         float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt         float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt         float64 `gorm:"type:double precision;" json:"deleted_at"`
}
