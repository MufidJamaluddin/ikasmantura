package album

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var albumSearchFields []string

func init() {
	albumSearchFields = []string{
		"title",
	}
}

func GetTotal(db *gorm.DB, search *viewmodels.AlbumParam) (uint, error) {
	var (
		err   error
		model models.Album
		tx    *gorm.DB
		total int64
	)

	tx = db.Model(&model)

	search.Filter(tx, albumSearchFields, false)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func Find(db *gorm.DB, search *viewmodels.AlbumParam, callback func(*viewmodels.AlbumDto)) error {
	var (
		err   error
		model models.Album
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = db.Model(&model)
	search.Filter(tx, albumSearchFields, true)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		toViewModel(&model, &search.AlbumDto)
		callback(&search.AlbumDto)
	}

	return err
}

func FindById(db *gorm.DB, id uint, out *viewmodels.AlbumDto) error {
	var (
		err   error
		model models.Album
	)

	if err = repository.FindById(db, id, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
