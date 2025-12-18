package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type AsuransiRepository interface {
	CountAsuransiAll() (count int64, err error)
	FindAsuransis() (AsuransiOutput []model.Asuransi, err error)
	FindAsuransisOffset(limit int, offset int, order string, dir string) (AsuransiOutput []model.Asuransi, err error)
	SearchAsuransi(limit int, offset int, order string, dir string, search string) (AsuransiOutput []model.Asuransi, err error)
	CountSearchAsuransi(search string) (count int64, err error)
	FindAsuransiById(id uint) (AsuransiOutput model.Asuransi, err error)
	FindExcAsuransi(id uint) (AsuransiOutput []model.Asuransi, err error)
	InsertAsuransi(Asuransi model.Asuransi) (AsuransiOutput model.Asuransi, err error)
	UpdateAsuransi(Asuransi model.Asuransi, id uint) (AsuransiOutput model.Asuransi, err error)
}

type AsuransiConnection struct {
	connection *gorm.DB
}

func NewAsuransiRepository(db *gorm.DB) AsuransiRepository {
	return &AsuransiConnection{
		connection: db,
	}
}
func (db *AsuransiConnection) CountAsuransiAll() (count int64, err error) {
	res := db.connection.Table("asuransis").Where("deleted_at = 0").Count(&count)
	return count, res.Error
}

func (db *AsuransiConnection) FindAsuransis() (AsuransiOutput []model.Asuransi, err error) {
	var (
		Asuransis []model.Asuransi
	)
	res := db.connection.Where("deleted_at = 0").Order("asuransi_name").Find(&Asuransis)
	return Asuransis, res.Error
}

func (db *AsuransiConnection) FindAsuransisOffset(limit int, offset int, order string, dir string) (AsuransiOutput []model.Asuransi, err error) {
	var (
		orderDirection string
		Asuransis      []model.Asuransi
	)
	orderDirection = order + " " + dir
	res := db.connection.Where("deleted_at = 0").Order(orderDirection).Limit(limit).Offset(offset).Find(&Asuransis)
	return Asuransis, res.Error
}

func (db *AsuransiConnection) SearchAsuransi(limit int, offset int, order string, dir string, search string) (AsuransiOutput []model.Asuransi, err error) {
	var (
		orderDirection string
		final          string
		Asuransis      []model.Asuransi
	)
	orderDirection = order + " " + dir
	final = "%" + strings.ToLower(search) + "%"
	res := db.connection.Where("(lower(asuransi_name) LIKE ? OR lower(nomor_hp) LIKE ? OR lower(email) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final, final).Order(orderDirection).Limit(limit).Offset(offset).Find(&Asuransis)
	return Asuransis, res.Error
}

func (db *AsuransiConnection) CountSearchAsuransi(search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("asuransis").Where("(lower(asuransi_name) LIKE ? OR lower(nomor_hp) LIKE ? OR lower(email) LIKE ? OR lower(remark) LIKE ?) AND deleted_at = 0", final, final, final, final).Count(&count)
	return count, res.Error
}

func (db *AsuransiConnection) FindAsuransiById(id uint) (AsuransiOutput model.Asuransi, err error) {
	var (
		Asuransi model.Asuransi
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&Asuransi)
	return Asuransi, res.Error
}

func (db *AsuransiConnection) FindExcAsuransi(id uint) (AsuransiOutput []model.Asuransi, err error) {
	var (
		Asuransis []model.Asuransi
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("asuransi_name").Find(&Asuransis)
	return Asuransis, res.Error
}

func (db *AsuransiConnection) InsertAsuransi(Asuransi model.Asuransi) (AsuransiOutput model.Asuransi, err error) {
	res := db.connection.Save(&Asuransi)
	return Asuransi, res.Error
}

func (db *AsuransiConnection) UpdateAsuransi(Asuransi model.Asuransi, id uint) (AsuransiOutput model.Asuransi, err error) {
	res := db.connection.Where("id=?", id).Updates(&Asuransi)
	return Asuransi, res.Error
}
