package model

type AsuransiRekening struct {
	ID               uint    `json:"asuransi_rekening_id"`
	AsuransiID       uint    `json:"asuransi_id"`
	NomorRekening    string  `gorm:"type:varchar(100);not null;" json:"nomor_rekening"`
	AtasNamaRekening string  `gorm:"type:varchar(50);not null;" json:"atas_nama_rekening"`
	BankName         string  `gorm:"type:varchar(100);not null;" json:"bank_name"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;" json:"deleted_at"`
}

type CreateAsuransiRekeningParameter struct {
	AsuransiID       uint    `json:"asuransi_id"`
	NomorRekening    string  `gorm:"type:varchar(100);not null;" json:"nomor_rekening"`
	AtasNamaRekening string  `gorm:"type:varchar(50);not null;" json:"atas_nama_rekening"`
	BankName         string  `gorm:"type:varchar(100);not null;" json:"bank_name"`
	CreatedUserID    uint    `gorm:"type:bigint;not null" json:"created_user_id"`
	UpdatedUserID    uint    `gorm:"type:bigint;not null" json:"updated_user_id"`
	DeletedUserID    uint    `gorm:"type:bigint;typedefault:null" json:"deleted_user_id"`
	CreatedAt        float64 `gorm:"type:double precision;not null" json:"created_at"`
	UpdatedAt        float64 `gorm:"type:double precision;not null" json:"updated_at"`
	DeletedAt        float64 `gorm:"type:double precision;" json:"deleted_at"`
}
