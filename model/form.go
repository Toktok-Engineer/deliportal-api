package model

type Form struct {
	ID              uint    `json:"form_id"`
	FormCode        string  `gorm:"type:varchar(50);not null;unique" json:"form_code"`
	FormPHP         string  `gorm:"type:varchar(50);typedefault:null" json:"form_php"`
	FormDescription string  `gorm:"type:varchar(100);not null" json:"form_description"`
	FormTypeID      uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"form_type_id"`
	FormParentID    uint    `gorm:"type:bigint;foreign_key;index:" json:"form_parent_id"`
	SequenceNo      uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	ClassTag        string  `gorm:"type:varchar(50)" json:"class_tag"`
	Remark          string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID   uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID   uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID   uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt       float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt       float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt       float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectFormParameter struct {
	ID                  uint    `json:"form_id"`
	FormCode            string  `gorm:"type:varchar(50);not null;unique" json:"form_code"`
	FormPHP             string  `gorm:"type:varchar(50);typedefault:null" json:"form_php"`
	FormDescription     string  `gorm:"type:varchar(100);not null" json:"form_description"`
	FormTypeID          uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"form_type_id"`
	FormTypeCode        string  `gorm:"type:varchar(50);not null;unique" json:"form_type_code"`
	FormTypeDescription string  `gorm:"type:varchar(100);not null" json:"form_type_description"`
	FormParentID        uint    `gorm:"type:bigint;foreign_key;index:" json:"form_parent_id"`
	SequenceNo          uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	ClassTag            string  `gorm:"type:varchar(50)" json:"class_tag"`
	Remark              string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID       uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID       uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID       uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt           float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt           float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt           float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectFormCRUDParameter struct {
	ID                  uint    `json:"form_id"`
	FormCode            string  `gorm:"type:varchar(50);not null;unique" json:"form_code"`
	FormPHP             string  `gorm:"type:varchar(50);typedefault:null" json:"form_php"`
	FormDescription     string  `gorm:"type:varchar(100);not null" json:"form_description"`
	FormTypeID          uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"form_type_id"`
	FormTypeCode        string  `gorm:"type:varchar(50);not null;unique" json:"form_type_code"`
	FormTypeDescription string  `gorm:"type:varchar(100);not null" json:"form_type_description"`
	FormParentID        uint    `gorm:"type:bigint;foreign_key;index:" json:"form_parent_id"`
	SequenceNo          uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	ClassTag            string  `gorm:"type:varchar(50)" json:"class_tag"`
	CreateFlag          bool    `gorm:"not null;typedefault:null" json:"create_flag"`
	ReadFlag            bool    `gorm:"not null;typedefault:null" json:"read_flag"`
	UpdateFlag          bool    `gorm:"not null;typedefault:null" json:"update_flag"`
	DeleteFlag          bool    `gorm:"not null;typedefault:null" json:"delete_flag"`
	Remark              string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID       uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID       uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID       uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt           float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt           float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt           float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateFormParameter struct {
	FormCode        string  `gorm:"type:varchar(50);not null;unique" json:"form_code"`
	FormPHP         string  `gorm:"type:varchar(50);typedefault:null" json:"form_php"`
	FormDescription string  `gorm:"type:varchar(100);not null" json:"form_description"`
	FormTypeID      uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"form_type_id"`
	FormParentID    uint    `gorm:"type:bigint;foreign_key;index:" json:"form_parent_id"`
	SequenceNo      uint    `gorm:"type:bigint;not null" json:"sequence_no"`
	ClassTag        string  `gorm:"type:varchar(50)" json:"class_tag"`
	Remark          string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID   uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID   uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID   uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt       float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt       float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt       float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
