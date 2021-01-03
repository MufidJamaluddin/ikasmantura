package albumphoto

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var albumPhotoSearchFields []string

func init() {
	albumPhotoSearchFields = []string{
		"title", "album_id",
	}
}

func GetTotal(db *gorm.DB, search *viewmodels.AlbumPhotoParam) (uint, error) {
	var (
		err   error
		model models.AlbumPhoto
		tx    *gorm.DB
		total int64
	)

	tx = db.Model(&model)

	search.Filter(tx, albumPhotoSearchFields, false)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func Find(db *gorm.DB, search *viewmodels.AlbumPhotoParam, callback func(*viewmodels.AlbumPhotoDto)) error {
	var (
		err   error
		model models.AlbumPhoto
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = db.Model(&model)
	search.Filter(tx, albumPhotoSearchFields, true)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		toViewModel(&model, &search.AlbumPhotoDto)

		callback(&search.AlbumPhotoDto)
	}

	return err
}

func FindById(db *gorm.DB, id uint, out *viewmodels.AlbumPhotoDto) error {
	var (
		err   error
		model models.AlbumPhoto
	)

	if err = repository.FindById(db, id, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
