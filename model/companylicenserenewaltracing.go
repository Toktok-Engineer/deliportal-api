package model

type CompanyLicenseRenewalTracing struct {
	ID               uint    `json:"company_license_renewal_tracing_id"`
	CompanyLicenseID uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_license_id"`
	RenewalStatus    uint    `gorm:"type:bigint" json:"renewal_status"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateCompanyLicenseRenewalTracingParameter struct {
	CompanyLicenseID uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_license_id"`
	RenewalStatus    uint    `gorm:"type:bigint" json:"renewal_status"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
