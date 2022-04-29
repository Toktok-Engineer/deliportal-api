package model

type EmailQueueType struct {
	ID                 uint    `json:"email_queue_type_id"`
	EmailQueueTypeName string  `gorm:"type:varchar(50);not null" json:"email_queue_type_name"`
	TableReference     string  `gorm:"type:varchar(50)" json:"table_reference"`
	FieldReference     string  `gorm:"type:varchar(50)" json:"field_reference"`
	RecipientName      string  `gorm:"type:varchar(200)" json:"recipient_name"`
	EmailRecipient     string  `gorm:"type:varchar(200)" json:"email_recipient"`
	EmailCC            string  `gorm:"type:varchar(200)" json:"email_cc"`
	EmailSubject       string  `gorm:"type:varchar(200)" json:"email_subject"`
	EmailBody          string  `gorm:"type:text" json:"email_body"`
	ParamValue         string  `gorm:"type:varchar(200)" json:"param_value"`
	Remark             string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID      uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID      uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID      uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt          float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt          float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt          float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateEmailQueueTypeParameter struct {
	EmailQueueTypeName string  `gorm:"type:varchar(50);not null" json:"email_queue_type_name"`
	TableReference     string  `gorm:"type:varchar(50)" json:"table_reference"`
	FieldReference     string  `gorm:"type:varchar(50)" json:"field_reference"`
	RecipientName      string  `gorm:"type:varchar(200)" json:"recipient_name"`
	EmailRecipient     string  `gorm:"type:varchar(200)" json:"email_recipient"`
	EmailCC            string  `gorm:"type:varchar(200)" json:"email_cc"`
	EmailSubject       string  `gorm:"type:varchar(200)" json:"email_subject"`
	EmailBody          string  `gorm:"type:text" json:"email_body"`
	ParamValue         string  `gorm:"type:varchar(200)" json:"param_value"`
	Remark             string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID      uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID      uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID      uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt          float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt          float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt          float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
