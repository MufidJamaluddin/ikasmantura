package temp_user

import (
	"backend/models"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
	"strings"
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
		model models.TempUser
		tx    *gorm.DB
		total int64
	)

	tx = db.Model(&model)

	search.Filter(tx, userSearchFields, false)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func Find(db *gorm.DB, search *viewmodels.UserParam, callback func(*viewmodels.UserDto)) error {
	var (
		err   error
		model models.TempUser
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = db.Model(&model)
	search.Filter(tx, userSearchFields, true)

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
		model models.TempUser
		tx    *gorm.DB
	)

	tx = db.Model(&model).
		Preload("Address").
		Preload("Classrooms")

	if err = tx.First(&model, "id = ?", id).Error; err == nil {
		toViewModel(&model, out)
	}
	return err
}

func IsUsernameOrEmailAvailable(
	db *gorm.DB,
	user *viewmodels.UserAvailabilityDto,
	response *viewmodels.UserAvailabilityResponseDto) error {

	var err error

	user.Username = strings.Trim(user.Username, " ")
	user.Email = strings.Trim(user.Email, " ")

	if user.Username == "" {
		user.Username = "-"
	}

	if user.Email == "" {
		user.Email = "-"
	}

	err = db.Raw(
		"SELECT EXISTS(SELECT 1 FROM users WHERE username = @username OR email = @email) "+
			"OR EXISTS(SELECT 1 FROM temp_users WHERE username = @username OR email = @email) "+
			"AS exist",
		sql.Named("username", user.Username),
		sql.Named("email", user.Email),
	).Scan(response).Error

	return err
}
