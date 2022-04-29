package model

type EmailQueueReference struct {
	ID               uint    `json:"email_queue_reference_id"`
	EmailQueueID     uint    `gorm:"type:bigint;foreign_key;not null" json:"email_queue_id"`
	EmailQueueTypeID uint    `gorm:"type:bigint;foreign_key;not null" json:"email_queue_type_id"`
	ReferenceID      uint    `gorm:"type:bigint;not nullt" json:"reference_id"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectEmailQueueReferenceParameter struct {
	ID                 uint    `json:"email_queue_reference_id"`
	EmailQueueID       uint    `gorm:"type:bigint;foreign_key;not null" json:"email_queue_id"`
	EmailQueueTypeID   uint    `gorm:"type:bigint;foreign_key;not null" json:"email_queue_type_id"`
	EmailQueueTypeName string  `gorm:"type:varchar(50);not null" json:"email_queue_type_name"`
	ReferenceID        uint    `gorm:"type:bigint;not nullt" json:"reference_id"`
	Remark             string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID      uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID      uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID      uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt          float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt          float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt          float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateEmailQueueReferenceParameter struct {
	EmailQueueID     uint    `gorm:"type:bigint;foreign_key;not null" json:"email_queue_id"`
	EmailQueueTypeID uint    `gorm:"type:bigint;foreign_key;not null" json:"email_queue_type_id"`
	ReferenceID      uint    `gorm:"type:bigint;not nullt" json:"reference_id"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
