package model

type Vehicle struct {
	ID                      uint    `json:"vehicle_id"`
	VehicleHullNumber       string  `gorm:"type:varchar(50);not null" json:"vehicle_hull_number"`
	VehiclePlateName        string  `gorm:"type:varchar(50);not null" json:"vehicle_plate_name"`
	VehicleCategoryID       uint    `json:"vehicle_category_id"`
	VehicleMerkID           uint    `json:"vehicle_merk_id"`
	VehicleTypeID           uint    `json:"vehicle_type_id"`
	VehicleRegistrationName string  `gorm:"type:varchar(100);not null" json:"vehicle_registration_name"`
	NamaPengguna            string  `gorm:"type:varchar(100);" json:"nama_pengguna"`
	ProductionYear          uint    `gorm:"type:bigint;not null" json:"production_year"`
	CompanyID               uint    `json:"company_id"`
	ChassisNumber           string  `gorm:"type:varchar(50);not null" json:"chassis_number"`
	MachineNumber           string  `gorm:"type:varchar(50);not null" json:"machine_number"`
	OriginalBPKBStorage     string  `gorm:"type:varchar(50);not null" json:"original_bpkb_storage"`
	BPKBNumber              string  `gorm:"type:varchar(50);not null" json:"bpkb_number"`
	Price                   float64 `gorm:"type:double precision" json:"price"`
	Flag                    string  `gorm:"type:varchar(50);not null" json:"flag"`
	NamaPTPembiayaan        uint    `json:"nama_pt_pembiayaan"`
	FlagReminder            uint    `json:"flag_reminder"`
	Status                  string  `json:"status"`
	EmailPengguna           string  `gorm:"type:varchar(100);not null" json:"email_pengguna"`
	LocationID              uint    `json:"location_id"`
	Remark                  string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID           uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID           uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID           uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt               float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt               float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt               float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type SelectVehicleParameter struct {
	ID                      uint    `json:"vehicle_id"`
	VehicleHullNumber       string  `gorm:"type:varchar(50);not null" json:"vehicle_hull_number"`
	VehiclePlateName        string  `gorm:"type:varchar(50);not null" json:"vehicle_plate_name"`
	VehicleCategoryID       uint    `json:"vehicle_category_id"`
	VehicleCategoryName     string  `gorm:"type:varchar(50);not null" json:"vehicle_category_name"`
	VehicleMerkID           uint    `json:"vehicle_merk_id"`
	VehicleMerkName         string  `gorm:"type:varchar(50);not null" json:"vehicle_merk_name"`
	VehicleTypeID           uint    `json:"vehicle_type_id"`
	VehicleTypeName         string  `gorm:"type:varchar(50);not null" json:"vehicle_type_name"`
	VehicleRegistrationName string  `gorm:"type:varchar(100);not null" json:"vehicle_registration_name"`
	NamaPengguna            string  `gorm:"type:varchar(100);" json:"nama_pengguna"`
	ProductionYear          uint    `gorm:"type:bigint;not null" json:"production_year"`
	CompanyID               uint    `json:"company_id"`
	CompanyName             string  `gorm:"type:varchar(50);not null" json:"company_name"`
	ChassisNumber           string  `gorm:"type:varchar(50);not null" json:"chassis_number"`
	MachineNumber           string  `gorm:"type:varchar(50);not null" json:"machine_number"`
	OriginalBPKBStorage     string  `gorm:"type:varchar(50);not null" json:"original_bpkb_storage"`
	BPKBNumber              string  `gorm:"type:varchar(50);not null" json:"bpkb_number"`
	Price                   float64 `gorm:"type:double precision" json:"price"`
	Flag                    string  `gorm:"type:varchar(50);not null" json:"flag"`
	NamaPTPembiayaan        string  `json:"nama_pt_pembiayaan"`
	FlagReminder            uint    `json:"flag_reminder"`
	Status                  string  `json:"status"`
	EmailPengguna           string  `gorm:"type:varchar(100);not null" json:"email_pengguna"`
	LocationID              uint    `json:"location_id"`
	LocationName            string  `gorm:"type:varchar(50);not null" json:"location_name"`
	Remark                  string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID           uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	CreatedUser             string  `gorm:"type:varchar(100);" json:"created_user"`
	UpdatedUserID           uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	UpdatedUser             string  `gorm:"type:varchar(100);" json:"updated_user"`
	DeletedUserID           uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	DeletedUser             string  `gorm:"type:varchar(100);" json:"deleted_user"`
	CreatedAt               string  `gorm:"type:varchar(100);not null" json:"created_at"`
	UpdatedAt               string  `gorm:"type:varchar(100);not null" json:"updated_at"`
	DeletedAt               string  `gorm:"type:varchar(100);typedefault:null" json:"deleted_at"`
}

type CreateVehicleParameter struct {
	VehicleHullNumber       string  `gorm:"type:varchar(50);not null" json:"vehicle_hull_number"`
	VehiclePlateName        string  `gorm:"type:varchar(50);not null" json:"vehicle_plate_name"`
	VehicleCategoryID       uint    `json:"vehicle_category_id"`
	VehicleMerkID           uint    `json:"vehicle_merk_id"`
	VehicleTypeID           uint    `json:"vehicle_type_id"`
	VehicleRegistrationName string  `gorm:"type:varchar(100);not null" json:"vehicle_registration_name"`
	NamaPengguna            string  `gorm:"type:varchar(100);" json:"nama_pengguna"`
	ProductionYear          uint    `gorm:"type:bigint;not null" json:"production_year"`
	CompanyID               uint    `json:"company_id"`
	ChassisNumber           string  `gorm:"type:varchar(50);not null" json:"chassis_number"`
	MachineNumber           string  `gorm:"type:varchar(50);not null" json:"machine_number"`
	OriginalBPKBStorage     string  `gorm:"type:varchar(50);not null" json:"original_bpkb_storage"`
	BPKBNumber              string  `gorm:"type:varchar(50);not null" json:"bpkb_number"`
	Price                   float64 `gorm:"type:double precision" json:"price"`
	Flag                    string  `gorm:"type:varchar(50);not null" json:"flag"`
	NamaPTPembiayaan        uint    `json:"nama_pt_pembiayaan"`
	FlagReminder            uint    `json:"flag_reminder"`
	Status                  string  `json:"status"`
	EmailPengguna           string  `gorm:"type:varchar(100);not null" json:"email_pengguna"`
	LocationID              uint    `json:"location_id"`
	Remark                  string  `gorm:"type:varchar(200)" json:"remark"`
	CreatedUserID           uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID           uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID           uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt               float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt               float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt               float64 `gorm:"type:double precision;" json:"deleted_at"`
}
