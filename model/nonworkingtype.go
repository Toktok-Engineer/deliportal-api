package model

type NonWorkingType struct {
	ID                 uint    `json:"non_working_type_id"`
	NonWorkingTypeName string  `gorm:"type:varchar(150);not null;unique" json:"non_working_type_name"`
	DeductLeave        bool    `gorm:"not null;typedefault:null" json:"deduct_leave"`
	Remark             string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID      uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID      uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID      uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt          float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt          float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt          float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateNonWorkingTypeParameter struct {
	NonWorkingTypeName string  `gorm:"type:varchar(150);not null;unique" json:"non_working_type_name"`
	DeductLeave        bool    `gorm:"not null;typedefault:null" json:"deduct_leave"`
	Remark             string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID      uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID      uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID      uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt          float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt          float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt          float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
