package model

type CompanyGroupCompany struct {
	ID             uint    `json:"company_group_company_id"`
	CompanyGroupID uint    `gorm:"type:bigint;foreign_key;index:" json:"company_group_id"`
	CompanyID      uint    `gorm:"type:bigint;foreign_key;index:" json:"company_id"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectCompanyGroupCompanyParameter struct {
	ID               uint    `json:"company_group_company_id"`
	CompanyGroupID   uint    `gorm:"type:bigint;foreign_key;index:" json:"company_group_id"`
	CompanyGroupName string  `gorm:"type:varchar(150);not null" json:"company_group_name"`
	CompanyID        uint    `gorm:"type:bigint;foreign_key;index:" json:"company_id"`
	CompanyName      string  `gorm:"type:varchar(50);not null" json:"company_name"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateCompanyGroupCompanyParameter struct {
	CompanyGroupID uint    `gorm:"type:bigint;foreign_key;index:" json:"company_group_id"`
	CompanyID      uint    `gorm:"type:bigint;foreign_key;index:" json:"company_id"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
