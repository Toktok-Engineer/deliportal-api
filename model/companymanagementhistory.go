package model

type CompanyManagementHistory struct {
	ID            uint    `json:"company_management_history_id"`
	CompanyID     uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	CompanyAkteId uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_akte_id"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectCompanyManagementHistoryParameter struct {
	ID                uint   `json:"company_management_history_id"`
	CompanyID         uint   `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	CompanyName       string `gorm:"type:varchar(50);not null;unique" json:"company_name"`
	CompanyAkteId     uint   `gorm:"type:bigint;foreign_key;not null;index:" json:"company_akte_id"`
	Year              string `gorm:"type:varchar(100);typedefault:null" json:"year"`
	AkteDate          string `gorm:"type:varchar(100);typedefault:null" json:"akte_date"`
	AkteNo            string `gorm:"type:varchar(150)" json:"akte_no"`
	ChangeInformation string `gorm:"type:varchar(200)" json:"change_information"`
	Remark            string `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID     uint   `gorm:"type:bigint;not null" json:"created_user_id"`
	CreatedUser       string `gorm:"type:varchar(100);" json:"created_user"`
	UpdatedUserID     uint   `gorm:"type:bigint;not null" json:"updated_user_id"`
	UpdatedUser       string `gorm:"type:varchar(100);" json:"updated_user"`
	DeletedUserID     uint   `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	DeletedUser       string `gorm:"type:varchar(100);" json:"deleted_user"`
	CreatedAt         string `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt         string `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt         string `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateCompanyManagementHistoryParameter struct {
	CompanyID     uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_id"`
	CompanyAkteId uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_akte_id"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
