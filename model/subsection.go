package model

type SubSection struct {
	ID             uint    `json:"sub_section_id"`
	SubSectionName string  `gorm:"type:varchar(50);not null" json:"sub_section_name"`
	DivisionID     uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"division_id"`
	DepartmentID   uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"department_id"`
	SectionID      uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"section_id"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectSubSectionParameter struct {
	ID             uint    `json:"sub_section_id"`
	SubSectionName string  `gorm:"type:varchar(50);not null" json:"sub_section_name"`
	DivisionID     uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	DivisionName   string  `gorm:"type:varchar(50);not null" json:"division_name"`
	DepartmentID   uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"department_id"`
	DepartmentName string  `gorm:"type:varchar(50);not null" json:"department_name"`
	SectionID      uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"section_id"`
	SectionName    string  `gorm:"type:varchar(50);not null" json:"section_name"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateSubSectionParameter struct {
	SubSectionName string  `gorm:"type:varchar(50);not null" json:"sub_section_name"`
	DivisionID     uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"division_id"`
	DepartmentID   uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"department_id"`
	SectionID      uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"section_id"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
