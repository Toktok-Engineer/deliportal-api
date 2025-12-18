package model

type STNKReminder struct {
	ID            uint    `json:"stnk_reminder_id"`
	VehicleID     uint    `json:"vehicle_id"`
	STNKEndDate   float64 `gorm:"type:double precision;" json:"stnk_end_date"`
	Status        bool    `gorm:"type:bool;" json:"status"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type SelectSTNKReminderParameter struct {
	ID            uint   `json:"stnk_reminder_id"`
	VehicleID     uint   `json:"vehicle_id"`
	STNKEndDate   string `gorm:"type:varchar(50);" json:"stnk_end_date"`
	Status        bool   `gorm:"type:bool;" json:"status"`
	CreatedUserID uint   `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint   `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint   `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     string `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt     string `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt     string `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateSTNKReminderParameter struct {
	VehicleID     uint    `json:"vehicle_id"`
	STNKEndDate   float64 `gorm:"type:double precision;" json:"stnk_end_date"`
	Status        bool    `gorm:"type:bool;" json:"status"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;" json:"deleted_at"`
}
