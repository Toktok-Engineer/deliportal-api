package repository

import (
	"deliportal-api/model"
	"strings"

	"gorm.io/gorm"
)

type AsuransiRekeningRepository interface {
	CountAsuransiRekeningAll(AsuransiID int) (count int64, err error)
	FindAsuransiRekenings() (AsuransiRekeningOutput []model.AsuransiRekening, err error)
	FindAsuransiRekeningsOffset(AsuransiID int, limit int, offset int, order string, dir string) (AsuransiRekeningOutput []model.AsuransiRekening, err error)
	SearchAsuransiRekening(AsuransiID int, limit int, offset int, order string, dir string, search string) (AsuransiRekeningOutput []model.AsuransiRekening, err error)
	CountSearchAsuransiRekening(AsuransiID int, search string) (count int64, err error)
	FindAsuransiRekeningById(id uint) (AsuransiRekeningOutput model.AsuransiRekening, err error)
	FindExcAsuransiRekening(id uint) (AsuransiRekeningOutput []model.AsuransiRekening, err error)
	InsertAsuransiRekening(AsuransiRekening model.AsuransiRekening) (AsuransiRekeningOutput model.AsuransiRekening, err error)
	UpdateAsuransiRekening(AsuransiRekening model.AsuransiRekening, id uint) (AsuransiRekeningOutput model.AsuransiRekening, err error)
}

type AsuransiRekeningConnection struct {
	connection *gorm.DB
}

func NewAsuransiRekeningRepository(db *gorm.DB) AsuransiRekeningRepository {
	return &AsuransiRekeningConnection{
		connection: db,
	}
}
func (db *AsuransiRekeningConnection) CountAsuransiRekeningAll(AsuransiID int) (count int64, err error) {
	res := db.connection.Table("asuransi_rekenings").Where("deleted_at = 0 and asuransi_id = ?", AsuransiID).Count(&count)
	return count, res.Error
}

func (db *AsuransiRekeningConnection) FindAsuransiRekenings() (AsuransiRekeningOutput []model.AsuransiRekening, err error) {
	var (
		AsuransiRekenings []model.AsuransiRekening
	)
	res := db.connection.Where("deleted_at = 0").Order("asuransi_id").Find(&AsuransiRekenings)
	return AsuransiRekenings, res.Error
}

func (db *AsuransiRekeningConnection) FindAsuransiRekeningsOffset(AsuransiID int, limit int, offset int, order string, dir string) (AsuransiRekeningOutput []model.AsuransiRekening, err error) {
	var (
		orderDirection    string
		AsuransiRekenings []model.AsuransiRekening
	)
	orderDirection = order + " " + dir
	res := db.connection.Table("asuransi_rekenings").Select("asuransi_rekenings.id, asuransis.asuransi_name, asuransi_rekenings.nomor_rekening, asuransi_rekenings.atas_nama_rekening, asuransi_rekenings.bank_name").Joins("left join asuransis on asuransi_rekenings.asuransi_id = asuransis.id").Where("asuransi_rekenings.asuransi_id = ? AND asuransi_rekenings.deleted_at = 0", AsuransiID).Order(orderDirection).Limit(limit).Offset(offset).Find(&AsuransiRekenings)
	return AsuransiRekenings, res.Error
}

func (db *AsuransiRekeningConnection) SearchAsuransiRekening(AsuransiID int, limit int, offset int, order string, dir string, search string) (AsuransiRekeningOutput []model.AsuransiRekening, err error) {
	var (
		orderDirection    string
		final             string
		AsuransiRekenings []model.AsuransiRekening
	)
	final = "%" + strings.ToLower(search) + "%"
	orderDirection = order + " " + dir
	res := db.connection.Table("asuransi_rekenings").Select("nomor_rekening, atas_nama_rekening, bank_name").Where("(lower(nomor_rekening) LIKE ? OR lower(atas_nama_rekening) LIKE ? OR lower(bank_name) LIKE ? ) AND deleted_at = 0 AND asuransi_id = ?", final, final, final, AsuransiID).Order(orderDirection).Limit(limit).Offset(offset).Find(&AsuransiRekenings)
	return AsuransiRekenings, res.Error
}

func (db *AsuransiRekeningConnection) CountSearchAsuransiRekening(AsuransiID int, search string) (count int64, err error) {
	final := "%" + strings.ToLower(search) + "%"
	res := db.connection.Table("asuransi_rekenings").Select("nomor_rekening, atas_nama_rekening, bank_name").Where("(lower(nomor_rekening) LIKE ? OR lower(atas_nama_rekening) LIKE ? OR lower(bank_name) LIKE ? ) AND deleted_at = 0 AND asuransi_id = ?", final, final, final, AsuransiID).Count(&count)
	return count, res.Error
}

func (db *AsuransiRekeningConnection) FindAsuransiRekeningById(id uint) (AsuransiRekeningOutput model.AsuransiRekening, err error) {
	var (
		AsuransiRekening model.AsuransiRekening
	)
	res := db.connection.Where("id=? AND deleted_at = 0", id).Take(&AsuransiRekening)
	return AsuransiRekening, res.Error
}

func (db *AsuransiRekeningConnection) FindExcAsuransiRekening(id uint) (AsuransiRekeningOutput []model.AsuransiRekening, err error) {
	var (
		AsuransiRekenings []model.AsuransiRekening
	)
	res := db.connection.Where("id!=? AND deleted_at = 0", id).Order("asuransi_id").Find(&AsuransiRekenings)
	return AsuransiRekenings, res.Error
}

func (db *AsuransiRekeningConnection) InsertAsuransiRekening(AsuransiRekening model.AsuransiRekening) (AsuransiRekeningOutput model.AsuransiRekening, err error) {
	res := db.connection.Save(&AsuransiRekening)
	return AsuransiRekening, res.Error
}

func (db *AsuransiRekeningConnection) UpdateAsuransiRekening(AsuransiRekening model.AsuransiRekening, id uint) (AsuransiRekeningOutput model.AsuransiRekening, err error) {
	res := db.connection.Where("id=?", id).Updates(&AsuransiRekening)
	return AsuransiRekening, res.Error
}
