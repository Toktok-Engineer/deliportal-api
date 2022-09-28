package model

type InternalMemoAttachment struct {
	ID             uint    `json:"internal_memo_attachment_id"`
	InternalMemoID uint    `json:"internal_memo_id"`
	Description    string  `gorm:"type:varchar(200)" json:"description"`
	FileName       string  `gorm:"type:varchar(200)" json:"file_name"`
	FileUrl        string  `gorm:"type:varchar(200)" json:"file_url"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type SelectInternalMemoAttachmentParameter struct {
	ID             uint   `json:"internal_memo_attachment_id"`
	InternalMemoID uint   `json:"internal_memo_id"`
	Description    string `gorm:"type:varchar(200)" json:"description"`
	FileName       string `gorm:"type:varchar(200)" json:"file_name"`
	FileUrl        string `gorm:"type:varchar(200)" json:"file_url"`
	Remark         string `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint   `gorm:"type:bigint;not null" json:"created_user_id"`
	CreatedUser    string `gorm:"type:varchar(100);" json:"created_user"`
	UpdatedUserID  uint   `gorm:"type:bigint;not null" json:"updated_user_id"`
	UpdatedUser    string `gorm:"type:varchar(100);" json:"updated_user"`
	DeletedUserID  uint   `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	DeletedUser    string `gorm:"type:varchar(100);" json:"deleted_user"`
	CreatedAt      string `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt      string `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt      string `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateInternalMemoAttachmentParameter struct {
	InternalMemoID uint    `json:"internal_memo_id"`
	Description    string  `gorm:"type:varchar(200)" json:"description"`
	FileName       string  `gorm:"type:varchar(200)" json:"file_name"`
	FileUrl        string  `gorm:"type:varchar(200)" json:"file_url"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;" json:"deleted_at"`
}
