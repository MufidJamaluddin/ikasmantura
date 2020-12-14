package classroom

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"database/sql"
	"gorm.io/gorm"
)

var classroomSearchFields []string

func init() {
	classroomSearchFields = []string{
		"major", "level", "seq",
	}
}

func GetTotal(db *gorm.DB, search *viewmodels.ClassroomParam) (uint, error) {
	var (
		err   error
		model models.Classroom
		tx    *gorm.DB
		total int64
	)

	tx = db.Model(&model)

	search.Filter(tx, classroomSearchFields)

	if err = tx.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint(total), err
}

func Find(db *gorm.DB, search *viewmodels.ClassroomParam, callback func(*viewmodels.ClassroomDto)) error {
	var (
		err   error
		model models.Classroom
		tx    *gorm.DB
		rows  *sql.Rows
	)

	tx = db.Model(&model)
	search.Filter(tx, classroomSearchFields)

	if rows, err = tx.Rows(); err != nil {
		return err
	}

	for rows.Next() {
		err = tx.ScanRows(rows, &model)

		toViewModel(&model, &search.ClassroomDto)
		callback(&search.ClassroomDto)
	}

	return err
}

func FindById(db *gorm.DB, id uint, out *viewmodels.ClassroomDto) error {
	var (
		err   error
		model models.Classroom
	)

	if err = repository.FindById(db, id, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
