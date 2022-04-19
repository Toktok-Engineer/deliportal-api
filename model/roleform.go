package model

type RoleForm struct {
	ID            uint    `json:"role_form_id"`
	RoleID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"role_id"`
	FormID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"form_id"`
	CreateFlag    bool    `gorm:"not null;typedefault:null" json:"create_flag"`
	ReadFlag      bool    `gorm:"not null;typedefault:null" json:"read_flag"`
	UpdateFlag    bool    `gorm:"not null;typedefault:null" json:"update_flag"`
	DeleteFlag    bool    `gorm:"not null;typedefault:null" json:"delete_flag"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectRoleFormParameter struct {
	ID              uint    `json:"role_form_id"`
	RoleID          uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"role_id"`
	RoleCode        string  `gorm:"type:varchar(50);not null;unique" json:"role_code"`
	RoleDescription string  `gorm:"type:varchar(100);not null" json:"role_description"`
	FormID          uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"form_id"`
	FormCode        string  `gorm:"type:varchar(50);not null;unique" json:"form_code"`
	FormDescription string  `gorm:"type:varchar(100);not null" json:"form_description"`
	CreateFlag      bool    `gorm:"not null;typedefault:null" json:"create_flag"`
	ReadFlag        bool    `gorm:"not null;typedefault:null" json:"read_flag"`
	UpdateFlag      bool    `gorm:"not null;typedefault:null" json:"update_flag"`
	DeleteFlag      bool    `gorm:"not null;typedefault:null" json:"delete_flag"`
	Remark          string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID   uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID   uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID   uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt       float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt       float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt       float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateRoleFormParameter struct {
	RoleID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"role_id"`
	FormID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"form_id"`
	CreateFlag    bool    `gorm:"not null;typedefault:null" json:"create_flag"`
	ReadFlag      bool    `gorm:"not null;typedefault:null" json:"read_flag"`
	UpdateFlag    bool    `gorm:"not null;typedefault:null" json:"update_flag"`
	DeleteFlag    bool    `gorm:"not null;typedefault:null" json:"delete_flag"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
