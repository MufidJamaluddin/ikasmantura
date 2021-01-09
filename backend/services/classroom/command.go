package classroom

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, id uint, data *viewmodels.ClassroomDto) error {
	var (
		err   error
		model models.Classroom
	)

	data.Id = int(id)

	toModel(data, &model)

	err = db.Transaction(func(tx *gorm.DB) error {
		return repository.Update(tx, &model)
	})

	return err
}

func Save(db *gorm.DB, out *viewmodels.ClassroomDto) error {
	var (
		err   error
		model models.Classroom
	)

	toModel(out, &model)
	err = db.Transaction(func(tx *gorm.DB) error {
		var errT error
		if errT = repository.Save(tx, &model); err == nil {
			out.Id = int(model.ID)
		}
		return errT
	})
	return err
}

func Delete(db *gorm.DB, id uint, out *viewmodels.ClassroomDto) error {
	var (
		err   error
		model models.Classroom
	)

	model.ID = id

	err = db.Transaction(func(tx *gorm.DB) error {
		var errT error
		if errT = repository.Delete(tx, &model); err == nil {
			toViewModel(&model, out)
		}
		return errT
	})

	return err
}
