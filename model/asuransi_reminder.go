package model

type AsuransiReminder struct {
	ID              uint    `json:"asuransi_reminder_id"`
	VehicleID       uint    `json:"vehicle_id"`
	AsuransiEndDate float64 `gorm:"type:double precision;" json:"asuransi_end_date"`
	AsuransiID      uint    `json:"asuransi_id"`
	NomorPolis      string  `gorm:"type:varchar(100);" json:"nomor_polis"`
	Status          bool    `gorm:"type:bool;" json:"status"`
	BuktiBayar      string  `gorm:"type:varchar(500);not null" json:"bukti_bayar"`
	Price           float64 `gorm:"type:double precision" json:"price"`
	CreatedUserID   uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID   uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID   uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt       float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt       float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt       float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type SelectAsuransiReminderParameter struct {
	ID              uint    `json:"asuransi_reminder_id"`
	VehicleID       uint    `json:"vehicle_id"`
	AsuransiEndDate string  `gorm:"type:varchar(50);" json:"asuransi_end_date"`
	AsuransiID      uint    `json:"asuransi_id"`
	AsuransiName    string  `gorm:"type:varchar(200);" json:"asuransi_name"`
	NomorPolis      string  `gorm:"type:varchar(100);" json:"nomor_polis"`
	Status          bool    `gorm:"type:bool;" json:"status"`
	BuktiBayar      string  `gorm:"type:varchar(500);not null" json:"bukti_bayar"`
	Price           float64 `gorm:"type:double precision" json:"price"`
	CreatedUserID   uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID   uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID   uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt       string  `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt       string  `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt       string  `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateAsuransiReminderParameter struct {
	VehicleID       uint    `json:"vehicle_id"`
	AsuransiEndDate float64 `gorm:"type:double precision;" json:"asuransi_end_date"`
	AsuransiID      uint    `json:"asuransi_id"`
	NomorPolis      string  `gorm:"type:varchar(100);" json:"nomor_polis"`
	Status          bool    `gorm:"type:bool;" json:"status"`
	BuktiBayar      string  `gorm:"type:varchar(500);not null" json:"bukti_bayar"`
	Price           float64 `gorm:"type:double precision" json:"price"`
	CreatedUserID   uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID   uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID   uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt       float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt       float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt       float64 `gorm:"type:double precision;" json:"deleted_at"`
}
