package model

type CompanyLicense struct {
	ID                    uint    `json:"company_license_id"`
	ParentLicenseID       uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"parent_license_id"`
	LicenseNo             string  `gorm:"type:varchar(50);not null" json:"license_no"`
	LicenseTypeID         uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"license_type_id"`
	CompanyID             uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	Renewable             bool    `gorm:"not null;typedefault:null" json:"renewable"`
	ReminderCounter       uint    `gorm:"type:bigint;not null" json:"reminder_counter"`
	IssuedBy              string  `gorm:"type:varchar(50);not null" json:"issued_by"`
	IssuedDate            float64 `gorm:"type:double precision;not null" json:"issued_date"`
	ExpiredDate           float64 `gorm:"type:double precision;not null" json:"expired_date"`
	EarliestRenewalDate   float64 `gorm:"type:double precision;not null" json:"earliest_renewal_date"`
	LastRenewalDate       float64 `gorm:"type:double precision;not null" json:"last_renewal_date"`
	Status                int     `json:"status"`
	RenewalStatus         int     `json:"renewal_status"`
	ApprovedUserID        uint    `gorm:"type:bigint" json:"approved_user_id"`
	RenewalApprovedUserID uint    `gorm:"type:bigint" json:"renewal_approved_user_id"`
	ApprovedDate          float64 `gorm:"type:double precision " json:"approved_date"`
	RenewalApprovedDate   float64 `gorm:"type:double precision " json:"renewal_approved_date"`
	Remark                string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID         uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID         uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID         uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt             float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt             float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt             float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectCompanyLicenseParameter struct {
	ID                    uint    `json:"company_license_id"`
	ParentLicenseID       uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"parent_license_id"`
	LicenseNo             string  `gorm:"type:varchar(50);not null" json:"license_no"`
	LicenseTypeID         uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"license_type_id"`
	LicenseTypeName       string  `gorm:"type:varchar(50);not null;unique" json:"license_type_name"`
	CompanyID             uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	Renewable             bool    `gorm:"not null;typedefault:null" json:"renewable"`
	ReminderCounter       uint    `gorm:"type:bigint;not null" json:"reminder_counter"`
	IssuedBy              string  `gorm:"type:varchar(150);not null" json:"issued_by"`
	IssuedDate            float64 `gorm:"type:double precision;not null" json:"issued_date"`
	ExpiredDate           float64 `gorm:"type:double precision;not null" json:"expired_date"`
	EarliestRenewalDate   float64 `gorm:"type:double precision;not null" json:"earliest_renewal_date"`
	LastRenewalDate       float64 `gorm:"type:double precision;not null" json:"last_renewal_date"`
	Status                int     `json:"status"`
	RenewalStatus         int     `json:"renewal_status"`
	ApprovedUserID        uint    `gorm:"type:bigint" json:"approved_user_id"`
	RenewalApprovedUserID uint    `gorm:"type:bigint" json:"renewal_approved_user_id"`
	ApprovedDate          float64 `gorm:"type:double precision" json:"approved_date"`
	RenewalApprovedDate   float64 `gorm:"type:double precision" json:"renewal_approved_date"`
	Remark                string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID         uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID         uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID         uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt             float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt             float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt             float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateCompanyLicenseParameter struct {
	ParentLicenseID       uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"parent_license_id"`
	LicenseNo             string  `gorm:"type:varchar(50);not null" json:"license_no"`
	LicenseTypeID         uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"license_type_id"`
	CompanyID             uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	Renewable             bool    `gorm:"not null;typedefault:null" json:"renewable"`
	ReminderCounter       uint    `gorm:"type:bigint;not null" json:"reminder_counter"`
	IssuedBy              string  `gorm:"type:varchar(150);not null" json:"issued_by"`
	IssuedDate            float64 `gorm:"type:double precision;not null" json:"issued_date"`
	ExpiredDate           float64 `gorm:"type:double precision;not null" json:"expired_date"`
	EarliestRenewalDate   float64 `gorm:"type:double precision;not null" json:"earliest_renewal_date"`
	LastRenewalDate       float64 `gorm:"type:double precision;not null" json:"last_renewal_date"`
	Status                int     `json:"status"`
	RenewalStatus         int     `json:"renewal_status"`
	ApprovedUserID        uint    `gorm:"type:bigint" json:"approved_user_id"`
	RenewalApprovedUserID uint    `gorm:"type:bigint" json:"renewal_approved_user_id"`
	ApprovedDate          float64 `gorm:"type:double precision " json:"approved_date"`
	RenewalApprovedDate   float64 `gorm:"type:double precision " json:"renewal_approved_date"`
	Remark                string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID         uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID         uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID         uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt             float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt             float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt             float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectCompanyLicenseExpiredParameter struct {
	ID                  uint    `json:"company_license_id"`
	LicenseTypeName     string  `gorm:"type:varchar(50);not null;unique" json:"license_type_name"`
	LicenseNo           string  `gorm:"type:varchar(50);not null" json:"license_no"`
	CompanyName         string  `gorm:"type:varchar(50);not null;unique" json:"company_name"`
	ExpiredDate         float64 `gorm:"type:double precision;not null" json:"expired_date"`
	EarliestRenewalDate float64 `gorm:"type:double precision;not null" json:"earliest_renewal_date"`
	LastRenewalDate     float64 `gorm:"type:double precision;not null" json:"last_renewal_date"`
	RenewalStatus       int     `json:"renewal_status"`
	Status              int     `json:"status"`
	ReminderCounter     uint    `gorm:"type:bigint;not null" json:"reminder_counter"`
	Remark              string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID       uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID       uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID       uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt           float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt           float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt           float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
