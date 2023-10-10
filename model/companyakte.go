package model

type CompanyAkte struct {
	ID            uint    `json:"company_akte_id"`
	CompanyID     uint    `gorm:"type:bigint;not null" json:"company_id"`
	AkteNo        string  `gorm:"type:varchar(150)" json:"akte_no"`
	AkteDate      uint    `gorm:"type:bigint;not null" json:"akte_date"`
	Year          uint    `gorm:"type:bigint;not null" json:"year"`
	SKNO1         string  `gorm:"type:varchar(150);not null" json:"sk_no_1"`
	SKNO2         string  `gorm:"type:varchar(150);not null" json:"sk_no_2"`
	SKNO3         string  `gorm:"type:varchar(150);not null" json:"sk_no_3"`
	FileName      string  `gorm:"type:varchar(200)" json:"file_name"`
	FileURL       string  `gorm:"type:varchar(200)" json:"file_url"`
	Remark        string  `gorm:"type:varchar(5000)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectCompanyAkteParameter struct {
	ID            uint   `json:"company_akte_id"`
	CompanyID     uint   `gorm:"type:bigint;not null" json:"company_id"`
	CompanyName   string `gorm:"type:varchar(50);not null;unique" json:"company_name"`
	AkteNo        string `gorm:"type:varchar(150)" json:"akte_no"`
	AkteDate      string `gorm:"type:varchar(100);typedefault:null" json:"akte_date"`
	Year          string `gorm:"type:varchar(100);typedefault:null" json:"year"`
	SKNO1         string `gorm:"type:varchar(150);not null" json:"sk_no_1"`
	SKNO2         string `gorm:"type:varchar(150);not null" json:"sk_no_2"`
	SKNO3         string `gorm:"type:varchar(150);not null" json:"sk_no_3"`
	FileName      string `gorm:"type:varchar(200)" json:"file_name"`
	FileURL       string `gorm:"type:varchar(200)" json:"file_url"`
	Remark        string `gorm:"type:varchar(5000)" json:"remark"`
	CreatedUserID uint   `gorm:"type:bigint;not null" json:"created_user_id"`
	CreatedUser   string `gorm:"type:varchar(100);" json:"created_user"`
	UpdatedUserID uint   `gorm:"type:bigint;not null" json:"updated_user_id"`
	UpdatedUser   string `gorm:"type:varchar(100);" json:"updated_user"`
	DeletedUserID uint   `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	DeletedUser   string `gorm:"type:varchar(100);" json:"deleted_user"`
	CreatedAt     string `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt     string `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt     string `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateCompanyAkteParameter struct {
	CompanyID     uint    `gorm:"type:bigint;not null" json:"company_id"`
	AkteNo        string  `gorm:"type:varchar(150)" json:"akte_no"`
	AkteDate      uint    `gorm:"type:bigint;not null" json:"akte_date"`
	Year          uint    `gorm:"type:bigint;not null" json:"year"`
	SKNO1         string  `gorm:"type:varchar(150);not null" json:"sk_no_1"`
	SKNO2         string  `gorm:"type:varchar(150);not null" json:"sk_no_2"`
	SKNO3         string  `gorm:"type:varchar(150);not null" json:"sk_no_3"`
	FileName      string  `gorm:"type:varchar(200)" json:"file_name"`
	FileURL       string  `gorm:"type:varchar(200)" json:"file_url"`
	Remark        string  `gorm:"type:varchar(5000)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
