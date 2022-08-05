package model

type EmailQueue struct {
	ID               uint    `json:"email_queue_id"`
	EmailQueueTypeID uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"email_queue_type_id"`
	EmailRecipient   string  `gorm:"type:varchar(200)" json:"email_recipient"`
	EmailCC          string  `gorm:"type:varchar(200)" json:"email_cc"`
	EmailSubject     string  `gorm:"type:varchar(200)" json:"email_subject"`
	EmailBody        string  `gorm:"type:text" json:"email_body"`
	Status           int     `json:"status"`
	ErrorMessage     string  `gorm:"type:varchar(200)" json:"error_message"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectEmailQueueParameter struct {
	ID                 uint   `json:"email_queue_id"`
	EmailQueueTypeID   uint   `gorm:"type:bigint;foreign_key;not null;index:" json:"email_queue_type_id"`
	EmailQueueTypeName string `gorm:"type:varchar(50);not null" json:"email_queue_type_name"`
	EmailRecipient     string `gorm:"type:varchar(200)" json:"email_recipient"`
	EmailCC            string `gorm:"type:varchar(200)" json:"email_cc"`
	EmailSubject       string `gorm:"type:varchar(200)" json:"email_subject"`
	EmailBody          string `gorm:"type:text" json:"email_body"`
	Status             int    `json:"status"`
	ErrorMessage       string `gorm:"type:varchar(200)" json:"error_message"`
	Remark             string `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID      uint   `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID      uint   `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID      uint   `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt          string `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt          string `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt          string `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateEmailQueueParameter struct {
	EmailQueueTypeID uint    `gorm:"type:bigint;foreign_key;not null;index:" json:"email_queue_type_id"`
	EmailRecipient   string  `gorm:"type:varchar(200)" json:"email_recipient"`
	EmailCC          string  `gorm:"type:varchar(200)" json:"email_cc"`
	EmailSubject     string  `gorm:"type:varchar(200)" json:"email_subject"`
	EmailBody        string  `gorm:"type:text" json:"email_body"`
	Status           int     `json:"status"`
	ErrorMessage     string  `gorm:"type:varchar(200)" json:"error_message"`
	Remark           string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
