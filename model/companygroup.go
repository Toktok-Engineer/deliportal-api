package model

type CompanyGroup struct {
	ID               uint    `json:"company_group_id"`
	CompanyGroupName string  `gorm:"type:varchar(150);not null;unique" json:"company_group_name" binding:"required"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type CreateCompanyGroupParameter struct {
	CompanyGroupName string  `gorm:"type:varchar(150);not null;unique" json:"company_group_name" binding:"required"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;" json:"deleted_at"`
}
