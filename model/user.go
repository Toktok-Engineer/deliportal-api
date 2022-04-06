package model

import "database/sql"

type User struct {
	ID            uint          `json:"user_id"`
	Username      string        `gorm:"type:varchar(50);not null;unique" json:"username"`
	Password      string        `gorm:"type:varchar(100);typedefault:null" json:"password"`
	EmployeeID    sql.NullInt64 `gorm:"type:bigint;typedefault:null;foreign_key;index:" json:"employee_id"`
	Email         string        `gorm:"type:varchar(30);not null" json:"email"`
	Remark        string        `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserId uint          `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserId uint          `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserId uint          `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64       `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64       `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64       `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateUserParameter struct {
	Username      string        `gorm:"type:varchar(50);not null;unique" json:"username"`
	Password      string        `gorm:"type:varchar(100);typedefault:null" json:"password"`
	EmployeeID    sql.NullInt64 `gorm:"type:bigint;typedefault:null;foreign_key;index:" json:"employee_id"`
	Email         string        `gorm:"type:varchar(30);not null" json:"email"`
	Remark        string        `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserId uint          `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserId uint          `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserId uint          `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64       `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64       `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64       `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
