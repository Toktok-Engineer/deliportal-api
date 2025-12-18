package model

type Asuransi struct {
	ID            uint    `json:"asuransi_id"`
	AsuransiName  string  `gorm:"type:varchar(100);not null;" json:"asuransi_name" binding:"required"`
	NomorHP       string  `gorm:"type:varchar(50);not null;" json:"nomor_hp"`
	Email         string  `gorm:"type:varchar(100);not null;" json:"email"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type CreateAsuransiParameter struct {
	AsuransiName  string  `gorm:"type:varchar(100);not null;" json:"asuransi_name" binding:"required"`
	NomorHP       string  `gorm:"type:varchar(50);not null;" json:"nomor_hp"`
	Email         string  `gorm:"type:varchar(100);not null;" json:"email"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;" json:"deleted_at"`
}
