package model

type CompanyManagementHistoryDetail struct {
	ID                         uint    `json:"company_management_history_detail_id"`
	CompanyManagementHistoryID uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_management_history_id"`
	CompanyManagementTypeID    uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_management_type_id"`
	ManagementName             string  `gorm:"type:varchar(100);not null" json:"management_name"`
	Remark                     string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID              uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID              uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID              uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                  float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                  float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                  float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectCompanyManagementHistoryDetail struct {
	ID                         uint   `json:"company_management_history_detail_id"`
	CompanyManagementHistoryID uint   `gorm:"type:bigint;foreign_key;not null;index:" json:"company_management_history_id"`
	CompanyManagementTypeID    uint   `gorm:"type:bigint;foreign_key;not null;index:" json:"company_management_type_id"`
	CompanyManagementTypeName  string `gorm:"type:varchar(50);not null;unique" json:"company_management_type_name"`
	ManagementName             string `gorm:"type:varchar(100);not null" json:"management_name"`
	Remark                     string `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID              uint   `gorm:"type:bigint;not null" json:"created_user_id"`
	CreatedUser                string `gorm:"type:varchar(100);" json:"created_user"`
	UpdatedUserID              uint   `gorm:"type:bigint;not null" json:"updated_user_id"`
	UpdatedUser                string `gorm:"type:varchar(100);" json:"updated_user"`
	DeletedUserID              uint   `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	DeletedUser                string `gorm:"type:varchar(100);" json:"deleted_user"`
	CreatedAt                  string `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt                  string `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt                  string `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateCompanyManagementHistoryDetailParameter struct {
	CompanyManagementHistoryID uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_management_history_id"`
	CompanyManagementTypeID    uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"company_management_type_id"`
	ManagementName             string  `gorm:"type:varchar(100);not null" json:"management_name"`
	Remark                     string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID              uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID              uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID              uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                  float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                  float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                  float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
