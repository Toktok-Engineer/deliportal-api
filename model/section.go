package model

type Section struct {
	ID            uint    `json:"section_id"`
	SectionName   string  `gorm:"type:varchar(50);not null;unique" json:"section_name"`
	DivisionID    uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"division_id"`
	DepartmentID  uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"department_id"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectSectionParameter struct {
	ID             uint    `json:"section_id"`
	SectionName    string  `gorm:"type:varchar(50);not null;unique" json:"section_name"`
	DivisionID     uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	DivisionName   string  `gorm:"type:varchar(50);not null;unique" json:"division_name"`
	DepartmentID   uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"department_id"`
	DepartmentName string  `gorm:"type:varchar(50);not null;unique" json:"department_name"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateSectionParameter struct {
	SectionName   string  `gorm:"type:varchar(50);not null;unique" json:"section_name"`
	DivisionID    uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"division_id"`
	DepartmentID  uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"department_id"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
