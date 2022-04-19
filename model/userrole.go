package model

type UserRole struct {
	ID            uint    `json:"user_role_id"`
	UserID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"user_id"`
	RoleID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"role_id"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectUserRoleParameter struct {
	ID              uint    `json:"user_role_id"`
	UserID          uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"user_id"`
	Username        string  `gorm:"type:varchar(50);not null;unique" json:"username"`
	Password        string  `gorm:"type:varchar(100);typedefault:null" json:"password"`
	EmployeeID      uint    `gorm:"type:bigint;typedefault:null;index:" json:"employee_id"`
	Email           string  `gorm:"type:varchar(30);not null" json:"email"`
	RoleID          uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"role_id"`
	RoleCode        string  `gorm:"type:varchar(50);not null;unique" json:"role_code"`
	RoleDescription string  `gorm:"type:varchar(100);not null" json:"role_description"`
	Remark          string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserId   uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserId   uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserId   uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt       float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt       float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt       float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateUserRoleParameter struct {
	UserID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"user_id"`
	RoleID        uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"role_id"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
