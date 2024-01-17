package model

type CompanyShareholder struct {
	ID                uint    `json:"company_shareholder_id"`
	CompanyID         uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	ShareholderName   string  `gorm:"type:varchar(100);not null" json:"shareholder_name"`
	NumberOfShare     float64 `gorm:"type:double precision;not null" json:"number_of_share"`
	PercentageOfShare float64 `gorm:"type:double precision;not null" json:"percentage_of_share"`
	ShareAmount       float64 `gorm:"type:double precision;not null" json:"share_amount"`
	Remark            string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID     uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID     uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt         float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt         float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt         float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectCompanyShareholder struct {
	ID                uint    `json:"company_shareholder_id"`
	CompanyID         uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	ShareholderName   string  `gorm:"type:varchar(100);not null" json:"shareholder_name"`
	NumberOfShare     float64 `gorm:"type:double precision;not null" json:"number_of_share"`
	PercentageOfShare float64 `gorm:"type:double precision;not null" json:"percentage_of_share"`
	ShareAmount       float64 `gorm:"type:double precision;not null" json:"share_amount"`
	Remark            string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	CreatedUser       string  `gorm:"type:varchar(100);" json:"created_user"`
	UpdatedUserID     uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	UpdatedUser       string  `gorm:"type:varchar(100);" json:"updated_user"`
	DeletedUserID     uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	DeletedUser       string  `gorm:"type:varchar(100);" json:"deleted_user"`
	CreatedAt         string  `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt         string  `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt         string  `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateCompanyShareholderParameter struct {
	CompanyID         uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	ShareholderName   string  `gorm:"type:varchar(100);not null" json:"shareholder_name"`
	NumberOfShare     float64 `gorm:"type:double precision;not null" json:"number_of_share"`
	PercentageOfShare float64 `gorm:"type:double precision;not null" json:"percentage_of_share"`
	ShareAmount       float64 `gorm:"type:double precision;not null" json:"share_amount"`
	Remark            string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID     uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID     uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt         float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt         float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt         float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectCompanyShareholderReport struct {
	ID                int     `json:"company_shareholder_id"`
	CompanyName       string  `gorm:"type:varchar(100);" json:"company_name"`
	ShareholderName   string  `gorm:"type:varchar(100);not null" json:"shareholder_name"`
	NumberOfShare     float64 `gorm:"type:double precision;not null" json:"number_of_share"`
	PercentageOfShare float64 `gorm:"type:double precision;not null" json:"percentage_of_share"`
	ShareAmount       float64 `gorm:"type:double precision;not null" json:"share_amount"`
}
