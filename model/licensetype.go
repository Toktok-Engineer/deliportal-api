package model

type LicenseType struct {
	ID                             uint    `json:"license_type_id"`
	LicenseTypeName                string  `gorm:"type:varchar(50);not null;unique" json:"license_type_name"`
	ReminderBeforeMonth            int     `json:"reminder_before_month"`
	ManagementReminderBeforeMonth  int     `json:"management_reminder_before_month"`
	ReminderFrequencyDay           int     `json:"reminder_frequency_day"`
	ManagementReminderFrequencyDay int     `json:"management_reminder_frequency_day"`
	Remark                         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID                  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID                  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID                  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateLicenseTypeParameter struct {
	LicenseTypeName                string  `gorm:"type:varchar(50);not null;unique" json:"license_type_name"`
	ReminderBeforeMonth            int     `json:"reminder_before_month"`
	ManagementReminderBeforeMonth  int     `json:"management_reminder_before_month"`
	ReminderFrequencyDay           int     `json:"reminder_frequency_day"`
	ManagementReminderFrequencyDay int     `json:"management_reminder_frequency_day"`
	Remark                         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID                  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID                  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID                  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt                      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt                      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt                      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
