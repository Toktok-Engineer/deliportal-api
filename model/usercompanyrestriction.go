package model

type UserCompanyRestriction struct {
	ID            uint    `json:"user_company_restriction_id"`
	UserID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"user_id"`
	CompanyID     uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectUserCompanyRestrictionParameter struct {
	ID            uint    `json:"user_company_restriction_id"`
	UserID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"user_id"`
	Username      string  `gorm:"type:varchar(50);not null;unique" json:"username" binding:"required"`
	CompanyID     uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	CompanyName   string  `gorm:"type:varchar(50);not null;unique" json:"company_name"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateUserCompanyRestrictionParameter struct {
	UserID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"user_id"`
	CompanyID     uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
