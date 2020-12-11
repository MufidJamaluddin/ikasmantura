package user

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var userSearchFields []string

func init() {
	userSearchFields = []string{
		"name", "username", "email", "phone", "is_admin",
	}
}

func GetTotal(db *gorm.DB, search *viewmodels.UserParam) (uint, error) {
	var (
		err   error
		model models.User
		tx    *gorm.DB
		total int64
	)

	tx = db.Model(&model)

	search.Filter(tx, userSearchFields)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func Find(db *gorm.DB, search *viewmodels.UserParam, callback func(*viewmodels.UserDto)) error {
	var (
		err   error
		model models.User
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = db.Model(&model)
	search.Filter(tx, userSearchFields)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		toViewModel(&model, &search.UserDto)
		callback(&search.UserDto)
	}

	return err
}

func FindById(db *gorm.DB, id uint, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.User
	)

	if err = repository.FindById(db, id, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
