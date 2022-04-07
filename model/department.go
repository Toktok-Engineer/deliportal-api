package model

type Department struct {
	ID             uint    `json:"department_id"`
	DepartmentName string  `gorm:"type:varchar(50);not null;unique" json:"department_name"`
	DivisionID     uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectDepartmentParameter struct {
	ID             uint    `json:"department_id"`
	DepartmentName string  `gorm:"type:varchar(50);not null;unique" json:"department_name"`
	DivisionID     uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	DivisionName   string  `gorm:"type:varchar(50);not null;unique" json:"division_name"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateDepartmentParameter struct {
	DepartmentName string  `gorm:"type:varchar(50);not null;unique" json:"department_name"`
	DivisionID     uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
