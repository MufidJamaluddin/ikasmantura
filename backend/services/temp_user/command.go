package temp_user

import (
	"backend/models"
	"backend/repository"
	"backend/viewmodels"
	"gorm.io/gorm"
)

func Verify(db *gorm.DB, id uint, out *viewmodels.UserDto) error {
	var (
		err            error
		tempModel      models.TempUser
		permanentModel models.User
	)

	out.Id = int(id)

	toTempModel(out, &tempModel)

	err = db.Transaction(func(tx *gorm.DB) error {
		var errorTransact error

		tx.Save(&tempModel)

		toPermanentModel(&tempModel, &permanentModel)

		tx.Save(&permanentModel)

		return errorTransact
	})

	if err == nil {
		db.Delete(&tempModel)
	}

	return err
}

func Update(db *gorm.DB, id uint, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.TempUser
	)

	out.Id = int(id)

	toTempModel(out, &model)
	err = repository.Update(db, &model)
	return err
}

func Save(db *gorm.DB, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.TempUser
	)

	toTempModel(out, &model)
	if err = db.Create(&model).Error; err == nil {
		out.Id = int(model.ID)
	}
	return err
}

func Delete(db *gorm.DB, id uint, out *viewmodels.UserDto) error {
	var (
		err   error
		model models.TempUser
	)

	model.ID = id

	if err = repository.Delete(db, &model); err == nil {
		toViewModel(&model, out)
	}
	return err
}
