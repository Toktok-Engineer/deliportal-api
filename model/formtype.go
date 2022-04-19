package model

type FormType struct {
	ID                  uint    `json:"form_type_id"`
	FormTypeCode        string  `gorm:"type:varchar(50);not null;unique" json:"form_type_code"`
	FormTypeDescription string  `gorm:"type:varchar(100);not null" json:"form_type_description"`
	Remark              string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID       uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID       uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID       uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt           float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt           float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt           float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateFormTypeParameter struct {
	FormTypeCode        string  `gorm:"type:varchar(50);not null;unique" json:"form_type_code"`
	FormTypeDescription string  `gorm:"type:varchar(100);not null" json:"form_type_description"`
	Remark              string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID       uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID       uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID       uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt           float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt           float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt           float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
