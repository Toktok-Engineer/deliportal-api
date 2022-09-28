package model

type ImDefaultCirculation struct {
	ID             uint    `json:"im_default_circulation_id"`
	SequenceNo     uint    `json:"sequence_no"`
	CompanyGroupID uint    `json:"company_group_id"`
	EmployeeID     uint    `json:"employee_id"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type SelectImDefaultCirculationParameter struct {
	ID             uint   `json:"im_default_circulation_id"`
	SequenceNo     uint   `json:"sequence_no"`
	CompanyGroupID uint   `json:"company_group_id"`
	EmployeeID     uint   `json:"employee_id"`
	Firstname      string `gorm:"type:varchar(30);not null" json:"first_name"`
	Lastname       string `gorm:"type:varchar(30)" json:"last_name"`
	Remark         string `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint   `gorm:"type:bigint;not null" json:"created_user_id"`
	CreatedUser    string `gorm:"type:varchar(100);" json:"created_user"`
	UpdatedUserID  uint   `gorm:"type:bigint;not null" json:"updated_user_id"`
	UpdatedUser    string `gorm:"type:varchar(100);" json:"updated_user"`
	DeletedUserID  uint   `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	DeletedUser    string `gorm:"type:varchar(100);" json:"deleted_user"`
	CreatedAt      string `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt      string `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt      string `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateImDefaultCirculationParameter struct {
	SequenceNo     uint    `json:"sequence_no"`
	CompanyGroupID uint    `json:"company_group_id"`
	EmployeeID     uint    `json:"employee_id"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;" json:"deleted_at"`
}
