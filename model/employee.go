package model

type Employee struct {
	ID            uint    `json:"employee_id"`
	NIK           string  `gorm:"type:varchar(20);not null" json:"nik"`
	Firstname     string  `gorm:"type:varchar(30);not null" json:"first_name"`
	Lastname      string  `gorm:"type:varchar(30)" json:"last_name"`
	Email         string  `gorm:"type:varchar(30);not null" json:"email"`
	Initials      []byte  `json:"initials"`
	Signature     []byte  `json:"signature"`
	DivisionID    uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	DepartmentID  uint    `gorm:"type:bigint;foreign_key;index:" json:"department_id"`
	SectionID     uint    `gorm:"type:bigint;foreign_key;index:" json:"section_id"`
	PositionID    uint    `gorm:"type:bigint;foreign_key;index:" json:"position_id"`
	LocationID    uint    `gorm:"type:bigint;foreign_key;index:" json:"location_id"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type SelectEmployeeParameter struct {
	ID             uint    `json:"employee_id"`
	NIK            string  `gorm:"type:varchar(20);not null" json:"nik"`
	Firstname      string  `gorm:"type:varchar(30);not null" json:"first_name"`
	Lastname       string  `gorm:"type:varchar(30)" json:"last_name"`
	Email          string  `gorm:"type:varchar(30);not null" json:"email"`
	Initials       []byte  `json:"initials"`
	Signature      []byte  `json:"signature"`
	DivisionID     uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	DivisionName   string  `gorm:"type:varchar(50);not null;unique" json:"division_name"`
	DepartmentID   uint    `gorm:"type:bigint;foreign_key;index:" json:"department_id"`
	DepartmentName string  `gorm:"type:varchar(50);not null;unique" json:"department_name"`
	SectionID      uint    `gorm:"type:bigint;foreign_key;index:" json:"section_id"`
	SectionName    string  `gorm:"type:varchar(50);not null;unique" json:"section_name"`
	PositionID     uint    `gorm:"type:bigint;foreign_key;index:" json:"position_id"`
	PositionName   string  `gorm:"type:varchar(50);not null;unique" json:"position_name"`
	LocationID     uint    `gorm:"type:bigint;foreign_key;index:" json:"location_id"`
	LocationName   string  `gorm:"type:varchar(50);not null;unique" json:"location_name"`
	Remark         string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID  uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID  uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID  uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt      float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt      float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt      float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}

type CreateEmployeeParameter struct {
	NIK           string  `gorm:"type:varchar(20);not null" json:"nik"`
	Firstname     string  `gorm:"type:varchar(30);not null" json:"first_name"`
	Lastname      string  `gorm:"type:varchar(30)" json:"last_name"`
	Email         string  `gorm:"type:varchar(30);not null" json:"email"`
	Initials      []byte  `json:"initials"`
	Signature     []byte  `json:"signature"`
	DivisionID    uint    `gorm:"type:bigint;foreign_key;index:" json:"division_id"`
	DepartmentID  uint    `gorm:"type:bigint;foreign_key;index:" json:"department_id"`
	SectionID     uint    `gorm:"type:bigint;foreign_key;index:" json:"section_id"`
	PositionID    uint    `gorm:"type:bigint;foreign_key;index:" json:"position_id"`
	LocationID    uint    `gorm:"type:bigint;foreign_key;index:" json:"location_id"`
	Remark        string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt     float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt     float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt     float64 `gorm:"type:double precision;typedefault:null" json:"deleted_at"`
}
