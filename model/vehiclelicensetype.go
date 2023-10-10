package model

type VehicleLicenseType struct {
	ID                     uint    `json:"vehicle_license_type_id"`
	VehicleLicenseTypeName string  `gorm:"type:varchar(50);not null;unique" json:"vehicle_license_type_name" binding:"required"`
	ReminderFirstMonth     bool    `gorm:"not null;typedefault:null" json:"reminder_first_month"`
	ReminderMonth          uint    `gorm:"type:bigint;not null" json:"reminder_month"`
	Remark                 string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID          uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID          uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID          uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt              float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt              float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt              float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type CreateVehicleLicenseTypeParameter struct {
	VehicleLicenseTypeName string  `gorm:"type:varchar(50);not null;unique" json:"vehicle_license_type_name" binding:"required"`
	ReminderFirstMonth     bool    `gorm:"not null;typedefault:null" json:"reminder_first_month"`
	ReminderMonth          uint    `gorm:"type:bigint;not null" json:"reminder_month"`
	Remark                 string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID          uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID          uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID          uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt              float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt              float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt              float64 `gorm:"type:double precision;" json:"deleted_at"`
}
