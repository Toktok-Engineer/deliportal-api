package model

type NonWorkingDay struct {
	ID               uint    `json:"non_working_day_id"`
	NonWorkingTypeID uint    `json:"non_working_type_id"`
	PeriodYear       uint    `gorm:"type:bigint;not null" json:"period_year"`
	Description      string  `gorm:"type:varchar(150)" json:"description"`
	EffectiveFrom    float64 `gorm:"type:double precision;not null" json:"effective_from"`
	EffectiveTo      float64 `gorm:"type:double precision;not null" json:"effective_to"`
	Total            uint    `gorm:"type:bigint;not null" json:"total"`
	Status           uint    `gorm:"type:bigint;not null" json:"status"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectNonWorkingDayParameter struct {
	ID                 uint    `json:"non_working_day_id"`
	NonWorkingTypeID   uint    `json:"non_working_type_id"`
	NonWorkingTypeName string  `gorm:"type:varchar(150);not null;unique" json:"non_working_type_name"`
	PeriodYear         uint    `gorm:"type:bigint;not null" json:"period_year"`
	Description        string  `gorm:"type:varchar(150)" json:"description"`
	EffectiveFrom      string  `gorm:"type:varchar(100);not null" json:"effective_from"`
	EffectiveFromUnix  float64 `gorm:"type:double precision;not null" json:"effective_from_unix"`
	EffectiveTo        string  `gorm:"type:varchar(100);not null" json:"effective_to"`
	Total              uint    `gorm:"type:bigint;not null" json:"total"`
	Status             uint    `gorm:"type:bigint;not null" json:"status"`
	StatusName         string  `gorm:"type:varchar(50)" json:"status_name"`
	Remark             string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID      uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID      uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID      uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt          string  `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt          string  `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt          string  `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateNonWorkingDayParameter struct {
	NonWorkingTypeID uint    `json:"non_working_type_id"`
	PeriodYear       uint    `gorm:"type:bigint;not null" json:"period_year"`
	Description      string  `gorm:"type:varchar(150)" json:"description"`
	EffectiveFrom    float64 `gorm:"type:double precision;not null" json:"effective_from"`
	EffectiveTo      float64 `gorm:"type:double precision;not null" json:"effective_to"`
	Total            uint    `gorm:"type:bigint;not null" json:"total"`
	Status           uint    `gorm:"type:bigint;not null" json:"status"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
