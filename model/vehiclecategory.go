package model

type VehicleCategory struct {
	ID                  uint    `json:"vehicle_category_id"`
	VehicleCategoryName string  `gorm:"type:varchar(50);not null;unique" json:"vehicle_category_name" binding:"required"`
	Remark              string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID       uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID       uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID       uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt           float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt           float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt           float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type CreateVehicleCategoryParameter struct {
	VehicleCategoryName string  `gorm:"type:varchar(50);not null;unique" json:"vehicle_category_name" binding:"required"`
	Remark              string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID       uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID       uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID       uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt           float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt           float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt           float64 `gorm:"type:double precision;" json:"deleted_at"`
}
