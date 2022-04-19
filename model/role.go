package model

type Role struct {
	ID              uint    `json:"role_id"`
	RoleCode        string  `gorm:"type:varchar(50);not null;unique" json:"role_code"`
	RoleDescription string  `gorm:"type:varchar(100);not null" json:"role_description"`
	Remark          string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID   uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID   uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID   uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt       float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt       float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt       float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateRoleParameter struct {
	RoleCode        string  `gorm:"type:varchar(50);not null;unique" json:"role_code"`
	RoleDescription string  `gorm:"type:varchar(100);not null" json:"role_description"`
	Remark          string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID   uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID   uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID   uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt       float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt       float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt       float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
