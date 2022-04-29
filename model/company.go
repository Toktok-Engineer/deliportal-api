package model

type Company struct {
	ID                  uint    `json:"company_id"`
	CompanyName         string  `gorm:"type:varchar(50);not null;unique" json:"company_name"`
	BusinessUnitID      uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"business_unit_id"`
	Address             string  `gorm:"type:varchar(500);not null" json:"address"`
	LegalLicenseFileUrl string  `gorm:"type:varchar(200);not null" json:"legal_license_file_url"`
	Status              uint    `gorm:"type:bigint;not null" json:"status"`
	ApprovedUserID      uint    `gorm:"type:bigint" json:"approved_user_id"`
	DeactivedUserID     uint    `gorm:"type:bigint" json:"deactived_user_id"`
	ApprovedDate        float64 `gorm:"type:double precision " json:"approved_date"`
	DeactivedDate       float64 `gorm:"type:double precision " json:"deactived_date"`
	Remark              string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID       uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID       uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID       uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt           float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt           float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt           float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectCompanyParameter struct {
	ID                  uint    `json:"company_id"`
	CompanyName         string  `gorm:"type:varchar(50);not null;unique" json:"company_name"`
	BusinessUnitID      uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"business_unit_id"`
	BusinessUnitName    string  `gorm:"type:varchar(100);not null;unique" json:"business_unit_name"`
	Address             string  `gorm:"type:varchar(500);not null" json:"address"`
	LegalLicenseFileUrl string  `gorm:"type:varchar(200);not null" json:"legal_license_file_url"`
	Status              uint    `gorm:"type:bigint;not null" json:"status"`
	ApprovedUserID      uint    `gorm:"type:bigint" json:"approved_user_id"`
	DeactivedUserID     uint    `gorm:"type:bigint" json:"deactived_user_id"`
	ApprovedDate        float64 `gorm:"type:double precision " json:"approved_date"`
	DeactivedDate       float64 `gorm:"type:double precision " json:"deactived_date"`
	Remark              string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID       uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID       uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID       uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt           float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt           float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt           float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateCompanyParameter struct {
	CompanyName         string  `gorm:"type:varchar(50);not null;unique" json:"company_name"`
	BusinessUnitID      uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"business_unit_id"`
	Address             string  `gorm:"type:varchar(500);not null" json:"address"`
	LegalLicenseFileUrl string  `gorm:"type:varchar(200);not null" json:"legal_license_file_url"`
	Status              uint    `gorm:"type:bigint;not null" json:"status"`
	ApprovedUserID      uint    `gorm:"type:bigint" json:"approved_user_id"`
	DeactivedUserID     uint    `gorm:"type:bigint" json:"deactived_user_id"`
	ApprovedDate        float64 `gorm:"type:double precision " json:"approved_date"`
	DeactivedDate       float64 `gorm:"type:double precision " json:"deactived_date"`
	Remark              string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID       uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID       uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID       uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt           float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt           float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt           float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
