package model

type CompanyLicense struct {
	ID                    uint    `json:"company_license_id"`
	ParentLicenseID       uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"parent_license_id"`
	LicenseNo             string  `gorm:"type:varchar(150);not null" json:"license_no"`
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
	GroupLicenseTypeID    uint    `gorm:"type:bigint;typedefault:null;index:" json:"group_license_type_id"`
	FileName              string  `gorm:"type:varchar(200)" json:"file_name"`
	FileUrl               string  `gorm:"type:varchar(200)" json:"file_url"`
}

type SelectCompanyLicenseParameter struct {
	ID                    uint   `json:"company_license_id"`
	CompanyName           string `gorm:"type:varchar(50);not null" json:"company_name"`
	BusinessUnitName      string `gorm:"type:varchar(100);not null" json:"business_unit_name"`
	ParentLicenseID       uint   `gorm:"type:bigint;foreign_key;not null;index:" json:"parent_license_id"`
	LicenseNo             string `gorm:"type:varchar(150);not null" json:"license_no"`
	LicenseParentNo       string `gorm:"type:varchar(50)" json:"license_parent_no"`
	LicenseTypeID         uint   `gorm:"type:bigint;foreign_key;not null;index:" json:"license_type_id"`
	LicenseTypeName       string `gorm:"type:varchar(50);not null;unique" json:"license_type_name"`
	CompanyID             uint   `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	Renewable             bool   `gorm:"not null;typedefault:null" json:"renewable"`
	ReminderCounter       uint   `gorm:"type:bigint;not null" json:"reminder_counter"`
	IssuedBy              string `gorm:"type:varchar(150);not null" json:"issued_by"`
	IssuedDate            string `gorm:"type:varchar(100);" json:"issued_date"`
	ExpiredDate           string `gorm:"type:varchar(100);" json:"expired_date"`
	EarliestRenewalDate   string `gorm:"type:varchar(100);" json:"earliest_renewal_date"`
	LastRenewalDate       string `gorm:"type:varchar(100);" json:"last_renewal_date"`
	Status                int    `json:"status"`
	StatusName            string `gorm:"type:varchar(50);" json:"status_name"`
	Severity              string `gorm:"type:varchar(50);" json:"severity"`
	RenewalStatus         int    `json:"renewal_status"`
	RenewalStatusName     string `gorm:"type:varchar(50);" json:"renewal_status_name"`
	ApprovedUserID        uint   `gorm:"type:bigint" json:"approved_user_id"`
	ApprovedUser          string `gorm:"type:varchar(100);" json:"approved_user"`
	RenewalApprovedUserID uint   `gorm:"type:bigint" json:"renewal_approved_user_id"`
	RenewalApprovedUser   string `gorm:"type:varchar(100);" json:"renewal_approved_user"`
	ApprovedDate          string `gorm:"type:varchar(100);not null"  json:"approved_date"`
	RenewalApprovedDate   string `gorm:"type:varchar(100);not null"  json:"renewal_approved_date"`
	Remark                string `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID         uint   `gorm:"type:bigint;not null" json:"created_user_id"`
	CreatedUser           string `gorm:"type:varchar(100);" json:"created_user"`
	UpdatedUserID         uint   `gorm:"type:bigint;not null" json:"updated_user_id"`
	UpdatedUser           string `gorm:"type:varchar(100);" json:"updated_user"`
	DeletedUserID         uint   `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	DeletedUser           string `gorm:"type:varchar(100);" json:"deleted_user"`
	CreatedAt             string `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt             string `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt             string `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
	GroupLicenseTypeID    uint   `gorm:"type:bigint;typedefault:null;index:" json:"group_license_type_id"`
	GroupLicenseTypeName  string `gorm:"type:varchar(150);not null;unique" json:"group_license_type_name" binding:"required"`
	FileName              string `gorm:"type:varchar(200)" json:"file_name"`
	FileUrl               string `gorm:"type:varchar(200)" json:"file_url"`
}

type CreateCompanyLicenseParameter struct {
	ParentLicenseID       uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"parent_license_id"`
	LicenseNo             string  `gorm:"type:varchar(150);not null" json:"license_no"`
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
	GroupLicenseTypeID    uint    `gorm:"type:bigint;typedefault:null;index:" json:"group_license_type_id"`
	FileName              string  `gorm:"type:varchar(200)" json:"file_name"`
	FileUrl               string  `gorm:"type:varchar(200)" json:"file_url"`
}
