package department

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var departmentSearchFields []string

func init() {
	departmentSearchFields = []string{
		"name", "user_id", "type",
	}
}

func GetTotal(db *gorm.DB, search *viewmodels.DepartmentParam) (uint, error) {
	var (
		err   error
		model models.Department
		tx    *gorm.DB
		total int64
	)

	tx = db.Model(&model)
	search.Filter(tx, departmentSearchFields)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func Find(db *gorm.DB, search *viewmodels.DepartmentParam, callback func(*viewmodels.DepartmentDto)) error {
	var (
		err   error
		model models.Department
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = db.Model(&model).Joins("User")
	search.Filter(tx, departmentSearchFields)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		toViewModel(&model, &search.DepartmentDto)
		callback(&search.DepartmentDto)
	}

	return err
}

func FindById(db *gorm.DB, id uint, out *viewmodels.DepartmentDto) error {
	var (
		err   error
		model models.Department
	)

	if err = repository.FindById(db, id, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
